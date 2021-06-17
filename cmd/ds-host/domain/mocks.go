// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/teleclimber/DropServer/cmd/ds-host/domain (interfaces: MetricsI,SandboxI,DbConn,AppspaceMetaDB,AppspaceInfoModel,V0RouteModel,AppspaceRouteModels,StdInput)

// Package domain is a generated GoMock package.
package domain

import (
	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
	twine "github.com/teleclimber/DropServer/internal/twine"
	http "net/http"
	reflect "reflect"
	time "time"
)

// MockMetricsI is a mock of MetricsI interface
type MockMetricsI struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsIMockRecorder
}

// MockMetricsIMockRecorder is the mock recorder for MockMetricsI
type MockMetricsIMockRecorder struct {
	mock *MockMetricsI
}

// NewMockMetricsI creates a new mock instance
func NewMockMetricsI(ctrl *gomock.Controller) *MockMetricsI {
	mock := &MockMetricsI{ctrl: ctrl}
	mock.recorder = &MockMetricsIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetricsI) EXPECT() *MockMetricsIMockRecorder {
	return m.recorder
}

// HostHandleReq mocks base method
func (m *MockMetricsI) HostHandleReq(arg0 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HostHandleReq", arg0)
}

// HostHandleReq indicates an expected call of HostHandleReq
func (mr *MockMetricsIMockRecorder) HostHandleReq(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HostHandleReq", reflect.TypeOf((*MockMetricsI)(nil).HostHandleReq), arg0)
}

// MockSandboxI is a mock of SandboxI interface
type MockSandboxI struct {
	ctrl     *gomock.Controller
	recorder *MockSandboxIMockRecorder
}

// MockSandboxIMockRecorder is the mock recorder for MockSandboxI
type MockSandboxIMockRecorder struct {
	mock *MockSandboxI
}

// NewMockSandboxI creates a new mock instance
func NewMockSandboxI(ctrl *gomock.Controller) *MockSandboxI {
	mock := &MockSandboxI{ctrl: ctrl}
	mock.recorder = &MockSandboxIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSandboxI) EXPECT() *MockSandboxIMockRecorder {
	return m.recorder
}

// ExecFn mocks base method
func (m *MockSandboxI) ExecFn(arg0 AppspaceRouteHandler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecFn", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecFn indicates an expected call of ExecFn
func (mr *MockSandboxIMockRecorder) ExecFn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecFn", reflect.TypeOf((*MockSandboxI)(nil).ExecFn), arg0)
}

// GetTransport mocks base method
func (m *MockSandboxI) GetTransport() http.RoundTripper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransport")
	ret0, _ := ret[0].(http.RoundTripper)
	return ret0
}

// GetTransport indicates an expected call of GetTransport
func (mr *MockSandboxIMockRecorder) GetTransport() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransport", reflect.TypeOf((*MockSandboxI)(nil).GetTransport))
}

// Graceful mocks base method
func (m *MockSandboxI) Graceful() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Graceful")
}

// Graceful indicates an expected call of Graceful
func (mr *MockSandboxIMockRecorder) Graceful() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Graceful", reflect.TypeOf((*MockSandboxI)(nil).Graceful))
}

// ID mocks base method
func (m *MockSandboxI) ID() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(int)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockSandboxIMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockSandboxI)(nil).ID))
}

// Kill mocks base method
func (m *MockSandboxI) Kill() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Kill")
}

// Kill indicates an expected call of Kill
func (mr *MockSandboxIMockRecorder) Kill() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kill", reflect.TypeOf((*MockSandboxI)(nil).Kill))
}

// LastActive mocks base method
func (m *MockSandboxI) LastActive() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastActive")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// LastActive indicates an expected call of LastActive
func (mr *MockSandboxIMockRecorder) LastActive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastActive", reflect.TypeOf((*MockSandboxI)(nil).LastActive))
}

// SendMessage mocks base method
func (m *MockSandboxI) SendMessage(arg0, arg1 int, arg2 []byte) (twine.SentMessageI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", arg0, arg1, arg2)
	ret0, _ := ret[0].(twine.SentMessageI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessage indicates an expected call of SendMessage
func (mr *MockSandboxIMockRecorder) SendMessage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockSandboxI)(nil).SendMessage), arg0, arg1, arg2)
}

// SetStatus mocks base method
func (m *MockSandboxI) SetStatus(arg0 SandboxStatus) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStatus", arg0)
}

// SetStatus indicates an expected call of SetStatus
func (mr *MockSandboxIMockRecorder) SetStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatus", reflect.TypeOf((*MockSandboxI)(nil).SetStatus), arg0)
}

// Start mocks base method
func (m *MockSandboxI) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockSandboxIMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockSandboxI)(nil).Start))
}

// Status mocks base method
func (m *MockSandboxI) Status() SandboxStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(SandboxStatus)
	return ret0
}

// Status indicates an expected call of Status
func (mr *MockSandboxIMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockSandboxI)(nil).Status))
}

