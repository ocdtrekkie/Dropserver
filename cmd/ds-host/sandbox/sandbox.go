package sandbox

// Keep this module but make significant changes to
// have it just manage deno processes?

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/teleclimber/DropServer/cmd/ds-host/domain"
	"github.com/teleclimber/DropServer/cmd/ds-host/record"
	"github.com/teleclimber/DropServer/internal/twine"
	"golang.org/x/sys/unix"
)

// OK so this is going to be significantly different.
// ..we may even need to delete entirely and create a fresh new module.
// - create sandbox on-demand on incoming request (this is a limitation from Deno, for now)
// - need to detect when it crashes/hangs
// - need to be able to send a die message
// - need to be able to kill if misbehaving
// - We need a appspace-api server on host for all api requests that emanate from appspaces.

type appSpaceSession struct {
	tasks      []*Task
	lastActive time.Time
	tiedUp     bool
}

// Task tracks the container being tied up for one request
type Task struct {
	Finished bool //build up with start time, elapsed and any other metadata
}

const sandboxService = 11
const executeService = 12

// MigrateService is for appspace migration
const MigrateService = 13 // hmm...

const routesService = 14

const execFnCommand = 11

// Sandbox holds the data necessary to interact with the container
type Sandbox struct {
	id              int // getter only (const), unexported
	appVersion      *domain.AppVersion
	appspace        *domain.Appspace
	status          domain.SandboxStatus // getter/setter, so make it unexported.
	socketsDir      string
	cmd             *exec.Cmd
	twine           *twine.Twine
	services        *domain.ReverseServices
	statusMux       sync.Mutex
	statusSub       map[domain.SandboxStatus][]chan domain.SandboxStatus
	transport       http.RoundTripper
	appSpaceSession appSpaceSession // put a getter for that?
	killScore       float64         // this should not be here.
	Config          *domain.RuntimeConfig
}

// NewSandbox creates a new sandbox with the passed parameters
func NewSandbox(sandboxID int, services *domain.ReverseServices, config *domain.RuntimeConfig) *Sandbox {
	newSandbox := &Sandbox{ // <-- this really needs a maker fn of some sort??
		id:        sandboxID,
		services:  services,
		status:    domain.SandboxStarting,
		statusSub: make(map[domain.SandboxStatus][]chan domain.SandboxStatus),
		Config:    config}

	return newSandbox
}

// Start Should start() return a channel or something?
// or should callers just do go start()?
func (s *Sandbox) Start(appVersion *domain.AppVersion, appspace *domain.Appspace) error { // TODO: return an error, presumably?
	s.getLogger("Start()").Debug("starting...")

	s.appVersion = appVersion
	s.appspace = appspace

	logger := s.getLogger("Start()")

	socketsDir, err := makeSocketsDir(s.Config.Sandbox.SocketsDir, appspace.AppspaceID)
	if err != nil {
		s.getLogger(fmt.Sprintf("Start(), makeSocketsDir() dir: %v", s.Config.Sandbox.SocketsDir)).Error(err)
		return err
	}
	s.socketsDir = socketsDir

	err = s.writeImportMap()
	if err != nil {
		return err
	}

	//fwdSock := path.Join(socketsDir, "fwd.sock") // forward socket is where the ds runner will create the sandbox server

	// Here start should take necessary data about appspace
	// ..in order to pass in the right permissions to deno.

	twineServer, err := twine.NewServer(path.Join(socketsDir, "rev.sock"))
	if err != nil {
		logger.AddNote("twine.NewServer").Error(err)
		return err // maybe return a user-centered error
	}
	s.twine = twineServer

	err = os.Setenv("NO_COLOR", "true")
	if err != nil {
		logger.AddNote("os.Setenv").Error(err)
		return err // return user-centered error
	}

	// Give deno's current sandboxing ideas,
	// Probably need to think more about flags we pass, such as --no-remote?
	cmd := exec.Command(
		"deno",
		"run",
		"--unstable", // needed for unix domain sockets
		"--importmap="+s.getImportPathFile(),
		"--allow-read",  // TODO app dir and appspace dir, and sockets
		"--allow-write", // TODO appspace dir, sockets
		s.Config.Exec.SandboxRunnerPath,
		s.socketsDir,
		filepath.Join(s.Config.Exec.AppsPath, appVersion.LocationKey), // while we have an import-map, these are stil needed to read files without importing
		filepath.Join(s.Config.Exec.AppspacesFilesPath, appspace.LocationKey),
	)
	s.cmd = cmd
	// Note that ultimately we need to stick this in a Cgroup

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logger.AddNote("cmd.StdoutPipe()").Error(err)
		return err // user centered error
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		logger.AddNote("cmd.StderrPipe()").Error(err)
		return err
	}

	err = cmd.Start() // returns right away
	if err != nil {
		logger.AddNote("cmd.Start()").Error(err)
		return err
	}

	go s.monitor(stdout, stderr)

	_, ok := <-s.twine.ReadyChan
	if !ok {
		logger.Log("Apparent failed start. ReadyChan closed")
		s.Stop()
		return errors.New("Failed to start sandbox")
	}

	go s.listenMessages()

	s.transport = http.DefaultTransport // really not sure what this means or what it's for anymore....
	s.SetStatus(domain.SandboxReady)

	return nil
}

