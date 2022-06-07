// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: infra/appengine/weetbix/proto/bq/test_variant_row.proto

package weetbixpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	analyzedtestvariant "infra/appengine/weetbix/proto/analyzedtestvariant"
	v1 "infra/appengine/weetbix/proto/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Verdict represent results of a test variant within an invocation.
type Verdict struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id of the invocation that contains the verdict.
	Invocation string `protobuf:"bytes,1,opt,name=invocation,proto3" json:"invocation,omitempty"`
	// Status of the verdict.
	// String representation of weetbix.v1.VerdictStatus.
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	// Invocation creation time.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"` // TODO: Add information about clusters and bugs.
}

func (x *Verdict) Reset() {
	*x = Verdict{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Verdict) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Verdict) ProtoMessage() {}

func (x *Verdict) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Verdict.ProtoReflect.Descriptor instead.
func (*Verdict) Descriptor() ([]byte, []int) {
	return file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_rawDescGZIP(), []int{0}
}

func (x *Verdict) GetInvocation() string {
	if x != nil {
		return x.Invocation
	}
	return ""
}

func (x *Verdict) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Verdict) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

// TestVariantRow represents a row in a BigQuery table for a Weetbix analyzed
// test variant.
type TestVariantRow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Can be used to refer to this test variant.
	// Format:
	// "realms/{REALM}/tests/{URL_ESCAPED_TEST_ID}/variants/{VARIANT_HASH}"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Realm that the test variant exists under.
	// See https://source.chromium.org/chromium/infra/infra/+/main:go/src/go.chromium.org/luci/common/proto/realms/realms_config.proto
	Realm string `protobuf:"bytes,2,opt,name=realm,proto3" json:"realm,omitempty"`
	// Test id, identifier of the test. Unique in a LUCI realm.
	TestId string `protobuf:"bytes,3,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
	// Hash of the variant.
	VariantHash string `protobuf:"bytes,4,opt,name=variant_hash,json=variantHash,proto3" json:"variant_hash,omitempty"`
	// Description of one specific way of running the test,
	// e.g. a specific bucket, builder and a test suite.
	Variant []*v1.StringPair `protobuf:"bytes,5,rep,name=variant,proto3" json:"variant,omitempty"`
	// Information about the test at the time of its execution.
	TestMetadata *v1.TestMetadata `protobuf:"bytes,6,opt,name=test_metadata,json=testMetadata,proto3" json:"test_metadata,omitempty"`
	// Metadata for the test variant.
	// See luci.resultdb.v1.Tags for details.
	Tags []*v1.StringPair `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	// A range of time. Flake statistics are calculated using test results
	// in the verdicts that were finalized within that range.
	TimeRange *v1.TimeRange `protobuf:"bytes,8,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
	// Status of the test variant.
	// String representation of weetbix.analyzedtestvariant.Status.
	Status string `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	// Flakiness statistics of the test variant.
	FlakeStatistics *analyzedtestvariant.FlakeStatistics `protobuf:"bytes,10,opt,name=flake_statistics,json=flakeStatistics,proto3" json:"flake_statistics,omitempty"`
	// Verdicts of the test variant during the time range.
	Verdicts []*Verdict `protobuf:"bytes,11,rep,name=verdicts,proto3" json:"verdicts,omitempty"`
	// Partition_time is used to partition the table.
	// It's the same as the latest of time_range.
	PartitionTime *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=partition_time,json=partitionTime,proto3" json:"partition_time,omitempty"`
}

func (x *TestVariantRow) Reset() {
	*x = TestVariantRow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestVariantRow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestVariantRow) ProtoMessage() {}

func (x *TestVariantRow) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestVariantRow.ProtoReflect.Descriptor instead.
func (*TestVariantRow) Descriptor() ([]byte, []int) {
	return file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_rawDescGZIP(), []int{1}
}

func (x *TestVariantRow) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestVariantRow) GetRealm() string {
	if x != nil {
		return x.Realm
	}
	return ""
}

func (x *TestVariantRow) GetTestId() string {
	if x != nil {
		return x.TestId
	}
	return ""
}

func (x *TestVariantRow) GetVariantHash() string {
	if x != nil {
		return x.VariantHash
	}
	return ""
}

func (x *TestVariantRow) GetVariant() []*v1.StringPair {
	if x != nil {
		return x.Variant
	}
	return nil
}

func (x *TestVariantRow) GetTestMetadata() *v1.TestMetadata {
	if x != nil {
		return x.TestMetadata
	}
	return nil
}

func (x *TestVariantRow) GetTags() []*v1.StringPair {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *TestVariantRow) GetTimeRange() *v1.TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

