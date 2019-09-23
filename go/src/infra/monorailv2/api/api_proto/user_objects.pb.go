// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api_proto/user_objects.proto

package monorail

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// TODO(jojwang): monorail:1701, fill User with all info necessary for
// creating a user profile page.
// Next available tag: 7
type User struct {
	DisplayName          string     `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	UserId               int64      `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	IsSiteAdmin          bool       `protobuf:"varint,3,opt,name=is_site_admin,json=isSiteAdmin,proto3" json:"is_site_admin,omitempty"`
	Availability         string     `protobuf:"bytes,4,opt,name=availability,proto3" json:"availability,omitempty"`
	LinkedParentRef      *UserRef   `protobuf:"bytes,5,opt,name=linked_parent_ref,json=linkedParentRef,proto3" json:"linked_parent_ref,omitempty"`
	LinkedChildRefs      []*UserRef `protobuf:"bytes,6,rep,name=linked_child_refs,json=linkedChildRefs,proto3" json:"linked_child_refs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_e651956a3fdc871c, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *User) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *User) GetIsSiteAdmin() bool {
	if m != nil {
		return m.IsSiteAdmin
	}
	return false
}

func (m *User) GetAvailability() string {
	if m != nil {
		return m.Availability
	}
	return ""
}

func (m *User) GetLinkedParentRef() *UserRef {
	if m != nil {
		return m.LinkedParentRef
	}
	return nil
}

func (m *User) GetLinkedChildRefs() []*UserRef {
	if m != nil {
		return m.LinkedChildRefs
	}
	return nil
}

// Next available tag: 3
type UserPrefValue struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserPrefValue) Reset()         { *m = UserPrefValue{} }
func (m *UserPrefValue) String() string { return proto.CompactTextString(m) }
func (*UserPrefValue) ProtoMessage()    {}
func (*UserPrefValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_e651956a3fdc871c, []int{1}
}

func (m *UserPrefValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserPrefValue.Unmarshal(m, b)
}
func (m *UserPrefValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserPrefValue.Marshal(b, m, deterministic)
}
func (m *UserPrefValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserPrefValue.Merge(m, src)
}
func (m *UserPrefValue) XXX_Size() int {
	return xxx_messageInfo_UserPrefValue.Size(m)
}
func (m *UserPrefValue) XXX_DiscardUnknown() {
	xxx_messageInfo_UserPrefValue.DiscardUnknown(m)
}

var xxx_messageInfo_UserPrefValue proto.InternalMessageInfo

func (m *UserPrefValue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserPrefValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Next available tag: 6
type UserProjects struct {
	UserRef              *UserRef `protobuf:"bytes,1,opt,name=user_ref,json=userRef,proto3" json:"user_ref,omitempty"`
	OwnerOf              []string `protobuf:"bytes,2,rep,name=owner_of,json=ownerOf,proto3" json:"owner_of,omitempty"`
	MemberOf             []string `protobuf:"bytes,3,rep,name=member_of,json=memberOf,proto3" json:"member_of,omitempty"`
	ContributorTo        []string `protobuf:"bytes,4,rep,name=contributor_to,json=contributorTo,proto3" json:"contributor_to,omitempty"`
	StarredProjects      []string `protobuf:"bytes,5,rep,name=starred_projects,json=starredProjects,proto3" json:"starred_projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserProjects) Reset()         { *m = UserProjects{} }
func (m *UserProjects) String() string { return proto.CompactTextString(m) }
func (*UserProjects) ProtoMessage()    {}
func (*UserProjects) Descriptor() ([]byte, []int) {
	return fileDescriptor_e651956a3fdc871c, []int{2}
}

func (m *UserProjects) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserProjects.Unmarshal(m, b)
}
func (m *UserProjects) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserProjects.Marshal(b, m, deterministic)
}
func (m *UserProjects) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserProjects.Merge(m, src)
}
func (m *UserProjects) XXX_Size() int {
	return xxx_messageInfo_UserProjects.Size(m)
}
func (m *UserProjects) XXX_DiscardUnknown() {
	xxx_messageInfo_UserProjects.DiscardUnknown(m)
}

var xxx_messageInfo_UserProjects proto.InternalMessageInfo

func (m *UserProjects) GetUserRef() *UserRef {
	if m != nil {
		return m.UserRef
	}
	return nil
}

func (m *UserProjects) GetOwnerOf() []string {
	if m != nil {
		return m.OwnerOf
	}
	return nil
}

func (m *UserProjects) GetMemberOf() []string {
	if m != nil {
		return m.MemberOf
	}
	return nil
}

func (m *UserProjects) GetContributorTo() []string {
	if m != nil {
		return m.ContributorTo
	}
	return nil
}

func (m *UserProjects) GetStarredProjects() []string {
	if m != nil {
		return m.StarredProjects
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "monorail.User")
	proto.RegisterType((*UserPrefValue)(nil), "monorail.UserPrefValue")
	proto.RegisterType((*UserProjects)(nil), "monorail.UserProjects")
}

func init() { proto.RegisterFile("api/api_proto/user_objects.proto", fileDescriptor_e651956a3fdc871c) }

