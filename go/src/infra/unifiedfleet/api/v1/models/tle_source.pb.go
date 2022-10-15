// Copyright 2022 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// This proto definition describes the schedulable label protos for TLE sources
// exposed by UFS.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: infra/unifiedfleet/api/v1/models/tle_source.proto

package ufspb

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

// TleSourceType refers to the entity type that a TleSource label is based upon.
//
// Next Tag: 3
type TleSourceType int32

const (
	TleSourceType_TLE_SOURCE_TYPE_UNKNOWN TleSourceType = 0
	// Refers to DutState in infra/unifiedfleet/api/v1/models/chromeos/lab/dut_state.proto
	TleSourceType_TLE_SOURCE_TYPE_DUT_STATE TleSourceType = 1
	// Refers to MachineLSE in infra/unifiedfleet/api/v1/models/machine_lse.proto
	TleSourceType_TLE_SOURCE_TYPE_LAB_CONFIG TleSourceType = 2
)

// Enum value maps for TleSourceType.
var (
	TleSourceType_name = map[int32]string{
		0: "TLE_SOURCE_TYPE_UNKNOWN",
		1: "TLE_SOURCE_TYPE_DUT_STATE",
		2: "TLE_SOURCE_TYPE_LAB_CONFIG",
	}
	TleSourceType_value = map[string]int32{
		"TLE_SOURCE_TYPE_UNKNOWN":    0,
		"TLE_SOURCE_TYPE_DUT_STATE":  1,
		"TLE_SOURCE_TYPE_LAB_CONFIG": 2,
	}
)

func (x TleSourceType) Enum() *TleSourceType {
	p := new(TleSourceType)
	*p = x
	return p
}

func (x TleSourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TleSourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_unifiedfleet_api_v1_models_tle_source_proto_enumTypes[0].Descriptor()
}

func (TleSourceType) Type() protoreflect.EnumType {
	return &file_infra_unifiedfleet_api_v1_models_tle_source_proto_enumTypes[0]
}

func (x TleSourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TleSourceType.Descriptor instead.
func (TleSourceType) EnumDescriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_tle_source_proto_rawDescGZIP(), []int{0}
}

// TleConverterType refers to the converter type that should be used to extract
// the label value from the TleSource.
//
// Next Tag: 4
type TleConverterType int32

const (
	TleConverterType_TLE_CONVERTER_TYPE_UNKNOWN TleConverterType = 0
	// Refers to the TleConverterStandard type.
	TleConverterType_TLE_CONVERTER_TYPE_STANDARD TleConverterType = 1
	// Refers to the TleConverterExistence type.
	TleConverterType_TLE_CONVERTER_TYPE_EXISTENCE TleConverterType = 2
	// Refers to the TleConverterDynamic type.
	TleConverterType_TLE_CONVERTER_TYPE_DYNAMIC TleConverterType = 3
)

// Enum value maps for TleConverterType.
var (
	TleConverterType_name = map[int32]string{
		0: "TLE_CONVERTER_TYPE_UNKNOWN",
		1: "TLE_CONVERTER_TYPE_STANDARD",
		2: "TLE_CONVERTER_TYPE_EXISTENCE",
		3: "TLE_CONVERTER_TYPE_DYNAMIC",
	}
	TleConverterType_value = map[string]int32{
		"TLE_CONVERTER_TYPE_UNKNOWN":   0,
		"TLE_CONVERTER_TYPE_STANDARD":  1,
		"TLE_CONVERTER_TYPE_EXISTENCE": 2,
		"TLE_CONVERTER_TYPE_DYNAMIC":   3,
	}
)

func (x TleConverterType) Enum() *TleConverterType {
	p := new(TleConverterType)
	*p = x
	return p
}

func (x TleConverterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TleConverterType) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_unifiedfleet_api_v1_models_tle_source_proto_enumTypes[1].Descriptor()
}

func (TleConverterType) Type() protoreflect.EnumType {
	return &file_infra_unifiedfleet_api_v1_models_tle_source_proto_enumTypes[1]
}

