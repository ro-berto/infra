// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.1
// source: infra/tricium/api/admin/v1/reporter.proto

package admin

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReportResultsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunId    int64  `protobuf:"varint,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	Analyzer string `protobuf:"bytes,2,opt,name=analyzer,proto3" json:"analyzer,omitempty"`
}

func (x *ReportResultsRequest) Reset() {
	*x = ReportResultsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_tricium_api_admin_v1_reporter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportResultsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportResultsRequest) ProtoMessage() {}

func (x *ReportResultsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infra_tricium_api_admin_v1_reporter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportResultsRequest.ProtoReflect.Descriptor instead.
func (*ReportResultsRequest) Descriptor() ([]byte, []int) {
	return file_infra_tricium_api_admin_v1_reporter_proto_rawDescGZIP(), []int{0}
}

func (x *ReportResultsRequest) GetRunId() int64 {
	if x != nil {
		return x.RunId
	}
	return 0
}

func (x *ReportResultsRequest) GetAnalyzer() string {
	if x != nil {
		return x.Analyzer
	}
	return ""
}

type ReportResultsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReportResultsResponse) Reset() {
	*x = ReportResultsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_tricium_api_admin_v1_reporter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportResultsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportResultsResponse) ProtoMessage() {}

func (x *ReportResultsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infra_tricium_api_admin_v1_reporter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportResultsResponse.ProtoReflect.Descriptor instead.
func (*ReportResultsResponse) Descriptor() ([]byte, []int) {
	return file_infra_tricium_api_admin_v1_reporter_proto_rawDescGZIP(), []int{1}
}

var File_infra_tricium_api_admin_v1_reporter_proto protoreflect.FileDescriptor

var file_infra_tricium_api_admin_v1_reporter_proto_rawDesc = []byte{
	0x0a, 0x29, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x22, 0x49, 0x0a, 0x14, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x75,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x22, 0x17, 0x0a,
	0x15, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x56, 0x0a, 0x08, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x12, 0x4a, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x12, 0x1b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x22,
	0x5a, 0x20, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_tricium_api_admin_v1_reporter_proto_rawDescOnce sync.Once
	file_infra_tricium_api_admin_v1_reporter_proto_rawDescData = file_infra_tricium_api_admin_v1_reporter_proto_rawDesc
)

func file_infra_tricium_api_admin_v1_reporter_proto_rawDescGZIP() []byte {
	file_infra_tricium_api_admin_v1_reporter_proto_rawDescOnce.Do(func() {
		file_infra_tricium_api_admin_v1_reporter_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_tricium_api_admin_v1_reporter_proto_rawDescData)
	})
	return file_infra_tricium_api_admin_v1_reporter_proto_rawDescData
}

var file_infra_tricium_api_admin_v1_reporter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_infra_tricium_api_admin_v1_reporter_proto_goTypes = []interface{}{
	(*ReportResultsRequest)(nil),  // 0: admin.ReportResultsRequest
	(*ReportResultsResponse)(nil), // 1: admin.ReportResultsResponse
}
var file_infra_tricium_api_admin_v1_reporter_proto_depIdxs = []int32{
	0, // 0: admin.Reporter.ReportResults:input_type -> admin.ReportResultsRequest
	1, // 1: admin.Reporter.ReportResults:output_type -> admin.ReportResultsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_infra_tricium_api_admin_v1_reporter_proto_init() }
func file_infra_tricium_api_admin_v1_reporter_proto_init() {
	if File_infra_tricium_api_admin_v1_reporter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_tricium_api_admin_v1_reporter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportResultsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_infra_tricium_api_admin_v1_reporter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportResultsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_infra_tricium_api_admin_v1_reporter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infra_tricium_api_admin_v1_reporter_proto_goTypes,
		DependencyIndexes: file_infra_tricium_api_admin_v1_reporter_proto_depIdxs,
		MessageInfos:      file_infra_tricium_api_admin_v1_reporter_proto_msgTypes,
	}.Build()
	File_infra_tricium_api_admin_v1_reporter_proto = out.File
	file_infra_tricium_api_admin_v1_reporter_proto_rawDesc = nil
	file_infra_tricium_api_admin_v1_reporter_proto_goTypes = nil
	file_infra_tricium_api_admin_v1_reporter_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReporterClient is the client API for Reporter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReporterClient interface {
	// ReportResults reports Tricium results.
	ReportResults(ctx context.Context, in *ReportResultsRequest, opts ...grpc.CallOption) (*ReportResultsResponse, error)
}
type reporterPRPCClient struct {
	client *prpc.Client
}

func NewReporterPRPCClient(client *prpc.Client) ReporterClient {
	return &reporterPRPCClient{client}
}

func (c *reporterPRPCClient) ReportResults(ctx context.Context, in *ReportResultsRequest, opts ...grpc.CallOption) (*ReportResultsResponse, error) {
	out := new(ReportResultsResponse)
	err := c.client.Call(ctx, "admin.Reporter", "ReportResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type reporterClient struct {
	cc grpc.ClientConnInterface
}

func NewReporterClient(cc grpc.ClientConnInterface) ReporterClient {
	return &reporterClient{cc}
}

func (c *reporterClient) ReportResults(ctx context.Context, in *ReportResultsRequest, opts ...grpc.CallOption) (*ReportResultsResponse, error) {
	out := new(ReportResultsResponse)
	err := c.cc.Invoke(ctx, "/admin.Reporter/ReportResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReporterServer is the server API for Reporter service.
type ReporterServer interface {
	// ReportResults reports Tricium results.
	ReportResults(context.Context, *ReportResultsRequest) (*ReportResultsResponse, error)
}

// UnimplementedReporterServer can be embedded to have forward compatible implementations.
type UnimplementedReporterServer struct {
}

func (*UnimplementedReporterServer) ReportResults(context.Context, *ReportResultsRequest) (*ReportResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportResults not implemented")
}

func RegisterReporterServer(s prpc.Registrar, srv ReporterServer) {
	s.RegisterService(&_Reporter_serviceDesc, srv)
}

func _Reporter_ReportResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReporterServer).ReportResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.Reporter/ReportResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReporterServer).ReportResults(ctx, req.(*ReportResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Reporter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "admin.Reporter",
	HandlerType: (*ReporterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportResults",
			Handler:    _Reporter_ReportResults_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infra/tricium/api/admin/v1/reporter.proto",
}
