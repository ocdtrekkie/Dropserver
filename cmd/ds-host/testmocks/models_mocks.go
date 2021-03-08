// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/teleclimber/DropServer/cmd/ds-host/testmocks (interfaces: UserModel,AppFilesModel,AppModel,AppspaceModel,AppspaceFilesModel,AppspaceInfoModels,AppspaceContactModel,DropIDModel,MigrationJobModel)

// Package testmocks is a generated GoMock package.
package testmocks

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/teleclimber/DropServer/cmd/ds-host/domain"
	nulltypes "github.com/teleclimber/DropServer/internal/nulltypes"
	reflect "reflect"
)

// MockUserModel is a mock of UserModel interface
type MockUserModel struct {
	ctrl     *gomock.Controller
	recorder *MockUserModelMockRecorder
}

// MockUserModelMockRecorder is the mock recorder for MockUserModel
type MockUserModelMockRecorder struct {
	mock *MockUserModel
}

// NewMockUserModel creates a new mock instance
func NewMockUserModel(ctrl *gomock.Controller) *MockUserModel {
	mock := &MockUserModel{ctrl: ctrl}
	mock.recorder = &MockUserModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserModel) EXPECT() *MockUserModelMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockUserModel) Create(arg0, arg1 string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockUserModelMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserModel)(nil).Create), arg0, arg1)
}

// DeleteAdmin mocks base method
func (m *MockUserModel) DeleteAdmin(arg0 domain.UserID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAdmin", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAdmin indicates an expected call of DeleteAdmin
func (mr *MockUserModelMockRecorder) DeleteAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAdmin", reflect.TypeOf((*MockUserModel)(nil).DeleteAdmin), arg0)
}

// GetAll mocks base method
func (m *MockUserModel) GetAll() ([]domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockUserModelMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockUserModel)(nil).GetAll))
}

// GetAllAdmins mocks base method
func (m *MockUserModel) GetAllAdmins() ([]domain.UserID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAdmins")
	ret0, _ := ret[0].([]domain.UserID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAdmins indicates an expected call of GetAllAdmins
func (mr *MockUserModelMockRecorder) GetAllAdmins() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAdmins", reflect.TypeOf((*MockUserModel)(nil).GetAllAdmins))
}

// GetFromEmail mocks base method
func (m *MockUserModel) GetFromEmail(arg0 string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFromEmail", arg0)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFromEmail indicates an expected call of GetFromEmail
func (mr *MockUserModelMockRecorder) GetFromEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFromEmail", reflect.TypeOf((*MockUserModel)(nil).GetFromEmail), arg0)
}

// GetFromEmailPassword mocks base method
func (m *MockUserModel) GetFromEmailPassword(arg0, arg1 string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFromEmailPassword", arg0, arg1)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFromEmailPassword indicates an expected call of GetFromEmailPassword
func (mr *MockUserModelMockRecorder) GetFromEmailPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFromEmailPassword", reflect.TypeOf((*MockUserModel)(nil).GetFromEmailPassword), arg0, arg1)
}

// GetFromID mocks base method
func (m *MockUserModel) GetFromID(arg0 domain.UserID) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFromID", arg0)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFromID indicates an expected call of GetFromID
func (mr *MockUserModelMockRecorder) GetFromID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFromID", reflect.TypeOf((*MockUserModel)(nil).GetFromID), arg0)
}

// IsAdmin mocks base method
func (m *MockUserModel) IsAdmin(arg0 domain.UserID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAdmin", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAdmin indicates an expected call of IsAdmin
func (mr *MockUserModelMockRecorder) IsAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAdmin", reflect.TypeOf((*MockUserModel)(nil).IsAdmin), arg0)
}

// MakeAdmin mocks base method
func (m *MockUserModel) MakeAdmin(arg0 domain.UserID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeAdmin", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeAdmin indicates an expected call of MakeAdmin
func (mr *MockUserModelMockRecorder) MakeAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeAdmin", reflect.TypeOf((*MockUserModel)(nil).MakeAdmin), arg0)
}

// UpdatePassword mocks base method
func (m *MockUserModel) UpdatePassword(arg0 domain.UserID, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword
func (mr *MockUserModelMockRecorder) UpdatePassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockUserModel)(nil).UpdatePassword), arg0, arg1)
}

// MockAppFilesModel is a mock of AppFilesModel interface
type MockAppFilesModel struct {
	ctrl     *gomock.Controller
	recorder *MockAppFilesModelMockRecorder
}

