// Copyright 2016 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: infra/tricium/api/v1/data.proto

package tricium

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

// Available data types should be listed in this enum and have a
// corresponding nested message with a mandatory platforms fields,
// see GitFileDetails for field details.
type Data_Type int32

const (
	Data_NONE             Data_Type = 0
	Data_GIT_FILE_DETAILS Data_Type = 1
	Data_FILES            Data_Type = 2
	Data_RESULTS          Data_Type = 4
)

// Enum value maps for Data_Type.
var (
	Data_Type_name = map[int32]string{
		0: "NONE",
		1: "GIT_FILE_DETAILS",
		2: "FILES",
		4: "RESULTS",
	}
	Data_Type_value = map[string]int32{
		"NONE":             0,
		"GIT_FILE_DETAILS": 1,
		"FILES":            2,
		"RESULTS":          4,
	}
)

func (x Data_Type) Enum() *Data_Type {
	p := new(Data_Type)
	*p = x
	return p
}

func (x Data_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Data_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_tricium_api_v1_data_proto_enumTypes[0].Descriptor()
}

func (Data_Type) Type() protoreflect.EnumType {
	return &file_infra_tricium_api_v1_data_proto_enumTypes[0]
}

func (x Data_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Data_Type.Descriptor instead.
func (Data_Type) EnumDescriptor() ([]byte, []int) {
	return file_infra_tricium_api_v1_data_proto_rawDescGZIP(), []int{0, 0}
}

// File change status.
//
// This corresponds to the status field provided by Gerrit in FileInfo:
// https://goo.gl/ABFHDg
type Data_Status int32

const (
	Data_MODIFIED  Data_Status = 0
	Data_ADDED     Data_Status = 1
	Data_DELETED   Data_Status = 2
	Data_RENAMED   Data_Status = 3
	Data_COPIED    Data_Status = 4
	Data_REWRITTEN Data_Status = 5
)

// Enum value maps for Data_Status.
var (
	Data_Status_name = map[int32]string{
		0: "MODIFIED",
		1: "ADDED",
		2: "DELETED",
		3: "RENAMED",
		4: "COPIED",
		5: "REWRITTEN",
	}
	Data_Status_value = map[string]int32{
		"MODIFIED":  0,
		"ADDED":     1,
		"DELETED":   2,
		"RENAMED":   3,
		"COPIED":    4,
		"REWRITTEN": 5,
	}
)

func (x Data_Status) Enum() *Data_Status {
	p := new(Data_Status)
	*p = x
	return p
}

func (x Data_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Data_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_tricium_api_v1_data_proto_enumTypes[1].Descriptor()
}

func (Data_Status) Type() protoreflect.EnumType {
	return &file_infra_tricium_api_v1_data_proto_enumTypes[1]
}

func (x Data_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Data_Status.Descriptor instead.
func (Data_Status) EnumDescriptor() ([]byte, []int) {
	return file_infra_tricium_api_v1_data_proto_rawDescGZIP(), []int{0, 1}
}

// Tricium data types.
//
// Any data type provided or needed by a Tricium function.
type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_tricium_api_v1_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_infra_tricium_api_v1_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_infra_tricium_api_v1_data_proto_rawDescGZIP(), []int{0}
}

// Details for supported types, specifically whether a type is tied to
// a platform.
//
// These type details are used to resolve data dependencies when
// generating workflows.
type Data_TypeDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type               Data_Type `protobuf:"varint,1,opt,name=type,proto3,enum=tricium.Data_Type" json:"type,omitempty"`
	IsPlatformSpecific bool      `protobuf:"varint,2,opt,name=is_platform_specific,json=isPlatformSpecific,proto3" json:"is_platform_specific,omitempty"`
}

func (x *Data_TypeDetails) Reset() {
	*x = Data_TypeDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_tricium_api_v1_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_TypeDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_TypeDetails) ProtoMessage() {}

