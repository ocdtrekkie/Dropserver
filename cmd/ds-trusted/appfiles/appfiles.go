package appfiles

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/teleclimber/DropServer/cmd/ds-host/domain"
	"github.com/teleclimber/DropServer/internal/dserror"
)

// performa operations on application files
// This might be a genenric batch file module
// ..so that it can be used for app-space files too.

// Gah should we inject a filesystem abstraction, or is it not worth it?
// Probably inject a config with file locations, and use that to mock files?

var appsPath = "/data/apps"

// AppFiles is struct for application files manager
type AppFiles struct {
	// maybe a config?
	// Logger
}

// Save puts the data passed in files in an apps directory
func (a *AppFiles) Save(files *domain.TrustedSaveAppFiles) (string, domain.Error) {

	dir, err := ioutil.TempDir(appsPath, "app")
	if err != nil {
		// TODO: log and return internal error?
		return "", dserror.New(dserror.InternalError)
	}

	for f, data := range *files.Files {
		f = filepath.Join(dir, f)
		err = ioutil.WriteFile(f, data, 0666) // TODO: permissions?
		if err != nil {
			// TODO: log and return internal error?
			return "", dserror.New(dserror.InternalError, err.Error())
		}
	}

	return filepath.Base(dir), nil
}

// ReadMeta reads metadata from the files at location key
func (a *AppFiles) ReadMeta(locationKey string) (*domain.AppFilesMetadata, domain.Error) {

	jsonPath := filepath.Join(appsPath, locationKey, "application.json")
	jsonHandle, err := os.Open(jsonPath)
	if err != nil {
		// here the error might be that application.json is not in app?
		// Or it could be a more internal problem, like directory of apps not where it's expected to be.
		// Or it could be a bad location key, like it was deleted but DB doesn't know.
		if !a.locationKeyExists(locationKey) {
			// TODO: log that location key does not exist
			return nil, dserror.New(dserror.InternalError, "ReadMeta: Location key not found "+locationKey)
		}
		return nil, dserror.New(dserror.AppConfigNotFound)
	}
	defer jsonHandle.Close()

	meta, dsErr := decodeAppJSON(jsonHandle)
	if dsErr != nil {
		return nil, dsErr
	}

	dsErr = validateAppMeta(meta)
	if dsErr != nil {
		return nil, dsErr
	}

	return meta, nil
}

func decodeAppJSON(r io.Reader) (*domain.AppFilesMetadata, domain.Error) {
	var meta domain.AppFilesMetadata
	dec := json.NewDecoder(r)
	err := dec.Decode(&meta)
	if err != nil {
		return nil, dserror.New(dserror.AppConfigParseFailed, err.Error())
	}
	return &meta, nil
}

func validateAppMeta(meta *domain.AppFilesMetadata) domain.Error {
	if meta.AppName == "" {
		return dserror.New(dserror.AppConfigProblem, "Name can not be blank")
	}
	if meta.AppVersion == "" {
		return dserror.New(dserror.AppConfigProblem, "Version can not be blank")
	}
	return nil
}

func (a *AppFiles) locationKeyExists(locationKey string) bool {
	_, err := os.Stat(filepath.Join(appsPath, locationKey))
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true // OK but there could be aonther problem, like permissions out of whack?
	// Should probably log that as warning at least.
}