var fileDescriptor_e651956a3fdc871c = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x56, 0x9a, 0xb6, 0x9b, 0x4e, 0x5b, 0x96, 0xb5, 0x90, 0x30, 0xcb, 0x25, 0x44, 0x42, 0x0a,
	0x12, 0xea, 0x4a, 0x70, 0xe2, 0xc0, 0x01, 0x71, 0xe2, 0xc2, 0xae, 0xcc, 0xcf, 0x35, 0x72, 0x92,
	0xb1, 0x18, 0x88, 0xed, 0xc8, 0x76, 0x17, 0xed, 0x1b, 0xf0, 0x64, 0x3c, 0x17, 0xb2, 0xd3, 0x95,
	0xda, 0x03, 0x7b, 0xcb, 0x7c, 0x3f, 0xc9, 0x37, 0xdf, 0x04, 0x4a, 0x39, 0xd2, 0x95, 0x1c, 0xa9,
	0x19, 0x9d, 0x0d, 0xf6, 0x6a, 0xef, 0xd1, 0x35, 0xb6, 0xfd, 0x89, 0x5d, 0xf0, 0xbb, 0x04, 0xb1,
	0x42, 0x5b, 0x63, 0x9d, 0xa4, 0xe1, 0xf2, 0xf2, 0x54, 0xdb, 0x59, 0xad, 0xad, 0x99, 0x54, 0xd5,
	0x9f, 0x19, 0xcc, 0xbf, 0x79, 0x74, 0xec, 0x05, 0x6c, 0x7a, 0xf2, 0xe3, 0x20, 0xef, 0x1a, 0x23,
	0x35, 0xf2, 0xac, 0xcc, 0xea, 0x95, 0x58, 0x1f, 0xb0, 0xcf, 0x52, 0x23, 0x7b, 0x0a, 0x67, 0xe9,
	0x3b, 0xd4, 0xf3, 0x59, 0x99, 0xd5, 0xb9, 0x58, 0xc6, 0xf1, 0x53, 0xcf, 0x2a, 0xd8, 0x92, 0x6f,
	0x3c, 0x05, 0x6c, 0x64, 0xaf, 0xc9, 0xf0, 0xbc, 0xcc, 0xea, 0x42, 0xac, 0xc9, 0x7f, 0xa1, 0x80,
	0x1f, 0x22, 0xc4, 0x2a, 0xd8, 0xc8, 0x5b, 0x49, 0x83, 0x6c, 0x69, 0xa0, 0x70, 0xc7, 0xe7, 0xe9,
	0xfd, 0x27, 0x18, 0x7b, 0x0f, 0x17, 0x03, 0x99, 0x5f, 0xd8, 0x37, 0xa3, 0x74, 0x68, 0x42, 0xe3,
	0x50, 0xf1, 0x45, 0x99, 0xd5, 0xeb, 0x37, 0x17, 0xbb, 0xfb, 0x75, 0x76, 0x31, 0xae, 0x40, 0x25,
	0xce, 0x27, 0xed, 0x4d, 0x92, 0x0a, 0x54, 0x47, 0xf6, 0xee, 0x07, 0x0d, 0x7d, 0x74, 0x7b, 0xbe,
	0x2c, 0xf3, 0x07, 0xed, 0x1f, 0xa3, 0x54, 0xa0, 0xf2, 0xd5, 0x3b, 0xd8, 0x46, 0xee, 0xc6, 0xa1,
	0xfa, 0x2e, 0x87, 0x3d, 0x32, 0x06, 0xf3, 0xa3, 0x2a, 0xd2, 0x33, 0x7b, 0x02, 0x8b, 0xdb, 0x48,
	0xa6, 0x06, 0x56, 0x62, 0x1a, 0xaa, 0xbf, 0x19, 0x6c, 0x26, 0xaf, 0x4d, 0x27, 0x60, 0xaf, 0xa1,
	0x48, 0x55, 0xc5, 0x05, 0xb2, 0xff, 0x2d, 0x90, 0xda, 0x8c, 0xc1, 0x9f, 0x41, 0x61, 0x7f, 0x9b,
	0x78, 0x41, 0xc5, 0x67, 0x65, 0x5e, 0xaf, 0xc4, 0x59, 0x9a, 0xaf, 0x15, 0x7b, 0x0e, 0x2b, 0x8d,
	0xba, 0x9d, 0xb8, 0x3c, 0x71, 0xc5, 0x04, 0x5c, 0x2b, 0xf6, 0x12, 0x1e, 0x75, 0xd6, 0x04, 0x47,
	0xed, 0x3e, 0x58, 0xd7, 0x04, 0xcb, 0xe7, 0x49, 0xb1, 0x3d, 0x42, 0xbf, 0x5a, 0xf6, 0x0a, 0x1e,
	0xfb, 0x20, 0x9d, 0x8b, 0xbd, 0x1e, 0x02, 0xf2, 0x45, 0x12, 0x9e, 0x1f, 0xf0, 0xfb, 0xdc, 0xed,
	0x32, 0xfd, 0x15, 0x6f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x53, 0x81, 0xce, 0x5f, 0x02,
	0x00, 0x00,
}
