// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/teleclimber/DropServer/cmd/ds-host/testmocks (interfaces: AppspacePausedEvents,AppspaceFilesEvents,AppspaceStatusEvents)

// Package testmocks is a generated GoMock package.
package testmocks

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/teleclimber/DropServer/cmd/ds-host/domain"
	reflect "reflect"
)

// MockAppspacePausedEvents is a mock of AppspacePausedEvents interface
type MockAppspacePausedEvents struct {
	ctrl     *gomock.Controller
	recorder *MockAppspacePausedEventsMockRecorder
}

// MockAppspacePausedEventsMockRecorder is the mock recorder for MockAppspacePausedEvents
type MockAppspacePausedEventsMockRecorder struct {
	mock *MockAppspacePausedEvents
}

// NewMockAppspacePausedEvents creates a new mock instance
func NewMockAppspacePausedEvents(ctrl *gomock.Controller) *MockAppspacePausedEvents {
	mock := &MockAppspacePausedEvents{ctrl: ctrl}
	mock.recorder = &MockAppspacePausedEventsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppspacePausedEvents) EXPECT() *MockAppspacePausedEventsMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockAppspacePausedEvents) Send(arg0 domain.AppspaceID, arg1 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Send", arg0, arg1)
}

// Send indicates an expected call of Send
func (mr *MockAppspacePausedEventsMockRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockAppspacePausedEvents)(nil).Send), arg0, arg1)
}

// Subscribe mocks base method
func (m *MockAppspacePausedEvents) Subscribe(arg0 chan<- domain.AppspacePausedEvent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Subscribe", arg0)
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockAppspacePausedEventsMockRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockAppspacePausedEvents)(nil).Subscribe), arg0)
}

// Unsubscribe mocks base method
func (m *MockAppspacePausedEvents) Unsubscribe(arg0 chan<- domain.AppspacePausedEvent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unsubscribe", arg0)
}

// Unsubscribe indicates an expected call of Unsubscribe
func (mr *MockAppspacePausedEventsMockRecorder) Unsubscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockAppspacePausedEvents)(nil).Unsubscribe), arg0)
}

// MockAppspaceFilesEvents is a mock of AppspaceFilesEvents interface
type MockAppspaceFilesEvents struct {
	ctrl     *gomock.Controller
	recorder *MockAppspaceFilesEventsMockRecorder
}

// MockAppspaceFilesEventsMockRecorder is the mock recorder for MockAppspaceFilesEvents
type MockAppspaceFilesEventsMockRecorder struct {
	mock *MockAppspaceFilesEvents
}

// NewMockAppspaceFilesEvents creates a new mock instance
func NewMockAppspaceFilesEvents(ctrl *gomock.Controller) *MockAppspaceFilesEvents {
	mock := &MockAppspaceFilesEvents{ctrl: ctrl}
	mock.recorder = &MockAppspaceFilesEventsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppspaceFilesEvents) EXPECT() *MockAppspaceFilesEventsMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockAppspaceFilesEvents) Send(arg0 domain.AppspaceID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Send", arg0)
}

// Send indicates an expected call of Send
func (mr *MockAppspaceFilesEventsMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockAppspaceFilesEvents)(nil).Send), arg0)
}

// Subscribe mocks base method
func (m *MockAppspaceFilesEvents) Subscribe(arg0 chan<- domain.AppspaceID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Subscribe", arg0)
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockAppspaceFilesEventsMockRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockAppspaceFilesEvents)(nil).Subscribe), arg0)
}

// Unsubscribe mocks base method
func (m *MockAppspaceFilesEvents) Unsubscribe(arg0 chan<- domain.AppspaceID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unsubscribe", arg0)
}

// Unsubscribe indicates an expected call of Unsubscribe
func (mr *MockAppspaceFilesEventsMockRecorder) Unsubscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockAppspaceFilesEvents)(nil).Unsubscribe), arg0)
}

// MockAppspaceStatusEvents is a mock of AppspaceStatusEvents interface
type MockAppspaceStatusEvents struct {
	ctrl     *gomock.Controller
	recorder *MockAppspaceStatusEventsMockRecorder
}

// MockAppspaceStatusEventsMockRecorder is the mock recorder for MockAppspaceStatusEvents
type MockAppspaceStatusEventsMockRecorder struct {
	mock *MockAppspaceStatusEvents
}

// NewMockAppspaceStatusEvents creates a new mock instance
func NewMockAppspaceStatusEvents(ctrl *gomock.Controller) *MockAppspaceStatusEvents {
	mock := &MockAppspaceStatusEvents{ctrl: ctrl}
	mock.recorder = &MockAppspaceStatusEventsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppspaceStatusEvents) EXPECT() *MockAppspaceStatusEventsMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockAppspaceStatusEvents) Send(arg0 domain.AppspaceID, arg1 domain.AppspaceStatusEvent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Send", arg0, arg1)
}

// Send indicates an expected call of Send
func (mr *MockAppspaceStatusEventsMockRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockAppspaceStatusEvents)(nil).Send), arg0, arg1)
}

// Subscribe mocks base method
func (m *MockAppspaceStatusEvents) Subscribe(arg0 domain.AppspaceID, arg1 chan<- domain.AppspaceStatusEvent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Subscribe", arg0, arg1)
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockAppspaceStatusEventsMockRecorder) Subscribe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockAppspaceStatusEvents)(nil).Subscribe), arg0, arg1)
}

// Unsubscribe mocks base method
func (m *MockAppspaceStatusEvents) Unsubscribe(arg0 domain.AppspaceID, arg1 chan<- domain.AppspaceStatusEvent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unsubscribe", arg0, arg1)
}

// Unsubscribe indicates an expected call of Unsubscribe
func (mr *MockAppspaceStatusEventsMockRecorder) Unsubscribe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockAppspaceStatusEvents)(nil).Unsubscribe), arg0, arg1)
}

// UnsubscribeChannel mocks base method
func (m *MockAppspaceStatusEvents) UnsubscribeChannel(arg0 chan<- domain.AppspaceStatusEvent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnsubscribeChannel", arg0)
}

// UnsubscribeChannel indicates an expected call of UnsubscribeChannel
func (mr *MockAppspaceStatusEventsMockRecorder) UnsubscribeChannel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribeChannel", reflect.TypeOf((*MockAppspaceStatusEvents)(nil).UnsubscribeChannel), arg0)
}
