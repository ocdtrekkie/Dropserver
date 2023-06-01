package appfilesmodel

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/teleclimber/DropServer/cmd/ds-host/domain"
	"github.com/teleclimber/DropServer/cmd/ds-host/record"
)

// performs operations on application files

// AppFilesModel is struct for application files manager
type AppFilesModel struct {
	AppLocation2Path interface {
		Base(string) string
		Meta(string) string
		Files(string) string
	} `checkinject:"required"`
	Config *domain.RuntimeConfig `checkinject:"required"`
}

// SavePackage creates a location key and saves the package file under /package
// because we want to save it as-is for later download.
func (a *AppFilesModel) SavePackage(r io.Reader) (string, error) { // this should be areader not bytes
	logger := a.getLogger("SavePackage()")

	locationKey, err := a.createLocation()
	if err != nil {
		return "", err
	}

	f, err := os.OpenFile(a.getPackagePath(locationKey), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		logger.AddNote("OpenFile").Error(err)
		return locationKey, err
	}
	defer f.Close()
	_, err = io.Copy(f, r)
	if err != nil {
		logger.AddNote("io.Copy").Error(err)
		return locationKey, fmt.Errorf("internal error saving package: %w", err)
	}

	return locationKey, nil
}

func (a *AppFilesModel) createLocation() (string, error) {
	p := a.Config.Exec.AppsPath

	err := os.MkdirAll(p, 0766)
	if err != nil {
		a.getLogger("createLocation()").AddNote("os.MkdirAll() apps").Error(err)
		return "", fmt.Errorf("internal error creating location: %w", err)
	}

	p, err = ioutil.TempDir(p, "app")
	if err != nil {
		a.getLogger("createLocation()").AddNote("ioutil.TempDir()").Error(err)
		return "", fmt.Errorf("internal error creating location: %w", err)
	}

	// maybe assert that the L2P returns the path we created?

	return filepath.Base(p), nil
}

// ExtractPackage expands the package contents.
func (a *AppFilesModel) ExtractPackage(locationKey string) error {
	// should ensure that the directory is empty before?

	logger := a.getLogger("SavePackage()")

	packageFD, err := os.Open(a.getPackagePath(locationKey))
	if err != nil {
		logger.AddNote("os.Open package").Error(err)
		return err
	}
	defer packageFD.Close()

	appFilesPath := a.AppLocation2Path.Files(locationKey)
	err = os.MkdirAll(appFilesPath, 0766)
	if err != nil {
		logger.AddNote("os.MkDirAll appFilesPath").Error(err)
		return err
	}

	err = ExtractPackageLow(packageFD, appFilesPath)
	if err != nil {
		logger.AddNote("extractPackage").Error(err)
		return err
	}
	return nil
}

func ExtractPackageLow(r io.Reader, appFilesPath string) error {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return err
		}
		name := filepath.Clean(hdr.Name)
		if filepath.IsAbs(name) {
			return errors.New("invalid (absolute) path in tar header: " + name)
		}
		if strings.Contains(name, "..") {
			return errors.New("invalid (outside) path in tar header: " + name)
		}
		p := filepath.Join(appFilesPath, name)
		err = os.MkdirAll(filepath.Dir(p), 0766)
		if err != nil {
			return fmt.Errorf("failed to create dir: %s: %w", filepath.Dir(p), err)
		}
		fd, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
		if os.IsExist(err) {
			return errors.New("file already exists: " + name)
		}
		if err != nil {
			return fmt.Errorf("failed to create file: %s: %w", name, err)
		}
		_, err = io.Copy(fd, tr)
		fd.Close()
		if err != nil {
			return fmt.Errorf("failed to write file: %s: %w", name, err)
		}
	}
	return nil
}

func (a *AppFilesModel) getPackagePath(locationKey string) string {
	return filepath.Join(a.AppLocation2Path.Base(locationKey), "package.tar.gz")
}

