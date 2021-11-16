// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.0
// source: infra/cros/karte/api/action.proto

package kartepb

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

// The status of an action is whether the action succeeded or failed.
type Action_Status int32

const (
	Action_STATUS_UNSPECIFIED Action_Status = 0
	Action_SUCCESS            Action_Status = 1
	Action_FAIL               Action_Status = 2
)

// Enum value maps for Action_Status.
var (
	Action_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "SUCCESS",
		2: "FAIL",
	}
	Action_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"SUCCESS":            1,
		"FAIL":               2,
	}
)

func (x Action_Status) Enum() *Action_Status {
	p := new(Action_Status)
	*p = x
	return p
}

func (x Action_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_cros_karte_api_action_proto_enumTypes[0].Descriptor()
}

func (Action_Status) Type() protoreflect.EnumType {
	return &file_infra_cros_karte_api_action_proto_enumTypes[0]
}

func (x Action_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action_Status.Descriptor instead.
func (Action_Status) EnumDescriptor() ([]byte, []int) {
	return file_infra_cros_karte_api_action_proto_rawDescGZIP(), []int{0, 0}
}

// An action represents an event that was intentionally performed on a DUT.
// Examples include running a command on a DUT or resetting the servo
// attached to a DUT.
//
type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the action. Names are generated
	// automatically when a new action is created.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A kind is a coarse-grained type of an action, such as
	// ssh-attempt. New action_kinds will be created frequently so this field
	// is a string; see https://google.aip.dev/126 for details.
	Kind string `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	// A swarming task ID is the ID of a single swarming task.
	// The swarming task of an action is the swarming task that invoked the
	// action.
	// For example, "4f6c0ba2ef3fc610" is a swarming task ID.
	SwarmingTaskId string `protobuf:"bytes,4,opt,name=swarming_task_id,json=swarmingTaskId,proto3" json:"swarming_task_id,omitempty"`
	// An asset tag is the tag of a given asset in UFS.
	// An asset tag may be a short number such as C444444 printed on a device,
	// or it may be a UUID in some circumstances.
	AssetTag string `protobuf:"bytes,5,opt,name=asset_tag,json=assetTag,proto3" json:"asset_tag,omitempty"`
	// The start time is the time that an action started.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The stop time is the time that an action finished.
	StopTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=stop_time,json=stopTime,proto3" json:"stop_time,omitempty"`
	// The create time is the time that an action was created by Karte.
	// This is the time that the event was first received, since events are
	// immutable outside of rare cases.
	// This field is managed by Karte itself.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Status     Action_Status          `protobuf:"varint,9,opt,name=status,proto3,enum=chromeos.karte.Action_Status" json:"status,omitempty"`
	// The fail reason of an event is a diagnostic message that is emitted when
	// the action in question has failed.
	FailReason string `protobuf:"bytes,10,opt,name=fail_reason,json=failReason,proto3" json:"fail_reason,omitempty"`
	// The seal time is when the particular Karte record is sealed and no further changes can be made.
	SealTime *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=seal_time,json=sealTime,proto3" json:"seal_time,omitempty"`
	// The client is the name of the entity creating the Action entry, e.g. "paris".
	ClientName string `protobuf:"bytes,12,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
	// The client version is the version of the entity creating the Action entry, e.g. "0.0.1".
	ClientVersion string `protobuf:"bytes,13,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_cros_karte_api_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_infra_cros_karte_api_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_infra_cros_karte_api_action_proto_rawDescGZIP(), []int{0}
}

func (x *Action) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Action) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Action) GetSwarmingTaskId() string {
	if x != nil {
		return x.SwarmingTaskId
	}
	return ""
}

func (x *Action) GetAssetTag() string {
	if x != nil {
		return x.AssetTag
	}
	return ""
}

func (x *Action) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Action) GetStopTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StopTime
	}
	return nil
}

