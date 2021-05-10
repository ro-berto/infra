// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.1
// source: infra/appengine/depot_tools_metrics/schema/schema.proto

package schema

import (
	proto "go.chromium.org/luci/buildbucket/proto"
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

// HttpRequest stores information on the HTTP requests made by the command.
type HttpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The host the request was made to. Must be one of the |knownHTTPHosts| in
	// metrics/constants.go.
	// e.g. chromium-review.googlesource.com
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// The HTTP method used to make the request (e.g. GET, POST).
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	// The path and URL arguments of the request.
	// The path must be one of the |knownHTTPPaths| and the arguments must be
	// |knownHTTPArguments| as defined in metrics/constants.go.
	//
	// The URL is not recorded since it might contain PII. Similarly, in most
	// cases, only the name of the arguments (and not their values) are recorded.
	// When the possible values for an argument is a fixed set, as is the case for
	// "o-parameters" in Gerrit, they'll be recorded as arguments.
	// Each argument is recorded separately, so as to make it easier to query.
	//
	// e.g. If the request was to
	// '/changes/?q=owner:foo@example.com+is:open&n=3&o=LABELS&o=ALL_REVISIONS'
	// The path will be '/changes' and the arguments will be 'q', 'n', 'o',
	// 'LABELS' and 'ALL_REVISIONS'.
	Path      string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Arguments []string `protobuf:"bytes,4,rep,name=arguments,proto3" json:"arguments,omitempty"`
	// The HTTP response status.
	Status int64 `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	// The latency of the HTTP request in seconds.
	// TODO(ehmaldonado): Consider converting to google.protobuf.Duration.
	ResponseTime float64 `protobuf:"fixed64,6,opt,name=response_time,json=responseTime,proto3" json:"response_time,omitempty"`
}

func (x *HttpRequest) Reset() {
	*x = HttpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpRequest) ProtoMessage() {}

func (x *HttpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpRequest.ProtoReflect.Descriptor instead.
func (*HttpRequest) Descriptor() ([]byte, []int) {
	return file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescGZIP(), []int{0}
}

func (x *HttpRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *HttpRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *HttpRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *HttpRequest) GetArguments() []string {
	if x != nil {
		return x.Arguments
	}
	return nil
}

func (x *HttpRequest) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *HttpRequest) GetResponseTime() float64 {
	if x != nil {
		return x.ResponseTime
	}
	return 0
}

// SubCommand stores information on the sub-commands executed by the command.
type SubCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The sub-command that was executed. Must be one of the |knownSubCommands| in
	// metrics/constans.go.
	Command string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	// The arguments passed to the sub-command. All arguments must be
	// |knownSubCommandArguments| as defined in metrics/constants.go.
	Arguments []string `protobuf:"bytes,2,rep,name=arguments,proto3" json:"arguments,omitempty"`
	// The runtime of the sub-command runtime in seconds.
	// TODO(ehmaldonado): Consider converting to google.protobuf.Duration.
	ExecutionTime float64 `protobuf:"fixed64,3,opt,name=execution_time,json=executionTime,proto3" json:"execution_time,omitempty"`
	// The exit code of the sub-command.
	ExitCode int64 `protobuf:"varint,4,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
}

func (x *SubCommand) Reset() {
	*x = SubCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubCommand) ProtoMessage() {}

func (x *SubCommand) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubCommand.ProtoReflect.Descriptor instead.
func (*SubCommand) Descriptor() ([]byte, []int) {
	return file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescGZIP(), []int{1}
}

func (x *SubCommand) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *SubCommand) GetArguments() []string {
	if x != nil {
		return x.Arguments
	}
	return nil
}

func (x *SubCommand) GetExecutionTime() float64 {
	if x != nil {
		return x.ExecutionTime
	}
	return 0
}

func (x *SubCommand) GetExitCode() int64 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