// MockAppFilesModelMockRecorder is the mock recorder for MockAppFilesModel
type MockAppFilesModelMockRecorder struct {
	mock *MockAppFilesModel
}

// NewMockAppFilesModel creates a new mock instance
func NewMockAppFilesModel(ctrl *gomock.Controller) *MockAppFilesModel {
	mock := &MockAppFilesModel{ctrl: ctrl}
	mock.recorder = &MockAppFilesModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppFilesModel) EXPECT() *MockAppFilesModelMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockAppFilesModel) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockAppFilesModelMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAppFilesModel)(nil).Delete), arg0)
}

// ReadMeta mocks base method
func (m *MockAppFilesModel) ReadMeta(arg0 string) (*domain.AppFilesMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMeta", arg0)
	ret0, _ := ret[0].(*domain.AppFilesMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadMeta indicates an expected call of ReadMeta
func (mr *MockAppFilesModelMockRecorder) ReadMeta(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMeta", reflect.TypeOf((*MockAppFilesModel)(nil).ReadMeta), arg0)
}

// Save mocks base method
func (m *MockAppFilesModel) Save(arg0 *map[string][]byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save
func (mr *MockAppFilesModelMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockAppFilesModel)(nil).Save), arg0)
}

// MockAppModel is a mock of AppModel interface
type MockAppModel struct {
	ctrl     *gomock.Controller
	recorder *MockAppModelMockRecorder
}

// MockAppModelMockRecorder is the mock recorder for MockAppModel
type MockAppModelMockRecorder struct {
	mock *MockAppModel
}

// NewMockAppModel creates a new mock instance
func NewMockAppModel(ctrl *gomock.Controller) *MockAppModel {
	mock := &MockAppModel{ctrl: ctrl}
	mock.recorder = &MockAppModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppModel) EXPECT() *MockAppModelMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockAppModel) Create(arg0 domain.UserID, arg1 string) (*domain.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*domain.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockAppModelMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAppModel)(nil).Create), arg0, arg1)
}

// CreateVersion mocks base method
func (m *MockAppModel) CreateVersion(arg0 domain.AppID, arg1 domain.Version, arg2 int, arg3 domain.APIVersion, arg4 string) (*domain.AppVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVersion", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*domain.AppVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVersion indicates an expected call of CreateVersion
func (mr *MockAppModelMockRecorder) CreateVersion(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVersion", reflect.TypeOf((*MockAppModel)(nil).CreateVersion), arg0, arg1, arg2, arg3, arg4)
}

// DeleteVersion mocks base method
func (m *MockAppModel) DeleteVersion(arg0 domain.AppID, arg1 domain.Version) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVersion", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVersion indicates an expected call of DeleteVersion
func (mr *MockAppModelMockRecorder) DeleteVersion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVersion", reflect.TypeOf((*MockAppModel)(nil).DeleteVersion), arg0, arg1)
}

// GetForOwner mocks base method
func (m *MockAppModel) GetForOwner(arg0 domain.UserID) ([]*domain.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetForOwner", arg0)
	ret0, _ := ret[0].([]*domain.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetForOwner indicates an expected call of GetForOwner
func (mr *MockAppModelMockRecorder) GetForOwner(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetForOwner", reflect.TypeOf((*MockAppModel)(nil).GetForOwner), arg0)
}

// GetFromID mocks base method
func (m *MockAppModel) GetFromID(arg0 domain.AppID) (*domain.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFromID", arg0)
	ret0, _ := ret[0].(*domain.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFromID indicates an expected call of GetFromID
func (mr *MockAppModelMockRecorder) GetFromID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFromID", reflect.TypeOf((*MockAppModel)(nil).GetFromID), arg0)
}

// GetVersion mocks base method
func (m *MockAppModel) GetVersion(arg0 domain.AppID, arg1 domain.Version) (*domain.AppVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersion", arg0, arg1)
	ret0, _ := ret[0].(*domain.AppVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersion indicates an expected call of GetVersion
func (mr *MockAppModelMockRecorder) GetVersion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockAppModel)(nil).GetVersion), arg0, arg1)
}

// GetVersionsForApp mocks base method
func (m *MockAppModel) GetVersionsForApp(arg0 domain.AppID) ([]*domain.AppVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersionsForApp", arg0)
	ret0, _ := ret[0].([]*domain.AppVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersionsForApp indicates an expected call of GetVersionsForApp
func (mr *MockAppModelMockRecorder) GetVersionsForApp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersionsForApp", reflect.TypeOf((*MockAppModel)(nil).GetVersionsForApp), arg0)
}

// MockAppspaceModel is a mock of AppspaceModel interface
type MockAppspaceModel struct {
	ctrl     *gomock.Controller
	recorder *MockAppspaceModelMockRecorder
}

// MockAppspaceModelMockRecorder is the mock recorder for MockAppspaceModel
type MockAppspaceModelMockRecorder struct {
	mock *MockAppspaceModel
}

// NewMockAppspaceModel creates a new mock instance
func NewMockAppspaceModel(ctrl *gomock.Controller) *MockAppspaceModel {
	mock := &MockAppspaceModel{ctrl: ctrl}
	mock.recorder = &MockAppspaceModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppspaceModel) EXPECT() *MockAppspaceModelMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockAppspaceModel) Create(arg0 domain.Appspace) (*domain.Appspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*domain.Appspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockAppspaceModelMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAppspaceModel)(nil).Create), arg0)
}

