// Code generated by protoc-gen-go. DO NOT EDIT.
// source: identity.proto

package proto // import "."

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_identity_e68f189c05486efb, []int{0}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type GetByIDRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByIDRequest) Reset()         { *m = GetByIDRequest{} }
func (m *GetByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetByIDRequest) ProtoMessage()    {}
func (*GetByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_identity_e68f189c05486efb, []int{1}
}
func (m *GetByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByIDRequest.Unmarshal(m, b)
}
func (m *GetByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByIDRequest.Marshal(b, m, deterministic)
}
func (dst *GetByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByIDRequest.Merge(dst, src)
}
func (m *GetByIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetByIDRequest.Size(m)
}
func (m *GetByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetByIDRequest proto.InternalMessageInfo

func (m *GetByIDRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListRequest struct {
	Token                string            `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Limit                int64             `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	OrderBy              string            `protobuf:"bytes,3,opt,name=orderBy,proto3" json:"orderBy,omitempty"`
	Filter               map[string]string `protobuf:"bytes,4,rep,name=filter,proto3" json:"filter,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_identity_e68f189c05486efb, []int{2}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (dst *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(dst, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ListRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListRequest) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

func (m *ListRequest) GetFilter() map[string]string {
	if m != nil {
		return m.Filter
	}
	return nil
}

type CommandRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandRequest) Reset()         { *m = CommandRequest{} }
func (m *CommandRequest) String() string { return proto.CompactTextString(m) }
func (*CommandRequest) ProtoMessage()    {}
func (*CommandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_identity_e68f189c05486efb, []int{3}
}
func (m *CommandRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandRequest.Unmarshal(m, b)
}
func (m *CommandRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandRequest.Marshal(b, m, deterministic)
}
func (dst *CommandRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandRequest.Merge(dst, src)
}
func (m *CommandRequest) XXX_Size() int {
	return xxx_messageInfo_CommandRequest.Size(m)
}
func (m *CommandRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommandRequest proto.InternalMessageInfo

func (m *CommandRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_identity_e68f189c05486efb, []int{4}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	PreferredUsername    string   `protobuf:"bytes,3,opt,name=preferredUsername,proto3" json:"preferredUsername,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	MiddleName           string   `protobuf:"bytes,6,opt,name=middleName,proto3" json:"middleName,omitempty"`
	FamilyName           string   `protobuf:"bytes,7,opt,name=familyName,proto3" json:"familyName,omitempty"`
	Locale               string   `protobuf:"bytes,8,opt,name=locale,proto3" json:"locale,omitempty"`
	Picture              string   `protobuf:"bytes,9,opt,name=picture,proto3" json:"picture,omitempty"`
	Status               string   `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	CreateTime           string   `protobuf:"bytes,11,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime           string   `protobuf:"bytes,12,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	Enabled              bool     `protobuf:"varint,13,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_identity_e68f189c05486efb, []int{5}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPreferredUsername() string {
	if m != nil {
		return m.PreferredUsername
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetMiddleName() string {
	if m != nil {
		return m.MiddleName
	}
	return ""
}

func (m *User) GetFamilyName() string {
	if m != nil {
		return m.FamilyName
	}
	return ""
}

func (m *User) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

func (m *User) GetPicture() string {
	if m != nil {
		return m.Picture
	}
	return ""
}

func (m *User) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *User) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *User) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *User) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type ListResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	NextPageToken        string   `protobuf:"bytes,2,opt,name=nextPageToken,proto3" json:"nextPageToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_identity_e68f189c05486efb, []int{6}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (dst *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(dst, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *ListResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "proto.GetRequest")
	proto.RegisterType((*GetByIDRequest)(nil), "proto.GetByIDRequest")
	proto.RegisterType((*ListRequest)(nil), "proto.ListRequest")
	proto.RegisterMapType((map[string]string)(nil), "proto.ListRequest.FilterEntry")
	proto.RegisterType((*CommandRequest)(nil), "proto.CommandRequest")
	proto.RegisterType((*Empty)(nil), "proto.Empty")
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*ListResponse)(nil), "proto.ListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IdentityClient is the client API for Identity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdentityClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*User, error)
	GetByID(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*User, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Enable(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*Empty, error)
	Disable(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*Empty, error)
	Remove(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*Empty, error)
}

type identityClient struct {
	cc *grpc.ClientConn
}

func NewIdentityClient(cc *grpc.ClientConn) IdentityClient {
	return &identityClient{cc}
}

func (c *identityClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.Identity/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) GetByID(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.Identity/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/proto.Identity/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) Enable(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Identity/Enable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) Disable(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Identity/Disable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) Remove(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Identity/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServer is the server API for Identity service.
type IdentityServer interface {
	Get(context.Context, *GetRequest) (*User, error)
	GetByID(context.Context, *GetByIDRequest) (*User, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Enable(context.Context, *CommandRequest) (*Empty, error)
	Disable(context.Context, *CommandRequest) (*Empty, error)
	Remove(context.Context, *CommandRequest) (*Empty, error)
}

func RegisterIdentityServer(s *grpc.Server, srv IdentityServer) {
	s.RegisterService(&_Identity_serviceDesc, srv)
}

func _Identity_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Identity/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Identity/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).GetByID(ctx, req.(*GetByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Identity/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Identity/Enable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).Enable(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Identity/Disable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).Disable(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Identity/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).Remove(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Identity_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Identity",
	HandlerType: (*IdentityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Identity_Get_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _Identity_GetByID_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Identity_List_Handler,
		},
		{
			MethodName: "Enable",
			Handler:    _Identity_Enable_Handler,
		},
		{
			MethodName: "Disable",
			Handler:    _Identity_Disable_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Identity_Remove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity.proto",
}

func init() { proto.RegisterFile("identity.proto", fileDescriptor_identity_e68f189c05486efb) }

var fileDescriptor_identity_e68f189c05486efb = []byte{
	// 534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0x46, 0xb2, 0x2d, 0xd9, 0x63, 0xc7, 0x34, 0xdb, 0x1f, 0x16, 0x1f, 0x82, 0x2b, 0x5a, 0xf0,
	0x21, 0x76, 0x21, 0x85, 0xd2, 0x9f, 0x9b, 0x1b, 0xd7, 0x04, 0x4a, 0x29, 0x22, 0xa1, 0xd0, 0x9b,
	0x62, 0x8d, 0xcb, 0x12, 0xad, 0xa4, 0xee, 0xae, 0x42, 0xf5, 0x34, 0x7d, 0x9b, 0x3e, 0x4e, 0x9f,
	0xa1, 0xec, 0x8f, 0x6d, 0xa9, 0xb9, 0xf8, 0x24, 0x7d, 0x3f, 0xb3, 0x33, 0x3b, 0x3b, 0x03, 0x63,
	0x96, 0x62, 0xae, 0x98, 0xaa, 0x17, 0xa5, 0x28, 0x54, 0x41, 0x7a, 0xe6, 0x13, 0xcd, 0x00, 0xd6,
	0xa8, 0x62, 0xfc, 0x59, 0xa1, 0x54, 0x64, 0x02, 0xfd, 0x4a, 0xa2, 0xc8, 0x13, 0x8e, 0xd4, 0x9b,
	0x7a, 0xb3, 0x41, 0xbc, 0xc7, 0xd1, 0x14, 0xc6, 0x6b, 0x54, 0xcb, 0xfa, 0xea, 0x72, 0xe7, 0x1e,
	0x83, 0xcf, 0x52, 0xe7, 0xf3, 0x59, 0x1a, 0xfd, 0xf1, 0x60, 0xf8, 0x99, 0xc9, 0xfd, 0x69, 0x4f,
	0xa0, 0xa7, 0x8a, 0x3b, 0xcc, 0x9d, 0xc5, 0x02, 0xcd, 0x66, 0x8c, 0x33, 0x45, 0xfd, 0xa9, 0x37,
	0xeb, 0xc4, 0x16, 0x10, 0x0a, 0x61, 0x21, 0x52, 0x14, 0xcb, 0x9a, 0x76, 0x8c, 0x7b, 0x07, 0xc9,
	0x1b, 0x08, 0xb6, 0x2c, 0x53, 0x28, 0x68, 0x77, 0xda, 0x99, 0x0d, 0x2f, 0xce, 0xec, 0x05, 0x16,
	0x8d, 0x4c, 0x8b, 0x4f, 0xc6, 0xb0, 0xca, 0x95, 0xa8, 0x63, 0xe7, 0x9e, 0xbc, 0x83, 0x61, 0x83,
	0x26, 0x8f, 0xa0, 0x73, 0x87, 0xb5, 0x2b, 0x45, 0xff, 0xea, 0x42, 0xee, 0x93, 0xac, 0x42, 0x53,
	0xc8, 0x20, 0xb6, 0xe0, 0xbd, 0xff, 0xd6, 0x8b, 0xce, 0x61, 0xfc, 0xb1, 0xe0, 0x3c, 0xc9, 0xd3,
	0x63, 0x1a, 0x13, 0x42, 0x6f, 0xc5, 0x4b, 0x55, 0x47, 0x7f, 0x7d, 0xe8, 0xde, 0x48, 0x14, 0xff,
	0x37, 0xa6, 0x15, 0xed, 0xb7, 0xa3, 0xc9, 0x39, 0x9c, 0x96, 0x02, 0xb7, 0x28, 0x04, 0xa6, 0x37,
	0x3b, 0x93, 0x6d, 0xc1, 0x43, 0x41, 0xd7, 0x8c, 0x3c, 0x61, 0x19, 0xed, 0xda, 0x9a, 0x0d, 0x20,
	0x04, 0xba, 0x26, 0xac, 0x67, 0x48, 0xf3, 0x4f, 0xce, 0x00, 0x38, 0x4b, 0xd3, 0x0c, 0xbf, 0x68,
	0x25, 0x30, 0x4a, 0x83, 0xd1, 0xfa, 0x36, 0xe1, 0x2c, 0xab, 0x8d, 0x1e, 0x5a, 0xfd, 0xc0, 0x90,
	0x67, 0x10, 0x64, 0xc5, 0x26, 0xc9, 0x90, 0xf6, 0x8d, 0xe6, 0x90, 0x7e, 0xa8, 0x92, 0x6d, 0x54,
	0x25, 0x90, 0x0e, 0xec, 0x43, 0x39, 0xa8, 0x23, 0xa4, 0x4a, 0x54, 0x25, 0x29, 0xd8, 0x08, 0x8b,
	0x74, 0xa6, 0x8d, 0xc0, 0x44, 0xe1, 0x35, 0xe3, 0x48, 0x87, 0x36, 0xd3, 0x81, 0xd1, 0x7a, 0x55,
	0xa6, 0x3b, 0x7d, 0x64, 0xf5, 0x03, 0xa3, 0x33, 0x62, 0x9e, 0xdc, 0x66, 0x98, 0xd2, 0x93, 0xa9,
	0x37, 0xeb, 0xc7, 0x3b, 0x18, 0x7d, 0x83, 0x91, 0x9d, 0x02, 0x59, 0x16, 0xb9, 0x44, 0xf2, 0x1c,
	0x7a, 0xba, 0xaf, 0x92, 0x7a, 0x66, 0x52, 0x86, 0x6e, 0x52, 0x74, 0xf7, 0x62, 0xab, 0x90, 0x17,
	0x70, 0x92, 0xe3, 0x2f, 0xf5, 0x35, 0xf9, 0x81, 0xd7, 0x66, 0x36, 0xed, 0x7b, 0xb4, 0xc9, 0x8b,
	0xdf, 0x3e, 0xf4, 0xaf, 0xdc, 0xbe, 0x90, 0x97, 0xd0, 0x59, 0xa3, 0x22, 0xa7, 0xee, 0xb4, 0xc3,
	0xba, 0x4c, 0x9a, 0x09, 0xc8, 0x1c, 0x42, 0xb7, 0x1f, 0xe4, 0xe9, 0xc1, 0xda, 0xd8, 0x97, 0xb6,
	0xfd, 0x15, 0x74, 0x75, 0xed, 0x84, 0x3c, 0x1c, 0xe7, 0xc9, 0xe3, 0x16, 0xe7, 0x2e, 0x37, 0x87,
	0x60, 0x65, 0xee, 0xbd, 0x3f, 0xbe, 0x3d, 0xa3, 0x93, 0x91, 0xa3, 0xcd, 0x30, 0x92, 0x05, 0x84,
	0x97, 0x4c, 0x1e, 0xef, 0x9f, 0x43, 0x10, 0x23, 0x2f, 0xee, 0x8f, 0xb3, 0x2f, 0x07, 0xdf, 0xc3,
	0xc5, 0x07, 0x43, 0xdc, 0x06, 0xe6, 0xf3, 0xfa, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x5b,
	0xbf, 0xe3, 0x62, 0x04, 0x00, 0x00,
}
