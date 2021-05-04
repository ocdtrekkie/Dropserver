package appspacefilesmodel

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/teleclimber/DropServer/cmd/ds-host/domain"
)

func TestDelete(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// create temp dir and put that in runtime config.
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	cfg := &domain.RuntimeConfig{}
	cfg.DataDir = dir
	cfg.Exec.AppspacesPath = dir

	m := AppspaceFilesModel{
		Config: cfg}

	locKey, err := m.CreateLocation()
	if err != nil {
		t.Fatal(err)
	}

	if locKey == "" {
		t.Fatal("location key can not be empty string")
	}

	appspacePath := filepath.Join(cfg.Exec.AppspacesPath, locKey)

	if _, err := os.Stat(appspacePath); os.IsNotExist(err) {
		// path/to/whatever does not exist
		t.Fatal("appspace path should exist")
	}

	// err = m.Delete(locKey)
	// if dsErr != nil {
	// 	t.Fatal(err)
	// }

	// _, err = os.Stat(filepath.Join(appsPath, locKey))
	// if err == nil || !os.IsNotExist(err) {
	// 	t.Fatal("expected not exist error", err)
	// }
}

func TestGetBackups(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// create temp dir and put that in runtime config.
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	cfg := &domain.RuntimeConfig{}
	cfg.DataDir = dir
	cfg.Exec.AppspacesPath = dir

	m := AppspaceFilesModel{
		Config: cfg}

	loc := "abcLOC"

	backupsDir := filepath.Join(dir, loc, "backups")
	err = os.MkdirAll(backupsDir, 0755)
	if err != nil {
		t.Fatal(err)
	}

	entries, err := m.GetBackups(loc)
	if err != nil {
		t.Error(err)
	}
	if len(entries) != 0 {
		t.Error("expected zero entries")
	}

	file1 := "1234-56-78_7890.zip"
	file2 := "9999-56-78_7890_1.zip"

	err = ioutil.WriteFile(filepath.Join(backupsDir, file2), []byte("test data"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile(filepath.Join(backupsDir, file1), []byte("test data"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	entries, err = m.GetBackups(loc)
	if err != nil {
		t.Error(err)
	}
	if len(entries) != 2 {
		t.Error("expected 2 entries")
	}
	if entries[0] != file2 {
		t.Error("expected 9999-* entry first " + entries[0])
	}

	// I'm just going to test delete while we're all set up:
	err = m.DeleteBackup(loc, file2)
	if err != nil {
		t.Error(err)
	}
	entries, err = m.GetBackups(loc)
	if err != nil {
		t.Error(err)
	}
	if len(entries) != 1 {
		t.Error("expected 1 entries")
	}
	if entries[0] != file1 {
		t.Error("expected 1234-* entry " + entries[0])
	}

}
