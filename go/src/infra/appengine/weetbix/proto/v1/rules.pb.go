// Copyright 2022 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: infra/appengine/weetbix/proto/v1/rules.proto

package weetbixpb

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

// A rule associating failures with a bug.
type Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Can be used to refer to this rule, e.g. in RulesService.Get RPC.
	// Format: projects/{project}/rules/{rule_id}.
	// See also https://google.aip.dev/122.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The LUCI Project for which this rule is defined.
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// The unique identifier for the failure association rule,
	// as 32 lowercase hexadecimal characters.
	RuleId string `protobuf:"bytes,3,opt,name=rule_id,json=ruleId,proto3" json:"rule_id,omitempty"`
	// The rule predicate, defining which failures are being associated.
	RuleDefinition string `protobuf:"bytes,4,opt,name=rule_definition,json=ruleDefinition,proto3" json:"rule_definition,omitempty"`
	// The bug that the failures are associated with.
	Bug *AssociatedBug `protobuf:"bytes,5,opt,name=bug,proto3" json:"bug,omitempty"`
	// Whether the bug should be updated by Weetbix, and whether failures
	// should still be matched against the rule.
	IsActive bool `protobuf:"varint,6,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	// The suggested cluster this rule was created from (if any).
	// Until re-clustering is complete and has reduced the residual impact
	// of the source cluster, this cluster ID tells bug filing to ignore
	// the source cluster when determining whether new bugs need to be filed.
	SourceCluster *ClusterId `protobuf:"bytes,7,opt,name=source_cluster,json=sourceCluster,proto3" json:"source_cluster,omitempty"`
	// The time the rule was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The user which created the rule.
	CreateUser string `protobuf:"bytes,9,opt,name=create_user,json=createUser,proto3" json:"create_user,omitempty"`
	// The time the rule was last updated.
	LastUpdateTime *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=last_update_time,json=lastUpdateTime,proto3" json:"last_update_time,omitempty"`
	// The user which last updated the rule.
	LastUpdateUser string `protobuf:"bytes,11,opt,name=last_update_user,json=lastUpdateUser,proto3" json:"last_update_user,omitempty"`
	// This checksum is computed by the server based on the value of other
	// fields, and may be sent on update requests to ensure the client
	// has an up-to-date value before proceeding.
	// See also https://google.aip.dev/154.
	Etag string `protobuf:"bytes,12,opt,name=etag,proto3" json:"etag,omitempty"`
}

func (x *Rule) Reset() {
	*x = Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_weetbix_proto_v1_rules_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

func (x *Rule) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_weetbix_proto_v1_rules_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rule.ProtoReflect.Descriptor instead.
func (*Rule) Descriptor() ([]byte, []int) {
	return file_infra_appengine_weetbix_proto_v1_rules_proto_rawDescGZIP(), []int{0}
}

func (x *Rule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Rule) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *Rule) GetRuleId() string {
	if x != nil {
		return x.RuleId
	}
	return ""
}

func (x *Rule) GetRuleDefinition() string {
	if x != nil {
		return x.RuleDefinition
	}
	return ""
}

func (x *Rule) GetBug() *AssociatedBug {
	if x != nil {
		return x.Bug
	}
	return nil
}

func (x *Rule) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Rule) GetSourceCluster() *ClusterId {
	if x != nil {
		return x.SourceCluster
	}
	return nil
}

func (x *Rule) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Rule) GetCreateUser() string {
	if x != nil {
		return x.CreateUser
	}
	return ""
}

func (x *Rule) GetLastUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdateTime
	}
	return nil
}

func (x *Rule) GetLastUpdateUser() string {
	if x != nil {
		return x.LastUpdateUser
	}
	return ""
}

func (x *Rule) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

type GetRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the rule to retrieve.
	// Format: projects/{project}/rules/{rule_id}.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetRuleRequest) Reset() {
	*x = GetRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_weetbix_proto_v1_rules_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRuleRequest) ProtoMessage() {}

