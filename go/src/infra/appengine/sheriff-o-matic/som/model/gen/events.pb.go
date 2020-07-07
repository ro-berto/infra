// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: infra/appengine/sheriff-o-matic/som/model/gen/events.proto

package gen

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// AlertType is the type of alert, from AlertType constants in
// infra/monitoring/messages/alerts.go. Until those constants are replaced
// in the analyzer code with these enum values, this list will have to be
// manually kept in sync.
type SOMAlertsEvent_Alert_AlertType int32

const (
	// STALE_MASTER indicates that we have no recent updates from the master.
	SOMAlertsEvent_Alert_STALE_MASTER SOMAlertsEvent_Alert_AlertType = 0
	// HUNG_BUILDER indicates that a builder has been executing a step for too long.
	SOMAlertsEvent_Alert_HUNG_BUILDER SOMAlertsEvent_Alert_AlertType = 1
	// OFFLINE_BUILDER indicates that we have no recent updates from the builder.
	SOMAlertsEvent_Alert_OFFLINE_BUILDER SOMAlertsEvent_Alert_AlertType = 2
	// IDLE_BUILDER indicates that a builder has not executed any builds recently
	// even though it has requests queued up.
	SOMAlertsEvent_Alert_IDLE_BUILDER SOMAlertsEvent_Alert_AlertType = 3
	// INFRA_FAILURE indicates that a builder step failed due to infrastructure.
	SOMAlertsEvent_Alert_INFRA_FAILURE SOMAlertsEvent_Alert_AlertType = 4
	// BUILD_FAILURE indicates that one of the build steps failed, most likely
	// due to the patch it's building/running with.
	SOMAlertsEvent_Alert_BUILD_FAILURE SOMAlertsEvent_Alert_AlertType = 5
	// TEST_FAILURE indicates that one or more of the tests in the build failed.
	SOMAlertsEvent_Alert_TEST_FAILURE SOMAlertsEvent_Alert_AlertType = 6
)

// Enum value maps for SOMAlertsEvent_Alert_AlertType.
var (
	SOMAlertsEvent_Alert_AlertType_name = map[int32]string{
		0: "STALE_MASTER",
		1: "HUNG_BUILDER",
		2: "OFFLINE_BUILDER",
		3: "IDLE_BUILDER",
		4: "INFRA_FAILURE",
		5: "BUILD_FAILURE",
		6: "TEST_FAILURE",
	}
	SOMAlertsEvent_Alert_AlertType_value = map[string]int32{
		"STALE_MASTER":    0,
		"HUNG_BUILDER":    1,
		"OFFLINE_BUILDER": 2,
		"IDLE_BUILDER":    3,
		"INFRA_FAILURE":   4,
		"BUILD_FAILURE":   5,
		"TEST_FAILURE":    6,
	}
)

func (x SOMAlertsEvent_Alert_AlertType) Enum() *SOMAlertsEvent_Alert_AlertType {
	p := new(SOMAlertsEvent_Alert_AlertType)
	*p = x
	return p
}

func (x SOMAlertsEvent_Alert_AlertType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SOMAlertsEvent_Alert_AlertType) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_enumTypes[0].Descriptor()
}

func (SOMAlertsEvent_Alert_AlertType) Type() protoreflect.EnumType {
	return &file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_enumTypes[0]
}