func (x *Data_TypeDetails) ProtoReflect() protoreflect.Message {
	mi := &file_infra_tricium_api_v1_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_TypeDetails.ProtoReflect.Descriptor instead.
func (*Data_TypeDetails) Descriptor() ([]byte, []int) {
	return file_infra_tricium_api_v1_data_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Data_TypeDetails) GetType() Data_Type {
	if x != nil {
		return x.Type
	}
	return Data_NONE
}

func (x *Data_TypeDetails) GetIsPlatformSpecific() bool {
	if x != nil {
		return x.IsPlatformSpecific
	}
	return false
}

// Details for retrieval of file content from a Git repository.
//
// In practice this was only used as an input to GitFileDetails,
// and is now DEPRECATED.
//
// PATH: tricium/data/git_file_details.json
type Data_GitFileDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The platforms this data is tied to encoded as a bitmap.
	//
	// The bit number for each platform should correspond to the enum
	// position number of the same platform in the Platform.Name enum.
	//
	// This includes the ANY platform, encoded as zero, which should
	// be used for any data that is not platform-specific.
	Platforms     int64        `protobuf:"varint,1,opt,name=platforms,proto3" json:"platforms,omitempty"`
	Repository    string       `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	Ref           string       `protobuf:"bytes,3,opt,name=ref,proto3" json:"ref,omitempty"`
	Files         []*Data_File `protobuf:"bytes,4,rep,name=files,proto3" json:"files,omitempty"`
	CommitMessage string       `protobuf:"bytes,5,opt,name=commit_message,json=commitMessage,proto3" json:"commit_message,omitempty"`
}

func (x *Data_GitFileDetails) Reset() {
	*x = Data_GitFileDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_tricium_api_v1_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_GitFileDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_GitFileDetails) ProtoMessage() {}

func (x *Data_GitFileDetails) ProtoReflect() protoreflect.Message {
	mi := &file_infra_tricium_api_v1_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_GitFileDetails.ProtoReflect.Descriptor instead.
func (*Data_GitFileDetails) Descriptor() ([]byte, []int) {
	return file_infra_tricium_api_v1_data_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Data_GitFileDetails) GetPlatforms() int64 {
	if x != nil {
		return x.Platforms
	}
	return 0
}

func (x *Data_GitFileDetails) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *Data_GitFileDetails) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *Data_GitFileDetails) GetFiles() []*Data_File {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *Data_GitFileDetails) GetCommitMessage() string {
	if x != nil {
		return x.CommitMessage
	}
	return ""
}

// List of paths included in the analyzer input.
//
// PATH: tricium/data/files.json
type Data_Files struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Platforms     int64        `protobuf:"varint,1,opt,name=platforms,proto3" json:"platforms,omitempty"`
	Files         []*Data_File `protobuf:"bytes,3,rep,name=files,proto3" json:"files,omitempty"`
	CommitMessage string       `protobuf:"bytes,4,opt,name=commit_message,json=commitMessage,proto3" json:"commit_message,omitempty"`
}

func (x *Data_Files) Reset() {
	*x = Data_Files{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_tricium_api_v1_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Files) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Files) ProtoMessage() {}

func (x *Data_Files) ProtoReflect() protoreflect.Message {
	mi := &file_infra_tricium_api_v1_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Files.ProtoReflect.Descriptor instead.
func (*Data_Files) Descriptor() ([]byte, []int) {
	return file_infra_tricium_api_v1_data_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Data_Files) GetPlatforms() int64 {
	if x != nil {
		return x.Platforms
	}
	return 0
}

func (x *Data_Files) GetFiles() []*Data_File {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *Data_Files) GetCommitMessage() string {
	if x != nil {
		return x.CommitMessage
	}
	return ""
}

type Data_File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Relative file path.
	//
	// The path is relative to the root of the repository being analyzed,
	// and the path separator character is "/".
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Whether or not this file contains binary content (not text).
	IsBinary bool `protobuf:"varint,2,opt,name=is_binary,json=isBinary,proto3" json:"is_binary,omitempty"`
	// How the file was changed.
	Status Data_Status `protobuf:"varint,3,opt,name=status,proto3,enum=tricium.Data_Status" json:"status,omitempty"`
}

func (x *Data_File) Reset() {
	*x = Data_File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_tricium_api_v1_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_File) ProtoMessage() {}

func (x *Data_File) ProtoReflect() protoreflect.Message {
	mi := &file_infra_tricium_api_v1_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_File.ProtoReflect.Descriptor instead.
func (*Data_File) Descriptor() ([]byte, []int) {
	return file_infra_tricium_api_v1_data_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Data_File) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Data_File) GetIsBinary() bool {
	if x != nil {
		return x.IsBinary
	}
	return false
}

func (x *Data_File) GetStatus() Data_Status {
	if x != nil {
		return x.Status
	}
	return Data_MODIFIED
}

// Results from running a Tricium analyzer.
//
// Results are returned to the Tricium service from Buildbucket
// properties on executed Tricium recipes.
//
// PATH: tricium/data/results.json
// BUILDBUCKET PROPERTIES: output.properties.comments
//
//	output.properties.num_comments
type Data_Results struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Platforms int64 `protobuf:"varint,1,opt,name=platforms,proto3" json:"platforms,omitempty"`
	// Zero or more results found as comments, either inline comments or change
	// comments (comments without line positions).
	Comments []*Data_Comment `protobuf:"bytes,2,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *Data_Results) Reset() {
	*x = Data_Results{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_tricium_api_v1_data_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Results) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Results) ProtoMessage() {}