// TaskBegin mocks base method
func (m *MockSandboxI) TaskBegin() chan bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TaskBegin")
	ret0, _ := ret[0].(chan bool)
	return ret0
}

// TaskBegin indicates an expected call of TaskBegin
func (mr *MockSandboxIMockRecorder) TaskBegin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TaskBegin", reflect.TypeOf((*MockSandboxI)(nil).TaskBegin))
}

// TiedUp mocks base method
func (m *MockSandboxI) TiedUp() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TiedUp")
	ret0, _ := ret[0].(bool)
	return ret0
}

// TiedUp indicates an expected call of TiedUp
func (mr *MockSandboxIMockRecorder) TiedUp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TiedUp", reflect.TypeOf((*MockSandboxI)(nil).TiedUp))
}

// WaitFor mocks base method
func (m *MockSandboxI) WaitFor(arg0 SandboxStatus) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WaitFor", arg0)
}

// WaitFor indicates an expected call of WaitFor
func (mr *MockSandboxIMockRecorder) WaitFor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitFor", reflect.TypeOf((*MockSandboxI)(nil).WaitFor), arg0)
}

// MockDbConn is a mock of DbConn interface
type MockDbConn struct {
	ctrl     *gomock.Controller
	recorder *MockDbConnMockRecorder
}

// MockDbConnMockRecorder is the mock recorder for MockDbConn
type MockDbConnMockRecorder struct {
	mock *MockDbConn
}

// NewMockDbConn creates a new mock instance
func NewMockDbConn(ctrl *gomock.Controller) *MockDbConn {
	mock := &MockDbConn{ctrl: ctrl}
	mock.recorder = &MockDbConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDbConn) EXPECT() *MockDbConnMockRecorder {
	return m.recorder
}

// GetHandle mocks base method
func (m *MockDbConn) GetHandle() *sqlx.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHandle")
	ret0, _ := ret[0].(*sqlx.DB)
	return ret0
}

// GetHandle indicates an expected call of GetHandle
func (mr *MockDbConnMockRecorder) GetHandle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHandle", reflect.TypeOf((*MockDbConn)(nil).GetHandle))
}

// MockAppspaceMetaDB is a mock of AppspaceMetaDB interface
type MockAppspaceMetaDB struct {
	ctrl     *gomock.Controller
	recorder *MockAppspaceMetaDBMockRecorder
}

// MockAppspaceMetaDBMockRecorder is the mock recorder for MockAppspaceMetaDB
type MockAppspaceMetaDBMockRecorder struct {
	mock *MockAppspaceMetaDB
}

// NewMockAppspaceMetaDB creates a new mock instance
func NewMockAppspaceMetaDB(ctrl *gomock.Controller) *MockAppspaceMetaDB {
	mock := &MockAppspaceMetaDB{ctrl: ctrl}
	mock.recorder = &MockAppspaceMetaDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppspaceMetaDB) EXPECT() *MockAppspaceMetaDBMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockAppspaceMetaDB) Create(arg0 AppspaceID, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockAppspaceMetaDBMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAppspaceMetaDB)(nil).Create), arg0, arg1)
}

// GetConn mocks base method
func (m *MockAppspaceMetaDB) GetConn(arg0 AppspaceID) (DbConn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConn", arg0)
	ret0, _ := ret[0].(DbConn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConn indicates an expected call of GetConn
func (mr *MockAppspaceMetaDBMockRecorder) GetConn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConn", reflect.TypeOf((*MockAppspaceMetaDB)(nil).GetConn), arg0)
}

// MockAppspaceInfoModel is a mock of AppspaceInfoModel interface
type MockAppspaceInfoModel struct {
	ctrl     *gomock.Controller
	recorder *MockAppspaceInfoModelMockRecorder
}

// MockAppspaceInfoModelMockRecorder is the mock recorder for MockAppspaceInfoModel
type MockAppspaceInfoModelMockRecorder struct {
	mock *MockAppspaceInfoModel
}

// NewMockAppspaceInfoModel creates a new mock instance
func NewMockAppspaceInfoModel(ctrl *gomock.Controller) *MockAppspaceInfoModel {
	mock := &MockAppspaceInfoModel{ctrl: ctrl}
	mock.recorder = &MockAppspaceInfoModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppspaceInfoModel) EXPECT() *MockAppspaceInfoModelMockRecorder {
	return m.recorder
}

// GetSchema mocks base method
func (m *MockAppspaceInfoModel) GetSchema() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchema")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchema indicates an expected call of GetSchema
func (mr *MockAppspaceInfoModelMockRecorder) GetSchema() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchema", reflect.TypeOf((*MockAppspaceInfoModel)(nil).GetSchema))
}