func (x SOMAlertsEvent_Alert_AlertType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SOMAlertsEvent_Alert_AlertType.Descriptor instead.
func (SOMAlertsEvent_Alert_AlertType) EnumDescriptor() ([]byte, []int) {
	return file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_rawDescGZIP(), []int{0, 0, 0}
}

type SOMAnnotationEvent_OperationType int32

const (
	SOMAnnotationEvent_ADD    SOMAnnotationEvent_OperationType = 0
	SOMAnnotationEvent_DELETE SOMAnnotationEvent_OperationType = 1
)

// Enum value maps for SOMAnnotationEvent_OperationType.
var (
	SOMAnnotationEvent_OperationType_name = map[int32]string{
		0: "ADD",
		1: "DELETE",
	}
	SOMAnnotationEvent_OperationType_value = map[string]int32{
		"ADD":    0,
		"DELETE": 1,
	}
)

func (x SOMAnnotationEvent_OperationType) Enum() *SOMAnnotationEvent_OperationType {
	p := new(SOMAnnotationEvent_OperationType)
	*p = x
	return p
}

func (x SOMAnnotationEvent_OperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SOMAnnotationEvent_OperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_enumTypes[1].Descriptor()
}

func (SOMAnnotationEvent_OperationType) Type() protoreflect.EnumType {
	return &file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_enumTypes[1]
}

func (x SOMAnnotationEvent_OperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SOMAnnotationEvent_OperationType.Descriptor instead.
func (SOMAnnotationEvent_OperationType) EnumDescriptor() ([]byte, []int) {
	return file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_rawDescGZIP(), []int{1, 0}
}

// Alerts contains alerts generated by sheriff-o-matic analyzer cron jobs.
type SOMAlertsEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tree is the name of the tree.
	Tree string `protobuf:"bytes,1,opt,name=tree,proto3" json:"tree,omitempty"`
	// Timestamp is when the alerts were generated.
	Timestamp *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// RequestId is the id of the incoming http request for the cron handler
	// that generated these alerts. This ID appears in the GAE request logs as
	// protoPayload.requestId.
	RequestId string `protobuf:"bytes,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Alerts are the set of alerts generated by an analyzer cron job run.
	Alerts []*SOMAlertsEvent_Alert `protobuf:"bytes,4,rep,name=Alerts,proto3" json:"Alerts,omitempty"`
}

func (x *SOMAlertsEvent) Reset() {
	*x = SOMAlertsEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SOMAlertsEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SOMAlertsEvent) ProtoMessage() {}

func (x *SOMAlertsEvent) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SOMAlertsEvent.ProtoReflect.Descriptor instead.
func (*SOMAlertsEvent) Descriptor() ([]byte, []int) {
	return file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_rawDescGZIP(), []int{0}
}

func (x *SOMAlertsEvent) GetTree() string {
	if x != nil {
		return x.Tree
	}
	return ""
}

func (x *SOMAlertsEvent) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *SOMAlertsEvent) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *SOMAlertsEvent) GetAlerts() []*SOMAlertsEvent_Alert {
	if x != nil {
		return x.Alerts
	}
	return nil
}

type SOMAnnotationEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tree is the name of the tree.
	Tree string `protobuf:"bytes,1,opt,name=tree,proto3" json:"tree,omitempty"`
	// Timestamp is when the alerts were generated.
	Timestamp *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// RequestId is the id of the incoming http request that generated this
	// annotation. This ID appears in the GAE request logs as protoPayload.requestId.
	RequestId string `protobuf:"bytes,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// User is the ID of the user making the annotation change.
	User string `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	// AlertKeyDigest is the key digest for the alert.
	AlertKeyDigest string `protobuf:"bytes,5,opt,name=alert_key_digest,json=alertKeyDigest,proto3" json:"alert_key_digest,omitempty"`
	// AlertKey is an opaque key for the alert being annotated.
	AlertKey string `protobuf:"bytes,6,opt,name=alert_key,json=alertKey,proto3" json:"alert_key,omitempty"`
	// Operation is the annoation operation.
	Operation SOMAnnotationEvent_OperationType `protobuf:"varint,7,opt,name=operation,proto3,enum=som.events.SOMAnnotationEvent_OperationType" json:"operation,omitempty"`
	// Deprecated, use bug_list instead
	//
	// Deprecated: Do not use.
	Bugs []string `protobuf:"bytes,8,rep,name=bugs,proto3" json:"bugs,omitempty"`
	// bug_list is the list of MonorailBugs attached to the alert.
	BugList []*SOMAnnotationEvent_MonorailBug `protobuf:"bytes,13,rep,name=bug_list,json=bugList,proto3" json:"bug_list,omitempty"`
	// Comments is the list of comments attached to the alert.
	Comments []*SOMAnnotationEvent_Comment `protobuf:"bytes,9,rep,name=comments,proto3" json:"comments,omitempty"`
	// SnoozeTime is the time until which to snooze the alert.
	SnoozeTime *timestamp.Timestamp `protobuf:"bytes,10,opt,name=snooze_time,json=snoozeTime,proto3" json:"snooze_time,omitempty"`
	// GroupId is the name of the alert group to which the alert belongs.
	GroupId string `protobuf:"bytes,11,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// ModificationTime is the time the annotation was modified.
	ModificationTime *timestamp.Timestamp `protobuf:"bytes,12,opt,name=modification_time,json=modificationTime,proto3" json:"modification_time,omitempty"`
}

func (x *SOMAnnotationEvent) Reset() {
	*x = SOMAnnotationEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SOMAnnotationEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SOMAnnotationEvent) ProtoMessage() {}

func (x *SOMAnnotationEvent) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SOMAnnotationEvent.ProtoReflect.Descriptor instead.
func (*SOMAnnotationEvent) Descriptor() ([]byte, []int) {
	return file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_rawDescGZIP(), []int{1}
}

func (x *SOMAnnotationEvent) GetTree() string {
	if x != nil {
		return x.Tree
	}
	return ""
}

func (x *SOMAnnotationEvent) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *SOMAnnotationEvent) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *SOMAnnotationEvent) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *SOMAnnotationEvent) GetAlertKeyDigest() string {
	if x != nil {
		return x.AlertKeyDigest
	}
	return ""
}