// monitor waits for cmd to end or an error gets sent
// It also collects output for use somewhere.
func (s *Sandbox) monitor(stdout io.ReadCloser, stderr io.ReadCloser) {

	go func() {
		for { // you need to be in a loop to keep the channel "flowing"
			err := <-s.twine.ErrorChan
			if err != nil {
				s.getLogger("ErrorChan").Error(err)
				// We may want to stash a message on s. to enlighten user as to what happened?
				s.Stop()
			} else {
				break // errorChan was closed, so exit loop
			}
		}
	}()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		printLogs(stdout)
	}()

	go func() {
		defer wg.Done()
		printLogs(stderr)
	}()

	//wg.Wait()

	err := s.cmd.Wait()
	if err != nil {
		// TODO check error type (see Wait comment)
		s.getLogger("monitor(), s.cmd.Wait()").Error(err)
		s.Stop()
		// Here we should kill off reverse listener?
		// This is where we want to log things for the benefit of the dropapp user.
		// Also somehow whoever started the sandbox needs to know it exited with error
	}

	wg.Wait()

	s.SetStatus(domain.SandboxDead)
	// -> it may have crashed.
	// TODO ..should probably call regular shutdown procdure to clean everything up.

	// now kill the reverse channel? Otherwise we risk killing it shile it could still provide valuable data?
}

func printLogs(r io.ReadCloser) {
	buf := make([]byte, 80)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			fmt.Printf("%s", buf[0:n]) // OK not sure about these. IS this effectively console.log output?
		}
		if err != nil {
			break
		}
	}
}

func (s *Sandbox) listenMessages() {
	for message := range s.twine.MessageChan {
		switch message.ServiceID() {
		case routesService:
			s.services.Routes.Command(s.appspace, message)
		default:
			s.getLogger("listenMessages()").Log(fmt.Sprintf("Service not recognized: %v", message.ServiceID()))
			message.SendError("service not recognized")
		}
	}
}

// Stop stops the sandbox and its associated open connections
func (s *Sandbox) Stop() {
	// reverse listener...

	// get state and then send kill signal?
	// Then loop over and check pid.
	// -> I think ds-sandbox-d had this nailed down.
	// update status and clean up after?

	// if sandbox status is killed, then do nothing, the killing system is working.

	// get status from pid, if running, send kill sig, then wait some.
	// follow up as with ds-sandbox-d

	// how do we avoid getting into a dysfunction acondition with Proxy?
	// Proxy should probably lock the mutex when it creates a task,
	// ..so that task count can be considered acurate in sandox manager's killing function

	s.SetStatus(domain.SandboxKilling)

	// TODO send a command on sandboxService to shutdown.

	err := s.kill(false)
	if err != nil {
		// force kill
		err = s.kill(true)
		if err != nil {
			// ???
		}
	}

	s.twine.Stop()
	// the shutdown command sent to via twine should automatically call graceful shutdown on this side.
	// But still ned to handle the case that the Twine server never got going because it never got "hi" because client died prior to sending it.

	// after you kill, whether successful or not,
	// sandbox manager ought to remove the sandbox from sandboxes.
	// If had to forcekill then quarantine the

	// Not sure if we can do this now, or have to wait til dead...
	err = cleanSocketsDir(s.socketsDir)
	if err != nil {
		s.getLogger("Stop(), cleanSocketDir()").Error(err)
	}
}

