// Code generated by MockGen. DO NOT EDIT.
// Source: inventory.pb.go

// Package fleet is a generated GoMock package.
package fleet

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockisSetSatlabStableVersionRequest_Strategy is a mock of isSetSatlabStableVersionRequest_Strategy interface.
type MockisSetSatlabStableVersionRequest_Strategy struct {
	ctrl     *gomock.Controller
	recorder *MockisSetSatlabStableVersionRequest_StrategyMockRecorder
}

// MockisSetSatlabStableVersionRequest_StrategyMockRecorder is the mock recorder for MockisSetSatlabStableVersionRequest_Strategy.
type MockisSetSatlabStableVersionRequest_StrategyMockRecorder struct {
	mock *MockisSetSatlabStableVersionRequest_Strategy
}

// NewMockisSetSatlabStableVersionRequest_Strategy creates a new mock instance.
func NewMockisSetSatlabStableVersionRequest_Strategy(ctrl *gomock.Controller) *MockisSetSatlabStableVersionRequest_Strategy {
	mock := &MockisSetSatlabStableVersionRequest_Strategy{ctrl: ctrl}
	mock.recorder = &MockisSetSatlabStableVersionRequest_StrategyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockisSetSatlabStableVersionRequest_Strategy) EXPECT() *MockisSetSatlabStableVersionRequest_StrategyMockRecorder {
	return m.recorder
}

// isSetSatlabStableVersionRequest_Strategy mocks base method.
func (m *MockisSetSatlabStableVersionRequest_Strategy) isSetSatlabStableVersionRequest_Strategy() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isSetSatlabStableVersionRequest_Strategy")
}

// isSetSatlabStableVersionRequest_Strategy indicates an expected call of isSetSatlabStableVersionRequest_Strategy.
func (mr *MockisSetSatlabStableVersionRequest_StrategyMockRecorder) isSetSatlabStableVersionRequest_Strategy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isSetSatlabStableVersionRequest_Strategy", reflect.TypeOf((*MockisSetSatlabStableVersionRequest_Strategy)(nil).isSetSatlabStableVersionRequest_Strategy))
}

// MockisDeleteSatlabStableVersionRequest_Strategy is a mock of isDeleteSatlabStableVersionRequest_Strategy interface.
type MockisDeleteSatlabStableVersionRequest_Strategy struct {
	ctrl     *gomock.Controller
	recorder *MockisDeleteSatlabStableVersionRequest_StrategyMockRecorder
}

// MockisDeleteSatlabStableVersionRequest_StrategyMockRecorder is the mock recorder for MockisDeleteSatlabStableVersionRequest_Strategy.
type MockisDeleteSatlabStableVersionRequest_StrategyMockRecorder struct {
	mock *MockisDeleteSatlabStableVersionRequest_Strategy
}

// NewMockisDeleteSatlabStableVersionRequest_Strategy creates a new mock instance.
func NewMockisDeleteSatlabStableVersionRequest_Strategy(ctrl *gomock.Controller) *MockisDeleteSatlabStableVersionRequest_Strategy {
	mock := &MockisDeleteSatlabStableVersionRequest_Strategy{ctrl: ctrl}
	mock.recorder = &MockisDeleteSatlabStableVersionRequest_StrategyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockisDeleteSatlabStableVersionRequest_Strategy) EXPECT() *MockisDeleteSatlabStableVersionRequest_StrategyMockRecorder {
	return m.recorder
}

// isDeleteSatlabStableVersionRequest_Strategy mocks base method.
func (m *MockisDeleteSatlabStableVersionRequest_Strategy) isDeleteSatlabStableVersionRequest_Strategy() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isDeleteSatlabStableVersionRequest_Strategy")
}

// isDeleteSatlabStableVersionRequest_Strategy indicates an expected call of isDeleteSatlabStableVersionRequest_Strategy.
func (mr *MockisDeleteSatlabStableVersionRequest_StrategyMockRecorder) isDeleteSatlabStableVersionRequest_Strategy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isDeleteSatlabStableVersionRequest_Strategy", reflect.TypeOf((*MockisDeleteSatlabStableVersionRequest_Strategy)(nil).isDeleteSatlabStableVersionRequest_Strategy))
}

// MockInventoryClient is a mock of InventoryClient interface.
type MockInventoryClient struct {
	ctrl     *gomock.Controller
	recorder *MockInventoryClientMockRecorder
}

// MockInventoryClientMockRecorder is the mock recorder for MockInventoryClient.
type MockInventoryClientMockRecorder struct {
	mock *MockInventoryClient
}

// NewMockInventoryClient creates a new mock instance.
func NewMockInventoryClient(ctrl *gomock.Controller) *MockInventoryClient {
	mock := &MockInventoryClient{ctrl: ctrl}
	mock.recorder = &MockInventoryClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInventoryClient) EXPECT() *MockInventoryClientMockRecorder {
	return m.recorder
}

