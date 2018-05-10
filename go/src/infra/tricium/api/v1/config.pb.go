// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/tricium/api/v1/config.proto

/*
Package tricium is a generated protocol buffer package.

It is generated from these files:
	infra/tricium/api/v1/config.proto
	infra/tricium/api/v1/data.proto
	infra/tricium/api/v1/function.proto
	infra/tricium/api/v1/platform.proto
	infra/tricium/api/v1/tricium.proto

It has these top-level messages:
	ServiceConfig
	ProjectConfig
	RepoDetails
	GerritProject
	GitRepo
	Acl
	Selection
	Config
	Data
	Function
	ConfigDef
	Impl
	Recipe
	Property
	Cmd
	CipdPackage
	Platform
	AnalyzeRequest
	GerritRevision
	GitCommit
	AnalyzeResponse
	ProgressRequest
	ProgressResponse
	FunctionProgress
	ProjectProgressRequest
	ProjectProgressResponse
	RunProgress
	ResultsRequest
	ResultsResponse
	FeedbackRequest
	FeedbackResponse
	ReportNotUsefulRequest
	ReportNotUsefulResponse
*/
package tricium

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Roles relevant to Tricium.
type Acl_Role int32

const (
	// Can read progress/results.
	Acl_READER Acl_Role = 0
	// Can request analysis.
	Acl_REQUESTER Acl_Role = 1
)

var Acl_Role_name = map[int32]string{
	0: "READER",
	1: "REQUESTER",
}
var Acl_Role_value = map[string]int32{
	"READER":    0,
	"REQUESTER": 1,
}

func (x Acl_Role) String() string {
	return proto.EnumName(Acl_Role_name, int32(x))
}
func (Acl_Role) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

// Tricium service configuration.
//
// Listing supported platforms and analyzers shared between projects connected
// to Tricium.
type ServiceConfig struct {
	// Supported platforms.
	Platforms []*Platform_Details `protobuf:"bytes,1,rep,name=platforms" json:"platforms,omitempty"`
	// Supported data types.
	DataDetails []*Data_TypeDetails `protobuf:"bytes,2,rep,name=data_details,json=dataDetails" json:"data_details,omitempty"`
	// List of shared functions.
	Functions []*Function `protobuf:"bytes,3,rep,name=functions" json:"functions,omitempty"`
	// Base recipe command used for workers implemented as recipes.
	//
	// Specific recipe details for the worker will be added as flags at the
	// end of the argument list.
	RecipeCmd *Cmd `protobuf:"bytes,4,opt,name=recipe_cmd,json=recipeCmd" json:"recipe_cmd,omitempty"`
	// Base recipe packages used for workers implemented as recipes.
	//
	// These packages will be adjusted for the platform in question, by appending
	// platform name details to the end of the package name.
	RecipePackages []*CipdPackage `protobuf:"bytes,5,rep,name=recipe_packages,json=recipePackages" json:"recipe_packages,omitempty"`
	// Swarming server to use for this service instance.
	//
	// This should be a full URL with no trailing slash.
	SwarmingServer string `protobuf:"bytes,6,opt,name=swarming_server,json=swarmingServer" json:"swarming_server,omitempty"`
	// Isolate server to use for this service instance.
	//
	// This should be a full URL with no trailing slash.
	IsolateServer string `protobuf:"bytes,7,opt,name=isolate_server,json=isolateServer" json:"isolate_server,omitempty"`
}

func (m *ServiceConfig) Reset()                    { *m = ServiceConfig{} }
func (m *ServiceConfig) String() string            { return proto.CompactTextString(m) }
func (*ServiceConfig) ProtoMessage()               {}
func (*ServiceConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServiceConfig) GetPlatforms() []*Platform_Details {
	if m != nil {
		return m.Platforms
	}
	return nil
}

func (m *ServiceConfig) GetDataDetails() []*Data_TypeDetails {
	if m != nil {
		return m.DataDetails
	}
	return nil
}

func (m *ServiceConfig) GetFunctions() []*Function {
	if m != nil {
		return m.Functions
	}
	return nil
}

func (m *ServiceConfig) GetRecipeCmd() *Cmd {
	if m != nil {
		return m.RecipeCmd
	}
	return nil
}