func (s *Sandbox) pidAlive() bool {
	process := s.cmd.Process

	if process == nil {
		return false
	}

	// what does proces look like after the underlying process has dies?

	err := process.Signal(syscall.Signal(0))
	if err == nil {
		return true
	}
	return false
}

// kill sandbox, which means send it the kill sig
// This should get picked up internally and it should shut itself down.
func (s *Sandbox) kill(force bool) error {
	process := s.cmd.Process

	if process == nil {
		return nil
	}

	sig := unix.SIGTERM
	if force {
		sig = unix.SIGKILL
	}
	err := process.Signal(sig)
	if err != nil {
		s.getLogger("kill()").AddNote(fmt.Sprintf("process.Signal(sig). Force: %t", force)).Error(err)
	}

	var alive bool
	ms := 5
	if force {
		ms = 50
	}
	for i := 1; i < 21; i++ {
		time.Sleep(time.Duration(ms) * time.Millisecond)

		alive = s.pidAlive()
		if !alive {
			break
		}
	}

	if alive {
		return errors.New("sandbox failed to terminate") // is this a sentinel error?
	}
	return nil
}

// basic getters

// ID returns the ID of the sandbox
func (s *Sandbox) ID() int {
	return s.id
}

// Status returns the status of the Sandbox
func (s *Sandbox) Status() domain.SandboxStatus {
	s.statusMux.Lock()
	defer s.statusMux.Unlock()
	return s.status
}

// ExecFn executes a function in athesandbox, based on the params of AppspaceRouteHandler
func (s *Sandbox) ExecFn(handler domain.AppspaceRouteHandler) error {
	// check status of sandbox first?
	taskCh := s.TaskBegin()
	defer func() {
		taskCh <- true
	}()

	// need to change file path so it can be resolved from inside sandbox.
	// for now assemble an absolute path
	//handler.File = path.Join(s.Config.Exec.AppsPath, s.appVersion.LocationKey, handler.File)

	payload, err := json.Marshal(handler)
	if err != nil {
		// this is an input error. The caller is at fault probably. Don't log.
		return err //return a "bad input" error, and wrap it?
	}

	sent, err := s.twine.Send(executeService, execFnCommand, &payload)
	if err != nil {
		s.getLogger("ExecFn(), s.twine.Send()").Error(err)
		return errors.New("Error sending to sandbox")
	}

	// maybe receive interim progress messages?

	reply, err := sent.WaitReply()
	if err != nil {
		s.getLogger("ExecFn(), sent.WaitReply()").Error(err)
		return err
	}

	if !reply.OK() {
		return reply.Error()
	}

	return nil
}

// SendMessage sends to sandbox via twine
func (s *Sandbox) SendMessage(serviceID int, commandID int, payload *[]byte) (twine.SentMessageI, error) {
	sent, err := s.twine.Send(serviceID, commandID, payload)
	if err != nil {
		return nil, err
	}
	return sent, nil
}

// GetTransport gets the http transport of the sandbox
func (s *Sandbox) GetTransport() http.RoundTripper {
	return s.transport
}

// getLogger retuns the Logging client
func (s *Sandbox) getLogger(note string) *record.DsLogger {
	l := record.NewDsLogger().AddNote("Sandbox")
	if s.appVersion != nil {
		l.AppID(s.appVersion.AppID).AppVersion(s.appVersion.Version)
	}
	if s.appspace != nil {
		l.AppspaceID(s.appspace.AppspaceID)
	}
	if note != "" {
		l.AddNote(note)
	}
	return l
}

// TiedUp returns the appspaceSession
func (s *Sandbox) TiedUp() bool {
	return s.appSpaceSession.tiedUp
}

// LastActive returns the last used time
func (s *Sandbox) LastActive() time.Time {
	return s.appSpaceSession.lastActive
}

// TaskBegin adds a new task to session tasks and returns it
func (s *Sandbox) TaskBegin() chan bool {
	reqTask := Task{}
	s.appSpaceSession.tasks = append(s.appSpaceSession.tasks, &reqTask)
	s.appSpaceSession.lastActive = time.Now()
	s.appSpaceSession.tiedUp = true

	ch := make(chan bool)

	// go func here that blocks on chanel.
	go func() {
		<-ch
		reqTask.Finished = true
		s.appSpaceSession.lastActive = time.Now()
		s.appSpaceSession.tiedUp = s.isTiedUp()
	}()

	return ch //instead of returning a task we should return a chanel.
}

