// Code generated by MockGen. DO NOT EDIT.
// Source: /usr/local/google/home/azrahman/chromiumos/infra/infra/go/src/go.chromium.org/chromiumos/config/go/test/api/provision_grpc.pb.go

// Package mocked_services is a generated GoMock package.
package mocked_services

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	longrunning "go.chromium.org/chromiumos/config/go/longrunning"
	api "go.chromium.org/chromiumos/config/go/test/api"
	grpc "google.golang.org/grpc"
)

// MockGenericProvisionServiceClient is a mock of GenericProvisionServiceClient interface.
type MockGenericProvisionServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockGenericProvisionServiceClientMockRecorder
}

// MockGenericProvisionServiceClientMockRecorder is the mock recorder for MockGenericProvisionServiceClient.
type MockGenericProvisionServiceClientMockRecorder struct {
	mock *MockGenericProvisionServiceClient
}

// NewMockGenericProvisionServiceClient creates a new mock instance.
func NewMockGenericProvisionServiceClient(ctrl *gomock.Controller) *MockGenericProvisionServiceClient {
	mock := &MockGenericProvisionServiceClient{ctrl: ctrl}
	mock.recorder = &MockGenericProvisionServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenericProvisionServiceClient) EXPECT() *MockGenericProvisionServiceClientMockRecorder {
	return m.recorder
}

// Install mocks base method.
func (m *MockGenericProvisionServiceClient) Install(ctx context.Context, in *api.InstallRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Install", varargs...)
	ret0, _ := ret[0].(*longrunning.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Install indicates an expected call of Install.
func (mr *MockGenericProvisionServiceClientMockRecorder) Install(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockGenericProvisionServiceClient)(nil).Install), varargs...)
}

// MockGenericProvisionServiceServer is a mock of GenericProvisionServiceServer interface.
type MockGenericProvisionServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockGenericProvisionServiceServerMockRecorder
}

// MockGenericProvisionServiceServerMockRecorder is the mock recorder for MockGenericProvisionServiceServer.
type MockGenericProvisionServiceServerMockRecorder struct {
	mock *MockGenericProvisionServiceServer
}

// NewMockGenericProvisionServiceServer creates a new mock instance.
func NewMockGenericProvisionServiceServer(ctrl *gomock.Controller) *MockGenericProvisionServiceServer {
	mock := &MockGenericProvisionServiceServer{ctrl: ctrl}
	mock.recorder = &MockGenericProvisionServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenericProvisionServiceServer) EXPECT() *MockGenericProvisionServiceServerMockRecorder {
	return m.recorder
}

// Install mocks base method.
func (m *MockGenericProvisionServiceServer) Install(arg0 context.Context, arg1 *api.InstallRequest) (*longrunning.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install", arg0, arg1)
	ret0, _ := ret[0].(*longrunning.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Install indicates an expected call of Install.
func (mr *MockGenericProvisionServiceServerMockRecorder) Install(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockGenericProvisionServiceServer)(nil).Install), arg0, arg1)
}

// MockUnsafeGenericProvisionServiceServer is a mock of UnsafeGenericProvisionServiceServer interface.
type MockUnsafeGenericProvisionServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeGenericProvisionServiceServerMockRecorder
}

// MockUnsafeGenericProvisionServiceServerMockRecorder is the mock recorder for MockUnsafeGenericProvisionServiceServer.
type MockUnsafeGenericProvisionServiceServerMockRecorder struct {
	mock *MockUnsafeGenericProvisionServiceServer
}

// NewMockUnsafeGenericProvisionServiceServer creates a new mock instance.
func NewMockUnsafeGenericProvisionServiceServer(ctrl *gomock.Controller) *MockUnsafeGenericProvisionServiceServer {
	mock := &MockUnsafeGenericProvisionServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeGenericProvisionServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeGenericProvisionServiceServer) EXPECT() *MockUnsafeGenericProvisionServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedGenericProvisionServiceServer mocks base method.
func (m *MockUnsafeGenericProvisionServiceServer) mustEmbedUnimplementedGenericProvisionServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedGenericProvisionServiceServer")
}

// mustEmbedUnimplementedGenericProvisionServiceServer indicates an expected call of mustEmbedUnimplementedGenericProvisionServiceServer.
func (mr *MockUnsafeGenericProvisionServiceServerMockRecorder) mustEmbedUnimplementedGenericProvisionServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedGenericProvisionServiceServer", reflect.TypeOf((*MockUnsafeGenericProvisionServiceServer)(nil).mustEmbedUnimplementedGenericProvisionServiceServer))
}