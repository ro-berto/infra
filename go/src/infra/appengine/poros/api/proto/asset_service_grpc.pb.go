// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: poros/api/proto/asset_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AssetClient is the client API for Asset service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssetClient interface {
	// Creates the given Asset.
	Create(ctx context.Context, in *CreateAssetRequest, opts ...grpc.CallOption) (*CreateAssetResponse, error)
	// Retrieves a Asset for a given unique value.
	Get(ctx context.Context, in *GetAssetRequest, opts ...grpc.CallOption) (*AssetModel, error)
	// Update a single asset in EnterpriseAsset.
	Update(ctx context.Context, in *UpdateAssetRequest, opts ...grpc.CallOption) (*UpdateAssetResponse, error)
	// Deletes the given Asset.
	Delete(ctx context.Context, in *DeleteAssetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Lists all Assets.
	List(ctx context.Context, in *ListAssetsRequest, opts ...grpc.CallOption) (*ListAssetsResponse, error)
	// Get AssetConfiguration
	GetAssetConfiguration(ctx context.Context, in *GetAssetConfigurationRequest, opts ...grpc.CallOption) (*GetAssetConfigurationResponse, error)
}

type assetClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetClient(cc grpc.ClientConnInterface) AssetClient {
	return &assetClient{cc}
}

func (c *assetClient) Create(ctx context.Context, in *CreateAssetRequest, opts ...grpc.CallOption) (*CreateAssetResponse, error) {
	out := new(CreateAssetResponse)
	err := c.cc.Invoke(ctx, "/poros.Asset/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) Get(ctx context.Context, in *GetAssetRequest, opts ...grpc.CallOption) (*AssetModel, error) {
	out := new(AssetModel)
	err := c.cc.Invoke(ctx, "/poros.Asset/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) Update(ctx context.Context, in *UpdateAssetRequest, opts ...grpc.CallOption) (*UpdateAssetResponse, error) {
	out := new(UpdateAssetResponse)
	err := c.cc.Invoke(ctx, "/poros.Asset/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) Delete(ctx context.Context, in *DeleteAssetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/poros.Asset/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) List(ctx context.Context, in *ListAssetsRequest, opts ...grpc.CallOption) (*ListAssetsResponse, error) {
	out := new(ListAssetsResponse)
	err := c.cc.Invoke(ctx, "/poros.Asset/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) GetAssetConfiguration(ctx context.Context, in *GetAssetConfigurationRequest, opts ...grpc.CallOption) (*GetAssetConfigurationResponse, error) {
	out := new(GetAssetConfigurationResponse)
	err := c.cc.Invoke(ctx, "/poros.Asset/GetAssetConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetServer is the server API for Asset service.
// All implementations must embed UnimplementedAssetServer
// for forward compatibility
type AssetServer interface {
	// Creates the given Asset.
	Create(context.Context, *CreateAssetRequest) (*CreateAssetResponse, error)
	// Retrieves a Asset for a given unique value.
	Get(context.Context, *GetAssetRequest) (*AssetModel, error)
	// Update a single asset in EnterpriseAsset.
	Update(context.Context, *UpdateAssetRequest) (*UpdateAssetResponse, error)
	// Deletes the given Asset.
	Delete(context.Context, *DeleteAssetRequest) (*emptypb.Empty, error)
	// Lists all Assets.
	List(context.Context, *ListAssetsRequest) (*ListAssetsResponse, error)
	// Get AssetConfiguration
	GetAssetConfiguration(context.Context, *GetAssetConfigurationRequest) (*GetAssetConfigurationResponse, error)
	mustEmbedUnimplementedAssetServer()
}

// UnimplementedAssetServer must be embedded to have forward compatible implementations.
type UnimplementedAssetServer struct {
}

func (UnimplementedAssetServer) Create(context.Context, *CreateAssetRequest) (*CreateAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAssetServer) Get(context.Context, *GetAssetRequest) (*AssetModel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAssetServer) Update(context.Context, *UpdateAssetRequest) (*UpdateAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAssetServer) Delete(context.Context, *DeleteAssetRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAssetServer) List(context.Context, *ListAssetsRequest) (*ListAssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAssetServer) GetAssetConfiguration(context.Context, *GetAssetConfigurationRequest) (*GetAssetConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssetConfiguration not implemented")
}
func (UnimplementedAssetServer) mustEmbedUnimplementedAssetServer() {}

// UnsafeAssetServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetServer will
// result in compilation errors.
type UnsafeAssetServer interface {
	mustEmbedUnimplementedAssetServer()
}

func RegisterAssetServer(s grpc.ServiceRegistrar, srv AssetServer) {
	s.RegisterService(&Asset_ServiceDesc, srv)
}

func _Asset_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poros.Asset/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServer).Create(ctx, req.(*CreateAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Asset_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poros.Asset/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServer).Get(ctx, req.(*GetAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Asset_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poros.Asset/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServer).Update(ctx, req.(*UpdateAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Asset_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poros.Asset/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServer).Delete(ctx, req.(*DeleteAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Asset_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poros.Asset/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServer).List(ctx, req.(*ListAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Asset_GetAssetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssetConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServer).GetAssetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poros.Asset/GetAssetConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServer).GetAssetConfiguration(ctx, req.(*GetAssetConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Asset_ServiceDesc is the grpc.ServiceDesc for Asset service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Asset_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "poros.Asset",
	HandlerType: (*AssetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Asset_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Asset_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Asset_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Asset_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Asset_List_Handler,
		},
		{
			MethodName: "GetAssetConfiguration",
			Handler:    _Asset_GetAssetConfiguration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "poros/api/proto/asset_service.proto",
}