func (s *Sandbox) isTiedUp() (tiedUp bool) {
	for _, task := range s.appSpaceSession.tasks {
		if !task.Finished {
			tiedUp = true
			break
		}
	}
	return
}

// SetStatus sets the status
func (s *Sandbox) SetStatus(status domain.SandboxStatus) {
	s.statusMux.Lock()
	defer s.statusMux.Unlock()

	if status > s.status {
		s.status = status
		for stat, subs := range s.statusSub {
			if stat <= s.status {
				for _, sub := range subs {
					sub <- s.status // this might block if nobody is actually waiting on the channel?
				}
				delete(s.statusSub, stat)
			}
		}
	}
}

// WaitFor blocks until status is met, or greater status is met
func (s *Sandbox) WaitFor(status domain.SandboxStatus) {
	s.statusMux.Lock()
	if s.status >= status {
		s.statusMux.Unlock()
		return
	}
	fmt.Println(s.id, "waiting for sandbox status", status)

	if _, ok := s.statusSub[status]; !ok {
		s.statusSub[status] = []chan domain.SandboxStatus{}
	}
	statusMet := make(chan domain.SandboxStatus)
	s.statusSub[status] = append(s.statusSub[status], statusMet)

	s.statusMux.Unlock()

	<-statusMet
}

// ImportPaths defines a type for creating imopsts.json for Deno
type ImportPaths struct {
	Imports map[string]string `json:"imports"`
}

func (s *Sandbox) makeImportMap() (*[]byte, error) {
	appPath := trailingSlash(filepath.Join(s.Config.Exec.AppsPath, s.appVersion.LocationKey))
	appspacePath := trailingSlash(filepath.Join(s.Config.Exec.AppspacesFilesPath, s.appspace.LocationKey))
	dropserverPath := trailingSlash(s.Config.Exec.SandboxCodePath)
	// TODO: check that none of these paths are "/" as this can defeat protection against forbidden imports.
	im := ImportPaths{
		Imports: map[string]string{
			"/":            "undefined:", // Defeat imports from outside the app dir. See:
			"./":           "./",         // https://github.com/denoland/deno/issues/6294#issuecomment-663256029
			"@app/":        appPath,
			"@appspace/":   appspacePath,
			"@dropserver/": dropserverPath,
			appPath:        appPath,
			appspacePath:   appspacePath,
			dropserverPath: dropserverPath,
		}}

	j, err := json.Marshal(im)
	if err != nil {
		s.getLogger("makeImportMap()").Error(err)
		return nil, err
	}

	return &j, nil
}
func trailingSlash(p string) string {
	if strings.HasSuffix(p, string(os.PathSeparator)) {
		return p
	}
	return p + string(os.PathSeparator)
}

func (s *Sandbox) writeImportMap() error {
	data, err := s.makeImportMap()
	if err != nil {
		return err
	}

	// this should be taken care of externally on appspace install probably.
	err = os.MkdirAll(s.getAppspaceMetaPath(), 0700)
	if err != nil {
		s.getLogger("writeImportMap()").AddNote("os.MkdirAll dir: " + s.getAppspaceMetaPath()).Error(err)
		return err
	}

	err = ioutil.WriteFile(s.getImportPathFile(), *data, 0600)
	if err != nil {
		s.getLogger("writeImportMap()").AddNote("ioutil.WriteFile file: " + s.getImportPathFile()).Error(err)
		return err
	}

	return nil
}
func (s *Sandbox) getImportPathFile() string {
	return filepath.Join(s.getAppspaceMetaPath(), "import-paths.json")
}

// this really needs to be extracted out!
func (s *Sandbox) getAppspaceMetaPath() string {
	return filepath.Join(s.Config.Exec.AppspacesMetaPath, fmt.Sprintf("appspace-%v", s.appspace.AppspaceID))
}

/////////////
func cleanSocketsDir(sockDir string) error {
	if err := os.RemoveAll(sockDir); err != nil {
		return err
	}
	return nil
}
func makeSocketsDir(baseDir string, appspaceID domain.AppspaceID) (string, error) {
	sockDir := getSocketsDir(baseDir, appspaceID)

	err := cleanSocketsDir(sockDir)
	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(sockDir, 0700); err != nil {
		return "", err
	}

	return sockDir, nil
}

func getSocketsDir(baseDir string, appspaceID domain.AppspaceID) string {
	return path.Join(baseDir, fmt.Sprintf("as-%d", appspaceID))
}