// GetForApp mocks base method
func (m *MockAppspaceModel) GetForApp(arg0 domain.AppID) ([]*domain.Appspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetForApp", arg0)
	ret0, _ := ret[0].([]*domain.Appspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetForApp indicates an expected call of GetForApp
func (mr *MockAppspaceModelMockRecorder) GetForApp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetForApp", reflect.TypeOf((*MockAppspaceModel)(nil).GetForApp), arg0)
}

// GetForAppVersion mocks base method
func (m *MockAppspaceModel) GetForAppVersion(arg0 domain.AppID, arg1 domain.Version) ([]*domain.Appspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetForAppVersion", arg0, arg1)
	ret0, _ := ret[0].([]*domain.Appspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetForAppVersion indicates an expected call of GetForAppVersion
func (mr *MockAppspaceModelMockRecorder) GetForAppVersion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetForAppVersion", reflect.TypeOf((*MockAppspaceModel)(nil).GetForAppVersion), arg0, arg1)
}

// GetForOwner mocks base method
func (m *MockAppspaceModel) GetForOwner(arg0 domain.UserID) ([]*domain.Appspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetForOwner", arg0)
	ret0, _ := ret[0].([]*domain.Appspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetForOwner indicates an expected call of GetForOwner
func (mr *MockAppspaceModelMockRecorder) GetForOwner(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetForOwner", reflect.TypeOf((*MockAppspaceModel)(nil).GetForOwner), arg0)
}

// GetFromDomain mocks base method
func (m *MockAppspaceModel) GetFromDomain(arg0 string) (*domain.Appspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFromDomain", arg0)
	ret0, _ := ret[0].(*domain.Appspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFromDomain indicates an expected call of GetFromDomain
func (mr *MockAppspaceModelMockRecorder) GetFromDomain(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFromDomain", reflect.TypeOf((*MockAppspaceModel)(nil).GetFromDomain), arg0)
}

// GetFromID mocks base method
func (m *MockAppspaceModel) GetFromID(arg0 domain.AppspaceID) (*domain.Appspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFromID", arg0)
	ret0, _ := ret[0].(*domain.Appspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFromID indicates an expected call of GetFromID
func (mr *MockAppspaceModelMockRecorder) GetFromID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFromID", reflect.TypeOf((*MockAppspaceModel)(nil).GetFromID), arg0)
}

// Pause mocks base method
func (m *MockAppspaceModel) Pause(arg0 domain.AppspaceID, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pause", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Pause indicates an expected call of Pause
func (mr *MockAppspaceModelMockRecorder) Pause(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pause", reflect.TypeOf((*MockAppspaceModel)(nil).Pause), arg0, arg1)
}

// SetVersion mocks base method
func (m *MockAppspaceModel) SetVersion(arg0 domain.AppspaceID, arg1 domain.Version) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetVersion", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetVersion indicates an expected call of SetVersion
func (mr *MockAppspaceModelMockRecorder) SetVersion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVersion", reflect.TypeOf((*MockAppspaceModel)(nil).SetVersion), arg0, arg1)
}

// MockAppspaceFilesModel is a mock of AppspaceFilesModel interface
type MockAppspaceFilesModel struct {
	ctrl     *gomock.Controller
	recorder *MockAppspaceFilesModelMockRecorder
}

// MockAppspaceFilesModelMockRecorder is the mock recorder for MockAppspaceFilesModel
type MockAppspaceFilesModelMockRecorder struct {
	mock *MockAppspaceFilesModel
}

// NewMockAppspaceFilesModel creates a new mock instance
func NewMockAppspaceFilesModel(ctrl *gomock.Controller) *MockAppspaceFilesModel {
	mock := &MockAppspaceFilesModel{ctrl: ctrl}
	mock.recorder = &MockAppspaceFilesModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppspaceFilesModel) EXPECT() *MockAppspaceFilesModelMockRecorder {
	return m.recorder
}

// CreateLocation mocks base method
func (m *MockAppspaceFilesModel) CreateLocation() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLocation")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLocation indicates an expected call of CreateLocation
func (mr *MockAppspaceFilesModelMockRecorder) CreateLocation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLocation", reflect.TypeOf((*MockAppspaceFilesModel)(nil).CreateLocation))
}

// MockAppspaceInfoModels is a mock of AppspaceInfoModels interface
type MockAppspaceInfoModels struct {
	ctrl     *gomock.Controller
	recorder *MockAppspaceInfoModelsMockRecorder
}

// MockAppspaceInfoModelsMockRecorder is the mock recorder for MockAppspaceInfoModels
type MockAppspaceInfoModelsMockRecorder struct {
	mock *MockAppspaceInfoModels
}

// NewMockAppspaceInfoModels creates a new mock instance
func NewMockAppspaceInfoModels(ctrl *gomock.Controller) *MockAppspaceInfoModels {
	mock := &MockAppspaceInfoModels{ctrl: ctrl}
	mock.recorder = &MockAppspaceInfoModelsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppspaceInfoModels) EXPECT() *MockAppspaceInfoModelsMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockAppspaceInfoModels) Get(arg0 domain.AppspaceID) domain.AppspaceInfoModel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(domain.AppspaceInfoModel)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockAppspaceInfoModelsMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAppspaceInfoModels)(nil).Get), arg0)
}

// GetSchema mocks base method
func (m *MockAppspaceInfoModels) GetSchema(arg0 domain.AppspaceID) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchema", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchema indicates an expected call of GetSchema
func (mr *MockAppspaceInfoModelsMockRecorder) GetSchema(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchema", reflect.TypeOf((*MockAppspaceInfoModels)(nil).GetSchema), arg0)
}

// Init mocks base method
func (m *MockAppspaceInfoModels) Init() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Init")
}