// DeleteSatlabStableVersion mocks base method.
func (m *MockInventoryClient) DeleteSatlabStableVersion(ctx context.Context, in *DeleteSatlabStableVersionRequest, opts ...grpc.CallOption) (*DeleteSatlabStableVersionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteSatlabStableVersion", varargs...)
	ret0, _ := ret[0].(*DeleteSatlabStableVersionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSatlabStableVersion indicates an expected call of DeleteSatlabStableVersion.
func (mr *MockInventoryClientMockRecorder) DeleteSatlabStableVersion(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSatlabStableVersion", reflect.TypeOf((*MockInventoryClient)(nil).DeleteSatlabStableVersion), varargs...)
}

// DumpStableVersionToDatastore mocks base method.
func (m *MockInventoryClient) DumpStableVersionToDatastore(ctx context.Context, in *DumpStableVersionToDatastoreRequest, opts ...grpc.CallOption) (*DumpStableVersionToDatastoreResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DumpStableVersionToDatastore", varargs...)
	ret0, _ := ret[0].(*DumpStableVersionToDatastoreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DumpStableVersionToDatastore indicates an expected call of DumpStableVersionToDatastore.
func (mr *MockInventoryClientMockRecorder) DumpStableVersionToDatastore(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DumpStableVersionToDatastore", reflect.TypeOf((*MockInventoryClient)(nil).DumpStableVersionToDatastore), varargs...)
}

// GetStableVersion mocks base method.
func (m *MockInventoryClient) GetStableVersion(ctx context.Context, in *GetStableVersionRequest, opts ...grpc.CallOption) (*GetStableVersionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStableVersion", varargs...)
	ret0, _ := ret[0].(*GetStableVersionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStableVersion indicates an expected call of GetStableVersion.
func (mr *MockInventoryClientMockRecorder) GetStableVersion(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStableVersion", reflect.TypeOf((*MockInventoryClient)(nil).GetStableVersion), varargs...)
}

// SetSatlabStableVersion mocks base method.
func (m *MockInventoryClient) SetSatlabStableVersion(ctx context.Context, in *SetSatlabStableVersionRequest, opts ...grpc.CallOption) (*SetSatlabStableVersionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetSatlabStableVersion", varargs...)
	ret0, _ := ret[0].(*SetSatlabStableVersionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetSatlabStableVersion indicates an expected call of SetSatlabStableVersion.
func (mr *MockInventoryClientMockRecorder) SetSatlabStableVersion(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSatlabStableVersion", reflect.TypeOf((*MockInventoryClient)(nil).SetSatlabStableVersion), varargs...)
}

// MockInventoryServer is a mock of InventoryServer interface.
type MockInventoryServer struct {
	ctrl     *gomock.Controller
	recorder *MockInventoryServerMockRecorder
}

// MockInventoryServerMockRecorder is the mock recorder for MockInventoryServer.
type MockInventoryServerMockRecorder struct {
	mock *MockInventoryServer
}

// NewMockInventoryServer creates a new mock instance.
func NewMockInventoryServer(ctrl *gomock.Controller) *MockInventoryServer {
	mock := &MockInventoryServer{ctrl: ctrl}
	mock.recorder = &MockInventoryServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInventoryServer) EXPECT() *MockInventoryServerMockRecorder {
	return m.recorder
}

// DeleteSatlabStableVersion mocks base method.
func (m *MockInventoryServer) DeleteSatlabStableVersion(arg0 context.Context, arg1 *DeleteSatlabStableVersionRequest) (*DeleteSatlabStableVersionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSatlabStableVersion", arg0, arg1)
	ret0, _ := ret[0].(*DeleteSatlabStableVersionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSatlabStableVersion indicates an expected call of DeleteSatlabStableVersion.
func (mr *MockInventoryServerMockRecorder) DeleteSatlabStableVersion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSatlabStableVersion", reflect.TypeOf((*MockInventoryServer)(nil).DeleteSatlabStableVersion), arg0, arg1)
}

// DumpStableVersionToDatastore mocks base method.
func (m *MockInventoryServer) DumpStableVersionToDatastore(arg0 context.Context, arg1 *DumpStableVersionToDatastoreRequest) (*DumpStableVersionToDatastoreResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DumpStableVersionToDatastore", arg0, arg1)
	ret0, _ := ret[0].(*DumpStableVersionToDatastoreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DumpStableVersionToDatastore indicates an expected call of DumpStableVersionToDatastore.
func (mr *MockInventoryServerMockRecorder) DumpStableVersionToDatastore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DumpStableVersionToDatastore", reflect.TypeOf((*MockInventoryServer)(nil).DumpStableVersionToDatastore), arg0, arg1)
}

// GetStableVersion mocks base method.
func (m *MockInventoryServer) GetStableVersion(arg0 context.Context, arg1 *GetStableVersionRequest) (*GetStableVersionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStableVersion", arg0, arg1)
	ret0, _ := ret[0].(*GetStableVersionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStableVersion indicates an expected call of GetStableVersion.
func (mr *MockInventoryServerMockRecorder) GetStableVersion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStableVersion", reflect.TypeOf((*MockInventoryServer)(nil).GetStableVersion), arg0, arg1)
}

// SetSatlabStableVersion mocks base method.
func (m *MockInventoryServer) SetSatlabStableVersion(arg0 context.Context, arg1 *SetSatlabStableVersionRequest) (*SetSatlabStableVersionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSatlabStableVersion", arg0, arg1)
	ret0, _ := ret[0].(*SetSatlabStableVersionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetSatlabStableVersion indicates an expected call of SetSatlabStableVersion.
func (mr *MockInventoryServerMockRecorder) SetSatlabStableVersion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSatlabStableVersion", reflect.TypeOf((*MockInventoryServer)(nil).SetSatlabStableVersion), arg0, arg1)
}