func (x *Data_Results) ProtoReflect() protoreflect.Message {
	mi := &file_infra_tricium_api_v1_data_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Results.ProtoReflect.Descriptor instead.
func (*Data_Results) Descriptor() ([]byte, []int) {
	return file_infra_tricium_api_v1_data_proto_rawDescGZIP(), []int{0, 4}
}

func (x *Data_Results) GetPlatforms() int64 {
	if x != nil {
		return x.Platforms
	}
	return 0
}

func (x *Data_Results) GetComments() []*Data_Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

// Results.Comment, results as comments.
//
// Similar content as that needed to provide robot comments in Gerrit,
// https://gerrit-review.googlesource.com/Documentation/config-robot-comments.html
type Data_Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Comment ID.
	//
	// This is an UUID generated by the Tricium service and used for tracking
	// of comment feedback. Analyzers should leave this field empty.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Category of the result, encoded as a path with the analyzer name as the
	// root, followed by an arbitrary number of subcategories, for example
	// "ClangTidy/llvm-header-guard".
	Category string `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	// Comment message. This should be a short message suitable as a code
	// review comment.
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// Path to the file this comment is for.
	//
	// If this path is the empty string, then the comment is on the commit
	// message text, rather than an actual file.
	Path string `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	// Position information. If start_line is omitted, then the comment
	// will be a file-level comment.
	StartLine int32 `protobuf:"varint,6,opt,name=start_line,json=startLine,proto3" json:"start_line,omitempty"` // 1-based, inclusive.
	EndLine   int32 `protobuf:"varint,7,opt,name=end_line,json=endLine,proto3" json:"end_line,omitempty"`       // 1-based, inclusive.
	StartChar int32 `protobuf:"varint,8,opt,name=start_char,json=startChar,proto3" json:"start_char,omitempty"` // 0-based, inclusive.
	EndChar   int32 `protobuf:"varint,9,opt,name=end_char,json=endChar,proto3" json:"end_char,omitempty"`       // 0-based, exclusive.
	// Suggested fixes for the identified issue.
	Suggestions []*Data_Suggestion `protobuf:"bytes,10,rep,name=suggestions,proto3" json:"suggestions,omitempty"`
	// When true, show on both changed and unchanged lines.
	// When false, only show on changed lines.
	ShowOnUnchangedLines bool `protobuf:"varint,11,opt,name=show_on_unchanged_lines,json=showOnUnchangedLines,proto3" json:"show_on_unchanged_lines,omitempty"`
}

func (x *Data_Comment) Reset() {
	*x = Data_Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_tricium_api_v1_data_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Comment) ProtoMessage() {}

func (x *Data_Comment) ProtoReflect() protoreflect.Message {
	mi := &file_infra_tricium_api_v1_data_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Comment.ProtoReflect.Descriptor instead.
func (*Data_Comment) Descriptor() ([]byte, []int) {
	return file_infra_tricium_api_v1_data_proto_rawDescGZIP(), []int{0, 5}
}

func (x *Data_Comment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Data_Comment) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Data_Comment) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Data_Comment) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Data_Comment) GetStartLine() int32 {
	if x != nil {
		return x.StartLine
	}
	return 0
}

func (x *Data_Comment) GetEndLine() int32 {
	if x != nil {
		return x.EndLine
	}
	return 0
}

func (x *Data_Comment) GetStartChar() int32 {
	if x != nil {
		return x.StartChar
	}
	return 0
}

func (x *Data_Comment) GetEndChar() int32 {
	if x != nil {
		return x.EndChar
	}
	return 0
}

func (x *Data_Comment) GetSuggestions() []*Data_Suggestion {
	if x != nil {
		return x.Suggestions
	}
	return nil
}

func (x *Data_Comment) GetShowOnUnchangedLines() bool {
	if x != nil {
		return x.ShowOnUnchangedLines
	}
	return false
}

// Suggested fix.
//
// A fix may include replacements in any file in the same repo as the file of
// the corresponding comment.
type Data_Suggestion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A brief description of the suggested fix.
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// Fix as a list of replacements.
	Replacements []*Data_Replacement `protobuf:"bytes,2,rep,name=replacements,proto3" json:"replacements,omitempty"`
}