// Init indicates an expected call of Init
func (mr *MockAppspaceInfoModelsMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockAppspaceInfoModels)(nil).Init))
}

// MockAppspaceContactModel is a mock of AppspaceContactModel interface
type MockAppspaceContactModel struct {
	ctrl     *gomock.Controller
	recorder *MockAppspaceContactModelMockRecorder
}

// MockAppspaceContactModelMockRecorder is the mock recorder for MockAppspaceContactModel
type MockAppspaceContactModelMockRecorder struct {
	mock *MockAppspaceContactModel
}

// NewMockAppspaceContactModel creates a new mock instance
func NewMockAppspaceContactModel(ctrl *gomock.Controller) *MockAppspaceContactModel {
	mock := &MockAppspaceContactModel{ctrl: ctrl}
	mock.recorder = &MockAppspaceContactModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppspaceContactModel) EXPECT() *MockAppspaceContactModelMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockAppspaceContactModel) Create(arg0 domain.UserID, arg1, arg2 string) (domain.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(domain.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockAppspaceContactModelMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAppspaceContactModel)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method
func (m *MockAppspaceContactModel) Delete(arg0 domain.UserID, arg1 domain.ContactID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockAppspaceContactModelMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAppspaceContactModel)(nil).Delete), arg0, arg1)
}

// DeleteAppspaceContact mocks base method
func (m *MockAppspaceContactModel) DeleteAppspaceContact(arg0 domain.AppspaceID, arg1 domain.ContactID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAppspaceContact", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAppspaceContact indicates an expected call of DeleteAppspaceContact
func (mr *MockAppspaceContactModelMockRecorder) DeleteAppspaceContact(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAppspaceContact", reflect.TypeOf((*MockAppspaceContactModel)(nil).DeleteAppspaceContact), arg0, arg1)
}

// Get mocks base method
func (m *MockAppspaceContactModel) Get(arg0 domain.ContactID) (domain.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(domain.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockAppspaceContactModelMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAppspaceContactModel)(nil).Get), arg0)
}

