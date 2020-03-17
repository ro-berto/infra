// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/v1/api_proto/hotlists.proto

package monorail_v1

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

// Request message for GetHotlist method.
// Next available tag: 2
type GetHotlistRequest struct {
	// The name of the hotlist to retrieve.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHotlistRequest) Reset()         { *m = GetHotlistRequest{} }
func (m *GetHotlistRequest) String() string { return proto.CompactTextString(m) }
func (*GetHotlistRequest) ProtoMessage()    {}
func (*GetHotlistRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a2c4cb0040a55a, []int{0}
}

func (m *GetHotlistRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHotlistRequest.Unmarshal(m, b)
}
func (m *GetHotlistRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHotlistRequest.Marshal(b, m, deterministic)
}
func (m *GetHotlistRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHotlistRequest.Merge(m, src)
}
func (m *GetHotlistRequest) XXX_Size() int {
	return xxx_messageInfo_GetHotlistRequest.Size(m)
}
func (m *GetHotlistRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHotlistRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHotlistRequest proto.InternalMessageInfo

func (m *GetHotlistRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for UpdateHotlist method.
// Next available tag: 2
type UpdateHotlistRequest struct {
	// The hotlist's `name` field is used to identify the hotlist to be updated.
	Hotlist *Hotlist `protobuf:"bytes,1,opt,name=hotlist,proto3" json:"hotlist,omitempty"`
	// The list of fields to be updated.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateHotlistRequest) Reset()         { *m = UpdateHotlistRequest{} }
func (m *UpdateHotlistRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateHotlistRequest) ProtoMessage()    {}
func (*UpdateHotlistRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a2c4cb0040a55a, []int{1}
}

func (m *UpdateHotlistRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateHotlistRequest.Unmarshal(m, b)
}
func (m *UpdateHotlistRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateHotlistRequest.Marshal(b, m, deterministic)
}
func (m *UpdateHotlistRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateHotlistRequest.Merge(m, src)
}
func (m *UpdateHotlistRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateHotlistRequest.Size(m)
}
func (m *UpdateHotlistRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateHotlistRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateHotlistRequest proto.InternalMessageInfo

func (m *UpdateHotlistRequest) GetHotlist() *Hotlist {
	if m != nil {
		return m.Hotlist
	}
	return nil
}

func (m *UpdateHotlistRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// Request message for ListHotlistItems method.
// Next available tag: 5
type ListHotlistItemsRequest struct {
	// The parent hotlist, which owns this collection of items.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of items to return. The service may return fewer than
	// this value.
	// If unspecified, at most 1000 items will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The string of comma separated field names used to order the items.
	// Adding '-' before a field, reverses the sort order.
	// E.g. 'stars,-status' sorts the items by number of stars low to high, then
	// status high to low.
	// If unspecified, items will be ordered by their rank in the parent.
	OrderBy string `protobuf:"bytes,3,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// A page token, received from a previous `ListHotlistItems` call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to `ListHotlistItems` must
	// match the call that provided the page token.
	PageToken            string   `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListHotlistItemsRequest) Reset()         { *m = ListHotlistItemsRequest{} }
func (m *ListHotlistItemsRequest) String() string { return proto.CompactTextString(m) }
func (*ListHotlistItemsRequest) ProtoMessage()    {}
func (*ListHotlistItemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a2c4cb0040a55a, []int{2}
}

func (m *ListHotlistItemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListHotlistItemsRequest.Unmarshal(m, b)
}
func (m *ListHotlistItemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListHotlistItemsRequest.Marshal(b, m, deterministic)
}
func (m *ListHotlistItemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListHotlistItemsRequest.Merge(m, src)
}
func (m *ListHotlistItemsRequest) XXX_Size() int {
	return xxx_messageInfo_ListHotlistItemsRequest.Size(m)
}
func (m *ListHotlistItemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListHotlistItemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListHotlistItemsRequest proto.InternalMessageInfo

func (m *ListHotlistItemsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListHotlistItemsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListHotlistItemsRequest) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

func (m *ListHotlistItemsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Response to ListHotlistItems call.
// Next available tag: 3
type ListHotlistItemsResponse struct {
	// The items from the specified hotlist.
	Items []*HotlistItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListHotlistItemsResponse) Reset()         { *m = ListHotlistItemsResponse{} }
func (m *ListHotlistItemsResponse) String() string { return proto.CompactTextString(m) }
func (*ListHotlistItemsResponse) ProtoMessage()    {}
func (*ListHotlistItemsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a2c4cb0040a55a, []int{3}
}

func (m *ListHotlistItemsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListHotlistItemsResponse.Unmarshal(m, b)
}
func (m *ListHotlistItemsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListHotlistItemsResponse.Marshal(b, m, deterministic)
}
func (m *ListHotlistItemsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListHotlistItemsResponse.Merge(m, src)
}
func (m *ListHotlistItemsResponse) XXX_Size() int {
	return xxx_messageInfo_ListHotlistItemsResponse.Size(m)
}
func (m *ListHotlistItemsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListHotlistItemsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListHotlistItemsResponse proto.InternalMessageInfo

func (m *ListHotlistItemsResponse) GetItems() []*HotlistItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ListHotlistItemsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The request used to rerank a Hotlist.
// Next available tag: 4
type RerankHotlistItemsRequest struct {
	// Resource name of the Hotlist to rerank.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// HotlistItems to be moved. The order of `hotlist_items` will
	// determine the order of these items after they have been moved.
	// E.g. With items [a, b, c, d, e], moving [d, c] to `target_position` 3, will
	// result in items [a, b, e, d, c].
	HotlistItems []string `protobuf:"bytes,2,rep,name=hotlist_items,json=hotlistItems,proto3" json:"hotlist_items,omitempty"`
	// Target starting position of the moved items.
	// `target_position` must be between 0 and (# hotlist items - # items being moved).
	TargetPosition       uint32   `protobuf:"varint,3,opt,name=target_position,json=targetPosition,proto3" json:"target_position,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RerankHotlistItemsRequest) Reset()         { *m = RerankHotlistItemsRequest{} }
func (m *RerankHotlistItemsRequest) String() string { return proto.CompactTextString(m) }
func (*RerankHotlistItemsRequest) ProtoMessage()    {}
func (*RerankHotlistItemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a2c4cb0040a55a, []int{4}
}

func (m *RerankHotlistItemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RerankHotlistItemsRequest.Unmarshal(m, b)
}
func (m *RerankHotlistItemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RerankHotlistItemsRequest.Marshal(b, m, deterministic)
}
func (m *RerankHotlistItemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RerankHotlistItemsRequest.Merge(m, src)
}
func (m *RerankHotlistItemsRequest) XXX_Size() int {
	return xxx_messageInfo_RerankHotlistItemsRequest.Size(m)
}
func (m *RerankHotlistItemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RerankHotlistItemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RerankHotlistItemsRequest proto.InternalMessageInfo

func (m *RerankHotlistItemsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RerankHotlistItemsRequest) GetHotlistItems() []string {
	if m != nil {
		return m.HotlistItems
	}
	return nil
}

func (m *RerankHotlistItemsRequest) GetTargetPosition() uint32 {
	if m != nil {
		return m.TargetPosition
	}
	return 0
}

// Request message for an AddHotlistItems call.
// Next available tag: 4
type AddHotlistItemsRequest struct {
	// Resource name of the Hotlist to add new items to.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Resource names of Issues to associate with new HotlistItems added to `parent`.
	Issues []string `protobuf:"bytes,2,rep,name=issues,proto3" json:"issues,omitempty"`
	// Target starting position of the new items.
	// `target_position` must be between [0 and # of items that currently exist in
	// `parent`]. The request will fail if a specified `target_position` is outside
	// of this range.
	// New HotlistItems added to a non-last position of the hotlist will
	// cause ranks of existing HotlistItems below `target_position` to be adjusted.
	// If no `target_position` is given, new items will be added to the end of
	// `parent`.
	TargetPosition       uint32   `protobuf:"varint,3,opt,name=target_position,json=targetPosition,proto3" json:"target_position,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddHotlistItemsRequest) Reset()         { *m = AddHotlistItemsRequest{} }
func (m *AddHotlistItemsRequest) String() string { return proto.CompactTextString(m) }
func (*AddHotlistItemsRequest) ProtoMessage()    {}
func (*AddHotlistItemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a2c4cb0040a55a, []int{5}
}

func (m *AddHotlistItemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddHotlistItemsRequest.Unmarshal(m, b)
}
func (m *AddHotlistItemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddHotlistItemsRequest.Marshal(b, m, deterministic)
}
func (m *AddHotlistItemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddHotlistItemsRequest.Merge(m, src)
}
func (m *AddHotlistItemsRequest) XXX_Size() int {
	return xxx_messageInfo_AddHotlistItemsRequest.Size(m)
}
func (m *AddHotlistItemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddHotlistItemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddHotlistItemsRequest proto.InternalMessageInfo

func (m *AddHotlistItemsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *AddHotlistItemsRequest) GetIssues() []string {
	if m != nil {
		return m.Issues
	}
	return nil
}

func (m *AddHotlistItemsRequest) GetTargetPosition() uint32 {
	if m != nil {
		return m.TargetPosition
	}
	return 0
}

// Request message for a RemoveHotlistItems call.
// Next available tag: 3
type RemoveHotlistItemsRequest struct {
	// Resource name of the Hotlist to remove items from.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Resource names of Issues associated with HotlistItems that should be removed.
	Issues               []string `protobuf:"bytes,2,rep,name=issues,proto3" json:"issues,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveHotlistItemsRequest) Reset()         { *m = RemoveHotlistItemsRequest{} }
func (m *RemoveHotlistItemsRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveHotlistItemsRequest) ProtoMessage()    {}
func (*RemoveHotlistItemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a2c4cb0040a55a, []int{6}
}

func (m *RemoveHotlistItemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveHotlistItemsRequest.Unmarshal(m, b)
}
func (m *RemoveHotlistItemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveHotlistItemsRequest.Marshal(b, m, deterministic)
}
func (m *RemoveHotlistItemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveHotlistItemsRequest.Merge(m, src)
}
func (m *RemoveHotlistItemsRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveHotlistItemsRequest.Size(m)
}
func (m *RemoveHotlistItemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveHotlistItemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveHotlistItemsRequest proto.InternalMessageInfo

func (m *RemoveHotlistItemsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *RemoveHotlistItemsRequest) GetIssues() []string {
	if m != nil {
		return m.Issues
	}
	return nil
}

func init() {
	proto.RegisterType((*GetHotlistRequest)(nil), "monorail.v1.GetHotlistRequest")
	proto.RegisterType((*UpdateHotlistRequest)(nil), "monorail.v1.UpdateHotlistRequest")
	proto.RegisterType((*ListHotlistItemsRequest)(nil), "monorail.v1.ListHotlistItemsRequest")
	proto.RegisterType((*ListHotlistItemsResponse)(nil), "monorail.v1.ListHotlistItemsResponse")
	proto.RegisterType((*RerankHotlistItemsRequest)(nil), "monorail.v1.RerankHotlistItemsRequest")
	proto.RegisterType((*AddHotlistItemsRequest)(nil), "monorail.v1.AddHotlistItemsRequest")
	proto.RegisterType((*RemoveHotlistItemsRequest)(nil), "monorail.v1.RemoveHotlistItemsRequest")
}

func init() { proto.RegisterFile("api/v1/api_proto/hotlists.proto", fileDescriptor_56a2c4cb0040a55a) }

var fileDescriptor_56a2c4cb0040a55a = []byte{
	// 758 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xbe, 0x4e, 0xda, 0xb4, 0x3d, 0xb9, 0xb9, 0xbd, 0x77, 0x6e, 0x7f, 0xd2, 0x44, 0x85, 0xd4,
	0xd0, 0x50, 0x4a, 0x6b, 0xb7, 0xa9, 0x40, 0xc0, 0x8a, 0x54, 0x50, 0xa8, 0x44, 0xa5, 0x2a, 0xfc,
	0x6c, 0xad, 0x49, 0x32, 0x4d, 0x4d, 0x12, 0x8f, 0xf1, 0x4c, 0x22, 0x5a, 0xc1, 0x86, 0x0d, 0x3b,
	0x36, 0x2c, 0x78, 0x04, 0x16, 0x88, 0xb7, 0xe0, 0x09, 0xe0, 0x05, 0xba, 0xe0, 0x29, 0x58, 0xa1,
	0x19, 0x8f, 0x49, 0x1d, 0xdb, 0x09, 0xac, 0x58, 0xfa, 0xcc, 0x77, 0xbe, 0xef, 0x3b, 0xdf, 0x78,
	0x0e, 0x5c, 0xc4, 0xae, 0x6d, 0xf6, 0xb7, 0x4d, 0xec, 0xda, 0x96, 0xeb, 0x51, 0x4e, 0xcd, 0x63,
	0xca, 0x3b, 0x36, 0xe3, 0xcc, 0x90, 0x9f, 0x28, 0xdb, 0xa5, 0x0e, 0xf5, 0xb0, 0xdd, 0x31, 0xfa,
	0xdb, 0x85, 0x72, 0x04, 0x7d, 0x44, 0x30, 0xef, 0x79, 0xc4, 0xa2, 0xf5, 0x67, 0xa4, 0x11, 0x34,
	0x15, 0x4a, 0x2d, 0x4a, 0x5b, 0x1d, 0x62, 0xca, 0xaf, 0x7a, 0xef, 0xc8, 0x3c, 0xb2, 0x49, 0xa7,
	0x69, 0x75, 0x31, 0x6b, 0x2b, 0x44, 0x71, 0x18, 0x41, 0xba, 0x2e, 0x3f, 0x51, 0x87, 0x1b, 0xfe,
	0xa1, 0x92, 0x50, 0x48, 0x21, 0xed, 0xd3, 0xd4, 0xc9, 0x31, 0xee, 0xdb, 0xd4, 0x53, 0xe8, 0x72,
	0x12, 0xda, 0x23, 0x8c, 0xf6, 0xbc, 0x06, 0x51, 0xb8, 0xab, 0x49, 0x38, 0xec, 0x38, 0x94, 0x63,
	0x6e, 0x53, 0x47, 0xf9, 0xd7, 0xf7, 0xe0, 0xbf, 0xfb, 0x84, 0x3f, 0xf0, 0x93, 0xa8, 0x91, 0xe7,
	0x3d, 0xc2, 0x38, 0xda, 0x86, 0x09, 0x07, 0x77, 0x49, 0x5e, 0x2b, 0x69, 0x6b, 0x33, 0xbb, 0xcb,
	0x67, 0xd5, 0xd4, 0xf7, 0xea, 0x22, 0xcc, 0x63, 0xd7, 0x36, 0x1a, 0x5e, 0xbd, 0xd7, 0x32, 0x1a,
	0xb4, 0x6b, 0x06, 0x3d, 0x12, 0xaa, 0x7f, 0xd0, 0x60, 0xee, 0x89, 0xdb, 0xc4, 0x9c, 0x0c, 0x71,
	0x1d, 0xc0, 0x94, 0xca, 0x59, 0xd2, 0x65, 0x2b, 0x73, 0xc6, 0xb9, 0x9c, 0x0d, 0x85, 0x1e, 0x27,
	0x12, 0x70, 0xa0, 0x3b, 0x90, 0xed, 0x49, 0x19, 0x19, 0x71, 0x3e, 0x25, 0x29, 0x0b, 0x86, 0x3f,
	0xa3, 0x11, 0x64, 0x6c, 0xec, 0x89, 0xf8, 0x0e, 0x30, 0x6b, 0xef, 0xa6, 0xcf, 0xaa, 0xa9, 0x1a,
	0xf8, 0x3d, 0xa2, 0xa0, 0x7f, 0xd4, 0x60, 0xf1, 0xa1, 0xcd, 0x82, 0x99, 0xf7, 0x39, 0xe9, 0xb2,
	0xc0, 0xec, 0x2d, 0xc8, 0xb8, 0xd8, 0x23, 0x0e, 0x57, 0xa3, 0xaf, 0x48, 0x57, 0x45, 0xb4, 0x14,
	0xeb, 0x4a, 0xb4, 0xd6, 0x54, 0x03, 0x2a, 0xc2, 0x8c, 0x8b, 0x5b, 0xc4, 0x62, 0xf6, 0x29, 0x91,
	0xb6, 0x26, 0x6b, 0xd3, 0xa2, 0xf0, 0xc8, 0x3e, 0x25, 0x68, 0x09, 0xa6, 0xa9, 0xd7, 0x24, 0x9e,
	0x55, 0x3f, 0xc9, 0xa7, 0x05, 0x73, 0x6d, 0x4a, 0x7e, 0xef, 0x9e, 0xa0, 0x65, 0x00, 0xd9, 0xc7,
	0x69, 0x9b, 0x38, 0xf9, 0x09, 0x79, 0x28, 0x99, 0x1e, 0x8b, 0x82, 0xee, 0x41, 0x3e, 0x6a, 0x96,
	0xb9, 0xd4, 0x61, 0x04, 0x19, 0x30, 0x69, 0x8b, 0x42, 0x5e, 0x2b, 0xa5, 0xd7, 0xb2, 0x95, 0x7c,
	0x5c, 0xb0, 0xd2, 0xa3, 0x0f, 0x43, 0x65, 0x98, 0x75, 0xc8, 0x0b, 0x6e, 0x9d, 0xd3, 0x4b, 0x49,
	0xbd, 0x9c, 0x28, 0x1f, 0xfe, 0xd4, 0xfc, 0xac, 0xc1, 0x52, 0x8d, 0x78, 0xd8, 0x69, 0xc7, 0x65,
	0xf4, 0xfb, 0x3f, 0x07, 0xda, 0x83, 0x9c, 0xba, 0x3f, 0xcb, 0x37, 0x9c, 0x2a, 0xa5, 0x07, 0xe9,
	0xc2, 0x88, 0x74, 0xff, 0x3e, 0x3e, 0xe7, 0x00, 0x6d, 0xc0, 0x2c, 0xc7, 0x5e, 0x8b, 0x70, 0xcb,
	0xa5, 0xcc, 0x16, 0xbf, 0xb1, 0x4c, 0x33, 0xe7, 0x5f, 0xf2, 0x3f, 0xfe, 0xd9, 0xa1, 0x3a, 0xd2,
	0x3f, 0x69, 0xb0, 0x50, 0x6d, 0x36, 0xe3, 0x66, 0xb8, 0x3e, 0x74, 0xcf, 0x63, 0xa6, 0x08, 0xee,
	0x78, 0x07, 0x32, 0x36, 0x63, 0x3d, 0x12, 0x0c, 0x50, 0x94, 0x6d, 0xf3, 0xf0, 0x7f, 0xb8, 0x6d,
	0x5f, 0x60, 0x6a, 0x0a, 0x8a, 0xae, 0x24, 0x98, 0x8e, 0xf8, 0x7d, 0x23, 0x63, 0xef, 0xd2, 0x3e,
	0xf9, 0xc3, 0x96, 0x2b, 0x5f, 0x32, 0x30, 0xad, 0x88, 0x18, 0xe2, 0x00, 0x83, 0x0d, 0x81, 0x2e,
	0x84, 0x7e, 0xb2, 0xc8, 0xea, 0x28, 0xc4, 0xbe, 0x6e, 0x7d, 0xeb, 0xf5, 0xd7, 0x6f, 0xef, 0x52,
	0xeb, 0xfa, 0xaa, 0xe9, 0x7a, 0x6e, 0xc3, 0x8c, 0x81, 0x30, 0x73, 0xc0, 0x75, 0x5b, 0x5b, 0x47,
	0xaf, 0x20, 0x17, 0x5a, 0x27, 0x68, 0x25, 0x44, 0x1c, 0xb7, 0x6a, 0x12, 0xb4, 0x77, 0xa4, 0xf6,
	0xa6, 0xbe, 0x36, 0x42, 0x3b, 0x44, 0x27, 0xe4, 0x5f, 0x42, 0xee, 0x2e, 0xe9, 0x90, 0x81, 0xfc,
	0xb8, 0xb9, 0x17, 0x22, 0x2b, 0xe8, 0x9e, 0x58, 0xf3, 0xbf, 0xa4, 0x1e, 0x52, 0x12, 0xea, 0xef,
	0x35, 0xf8, 0x77, 0xf8, 0xd5, 0xa3, 0xcb, 0x21, 0x07, 0x09, 0x1b, 0xac, 0xb0, 0x3a, 0x06, 0xe5,
	0xaf, 0x0e, 0xfd, 0x86, 0xb4, 0xb5, 0xa5, 0x5f, 0x1b, 0x61, 0x6b, 0xb8, 0x59, 0x38, 0x7b, 0xab,
	0x01, 0x8a, 0xae, 0x06, 0x54, 0x0e, 0xa9, 0x26, 0xee, 0x8e, 0xc4, 0x94, 0x6e, 0x4a, 0x3b, 0x15,
	0x7d, 0x73, 0x84, 0x9d, 0x28, 0xab, 0x30, 0x74, 0x08, 0xb3, 0x43, 0x6f, 0x1c, 0x5d, 0x0a, 0x99,
	0x89, 0xdf, 0x00, 0x89, 0x4e, 0xfe, 0x42, 0x4f, 0xc5, 0x84, 0xc3, 0xaf, 0x30, 0x32, 0x61, 0xc2,
	0x33, 0x4d, 0xe6, 0xad, 0x67, 0x64, 0x65, 0xe7, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x8c,
	0x8e, 0x3d, 0x88, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HotlistsClient is the client API for Hotlists service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HotlistsClient interface {
	// Returns the requested Hotlist.
	//
	// Raises:
	//   NOT_FOUND if the requested hotlist is not found.
	//   PERMISSION_DENIED if the requester is now allowed to view the hotlist.
	//   INVALID_ARGUMENT if the given resource name is not valid.
	GetHotlist(ctx context.Context, in *GetHotlistRequest, opts ...grpc.CallOption) (*Hotlist, error)
	// Updates a hotlist.
	//
	// TODO(crbug/monorail/6988): Document possible errors when implemented.
	UpdateHotlist(ctx context.Context, in *UpdateHotlistRequest, opts ...grpc.CallOption) (*Hotlist, error)
	// Deletes a hotlist.
	//
	DeleteHotlist(ctx context.Context, in *GetHotlistRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Returns a list of all HotlistItems in the hotlist.
	//
	// Raises:
	//   NOT_FOUND if the parent hotlist is not found.
	//   PERMISSION_DENIED if the requester is not allowed to view the hotlist.
	//   INVALID_ARGUMENT if the page_token or given hotlist resource name is not
	//   valid.
	ListHotlistItems(ctx context.Context, in *ListHotlistItemsRequest, opts ...grpc.CallOption) (*ListHotlistItemsResponse, error)
	// Reranks a hotlist's items.
	//
	// Raises:
	//   NOT_FOUND if the hotlist to rerank is not found.
	//   PERMISSION_DENIED if the requester is not allowed to rerank the hotlist.
	//   INVALID_ARGUMENT if the `target_position` is invalid or `hotlist_items`
	//   is empty or contains items not in the Hotlist.
	RerankHotlistItems(ctx context.Context, in *RerankHotlistItemsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Adds new items associated with given issues to a hotlist.
	//
	// TODO(crbug/monorail/7104): Document possible errors when implemented.
	AddHotlistItems(ctx context.Context, in *AddHotlistItemsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Removes items associated with given issues from a hotlist.
	//
	// TODO(crbug/monorali:7104): Document possible errors when implemented.
	RemoveHotlistItems(ctx context.Context, in *RemoveHotlistItemsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}
type hotlistsPRPCClient struct {
	client *prpc.Client
}

func NewHotlistsPRPCClient(client *prpc.Client) HotlistsClient {
	return &hotlistsPRPCClient{client}
}

func (c *hotlistsPRPCClient) GetHotlist(ctx context.Context, in *GetHotlistRequest, opts ...grpc.CallOption) (*Hotlist, error) {
	out := new(Hotlist)
	err := c.client.Call(ctx, "monorail.v1.Hotlists", "GetHotlist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotlistsPRPCClient) UpdateHotlist(ctx context.Context, in *UpdateHotlistRequest, opts ...grpc.CallOption) (*Hotlist, error) {
	out := new(Hotlist)
	err := c.client.Call(ctx, "monorail.v1.Hotlists", "UpdateHotlist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotlistsPRPCClient) DeleteHotlist(ctx context.Context, in *GetHotlistRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.client.Call(ctx, "monorail.v1.Hotlists", "DeleteHotlist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotlistsPRPCClient) ListHotlistItems(ctx context.Context, in *ListHotlistItemsRequest, opts ...grpc.CallOption) (*ListHotlistItemsResponse, error) {
	out := new(ListHotlistItemsResponse)
	err := c.client.Call(ctx, "monorail.v1.Hotlists", "ListHotlistItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotlistsPRPCClient) RerankHotlistItems(ctx context.Context, in *RerankHotlistItemsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.client.Call(ctx, "monorail.v1.Hotlists", "RerankHotlistItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotlistsPRPCClient) AddHotlistItems(ctx context.Context, in *AddHotlistItemsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.client.Call(ctx, "monorail.v1.Hotlists", "AddHotlistItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotlistsPRPCClient) RemoveHotlistItems(ctx context.Context, in *RemoveHotlistItemsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.client.Call(ctx, "monorail.v1.Hotlists", "RemoveHotlistItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type hotlistsClient struct {
	cc grpc.ClientConnInterface
}

func NewHotlistsClient(cc grpc.ClientConnInterface) HotlistsClient {
	return &hotlistsClient{cc}
}

func (c *hotlistsClient) GetHotlist(ctx context.Context, in *GetHotlistRequest, opts ...grpc.CallOption) (*Hotlist, error) {
	out := new(Hotlist)
	err := c.cc.Invoke(ctx, "/monorail.v1.Hotlists/GetHotlist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotlistsClient) UpdateHotlist(ctx context.Context, in *UpdateHotlistRequest, opts ...grpc.CallOption) (*Hotlist, error) {
	out := new(Hotlist)
	err := c.cc.Invoke(ctx, "/monorail.v1.Hotlists/UpdateHotlist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotlistsClient) DeleteHotlist(ctx context.Context, in *GetHotlistRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/monorail.v1.Hotlists/DeleteHotlist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotlistsClient) ListHotlistItems(ctx context.Context, in *ListHotlistItemsRequest, opts ...grpc.CallOption) (*ListHotlistItemsResponse, error) {
	out := new(ListHotlistItemsResponse)
	err := c.cc.Invoke(ctx, "/monorail.v1.Hotlists/ListHotlistItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotlistsClient) RerankHotlistItems(ctx context.Context, in *RerankHotlistItemsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/monorail.v1.Hotlists/RerankHotlistItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotlistsClient) AddHotlistItems(ctx context.Context, in *AddHotlistItemsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/monorail.v1.Hotlists/AddHotlistItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotlistsClient) RemoveHotlistItems(ctx context.Context, in *RemoveHotlistItemsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/monorail.v1.Hotlists/RemoveHotlistItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HotlistsServer is the server API for Hotlists service.
type HotlistsServer interface {
	// Returns the requested Hotlist.
	//
	// Raises:
	//   NOT_FOUND if the requested hotlist is not found.
	//   PERMISSION_DENIED if the requester is now allowed to view the hotlist.
	//   INVALID_ARGUMENT if the given resource name is not valid.
	GetHotlist(context.Context, *GetHotlistRequest) (*Hotlist, error)
	// Updates a hotlist.
	//
	// TODO(crbug/monorail/6988): Document possible errors when implemented.
	UpdateHotlist(context.Context, *UpdateHotlistRequest) (*Hotlist, error)
	// Deletes a hotlist.
	//
	DeleteHotlist(context.Context, *GetHotlistRequest) (*empty.Empty, error)
	// Returns a list of all HotlistItems in the hotlist.
	//
	// Raises:
	//   NOT_FOUND if the parent hotlist is not found.
	//   PERMISSION_DENIED if the requester is not allowed to view the hotlist.
	//   INVALID_ARGUMENT if the page_token or given hotlist resource name is not
	//   valid.
	ListHotlistItems(context.Context, *ListHotlistItemsRequest) (*ListHotlistItemsResponse, error)
	// Reranks a hotlist's items.
	//
	// Raises:
	//   NOT_FOUND if the hotlist to rerank is not found.
	//   PERMISSION_DENIED if the requester is not allowed to rerank the hotlist.
	//   INVALID_ARGUMENT if the `target_position` is invalid or `hotlist_items`
	//   is empty or contains items not in the Hotlist.
	RerankHotlistItems(context.Context, *RerankHotlistItemsRequest) (*empty.Empty, error)
	// Adds new items associated with given issues to a hotlist.
	//
	// TODO(crbug/monorail/7104): Document possible errors when implemented.
	AddHotlistItems(context.Context, *AddHotlistItemsRequest) (*empty.Empty, error)
	// Removes items associated with given issues from a hotlist.
	//
	// TODO(crbug/monorali:7104): Document possible errors when implemented.
	RemoveHotlistItems(context.Context, *RemoveHotlistItemsRequest) (*empty.Empty, error)
}

// UnimplementedHotlistsServer can be embedded to have forward compatible implementations.
type UnimplementedHotlistsServer struct {
}

func (*UnimplementedHotlistsServer) GetHotlist(ctx context.Context, req *GetHotlistRequest) (*Hotlist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHotlist not implemented")
}
func (*UnimplementedHotlistsServer) UpdateHotlist(ctx context.Context, req *UpdateHotlistRequest) (*Hotlist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHotlist not implemented")
}
func (*UnimplementedHotlistsServer) DeleteHotlist(ctx context.Context, req *GetHotlistRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHotlist not implemented")
}
func (*UnimplementedHotlistsServer) ListHotlistItems(ctx context.Context, req *ListHotlistItemsRequest) (*ListHotlistItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHotlistItems not implemented")
}
func (*UnimplementedHotlistsServer) RerankHotlistItems(ctx context.Context, req *RerankHotlistItemsRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RerankHotlistItems not implemented")
}
func (*UnimplementedHotlistsServer) AddHotlistItems(ctx context.Context, req *AddHotlistItemsRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddHotlistItems not implemented")
}
func (*UnimplementedHotlistsServer) RemoveHotlistItems(ctx context.Context, req *RemoveHotlistItemsRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveHotlistItems not implemented")
}

func RegisterHotlistsServer(s prpc.Registrar, srv HotlistsServer) {
	s.RegisterService(&_Hotlists_serviceDesc, srv)
}

func _Hotlists_GetHotlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHotlistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotlistsServer).GetHotlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monorail.v1.Hotlists/GetHotlist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotlistsServer).GetHotlist(ctx, req.(*GetHotlistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hotlists_UpdateHotlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHotlistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotlistsServer).UpdateHotlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monorail.v1.Hotlists/UpdateHotlist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotlistsServer).UpdateHotlist(ctx, req.(*UpdateHotlistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hotlists_DeleteHotlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHotlistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotlistsServer).DeleteHotlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monorail.v1.Hotlists/DeleteHotlist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotlistsServer).DeleteHotlist(ctx, req.(*GetHotlistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hotlists_ListHotlistItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHotlistItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotlistsServer).ListHotlistItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monorail.v1.Hotlists/ListHotlistItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotlistsServer).ListHotlistItems(ctx, req.(*ListHotlistItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hotlists_RerankHotlistItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RerankHotlistItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotlistsServer).RerankHotlistItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monorail.v1.Hotlists/RerankHotlistItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotlistsServer).RerankHotlistItems(ctx, req.(*RerankHotlistItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hotlists_AddHotlistItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddHotlistItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotlistsServer).AddHotlistItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monorail.v1.Hotlists/AddHotlistItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotlistsServer).AddHotlistItems(ctx, req.(*AddHotlistItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hotlists_RemoveHotlistItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveHotlistItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotlistsServer).RemoveHotlistItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monorail.v1.Hotlists/RemoveHotlistItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotlistsServer).RemoveHotlistItems(ctx, req.(*RemoveHotlistItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hotlists_serviceDesc = grpc.ServiceDesc{
	ServiceName: "monorail.v1.Hotlists",
	HandlerType: (*HotlistsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHotlist",
			Handler:    _Hotlists_GetHotlist_Handler,
		},
		{
			MethodName: "UpdateHotlist",
			Handler:    _Hotlists_UpdateHotlist_Handler,
		},
		{
			MethodName: "DeleteHotlist",
			Handler:    _Hotlists_DeleteHotlist_Handler,
		},
		{
			MethodName: "ListHotlistItems",
			Handler:    _Hotlists_ListHotlistItems_Handler,
		},
		{
			MethodName: "RerankHotlistItems",
			Handler:    _Hotlists_RerankHotlistItems_Handler,
		},
		{
			MethodName: "AddHotlistItems",
			Handler:    _Hotlists_AddHotlistItems_Handler,
		},
		{
			MethodName: "RemoveHotlistItems",
			Handler:    _Hotlists_RemoveHotlistItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/api_proto/hotlists.proto",
}