func (m *ServiceConfig) GetRecipePackages() []*CipdPackage {
	if m != nil {
		return m.RecipePackages
	}
	return nil
}

func (m *ServiceConfig) GetSwarmingServer() string {
	if m != nil {
		return m.SwarmingServer
	}
	return ""
}

func (m *ServiceConfig) GetIsolateServer() string {
	if m != nil {
		return m.IsolateServer
	}
	return ""
}

// Tricium project configuration.
//
// Specifies details needed to connect a project to Tricium.
// Adds project-specific functions and selects shared function
// implementations.
type ProjectConfig struct {
	// Access control rules for the project.
	Acls []*Acl `protobuf:"bytes,2,rep,name=acls" json:"acls,omitempty"`
	// Project-specific function details.
	//
	// This includes project-specific analyzer implementations and full
	// project-specific analyzer specifications.
	Functions []*Function `protobuf:"bytes,3,rep,name=functions" json:"functions,omitempty"`
	// Selection of function implementations to run for this project.
	Selections []*Selection `protobuf:"bytes,4,rep,name=selections" json:"selections,omitempty"`
	// Repositories, including Git and Gerrit details.
	Repos []*RepoDetails `protobuf:"bytes,5,rep,name=repos" json:"repos,omitempty"`
	// General service account for this project.
	// Used for any service interaction, with the exception of swarming.
	ServiceAccount string `protobuf:"bytes,6,opt,name=service_account,json=serviceAccount" json:"service_account,omitempty"`
	// Project-specific swarming service account.
	SwarmingServiceAccount string `protobuf:"bytes,7,opt,name=swarming_service_account,json=swarmingServiceAccount" json:"swarming_service_account,omitempty"`
}

func (m *ProjectConfig) Reset()                    { *m = ProjectConfig{} }
func (m *ProjectConfig) String() string            { return proto.CompactTextString(m) }
func (*ProjectConfig) ProtoMessage()               {}
func (*ProjectConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ProjectConfig) GetAcls() []*Acl {
	if m != nil {
		return m.Acls
	}
	return nil
}

func (m *ProjectConfig) GetFunctions() []*Function {
	if m != nil {
		return m.Functions
	}
	return nil
}

func (m *ProjectConfig) GetSelections() []*Selection {
	if m != nil {
		return m.Selections
	}
	return nil
}

func (m *ProjectConfig) GetRepos() []*RepoDetails {
	if m != nil {
		return m.Repos
	}
	return nil
}

func (m *ProjectConfig) GetServiceAccount() string {
	if m != nil {
		return m.ServiceAccount
	}
	return ""
}

func (m *ProjectConfig) GetSwarmingServiceAccount() string {
	if m != nil {
		return m.SwarmingServiceAccount
	}
	return ""
}

// Repository details for a project.
// DEPRECATED, see https://crbug.com/824558
type RepoDetails struct {
	// Could be renamed to kind when the above kind is removed.
	//
	// Types that are valid to be assigned to Source:
	//	*RepoDetails_GerritProject
	//	*RepoDetails_GitRepo
	Source isRepoDetails_Source `protobuf_oneof:"source"`
	// Whether to disable reporting results back.
	DisableReporting bool `protobuf:"varint,3,opt,name=disable_reporting,json=disableReporting" json:"disable_reporting,omitempty"`
	// Whitelisted groups.
	//
	// The owner of a change will be checked for membership of a whitelisted
	// group. Absence of this field means all groups are whitelisted.
	//
	// Group names must be known to the Chrome infra auth service,
	// https://chrome-infra-auth.appspot.com. Contact a Chromium trooper
	// if you need to add or modify a group: g.co/bugatrooper.
	WhitelistedGroup []string `protobuf:"bytes,4,rep,name=whitelisted_group,json=whitelistedGroup" json:"whitelisted_group,omitempty"`
}

