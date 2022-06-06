// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// This proto definition describes the structure of the requirements for deployment
// of a host.
//
// Put this proto in chrome infra repo for 3 reasons:
//     * It's possible that we migrate automation services back to chrome infra, and we
//       don't need to move around protos at that time.
//     * There're some commonly used deployment units in UFS and automation services.
//     * All proto here is copied to google3 via copybara, so current Chrome MDM service
//       in google3 could also use it.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: infra/unifiedfleet/api/v1/models/machine_lse_deployment_requirement.proto

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

type HostsToProfileItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of matching between hosts to its specified profiles
	MatchingItem []*HostsToProfileItem `protobuf:"bytes,1,rep,name=matching_item,json=matchingItem,proto3" json:"matching_item,omitempty"`
}

func (x *HostsToProfileItems) Reset() {
	*x = HostsToProfileItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostsToProfileItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostsToProfileItems) ProtoMessage() {}

func (x *HostsToProfileItems) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostsToProfileItems.ProtoReflect.Descriptor instead.
func (*HostsToProfileItems) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_rawDescGZIP(), []int{0}
}

func (x *HostsToProfileItems) GetMatchingItem() []*HostsToProfileItem {
	if x != nil {
		return x.MatchingItem
	}
	return nil
}

type HostsToProfileItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// For each matching, the key is one host criterial, the value is
	// a list of profiles to be pushed
	HostCriterial *HostCriterial `protobuf:"bytes,1,opt,name=host_criterial,json=hostCriterial,proto3" json:"host_criterial,omitempty"`
	Profiles      []*Payload     `protobuf:"bytes,2,rep,name=profiles,proto3" json:"profiles,omitempty"`
}

func (x *HostsToProfileItem) Reset() {
	*x = HostsToProfileItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostsToProfileItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostsToProfileItem) ProtoMessage() {}

func (x *HostsToProfileItem) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostsToProfileItem.ProtoReflect.Descriptor instead.
func (*HostsToProfileItem) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_rawDescGZIP(), []int{1}
}

func (x *HostsToProfileItem) GetHostCriterial() *HostCriterial {
	if x != nil {
		return x.HostCriterial
	}
	return nil
}

func (x *HostsToProfileItem) GetProfiles() []*Payload {
	if x != nil {
		return x.Profiles
	}
	return nil
}

type HostCriterial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Users can specify multiple requirements in each host criterial
	//   * Search the bots based on the swarming instances and dimensions
	//   * List the hostnames to push
	//   * All enrolled hosts need to be pushed
	// The priority comparison:
	//   all > hostnames > swarming_criterials
	All       bool     `protobuf:"varint,1,opt,name=all,proto3" json:"all,omitempty"`
	Hostnames []string `protobuf:"bytes,2,rep,name=hostnames,proto3" json:"hostnames,omitempty"`
	// When swarming is replaced by RBE, this criterial will be deprecated and
	// replaced by rbe_criterials.
	SwarmingCriterials []*SwarmingCriterial `protobuf:"bytes,3,rep,name=swarming_criterials,json=swarmingCriterials,proto3" json:"swarming_criterials,omitempty"`
}

func (x *HostCriterial) Reset() {
	*x = HostCriterial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostCriterial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostCriterial) ProtoMessage() {}

func (x *HostCriterial) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostCriterial.ProtoReflect.Descriptor instead.
func (*HostCriterial) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_rawDescGZIP(), []int{2}
}

func (x *HostCriterial) GetAll() bool {
	if x != nil {
		return x.All
	}
	return false
}

func (x *HostCriterial) GetHostnames() []string {
	if x != nil {
		return x.Hostnames
	}
	return nil
}

func (x *HostCriterial) GetSwarmingCriterials() []*SwarmingCriterial {
	if x != nil {
		return x.SwarmingCriterials
	}
	return nil
}

