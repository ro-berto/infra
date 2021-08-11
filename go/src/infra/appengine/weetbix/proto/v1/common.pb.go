// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.0
// source: infra/appengine/weetbix/proto/v1/common.proto

package weetbixpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Status of a Verdict.
// It is determined by all the test results of the verdict, and exonerations are
// ignored(i.e. failure is treated as a failure, even if it is exonerated).
type VerdictStatus int32

const (
	// A verdict must not have this status.
	// This is only used when filtering verdicts.
	VerdictStatus_VERDICT_STATUS_UNSPECIFIED VerdictStatus = 0
	// All results of the verdict are unexpected.
	VerdictStatus_UNEXPECTED VerdictStatus = 10
	// The verdict has both expected and unexpected results.
	// To be differentiated with AnalyzedTestVariantStatus.FLAKY.
	VerdictStatus_VERDICT_FLAKY VerdictStatus = 30
	// All results of the verdict are expected.
	VerdictStatus_EXPECTED VerdictStatus = 50
)

// Enum value maps for VerdictStatus.
var (
	VerdictStatus_name = map[int32]string{
		0:  "VERDICT_STATUS_UNSPECIFIED",
		10: "UNEXPECTED",
		30: "VERDICT_FLAKY",
		50: "EXPECTED",
	}
	VerdictStatus_value = map[string]int32{
		"VERDICT_STATUS_UNSPECIFIED": 0,
		"UNEXPECTED":                 10,
		"VERDICT_FLAKY":              30,
		"EXPECTED":                   50,
	}
)

func (x VerdictStatus) Enum() *VerdictStatus {
	p := new(VerdictStatus)
	*p = x
	return p
}

func (x VerdictStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VerdictStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_appengine_weetbix_proto_v1_common_proto_enumTypes[0].Descriptor()
}

func (VerdictStatus) Type() protoreflect.EnumType {
	return &file_infra_appengine_weetbix_proto_v1_common_proto_enumTypes[0]
}

func (x VerdictStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VerdictStatus.Descriptor instead.
func (VerdictStatus) EnumDescriptor() ([]byte, []int) {
	return file_infra_appengine_weetbix_proto_v1_common_proto_rawDescGZIP(), []int{0}
}

// A string key-value pair.
type StringPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Regex: ^[a-z][a-z0-9_]*(/[a-z][a-z0-9_]*)*$
	// Max length: 64.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Max length: 256.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *StringPair) Reset() {
	*x = StringPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_weetbix_proto_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringPair) ProtoMessage() {}

func (x *StringPair) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_weetbix_proto_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringPair.ProtoReflect.Descriptor instead.
func (*StringPair) Descriptor() ([]byte, []int) {
	return file_infra_appengine_weetbix_proto_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *StringPair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *StringPair) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// A range of timestamps.
type TimeRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The oldest timestamp to include in the range.
	Earliest *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=earliest,proto3" json:"earliest,omitempty"`
	// Include only timestamps that are strictly older than this.
	Latest *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=latest,proto3" json:"latest,omitempty"`
}

func (x *TimeRange) Reset() {
	*x = TimeRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_weetbix_proto_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRange) ProtoMessage() {}

func (x *TimeRange) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_weetbix_proto_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRange.ProtoReflect.Descriptor instead.
func (*TimeRange) Descriptor() ([]byte, []int) {
	return file_infra_appengine_weetbix_proto_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *TimeRange) GetEarliest() *timestamppb.Timestamp {
	if x != nil {
		return x.Earliest
	}
	return nil
}

func (x *TimeRange) GetLatest() *timestamppb.Timestamp {
	if x != nil {
		return x.Latest
	}
	return nil
}

var File_infra_appengine_weetbix_proto_v1_common_proto protoreflect.FileDescriptor

var file_infra_appengine_weetbix_proto_v1_common_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x0a,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x77, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x36, 0x0a, 0x08, 0x65, 0x61, 0x72, 0x6c, 0x69, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x65,
	0x61, 0x72, 0x6c, 0x69, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x06, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x2a, 0x60, 0x0a, 0x0d, 0x56,
	0x65, 0x72, 0x64, 0x69, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x1a,
	0x56, 0x45, 0x52, 0x44, 0x49, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x55, 0x4e, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x11, 0x0a, 0x0d,
	0x56, 0x45, 0x52, 0x44, 0x49, 0x43, 0x54, 0x5f, 0x46, 0x4c, 0x41, 0x4b, 0x59, 0x10, 0x1e, 0x12,
	0x0c, 0x0a, 0x08, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x32, 0x42, 0x2c, 0x5a,
	0x2a, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2f, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x31, 0x3b, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_infra_appengine_weetbix_proto_v1_common_proto_rawDescOnce sync.Once
	file_infra_appengine_weetbix_proto_v1_common_proto_rawDescData = file_infra_appengine_weetbix_proto_v1_common_proto_rawDesc
)

func file_infra_appengine_weetbix_proto_v1_common_proto_rawDescGZIP() []byte {
	file_infra_appengine_weetbix_proto_v1_common_proto_rawDescOnce.Do(func() {
		file_infra_appengine_weetbix_proto_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_appengine_weetbix_proto_v1_common_proto_rawDescData)
	})
	return file_infra_appengine_weetbix_proto_v1_common_proto_rawDescData
}

var file_infra_appengine_weetbix_proto_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_infra_appengine_weetbix_proto_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_infra_appengine_weetbix_proto_v1_common_proto_goTypes = []interface{}{
	(VerdictStatus)(0),            // 0: weetbix.v1.VerdictStatus
	(*StringPair)(nil),            // 1: weetbix.v1.StringPair
	(*TimeRange)(nil),             // 2: weetbix.v1.TimeRange
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_infra_appengine_weetbix_proto_v1_common_proto_depIdxs = []int32{
	3, // 0: weetbix.v1.TimeRange.earliest:type_name -> google.protobuf.Timestamp
	3, // 1: weetbix.v1.TimeRange.latest:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_infra_appengine_weetbix_proto_v1_common_proto_init() }
func file_infra_appengine_weetbix_proto_v1_common_proto_init() {
	if File_infra_appengine_weetbix_proto_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_appengine_weetbix_proto_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringPair); i {
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
		file_infra_appengine_weetbix_proto_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeRange); i {
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
			RawDescriptor: file_infra_appengine_weetbix_proto_v1_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_appengine_weetbix_proto_v1_common_proto_goTypes,
		DependencyIndexes: file_infra_appengine_weetbix_proto_v1_common_proto_depIdxs,
		EnumInfos:         file_infra_appengine_weetbix_proto_v1_common_proto_enumTypes,
		MessageInfos:      file_infra_appengine_weetbix_proto_v1_common_proto_msgTypes,
	}.Build()
	File_infra_appengine_weetbix_proto_v1_common_proto = out.File
	file_infra_appengine_weetbix_proto_v1_common_proto_rawDesc = nil
	file_infra_appengine_weetbix_proto_v1_common_proto_goTypes = nil
	file_infra_appengine_weetbix_proto_v1_common_proto_depIdxs = nil
}