func (x *SOMAnnotationEvent) GetAlertKey() string {
	if x != nil {
		return x.AlertKey
	}
	return ""
}

func (x *SOMAnnotationEvent) GetOperation() SOMAnnotationEvent_OperationType {
	if x != nil {
		return x.Operation
	}
	return SOMAnnotationEvent_ADD
}

// Deprecated: Do not use.
func (x *SOMAnnotationEvent) GetBugs() []string {
	if x != nil {
		return x.Bugs
	}
	return nil
}

func (x *SOMAnnotationEvent) GetBugList() []*SOMAnnotationEvent_MonorailBug {
	if x != nil {
		return x.BugList
	}
	return nil
}

func (x *SOMAnnotationEvent) GetComments() []*SOMAnnotationEvent_Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *SOMAnnotationEvent) GetSnoozeTime() *timestamp.Timestamp {
	if x != nil {
		return x.SnoozeTime
	}
	return nil
}

func (x *SOMAnnotationEvent) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *SOMAnnotationEvent) GetModificationTime() *timestamp.Timestamp {
	if x != nil {
		return x.ModificationTime
	}
	return nil
}

type SOMAlertsEvent_Alert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key is an opaque key generated for each alert.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Title is the human-readable title of the alert.
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// Body is the human-readable plain text body of the alert.
	Body string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	// Type is the type of alert.
	Type SOMAlertsEvent_Alert_AlertType `protobuf:"varint,4,opt,name=type,proto3,enum=som.events.SOMAlertsEvent_Alert_AlertType" json:"type,omitempty"`
	// BuildbotFailures contains information about build failures grouped into this alert.
	BuildbotFailures []*SOMAlertsEvent_Alert_BuildbotFailure `protobuf:"bytes,6,rep,name=buildbot_failures,json=buildbotFailures,proto3" json:"buildbot_failures,omitempty"`
}

func (x *SOMAlertsEvent_Alert) Reset() {
	*x = SOMAlertsEvent_Alert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SOMAlertsEvent_Alert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SOMAlertsEvent_Alert) ProtoMessage() {}

func (x *SOMAlertsEvent_Alert) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SOMAlertsEvent_Alert.ProtoReflect.Descriptor instead.
func (*SOMAlertsEvent_Alert) Descriptor() ([]byte, []int) {
	return file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SOMAlertsEvent_Alert) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SOMAlertsEvent_Alert) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SOMAlertsEvent_Alert) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *SOMAlertsEvent_Alert) GetType() SOMAlertsEvent_Alert_AlertType {
	if x != nil {
		return x.Type
	}
	return SOMAlertsEvent_Alert_STALE_MASTER
}