// GetAppspaceContacts mocks base method
func (m *MockAppspaceContactModel) GetAppspaceContacts(arg0 domain.AppspaceID) ([]domain.AppspaceContact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppspaceContacts", arg0)
	ret0, _ := ret[0].([]domain.AppspaceContact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppspaceContacts indicates an expected call of GetAppspaceContacts
func (mr *MockAppspaceContactModelMockRecorder) GetAppspaceContacts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppspaceContacts", reflect.TypeOf((*MockAppspaceContactModel)(nil).GetAppspaceContacts), arg0)
}

// GetByProxy mocks base method
func (m *MockAppspaceContactModel) GetByProxy(arg0 domain.AppspaceID, arg1 domain.ProxyID) (domain.ContactID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByProxy", arg0, arg1)
	ret0, _ := ret[0].(domain.ContactID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByProxy indicates an expected call of GetByProxy
func (mr *MockAppspaceContactModelMockRecorder) GetByProxy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByProxy", reflect.TypeOf((*MockAppspaceContactModel)(nil).GetByProxy), arg0, arg1)
}

// GetContactAppspaces mocks base method
func (m *MockAppspaceContactModel) GetContactAppspaces(arg0 domain.ContactID) ([]domain.AppspaceContact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContactAppspaces", arg0)
	ret0, _ := ret[0].([]domain.AppspaceContact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContactAppspaces indicates an expected call of GetContactAppspaces
func (mr *MockAppspaceContactModelMockRecorder) GetContactAppspaces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContactAppspaces", reflect.TypeOf((*MockAppspaceContactModel)(nil).GetContactAppspaces), arg0)
}

// GetContactProxy mocks base method
func (m *MockAppspaceContactModel) GetContactProxy(arg0 domain.AppspaceID, arg1 domain.ContactID) (domain.ProxyID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContactProxy", arg0, arg1)
	ret0, _ := ret[0].(domain.ProxyID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContactProxy indicates an expected call of GetContactProxy
func (mr *MockAppspaceContactModelMockRecorder) GetContactProxy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContactProxy", reflect.TypeOf((*MockAppspaceContactModel)(nil).GetContactProxy), arg0, arg1)
}

// GetForUser mocks base method
func (m *MockAppspaceContactModel) GetForUser(arg0 domain.UserID) ([]domain.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetForUser", arg0)
	ret0, _ := ret[0].([]domain.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetForUser indicates an expected call of GetForUser
func (mr *MockAppspaceContactModelMockRecorder) GetForUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetForUser", reflect.TypeOf((*MockAppspaceContactModel)(nil).GetForUser), arg0)
}

// InsertAppspaceContact mocks base method
func (m *MockAppspaceContactModel) InsertAppspaceContact(arg0 domain.AppspaceID, arg1 domain.ContactID, arg2 domain.ProxyID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertAppspaceContact", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertAppspaceContact indicates an expected call of InsertAppspaceContact
func (mr *MockAppspaceContactModelMockRecorder) InsertAppspaceContact(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertAppspaceContact", reflect.TypeOf((*MockAppspaceContactModel)(nil).InsertAppspaceContact), arg0, arg1, arg2)
}

// Update mocks base method
func (m *MockAppspaceContactModel) Update(arg0 domain.UserID, arg1 domain.ContactID, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockAppspaceContactModelMockRecorder) Update(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAppspaceContactModel)(nil).Update), arg0, arg1, arg2, arg3)
}

// MockDropIDModel is a mock of DropIDModel interface
type MockDropIDModel struct {
	ctrl     *gomock.Controller
	recorder *MockDropIDModelMockRecorder
}

// MockDropIDModelMockRecorder is the mock recorder for MockDropIDModel
type MockDropIDModelMockRecorder struct {
	mock *MockDropIDModel
}

// NewMockDropIDModel creates a new mock instance
func NewMockDropIDModel(ctrl *gomock.Controller) *MockDropIDModel {
	mock := &MockDropIDModel{ctrl: ctrl}
	mock.recorder = &MockDropIDModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDropIDModel) EXPECT() *MockDropIDModelMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockDropIDModel) Create(arg0 domain.UserID, arg1, arg2, arg3 string) (domain.DropID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(domain.DropID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockDropIDModelMockRecorder) Create(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDropIDModel)(nil).Create), arg0, arg1, arg2, arg3)
}

// Delete mocks base method
func (m *MockDropIDModel) Delete(arg0 domain.UserID, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockDropIDModelMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDropIDModel)(nil).Delete), arg0, arg1, arg2)
}