func (x *Data_Suggestion) Reset() {
	*x = Data_Suggestion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_tricium_api_v1_data_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Suggestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Suggestion) ProtoMessage() {}

func (x *Data_Suggestion) ProtoReflect() protoreflect.Message {
	mi := &file_infra_tricium_api_v1_data_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Suggestion.ProtoReflect.Descriptor instead.
func (*Data_Suggestion) Descriptor() ([]byte, []int) {
	return file_infra_tricium_api_v1_data_proto_rawDescGZIP(), []int{0, 6}
}

func (x *Data_Suggestion) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Data_Suggestion) GetReplacements() []*Data_Replacement {
	if x != nil {
		return x.Replacements
	}
	return nil
}

// A suggested replacement.
//
// The replacement should be for one continuous section of a file.
type Data_Replacement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path to the file for this replacement.
	//
	// An empty string indicates the commit message.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// A replacement string.
	Replacement string `protobuf:"bytes,2,opt,name=replacement,proto3" json:"replacement,omitempty"`
	// A continuous section of the file to replace.
	StartLine int32 `protobuf:"varint,3,opt,name=start_line,json=startLine,proto3" json:"start_line,omitempty"` // 1-based, inclusive.
	EndLine   int32 `protobuf:"varint,4,opt,name=end_line,json=endLine,proto3" json:"end_line,omitempty"`       // 1-based, inclusive.
	StartChar int32 `protobuf:"varint,5,opt,name=start_char,json=startChar,proto3" json:"start_char,omitempty"` // 0-based, inclusive.
	EndChar   int32 `protobuf:"varint,6,opt,name=end_char,json=endChar,proto3" json:"end_char,omitempty"`       // 0-based, exclusive.
}

func (x *Data_Replacement) Reset() {
	*x = Data_Replacement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_tricium_api_v1_data_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Replacement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Replacement) ProtoMessage() {}

func (x *Data_Replacement) ProtoReflect() protoreflect.Message {
	mi := &file_infra_tricium_api_v1_data_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Replacement.ProtoReflect.Descriptor instead.
func (*Data_Replacement) Descriptor() ([]byte, []int) {
	return file_infra_tricium_api_v1_data_proto_rawDescGZIP(), []int{0, 7}
}

func (x *Data_Replacement) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Data_Replacement) GetReplacement() string {
	if x != nil {
		return x.Replacement
	}
	return ""
}

func (x *Data_Replacement) GetStartLine() int32 {
	if x != nil {
		return x.StartLine
	}
	return 0
}

func (x *Data_Replacement) GetEndLine() int32 {
	if x != nil {
		return x.EndLine
	}
	return 0
}

func (x *Data_Replacement) GetStartChar() int32 {
	if x != nil {
		return x.StartChar
	}
	return 0
}

func (x *Data_Replacement) GetEndChar() int32 {
	if x != nil {
		return x.EndChar
	}
	return 0
}

var File_infra_tricium_api_v1_data_proto protoreflect.FileDescriptor

var file_infra_tricium_api_v1_data_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x22, 0xf8, 0x0a, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x67, 0x0a, 0x0b, 0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x12, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x73,
	0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x69, 0x73, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x1a, 0xb1, 0x01, 0x0a,
	0x0e, 0x47, 0x69, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x72, 0x65, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12,
	0x28, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x76, 0x0a, 0x05, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x12, 0x28, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x65, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x62, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x42, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a,
	0x5a, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x12, 0x31, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x72, 0x69,
	0x63, 0x69, 0x75, 0x6d, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0xd0, 0x02, 0x0a, 0x07,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x68, 0x61, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x6e,
	0x64, 0x43, 0x68, 0x61, 0x72, 0x12, 0x3a, 0x0a, 0x0b, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x72, 0x69,
	0x63, 0x69, 0x75, 0x6d, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x35, 0x0a, 0x17, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x6f, 0x6e, 0x5f, 0x75, 0x6e, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x14, 0x73, 0x68, 0x6f, 0x77, 0x4f, 0x6e, 0x55, 0x6e, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x1a, 0x6d,
	0x0a, 0x0a, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d,
	0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x0c, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0xb7, 0x01,
	0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6c, 0x69, 0x6e,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4c, 0x69,
	0x6e, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x68, 0x61, 0x72, 0x12, 0x19, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x72, 0x22, 0x44, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x49, 0x54,
	0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x44, 0x45, 0x54, 0x41, 0x49, 0x4c, 0x53, 0x10, 0x01, 0x12,
	0x09, 0x0a, 0x05, 0x46, 0x49, 0x4c, 0x45, 0x53, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45,
	0x53, 0x55, 0x4c, 0x54, 0x53, 0x10, 0x04, 0x22, 0x04, 0x08, 0x03, 0x10, 0x03, 0x22, 0x56, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x4f, 0x44, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x44, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a,
	0x07, 0x52, 0x45, 0x4e, 0x41, 0x4d, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f,
	0x50, 0x49, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x57, 0x52, 0x49, 0x54,
	0x54, 0x45, 0x4e, 0x10, 0x05, 0x42, 0x1e, 0x5a, 0x1c, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x74,
	0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x72,
	0x69, 0x63, 0x69, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_tricium_api_v1_data_proto_rawDescOnce sync.Once
	file_infra_tricium_api_v1_data_proto_rawDescData = file_infra_tricium_api_v1_data_proto_rawDesc
)