func (x *SOMAlertsEvent_Alert) GetBuildbotFailures() []*SOMAlertsEvent_Alert_BuildbotFailure {
	if x != nil {
		return x.BuildbotFailures
	}
	return nil
}

// BuildBotFaiure describes a range of failing builds on a buildbot builder.
type SOMAlertsEvent_Alert_BuildbotFailure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Master is the name of the buildbot master.
	Master string `protobuf:"bytes,1,opt,name=master,proto3" json:"master,omitempty"`
	// Builder is the name of the builder.
	Builder string `protobuf:"bytes,2,opt,name=builder,proto3" json:"builder,omitempty"`
	// Step is the name of the failing build step.
	Step string `protobuf:"bytes,3,opt,name=step,proto3" json:"step,omitempty"`
	// FirstFailure is the eariest known build number for this run of failures.
	FirstFailure int64 `protobuf:"varint,4,opt,name=first_failure,json=firstFailure,proto3" json:"first_failure,omitempty"`
	// LatestFailure is the latest known build number for this run of failures.
	LatestFailure int64 `protobuf:"varint,5,opt,name=latest_failure,json=latestFailure,proto3" json:"latest_failure,omitempty"`
	// LatestPassing is the latest known build number where this step passed.
	LatestPassing int64 `protobuf:"varint,6,opt,name=latest_passing,json=latestPassing,proto3" json:"latest_passing,omitempty"`
}

func (x *SOMAlertsEvent_Alert_BuildbotFailure) Reset() {
	*x = SOMAlertsEvent_Alert_BuildbotFailure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SOMAlertsEvent_Alert_BuildbotFailure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SOMAlertsEvent_Alert_BuildbotFailure) ProtoMessage() {}

func (x *SOMAlertsEvent_Alert_BuildbotFailure) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SOMAlertsEvent_Alert_BuildbotFailure.ProtoReflect.Descriptor instead.
func (*SOMAlertsEvent_Alert_BuildbotFailure) Descriptor() ([]byte, []int) {
	return file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *SOMAlertsEvent_Alert_BuildbotFailure) GetMaster() string {
	if x != nil {
		return x.Master
	}
	return ""
}

func (x *SOMAlertsEvent_Alert_BuildbotFailure) GetBuilder() string {
	if x != nil {
		return x.Builder
	}
	return ""
}

func (x *SOMAlertsEvent_Alert_BuildbotFailure) GetStep() string {
	if x != nil {
		return x.Step
	}
	return ""
}

func (x *SOMAlertsEvent_Alert_BuildbotFailure) GetFirstFailure() int64 {
	if x != nil {
		return x.FirstFailure
	}
	return 0
}

func (x *SOMAlertsEvent_Alert_BuildbotFailure) GetLatestFailure() int64 {
	if x != nil {
		return x.LatestFailure
	}
	return 0
}

func (x *SOMAlertsEvent_Alert_BuildbotFailure) GetLatestPassing() int64 {
	if x != nil {
		return x.LatestPassing
	}
	return 0
}

type SOMAnnotationEvent_MonorailBug struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BugId     string `protobuf:"bytes,1,opt,name=bug_id,json=bugId,proto3" json:"bug_id,omitempty"`
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *SOMAnnotationEvent_MonorailBug) Reset() {
	*x = SOMAnnotationEvent_MonorailBug{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SOMAnnotationEvent_MonorailBug) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SOMAnnotationEvent_MonorailBug) ProtoMessage() {}

func (x *SOMAnnotationEvent_MonorailBug) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SOMAnnotationEvent_MonorailBug.ProtoReflect.Descriptor instead.
func (*SOMAnnotationEvent_MonorailBug) Descriptor() ([]byte, []int) {
	return file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SOMAnnotationEvent_MonorailBug) GetBugId() string {
	if x != nil {
		return x.BugId
	}
	return ""
}

