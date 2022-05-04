// Copyright 2022 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: infra/appengine/weetbix/proto/v1/test_verdict.proto

package weetbixpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Status of a test verdict.
// It is a mirror of luci.resultdb.v1.TestVariantStatus.
type TestVerdictStatus int32

const (
	// a test verdict must not have this status.
	// This is only used when filtering verdicts.
	TestVerdictStatus_TEST_VERDICT_STATUS_UNSPECIFIED TestVerdictStatus = 0
	// The test verdict has no exonerations, and all results are unexpected.
	TestVerdictStatus_UNEXPECTED TestVerdictStatus = 10
	// The test verdict has no exonerations, and all results are unexpectedly skipped.
	TestVerdictStatus_UNEXPECTEDLY_SKIPPED TestVerdictStatus = 20
	// The test verdict has no exonerations, and has both expected and unexpected
	// results.
	TestVerdictStatus_FLAKY TestVerdictStatus = 30
	// The test verdict has one or more test exonerations.
	TestVerdictStatus_EXONERATED TestVerdictStatus = 40
	// The test verdict has no exonerations, and all results are expected.
	TestVerdictStatus_EXPECTED TestVerdictStatus = 50
)

// Enum value maps for TestVerdictStatus.
var (
	TestVerdictStatus_name = map[int32]string{
		0:  "TEST_VERDICT_STATUS_UNSPECIFIED",
		10: "UNEXPECTED",
		20: "UNEXPECTEDLY_SKIPPED",
		30: "FLAKY",
		40: "EXONERATED",
		50: "EXPECTED",
	}
	TestVerdictStatus_value = map[string]int32{
		"TEST_VERDICT_STATUS_UNSPECIFIED": 0,
		"UNEXPECTED":                      10,
		"UNEXPECTEDLY_SKIPPED":            20,
		"FLAKY":                           30,
		"EXONERATED":                      40,
		"EXPECTED":                        50,
	}
)

func (x TestVerdictStatus) Enum() *TestVerdictStatus {
	p := new(TestVerdictStatus)
	*p = x
	return p
}

func (x TestVerdictStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestVerdictStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_appengine_weetbix_proto_v1_test_verdict_proto_enumTypes[0].Descriptor()
}

func (TestVerdictStatus) Type() protoreflect.EnumType {
	return &file_infra_appengine_weetbix_proto_v1_test_verdict_proto_enumTypes[0]
}

func (x TestVerdictStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestVerdictStatus.Descriptor instead.
func (TestVerdictStatus) EnumDescriptor() ([]byte, []int) {
	return file_infra_appengine_weetbix_proto_v1_test_verdict_proto_rawDescGZIP(), []int{0}
}

type TestVerdict struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier of the test.
	// This has the same value as luci.resultdb.v1.TestResult.test_id.
	TestId string `protobuf:"bytes,1,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
	// The hash of the variant.
	VariantHash string `protobuf:"bytes,2,opt,name=variant_hash,json=variantHash,proto3" json:"variant_hash,omitempty"`
	// The ID of the top-level invocation that the test verdict belongs to when
	// ingested.
	InvocationId string `protobuf:"bytes,3,opt,name=invocation_id,json=invocationId,proto3" json:"invocation_id,omitempty"`
	// The status of the test verdict.
	Status TestVerdictStatus `protobuf:"varint,4,opt,name=status,proto3,enum=weetbix.v1.TestVerdictStatus" json:"status,omitempty"`
	// Start time of the presubmit run (for results that are part of a presubmit
	// run) or start time of the buildbucket build (otherwise).
	PartitionTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=partition_time,json=partitionTime,proto3" json:"partition_time,omitempty"`
	// The average duration of the PASSED test results included in the test
	// verdict.
	PassedAvgDuration *durationpb.Duration `protobuf:"bytes,6,opt,name=passed_avg_duration,json=passedAvgDuration,proto3" json:"passed_avg_duration,omitempty"`
}

func (x *TestVerdict) Reset() {
	*x = TestVerdict{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_weetbix_proto_v1_test_verdict_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestVerdict) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestVerdict) ProtoMessage() {}

func (x *TestVerdict) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_weetbix_proto_v1_test_verdict_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestVerdict.ProtoReflect.Descriptor instead.
func (*TestVerdict) Descriptor() ([]byte, []int) {
	return file_infra_appengine_weetbix_proto_v1_test_verdict_proto_rawDescGZIP(), []int{0}
}

func (x *TestVerdict) GetTestId() string {
	if x != nil {
		return x.TestId
	}
	return ""
}

func (x *TestVerdict) GetVariantHash() string {
	if x != nil {
		return x.VariantHash
	}
	return ""
}

func (x *TestVerdict) GetInvocationId() string {
	if x != nil {
		return x.InvocationId
	}
	return ""
}

func (x *TestVerdict) GetStatus() TestVerdictStatus {
	if x != nil {
		return x.Status
	}
	return TestVerdictStatus_TEST_VERDICT_STATUS_UNSPECIFIED
}