func (m *RepoDetails) Reset()                    { *m = RepoDetails{} }
func (m *RepoDetails) String() string            { return proto.CompactTextString(m) }
func (*RepoDetails) ProtoMessage()               {}
func (*RepoDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isRepoDetails_Source interface {
	isRepoDetails_Source()
}

type RepoDetails_GerritProject struct {
	GerritProject *GerritProject `protobuf:"bytes,1,opt,name=gerrit_project,json=gerritProject,oneof"`
}
type RepoDetails_GitRepo struct {
	GitRepo *GitRepo `protobuf:"bytes,2,opt,name=git_repo,json=gitRepo,oneof"`
}

func (*RepoDetails_GerritProject) isRepoDetails_Source() {}
func (*RepoDetails_GitRepo) isRepoDetails_Source()       {}

func (m *RepoDetails) GetSource() isRepoDetails_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *RepoDetails) GetGerritProject() *GerritProject {
	if x, ok := m.GetSource().(*RepoDetails_GerritProject); ok {
		return x.GerritProject
	}
	return nil
}

func (m *RepoDetails) GetGitRepo() *GitRepo {
	if x, ok := m.GetSource().(*RepoDetails_GitRepo); ok {
		return x.GitRepo
	}
	return nil
}

func (m *RepoDetails) GetDisableReporting() bool {
	if m != nil {
		return m.DisableReporting
	}
	return false
}

func (m *RepoDetails) GetWhitelistedGroup() []string {
	if m != nil {
		return m.WhitelistedGroup
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RepoDetails) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RepoDetails_OneofMarshaler, _RepoDetails_OneofUnmarshaler, _RepoDetails_OneofSizer, []interface{}{
		(*RepoDetails_GerritProject)(nil),
		(*RepoDetails_GitRepo)(nil),
	}
}

func _RepoDetails_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RepoDetails)
	// source
	switch x := m.Source.(type) {
	case *RepoDetails_GerritProject:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GerritProject); err != nil {
			return err
		}
	case *RepoDetails_GitRepo:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GitRepo); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RepoDetails.Source has unexpected type %T", x)
	}
	return nil
}

func _RepoDetails_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RepoDetails)
	switch tag {
	case 1: // source.gerrit_project
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GerritProject)
		err := b.DecodeMessage(msg)
		m.Source = &RepoDetails_GerritProject{msg}
		return true, err
	case 2: // source.git_repo
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GitRepo)
		err := b.DecodeMessage(msg)
		m.Source = &RepoDetails_GitRepo{msg}
		return true, err
	default:
		return false, nil
	}
}

func _RepoDetails_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RepoDetails)
	// source
	switch x := m.Source.(type) {
	case *RepoDetails_GerritProject:
		s := proto.Size(x.GerritProject)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RepoDetails_GitRepo:
		s := proto.Size(x.GitRepo)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Specifies a Gerrit project and its corresponding git repo.
type GerritProject struct {
	// The Gerrit host to connect to.
	//
	// Value must not include the schema part; it will be assumed to be "https".
	Host string `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	// Gerrit project name.
	Project string `protobuf:"bytes,2,opt,name=project" json:"project,omitempty"`
	// Full URL for the corresponding git repo.
	GitUrl string `protobuf:"bytes,3,opt,name=git_url,json=gitUrl" json:"git_url,omitempty"`
}

func (m *GerritProject) Reset()                    { *m = GerritProject{} }
func (m *GerritProject) String() string            { return proto.CompactTextString(m) }
func (*GerritProject) ProtoMessage()               {}
func (*GerritProject) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GerritProject) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *GerritProject) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *GerritProject) GetGitUrl() string {
	if m != nil {
		return m.GitUrl
	}
	return ""
}

type GitRepo struct {
	// Full repository url, including schema.
	Url string `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
}

func (m *GitRepo) Reset()                    { *m = GitRepo{} }
func (m *GitRepo) String() string            { return proto.CompactTextString(m) }
func (*GitRepo) ProtoMessage()               {}
func (*GitRepo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GitRepo) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// Access control rules.
type Acl struct {
	// Role of a group or identity.
	Role Acl_Role `protobuf:"varint,1,opt,name=role,enum=tricium.Acl_Role" json:"role,omitempty"`
	// Name of group, as defined in the auth service. Specify either group or
	// identity, not both.
	Group string `protobuf:"bytes,2,opt,name=group" json:"group,omitempty"`
	// Identity, as defined by the auth service. Can be either an email address
	// or an identity string, for instance, "anonymous:anonymous" for anonymous
	// users. Specify either group or identity, not both.
	Identity string `protobuf:"bytes,3,opt,name=identity" json:"identity,omitempty"`
}

func (m *Acl) Reset()                    { *m = Acl{} }
func (m *Acl) String() string            { return proto.CompactTextString(m) }
func (*Acl) ProtoMessage()               {}
func (*Acl) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Acl) GetRole() Acl_Role {
	if m != nil {
		return m.Role
	}
	return Acl_READER
}