// SetSchema mocks base method
func (m *MockAppspaceInfoModel) SetSchema(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSchema", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSchema indicates an expected call of SetSchema
func (mr *MockAppspaceInfoModelMockRecorder) SetSchema(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSchema", reflect.TypeOf((*MockAppspaceInfoModel)(nil).SetSchema), arg0)
}

// MockV0RouteModel is a mock of V0RouteModel interface
type MockV0RouteModel struct {
	ctrl     *gomock.Controller
	recorder *MockV0RouteModelMockRecorder
}

// MockV0RouteModelMockRecorder is the mock recorder for MockV0RouteModel
type MockV0RouteModelMockRecorder struct {
	mock *MockV0RouteModel
}

// NewMockV0RouteModel creates a new mock instance
func NewMockV0RouteModel(ctrl *gomock.Controller) *MockV0RouteModel {
	mock := &MockV0RouteModel{ctrl: ctrl}
	mock.recorder = &MockV0RouteModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockV0RouteModel) EXPECT() *MockV0RouteModelMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockV0RouteModel) Create(arg0 []string, arg1 string, arg2 AppspaceRouteAuth, arg3 AppspaceRouteHandler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockV0RouteModelMockRecorder) Create(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockV0RouteModel)(nil).Create), arg0, arg1, arg2, arg3)
}

// Delete mocks base method
func (m *MockV0RouteModel) Delete(arg0 []string, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockV0RouteModelMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockV0RouteModel)(nil).Delete), arg0, arg1)
}

// Get mocks base method
func (m *MockV0RouteModel) Get(arg0 []string, arg1 string) (*[]AppspaceRouteConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*[]AppspaceRouteConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockV0RouteModelMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockV0RouteModel)(nil).Get), arg0, arg1)
}

// GetAll mocks base method
func (m *MockV0RouteModel) GetAll() (*[]AppspaceRouteConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(*[]AppspaceRouteConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockV0RouteModelMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockV0RouteModel)(nil).GetAll))
}

// GetPath mocks base method
func (m *MockV0RouteModel) GetPath(arg0 string) (*[]AppspaceRouteConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPath", arg0)
	ret0, _ := ret[0].(*[]AppspaceRouteConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPath indicates an expected call of GetPath
func (mr *MockV0RouteModelMockRecorder) GetPath(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPath", reflect.TypeOf((*MockV0RouteModel)(nil).GetPath), arg0)
}

// HandleMessage mocks base method
func (m *MockV0RouteModel) HandleMessage(arg0 twine.ReceivedMessageI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleMessage", arg0)
}

// HandleMessage indicates an expected call of HandleMessage
func (mr *MockV0RouteModelMockRecorder) HandleMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleMessage", reflect.TypeOf((*MockV0RouteModel)(nil).HandleMessage), arg0)
}

// Match mocks base method
func (m *MockV0RouteModel) Match(arg0, arg1 string) (*AppspaceRouteConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Match", arg0, arg1)
	ret0, _ := ret[0].(*AppspaceRouteConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Match indicates an expected call of Match
func (mr *MockV0RouteModelMockRecorder) Match(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Match", reflect.TypeOf((*MockV0RouteModel)(nil).Match), arg0, arg1)
}

// MockAppspaceRouteModels is a mock of AppspaceRouteModels interface
type MockAppspaceRouteModels struct {
	ctrl     *gomock.Controller
	recorder *MockAppspaceRouteModelsMockRecorder
}

// MockAppspaceRouteModelsMockRecorder is the mock recorder for MockAppspaceRouteModels
type MockAppspaceRouteModelsMockRecorder struct {
	mock *MockAppspaceRouteModels
}

// NewMockAppspaceRouteModels creates a new mock instance
func NewMockAppspaceRouteModels(ctrl *gomock.Controller) *MockAppspaceRouteModels {
	mock := &MockAppspaceRouteModels{ctrl: ctrl}
	mock.recorder = &MockAppspaceRouteModelsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppspaceRouteModels) EXPECT() *MockAppspaceRouteModelsMockRecorder {
	return m.recorder
}

// GetV0 mocks base method
func (m *MockAppspaceRouteModels) GetV0(arg0 AppspaceID) V0RouteModel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetV0", arg0)
	ret0, _ := ret[0].(V0RouteModel)
	return ret0
}

// GetV0 indicates an expected call of GetV0
func (mr *MockAppspaceRouteModelsMockRecorder) GetV0(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetV0", reflect.TypeOf((*MockAppspaceRouteModels)(nil).GetV0), arg0)
}

// MockStdInput is a mock of StdInput interface
type MockStdInput struct {
	ctrl     *gomock.Controller
	recorder *MockStdInputMockRecorder
}

// MockStdInputMockRecorder is the mock recorder for MockStdInput
type MockStdInputMockRecorder struct {
	mock *MockStdInput
}

// NewMockStdInput creates a new mock instance
func NewMockStdInput(ctrl *gomock.Controller) *MockStdInput {
	mock := &MockStdInput{ctrl: ctrl}
	mock.recorder = &MockStdInputMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStdInput) EXPECT() *MockStdInputMockRecorder {
	return m.recorder
}

// ReadLine mocks base method
func (m *MockStdInput) ReadLine(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadLine", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// ReadLine indicates an expected call of ReadLine
func (mr *MockStdInputMockRecorder) ReadLine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadLine", reflect.TypeOf((*MockStdInput)(nil).ReadLine), arg0)
}