func file_infra_tricium_api_v1_data_proto_rawDescGZIP() []byte {
	file_infra_tricium_api_v1_data_proto_rawDescOnce.Do(func() {
		file_infra_tricium_api_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_tricium_api_v1_data_proto_rawDescData)
	})
	return file_infra_tricium_api_v1_data_proto_rawDescData
}

var file_infra_tricium_api_v1_data_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_infra_tricium_api_v1_data_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_infra_tricium_api_v1_data_proto_goTypes = []interface{}{
	(Data_Type)(0),              // 0: tricium.Data.Type
	(Data_Status)(0),            // 1: tricium.Data.Status
	(*Data)(nil),                // 2: tricium.Data
	(*Data_TypeDetails)(nil),    // 3: tricium.Data.TypeDetails
	(*Data_GitFileDetails)(nil), // 4: tricium.Data.GitFileDetails
	(*Data_Files)(nil),          // 5: tricium.Data.Files
	(*Data_File)(nil),           // 6: tricium.Data.File
	(*Data_Results)(nil),        // 7: tricium.Data.Results
	(*Data_Comment)(nil),        // 8: tricium.Data.Comment
	(*Data_Suggestion)(nil),     // 9: tricium.Data.Suggestion
	(*Data_Replacement)(nil),    // 10: tricium.Data.Replacement
}
var file_infra_tricium_api_v1_data_proto_depIdxs = []int32{
	0,  // 0: tricium.Data.TypeDetails.type:type_name -> tricium.Data.Type
	6,  // 1: tricium.Data.GitFileDetails.files:type_name -> tricium.Data.File
	6,  // 2: tricium.Data.Files.files:type_name -> tricium.Data.File
	1,  // 3: tricium.Data.File.status:type_name -> tricium.Data.Status
	8,  // 4: tricium.Data.Results.comments:type_name -> tricium.Data.Comment
	9,  // 5: tricium.Data.Comment.suggestions:type_name -> tricium.Data.Suggestion
	10, // 6: tricium.Data.Suggestion.replacements:type_name -> tricium.Data.Replacement
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_infra_tricium_api_v1_data_proto_init() }
func file_infra_tricium_api_v1_data_proto_init() {
	if File_infra_tricium_api_v1_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_tricium_api_v1_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_infra_tricium_api_v1_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_TypeDetails); i {
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
		file_infra_tricium_api_v1_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_GitFileDetails); i {
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
		file_infra_tricium_api_v1_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Files); i {
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
		file_infra_tricium_api_v1_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_File); i {
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
		file_infra_tricium_api_v1_data_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Results); i {
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
		file_infra_tricium_api_v1_data_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Comment); i {
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
		file_infra_tricium_api_v1_data_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Suggestion); i {
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
		file_infra_tricium_api_v1_data_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Replacement); i {
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
			RawDescriptor: file_infra_tricium_api_v1_data_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_tricium_api_v1_data_proto_goTypes,
		DependencyIndexes: file_infra_tricium_api_v1_data_proto_depIdxs,
		EnumInfos:         file_infra_tricium_api_v1_data_proto_enumTypes,
		MessageInfos:      file_infra_tricium_api_v1_data_proto_msgTypes,
	}.Build()
	File_infra_tricium_api_v1_data_proto = out.File
	file_infra_tricium_api_v1_data_proto_rawDesc = nil
	file_infra_tricium_api_v1_data_proto_goTypes = nil
	file_infra_tricium_api_v1_data_proto_depIdxs = nil
}
