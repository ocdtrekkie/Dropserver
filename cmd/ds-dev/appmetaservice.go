package main

import (
	"encoding/json"
	"fmt"

	"github.com/teleclimber/DropServer/cmd/ds-host/domain"
	"github.com/teleclimber/twine-go/twine"
)

type AppMetaService struct {
	DevAppModel   *DevAppModel `checkinject:"required"`
	AppFilesModel interface {
		ReadMigrations(locationKey string) ([]byte, error)
	} `checkinject:"required"`
	AppGetter interface {
		ValidateMigrationSteps(migrations []domain.MigrationStep) ([]int, error)
	} `checkinject:"required"`
	DevAppProcessEvents interface {
		Subscribe() (AppProcessEvent, <-chan AppProcessEvent)
		Unsubscribe(<-chan AppProcessEvent)
	} `checkinject:"required"`
	AppVersionEvents interface {
		Subscribe(chan<- string)
		Unsubscribe(chan<- string)
	} `checkinject:"required"`
}

func (s *AppMetaService) HandleMessage(m twine.ReceivedMessageI) {

}

func (s *AppMetaService) Start(t *twine.Twine) {
	appVersionEvent := make(chan string)
	s.AppVersionEvents.Subscribe(appVersionEvent)
	go func() {
		for range appVersionEvent {
			go s.sendAppData(t)
		}
	}()
	s.sendAppData(t)

	ev, appProcessCh := s.DevAppProcessEvents.Subscribe()
	go func() {
		for appProcessEvent := range appProcessCh {
			go s.sendAppGetEvent(t, appProcessEvent)
		}
	}()
	go s.sendAppGetEvent(t, ev)

	t.WaitClose()

	s.DevAppProcessEvents.Unsubscribe(appProcessCh)
	s.AppVersionEvents.Unsubscribe(appVersionEvent)
}

const appDataCmd = 12
const appProcessEventCmd = 13

type AppMetaResp struct {
	AppName       string                 `json:"name"`
	AppVersion    domain.Version         `json:"version"`
	SchemaVersion int                    `json:"schema"`
	APIVersion    domain.APIVersion      `json:"api_version"`
	Migrations    []domain.MigrationStep `json:"migration_steps"`
	Schemas       []int                  `json:"schemas"`
}

func (s *AppMetaService) sendAppData(twine *twine.Twine) {
	resp := AppMetaResp{
		AppName:       s.DevAppModel.App.Name,
		AppVersion:    s.DevAppModel.Ver.Version,
		APIVersion:    s.DevAppModel.Ver.APIVersion,
		SchemaVersion: s.DevAppModel.Ver.Schema}
	migrationsBytes, err := s.AppFilesModel.ReadMigrations("")
	if err == nil {
		var migrations []domain.MigrationStep
		err = json.Unmarshal(migrationsBytes, &migrations)
		if err == nil {
			resp.Migrations = migrations
			schemas, _ := s.AppGetter.ValidateMigrationSteps(migrations)
			resp.Schemas = schemas
		}
	}

	bytes, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("sendAppData json Marshal Error: " + err.Error())
	}
	_, err = twine.SendBlock(appDataService, appDataCmd, bytes)
	if err != nil {
		fmt.Println("sendAppData SendBlock Error: " + err.Error())
	}
}

func (s *AppMetaService) sendAppGetEvent(twine *twine.Twine, ev AppProcessEvent) {
	bytes, err := json.Marshal(ev)
	if err != nil {
		fmt.Println("sendAppGetEvent json Marshal Error: " + err.Error())
	}
	_, err = twine.SendBlock(appDataService, appProcessEventCmd, bytes)
	if err != nil {
		fmt.Println("sendAppGetEvent SendBlock Error: " + err.Error())
	}
}