// ReadManifest reads metadata from the files at location key
// this just reads the JSON file. OK.
// In next iteration this will presumably read teh normalized JSON that was evaluated prior.
// OR, this is the original file from the package? And we store the augmented / normalized version in the DB?
// ^^ This is correct. The app is stored unmodified, and the normalized/processed/evaluated data is stashed in the DB.
func (a *AppFilesModel) ReadManifest(locationKey string) (domain.AppVersionManifest, error) {
	jsonPath := filepath.Join(a.AppLocation2Path.Files(locationKey), "dropapp.json")
	jsonBytes, err := os.ReadFile(jsonPath)
	if err != nil {
		// here the error might be that dropapp.json is not in app?
		// Or it could be a more internal problem, like directory of apps not where it's expected to be.
		// Or it could be a bad location key, like it was deleted but DB doesn't know.
		if !a.locationKeyExists(locationKey) {
			a.getLogger(fmt.Sprintf("ReadManifest(), location key: %v", locationKey)).Error(err)
			return domain.AppVersionManifest{}, errors.New("internal error reading app meta data")
		}
		return domain.AppVersionManifest{}, domain.ErrAppManifestNotFound // rename to app manifest
	}

	meta, err := unmarshalManifest(jsonBytes)
	if err != nil {
		a.getLogger(fmt.Sprintf("ReadManifest(), location key: %v, unmarshalManifest", locationKey)).Error(err)
		return domain.AppVersionManifest{}, err
	}

	return meta, nil
}

func (a *AppFilesModel) WriteRoutes(locationKey string, routesData []byte) error {
	routesFile := filepath.Join(a.AppLocation2Path.Meta(locationKey), "routes.json")
	err := ioutil.WriteFile(routesFile, routesData, 0644) // TODO: correct permissions?
	if err != nil {
		a.getLogger(fmt.Sprintf("WriteRoutes(), location key: %v", locationKey)).Error(err)
		return err
	}
	return nil
}

func (a *AppFilesModel) ReadRoutes(locationKey string) ([]byte, error) {
	routesFile := filepath.Join(a.AppLocation2Path.Meta(locationKey), "routes.json")
	routesData, err := ioutil.ReadFile(routesFile)
	if err != nil {
		a.getLogger(fmt.Sprintf("ReadRoutes(), location key: %v", locationKey)).Error(err)
		return nil, err
	}
	return routesData, nil
}

func (a *AppFilesModel) WriteAppIconLink(locationKey string, iconPath string) error {
	// first remove it
	linkPath := filepath.Join(a.AppLocation2Path.Meta(locationKey), "app-icon")
	os.Remove(linkPath)
	if iconPath == "" {
		return nil
	}
	iconPath = filepath.Join(a.AppLocation2Path.Files(locationKey), iconPath)
	err := os.Link(iconPath, linkPath)
	if err != nil {
		a.getLogger("WriteAppIconLink, os.Link").AddNote(iconPath).AddNote(linkPath).Error(err)
		return err
	}
	return nil
}

func (a *AppFilesModel) GetAppIconPath(locationKey string) string {
	linkPath := filepath.Join(a.AppLocation2Path.Meta(locationKey), "app-icon")
	_, err := os.Stat(linkPath)
	if err == os.ErrNotExist {
		return ""
	}
	if err != nil {
		a.getLogger("GetAppIconPath, os.Stat").AddNote(linkPath).Error(err)
		return ""
	}
	return linkPath
}

// Delete removes the files from the system
func (a *AppFilesModel) Delete(locationKey string) error {
	if locationKey == "" {
		e := errors.New("empty string is not a valid location key")
		a.getLogger("Delete()").Error(e)
		return e
	}
	if !a.locationKeyExists(locationKey) {
		return nil //is that an error or do we consider this OK?
	}
	err := os.RemoveAll(a.AppLocation2Path.Base(locationKey))
	if err != nil {
		a.getLogger("Delete()").Error(err)
		return err
	}

	return nil
}

func unmarshalManifest(b []byte) (domain.AppVersionManifest, error) {
	var manifest domain.AppVersionManifest

	err := json.Unmarshal(b, &manifest)
	if err != nil {
		return manifest, err
	}

	return manifest, nil
}

func (a *AppFilesModel) locationKeyExists(locationKey string) bool {
	_, err := os.Stat(a.AppLocation2Path.Base(locationKey))
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true // OK but there could be aonther problem, like permissions out of whack?
	// Should probably log that as warning at least.
}

func (a *AppFilesModel) getLogger(note string) *record.DsLogger {
	r := record.NewDsLogger().AddNote("AppFilesModel")
	if note != "" {
		r.AddNote(note)
	}
	return r
}

// pathInsidePath determines if A path is inside (contained within) path B
func pathInsidePath(p, root string) (bool, error) {
	rel, err := filepath.Rel(root, p)
	if err != nil {
		return false, err // not clear that is an error that is actually an error. Errors I think just mean "can't be inside"
	}

	if strings.Contains(rel, "..") {
		return false, nil
	}
	return true, nil
}