func (x *SOMAnnotationEvent_MonorailBug) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

// Comment is a user comment attached to an alert.
type SOMAnnotationEvent_Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Text is the text of the comment.
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// Time is the time when the comment was posted.
	Time *timestamp.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *SOMAnnotationEvent_Comment) Reset() {
	*x = SOMAnnotationEvent_Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SOMAnnotationEvent_Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SOMAnnotationEvent_Comment) ProtoMessage() {}

func (x *SOMAnnotationEvent_Comment) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SOMAnnotationEvent_Comment.ProtoReflect.Descriptor instead.
func (*SOMAnnotationEvent_Comment) Descriptor() ([]byte, []int) {
	return file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_rawDescGZIP(), []int{1, 1}
}

func (x *SOMAnnotationEvent_Comment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *SOMAnnotationEvent_Comment) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

var File_infra_appengine_sheriff_o_matic_som_model_gen_events_proto protoreflect.FileDescriptor

var file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x73, 0x68, 0x65, 0x72, 0x69, 0x66, 0x66, 0x2d, 0x6f, 0x2d, 0x6d, 0x61, 0x74, 0x69,
	0x63, 0x2f, 0x73, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x6f,
	0x6d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x05, 0x0a, 0x0e, 0x53, 0x4f,
	0x4d, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x72, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x72, 0x65, 0x65,
	0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x06, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x6f, 0x6d, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x4f, 0x4d, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x06, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x1a, 0xc0, 0x04, 0x0a, 0x05, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x3e, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x73, 0x6f, 0x6d, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x4f, 0x4d, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x5d, 0x0a, 0x11, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x62, 0x6f, 0x74, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x73, 0x6f, 0x6d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x53, 0x4f, 0x4d, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x6f, 0x74, 0x46,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x10, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x6f, 0x74,
	0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x1a, 0xca, 0x01, 0x0a, 0x0f, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x62, 0x6f, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x74,
	0x65, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x66, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x50, 0x61,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x22, 0x8e, 0x01, 0x0a, 0x09, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x4c, 0x45, 0x5f, 0x4d, 0x41, 0x53,
	0x54, 0x45, 0x52, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x48, 0x55, 0x4e, 0x47, 0x5f, 0x42, 0x55,
	0x49, 0x4c, 0x44, 0x45, 0x52, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x46, 0x46, 0x4c, 0x49,
	0x4e, 0x45, 0x5f, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x45, 0x52, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c,
	0x49, 0x44, 0x4c, 0x45, 0x5f, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x45, 0x52, 0x10, 0x03, 0x12, 0x11,
	0x0a, 0x0d, 0x49, 0x4e, 0x46, 0x52, 0x41, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10,
	0x04, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55,
	0x52, 0x45, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x55, 0x52, 0x45, 0x10, 0x06, 0x22, 0xa6, 0x06, 0x0a, 0x12, 0x53, 0x4f, 0x4d, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x72, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x72, 0x65,
	0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x28,
	0x0a, 0x10, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x4b,
	0x65, 0x79, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x4a, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x73, 0x6f, 0x6d, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x4f, 0x4d, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x04, 0x62, 0x75, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x04, 0x62, 0x75, 0x67, 0x73, 0x12, 0x45, 0x0a, 0x08, 0x62, 0x75, 0x67,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x6f,
	0x6d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x4f, 0x4d, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x6e, 0x6f,
	0x72, 0x61, 0x69, 0x6c, 0x42, 0x75, 0x67, 0x52, 0x07, 0x62, 0x75, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x42, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x6f, 0x6d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x53, 0x4f, 0x4d, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x73, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x73, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x11,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x10, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x43, 0x0a, 0x0b, 0x4d, 0x6f, 0x6e, 0x6f, 0x72, 0x61, 0x69,
	0x6c, 0x42, 0x75, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x75, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x75, 0x67, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x4d, 0x0a, 0x07, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x0d, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x44,
	0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x01, 0x42,
	0x2f, 0x5a, 0x2d, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x73, 0x68, 0x65, 0x72, 0x69, 0x66, 0x66, 0x2d, 0x6f, 0x2d, 0x6d, 0x61, 0x74,
	0x69, 0x63, 0x2f, 0x73, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_rawDescOnce sync.Once
	file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_rawDescData = file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_rawDesc
)

