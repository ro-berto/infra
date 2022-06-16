// Copyright 2022 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: src/infra/rts/cmd/rts-ml-chromium/proto/rts-ml-chromium.proto

package proto

import (
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

// The stability information for a test
type Stability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of times the test failed in six months
	SixMonthFailCount int64 `protobuf:"varint,1,opt,name=six_month_fail_count,json=sixMonthFailCount,proto3" json:"six_month_fail_count,omitempty"`
	// The number of times the ran in six months
	SixMonthRunCount int64 `protobuf:"varint,2,opt,name=six_month_run_count,json=sixMonthRunCount,proto3" json:"six_month_run_count,omitempty"`
	// The number of times the test failed in one month
	OneMonthFailCount int64 `protobuf:"varint,3,opt,name=one_month_fail_count,json=oneMonthFailCount,proto3" json:"one_month_fail_count,omitempty"`
	// The number of times the ran in one month
	OneMonthRunCount int64 `protobuf:"varint,4,opt,name=one_month_run_count,json=oneMonthRunCount,proto3" json:"one_month_run_count,omitempty"`
	// The number of times the test failed in one week
	OneWeekFailCount int64 `protobuf:"varint,5,opt,name=oneWeek_fail_count,json=oneWeekFailCount,proto3" json:"oneWeek_fail_count,omitempty"`
	// The number of times the ran in one week
	OneWeekRunCount int64 `protobuf:"varint,6,opt,name=oneWeek_run_count,json=oneWeekRunCount,proto3" json:"oneWeek_run_count,omitempty"`
}

func (x *Stability) Reset() {
	*x = Stability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stability) ProtoMessage() {}

func (x *Stability) ProtoReflect() protoreflect.Message {
	mi := &file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stability.ProtoReflect.Descriptor instead.
func (*Stability) Descriptor() ([]byte, []int) {
	return file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_rawDescGZIP(), []int{0}
}

func (x *Stability) GetSixMonthFailCount() int64 {
	if x != nil {
		return x.SixMonthFailCount
	}
	return 0
}

func (x *Stability) GetSixMonthRunCount() int64 {
	if x != nil {
		return x.SixMonthRunCount
	}
	return 0
}

func (x *Stability) GetOneMonthFailCount() int64 {
	if x != nil {
		return x.OneMonthFailCount
	}
	return 0
}

func (x *Stability) GetOneMonthRunCount() int64 {
	if x != nil {
		return x.OneMonthRunCount
	}
	return 0
}

func (x *Stability) GetOneWeekFailCount() int64 {
	if x != nil {
		return x.OneWeekFailCount
	}
	return 0
}

func (x *Stability) GetOneWeekRunCount() int64 {
	if x != nil {
		return x.OneWeekRunCount
	}
	return 0
}

// The number of times the test failed in six months
type TestStability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The test id of the test the stability belongs to
	TestId string `protobuf:"bytes,1,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
	// The test name of the test the stability belongs to (use for filtering)
	TestName string `protobuf:"bytes,2,opt,name=test_name,json=testName,proto3" json:"test_name,omitempty"`
	// The builder the test was run on
	Builder string `protobuf:"bytes,3,opt,name=builder,proto3" json:"builder,omitempty"`
	// The test suite the tes was run in
	TestSuite string `protobuf:"bytes,4,opt,name=test_suite,json=testSuite,proto3" json:"test_suite,omitempty"`
	// All the stability information about the test
	Stability *Stability `protobuf:"bytes,5,opt,name=stability,proto3" json:"stability,omitempty"`
}

func (x *TestStability) Reset() {
	*x = TestStability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestStability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestStability) ProtoMessage() {}

func (x *TestStability) ProtoReflect() protoreflect.Message {
	mi := &file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestStability.ProtoReflect.Descriptor instead.
func (*TestStability) Descriptor() ([]byte, []int) {
	return file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_rawDescGZIP(), []int{1}
}

func (x *TestStability) GetTestId() string {
	if x != nil {
		return x.TestId
	}
	return ""
}

func (x *TestStability) GetTestName() string {
	if x != nil {
		return x.TestName
	}
	return ""
}

func (x *TestStability) GetBuilder() string {
	if x != nil {
		return x.Builder
	}
	return ""
}

func (x *TestStability) GetTestSuite() string {
	if x != nil {
		return x.TestSuite
	}
	return ""
}

func (x *TestStability) GetStability() *Stability {
	if x != nil {
		return x.Stability
	}
	return nil
}

var File_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto protoreflect.FileDescriptor

var file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x73, 0x72, 0x63, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x72, 0x74, 0x73, 0x2f,
	0x63, 0x6d, 0x64, 0x2f, 0x72, 0x74, 0x73, 0x2d, 0x6d, 0x6c, 0x2d, 0x63, 0x68, 0x72, 0x6f, 0x6d,
	0x69, 0x75, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x74, 0x73, 0x2d, 0x6d, 0x6c,
	0x2d, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0xa5, 0x02, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x12, 0x2f, 0x0a, 0x14, 0x73, 0x69, 0x78, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x5f, 0x66, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x11, 0x73, 0x69, 0x78, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x46, 0x61, 0x69, 0x6c, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x13, 0x73, 0x69, 0x78, 0x5f, 0x6d, 0x6f, 0x6e, 0x74,
	0x68, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x10, 0x73, 0x69, 0x78, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x52, 0x75, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x14, 0x6f, 0x6e, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x5f, 0x66, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x11, 0x6f, 0x6e, 0x65, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x46, 0x61, 0x69, 0x6c, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x13, 0x6f, 0x6e, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x74,
	0x68, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x10, 0x6f, 0x6e, 0x65, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x52, 0x75, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x6f, 0x6e, 0x65, 0x57, 0x65, 0x65, 0x6b, 0x5f, 0x66,
	0x61, 0x69, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x10, 0x6f, 0x6e, 0x65, 0x57, 0x65, 0x65, 0x6b, 0x46, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2a, 0x0a, 0x11, 0x6f, 0x6e, 0x65, 0x57, 0x65, 0x65, 0x6b, 0x5f, 0x72, 0x75, 0x6e,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6f, 0x6e,
	0x65, 0x57, 0x65, 0x65, 0x6b, 0x52, 0x75, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xad, 0x01,
	0x0a, 0x0d, 0x54, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x12,
	0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x73, 0x74, 0x53, 0x75, 0x69, 0x74, 0x65, 0x12, 0x2d,
	0x0a, 0x09, 0x73, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x09, 0x73, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x42, 0x25, 0x5a,
	0x23, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x72, 0x74, 0x73, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x72,
	0x74, 0x73, 0x2d, 0x6d, 0x6c, 0x2d, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_rawDescOnce sync.Once
	file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_rawDescData = file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_rawDesc
)

func file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_rawDescGZIP() []byte {
	file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_rawDescOnce.Do(func() {
		file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_rawDescData)
	})
	return file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_rawDescData
}

var file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_goTypes = []interface{}{
	(*Stability)(nil),     // 0: main.Stability
	(*TestStability)(nil), // 1: main.TestStability
}
var file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_depIdxs = []int32{
	0, // 0: main.TestStability.stability:type_name -> main.Stability
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_init() }
func file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_init() {
	if File_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stability); i {
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
		file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestStability); i {
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
			RawDescriptor: file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_goTypes,
		DependencyIndexes: file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_depIdxs,
		MessageInfos:      file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_msgTypes,
	}.Build()
	File_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto = out.File
	file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_rawDesc = nil
	file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_goTypes = nil
	file_src_infra_rts_cmd_rts_ml_chromium_proto_rts_ml_chromium_proto_depIdxs = nil
}