func (x *TestVariantRow) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TestVariantRow) GetFlakeStatistics() *analyzedtestvariant.FlakeStatistics {
	if x != nil {
		return x.FlakeStatistics
	}
	return nil
}

func (x *TestVariantRow) GetVerdicts() []*Verdict {
	if x != nil {
		return x.Verdicts
	}
	return nil
}

func (x *TestVariantRow) GetPartitionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.PartitionTime
	}
	return nil
}

var File_infra_appengine_weetbix_proto_bq_test_variant_row_proto protoreflect.FileDescriptor

var file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_rawDesc = []byte{
	0x0a, 0x37, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x62, 0x71, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x5f,
	0x72, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x77, 0x65, 0x65, 0x74, 0x62,
	0x69, 0x78, 0x2e, 0x62, 0x71, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4d, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70,
	0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x64, 0x74, 0x65,
	0x73, 0x74, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a,
	0x65, 0x64, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x64, 0x69, 0x63, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0xae, 0x04, 0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x56, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x52, 0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x65, 0x61, 0x6c, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x61, 0x6c,
	0x6d, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x30, 0x0a,
	0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x50, 0x61, 0x69, 0x72, 0x52, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12,
	0x3d, 0x0a, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x0c, 0x74, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2a,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x77,
	0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x50, 0x61, 0x69, 0x72, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x34, 0x0a, 0x0a, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x57, 0x0a, 0x10, 0x66, 0x6c, 0x61, 0x6b,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x7a, 0x65, 0x64, 0x74, 0x65, 0x73, 0x74, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x2e, 0x46, 0x6c, 0x61, 0x6b, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x52, 0x0f, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x12, 0x2f, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x64, 0x69, 0x63, 0x74, 0x73, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x62, 0x71,
	0x2e, 0x56, 0x65, 0x72, 0x64, 0x69, 0x63, 0x74, 0x52, 0x08, 0x76, 0x65, 0x72, 0x64, 0x69, 0x63,
	0x74, 0x73, 0x12, 0x41, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61,
	0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x71, 0x3b, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69,
	0x78, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_rawDescOnce sync.Once
	file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_rawDescData = file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_rawDesc
)

func file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_rawDescGZIP() []byte {
	file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_rawDescOnce.Do(func() {
		file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_rawDescData)
	})
	return file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_rawDescData
}

var file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_goTypes = []interface{}{
	(*Verdict)(nil),                             // 0: weetbix.bq.Verdict
	(*TestVariantRow)(nil),                      // 1: weetbix.bq.TestVariantRow
	(*timestamppb.Timestamp)(nil),               // 2: google.protobuf.Timestamp
	(*v1.StringPair)(nil),                       // 3: weetbix.v1.StringPair
	(*v1.TestMetadata)(nil),                     // 4: weetbix.v1.TestMetadata
	(*v1.TimeRange)(nil),                        // 5: weetbix.v1.TimeRange
	(*analyzedtestvariant.FlakeStatistics)(nil), // 6: weetbix.analyzedtestvariant.FlakeStatistics
}
var file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_depIdxs = []int32{
	2, // 0: weetbix.bq.Verdict.create_time:type_name -> google.protobuf.Timestamp
	3, // 1: weetbix.bq.TestVariantRow.variant:type_name -> weetbix.v1.StringPair
	4, // 2: weetbix.bq.TestVariantRow.test_metadata:type_name -> weetbix.v1.TestMetadata
	3, // 3: weetbix.bq.TestVariantRow.tags:type_name -> weetbix.v1.StringPair
	5, // 4: weetbix.bq.TestVariantRow.time_range:type_name -> weetbix.v1.TimeRange
	6, // 5: weetbix.bq.TestVariantRow.flake_statistics:type_name -> weetbix.analyzedtestvariant.FlakeStatistics
	0, // 6: weetbix.bq.TestVariantRow.verdicts:type_name -> weetbix.bq.Verdict
	2, // 7: weetbix.bq.TestVariantRow.partition_time:type_name -> google.protobuf.Timestamp
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_init() }
func file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_init() {
	if File_infra_appengine_weetbix_proto_bq_test_variant_row_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Verdict); i {
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
		file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestVariantRow); i {
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
			RawDescriptor: file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_goTypes,
		DependencyIndexes: file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_depIdxs,
		MessageInfos:      file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_msgTypes,
	}.Build()
	File_infra_appengine_weetbix_proto_bq_test_variant_row_proto = out.File
	file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_rawDesc = nil
	file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_goTypes = nil
	file_infra_appengine_weetbix_proto_bq_test_variant_row_proto_depIdxs = nil
}
