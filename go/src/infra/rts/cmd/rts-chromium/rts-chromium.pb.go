// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.1
// source: infra/rts/cmd/rts-chromium/rts-chromium.proto

package main

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	proto "infra/rts/presubmit/eval/proto"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A test file used in Chromium.
//
// Used in an RTS model, in file "test-files.jsonl" encoded as JSON Lines
// of TestFile protojson messages.
type TestFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Source-absolute path to the test file, e.g.
	// "//chrome/renderer/autofill/password_autofill_agent_browsertest.cc".
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Names of tests known to be defined in the the file.
	// The names are native to the test framework, e.g.
	// "PasswordAutofillAgentTest.NoMayUsePlaceholderAndPlaceholderOnForm".
	TestNames []string `protobuf:"bytes,2,rep,name=test_names,json=testNames,proto3" json:"test_names,omitempty"`
	// Test targets where the test file was observed, e.g. "browser_tests".
	TestTargets []string `protobuf:"bytes,3,rep,name=test_targets,json=testTargets,proto3" json:"test_targets,omitempty"`
}

func (x *TestFile) Reset() {
	*x = TestFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_rts_cmd_rts_chromium_rts_chromium_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestFile) ProtoMessage() {}

func (x *TestFile) ProtoReflect() protoreflect.Message {
	mi := &file_infra_rts_cmd_rts_chromium_rts_chromium_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestFile.ProtoReflect.Descriptor instead.
func (*TestFile) Descriptor() ([]byte, []int) {
	return file_infra_rts_cmd_rts_chromium_rts_chromium_proto_rawDescGZIP(), []int{0}
}

func (x *TestFile) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *TestFile) GetTestNames() []string {
	if x != nil {
		return x.TestNames
	}
	return nil
}

func (x *TestFile) GetTestTargets() []string {
	if x != nil {
		return x.TestTargets
	}
	return nil
}

// Configuration of a git-based selection strategy.
type GitBasedStrategyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value for git.EdgeReader.ChangeLogDistanceFactor;
	ChangeLogDistanceFactor float32 `protobuf:"fixed32,1,opt,name=change_log_distance_factor,json=changeLogDistanceFactor,proto3" json:"change_log_distance_factor,omitempty"`
	// Value for git.EdgeReader.FileStructureDistanceFactor.
	FileStructureDistanceFactor float32 `protobuf:"fixed32,2,opt,name=file_structure_distance_factor,json=fileStructureDistanceFactor,proto3" json:"file_structure_distance_factor,omitempty"`
	// Thresholds with change recalls and savings,
	// ordered by change recall.
	Thresholds []*proto.Threshold `protobuf:"bytes,3,rep,name=thresholds,proto3" json:"thresholds,omitempty"`
}

func (x *GitBasedStrategyConfig) Reset() {
	*x = GitBasedStrategyConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_rts_cmd_rts_chromium_rts_chromium_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitBasedStrategyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitBasedStrategyConfig) ProtoMessage() {}

