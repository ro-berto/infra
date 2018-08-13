// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/qscheduler/qslib/types/types.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import account "infra/qscheduler/qslib/types/account"
import task "infra/qscheduler/qslib/types/task"
import vector "infra/qscheduler/qslib/types/vector"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Worker represents a resource that can run 1 task at a time. This corresponds
// to the swarming concept of a Bot. This representation considers only the
// subset of Labels that are Provisionable (can be changed by running a task),
// because the quota scheduler algorithm is expected to run against a pool of
// otherwise homogenous workers.
type Worker struct {
	Id                   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Labels               []string  `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	RunningTask          *task.Run `protobuf:"bytes,3,opt,name=running_task,json=runningTask,proto3" json:"running_task,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Worker) Reset()         { *m = Worker{} }
func (m *Worker) String() string { return proto.CompactTextString(m) }
func (*Worker) ProtoMessage()    {}
func (*Worker) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_405a09ae148a64be, []int{0}
}
func (m *Worker) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Worker.Unmarshal(m, b)
}
func (m *Worker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Worker.Marshal(b, m, deterministic)
}
func (dst *Worker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Worker.Merge(dst, src)
}
func (m *Worker) XXX_Size() int {
	return xxx_messageInfo_Worker.Size(m)
}
func (m *Worker) XXX_DiscardUnknown() {
	xxx_messageInfo_Worker.DiscardUnknown(m)
}

var xxx_messageInfo_Worker proto.InternalMessageInfo

func (m *Worker) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Worker) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Worker) GetRunningTask() *task.Run {
	if m != nil {
		return m.RunningTask
	}
	return nil
}

// State represents the overall state of a quota scheduler worker pool,
// account set, and task queue. This is represented separately from
// configuration information. The state is expected to be updated frequently,
// on each scheduler tick.
type State struct {
	// Requests that are waiting to be assigned to a worker, keyed by
	// request id.
	Requests map[string]*task.Request `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Balance of all quota accounts for this pool, keyed by account id.
	Balances map[string]*vector.Vector `protobuf:"bytes,2,rep,name=balances,proto3" json:"balances,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Workers that may run tasks, and their states, keyed by worker id.
	Workers              map[string]*Worker `protobuf:"bytes,3,rep,name=workers,proto3" json:"workers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_405a09ae148a64be, []int{1}
}
func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (dst *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(dst, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetRequests() map[string]*task.Request {
	if m != nil {
		return m.Requests
	}
	return nil
}

func (m *State) GetBalances() map[string]*vector.Vector {
	if m != nil {
		return m.Balances
	}
	return nil
}

func (m *State) GetWorkers() map[string]*Worker {
	if m != nil {
		return m.Workers
	}
	return nil
}

// Config represents configuration information about the behavior of accounts
// for this quota scheduler pool.
type Config struct {
	// Configuration for a given account, keyed by account id.
	AccountConfigs       map[string]*account.Config `protobuf:"bytes,1,rep,name=account_configs,json=accountConfigs,proto3" json:"account_configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_405a09ae148a64be, []int{2}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (dst *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(dst, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetAccountConfigs() map[string]*account.Config {
	if m != nil {
		return m.AccountConfigs
	}
	return nil
}

func init() {
	proto.RegisterType((*Worker)(nil), "types.Worker")
	proto.RegisterType((*State)(nil), "types.State")
	proto.RegisterMapType((map[string]*vector.Vector)(nil), "types.State.BalancesEntry")
	proto.RegisterMapType((map[string]*task.Request)(nil), "types.State.RequestsEntry")
	proto.RegisterMapType((map[string]*Worker)(nil), "types.State.WorkersEntry")
	proto.RegisterType((*Config)(nil), "types.Config")
	proto.RegisterMapType((map[string]*account.Config)(nil), "types.Config.AccountConfigsEntry")
}

func init() {
	proto.RegisterFile("infra/qscheduler/qslib/types/types.proto", fileDescriptor_types_405a09ae148a64be)
}

