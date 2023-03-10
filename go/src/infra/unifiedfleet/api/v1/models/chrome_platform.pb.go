// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: infra/unifiedfleet/api/v1/models/chrome_platform.proto

package ufspb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// ChromePlatform refers to static configs defined by Chrome Browser lab
// for their machines
type ChromePlatform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The format will be chromePlatforms/XXX
	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Manufacturer string `protobuf:"bytes,2,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	Description  string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Record the last update timestamp of this ChromePlatform (In UTC timezone)
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// tags user can attach for easy querying/search
	Tags []string `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *ChromePlatform) Reset() {
	*x = ChromePlatform{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChromePlatform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChromePlatform) ProtoMessage() {}

func (x *ChromePlatform) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChromePlatform.ProtoReflect.Descriptor instead.
func (*ChromePlatform) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_rawDescGZIP(), []int{0}
}

func (x *ChromePlatform) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChromePlatform) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *ChromePlatform) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ChromePlatform) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *ChromePlatform) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_infra_unifiedfleet_api_v1_models_chrome_platform_proto protoreflect.FileDescriptor

var file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_rawDesc = []byte{
	0x0a, 0x36, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xab, 0x02, 0x0a, 0x0e, 0x43, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x75,
	0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x3a, 0x56, 0xea, 0x41, 0x53, 0x0a, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x2d, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x61, 0x70, 0x70, 0x73, 0x70, 0x6f, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x68, 0x72, 0x6f,
	0x6d, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x20, 0x63, 0x68, 0x72, 0x6f,
	0x6d, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x2f, 0x7b, 0x63, 0x68, 0x72,
	0x6f, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x7d, 0x4a, 0x04, 0x08, 0x05,
	0x10, 0x06, 0x52, 0x0b, 0x76, 0x6d, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x42,
	0x28, 0x5a, 0x26, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x66, 0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x3b, 0x75, 0x66, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_rawDescOnce sync.Once
	file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_rawDescData = file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_rawDesc
)

func file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_rawDescGZIP() []byte {
	file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_rawDescOnce.Do(func() {
		file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_rawDescData)
	})
	return file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_rawDescData
}

var file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_goTypes = []interface{}{
	(*ChromePlatform)(nil),        // 0: unifiedfleet.api.v1.models.ChromePlatform
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_depIdxs = []int32{
	1, // 0: unifiedfleet.api.v1.models.ChromePlatform.update_time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_init() }
func file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_init() {
	if File_infra_unifiedfleet_api_v1_models_chrome_platform_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChromePlatform); i {
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
			RawDescriptor: file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_goTypes,
		DependencyIndexes: file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_depIdxs,
		MessageInfos:      file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_msgTypes,
	}.Build()
	File_infra_unifiedfleet_api_v1_models_chrome_platform_proto = out.File
	file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_rawDesc = nil
	file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_goTypes = nil
	file_infra_unifiedfleet_api_v1_models_chrome_platform_proto_depIdxs = nil
}