func (x *Action) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Action) GetStatus() Action_Status {
	if x != nil {
		return x.Status
	}
	return Action_STATUS_UNSPECIFIED
}

func (x *Action) GetFailReason() string {
	if x != nil {
		return x.FailReason
	}
	return ""
}

func (x *Action) GetSealTime() *timestamppb.Timestamp {
	if x != nil {
		return x.SealTime
	}
	return nil
}

func (x *Action) GetClientName() string {
	if x != nil {
		return x.ClientName
	}
	return ""
}

func (x *Action) GetClientVersion() string {
	if x != nil {
		return x.ClientVersion
	}
	return ""
}

var File_infra_cros_karte_api_action_proto protoreflect.FileDescriptor

var file_infra_cros_karte_api_action_proto_rawDesc = []byte{
	0x0a, 0x21, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x2f, 0x6b, 0x61, 0x72,
	0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x2e, 0x6b, 0x61,
	0x72, 0x74, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x83, 0x05, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x5f,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73,
	0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x54, 0x61, 0x67, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x35, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1d, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x2e, 0x6b, 0x61, 0x72, 0x74,
	0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x5f,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x6c,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x73, 0x65, 0x61, 0x6c, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c,
	0x10, 0x02, 0x3a, 0x2f, 0xea, 0x41, 0x2c, 0x0a, 0x18, 0x6b, 0x61, 0x72, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x70, 0x73, 0x70, 0x6f, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x10, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x7d, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x42, 0x1e, 0x5a, 0x1c, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f,
	0x63, 0x72, 0x6f, 0x73, 0x2f, 0x6b, 0x61, 0x72, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x6b,
	0x61, 0x72, 0x74, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_cros_karte_api_action_proto_rawDescOnce sync.Once
	file_infra_cros_karte_api_action_proto_rawDescData = file_infra_cros_karte_api_action_proto_rawDesc
)

func file_infra_cros_karte_api_action_proto_rawDescGZIP() []byte {
	file_infra_cros_karte_api_action_proto_rawDescOnce.Do(func() {
		file_infra_cros_karte_api_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_cros_karte_api_action_proto_rawDescData)
	})
	return file_infra_cros_karte_api_action_proto_rawDescData
}

var file_infra_cros_karte_api_action_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_infra_cros_karte_api_action_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_infra_cros_karte_api_action_proto_goTypes = []interface{}{
	(Action_Status)(0),            // 0: chromeos.karte.Action.Status
	(*Action)(nil),                // 1: chromeos.karte.Action
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_infra_cros_karte_api_action_proto_depIdxs = []int32{
	2, // 0: chromeos.karte.Action.start_time:type_name -> google.protobuf.Timestamp
	2, // 1: chromeos.karte.Action.stop_time:type_name -> google.protobuf.Timestamp
	2, // 2: chromeos.karte.Action.create_time:type_name -> google.protobuf.Timestamp
	0, // 3: chromeos.karte.Action.status:type_name -> chromeos.karte.Action.Status
	2, // 4: chromeos.karte.Action.seal_time:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_infra_cros_karte_api_action_proto_init() }
func file_infra_cros_karte_api_action_proto_init() {
	if File_infra_cros_karte_api_action_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_cros_karte_api_action_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
			RawDescriptor: file_infra_cros_karte_api_action_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_cros_karte_api_action_proto_goTypes,
		DependencyIndexes: file_infra_cros_karte_api_action_proto_depIdxs,
		EnumInfos:         file_infra_cros_karte_api_action_proto_enumTypes,
		MessageInfos:      file_infra_cros_karte_api_action_proto_msgTypes,
	}.Build()
	File_infra_cros_karte_api_action_proto = out.File
	file_infra_cros_karte_api_action_proto_rawDesc = nil
	file_infra_cros_karte_api_action_proto_goTypes = nil
	file_infra_cros_karte_api_action_proto_depIdxs = nil
}