func (x TleConverterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TleConverterType.Descriptor instead.
func (TleConverterType) EnumDescriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_tle_source_proto_rawDescGZIP(), []int{1}
}

// TleSource refers to the metadata related to a schedulable label that is
// specific to a Test Lab Environment. This metadata is used to extract the
// actual label value from the TleSource entity.
//
// Next Tag: 8
type TleSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the label. Should match 1-1 with a label id defined by a
	// DutAttribute.
	LabelName string `protobuf:"bytes,1,opt,name=label_name,json=labelName,proto3" json:"label_name,omitempty"`
	// The entity type that this label is based upon.
	SourceType TleSourceType `protobuf:"varint,2,opt,name=source_type,json=sourceType,proto3,enum=unifiedfleet.api.v1.models.TleSourceType" json:"source_type,omitempty"`
	// The proto field path to be used in the converter. The path is formatted as
	// a jsonpath.
	FieldPath string `protobuf:"bytes,3,opt,name=field_path,json=fieldPath,proto3" json:"field_path,omitempty"`
	// The converter type to be used to extract the label value.
	ConverterType TleConverterType `protobuf:"varint,4,opt,name=converter_type,json=converterType,proto3,enum=unifiedfleet.api.v1.models.TleConverterType" json:"converter_type,omitempty"`
	// TleConverter contains the metadata needed for the selected converter type.
	//
	// Types that are assignable to Converter:
	//	*TleSource_StandardConverter
	//	*TleSource_ExistenceConverter
	//	*TleSource_DynamicConverter
	Converter isTleSource_Converter `protobuf_oneof:"converter"`
}

func (x *TleSource) Reset() {
	*x = TleSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_tle_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TleSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TleSource) ProtoMessage() {}

func (x *TleSource) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_tle_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TleSource.ProtoReflect.Descriptor instead.
func (*TleSource) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_tle_source_proto_rawDescGZIP(), []int{0}
}

func (x *TleSource) GetLabelName() string {
	if x != nil {
		return x.LabelName
	}
	return ""
}

func (x *TleSource) GetSourceType() TleSourceType {
	if x != nil {
		return x.SourceType
	}
	return TleSourceType_TLE_SOURCE_TYPE_UNKNOWN
}

func (x *TleSource) GetFieldPath() string {
	if x != nil {
		return x.FieldPath
	}
	return ""
}

func (x *TleSource) GetConverterType() TleConverterType {
	if x != nil {
		return x.ConverterType
	}
	return TleConverterType_TLE_CONVERTER_TYPE_UNKNOWN
}

func (m *TleSource) GetConverter() isTleSource_Converter {
	if m != nil {
		return m.Converter
	}
	return nil
}

func (x *TleSource) GetStandardConverter() *TleConverterStandard {
	if x, ok := x.GetConverter().(*TleSource_StandardConverter); ok {
		return x.StandardConverter
	}
	return nil
}

func (x *TleSource) GetExistenceConverter() *TleConverterExistence {
	if x, ok := x.GetConverter().(*TleSource_ExistenceConverter); ok {
		return x.ExistenceConverter
	}
	return nil
}

func (x *TleSource) GetDynamicConverter() *TleConverterDynamic {
	if x, ok := x.GetConverter().(*TleSource_DynamicConverter); ok {
		return x.DynamicConverter
	}
	return nil
}

type isTleSource_Converter interface {
	isTleSource_Converter()
}

type TleSource_StandardConverter struct {
	StandardConverter *TleConverterStandard `protobuf:"bytes,5,opt,name=standard_converter,json=standardConverter,proto3,oneof"`
}

type TleSource_ExistenceConverter struct {
	ExistenceConverter *TleConverterExistence `protobuf:"bytes,6,opt,name=existence_converter,json=existenceConverter,proto3,oneof"`
}

type TleSource_DynamicConverter struct {
	DynamicConverter *TleConverterDynamic `protobuf:"bytes,7,opt,name=dynamic_converter,json=dynamicConverter,proto3,oneof"`
}

func (*TleSource_StandardConverter) isTleSource_Converter() {}

