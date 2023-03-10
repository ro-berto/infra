// Copyright 2017 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: infra/appengine/arquebus/app/backend/tq.proto

package backend

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

// ScheduleAssignerTask is to schedule a new task for a given assigner.
//
// Queue: "schedule-assigners".
type ScheduleAssignerTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssignerId string `protobuf:"bytes,1,opt,name=assigner_id,json=assignerId,proto3" json:"assigner_id,omitempty"`
}

func (x *ScheduleAssignerTask) Reset() {
	*x = ScheduleAssignerTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_arquebus_app_backend_tq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleAssignerTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleAssignerTask) ProtoMessage() {}

func (x *ScheduleAssignerTask) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_arquebus_app_backend_tq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleAssignerTask.ProtoReflect.Descriptor instead.
func (*ScheduleAssignerTask) Descriptor() ([]byte, []int) {
	return file_infra_appengine_arquebus_app_backend_tq_proto_rawDescGZIP(), []int{0}
}

func (x *ScheduleAssignerTask) GetAssignerId() string {
	if x != nil {
		return x.AssignerId
	}
	return ""
}

// RunAssignerTask is to trigger an assigner run for a given scheduled
// task.
//
// Queue: "run-assigners".
type RunAssignerTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssignerId string `protobuf:"bytes,1,opt,name=assigner_id,json=assignerId,proto3" json:"assigner_id,omitempty"`
	// the task ID of a scheduled tasks to trigger a run.
	TaskId int64 `protobuf:"varint,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *RunAssignerTask) Reset() {
	*x = RunAssignerTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_arquebus_app_backend_tq_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunAssignerTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunAssignerTask) ProtoMessage() {}

func (x *RunAssignerTask) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_arquebus_app_backend_tq_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunAssignerTask.ProtoReflect.Descriptor instead.
func (*RunAssignerTask) Descriptor() ([]byte, []int) {
	return file_infra_appengine_arquebus_app_backend_tq_proto_rawDescGZIP(), []int{1}
}

func (x *RunAssignerTask) GetAssignerId() string {
	if x != nil {
		return x.AssignerId
	}
	return ""
}

func (x *RunAssignerTask) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

var File_infra_appengine_arquebus_app_backend_tq_proto protoreflect.FileDescriptor

var file_infra_appengine_arquebus_app_backend_tq_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x61, 0x72, 0x71, 0x75, 0x65, 0x62, 0x75, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x74, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x74, 0x71, 0x22, 0x37, 0x0a, 0x14, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x0f, 0x52, 0x75, 0x6e, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49,
	0x64, 0x42, 0x26, 0x5a, 0x24, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x72, 0x71, 0x75, 0x65, 0x62, 0x75, 0x73, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_infra_appengine_arquebus_app_backend_tq_proto_rawDescOnce sync.Once
	file_infra_appengine_arquebus_app_backend_tq_proto_rawDescData = file_infra_appengine_arquebus_app_backend_tq_proto_rawDesc
)

func file_infra_appengine_arquebus_app_backend_tq_proto_rawDescGZIP() []byte {
	file_infra_appengine_arquebus_app_backend_tq_proto_rawDescOnce.Do(func() {
		file_infra_appengine_arquebus_app_backend_tq_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_appengine_arquebus_app_backend_tq_proto_rawDescData)
	})
	return file_infra_appengine_arquebus_app_backend_tq_proto_rawDescData
}

var file_infra_appengine_arquebus_app_backend_tq_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_infra_appengine_arquebus_app_backend_tq_proto_goTypes = []interface{}{
	(*ScheduleAssignerTask)(nil), // 0: backend.tq.ScheduleAssignerTask
	(*RunAssignerTask)(nil),      // 1: backend.tq.RunAssignerTask
}
var file_infra_appengine_arquebus_app_backend_tq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_infra_appengine_arquebus_app_backend_tq_proto_init() }
func file_infra_appengine_arquebus_app_backend_tq_proto_init() {
	if File_infra_appengine_arquebus_app_backend_tq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_appengine_arquebus_app_backend_tq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleAssignerTask); i {
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
		file_infra_appengine_arquebus_app_backend_tq_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunAssignerTask); i {
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
			RawDescriptor: file_infra_appengine_arquebus_app_backend_tq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_appengine_arquebus_app_backend_tq_proto_goTypes,
		DependencyIndexes: file_infra_appengine_arquebus_app_backend_tq_proto_depIdxs,
		MessageInfos:      file_infra_appengine_arquebus_app_backend_tq_proto_msgTypes,
	}.Build()
	File_infra_appengine_arquebus_app_backend_tq_proto = out.File
	file_infra_appengine_arquebus_app_backend_tq_proto_rawDesc = nil
	file_infra_appengine_arquebus_app_backend_tq_proto_goTypes = nil
	file_infra_appengine_arquebus_app_backend_tq_proto_depIdxs = nil
}