func (x *TestVerdict) GetPartitionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.PartitionTime
	}
	return nil
}

func (x *TestVerdict) GetPassedAvgDuration() *durationpb.Duration {
	if x != nil {
		return x.PassedAvgDuration
	}
	return nil
}

var File_infra_appengine_weetbix_proto_v1_test_verdict_proto protoreflect.FileDescriptor

var file_infra_appengine_weetbix_proto_v1_test_verdict_proto_rawDesc = []byte{
	0x0a, 0x33, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x64, 0x69, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76,
	0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb3, 0x02, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x56, 0x65, 0x72, 0x64, 0x69,
	0x63, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x76,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x23,
	0x0a, 0x0d, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x56, 0x65, 0x72, 0x64, 0x69, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x41, 0x0a, 0x0e, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x49, 0x0a,
	0x13, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x76, 0x67, 0x5f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x41, 0x76, 0x67,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x8b, 0x01, 0x0a, 0x11, 0x54, 0x65, 0x73,
	0x74, 0x56, 0x65, 0x72, 0x64, 0x69, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23,
	0x0a, 0x1f, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x56, 0x45, 0x52, 0x44, 0x49, 0x43, 0x54, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x45,
	0x44, 0x10, 0x0a, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x4e, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x45,
	0x44, 0x4c, 0x59, 0x5f, 0x53, 0x4b, 0x49, 0x50, 0x50, 0x45, 0x44, 0x10, 0x14, 0x12, 0x09, 0x0a,
	0x05, 0x46, 0x4c, 0x41, 0x4b, 0x59, 0x10, 0x1e, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x58, 0x4f, 0x4e,
	0x45, 0x52, 0x41, 0x54, 0x45, 0x44, 0x10, 0x28, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x58, 0x50, 0x45,
	0x43, 0x54, 0x45, 0x44, 0x10, 0x32, 0x42, 0x2c, 0x5a, 0x2a, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69,
	0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x65, 0x65, 0x74, 0x62,
	0x69, 0x78, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_appengine_weetbix_proto_v1_test_verdict_proto_rawDescOnce sync.Once
	file_infra_appengine_weetbix_proto_v1_test_verdict_proto_rawDescData = file_infra_appengine_weetbix_proto_v1_test_verdict_proto_rawDesc
)

func file_infra_appengine_weetbix_proto_v1_test_verdict_proto_rawDescGZIP() []byte {
	file_infra_appengine_weetbix_proto_v1_test_verdict_proto_rawDescOnce.Do(func() {
		file_infra_appengine_weetbix_proto_v1_test_verdict_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_appengine_weetbix_proto_v1_test_verdict_proto_rawDescData)
	})
	return file_infra_appengine_weetbix_proto_v1_test_verdict_proto_rawDescData
}

var file_infra_appengine_weetbix_proto_v1_test_verdict_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_infra_appengine_weetbix_proto_v1_test_verdict_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_infra_appengine_weetbix_proto_v1_test_verdict_proto_goTypes = []interface{}{
	(TestVerdictStatus)(0),        // 0: weetbix.v1.TestVerdictStatus
	(*TestVerdict)(nil),           // 1: weetbix.v1.TestVerdict
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 3: google.protobuf.Duration
}
var file_infra_appengine_weetbix_proto_v1_test_verdict_proto_depIdxs = []int32{
	0, // 0: weetbix.v1.TestVerdict.status:type_name -> weetbix.v1.TestVerdictStatus
	2, // 1: weetbix.v1.TestVerdict.partition_time:type_name -> google.protobuf.Timestamp
	3, // 2: weetbix.v1.TestVerdict.passed_avg_duration:type_name -> google.protobuf.Duration
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_infra_appengine_weetbix_proto_v1_test_verdict_proto_init() }
func file_infra_appengine_weetbix_proto_v1_test_verdict_proto_init() {
	if File_infra_appengine_weetbix_proto_v1_test_verdict_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_appengine_weetbix_proto_v1_test_verdict_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestVerdict); i {
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
			RawDescriptor: file_infra_appengine_weetbix_proto_v1_test_verdict_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_appengine_weetbix_proto_v1_test_verdict_proto_goTypes,
		DependencyIndexes: file_infra_appengine_weetbix_proto_v1_test_verdict_proto_depIdxs,
		EnumInfos:         file_infra_appengine_weetbix_proto_v1_test_verdict_proto_enumTypes,
		MessageInfos:      file_infra_appengine_weetbix_proto_v1_test_verdict_proto_msgTypes,
	}.Build()
	File_infra_appengine_weetbix_proto_v1_test_verdict_proto = out.File
	file_infra_appengine_weetbix_proto_v1_test_verdict_proto_rawDesc = nil
	file_infra_appengine_weetbix_proto_v1_test_verdict_proto_goTypes = nil
	file_infra_appengine_weetbix_proto_v1_test_verdict_proto_depIdxs = nil
}