func (x *GitBasedStrategyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_infra_rts_cmd_rts_chromium_rts_chromium_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitBasedStrategyConfig.ProtoReflect.Descriptor instead.
func (*GitBasedStrategyConfig) Descriptor() ([]byte, []int) {
	return file_infra_rts_cmd_rts_chromium_rts_chromium_proto_rawDescGZIP(), []int{1}
}

func (x *GitBasedStrategyConfig) GetChangeLogDistanceFactor() float32 {
	if x != nil {
		return x.ChangeLogDistanceFactor
	}
	return 0
}

func (x *GitBasedStrategyConfig) GetFileStructureDistanceFactor() float32 {
	if x != nil {
		return x.FileStructureDistanceFactor
	}
	return 0
}

func (x *GitBasedStrategyConfig) GetThresholds() []*proto.Threshold {
	if x != nil {
		return x.Thresholds
	}
	return nil
}

var File_infra_rts_cmd_rts_chromium_rts_chromium_proto protoreflect.FileDescriptor

var file_infra_rts_cmd_rts_chromium_rts_chromium_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x72, 0x74, 0x73, 0x2f, 0x63, 0x6d, 0x64, 0x2f,
	0x72, 0x74, 0x73, 0x2d, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2f, 0x72, 0x74, 0x73,
	0x2d, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x2e, 0x72, 0x74, 0x73, 0x1a, 0x2c, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x2f, 0x72, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x2f, 0x65, 0x76, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x08, 0x54, 0x65, 0x73,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x65, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x65, 0x73, 0x74,
	0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x74, 0x65, 0x73, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x22, 0xe0, 0x01, 0x0a, 0x16,
	0x47, 0x69, 0x74, 0x42, 0x61, 0x73, 0x65, 0x64, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3b, 0x0a, 0x1a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x66, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x17, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x4c, 0x6f, 0x67, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x43, 0x0a, 0x1e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x66,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x1b, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x44, 0x0a, 0x0a, 0x74, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x65, 0x2e, 0x72, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x65, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x2e, 0x65, 0x76, 0x61, 0x6c, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x52, 0x0a, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x42, 0x21,
	0x5a, 0x1f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x72, 0x74, 0x73, 0x2f, 0x63, 0x6d, 0x64, 0x2f,
	0x72, 0x74, 0x73, 0x2d, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x3b, 0x6d, 0x61, 0x69,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_rts_cmd_rts_chromium_rts_chromium_proto_rawDescOnce sync.Once
	file_infra_rts_cmd_rts_chromium_rts_chromium_proto_rawDescData = file_infra_rts_cmd_rts_chromium_rts_chromium_proto_rawDesc
)

func file_infra_rts_cmd_rts_chromium_rts_chromium_proto_rawDescGZIP() []byte {
	file_infra_rts_cmd_rts_chromium_rts_chromium_proto_rawDescOnce.Do(func() {
		file_infra_rts_cmd_rts_chromium_rts_chromium_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_rts_cmd_rts_chromium_rts_chromium_proto_rawDescData)
	})
	return file_infra_rts_cmd_rts_chromium_rts_chromium_proto_rawDescData
}

var file_infra_rts_cmd_rts_chromium_rts_chromium_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_infra_rts_cmd_rts_chromium_rts_chromium_proto_goTypes = []interface{}{
	(*TestFile)(nil),               // 0: chrome.rts.TestFile
	(*GitBasedStrategyConfig)(nil), // 1: chrome.rts.GitBasedStrategyConfig
	(*proto.Threshold)(nil),        // 2: chrome.rts.presubmit.eval.Threshold
}
var file_infra_rts_cmd_rts_chromium_rts_chromium_proto_depIdxs = []int32{
	2, // 0: chrome.rts.GitBasedStrategyConfig.thresholds:type_name -> chrome.rts.presubmit.eval.Threshold
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_infra_rts_cmd_rts_chromium_rts_chromium_proto_init() }
func file_infra_rts_cmd_rts_chromium_rts_chromium_proto_init() {
	if File_infra_rts_cmd_rts_chromium_rts_chromium_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_rts_cmd_rts_chromium_rts_chromium_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestFile); i {
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
		file_infra_rts_cmd_rts_chromium_rts_chromium_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GitBasedStrategyConfig); i {
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
			RawDescriptor: file_infra_rts_cmd_rts_chromium_rts_chromium_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_rts_cmd_rts_chromium_rts_chromium_proto_goTypes,
		DependencyIndexes: file_infra_rts_cmd_rts_chromium_rts_chromium_proto_depIdxs,
		MessageInfos:      file_infra_rts_cmd_rts_chromium_rts_chromium_proto_msgTypes,
	}.Build()
	File_infra_rts_cmd_rts_chromium_rts_chromium_proto = out.File
	file_infra_rts_cmd_rts_chromium_rts_chromium_proto_rawDesc = nil
	file_infra_rts_cmd_rts_chromium_rts_chromium_proto_goTypes = nil
	file_infra_rts_cmd_rts_chromium_rts_chromium_proto_depIdxs = nil
}
