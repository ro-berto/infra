// Copyright 2022 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: poros/api/proto/asset_service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request to create a single asset in AssetServ
type CreateAssetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The new asset to add to the database.
	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateAssetRequest) Reset() {
	*x = CreateAssetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poros_api_proto_asset_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAssetRequest) ProtoMessage() {}

func (x *CreateAssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_poros_api_proto_asset_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAssetRequest.ProtoReflect.Descriptor instead.
func (*CreateAssetRequest) Descriptor() ([]byte, []int) {
	return file_poros_api_proto_asset_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAssetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateAssetRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// Gets a Asset resource.
type GetAssetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of the asset to retrieve.
	// Format: publishers/{publisher}/assets/{asset}
	AssetId string `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
}

func (x *GetAssetRequest) Reset() {
	*x = GetAssetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poros_api_proto_asset_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetRequest) ProtoMessage() {}

func (x *GetAssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_poros_api_proto_asset_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetRequest.ProtoReflect.Descriptor instead.
func (*GetAssetRequest) Descriptor() ([]byte, []int) {
	return file_poros_api_proto_asset_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetAssetRequest) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

// Request to update a single asset in EnterpriseAsset.
type UpdateAssetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The existing asset to update in the database.
	Asset *AssetModel `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	// The list of fields to update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateAssetRequest) Reset() {
	*x = UpdateAssetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poros_api_proto_asset_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAssetRequest) ProtoMessage() {}

func (x *UpdateAssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_poros_api_proto_asset_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAssetRequest.ProtoReflect.Descriptor instead.
func (*UpdateAssetRequest) Descriptor() ([]byte, []int) {
	return file_poros_api_proto_asset_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAssetRequest) GetAsset() *AssetModel {
	if x != nil {
		return x.Asset
	}
	return nil
}

func (x *UpdateAssetRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// Request to delete a single asset from EnterpriseAsset.
type DeleteAssetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of the asset to delete from the database.
	AssetId string `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
}

func (x *DeleteAssetRequest) Reset() {
	*x = DeleteAssetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poros_api_proto_asset_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAssetRequest) ProtoMessage() {}

func (x *DeleteAssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_poros_api_proto_asset_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAssetRequest.ProtoReflect.Descriptor instead.
func (*DeleteAssetRequest) Descriptor() ([]byte, []int) {
	return file_poros_api_proto_asset_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteAssetRequest) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

// Request to list all assets in EnterpriseAsset.
type ListAssetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Fields to include on each result
	ReadMask *fieldmaskpb.FieldMask `protobuf:"bytes,1,opt,name=read_mask,json=readMask,proto3" json:"read_mask,omitempty"`
	// Number of results per page
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token from a previous page's ListAssetsResponse
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListAssetsRequest) Reset() {
	*x = ListAssetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poros_api_proto_asset_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAssetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAssetsRequest) ProtoMessage() {}

func (x *ListAssetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_poros_api_proto_asset_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAssetsRequest.ProtoReflect.Descriptor instead.
func (*ListAssetsRequest) Descriptor() ([]byte, []int) {
	return file_poros_api_proto_asset_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListAssetsRequest) GetReadMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.ReadMask
	}
	return nil
}

func (x *ListAssetsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListAssetsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response to ListAssetsRequest.
type ListAssetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The result set.
	Assets []*AssetModel `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets,omitempty"`
	// A page token that can be passed into future requests to get the next page.
	// Empty if there is no next page.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListAssetsResponse) Reset() {
	*x = ListAssetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poros_api_proto_asset_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAssetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAssetsResponse) ProtoMessage() {}

func (x *ListAssetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_poros_api_proto_asset_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAssetsResponse.ProtoReflect.Descriptor instead.
func (*ListAssetsResponse) Descriptor() ([]byte, []int) {
	return file_poros_api_proto_asset_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListAssetsResponse) GetAssets() []*AssetModel {
	if x != nil {
		return x.Assets
	}
	return nil
}

func (x *ListAssetsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// request for get the asset configuration by asset id
type GetAssetConfigurationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of the asset to retrieve.
	AssetId string `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
}

func (x *GetAssetConfigurationRequest) Reset() {
	*x = GetAssetConfigurationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poros_api_proto_asset_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssetConfigurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetConfigurationRequest) ProtoMessage() {}

func (x *GetAssetConfigurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_poros_api_proto_asset_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetConfigurationRequest.ProtoReflect.Descriptor instead.
func (*GetAssetConfigurationRequest) Descriptor() ([]byte, []int) {
	return file_poros_api_proto_asset_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetAssetConfigurationRequest) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

type GetAssetConfigurationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The configuration for the asset.
	Config string `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *GetAssetConfigurationResponse) Reset() {
	*x = GetAssetConfigurationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poros_api_proto_asset_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssetConfigurationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetConfigurationResponse) ProtoMessage() {}

func (x *GetAssetConfigurationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_poros_api_proto_asset_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetConfigurationResponse.ProtoReflect.Descriptor instead.
func (*GetAssetConfigurationResponse) Descriptor() ([]byte, []int) {
	return file_poros_api_proto_asset_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetAssetConfigurationResponse) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

var File_poros_api_proto_asset_service_proto protoreflect.FileDescriptor

var file_poros_api_proto_asset_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x6f, 0x72, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x6f, 0x72, 0x6f, 0x73, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x6f, 0x72,
	0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x22, 0x7a, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x05,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6f,
	0x72, 0x6f, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x05,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61,
	0x73, 0x6b, 0x22, 0x2f, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x49, 0x64, 0x22, 0x88, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x65, 0x61,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x4d, 0x61,
	0x73, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x67,
	0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6f, 0x72, 0x6f, 0x73, 0x2e, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x39, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x49, 0x64, 0x22, 0x37, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x32, 0x87, 0x03, 0x0a, 0x05,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x19, 0x2e, 0x70, 0x6f, 0x72, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x6f, 0x72,
	0x6f, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x30, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x6f, 0x72, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70,
	0x6f, 0x72, 0x6f, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x36, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x6f, 0x72, 0x6f,
	0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x6f, 0x72, 0x6f, 0x73, 0x2e, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x3b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x19, 0x2e, 0x70, 0x6f, 0x72, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x70,
	0x6f, 0x72, 0x6f, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x6f, 0x72, 0x6f, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x62, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x70, 0x6f, 0x72,
	0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x70, 0x6f, 0x72, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61,
	0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x6f, 0x72, 0x6f, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_poros_api_proto_asset_service_proto_rawDescOnce sync.Once
	file_poros_api_proto_asset_service_proto_rawDescData = file_poros_api_proto_asset_service_proto_rawDesc
)

func file_poros_api_proto_asset_service_proto_rawDescGZIP() []byte {
	file_poros_api_proto_asset_service_proto_rawDescOnce.Do(func() {
		file_poros_api_proto_asset_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_poros_api_proto_asset_service_proto_rawDescData)
	})
	return file_poros_api_proto_asset_service_proto_rawDescData
}

var file_poros_api_proto_asset_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_poros_api_proto_asset_service_proto_goTypes = []interface{}{
	(*CreateAssetRequest)(nil),            // 0: poros.CreateAssetRequest
	(*GetAssetRequest)(nil),               // 1: poros.GetAssetRequest
	(*UpdateAssetRequest)(nil),            // 2: poros.UpdateAssetRequest
	(*DeleteAssetRequest)(nil),            // 3: poros.DeleteAssetRequest
	(*ListAssetsRequest)(nil),             // 4: poros.ListAssetsRequest
	(*ListAssetsResponse)(nil),            // 5: poros.ListAssetsResponse
	(*GetAssetConfigurationRequest)(nil),  // 6: poros.GetAssetConfigurationRequest
	(*GetAssetConfigurationResponse)(nil), // 7: poros.GetAssetConfigurationResponse
	(*AssetModel)(nil),                    // 8: poros.AssetModel
	(*fieldmaskpb.FieldMask)(nil),         // 9: google.protobuf.FieldMask
	(*emptypb.Empty)(nil),                 // 10: google.protobuf.Empty
}
var file_poros_api_proto_asset_service_proto_depIdxs = []int32{
	8,  // 0: poros.UpdateAssetRequest.asset:type_name -> poros.AssetModel
	9,  // 1: poros.UpdateAssetRequest.update_mask:type_name -> google.protobuf.FieldMask
	9,  // 2: poros.ListAssetsRequest.read_mask:type_name -> google.protobuf.FieldMask
	8,  // 3: poros.ListAssetsResponse.assets:type_name -> poros.AssetModel
	0,  // 4: poros.Asset.Create:input_type -> poros.CreateAssetRequest
	1,  // 5: poros.Asset.Get:input_type -> poros.GetAssetRequest
	2,  // 6: poros.Asset.Update:input_type -> poros.UpdateAssetRequest
	3,  // 7: poros.Asset.Delete:input_type -> poros.DeleteAssetRequest
	4,  // 8: poros.Asset.List:input_type -> poros.ListAssetsRequest
	6,  // 9: poros.Asset.GetAssetConfiguration:input_type -> poros.GetAssetConfigurationRequest
	8,  // 10: poros.Asset.Create:output_type -> poros.AssetModel
	8,  // 11: poros.Asset.Get:output_type -> poros.AssetModel
	8,  // 12: poros.Asset.Update:output_type -> poros.AssetModel
	10, // 13: poros.Asset.Delete:output_type -> google.protobuf.Empty
	5,  // 14: poros.Asset.List:output_type -> poros.ListAssetsResponse
	7,  // 15: poros.Asset.GetAssetConfiguration:output_type -> poros.GetAssetConfigurationResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_poros_api_proto_asset_service_proto_init() }
func file_poros_api_proto_asset_service_proto_init() {
	if File_poros_api_proto_asset_service_proto != nil {
		return
	}
	file_poros_api_proto_resources_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_poros_api_proto_asset_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAssetRequest); i {
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
		file_poros_api_proto_asset_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAssetRequest); i {
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
		file_poros_api_proto_asset_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAssetRequest); i {
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
		file_poros_api_proto_asset_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAssetRequest); i {
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
		file_poros_api_proto_asset_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAssetsRequest); i {
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
		file_poros_api_proto_asset_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAssetsResponse); i {
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
		file_poros_api_proto_asset_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAssetConfigurationRequest); i {
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
		file_poros_api_proto_asset_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAssetConfigurationResponse); i {
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
			RawDescriptor: file_poros_api_proto_asset_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_poros_api_proto_asset_service_proto_goTypes,
		DependencyIndexes: file_poros_api_proto_asset_service_proto_depIdxs,
		MessageInfos:      file_poros_api_proto_asset_service_proto_msgTypes,
	}.Build()
	File_poros_api_proto_asset_service_proto = out.File
	file_poros_api_proto_asset_service_proto_rawDesc = nil
	file_poros_api_proto_asset_service_proto_goTypes = nil
	file_poros_api_proto_asset_service_proto_depIdxs = nil
}