func (x *GetRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_weetbix_proto_v1_rules_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRuleRequest.ProtoReflect.Descriptor instead.
func (*GetRuleRequest) Descriptor() ([]byte, []int) {
	return file_infra_appengine_weetbix_proto_v1_rules_proto_rawDescGZIP(), []int{1}
}

func (x *GetRuleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListRulesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The parent, which owns this collection of rules.
	// Format: projects/{project}.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
}

func (x *ListRulesRequest) Reset() {
	*x = ListRulesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_weetbix_proto_v1_rules_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRulesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRulesRequest) ProtoMessage() {}

func (x *ListRulesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_weetbix_proto_v1_rules_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRulesRequest.ProtoReflect.Descriptor instead.
func (*ListRulesRequest) Descriptor() ([]byte, []int) {
	return file_infra_appengine_weetbix_proto_v1_rules_proto_rawDescGZIP(), []int{2}
}

func (x *ListRulesRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

type ListRulesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The rules.
	Rules []*Rule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *ListRulesResponse) Reset() {
	*x = ListRulesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_weetbix_proto_v1_rules_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRulesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRulesResponse) ProtoMessage() {}

func (x *ListRulesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_weetbix_proto_v1_rules_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRulesResponse.ProtoReflect.Descriptor instead.
func (*ListRulesResponse) Descriptor() ([]byte, []int) {
	return file_infra_appengine_weetbix_proto_v1_rules_proto_rawDescGZIP(), []int{3}
}

func (x *ListRulesResponse) GetRules() []*Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

type CreateRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The parent resource where the rule will be created.
	// Format: projects/{project}.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The rule to create.
	Rule *Rule `protobuf:"bytes,2,opt,name=rule,proto3" json:"rule,omitempty"`
}

func (x *CreateRuleRequest) Reset() {
	*x = CreateRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_weetbix_proto_v1_rules_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRuleRequest) ProtoMessage() {}

func (x *CreateRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_weetbix_proto_v1_rules_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRuleRequest.ProtoReflect.Descriptor instead.
func (*CreateRuleRequest) Descriptor() ([]byte, []int) {
	return file_infra_appengine_weetbix_proto_v1_rules_proto_rawDescGZIP(), []int{4}
}

func (x *CreateRuleRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateRuleRequest) GetRule() *Rule {
	if x != nil {
		return x.Rule
	}
	return nil
}

type UpdateRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The rule to update.
	//
	// The rule's `name` field is used to identify the book to update.
	// Format: projects/{project}/rules/{rule_id}.
	Rule *Rule `protobuf:"bytes,1,opt,name=rule,proto3" json:"rule,omitempty"`
	// The list of fields to update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The current etag of the rule.
	// If an etag is provided and does not match the current etag of the rule,
	// update will be blocked and an ABORTED error will be returned.
	Etag string `protobuf:"bytes,3,opt,name=etag,proto3" json:"etag,omitempty"`
}

func (x *UpdateRuleRequest) Reset() {
	*x = UpdateRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_weetbix_proto_v1_rules_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRuleRequest) ProtoMessage() {}

func (x *UpdateRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_weetbix_proto_v1_rules_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRuleRequest.ProtoReflect.Descriptor instead.
func (*UpdateRuleRequest) Descriptor() ([]byte, []int) {
	return file_infra_appengine_weetbix_proto_v1_rules_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateRuleRequest) GetRule() *Rule {
	if x != nil {
		return x.Rule
	}
	return nil
}

func (x *UpdateRuleRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *UpdateRuleRequest) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

var File_infra_appengine_weetbix_proto_v1_rules_proto protoreflect.FileDescriptor

var file_infra_appengine_weetbix_proto_v1_rules_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f,
	0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x04,
	0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x07, 0x72, 0x75, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x06, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x0f, 0x72, 0x75, 0x6c, 0x65, 0x5f,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e, 0x72, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x03, 0x62, 0x75, 0x67, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x42, 0x75, 0x67, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x03, 0x62, 0x75, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x12, 0x3c, 0x0a, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x77,
	0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x52, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x10, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x22, 0x3b, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x75, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x77, 0x65, 0x65, 0x74, 0x62,
	0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x22, 0x5b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x75, 0x6c, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x22, 0x8f,
	0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x75, 0x6c, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x12,
	0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04,
	0x65, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67,
	0x32, 0xff, 0x01, 0x0a, 0x05, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x1a, 0x2e, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x22,
	0x00, 0x12, 0x45, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x77, 0x65, 0x65, 0x74,
	0x62, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69,
	0x78, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x1d, 0x2e, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x75, 0x6c, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1d, 0x2e, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65,
	0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_appengine_weetbix_proto_v1_rules_proto_rawDescOnce sync.Once
	file_infra_appengine_weetbix_proto_v1_rules_proto_rawDescData = file_infra_appengine_weetbix_proto_v1_rules_proto_rawDesc
)

func file_infra_appengine_weetbix_proto_v1_rules_proto_rawDescGZIP() []byte {
	file_infra_appengine_weetbix_proto_v1_rules_proto_rawDescOnce.Do(func() {
		file_infra_appengine_weetbix_proto_v1_rules_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_appengine_weetbix_proto_v1_rules_proto_rawDescData)
	})
	return file_infra_appengine_weetbix_proto_v1_rules_proto_rawDescData
}

var file_infra_appengine_weetbix_proto_v1_rules_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_infra_appengine_weetbix_proto_v1_rules_proto_goTypes = []interface{}{
	(*Rule)(nil),                  // 0: weetbix.v1.Rule
	(*GetRuleRequest)(nil),        // 1: weetbix.v1.GetRuleRequest
	(*ListRulesRequest)(nil),      // 2: weetbix.v1.ListRulesRequest
	(*ListRulesResponse)(nil),     // 3: weetbix.v1.ListRulesResponse
	(*CreateRuleRequest)(nil),     // 4: weetbix.v1.CreateRuleRequest
	(*UpdateRuleRequest)(nil),     // 5: weetbix.v1.UpdateRuleRequest
	(*AssociatedBug)(nil),         // 6: weetbix.v1.AssociatedBug
	(*ClusterId)(nil),             // 7: weetbix.v1.ClusterId
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*fieldmaskpb.FieldMask)(nil), // 9: google.protobuf.FieldMask
}
var file_infra_appengine_weetbix_proto_v1_rules_proto_depIdxs = []int32{
	6,  // 0: weetbix.v1.Rule.bug:type_name -> weetbix.v1.AssociatedBug
	7,  // 1: weetbix.v1.Rule.source_cluster:type_name -> weetbix.v1.ClusterId
	8,  // 2: weetbix.v1.Rule.create_time:type_name -> google.protobuf.Timestamp
	8,  // 3: weetbix.v1.Rule.last_update_time:type_name -> google.protobuf.Timestamp
	0,  // 4: weetbix.v1.ListRulesResponse.rules:type_name -> weetbix.v1.Rule
	0,  // 5: weetbix.v1.CreateRuleRequest.rule:type_name -> weetbix.v1.Rule
	0,  // 6: weetbix.v1.UpdateRuleRequest.rule:type_name -> weetbix.v1.Rule
	9,  // 7: weetbix.v1.UpdateRuleRequest.update_mask:type_name -> google.protobuf.FieldMask
	1,  // 8: weetbix.v1.Rules.Get:input_type -> weetbix.v1.GetRuleRequest
	2,  // 9: weetbix.v1.Rules.List:input_type -> weetbix.v1.ListRulesRequest
	4,  // 10: weetbix.v1.Rules.Create:input_type -> weetbix.v1.CreateRuleRequest
	5,  // 11: weetbix.v1.Rules.Update:input_type -> weetbix.v1.UpdateRuleRequest
	0,  // 12: weetbix.v1.Rules.Get:output_type -> weetbix.v1.Rule
	3,  // 13: weetbix.v1.Rules.List:output_type -> weetbix.v1.ListRulesResponse
	0,  // 14: weetbix.v1.Rules.Create:output_type -> weetbix.v1.Rule
	0,  // 15: weetbix.v1.Rules.Update:output_type -> weetbix.v1.Rule
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_infra_appengine_weetbix_proto_v1_rules_proto_init() }
func file_infra_appengine_weetbix_proto_v1_rules_proto_init() {
	if File_infra_appengine_weetbix_proto_v1_rules_proto != nil {
		return
	}
	file_infra_appengine_weetbix_proto_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_infra_appengine_weetbix_proto_v1_rules_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rule); i {
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
		file_infra_appengine_weetbix_proto_v1_rules_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRuleRequest); i {
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
		file_infra_appengine_weetbix_proto_v1_rules_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRulesRequest); i {
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
		file_infra_appengine_weetbix_proto_v1_rules_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRulesResponse); i {
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
		file_infra_appengine_weetbix_proto_v1_rules_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRuleRequest); i {
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
		file_infra_appengine_weetbix_proto_v1_rules_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRuleRequest); i {
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
			RawDescriptor: file_infra_appengine_weetbix_proto_v1_rules_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infra_appengine_weetbix_proto_v1_rules_proto_goTypes,
		DependencyIndexes: file_infra_appengine_weetbix_proto_v1_rules_proto_depIdxs,
		MessageInfos:      file_infra_appengine_weetbix_proto_v1_rules_proto_msgTypes,
	}.Build()
	File_infra_appengine_weetbix_proto_v1_rules_proto = out.File
	file_infra_appengine_weetbix_proto_v1_rules_proto_rawDesc = nil
	file_infra_appengine_weetbix_proto_v1_rules_proto_goTypes = nil
	file_infra_appengine_weetbix_proto_v1_rules_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RulesClient is the client API for Rules service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RulesClient interface {
	// Retrieves a rule.
	// Designed to conform to https://google.aip.dev/131.
	Get(ctx context.Context, in *GetRuleRequest, opts ...grpc.CallOption) (*Rule, error)
	// Lists rules.
	// TODO: implement pagination to make this
	// RPC compliant with https://google.aip.dev/132.
	// This RPC is incomplete. Future breaking changes are
	// expressly flagged.
	List(ctx context.Context, in *ListRulesRequest, opts ...grpc.CallOption) (*ListRulesResponse, error)
	// Creates a new rule.
	// Designed to conform to https://google.aip.dev/133.
	Create(ctx context.Context, in *CreateRuleRequest, opts ...grpc.CallOption) (*Rule, error)
	// Updates a rule.
	// Designed to conform to https://google.aip.dev/134.
	Update(ctx context.Context, in *UpdateRuleRequest, opts ...grpc.CallOption) (*Rule, error)
}
type rulesPRPCClient struct {
	client *prpc.Client
}

