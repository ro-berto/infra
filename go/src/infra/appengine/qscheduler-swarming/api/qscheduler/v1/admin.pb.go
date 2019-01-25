// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/appengine/qscheduler-swarming/api/qscheduler/v1/admin.proto

package qscheduler

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	fmt "fmt"
	scheduler "infra/qscheduler/qslib/scheduler"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type CreateSchedulerPoolRequest struct {
	// PoolId is the unique id of this scheduler pool.
	PoolId string `protobuf:"bytes,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	// Config is the scheduler configuration for the scheduler to create.
	Config               *SchedulerPoolConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateSchedulerPoolRequest) Reset()         { *m = CreateSchedulerPoolRequest{} }
func (m *CreateSchedulerPoolRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSchedulerPoolRequest) ProtoMessage()    {}
func (*CreateSchedulerPoolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19eb132fa8b54f73, []int{0}
}

func (m *CreateSchedulerPoolRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSchedulerPoolRequest.Unmarshal(m, b)
}
func (m *CreateSchedulerPoolRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSchedulerPoolRequest.Marshal(b, m, deterministic)
}
func (m *CreateSchedulerPoolRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSchedulerPoolRequest.Merge(m, src)
}
func (m *CreateSchedulerPoolRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSchedulerPoolRequest.Size(m)
}
func (m *CreateSchedulerPoolRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSchedulerPoolRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSchedulerPoolRequest proto.InternalMessageInfo

func (m *CreateSchedulerPoolRequest) GetPoolId() string {
	if m != nil {
		return m.PoolId
	}
	return ""
}

func (m *CreateSchedulerPoolRequest) GetConfig() *SchedulerPoolConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type CreateSchedulerPoolResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSchedulerPoolResponse) Reset()         { *m = CreateSchedulerPoolResponse{} }
func (m *CreateSchedulerPoolResponse) String() string { return proto.CompactTextString(m) }
func (*CreateSchedulerPoolResponse) ProtoMessage()    {}
func (*CreateSchedulerPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_19eb132fa8b54f73, []int{1}
}

func (m *CreateSchedulerPoolResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSchedulerPoolResponse.Unmarshal(m, b)
}
func (m *CreateSchedulerPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSchedulerPoolResponse.Marshal(b, m, deterministic)
}
func (m *CreateSchedulerPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSchedulerPoolResponse.Merge(m, src)
}
func (m *CreateSchedulerPoolResponse) XXX_Size() int {
	return xxx_messageInfo_CreateSchedulerPoolResponse.Size(m)
}
func (m *CreateSchedulerPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSchedulerPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSchedulerPoolResponse proto.InternalMessageInfo

type CreateAccountRequest struct {
	// PoolID is the id of the scheduler to create an account within.
	PoolId string `protobuf:"bytes,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	// AccountId is the unique id of the account (within the given pool).
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// Config is the quota account config for the quota account to create.
	Config               *scheduler.AccountConfig `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CreateAccountRequest) Reset()         { *m = CreateAccountRequest{} }
func (m *CreateAccountRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAccountRequest) ProtoMessage()    {}
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19eb132fa8b54f73, []int{2}
}

func (m *CreateAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountRequest.Unmarshal(m, b)
}
func (m *CreateAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountRequest.Marshal(b, m, deterministic)
}
func (m *CreateAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountRequest.Merge(m, src)
}
func (m *CreateAccountRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAccountRequest.Size(m)
}
func (m *CreateAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountRequest proto.InternalMessageInfo

func (m *CreateAccountRequest) GetPoolId() string {
	if m != nil {
		return m.PoolId
	}
	return ""
}

func (m *CreateAccountRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *CreateAccountRequest) GetConfig() *scheduler.AccountConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type CreateAccountResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccountResponse) Reset()         { *m = CreateAccountResponse{} }
func (m *CreateAccountResponse) String() string { return proto.CompactTextString(m) }
func (*CreateAccountResponse) ProtoMessage()    {}
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_19eb132fa8b54f73, []int{3}
}

func (m *CreateAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountResponse.Unmarshal(m, b)
}
func (m *CreateAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountResponse.Marshal(b, m, deterministic)
}
func (m *CreateAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountResponse.Merge(m, src)
}
func (m *CreateAccountResponse) XXX_Size() int {
	return xxx_messageInfo_CreateAccountResponse.Size(m)
}
func (m *CreateAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountResponse proto.InternalMessageInfo

type ListAccountsRequest struct {
	// PoolID is the id of the scheduler to list accounts from.
	PoolId               string   `protobuf:"bytes,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAccountsRequest) Reset()         { *m = ListAccountsRequest{} }
func (m *ListAccountsRequest) String() string { return proto.CompactTextString(m) }
func (*ListAccountsRequest) ProtoMessage()    {}
func (*ListAccountsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19eb132fa8b54f73, []int{4}
}

func (m *ListAccountsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAccountsRequest.Unmarshal(m, b)
}
func (m *ListAccountsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAccountsRequest.Marshal(b, m, deterministic)
}
func (m *ListAccountsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAccountsRequest.Merge(m, src)
}
func (m *ListAccountsRequest) XXX_Size() int {
	return xxx_messageInfo_ListAccountsRequest.Size(m)
}
func (m *ListAccountsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAccountsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAccountsRequest proto.InternalMessageInfo

func (m *ListAccountsRequest) GetPoolId() string {
	if m != nil {
		return m.PoolId
	}
	return ""
}

type ListAccountsResponse struct {
	Accounts             map[string]*scheduler.AccountConfig `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *ListAccountsResponse) Reset()         { *m = ListAccountsResponse{} }
func (m *ListAccountsResponse) String() string { return proto.CompactTextString(m) }
func (*ListAccountsResponse) ProtoMessage()    {}
func (*ListAccountsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_19eb132fa8b54f73, []int{5}
}

func (m *ListAccountsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAccountsResponse.Unmarshal(m, b)
}
func (m *ListAccountsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAccountsResponse.Marshal(b, m, deterministic)
}
func (m *ListAccountsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAccountsResponse.Merge(m, src)
}
func (m *ListAccountsResponse) XXX_Size() int {
	return xxx_messageInfo_ListAccountsResponse.Size(m)
}
func (m *ListAccountsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAccountsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAccountsResponse proto.InternalMessageInfo

func (m *ListAccountsResponse) GetAccounts() map[string]*scheduler.AccountConfig {
	if m != nil {
		return m.Accounts
	}
	return nil
}

type SchedulerPoolConfig struct {
	// Labels is a list of swarming dimensions in "key:value" form. This corresponds
	// to swarming.ExternalSchedulerConfig.dimensionsions; it is the minimal set
	// of dimensions for tasks or bots that will use this scheduler.
	Labels               []string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SchedulerPoolConfig) Reset()         { *m = SchedulerPoolConfig{} }
func (m *SchedulerPoolConfig) String() string { return proto.CompactTextString(m) }
func (*SchedulerPoolConfig) ProtoMessage()    {}
func (*SchedulerPoolConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_19eb132fa8b54f73, []int{6}
}

func (m *SchedulerPoolConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchedulerPoolConfig.Unmarshal(m, b)
}
func (m *SchedulerPoolConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchedulerPoolConfig.Marshal(b, m, deterministic)
}
func (m *SchedulerPoolConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchedulerPoolConfig.Merge(m, src)
}
func (m *SchedulerPoolConfig) XXX_Size() int {
	return xxx_messageInfo_SchedulerPoolConfig.Size(m)
}
func (m *SchedulerPoolConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SchedulerPoolConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SchedulerPoolConfig proto.InternalMessageInfo

func (m *SchedulerPoolConfig) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type InspectPoolRequest struct {
	PoolId               string   `protobuf:"bytes,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InspectPoolRequest) Reset()         { *m = InspectPoolRequest{} }
func (m *InspectPoolRequest) String() string { return proto.CompactTextString(m) }
func (*InspectPoolRequest) ProtoMessage()    {}
func (*InspectPoolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19eb132fa8b54f73, []int{7}
}

func (m *InspectPoolRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InspectPoolRequest.Unmarshal(m, b)
}
func (m *InspectPoolRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InspectPoolRequest.Marshal(b, m, deterministic)
}
func (m *InspectPoolRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InspectPoolRequest.Merge(m, src)
}
func (m *InspectPoolRequest) XXX_Size() int {
	return xxx_messageInfo_InspectPoolRequest.Size(m)
}
func (m *InspectPoolRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InspectPoolRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InspectPoolRequest proto.InternalMessageInfo

func (m *InspectPoolRequest) GetPoolId() string {
	if m != nil {
		return m.PoolId
	}
	return ""
}

type InspectPoolResponse struct {
	// NumWaitingTasks is the number of waiting tasks.
	NumWaitingTasks int32 `protobuf:"varint,1,opt,name=num_waiting_tasks,json=numWaitingTasks,proto3" json:"num_waiting_tasks,omitempty"`
	// NumRunningTasks is the number of running tasks.
	NumRunningTasks int32 `protobuf:"varint,2,opt,name=num_running_tasks,json=numRunningTasks,proto3" json:"num_running_tasks,omitempty"`
	// IdleBots is the number of idle bots.
	NumIdleBots int32 `protobuf:"varint,3,opt,name=num_idle_bots,json=numIdleBots,proto3" json:"num_idle_bots,omitempty"`
	// AccountBalances is the set of balances for all accounts.
	AccountBalances map[string]*scheduler.StateProto_Balance `protobuf:"bytes,4,rep,name=account_balances,json=accountBalances,proto3" json:"account_balances,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// RunningTasks is a description of the running tasks according to
	// quotascheduler.
	RunningTasks []*InspectPoolResponse_RunningTask `protobuf:"bytes,5,rep,name=running_tasks,json=runningTasks,proto3" json:"running_tasks,omitempty"`
	// WaitingTasks is a description of the tasks that are waiting
	// according to quotascheduler.
	WaitingTasks         []*InspectPoolResponse_WaitingTask `protobuf:"bytes,6,rep,name=waiting_tasks,json=waitingTasks,proto3" json:"waiting_tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *InspectPoolResponse) Reset()         { *m = InspectPoolResponse{} }
func (m *InspectPoolResponse) String() string { return proto.CompactTextString(m) }
func (*InspectPoolResponse) ProtoMessage()    {}
func (*InspectPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_19eb132fa8b54f73, []int{8}
}

func (m *InspectPoolResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InspectPoolResponse.Unmarshal(m, b)
}
func (m *InspectPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InspectPoolResponse.Marshal(b, m, deterministic)
}
func (m *InspectPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InspectPoolResponse.Merge(m, src)
}
func (m *InspectPoolResponse) XXX_Size() int {
	return xxx_messageInfo_InspectPoolResponse.Size(m)
}
func (m *InspectPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InspectPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InspectPoolResponse proto.InternalMessageInfo

func (m *InspectPoolResponse) GetNumWaitingTasks() int32 {
	if m != nil {
		return m.NumWaitingTasks
	}
	return 0
}

func (m *InspectPoolResponse) GetNumRunningTasks() int32 {
	if m != nil {
		return m.NumRunningTasks
	}
	return 0
}

func (m *InspectPoolResponse) GetNumIdleBots() int32 {
	if m != nil {
		return m.NumIdleBots
	}
	return 0
}

func (m *InspectPoolResponse) GetAccountBalances() map[string]*scheduler.StateProto_Balance {
	if m != nil {
		return m.AccountBalances
	}
	return nil
}

func (m *InspectPoolResponse) GetRunningTasks() []*InspectPoolResponse_RunningTask {
	if m != nil {
		return m.RunningTasks
	}
	return nil
}

func (m *InspectPoolResponse) GetWaitingTasks() []*InspectPoolResponse_WaitingTask {
	if m != nil {
		return m.WaitingTasks
	}
	return nil
}

type InspectPoolResponse_RunningTask struct {
	// Id is the id of the request.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// BotId is the id of the bot running the request.
	BotId string `protobuf:"bytes,2,opt,name=bot_id,json=botId,proto3" json:"bot_id,omitempty"`
	// Priority is the current quotascheduler priority that the task is
	// running at.
	Priority int32 `protobuf:"varint,3,opt,name=priority,proto3" json:"priority,omitempty"`
	// AccountID is the account id of the request.
	AccountId            string   `protobuf:"bytes,4,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InspectPoolResponse_RunningTask) Reset()         { *m = InspectPoolResponse_RunningTask{} }
func (m *InspectPoolResponse_RunningTask) String() string { return proto.CompactTextString(m) }
func (*InspectPoolResponse_RunningTask) ProtoMessage()    {}
func (*InspectPoolResponse_RunningTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_19eb132fa8b54f73, []int{8, 1}
}

func (m *InspectPoolResponse_RunningTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InspectPoolResponse_RunningTask.Unmarshal(m, b)
}
func (m *InspectPoolResponse_RunningTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InspectPoolResponse_RunningTask.Marshal(b, m, deterministic)
}
func (m *InspectPoolResponse_RunningTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InspectPoolResponse_RunningTask.Merge(m, src)
}
func (m *InspectPoolResponse_RunningTask) XXX_Size() int {
	return xxx_messageInfo_InspectPoolResponse_RunningTask.Size(m)
}
func (m *InspectPoolResponse_RunningTask) XXX_DiscardUnknown() {
	xxx_messageInfo_InspectPoolResponse_RunningTask.DiscardUnknown(m)
}

var xxx_messageInfo_InspectPoolResponse_RunningTask proto.InternalMessageInfo

func (m *InspectPoolResponse_RunningTask) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InspectPoolResponse_RunningTask) GetBotId() string {
	if m != nil {
		return m.BotId
	}
	return ""
}

func (m *InspectPoolResponse_RunningTask) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *InspectPoolResponse_RunningTask) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

type InspectPoolResponse_WaitingTask struct {
	// Id is the id of the request.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// AccountID is the account id of the request.
	AccountId            string   `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InspectPoolResponse_WaitingTask) Reset()         { *m = InspectPoolResponse_WaitingTask{} }
func (m *InspectPoolResponse_WaitingTask) String() string { return proto.CompactTextString(m) }
func (*InspectPoolResponse_WaitingTask) ProtoMessage()    {}
func (*InspectPoolResponse_WaitingTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_19eb132fa8b54f73, []int{8, 2}
}

func (m *InspectPoolResponse_WaitingTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InspectPoolResponse_WaitingTask.Unmarshal(m, b)
}
func (m *InspectPoolResponse_WaitingTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InspectPoolResponse_WaitingTask.Marshal(b, m, deterministic)
}
func (m *InspectPoolResponse_WaitingTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InspectPoolResponse_WaitingTask.Merge(m, src)
}
func (m *InspectPoolResponse_WaitingTask) XXX_Size() int {
	return xxx_messageInfo_InspectPoolResponse_WaitingTask.Size(m)
}
func (m *InspectPoolResponse_WaitingTask) XXX_DiscardUnknown() {
	xxx_messageInfo_InspectPoolResponse_WaitingTask.DiscardUnknown(m)
}

var xxx_messageInfo_InspectPoolResponse_WaitingTask proto.InternalMessageInfo

func (m *InspectPoolResponse_WaitingTask) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InspectPoolResponse_WaitingTask) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateSchedulerPoolRequest)(nil), "qscheduler.CreateSchedulerPoolRequest")
	proto.RegisterType((*CreateSchedulerPoolResponse)(nil), "qscheduler.CreateSchedulerPoolResponse")
	proto.RegisterType((*CreateAccountRequest)(nil), "qscheduler.CreateAccountRequest")
	proto.RegisterType((*CreateAccountResponse)(nil), "qscheduler.CreateAccountResponse")
	proto.RegisterType((*ListAccountsRequest)(nil), "qscheduler.ListAccountsRequest")
	proto.RegisterType((*ListAccountsResponse)(nil), "qscheduler.ListAccountsResponse")
	proto.RegisterMapType((map[string]*scheduler.AccountConfig)(nil), "qscheduler.ListAccountsResponse.AccountsEntry")
	proto.RegisterType((*SchedulerPoolConfig)(nil), "qscheduler.SchedulerPoolConfig")
	proto.RegisterType((*InspectPoolRequest)(nil), "qscheduler.InspectPoolRequest")
	proto.RegisterType((*InspectPoolResponse)(nil), "qscheduler.InspectPoolResponse")
	proto.RegisterMapType((map[string]*scheduler.StateProto_Balance)(nil), "qscheduler.InspectPoolResponse.AccountBalancesEntry")
	proto.RegisterType((*InspectPoolResponse_RunningTask)(nil), "qscheduler.InspectPoolResponse.RunningTask")
	proto.RegisterType((*InspectPoolResponse_WaitingTask)(nil), "qscheduler.InspectPoolResponse.WaitingTask")
}

func init() {
	proto.RegisterFile("infra/appengine/qscheduler-swarming/api/qscheduler/v1/admin.proto", fileDescriptor_19eb132fa8b54f73)
}

var fileDescriptor_19eb132fa8b54f73 = []byte{
	// 688 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x4f, 0x4f, 0xdb, 0x4a,
	0x10, 0x97, 0x03, 0xc9, 0x83, 0x09, 0x01, 0xde, 0x06, 0x1e, 0x91, 0x9f, 0x28, 0xa9, 0x0f, 0x2d,
	0x6a, 0x1b, 0xa7, 0x0d, 0x95, 0x5a, 0x55, 0xbd, 0x00, 0xea, 0x21, 0x55, 0x55, 0x05, 0x43, 0xdb,
	0x63, 0xb4, 0x89, 0x97, 0x74, 0x85, 0xb3, 0x6b, 0xbc, 0x6b, 0x22, 0x6e, 0xfd, 0x38, 0xbd, 0xf7,
	0xd4, 0x4f, 0xd1, 0xaf, 0x54, 0xd9, 0xbb, 0x49, 0xd6, 0xc1, 0xc1, 0xdc, 0xe2, 0x99, 0xdf, 0xcc,
	0xfc, 0xe6, 0xcf, 0xfe, 0x02, 0xc7, 0x94, 0x5d, 0x46, 0xb8, 0x8d, 0xc3, 0x90, 0xb0, 0x11, 0x65,
	0xa4, 0x7d, 0x2d, 0x86, 0xdf, 0x89, 0x1f, 0x07, 0x24, 0x6a, 0x89, 0x09, 0x8e, 0xc6, 0x94, 0x8d,
	0xda, 0x38, 0xa4, 0x86, 0xbd, 0x7d, 0xf3, 0xaa, 0x8d, 0xfd, 0x31, 0x65, 0x6e, 0x18, 0x71, 0xc9,
	0x11, 0xcc, 0x5d, 0x76, 0x4b, 0xa5, 0x33, 0xc0, 0xd7, 0x22, 0xa0, 0x83, 0xf6, 0xfc, 0x7b, 0xc8,
	0xd9, 0x25, 0x1d, 0xa9, 0x50, 0xfb, 0x45, 0x21, 0x5c, 0x48, 0x2c, 0x89, 0x42, 0x3b, 0x0c, 0xec,
	0xd3, 0x88, 0x60, 0x49, 0xce, 0xa7, 0xee, 0x1e, 0xe7, 0x81, 0x47, 0xae, 0x63, 0x22, 0x24, 0xda,
	0x83, 0x7f, 0x42, 0xce, 0x83, 0x3e, 0xf5, 0x1b, 0x56, 0xd3, 0x3a, 0x5c, 0xf7, 0x2a, 0xc9, 0x67,
	0xd7, 0x47, 0x6f, 0xa0, 0xa2, 0x8a, 0x36, 0x4a, 0x4d, 0xeb, 0xb0, 0xda, 0x39, 0x70, 0xe7, 0xf5,
	0xdc, 0x4c, 0xaa, 0xd3, 0x14, 0xe6, 0x69, 0xb8, 0xb3, 0x0f, 0xff, 0xe7, 0xd6, 0x13, 0x21, 0x67,
	0x82, 0x38, 0x3f, 0x2c, 0xd8, 0x51, 0xfe, 0xe3, 0xe1, 0x90, 0xc7, 0x4c, 0x16, 0x32, 0xd9, 0x07,
	0xc0, 0x0a, 0x9a, 0xf8, 0x4a, 0xa9, 0x6f, 0x5d, 0x5b, 0xba, 0x3e, 0x7a, 0x39, 0x23, 0xba, 0x92,
	0x12, 0x6d, 0xb8, 0x73, 0x9e, 0xba, 0xc4, 0x02, 0xc3, 0x3d, 0xd8, 0x5d, 0x60, 0xa0, 0xb9, 0xb9,
	0x50, 0xff, 0x44, 0x85, 0xd4, 0x66, 0x51, 0xc4, 0xcc, 0xf9, 0x6d, 0xc1, 0x4e, 0x36, 0x40, 0x25,
	0x42, 0x1f, 0x61, 0x4d, 0x13, 0x14, 0x0d, 0xab, 0xb9, 0x72, 0x58, 0xed, 0xb8, 0xe6, 0xf8, 0xf2,
	0x62, 0xa6, 0x5c, 0xc5, 0x07, 0x26, 0xa3, 0x5b, 0x6f, 0x16, 0x6f, 0x7f, 0x81, 0x5a, 0xc6, 0x85,
	0xb6, 0x61, 0xe5, 0x8a, 0xdc, 0x6a, 0x2a, 0xc9, 0x4f, 0xe4, 0x42, 0xf9, 0x06, 0x07, 0x31, 0xd1,
	0xab, 0x5a, 0x3e, 0x01, 0x05, 0x7b, 0x57, 0x7a, 0x6b, 0x39, 0x2d, 0xa8, 0xe7, 0x6c, 0x11, 0xfd,
	0x07, 0x95, 0x00, 0x0f, 0x48, 0xa0, 0x78, 0xaf, 0x7b, 0xfa, 0xcb, 0x69, 0x01, 0xea, 0x32, 0x11,
	0x92, 0xa1, 0x7c, 0xc8, 0xf5, 0x38, 0x3f, 0xcb, 0x50, 0xcf, 0xe0, 0xf5, 0x60, 0x9e, 0xc1, 0xbf,
	0x2c, 0x1e, 0xf7, 0x27, 0x98, 0x4a, 0xca, 0x46, 0x7d, 0x89, 0xc5, 0x95, 0x48, 0x43, 0xcb, 0xde,
	0x16, 0x8b, 0xc7, 0xdf, 0x94, 0xfd, 0x22, 0x31, 0x4f, 0xb1, 0x51, 0xcc, 0xd8, 0x1c, 0x5b, 0x9a,
	0x61, 0x3d, 0x65, 0x57, 0x58, 0x07, 0x6a, 0x09, 0x96, 0xfa, 0x01, 0xe9, 0x0f, 0xb8, 0x14, 0xe9,
	0x2d, 0x94, 0xbd, 0x2a, 0x8b, 0xc7, 0x5d, 0x3f, 0x20, 0x27, 0x5c, 0x0a, 0xd4, 0x87, 0xed, 0xe9,
	0x1d, 0x0d, 0x70, 0x80, 0xd9, 0x90, 0x88, 0xc6, 0x6a, 0xba, 0x9c, 0xd7, 0xe6, 0x72, 0x72, 0x68,
	0x4f, 0xa7, 0x78, 0xa2, 0xc3, 0xd4, 0x8a, 0xb6, 0x70, 0xd6, 0x8a, 0x7a, 0x50, 0xcb, 0x92, 0x2d,
	0xa7, 0xd9, 0x9f, 0x17, 0x65, 0x37, 0x3a, 0xf1, 0x36, 0x22, 0xb3, 0xad, 0x1e, 0xd4, 0xb2, 0xa3,
	0xaa, 0x3c, 0x2c, 0xa3, 0x31, 0x47, 0x6f, 0x63, 0x62, 0x0c, 0xd5, 0xc6, 0xb0, 0x93, 0xd7, 0x4c,
	0xce, 0x51, 0x1d, 0x65, 0x8f, 0x6a, 0xdf, 0x38, 0xaa, 0xf3, 0x44, 0x5e, 0x7a, 0x89, 0xba, 0xb8,
	0x3a, 0x8b, 0x71, 0x59, 0x36, 0x87, 0xaa, 0xd1, 0x11, 0xda, 0x84, 0xd2, 0xec, 0x3c, 0x4a, 0xd4,
	0x47, 0xbb, 0x50, 0x19, 0x70, 0xe3, 0x29, 0x97, 0x07, 0x3c, 0x79, 0xc6, 0x36, 0xac, 0x85, 0x11,
	0xe5, 0x11, 0x95, 0xb7, 0x7a, 0x79, 0xb3, 0xef, 0x05, 0x05, 0x58, 0x5d, 0x50, 0x00, 0xfb, 0x3d,
	0x54, 0x8d, 0x86, 0xef, 0x14, 0xbc, 0x5f, 0x3f, 0x3a, 0x7f, 0x2c, 0xd8, 0x3a, 0x9b, 0x3d, 0x85,
	0xe3, 0x44, 0xa2, 0xd1, 0x25, 0xd4, 0x73, 0x34, 0x0c, 0x3d, 0x31, 0xe7, 0xbe, 0x5c, 0x54, 0xed,
	0xa7, 0x85, 0x38, 0xfd, 0x1c, 0x2e, 0xa0, 0x96, 0x51, 0x22, 0xd4, 0xbc, 0x1b, 0x99, 0x95, 0x49,
	0xfb, 0xf1, 0x3d, 0x08, 0x95, 0xb5, 0xf3, 0xcb, 0x82, 0xcd, 0x79, 0x47, 0x5f, 0x29, 0x99, 0xa0,
	0x33, 0xd8, 0x30, 0x45, 0x07, 0x1d, 0x2c, 0x97, 0x23, 0x55, 0xa6, 0x59, 0xa4, 0x57, 0xe8, 0x33,
	0x54, 0x8d, 0xd3, 0x43, 0x8f, 0x96, 0xde, 0xa4, 0x4a, 0x78, 0x50, 0x70, 0xb3, 0x83, 0x4a, 0xfa,
	0x77, 0x75, 0xf4, 0x37, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x19, 0x72, 0x8e, 0x5c, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QSchedulerAdminClient is the client API for QSchedulerAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QSchedulerAdminClient interface {
	// CreateSchedulerPool creates a scheduler, with the given configuration
	// options.
	CreateSchedulerPool(ctx context.Context, in *CreateSchedulerPoolRequest, opts ...grpc.CallOption) (*CreateSchedulerPoolResponse, error)
	// CreateAccount creates a quota account within a scheduler, with the
	// given configuration options.
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
}
type qSchedulerAdminPRPCClient struct {
	client *prpc.Client
}

func NewQSchedulerAdminPRPCClient(client *prpc.Client) QSchedulerAdminClient {
	return &qSchedulerAdminPRPCClient{client}
}

func (c *qSchedulerAdminPRPCClient) CreateSchedulerPool(ctx context.Context, in *CreateSchedulerPoolRequest, opts ...grpc.CallOption) (*CreateSchedulerPoolResponse, error) {
	out := new(CreateSchedulerPoolResponse)
	err := c.client.Call(ctx, "qscheduler.QSchedulerAdmin", "CreateSchedulerPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qSchedulerAdminPRPCClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.client.Call(ctx, "qscheduler.QSchedulerAdmin", "CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type qSchedulerAdminClient struct {
	cc *grpc.ClientConn
}

func NewQSchedulerAdminClient(cc *grpc.ClientConn) QSchedulerAdminClient {
	return &qSchedulerAdminClient{cc}
}

func (c *qSchedulerAdminClient) CreateSchedulerPool(ctx context.Context, in *CreateSchedulerPoolRequest, opts ...grpc.CallOption) (*CreateSchedulerPoolResponse, error) {
	out := new(CreateSchedulerPoolResponse)
	err := c.cc.Invoke(ctx, "/qscheduler.QSchedulerAdmin/CreateSchedulerPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qSchedulerAdminClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/qscheduler.QSchedulerAdmin/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QSchedulerAdminServer is the server API for QSchedulerAdmin service.
type QSchedulerAdminServer interface {
	// CreateSchedulerPool creates a scheduler, with the given configuration
	// options.
	CreateSchedulerPool(context.Context, *CreateSchedulerPoolRequest) (*CreateSchedulerPoolResponse, error)
	// CreateAccount creates a quota account within a scheduler, with the
	// given configuration options.
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
}

func RegisterQSchedulerAdminServer(s prpc.Registrar, srv QSchedulerAdminServer) {
	s.RegisterService(&_QSchedulerAdmin_serviceDesc, srv)
}

func _QSchedulerAdmin_CreateSchedulerPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSchedulerPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QSchedulerAdminServer).CreateSchedulerPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qscheduler.QSchedulerAdmin/CreateSchedulerPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QSchedulerAdminServer).CreateSchedulerPool(ctx, req.(*CreateSchedulerPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QSchedulerAdmin_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QSchedulerAdminServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qscheduler.QSchedulerAdmin/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QSchedulerAdminServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QSchedulerAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "qscheduler.QSchedulerAdmin",
	HandlerType: (*QSchedulerAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSchedulerPool",
			Handler:    _QSchedulerAdmin_CreateSchedulerPool_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _QSchedulerAdmin_CreateAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infra/appengine/qscheduler-swarming/api/qscheduler/v1/admin.proto",
}

// QSchedulerViewClient is the client API for QSchedulerView service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QSchedulerViewClient interface {
	// ListAccounts returns the set of accounts for a given scheduler.
	ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error)
	// InspectPool returns a description of the state of a scheduler, for debugging
	// or diagnostic purposes.
	InspectPool(ctx context.Context, in *InspectPoolRequest, opts ...grpc.CallOption) (*InspectPoolResponse, error)
}
type qSchedulerViewPRPCClient struct {
	client *prpc.Client
}

func NewQSchedulerViewPRPCClient(client *prpc.Client) QSchedulerViewClient {
	return &qSchedulerViewPRPCClient{client}
}

func (c *qSchedulerViewPRPCClient) ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error) {
	out := new(ListAccountsResponse)
	err := c.client.Call(ctx, "qscheduler.QSchedulerView", "ListAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qSchedulerViewPRPCClient) InspectPool(ctx context.Context, in *InspectPoolRequest, opts ...grpc.CallOption) (*InspectPoolResponse, error) {
	out := new(InspectPoolResponse)
	err := c.client.Call(ctx, "qscheduler.QSchedulerView", "InspectPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type qSchedulerViewClient struct {
	cc *grpc.ClientConn
}

func NewQSchedulerViewClient(cc *grpc.ClientConn) QSchedulerViewClient {
	return &qSchedulerViewClient{cc}
}

func (c *qSchedulerViewClient) ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error) {
	out := new(ListAccountsResponse)
	err := c.cc.Invoke(ctx, "/qscheduler.QSchedulerView/ListAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qSchedulerViewClient) InspectPool(ctx context.Context, in *InspectPoolRequest, opts ...grpc.CallOption) (*InspectPoolResponse, error) {
	out := new(InspectPoolResponse)
	err := c.cc.Invoke(ctx, "/qscheduler.QSchedulerView/InspectPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QSchedulerViewServer is the server API for QSchedulerView service.
type QSchedulerViewServer interface {
	// ListAccounts returns the set of accounts for a given scheduler.
	ListAccounts(context.Context, *ListAccountsRequest) (*ListAccountsResponse, error)
	// InspectPool returns a description of the state of a scheduler, for debugging
	// or diagnostic purposes.
	InspectPool(context.Context, *InspectPoolRequest) (*InspectPoolResponse, error)
}

func RegisterQSchedulerViewServer(s prpc.Registrar, srv QSchedulerViewServer) {
	s.RegisterService(&_QSchedulerView_serviceDesc, srv)
}

func _QSchedulerView_ListAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QSchedulerViewServer).ListAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qscheduler.QSchedulerView/ListAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QSchedulerViewServer).ListAccounts(ctx, req.(*ListAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QSchedulerView_InspectPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InspectPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QSchedulerViewServer).InspectPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qscheduler.QSchedulerView/InspectPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QSchedulerViewServer).InspectPool(ctx, req.(*InspectPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QSchedulerView_serviceDesc = grpc.ServiceDesc{
	ServiceName: "qscheduler.QSchedulerView",
	HandlerType: (*QSchedulerViewServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAccounts",
			Handler:    _QSchedulerView_ListAccounts_Handler,
		},
		{
			MethodName: "InspectPool",
			Handler:    _QSchedulerView_InspectPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infra/appengine/qscheduler-swarming/api/qscheduler/v1/admin.proto",
}