func (m *Acl) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *Acl) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

// Selection of function implementations to run for a project.
type Selection struct {
	// Name of function to run.
	Function string `protobuf:"bytes,1,opt,name=function" json:"function,omitempty"`
	// Name of platform to retrieve results from.
	Platform Platform_Name `protobuf:"varint,2,opt,name=platform,enum=tricium.Platform_Name" json:"platform,omitempty"`
	// Function configuration to use on this platform.
	Configs []*Config `protobuf:"bytes,3,rep,name=configs" json:"configs,omitempty"`
}

func (m *Selection) Reset()                    { *m = Selection{} }
func (m *Selection) String() string            { return proto.CompactTextString(m) }
func (*Selection) ProtoMessage()               {}
func (*Selection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Selection) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

func (m *Selection) GetPlatform() Platform_Name {
	if m != nil {
		return m.Platform
	}
	return Platform_ANY
}

func (m *Selection) GetConfigs() []*Config {
	if m != nil {
		return m.Configs
	}
	return nil
}

// Function configuration used when selecting a function implementation.
type Config struct {
	// Name of the configuration option.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Value of the configuration.
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Config) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Config) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*ServiceConfig)(nil), "tricium.ServiceConfig")
	proto.RegisterType((*ProjectConfig)(nil), "tricium.ProjectConfig")
	proto.RegisterType((*RepoDetails)(nil), "tricium.RepoDetails")
	proto.RegisterType((*GerritProject)(nil), "tricium.GerritProject")
	proto.RegisterType((*GitRepo)(nil), "tricium.GitRepo")
	proto.RegisterType((*Acl)(nil), "tricium.Acl")
	proto.RegisterType((*Selection)(nil), "tricium.Selection")
	proto.RegisterType((*Config)(nil), "tricium.Config")
	proto.RegisterEnum("tricium.Acl_Role", Acl_Role_name, Acl_Role_value)
}