// Get mocks base method
func (m *MockDropIDModel) Get(arg0, arg1 string) (domain.DropID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(domain.DropID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockDropIDModelMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDropIDModel)(nil).Get), arg0, arg1)
}

// GetForUser mocks base method
func (m *MockDropIDModel) GetForUser(arg0 domain.UserID) ([]domain.DropID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetForUser", arg0)
	ret0, _ := ret[0].([]domain.DropID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetForUser indicates an expected call of GetForUser
func (mr *MockDropIDModelMockRecorder) GetForUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetForUser", reflect.TypeOf((*MockDropIDModel)(nil).GetForUser), arg0)
}

// Update mocks base method
func (m *MockDropIDModel) Update(arg0 domain.UserID, arg1, arg2, arg3 string) (domain.DropID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(domain.DropID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockDropIDModelMockRecorder) Update(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDropIDModel)(nil).Update), arg0, arg1, arg2, arg3)
}

// MockMigrationJobModel is a mock of MigrationJobModel interface
type MockMigrationJobModel struct {
	ctrl     *gomock.Controller
	recorder *MockMigrationJobModelMockRecorder
}

// MockMigrationJobModelMockRecorder is the mock recorder for MockMigrationJobModel
type MockMigrationJobModelMockRecorder struct {
	mock *MockMigrationJobModel
}

// NewMockMigrationJobModel creates a new mock instance
func NewMockMigrationJobModel(ctrl *gomock.Controller) *MockMigrationJobModel {
	mock := &MockMigrationJobModel{ctrl: ctrl}
	mock.recorder = &MockMigrationJobModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMigrationJobModel) EXPECT() *MockMigrationJobModelMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockMigrationJobModel) Create(arg0 domain.UserID, arg1 domain.AppspaceID, arg2 domain.Version, arg3 bool) (*domain.MigrationJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*domain.MigrationJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockMigrationJobModelMockRecorder) Create(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMigrationJobModel)(nil).Create), arg0, arg1, arg2, arg3)
}

// GetForAppspace mocks base method
func (m *MockMigrationJobModel) GetForAppspace(arg0 domain.AppspaceID) ([]*domain.MigrationJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetForAppspace", arg0)
	ret0, _ := ret[0].([]*domain.MigrationJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetForAppspace indicates an expected call of GetForAppspace
func (mr *MockMigrationJobModelMockRecorder) GetForAppspace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetForAppspace", reflect.TypeOf((*MockMigrationJobModel)(nil).GetForAppspace), arg0)
}

// GetJob mocks base method
func (m *MockMigrationJobModel) GetJob(arg0 domain.JobID) (*domain.MigrationJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJob", arg0)
	ret0, _ := ret[0].(*domain.MigrationJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJob indicates an expected call of GetJob
func (mr *MockMigrationJobModelMockRecorder) GetJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJob", reflect.TypeOf((*MockMigrationJobModel)(nil).GetJob), arg0)
}

// GetPending mocks base method
func (m *MockMigrationJobModel) GetPending() ([]*domain.MigrationJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPending")
	ret0, _ := ret[0].([]*domain.MigrationJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPending indicates an expected call of GetPending
func (mr *MockMigrationJobModelMockRecorder) GetPending() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPending", reflect.TypeOf((*MockMigrationJobModel)(nil).GetPending))
}

// GetRunning mocks base method
func (m *MockMigrationJobModel) GetRunning() ([]domain.MigrationJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunning")
	ret0, _ := ret[0].([]domain.MigrationJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRunning indicates an expected call of GetRunning
func (mr *MockMigrationJobModelMockRecorder) GetRunning() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunning", reflect.TypeOf((*MockMigrationJobModel)(nil).GetRunning))
}

// SetFinished mocks base method
func (m *MockMigrationJobModel) SetFinished(arg0 domain.JobID, arg1 nulltypes.NullString) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFinished", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetFinished indicates an expected call of SetFinished
func (mr *MockMigrationJobModelMockRecorder) SetFinished(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFinished", reflect.TypeOf((*MockMigrationJobModel)(nil).SetFinished), arg0, arg1)
}

// SetStarted mocks base method
func (m *MockMigrationJobModel) SetStarted(arg0 domain.JobID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStarted", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetStarted indicates an expected call of SetStarted
func (mr *MockMigrationJobModelMockRecorder) SetStarted(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStarted", reflect.TypeOf((*MockMigrationJobModel)(nil).SetStarted), arg0)
}
