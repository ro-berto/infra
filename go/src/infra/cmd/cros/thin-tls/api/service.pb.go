// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DutShellRequest struct {
	Command              string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DutShellRequest) Reset()         { *m = DutShellRequest{} }
func (m *DutShellRequest) String() string { return proto.CompactTextString(m) }
func (*DutShellRequest) ProtoMessage()    {}
func (*DutShellRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *DutShellRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DutShellRequest.Unmarshal(m, b)
}
func (m *DutShellRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DutShellRequest.Marshal(b, m, deterministic)
}
func (m *DutShellRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DutShellRequest.Merge(m, src)
}
func (m *DutShellRequest) XXX_Size() int {
	return xxx_messageInfo_DutShellRequest.Size(m)
}
func (m *DutShellRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DutShellRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DutShellRequest proto.InternalMessageInfo

func (m *DutShellRequest) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

// For the last response in the stream, exited will be true and status
// will be set.
type DutShellResponse struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Exited               bool     `protobuf:"varint,2,opt,name=exited,proto3" json:"exited,omitempty"`
	Output               []byte   `protobuf:"bytes,3,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DutShellResponse) Reset()         { *m = DutShellResponse{} }
func (m *DutShellResponse) String() string { return proto.CompactTextString(m) }
func (*DutShellResponse) ProtoMessage()    {}
func (*DutShellResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *DutShellResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DutShellResponse.Unmarshal(m, b)
}
func (m *DutShellResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DutShellResponse.Marshal(b, m, deterministic)
}
func (m *DutShellResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DutShellResponse.Merge(m, src)
}
func (m *DutShellResponse) XXX_Size() int {
	return xxx_messageInfo_DutShellResponse.Size(m)
}
func (m *DutShellResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DutShellResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DutShellResponse proto.InternalMessageInfo

func (m *DutShellResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *DutShellResponse) GetExited() bool {
	if m != nil {
		return m.Exited
	}
	return false
}

func (m *DutShellResponse) GetOutput() []byte {
	if m != nil {
		return m.Output
	}
	return nil
}

func init() {
	proto.RegisterType((*DutShellRequest)(nil), "cros.tls.experimental.DutShellRequest")
	proto.RegisterType((*DutShellResponse)(nil), "cros.tls.experimental.DutShellResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x89, 0xa5, 0xb5, 0x06, 0x45, 0x09, 0x28, 0xc1, 0x53, 0xe9, 0x41, 0x0b, 0x42, 0x10,
	0x7d, 0x03, 0xf1, 0x09, 0xa2, 0xa7, 0x82, 0x87, 0xd8, 0x0e, 0x18, 0x48, 0x9b, 0x98, 0x99, 0x48,
	0x1f, 0x5f, 0xda, 0x6d, 0x59, 0x58, 0x16, 0xf6, 0xf8, 0xfd, 0x7c, 0x30, 0x1f, 0xc3, 0xaf, 0x10,
	0xe2, 0x9f, 0xed, 0x40, 0x85, 0xe8, 0xc9, 0x8b, 0xdb, 0x2e, 0x7a, 0x54, 0xe4, 0x50, 0xc1, 0x14,
	0x20, 0xda, 0x01, 0x46, 0x32, 0xae, 0x7e, 0xe2, 0xd7, 0xef, 0x89, 0x3e, 0x7e, 0xc0, 0x39, 0x0d,
	0xbf, 0x09, 0x90, 0x84, 0xe4, 0xe7, 0x9d, 0x1f, 0x06, 0x33, 0xf6, 0x92, 0x55, 0xac, 0xb9, 0xd0,
	0x1b, 0xd6, 0x2d, 0xbf, 0xd9, 0xcb, 0x18, 0xfc, 0x88, 0x20, 0xee, 0x78, 0x81, 0x64, 0x28, 0xe1,
	0x22, 0xe7, 0x7a, 0xa5, 0x79, 0x87, 0xc9, 0x12, 0xf4, 0xf2, 0xac, 0x62, 0x4d, 0xa9, 0x57, 0x9a,
	0x77, 0x9f, 0x28, 0x24, 0x92, 0x59, 0xc5, 0x9a, 0x4b, 0xbd, 0xd2, 0x4b, 0xcf, 0xb3, 0x4f, 0x87,
	0xe2, 0x8b, 0x97, 0xdb, 0x09, 0xf1, 0xa0, 0x8e, 0x36, 0xab, 0x83, 0xe0, 0xfb, 0xc7, 0x93, 0xde,
	0xae, 0xf5, 0x99, 0xbd, 0xe5, 0x6d, 0x66, 0x82, 0xfd, 0x2e, 0x96, 0x9f, 0xbc, 0xfe, 0x07, 0x00,
	0x00, 0xff, 0xff, 0xba, 0x65, 0x42, 0x6e, 0x24, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TlsClient is the client API for Tls service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TlsClient interface {
	// Runs a shell command with the default shell.
	// Does not spawn a tty.
	DutShell(ctx context.Context, in *DutShellRequest, opts ...grpc.CallOption) (Tls_DutShellClient, error)
}

type tlsClient struct {
	cc *grpc.ClientConn
}

func NewTlsClient(cc *grpc.ClientConn) TlsClient {
	return &tlsClient{cc}
}

func (c *tlsClient) DutShell(ctx context.Context, in *DutShellRequest, opts ...grpc.CallOption) (Tls_DutShellClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Tls_serviceDesc.Streams[0], "/cros.tls.experimental.Tls/DutShell", opts...)
	if err != nil {
		return nil, err
	}
	x := &tlsDutShellClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Tls_DutShellClient interface {
	Recv() (*DutShellResponse, error)
	grpc.ClientStream
}

type tlsDutShellClient struct {
	grpc.ClientStream
}

func (x *tlsDutShellClient) Recv() (*DutShellResponse, error) {
	m := new(DutShellResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TlsServer is the server API for Tls service.
type TlsServer interface {
	// Runs a shell command with the default shell.
	// Does not spawn a tty.
	DutShell(*DutShellRequest, Tls_DutShellServer) error
}

// UnimplementedTlsServer can be embedded to have forward compatible implementations.
type UnimplementedTlsServer struct {
}

func (*UnimplementedTlsServer) DutShell(req *DutShellRequest, srv Tls_DutShellServer) error {
	return status.Errorf(codes.Unimplemented, "method DutShell not implemented")
}

func RegisterTlsServer(s *grpc.Server, srv TlsServer) {
	s.RegisterService(&_Tls_serviceDesc, srv)
}

func _Tls_DutShell_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DutShellRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TlsServer).DutShell(m, &tlsDutShellServer{stream})
}

type Tls_DutShellServer interface {
	Send(*DutShellResponse) error
	grpc.ServerStream
}

type tlsDutShellServer struct {
	grpc.ServerStream
}

func (x *tlsDutShellServer) Send(m *DutShellResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Tls_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cros.tls.experimental.Tls",
	HandlerType: (*TlsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DutShell",
			Handler:       _Tls_DutShell_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}