func init() { proto.RegisterFile("infra/tricium/api/v1/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 719 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcb, 0x6e, 0xdb, 0x38,
	0x14, 0x8d, 0x6c, 0xc7, 0xb6, 0xae, 0x63, 0xc7, 0x21, 0x82, 0x8c, 0x26, 0xb3, 0x18, 0x47, 0x83,
	0x60, 0x3c, 0x13, 0x8c, 0x8d, 0x71, 0x17, 0xed, 0xa2, 0x45, 0xe1, 0x26, 0x6e, 0xb2, 0x2a, 0x52,
	0x3a, 0xe9, 0x56, 0x60, 0x24, 0x5a, 0x61, 0xab, 0x17, 0x28, 0xda, 0x41, 0x96, 0xdd, 0xf4, 0x4f,
	0xfa, 0x61, 0xfd, 0x86, 0xfe, 0x40, 0x21, 0x3e, 0x24, 0xa5, 0x0d, 0x0a, 0x74, 0xc7, 0x7b, 0xcf,
	0x39, 0xbc, 0xf7, 0x1e, 0x3e, 0xe0, 0x88, 0x25, 0x2b, 0x4e, 0xa6, 0x82, 0x33, 0x9f, 0xad, 0xe3,
	0x29, 0xc9, 0xd8, 0x74, 0xf3, 0xff, 0xd4, 0x4f, 0x93, 0x15, 0x0b, 0x27, 0x19, 0x4f, 0x45, 0x8a,
	0x3a, 0x1a, 0x3c, 0xfc, 0xf3, 0x51, 0x6e, 0x40, 0x04, 0x51, 0xcc, 0xc3, 0xbf, 0x1e, 0x25, 0xac,
	0xd6, 0x89, 0x2f, 0x58, 0x9a, 0xfc, 0x94, 0x94, 0x45, 0x44, 0xac, 0x52, 0x1e, 0x2b, 0x92, 0xfb,
	0xb5, 0x01, 0xfd, 0x25, 0xe5, 0x1b, 0xe6, 0xd3, 0x53, 0xd9, 0x0b, 0x7a, 0x0a, 0xb6, 0xe1, 0xe4,
	0x8e, 0x35, 0x6a, 0x8e, 0x7b, 0xb3, 0xdf, 0x27, 0x7a, 0x93, 0xc9, 0xa5, 0x51, 0x9f, 0x51, 0x41,
	0x58, 0x94, 0xe3, 0x8a, 0x8b, 0x9e, 0xc3, 0x4e, 0xd1, 0xa2, 0x17, 0x28, 0xc8, 0x69, 0x7c, 0xa7,
	0x3d, 0x2b, 0xfa, 0xbf, 0xba, 0xcf, 0xa8, 0xd1, 0xf6, 0x0a, 0xba, 0x0e, 0xd0, 0x14, 0x6c, 0xd3,
	0x7f, 0xee, 0x34, 0xa5, 0x74, 0xaf, 0x94, 0xbe, 0xd6, 0x08, 0xae, 0x38, 0xe8, 0x04, 0x80, 0x53,
	0x9f, 0x65, 0xd4, 0xf3, 0xe3, 0xc0, 0x69, 0x8d, 0xac, 0x71, 0x6f, 0xb6, 0x53, 0x2a, 0x4e, 0xe3,
	0x00, 0xdb, 0x0a, 0x3f, 0x8d, 0x03, 0xf4, 0x02, 0x76, 0x35, 0x39, 0x23, 0xfe, 0x07, 0x12, 0xd2,
	0xdc, 0xd9, 0x96, 0x35, 0xf6, 0x2b, 0x05, 0xcb, 0x82, 0x4b, 0x05, 0xe2, 0x81, 0x22, 0xeb, 0x30,
	0x47, 0x7f, 0xc3, 0x6e, 0x7e, 0x47, 0x78, 0xcc, 0x92, 0xd0, 0xcb, 0x29, 0xdf, 0x50, 0xee, 0xb4,
	0x47, 0xd6, 0xd8, 0xc6, 0x03, 0x93, 0x5e, 0xca, 0x2c, 0x3a, 0x86, 0x01, 0xcb, 0xd3, 0x88, 0x08,
	0x6a, 0x78, 0x1d, 0xc9, 0xeb, 0xeb, 0xac, 0xa2, 0xb9, 0x9f, 0x1b, 0xd0, 0xbf, 0xe4, 0xe9, 0x7b,
	0xea, 0x0b, 0xed, 0xfa, 0x08, 0x5a, 0xc4, 0x2f, 0x4d, 0xab, 0xe6, 0x98, 0xfb, 0x11, 0x96, 0xc8,
	0xaf, 0x1b, 0x34, 0x03, 0xc8, 0x69, 0x44, 0xb5, 0xa2, 0x25, 0x15, 0xa8, 0x54, 0x2c, 0x0d, 0x84,
	0x6b, 0x2c, 0xf4, 0x2f, 0x6c, 0x73, 0x9a, 0xa5, 0x3f, 0xba, 0x83, 0x69, 0x96, 0x9a, 0x73, 0x53,
	0x14, 0x69, 0x8a, 0xba, 0x39, 0x1e, 0xf1, 0xfd, 0x74, 0x9d, 0x88, 0xd2, 0x14, 0x95, 0x9e, 0xab,
	0x2c, 0x7a, 0x06, 0xce, 0x03, 0xf7, 0xea, 0x0a, 0x65, 0xcf, 0x41, 0xdd, 0xc6, 0x4a, 0xe9, 0x7e,
	0xb1, 0xa0, 0x57, 0xab, 0x8c, 0x5e, 0xc2, 0x20, 0xa4, 0x9c, 0x33, 0xe1, 0x65, 0xca, 0x3d, 0xc7,
	0x92, 0xe7, 0x7e, 0x50, 0xf6, 0x79, 0x2e, 0x61, 0xed, 0xed, 0xc5, 0x16, 0xee, 0x87, 0xf5, 0x04,
	0xfa, 0x0f, 0xba, 0x21, 0x13, 0x5e, 0x31, 0x80, 0xd3, 0x90, 0xd2, 0x61, 0x25, 0x65, 0xa2, 0xa8,
	0x75, 0xb1, 0x85, 0x3b, 0xa1, 0x5a, 0xa2, 0x13, 0xd8, 0x0b, 0x58, 0x4e, 0x6e, 0x22, 0x2a, 0x25,
	0x5c, 0xb0, 0x24, 0x74, 0x9a, 0x23, 0x6b, 0xdc, 0xc5, 0x43, 0x0d, 0x60, 0x93, 0x2f, 0xc8, 0x77,
	0xb7, 0x4c, 0xd0, 0x88, 0xe5, 0x82, 0x06, 0x5e, 0xc8, 0xd3, 0x75, 0x26, 0x6d, 0xb7, 0xf1, 0xb0,
	0x06, 0x9c, 0x17, 0xf9, 0x57, 0x5d, 0x68, 0xe7, 0xe9, 0x9a, 0xfb, 0xd4, 0x7d, 0x07, 0xfd, 0x07,
	0x4d, 0x23, 0x04, 0xad, 0xdb, 0x34, 0x57, 0xa3, 0xd9, 0x58, 0xae, 0x91, 0x03, 0x1d, 0x33, 0x71,
	0x43, 0xa6, 0x4d, 0x88, 0x7e, 0x83, 0xa2, 0x5b, 0x6f, 0xcd, 0x23, 0xd9, 0x98, 0x8d, 0xdb, 0x21,
	0x13, 0xd7, 0x3c, 0x72, 0xff, 0x80, 0x8e, 0x9e, 0x08, 0x0d, 0xa1, 0x59, 0xe1, 0xc5, 0xd2, 0xfd,
	0x68, 0x41, 0x73, 0xee, 0x47, 0xe8, 0x18, 0x5a, 0x3c, 0x8d, 0xa8, 0xac, 0x35, 0xa8, 0xdd, 0xa7,
	0xb9, 0x1f, 0x4d, 0x70, 0x1a, 0x51, 0x2c, 0x61, 0xb4, 0x0f, 0xdb, 0x6a, 0x1c, 0x55, 0x5c, 0x05,
	0xe8, 0x10, 0xba, 0x2c, 0xa0, 0x89, 0x60, 0xe2, 0x5e, 0xef, 0x5d, 0xc6, 0xee, 0x11, 0xb4, 0x0a,
	0x3d, 0x02, 0x68, 0xe3, 0xc5, 0xfc, 0x6c, 0x81, 0x87, 0x5b, 0xa8, 0x0f, 0x36, 0x5e, 0xbc, 0xbd,
	0x5e, 0x2c, 0xaf, 0x16, 0x78, 0x68, 0xb9, 0x9f, 0x2c, 0xb0, 0xcb, 0x5b, 0x58, 0x6c, 0x66, 0xae,
	0xae, 0x9e, 0xbc, 0x8c, 0xd1, 0x0c, 0xba, 0xe6, 0x9b, 0x91, 0x1d, 0x0c, 0x6a, 0x07, 0x5e, 0xfe,
	0x48, 0x6f, 0x48, 0x4c, 0x71, 0xc9, 0x43, 0xff, 0x40, 0x47, 0x7d, 0xae, 0xe6, 0xb1, 0xec, 0x56,
	0x2f, 0x5d, 0xe6, 0xb1, 0xc1, 0xdd, 0x19, 0xb4, 0xf5, 0x2b, 0x44, 0xd0, 0x4a, 0x48, 0x4c, 0x8d,
	0xf5, 0xc5, 0xba, 0x98, 0x7d, 0x43, 0xa2, 0x35, 0x35, 0xb3, 0xcb, 0xe0, 0xa6, 0x2d, 0xbf, 0xcf,
	0x27, 0xdf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x61, 0xd1, 0x3f, 0x18, 0xd7, 0x05, 0x00, 0x00,
}