// BotMetrics stores information about the bot environment from which the
// command was executed.
type BotMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The build from which this command was executed.
	BuildId int64 `protobuf:"varint,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	// The builder corresponding to the build.
	Builder *proto.BuilderID `protobuf:"bytes,2,opt,name=builder,proto3" json:"builder,omitempty"`
}

func (x *BotMetrics) Reset() {
	*x = BotMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BotMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BotMetrics) ProtoMessage() {}

func (x *BotMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BotMetrics.ProtoReflect.Descriptor instead.
func (*BotMetrics) Descriptor() ([]byte, []int) {
	return file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescGZIP(), []int{2}
}

func (x *BotMetrics) GetBuildId() int64 {
	if x != nil {
		return x.BuildId
	}
	return 0
}

func (x *BotMetrics) GetBuilder() *proto.BuilderID {
	if x != nil {
		return x.Builder
	}
	return nil
}

// Metrics stores information for a depot_tools command's execution.
type Metrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The version of the format used to report the metrics.
	MetricsVersion int64 `protobuf:"varint,1,opt,name=metrics_version,json=metricsVersion,proto3" json:"metrics_version,omitempty"`
	// A UNIX timestamp for the time when the command was executed.
	// TODO(ehmaldonado): Consider converting to google.protobuf.Timestamp.
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The command that was executed. Must be one of the |knownCommands| defined
	// in metrics/constants.go.
	Command string `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"`
	// The arguments passed to the command. All arguments must be |knownArguments|
	// as defined in metrics/constants.go.
	Arguments []string `protobuf:"bytes,4,rep,name=arguments,proto3" json:"arguments,omitempty"`
	// The runtime of the command in seconds.
	// TODO(ehmaldonado): Consider converting to google.protobuf.Duration.
	ExecutionTime float64 `protobuf:"fixed64,5,opt,name=execution_time,json=executionTime,proto3" json:"execution_time,omitempty"`
	// The exit code of the command.
	ExitCode int64 `protobuf:"varint,6,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	// Information on the sub-commands executed by this command.
	SubCommands []*SubCommand `protobuf:"bytes,7,rep,name=sub_commands,json=subCommands,proto3" json:"sub_commands,omitempty"`
	// Information on the HTTP requests made by this command.
	HttpRequests []*HttpRequest `protobuf:"bytes,8,rep,name=http_requests,json=httpRequests,proto3" json:"http_requests,omitempty"`
	// The URLs of the current project(s).
	// e.g. The project to which git-cl uploads a change; the projects gclient is
	// configured to manage; etc.
	// Must be one of the |knownProjectURLs| as defined in metrics/constants.go.
	ProjectUrls []string `protobuf:"bytes,9,rep,name=project_urls,json=projectUrls,proto3" json:"project_urls,omitempty"`
	// A UNIX timestamp for the time depot_tools was last modified.
	// TODO(ehmaldonado): Consider converting to google.protobuf.Timestamp.
	DepotToolsAge float64 `protobuf:"fixed64,10,opt,name=depot_tools_age,json=depotToolsAge,proto3" json:"depot_tools_age,omitempty"`
	// The arch the command was executed on. Must be one of the |knownHostArchs|
	// as defined in metrics/constants.go.
	// e.g. x86, arm
	HostArch string `protobuf:"bytes,11,opt,name=host_arch,json=hostArch,proto3" json:"host_arch,omitempty"`
	// The OS the command was executed on. Must be one of the |knownOSs| as
	// defined in metrics/constants.go.
	HostOs string `protobuf:"bytes,12,opt,name=host_os,json=hostOs,proto3" json:"host_os,omitempty"`
	// The python version the command was executed with. Must match the
	// |pythonVersionRegex| defined in metrics/constants.go.
	PythonVersion string `protobuf:"bytes,13,opt,name=python_version,json=pythonVersion,proto3" json:"python_version,omitempty"`
	// The git version the command used. Must match the |gitVersionRegex| defined
	// in metrics/constants.go.
	GitVersion string      `protobuf:"bytes,14,opt,name=git_version,json=gitVersion,proto3" json:"git_version,omitempty"`
	BotMetrics *BotMetrics `protobuf:"bytes,15,opt,name=bot_metrics,json=botMetrics,proto3" json:"bot_metrics,omitempty"`
}

func (x *Metrics) Reset() {
	*x = Metrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metrics) ProtoMessage() {}

func (x *Metrics) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metrics.ProtoReflect.Descriptor instead.
func (*Metrics) Descriptor() ([]byte, []int) {
	return file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescGZIP(), []int{3}
}

func (x *Metrics) GetMetricsVersion() int64 {
	if x != nil {
		return x.MetricsVersion
	}
	return 0
}

func (x *Metrics) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Metrics) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Metrics) GetArguments() []string {
	if x != nil {
		return x.Arguments
	}
	return nil
}

func (x *Metrics) GetExecutionTime() float64 {
	if x != nil {
		return x.ExecutionTime
	}
	return 0
}

func (x *Metrics) GetExitCode() int64 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

func (x *Metrics) GetSubCommands() []*SubCommand {
	if x != nil {
		return x.SubCommands
	}
	return nil
}

func (x *Metrics) GetHttpRequests() []*HttpRequest {
	if x != nil {
		return x.HttpRequests
	}
	return nil
}

func (x *Metrics) GetProjectUrls() []string {
	if x != nil {
		return x.ProjectUrls
	}
	return nil
}

func (x *Metrics) GetDepotToolsAge() float64 {
	if x != nil {
		return x.DepotToolsAge
	}
	return 0
}

func (x *Metrics) GetHostArch() string {
	if x != nil {
		return x.HostArch
	}
	return ""
}

func (x *Metrics) GetHostOs() string {
	if x != nil {
		return x.HostOs
	}
	return ""
}

func (x *Metrics) GetPythonVersion() string {
	if x != nil {
		return x.PythonVersion
	}
	return ""
}

func (x *Metrics) GetGitVersion() string {
	if x != nil {
		return x.GitVersion
	}
	return ""
}

func (x *Metrics) GetBotMetrics() *BotMetrics {
	if x != nil {
		return x.BotMetrics
	}
	return nil
}

var File_infra_appengine_depot_tools_metrics_schema_schema_proto protoreflect.FileDescriptor

var file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDesc = []byte{
	0x0a, 0x37, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x5f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x5f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x1a, 0x34, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a, 0x0b, 0x48, 0x74, 0x74, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x88, 0x01, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x5c, 0x0a,
	0x0a, 0x42, 0x6f, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x49, 0x44, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x22, 0xbb, 0x04, 0x0a, 0x07,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x67,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x35, 0x0a, 0x0c, 0x73, 0x75,
	0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x75, 0x62, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x12, 0x38, 0x0a, 0x0d, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x68,
	0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x26,
	0x0a, 0x0f, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x5f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x5f, 0x61, 0x67,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x54, 0x6f,
	0x6f, 0x6c, 0x73, 0x41, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x61,
	0x72, 0x63, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x41,
	0x72, 0x63, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6f, 0x73, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x4f, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x69, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x69, 0x74, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x0b, 0x62, 0x6f, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2e, 0x42, 0x6f, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x0a, 0x62,
	0x6f, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x42, 0x33, 0x5a, 0x31, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x64, 0x65, 0x70,
	0x6f, 0x74, 0x5f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x3b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescOnce sync.Once
	file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescData = file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDesc
)

func file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescGZIP() []byte {
	file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescOnce.Do(func() {
		file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescData)
	})
	return file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescData
}

var file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_infra_appengine_depot_tools_metrics_schema_schema_proto_goTypes = []interface{}{
	(*HttpRequest)(nil),     // 0: schema.HttpRequest
	(*SubCommand)(nil),      // 1: schema.SubCommand
	(*BotMetrics)(nil),      // 2: schema.BotMetrics
	(*Metrics)(nil),         // 3: schema.Metrics
	(*proto.BuilderID)(nil), // 4: buildbucket.v2.BuilderID
}
var file_infra_appengine_depot_tools_metrics_schema_schema_proto_depIdxs = []int32{
	4, // 0: schema.BotMetrics.builder:type_name -> buildbucket.v2.BuilderID
	1, // 1: schema.Metrics.sub_commands:type_name -> schema.SubCommand
	0, // 2: schema.Metrics.http_requests:type_name -> schema.HttpRequest
	2, // 3: schema.Metrics.bot_metrics:type_name -> schema.BotMetrics
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_infra_appengine_depot_tools_metrics_schema_schema_proto_init() }
func file_infra_appengine_depot_tools_metrics_schema_schema_proto_init() {
	if File_infra_appengine_depot_tools_metrics_schema_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpRequest); i {
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
		file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubCommand); i {
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
		file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BotMetrics); i {
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
		file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metrics); i {
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
			RawDescriptor: file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_appengine_depot_tools_metrics_schema_schema_proto_goTypes,
		DependencyIndexes: file_infra_appengine_depot_tools_metrics_schema_schema_proto_depIdxs,
		MessageInfos:      file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes,
	}.Build()
	File_infra_appengine_depot_tools_metrics_schema_schema_proto = out.File
	file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDesc = nil
	file_infra_appengine_depot_tools_metrics_schema_schema_proto_goTypes = nil
	file_infra_appengine_depot_tools_metrics_schema_schema_proto_depIdxs = nil
}
