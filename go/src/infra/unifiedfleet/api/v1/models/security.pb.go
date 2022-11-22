// Copyright 2022 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: infra/unifiedfleet/api/v1/models/security.proto

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

type SecurityInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the security pool
	PoolName string `protobuf:"bytes,1,opt,name=pool_name,json=poolName,proto3" json:"pool_name,omitempty"`
	// customer group with access to this pool
	Customer string `protobuf:"bytes,2,opt,name=customer,proto3" json:"customer,omitempty"`
	// security level of the bot ex:trusted, untrusted etc
	SecurityLevel string `protobuf:"bytes,3,opt,name=security_level,json=securityLevel,proto3" json:"security_level,omitempty"`
	// custom MIBA realm for this pool
	// default is "<customer>_<security_level>"
	MibaRealm string `protobuf:"bytes,4,opt,name=miba_realm,json=mibaRealm,proto3" json:"miba_realm,omitempty"`
	// id of the swarming server that owns this pool
	SwarmingServerId string `protobuf:"bytes,5,opt,name=swarming_server_id,json=swarmingServerId,proto3" json:"swarming_server_id,omitempty"`
	// hosts that belong to this pool
	Hosts []string `protobuf:"bytes,6,rep,name=hosts,proto3" json:"hosts,omitempty"`
}

func (x *SecurityInfo) Reset() {
	*x = SecurityInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_security_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityInfo) ProtoMessage() {}

func (x *SecurityInfo) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_security_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityInfo.ProtoReflect.Descriptor instead.
func (*SecurityInfo) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_security_proto_rawDescGZIP(), []int{0}
}

func (x *SecurityInfo) GetPoolName() string {
	if x != nil {
		return x.PoolName
	}
	return ""
}

func (x *SecurityInfo) GetCustomer() string {
	if x != nil {
		return x.Customer
	}
	return ""
}

func (x *SecurityInfo) GetSecurityLevel() string {
	if x != nil {
		return x.SecurityLevel
	}
	return ""
}

func (x *SecurityInfo) GetMibaRealm() string {
	if x != nil {
		return x.MibaRealm
	}
	return ""
}

func (x *SecurityInfo) GetSwarmingServerId() string {
	if x != nil {
		return x.SwarmingServerId
	}
	return ""
}

func (x *SecurityInfo) GetHosts() []string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

type SecurityInfos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list of pools with security info
	Pools []*SecurityInfo `protobuf:"bytes,1,rep,name=pools,proto3" json:"pools,omitempty"`
}

func (x *SecurityInfos) Reset() {
	*x = SecurityInfos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_security_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityInfos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityInfos) ProtoMessage() {}

func (x *SecurityInfos) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_security_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityInfos.ProtoReflect.Descriptor instead.
func (*SecurityInfos) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_security_proto_rawDescGZIP(), []int{1}
}

func (x *SecurityInfos) GetPools() []*SecurityInfo {
	if x != nil {
		return x.Pools
	}
	return nil
}

var File_infra_unifiedfleet_api_v1_models_security_proto protoreflect.FileDescriptor

var file_infra_unifiedfleet_api_v1_models_security_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1a, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0xd1, 0x01,
	0x0a, 0x0c, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x6f, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1d,
	0x0a, 0x0a, 0x6d, 0x69, 0x62, 0x61, 0x5f, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x69, 0x62, 0x61, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x12, 0x2c, 0x0a,
	0x12, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x77, 0x61, 0x72, 0x6d,
	0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x68,
	0x6f, 0x73, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74,
	0x73, 0x22, 0x4f, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x12, 0x3e, 0x0a, 0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x53,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x70, 0x6f, 0x6f,
	0x6c, 0x73, 0x42, 0x28, 0x5a, 0x26, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x3b, 0x75, 0x66, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_unifiedfleet_api_v1_models_security_proto_rawDescOnce sync.Once
	file_infra_unifiedfleet_api_v1_models_security_proto_rawDescData = file_infra_unifiedfleet_api_v1_models_security_proto_rawDesc
)

func file_infra_unifiedfleet_api_v1_models_security_proto_rawDescGZIP() []byte {
	file_infra_unifiedfleet_api_v1_models_security_proto_rawDescOnce.Do(func() {
		file_infra_unifiedfleet_api_v1_models_security_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_unifiedfleet_api_v1_models_security_proto_rawDescData)
	})
	return file_infra_unifiedfleet_api_v1_models_security_proto_rawDescData
}

var file_infra_unifiedfleet_api_v1_models_security_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_infra_unifiedfleet_api_v1_models_security_proto_goTypes = []interface{}{
	(*SecurityInfo)(nil),  // 0: unifiedfleet.api.v1.models.SecurityInfo
	(*SecurityInfos)(nil), // 1: unifiedfleet.api.v1.models.SecurityInfos
}
var file_infra_unifiedfleet_api_v1_models_security_proto_depIdxs = []int32{
	0, // 0: unifiedfleet.api.v1.models.SecurityInfos.pools:type_name -> unifiedfleet.api.v1.models.SecurityInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_infra_unifiedfleet_api_v1_models_security_proto_init() }
func file_infra_unifiedfleet_api_v1_models_security_proto_init() {
	if File_infra_unifiedfleet_api_v1_models_security_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_unifiedfleet_api_v1_models_security_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityInfo); i {
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
		file_infra_unifiedfleet_api_v1_models_security_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityInfos); i {
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
			RawDescriptor: file_infra_unifiedfleet_api_v1_models_security_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_unifiedfleet_api_v1_models_security_proto_goTypes,
		DependencyIndexes: file_infra_unifiedfleet_api_v1_models_security_proto_depIdxs,
		MessageInfos:      file_infra_unifiedfleet_api_v1_models_security_proto_msgTypes,
	}.Build()
	File_infra_unifiedfleet_api_v1_models_security_proto = out.File
	file_infra_unifiedfleet_api_v1_models_security_proto_rawDesc = nil
	file_infra_unifiedfleet_api_v1_models_security_proto_goTypes = nil
	file_infra_unifiedfleet_api_v1_models_security_proto_depIdxs = nil
}