var fileDescriptor_types_405a09ae148a64be = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6a, 0xe3, 0x30,
	0x10, 0xc6, 0xb1, 0x4d, 0xbc, 0x9b, 0xc9, 0xbf, 0x45, 0x0b, 0x8b, 0xd7, 0x27, 0x6f, 0x76, 0x17,
	0x7c, 0x08, 0x4e, 0x49, 0xa0, 0x94, 0xde, 0xda, 0xd2, 0x43, 0xd3, 0x9b, 0x5a, 0xda, 0x5b, 0x83,
	0xec, 0x28, 0xa9, 0xb1, 0x91, 0x13, 0x49, 0x4e, 0xc9, 0x5b, 0xf4, 0x31, 0xfa, 0x98, 0x25, 0x92,
	0x5c, 0xec, 0x12, 0x42, 0x2f, 0x96, 0x47, 0xf3, 0xfd, 0x86, 0x6f, 0x46, 0x03, 0x61, 0xca, 0x96,
	0x9c, 0x8c, 0x37, 0x22, 0x79, 0xa6, 0x8b, 0x32, 0xa7, 0x7c, 0xbc, 0x11, 0x79, 0x1a, 0x8f, 0xe5,
	0x6e, 0x4d, 0x85, 0xfe, 0x46, 0x6b, 0x5e, 0xc8, 0x02, 0xb5, 0x54, 0xe0, 0x8f, 0x8e, 0x03, 0x44,
	0x64, 0xea, 0xa3, 0x21, 0xff, 0xe4, 0xa8, 0x7a, 0x4b, 0x13, 0x59, 0x70, 0x73, 0x18, 0x62, 0x72,
	0x94, 0x20, 0x49, 0x52, 0x94, 0x4c, 0x56, 0xa7, 0x66, 0x86, 0x4f, 0xe0, 0x3e, 0x16, 0x3c, 0xa3,
	0x1c, 0xf5, 0xc1, 0x4e, 0x17, 0x9e, 0x15, 0x58, 0x61, 0x1b, 0xdb, 0xe9, 0x02, 0xfd, 0x02, 0x37,
	0x27, 0x31, 0xcd, 0x85, 0x67, 0x07, 0x4e, 0xd8, 0xc6, 0x26, 0x42, 0x23, 0xe8, 0xf2, 0x92, 0xb1,
	0x94, 0xad, 0xe6, 0x7b, 0xb7, 0x9e, 0x13, 0x58, 0x61, 0x67, 0xd2, 0x8e, 0x94, 0x75, 0x5c, 0x32,
	0xdc, 0x31, 0xe9, 0x7b, 0x22, 0xb2, 0xe1, 0xab, 0x03, 0xad, 0x3b, 0x49, 0x24, 0x45, 0xa7, 0xf0,
	0x9d, 0xd3, 0x4d, 0x49, 0x85, 0x14, 0x9e, 0x15, 0x38, 0x61, 0x67, 0xe2, 0x47, 0x7a, 0x48, 0x2a,
	0x1f, 0x61, 0x93, 0xbc, 0x66, 0x92, 0xef, 0xf0, 0x87, 0x76, 0xcf, 0xc5, 0x24, 0x27, 0x2c, 0xa1,
	0xda, 0xc9, 0x67, 0xee, 0xd2, 0x24, 0x0d, 0x57, 0x69, 0xd1, 0x14, 0xbe, 0xbd, 0xa8, 0xce, 0x84,
	0xe7, 0x28, 0xec, 0x77, 0x03, 0xd3, 0x5d, 0x1b, 0xaa, 0x52, 0xfa, 0x33, 0xe8, 0x35, 0x7c, 0xa0,
	0x1f, 0xe0, 0x64, 0x74, 0x67, 0xc6, 0xb2, 0xff, 0x45, 0x7f, 0xa1, 0xb5, 0x25, 0x79, 0x49, 0x3d,
	0x5b, 0x35, 0xde, 0x33, 0x8d, 0x6b, 0x0a, 0xeb, 0xdc, 0xb9, 0x7d, 0x66, 0xf9, 0xb7, 0xd0, 0x6b,
	0x78, 0x3b, 0x50, 0xeb, 0x5f, 0xb3, 0x56, 0x3f, 0x32, 0xef, 0xf9, 0xa0, 0x8e, 0x7a, 0xb1, 0x1b,
	0xe8, 0xd6, 0x1d, 0x7f, 0xc9, 0x97, 0xea, 0x56, 0x53, 0xb5, 0x52, 0xc3, 0x37, 0x0b, 0xdc, 0xab,
	0x82, 0x2d, 0xd3, 0x15, 0x9a, 0xc1, 0xc0, 0xac, 0xc3, 0x3c, 0x51, 0x37, 0xd5, 0xd3, 0xfc, 0x31,
	0xb4, 0xd6, 0x45, 0x17, 0x5a, 0xa4, 0x23, 0x33, 0xb3, 0x3e, 0x69, 0x5c, 0xfa, 0x18, 0x7e, 0x1e,
	0x90, 0x1d, 0x30, 0xfa, 0xbf, 0x69, 0x74, 0x10, 0x55, 0x1b, 0xa9, 0xb9, 0x9a, 0xd5, 0xd8, 0x55,
	0x4b, 0x3a, 0x7d, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xce, 0x39, 0xcc, 0xbd, 0x6b, 0x03, 0x00, 0x00,
}
