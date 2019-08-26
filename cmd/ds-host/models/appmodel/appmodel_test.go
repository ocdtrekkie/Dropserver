package appmodel

import (
	"testing"

	"github.com/teleclimber/DropServer/cmd/ds-host/domain"
	"github.com/teleclimber/DropServer/cmd/ds-host/migrate"
)

func TestPrepareStatements(t *testing.T) {
	h := migrate.MakeSqliteDummyDB()
	defer h.Close()

	db := &domain.DB{
		Handle: h}

	appModel := &AppModel{
		DB: db}

	appModel.PrepareStatements()
}

func TestGetFromNonExistentID(t *testing.T) {
	h := migrate.MakeSqliteDummyDB()
	defer h.Close()

	db := &domain.DB{
		Handle: h}

	appModel := &AppModel{
		DB: db}

	appModel.PrepareStatements()

	// There should be an error, but no panics
	_, dsErr := appModel.GetFromID(10)
	if dsErr == nil {
		t.Error(dsErr)
	}
}

func TestCreate(t *testing.T) {

	h := migrate.MakeSqliteDummyDB()
	defer h.Close()

	db := &domain.DB{
		Handle: h}

	appModel := &AppModel{
		DB: db}

	appModel.PrepareStatements()

	app, dsErr := appModel.Create(domain.UserID(1), "test-app")
	if dsErr != nil {
		t.Error(dsErr)
	}

	if app.Name != "test-app" {
		t.Error("input name does not match output name", app)
	}
}

func TestGetFromID(t *testing.T) {
	h := migrate.MakeSqliteDummyDB()
	defer h.Close()

	db := &domain.DB{
		Handle: h}

	appModel := &AppModel{
		DB: db}

	appModel.PrepareStatements()

	_, dsErr := appModel.Create(7, "test-app")
	if dsErr != nil {
		t.Error(dsErr)
	}

	// There should now be one row so app id 1 should return something
	app, dsErr := appModel.GetFromID(1)
	if dsErr != nil {
		t.Error(dsErr)
	}

	if app.AppID != 1 {
		t.Error("app.AppID does not match requested ID", app)
	}
}

func TestGetOwner(t *testing.T) {
	h := migrate.MakeSqliteDummyDB()
	defer h.Close()

	db := &domain.DB{
		Handle: h}

	appModel := &AppModel{
		DB: db}

	appModel.PrepareStatements()

	_, dsErr := appModel.Create(7, "test-app")
	if dsErr != nil {
		t.Error(dsErr)
	}

	_, dsErr = appModel.Create(7, "test-app2")
	if dsErr != nil {
		t.Error(dsErr)
	}

	_, dsErr = appModel.Create(11, "bad-app")
	if dsErr != nil {
		t.Error(dsErr)
	}

	apps, dsErr := appModel.GetForOwner(7)
	if dsErr != nil {
		t.Error(dsErr)
	}

	if len(apps) != 2 {
		t.Error("There should be two apps found")
	}
}

func TestVersion(t *testing.T) {
	h := migrate.MakeSqliteDummyDB()
	defer h.Close()

	db := &domain.DB{
		Handle: h}

	appModel := &AppModel{
		DB: db}

	appModel.PrepareStatements()

	appVersion, dsErr := appModel.CreateVersion(1, "0.0.1", 7, "foo-location")
	if dsErr != nil {
		t.Error(dsErr)
	}

	if appVersion.Version != "0.0.1" {
		t.Error("input version does not match output version", appVersion)
	}
	if appVersion.Schema != 7 {
		t.Error("input schema does not match output schema", appVersion)
	}
	if appVersion.LocationKey != "foo-location" {
		t.Error("input does not match output location key", appVersion)
	}

	// just go ahead and test GetVersion here for completeness
	appVersion, dsErr = appModel.GetVersion(1, "0.0.1")
	if dsErr != nil {
		t.Error(dsErr)
	}
	if appVersion.Version != "0.0.1" {
		t.Error("input version does not match output version", appVersion)
	}
	if appVersion.LocationKey != "foo-location" {
		t.Error("input does not match output location key", appVersion)
	}
}

func TestGetVersionForApp(t *testing.T) {
	h := migrate.MakeSqliteDummyDB()
	defer h.Close()

	db := &domain.DB{
		Handle: h}

	appModel := &AppModel{
		DB: db}

	appModel.PrepareStatements()

	
	ins := []struct{
		appID domain.AppID
		version domain.Version
		schema int
		location string
	}{
		{7, "0.0.1", 1, "foo-location"},
		{7, "0.0.2", 2, "2foo-location"},
		{7, "0.0.3", 3, "3foo-location"},
		{11, "0.0.1", 1, "bar-location"},
	}

	for _, i := range ins {
		_, dsErr := appModel.CreateVersion(i.appID, i.version, i.schema, i.location)
		if dsErr != nil {
			t.Error(dsErr)
		}
	}

	vers, dsErr := appModel.GetVersionsForApp(7)
	if dsErr != nil {
		t.Error(dsErr)
	}
	if len(vers) != 3 {
		t.Error("Got wrong number of results: should be 3")
	}

	vers, dsErr = appModel.GetVersionsForApp(1)
	if dsErr != nil {
		t.Error(dsErr)
	}
	if len(vers) != 0 {
		t.Error("Got wrong number of results: should be 0")
	}
}