type SwarmingCriterial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// E.g. chromium-swarm.appspot.com
	Instance   string       `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	Dimensions []*Dimension `protobuf:"bytes,2,rep,name=dimensions,proto3" json:"dimensions,omitempty"`
}

func (x *SwarmingCriterial) Reset() {
	*x = SwarmingCriterial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwarmingCriterial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwarmingCriterial) ProtoMessage() {}

func (x *SwarmingCriterial) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwarmingCriterial.ProtoReflect.Descriptor instead.
func (*SwarmingCriterial) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_rawDescGZIP(), []int{3}
}

func (x *SwarmingCriterial) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

func (x *SwarmingCriterial) GetDimensions() []*Dimension {
	if x != nil {
		return x.Dimensions
	}
	return nil
}

// Representing a matching <string, []string>
type Dimension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Dimension) Reset() {
	*x = Dimension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dimension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dimension) ProtoMessage() {}

func (x *Dimension) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dimension.ProtoReflect.Descriptor instead.
func (*Dimension) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_rawDescGZIP(), []int{4}
}

func (x *Dimension) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Dimension) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto protoreflect.FileDescriptor

var file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_rawDesc = []byte{
	0x0a, 0x49, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x6c, 0x73, 0x65, 0x5f, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x75, 0x6e, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x31, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75,
	0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x13, 0x48, 0x6f,
	0x73, 0x74, 0x73, 0x54, 0x6f, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x53, 0x0a, 0x0d, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x74,
	0x65, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x54, 0x6f, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x22, 0xa7, 0x01, 0x0a, 0x12, 0x48, 0x6f, 0x73, 0x74, 0x73,
	0x54, 0x6f, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x50, 0x0a,
	0x0e, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x52, 0x0d, 0x68, 0x6f, 0x73, 0x74, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12,
	0x3f, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x22, 0x9f, 0x01, 0x0a, 0x0d, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x61, 0x6c, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x12, 0x5e, 0x0a, 0x13, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x5f, 0x63,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x53, 0x77, 0x61,
	0x72, 0x6d, 0x69, 0x6e, 0x67, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x12,
	0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x73, 0x22, 0x76, 0x0a, 0x11, 0x53, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x43, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a,
	0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x35, 0x0a, 0x09, 0x44, 0x69,
	0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x42, 0x28, 0x5a, 0x26, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x3b, 0x75, 0x66, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_rawDescOnce sync.Once
	file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_rawDescData = file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_rawDesc
)

func file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_rawDescGZIP() []byte {
	file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_rawDescOnce.Do(func() {
		file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_rawDescData)
	})
	return file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_rawDescData
}

var file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_goTypes = []interface{}{
	(*HostsToProfileItems)(nil), // 0: unifiedfleet.api.v1.models.HostsToProfileItems
	(*HostsToProfileItem)(nil),  // 1: unifiedfleet.api.v1.models.HostsToProfileItem
	(*HostCriterial)(nil),       // 2: unifiedfleet.api.v1.models.HostCriterial
	(*SwarmingCriterial)(nil),   // 3: unifiedfleet.api.v1.models.SwarmingCriterial
	(*Dimension)(nil),           // 4: unifiedfleet.api.v1.models.Dimension
	(*Payload)(nil),             // 5: unifiedfleet.api.v1.models.Payload
}
var file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_depIdxs = []int32{
	1, // 0: unifiedfleet.api.v1.models.HostsToProfileItems.matching_item:type_name -> unifiedfleet.api.v1.models.HostsToProfileItem
	2, // 1: unifiedfleet.api.v1.models.HostsToProfileItem.host_criterial:type_name -> unifiedfleet.api.v1.models.HostCriterial
	5, // 2: unifiedfleet.api.v1.models.HostsToProfileItem.profiles:type_name -> unifiedfleet.api.v1.models.Payload
	3, // 3: unifiedfleet.api.v1.models.HostCriterial.swarming_criterials:type_name -> unifiedfleet.api.v1.models.SwarmingCriterial
	4, // 4: unifiedfleet.api.v1.models.SwarmingCriterial.dimensions:type_name -> unifiedfleet.api.v1.models.Dimension
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_init() }
func file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_init() {
	if File_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto != nil {
		return
	}
	file_infra_unifiedfleet_api_v1_models_deployment_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostsToProfileItems); i {
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
		file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostsToProfileItem); i {
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
		file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostCriterial); i {
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
		file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwarmingCriterial); i {
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
		file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dimension); i {
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
			RawDescriptor: file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_goTypes,
		DependencyIndexes: file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_depIdxs,
		MessageInfos:      file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_msgTypes,
	}.Build()
	File_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto = out.File
	file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_rawDesc = nil
	file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_goTypes = nil
	file_infra_unifiedfleet_api_v1_models_machine_lse_deployment_requirement_proto_depIdxs = nil
}