func (*TleSource_ExistenceConverter) isTleSource_Converter() {}

func (*TleSource_DynamicConverter) isTleSource_Converter() {}

// A collection of TleSource
type TleSources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TleSources []*TleSource `protobuf:"bytes,1,rep,name=tle_sources,json=tleSources,proto3" json:"tle_sources,omitempty"`
}

func (x *TleSources) Reset() {
	*x = TleSources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_tle_source_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TleSources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TleSources) ProtoMessage() {}

func (x *TleSources) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_tle_source_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TleSources.ProtoReflect.Descriptor instead.
func (*TleSources) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_tle_source_proto_rawDescGZIP(), []int{1}
}

func (x *TleSources) GetTleSources() []*TleSource {
	if x != nil {
		return x.TleSources
	}
	return nil
}

// TleConverterStandard is the default converter that reads the value directly
// from a given config path.
type TleConverterStandard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If specified, the prefix will be used to append or truncate from the label
	// value.
	Prefix string `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// If true, it specifies append. If false, it specifies truncate.
	AppendPrefix bool `protobuf:"varint,2,opt,name=append_prefix,json=appendPrefix,proto3" json:"append_prefix,omitempty"`
}

func (x *TleConverterStandard) Reset() {
	*x = TleConverterStandard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_tle_source_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TleConverterStandard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TleConverterStandard) ProtoMessage() {}

func (x *TleConverterStandard) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_tle_source_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TleConverterStandard.ProtoReflect.Descriptor instead.
func (*TleConverterStandard) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_tle_source_proto_rawDescGZIP(), []int{2}
}

func (x *TleConverterStandard) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *TleConverterStandard) GetAppendPrefix() bool {
	if x != nil {
		return x.AppendPrefix
	}
	return false
}

// TleConverterExistence determines a boolean value based on a given proto from
// the config path.
type TleConverterExistence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The state_existence proto indicates whether existence is determined based
	// on a state proto. E.g. Servo is based on whether its PeripheralState is
	// set to a proper state or not.
	StateExistence *TleConverterExistence_StateExistence `protobuf:"bytes,1,opt,name=state_existence,json=stateExistence,proto3" json:"state_existence,omitempty"`
}

func (x *TleConverterExistence) Reset() {
	*x = TleConverterExistence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_tle_source_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TleConverterExistence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TleConverterExistence) ProtoMessage() {}

func (x *TleConverterExistence) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_tle_source_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TleConverterExistence.ProtoReflect.Descriptor instead.
func (*TleConverterExistence) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_tle_source_proto_rawDescGZIP(), []int{3}
}

func (x *TleConverterExistence) GetStateExistence() *TleConverterExistence_StateExistence {
	if x != nil {
		return x.StateExistence
	}
	return nil
}

// TleConverterDynamic is a converter that generates label names in runtime.
// The label names depend on another proto and cannot be hardcoded
// exhaustively.
type TleConverterDynamic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TleConverterDynamic) Reset() {
	*x = TleConverterDynamic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_tle_source_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TleConverterDynamic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TleConverterDynamic) ProtoMessage() {}

func (x *TleConverterDynamic) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_tle_source_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TleConverterDynamic.ProtoReflect.Descriptor instead.
func (*TleConverterDynamic) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_tle_source_proto_rawDescGZIP(), []int{4}
}

type TleConverterExistence_StateExistence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of states that indicate a false value for existence.
	// e.g. For Servo, the invalid states are the PeripheralStates UNKNOWN and
	// NOT_CONNECTED.
	InvalidStates []string `protobuf:"bytes,1,rep,name=invalid_states,json=invalidStates,proto3" json:"invalid_states,omitempty"`
}

func (x *TleConverterExistence_StateExistence) Reset() {
	*x = TleConverterExistence_StateExistence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_tle_source_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TleConverterExistence_StateExistence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TleConverterExistence_StateExistence) ProtoMessage() {}

func (x *TleConverterExistence_StateExistence) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_tle_source_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TleConverterExistence_StateExistence.ProtoReflect.Descriptor instead.
func (*TleConverterExistence_StateExistence) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_tle_source_proto_rawDescGZIP(), []int{3, 0}
}

func (x *TleConverterExistence_StateExistence) GetInvalidStates() []string {
	if x != nil {
		return x.InvalidStates
	}
	return nil
}

var File_infra_unifiedfleet_api_v1_models_tle_source_proto protoreflect.FileDescriptor

var file_infra_unifiedfleet_api_v1_models_tle_source_proto_rawDesc = []byte{
	0x0a, 0x31, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x74, 0x6c, 0x65, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22,
	0xa0, 0x04, 0x0a, 0x09, 0x54, 0x6c, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x0b,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x29, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x54,
	0x6c, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x50, 0x61, 0x74, 0x68, 0x12, 0x53, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2c, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x54, 0x6c, 0x65,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x61, 0x0a, 0x12,
	0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x54, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x48, 0x00, 0x52, 0x11, 0x73, 0x74,
	0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x12,
	0x64, 0x0a, 0x13, 0x65, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x75,
	0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x54, 0x6c, 0x65, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x48,
	0x00, 0x52, 0x12, 0x65, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x65, 0x72, 0x12, 0x5e, 0x0a, 0x11, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63,
	0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2f, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x54, 0x6c,
	0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69,
	0x63, 0x48, 0x00, 0x52, 0x10, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x65, 0x72, 0x42, 0x0b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x65, 0x72, 0x22, 0x54, 0x0a, 0x0a, 0x54, 0x6c, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x12, 0x46, 0x0a, 0x0b, 0x74, 0x6c, 0x65, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x54, 0x6c, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0a, 0x74, 0x6c,
	0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x53, 0x0a, 0x14, 0x54, 0x6c, 0x65, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x70, 0x70, 0x65,
	0x6e, 0x64, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0xbb, 0x01,
	0x0a, 0x15, 0x54, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x40, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x54, 0x6c,
	0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x63, 0x65, 0x1a, 0x37, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x54,
	0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x44, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x2a, 0x6b, 0x0a, 0x0d, 0x54, 0x6c, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x4c, 0x45, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x1d, 0x0a, 0x19, 0x54, 0x4c, 0x45, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x44, 0x55, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12,
	0x1e, 0x0a, 0x1a, 0x54, 0x4c, 0x45, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4c, 0x41, 0x42, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x10, 0x02, 0x2a,
	0x95, 0x01, 0x0a, 0x10, 0x54, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x4c, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x56,
	0x45, 0x52, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x54, 0x4c, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x56,
	0x45, 0x52, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x4e, 0x44,
	0x41, 0x52, 0x44, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x54, 0x4c, 0x45, 0x5f, 0x43, 0x4f, 0x4e,
	0x56, 0x45, 0x52, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x58, 0x49, 0x53,
	0x54, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x4c, 0x45, 0x5f, 0x43,
	0x4f, 0x4e, 0x56, 0x45, 0x52, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x59,
	0x4e, 0x41, 0x4d, 0x49, 0x43, 0x10, 0x03, 0x42, 0x28, 0x5a, 0x26, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x3b, 0x75, 0x66, 0x73, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_unifiedfleet_api_v1_models_tle_source_proto_rawDescOnce sync.Once
	file_infra_unifiedfleet_api_v1_models_tle_source_proto_rawDescData = file_infra_unifiedfleet_api_v1_models_tle_source_proto_rawDesc
)

func file_infra_unifiedfleet_api_v1_models_tle_source_proto_rawDescGZIP() []byte {
	file_infra_unifiedfleet_api_v1_models_tle_source_proto_rawDescOnce.Do(func() {
		file_infra_unifiedfleet_api_v1_models_tle_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_unifiedfleet_api_v1_models_tle_source_proto_rawDescData)
	})
	return file_infra_unifiedfleet_api_v1_models_tle_source_proto_rawDescData
}

var file_infra_unifiedfleet_api_v1_models_tle_source_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_infra_unifiedfleet_api_v1_models_tle_source_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_infra_unifiedfleet_api_v1_models_tle_source_proto_goTypes = []interface{}{
	(TleSourceType)(0),                           // 0: unifiedfleet.api.v1.models.TleSourceType
	(TleConverterType)(0),                        // 1: unifiedfleet.api.v1.models.TleConverterType
	(*TleSource)(nil),                            // 2: unifiedfleet.api.v1.models.TleSource
	(*TleSources)(nil),                           // 3: unifiedfleet.api.v1.models.TleSources
	(*TleConverterStandard)(nil),                 // 4: unifiedfleet.api.v1.models.TleConverterStandard
	(*TleConverterExistence)(nil),                // 5: unifiedfleet.api.v1.models.TleConverterExistence
	(*TleConverterDynamic)(nil),                  // 6: unifiedfleet.api.v1.models.TleConverterDynamic
	(*TleConverterExistence_StateExistence)(nil), // 7: unifiedfleet.api.v1.models.TleConverterExistence.StateExistence
}
var file_infra_unifiedfleet_api_v1_models_tle_source_proto_depIdxs = []int32{
	0, // 0: unifiedfleet.api.v1.models.TleSource.source_type:type_name -> unifiedfleet.api.v1.models.TleSourceType
	1, // 1: unifiedfleet.api.v1.models.TleSource.converter_type:type_name -> unifiedfleet.api.v1.models.TleConverterType
	4, // 2: unifiedfleet.api.v1.models.TleSource.standard_converter:type_name -> unifiedfleet.api.v1.models.TleConverterStandard
	5, // 3: unifiedfleet.api.v1.models.TleSource.existence_converter:type_name -> unifiedfleet.api.v1.models.TleConverterExistence
	6, // 4: unifiedfleet.api.v1.models.TleSource.dynamic_converter:type_name -> unifiedfleet.api.v1.models.TleConverterDynamic
	2, // 5: unifiedfleet.api.v1.models.TleSources.tle_sources:type_name -> unifiedfleet.api.v1.models.TleSource
	7, // 6: unifiedfleet.api.v1.models.TleConverterExistence.state_existence:type_name -> unifiedfleet.api.v1.models.TleConverterExistence.StateExistence
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_infra_unifiedfleet_api_v1_models_tle_source_proto_init() }
func file_infra_unifiedfleet_api_v1_models_tle_source_proto_init() {
	if File_infra_unifiedfleet_api_v1_models_tle_source_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_unifiedfleet_api_v1_models_tle_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TleSource); i {
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
		file_infra_unifiedfleet_api_v1_models_tle_source_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TleSources); i {
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
		file_infra_unifiedfleet_api_v1_models_tle_source_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TleConverterStandard); i {
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
		file_infra_unifiedfleet_api_v1_models_tle_source_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TleConverterExistence); i {
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
		file_infra_unifiedfleet_api_v1_models_tle_source_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TleConverterDynamic); i {
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
		file_infra_unifiedfleet_api_v1_models_tle_source_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TleConverterExistence_StateExistence); i {
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
	file_infra_unifiedfleet_api_v1_models_tle_source_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TleSource_StandardConverter)(nil),
		(*TleSource_ExistenceConverter)(nil),
		(*TleSource_DynamicConverter)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_infra_unifiedfleet_api_v1_models_tle_source_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_unifiedfleet_api_v1_models_tle_source_proto_goTypes,
		DependencyIndexes: file_infra_unifiedfleet_api_v1_models_tle_source_proto_depIdxs,
		EnumInfos:         file_infra_unifiedfleet_api_v1_models_tle_source_proto_enumTypes,
		MessageInfos:      file_infra_unifiedfleet_api_v1_models_tle_source_proto_msgTypes,
	}.Build()
	File_infra_unifiedfleet_api_v1_models_tle_source_proto = out.File
	file_infra_unifiedfleet_api_v1_models_tle_source_proto_rawDesc = nil
	file_infra_unifiedfleet_api_v1_models_tle_source_proto_goTypes = nil
	file_infra_unifiedfleet_api_v1_models_tle_source_proto_depIdxs = nil
}
