// Code generated by MockGen. DO NOT EDIT.
// Source: components.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	xmlrpc "go.chromium.org/chromiumos/config/go/api/test/xmlrpc"
)

// MockServod is a mock of Servod interface.
type MockServod struct {
	ctrl     *gomock.Controller
	recorder *MockServodMockRecorder
}

// MockServodMockRecorder is the mock recorder for MockServod.
type MockServodMockRecorder struct {
	mock *MockServod
}

// NewMockServod creates a new mock instance.
func NewMockServod(ctrl *gomock.Controller) *MockServod {
	mock := &MockServod{ctrl: ctrl}
	mock.recorder = &MockServodMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServod) EXPECT() *MockServodMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockServod) Get(ctx context.Context, cmd string) (*xmlrpc.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, cmd)
	ret0, _ := ret[0].(*xmlrpc.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockServodMockRecorder) Get(ctx, cmd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockServod)(nil).Get), ctx, cmd)
}

// Port mocks base method.
func (m *MockServod) Port() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Port")
	ret0, _ := ret[0].(int)
	return ret0
}

// Port indicates an expected call of Port.
func (mr *MockServodMockRecorder) Port() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Port", reflect.TypeOf((*MockServod)(nil).Port))
}

// Set mocks base method.
func (m *MockServod) Set(ctx context.Context, cmd string, val *xmlrpc.Value) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, cmd, val)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockServodMockRecorder) Set(ctx, cmd, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockServod)(nil).Set), ctx, cmd, val)
}
