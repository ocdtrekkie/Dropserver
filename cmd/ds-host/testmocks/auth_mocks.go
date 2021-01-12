// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/teleclimber/DropServer/cmd/ds-host/testmocks (interfaces: Authenticator,AppspaceLogin)

// Package testmocks is a generated GoMock package.
package testmocks

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/teleclimber/DropServer/cmd/ds-host/domain"
	http "net/http"
	url "net/url"
	reflect "reflect"
)

// MockAuthenticator is a mock of Authenticator interface
type MockAuthenticator struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticatorMockRecorder
}

// MockAuthenticatorMockRecorder is the mock recorder for MockAuthenticator
type MockAuthenticatorMockRecorder struct {
	mock *MockAuthenticator
}

// NewMockAuthenticator creates a new mock instance
func NewMockAuthenticator(ctrl *gomock.Controller) *MockAuthenticator {
	mock := &MockAuthenticator{ctrl: ctrl}
	mock.recorder = &MockAuthenticatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthenticator) EXPECT() *MockAuthenticatorMockRecorder {
	return m.recorder
}

// Authenticate mocks base method
func (m *MockAuthenticator) Authenticate(arg0 *http.Request) domain.Authentication {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", arg0)
	ret0, _ := ret[0].(domain.Authentication)
	return ret0
}

// Authenticate indicates an expected call of Authenticate
func (mr *MockAuthenticatorMockRecorder) Authenticate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockAuthenticator)(nil).Authenticate), arg0)
}

// SetForAccount mocks base method
func (m *MockAuthenticator) SetForAccount(arg0 http.ResponseWriter, arg1 domain.UserID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetForAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetForAccount indicates an expected call of SetForAccount
func (mr *MockAuthenticatorMockRecorder) SetForAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetForAccount", reflect.TypeOf((*MockAuthenticator)(nil).SetForAccount), arg0, arg1)
}

// SetForAppspace mocks base method
func (m *MockAuthenticator) SetForAppspace(arg0 http.ResponseWriter, arg1 domain.ProxyID, arg2 domain.AppspaceID, arg3 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetForAppspace", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetForAppspace indicates an expected call of SetForAppspace
func (mr *MockAuthenticatorMockRecorder) SetForAppspace(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetForAppspace", reflect.TypeOf((*MockAuthenticator)(nil).SetForAppspace), arg0, arg1, arg2, arg3)
}

// UnsetForAccount mocks base method
func (m *MockAuthenticator) UnsetForAccount(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnsetForAccount", arg0, arg1)
}

// UnsetForAccount indicates an expected call of UnsetForAccount
func (mr *MockAuthenticatorMockRecorder) UnsetForAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsetForAccount", reflect.TypeOf((*MockAuthenticator)(nil).UnsetForAccount), arg0, arg1)
}

// MockAppspaceLogin is a mock of AppspaceLogin interface
type MockAppspaceLogin struct {
	ctrl     *gomock.Controller
	recorder *MockAppspaceLoginMockRecorder
}

// MockAppspaceLoginMockRecorder is the mock recorder for MockAppspaceLogin
type MockAppspaceLoginMockRecorder struct {
	mock *MockAppspaceLogin
}

// NewMockAppspaceLogin creates a new mock instance
func NewMockAppspaceLogin(ctrl *gomock.Controller) *MockAppspaceLogin {
	mock := &MockAppspaceLogin{ctrl: ctrl}
	mock.recorder = &MockAppspaceLoginMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppspaceLogin) EXPECT() *MockAppspaceLoginMockRecorder {
	return m.recorder
}

// CheckRedirectToken mocks base method
func (m *MockAppspaceLogin) CheckRedirectToken(arg0 string) (domain.AppspaceLoginToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRedirectToken", arg0)
	ret0, _ := ret[0].(domain.AppspaceLoginToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckRedirectToken indicates an expected call of CheckRedirectToken
func (mr *MockAppspaceLoginMockRecorder) CheckRedirectToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRedirectToken", reflect.TypeOf((*MockAppspaceLogin)(nil).CheckRedirectToken), arg0)
}

// Create mocks base method
func (m *MockAppspaceLogin) Create(arg0 domain.AppspaceID, arg1 url.URL) domain.AppspaceLoginToken {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(domain.AppspaceLoginToken)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockAppspaceLoginMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAppspaceLogin)(nil).Create), arg0, arg1)
}

// LogIn mocks base method
func (m *MockAppspaceLogin) LogIn(arg0 string, arg1 domain.UserID) (domain.AppspaceLoginToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogIn", arg0, arg1)
	ret0, _ := ret[0].(domain.AppspaceLoginToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogIn indicates an expected call of LogIn
func (mr *MockAppspaceLoginMockRecorder) LogIn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogIn", reflect.TypeOf((*MockAppspaceLogin)(nil).LogIn), arg0, arg1)
}