func NewRulesPRPCClient(client *prpc.Client) RulesClient {
	return &rulesPRPCClient{client}
}

func (c *rulesPRPCClient) Get(ctx context.Context, in *GetRuleRequest, opts ...grpc.CallOption) (*Rule, error) {
	out := new(Rule)
	err := c.client.Call(ctx, "weetbix.v1.Rules", "Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesPRPCClient) List(ctx context.Context, in *ListRulesRequest, opts ...grpc.CallOption) (*ListRulesResponse, error) {
	out := new(ListRulesResponse)
	err := c.client.Call(ctx, "weetbix.v1.Rules", "List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesPRPCClient) Create(ctx context.Context, in *CreateRuleRequest, opts ...grpc.CallOption) (*Rule, error) {
	out := new(Rule)
	err := c.client.Call(ctx, "weetbix.v1.Rules", "Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesPRPCClient) Update(ctx context.Context, in *UpdateRuleRequest, opts ...grpc.CallOption) (*Rule, error) {
	out := new(Rule)
	err := c.client.Call(ctx, "weetbix.v1.Rules", "Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type rulesClient struct {
	cc grpc.ClientConnInterface
}

func NewRulesClient(cc grpc.ClientConnInterface) RulesClient {
	return &rulesClient{cc}
}

func (c *rulesClient) Get(ctx context.Context, in *GetRuleRequest, opts ...grpc.CallOption) (*Rule, error) {
	out := new(Rule)
	err := c.cc.Invoke(ctx, "/weetbix.v1.Rules/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) List(ctx context.Context, in *ListRulesRequest, opts ...grpc.CallOption) (*ListRulesResponse, error) {
	out := new(ListRulesResponse)
	err := c.cc.Invoke(ctx, "/weetbix.v1.Rules/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) Create(ctx context.Context, in *CreateRuleRequest, opts ...grpc.CallOption) (*Rule, error) {
	out := new(Rule)
	err := c.cc.Invoke(ctx, "/weetbix.v1.Rules/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) Update(ctx context.Context, in *UpdateRuleRequest, opts ...grpc.CallOption) (*Rule, error) {
	out := new(Rule)
	err := c.cc.Invoke(ctx, "/weetbix.v1.Rules/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RulesServer is the server API for Rules service.
type RulesServer interface {
	// Retrieves a rule.
	// Designed to conform to https://google.aip.dev/131.
	Get(context.Context, *GetRuleRequest) (*Rule, error)
	// Lists rules.
	// TODO: implement pagination to make this
	// RPC compliant with https://google.aip.dev/132.
	// This RPC is incomplete. Future breaking changes are
	// expressly flagged.
	List(context.Context, *ListRulesRequest) (*ListRulesResponse, error)
	// Creates a new rule.
	// Designed to conform to https://google.aip.dev/133.
	Create(context.Context, *CreateRuleRequest) (*Rule, error)
	// Updates a rule.
	// Designed to conform to https://google.aip.dev/134.
	Update(context.Context, *UpdateRuleRequest) (*Rule, error)
}

// UnimplementedRulesServer can be embedded to have forward compatible implementations.
type UnimplementedRulesServer struct {
}

func (*UnimplementedRulesServer) Get(context.Context, *GetRuleRequest) (*Rule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedRulesServer) List(context.Context, *ListRulesRequest) (*ListRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedRulesServer) Create(context.Context, *CreateRuleRequest) (*Rule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedRulesServer) Update(context.Context, *UpdateRuleRequest) (*Rule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterRulesServer(s prpc.Registrar, srv RulesServer) {
	s.RegisterService(&_Rules_serviceDesc, srv)
}

func _Rules_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weetbix.v1.Rules/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).Get(ctx, req.(*GetRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weetbix.v1.Rules/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).List(ctx, req.(*ListRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weetbix.v1.Rules/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).Create(ctx, req.(*CreateRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weetbix.v1.Rules/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).Update(ctx, req.(*UpdateRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rules_serviceDesc = grpc.ServiceDesc{
	ServiceName: "weetbix.v1.Rules",
	HandlerType: (*RulesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Rules_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Rules_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Rules_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Rules_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infra/appengine/weetbix/proto/v1/rules.proto",
}