func file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_rawDescGZIP() []byte {
	file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_rawDescOnce.Do(func() {
		file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_rawDescData)
	})
	return file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_rawDescData
}

var file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_goTypes = []interface{}{
	(SOMAlertsEvent_Alert_AlertType)(0),          // 0: som.events.SOMAlertsEvent.Alert.AlertType
	(SOMAnnotationEvent_OperationType)(0),        // 1: som.events.SOMAnnotationEvent.OperationType
	(*SOMAlertsEvent)(nil),                       // 2: som.events.SOMAlertsEvent
	(*SOMAnnotationEvent)(nil),                   // 3: som.events.SOMAnnotationEvent
	(*SOMAlertsEvent_Alert)(nil),                 // 4: som.events.SOMAlertsEvent.Alert
	(*SOMAlertsEvent_Alert_BuildbotFailure)(nil), // 5: som.events.SOMAlertsEvent.Alert.BuildbotFailure
	(*SOMAnnotationEvent_MonorailBug)(nil),       // 6: som.events.SOMAnnotationEvent.MonorailBug
	(*SOMAnnotationEvent_Comment)(nil),           // 7: som.events.SOMAnnotationEvent.Comment
	(*timestamp.Timestamp)(nil),                  // 8: google.protobuf.Timestamp
}
var file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_depIdxs = []int32{
	8,  // 0: som.events.SOMAlertsEvent.timestamp:type_name -> google.protobuf.Timestamp
	4,  // 1: som.events.SOMAlertsEvent.Alerts:type_name -> som.events.SOMAlertsEvent.Alert
	8,  // 2: som.events.SOMAnnotationEvent.timestamp:type_name -> google.protobuf.Timestamp
	1,  // 3: som.events.SOMAnnotationEvent.operation:type_name -> som.events.SOMAnnotationEvent.OperationType
	6,  // 4: som.events.SOMAnnotationEvent.bug_list:type_name -> som.events.SOMAnnotationEvent.MonorailBug
	7,  // 5: som.events.SOMAnnotationEvent.comments:type_name -> som.events.SOMAnnotationEvent.Comment
	8,  // 6: som.events.SOMAnnotationEvent.snooze_time:type_name -> google.protobuf.Timestamp
	8,  // 7: som.events.SOMAnnotationEvent.modification_time:type_name -> google.protobuf.Timestamp
	0,  // 8: som.events.SOMAlertsEvent.Alert.type:type_name -> som.events.SOMAlertsEvent.Alert.AlertType
	5,  // 9: som.events.SOMAlertsEvent.Alert.buildbot_failures:type_name -> som.events.SOMAlertsEvent.Alert.BuildbotFailure
	8,  // 10: som.events.SOMAnnotationEvent.Comment.time:type_name -> google.protobuf.Timestamp
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_init() }
func file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_init() {
	if File_infra_appengine_sheriff_o_matic_som_model_gen_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SOMAlertsEvent); i {
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
		file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SOMAnnotationEvent); i {
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
		file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SOMAlertsEvent_Alert); i {
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
		file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SOMAlertsEvent_Alert_BuildbotFailure); i {
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
		file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SOMAnnotationEvent_MonorailBug); i {
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
		file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SOMAnnotationEvent_Comment); i {
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
			RawDescriptor: file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_goTypes,
		DependencyIndexes: file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_depIdxs,
		EnumInfos:         file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_enumTypes,
		MessageInfos:      file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_msgTypes,
	}.Build()
	File_infra_appengine_sheriff_o_matic_som_model_gen_events_proto = out.File
	file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_rawDesc = nil
	file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_goTypes = nil
	file_infra_appengine_sheriff_o_matic_som_model_gen_events_proto_depIdxs = nil
}
