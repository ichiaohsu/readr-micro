// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/proto/users.proto

package go_micro_srv_users

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Nickname             string   `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_a4849b6f7057ecac, []int{0}
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

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

type GetRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_a4849b6f7057ecac, []int{1}
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

func (m *GetRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Users struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_a4849b6f7057ecac, []int{2}
}
func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (dst *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(dst, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "go.micro.srv.users.User")
	proto.RegisterType((*GetRequest)(nil), "go.micro.srv.users.GetRequest")
	proto.RegisterType((*Users)(nil), "go.micro.srv.users.Users")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Users, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.users"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Users, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Get", in)
	out := new(Users)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Get(context.Context, *GetRequest, *Users) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) Get(ctx context.Context, in *GetRequest, out *Users) error {
	return h.UserServiceHandler.Get(ctx, in, out)
}

func init() { proto.RegisterFile("server/proto/users.proto", fileDescriptor_users_a4849b6f7057ecac) }

var fileDescriptor_users_a4849b6f7057ecac = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0xcd, 0x4a, 0xc4, 0x30,
	0x10, 0xa6, 0xed, 0xae, 0xe8, 0x2c, 0x2a, 0x0c, 0x08, 0xb1, 0x2c, 0x52, 0x7a, 0xea, 0x29, 0x81,
	0xf5, 0xe0, 0x1b, 0xb8, 0xf7, 0x8a, 0x78, 0xee, 0xb6, 0x63, 0x09, 0xda, 0x64, 0x9d, 0xc9, 0xf6,
	0x22, 0x5e, 0x7c, 0x05, 0x1f, 0xcd, 0x57, 0xf0, 0x41, 0xa4, 0x51, 0x14, 0xfc, 0xb9, 0xcd, 0xf7,
	0xe5, 0xfb, 0x23, 0xa0, 0x84, 0x78, 0x24, 0x36, 0x5b, 0xf6, 0xc1, 0x9b, 0x9d, 0x10, 0x8b, 0x8e,
	0x37, 0x62, 0xef, 0xf5, 0x60, 0x5b, 0xf6, 0x5a, 0x78, 0xd4, 0xf1, 0x25, 0x5f, 0xf6, 0xde, 0xf7,
	0xf7, 0x64, 0x9a, 0xad, 0x35, 0x8d, 0x73, 0x3e, 0x34, 0xc1, 0x7a, 0xf7, 0xe9, 0x28, 0x2f, 0x61,
	0x76, 0x2d, 0xc4, 0x78, 0x04, 0xa9, 0xed, 0x54, 0x52, 0x24, 0x55, 0x56, 0xa7, 0xb6, 0x43, 0x84,
	0x99, 0x6b, 0x06, 0x52, 0x69, 0x91, 0x54, 0x07, 0x75, 0xbc, 0x31, 0x87, 0x7d, 0x67, 0xdb, 0xbb,
	0xc8, 0x67, 0x91, 0xff, 0xc2, 0xe5, 0x12, 0x60, 0x4d, 0xa1, 0xa6, 0x87, 0x1d, 0x49, 0xf8, 0x99,
	0x56, 0x5e, 0xc0, 0x7c, 0x6a, 0x11, 0xd4, 0x30, 0x8f, 0xab, 0x54, 0x52, 0x64, 0xd5, 0x62, 0xa5,
	0xf4, 0xef, 0xc1, 0x7a, 0x52, 0xd6, 0x1f, 0xb2, 0xd5, 0x2d, 0x2c, 0x26, 0x78, 0x45, 0x3c, 0xda,
	0x96, 0xf0, 0x06, 0xb2, 0x35, 0x05, 0x3c, 0xfb, 0xcb, 0xf6, 0x5d, 0x9f, 0x9f, 0xfe, 0x17, 0x2b,
	0xe5, 0xc9, 0xf3, 0xeb, 0xdb, 0x4b, 0x7a, 0x8c, 0x87, 0x66, 0xa0, 0x61, 0x43, 0x2c, 0xe6, 0xd1,
	0x76, 0x4f, 0x9b, 0xbd, 0xf8, 0x1b, 0xe7, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x36, 0xd6, 0x25,
	0x5b, 0x5b, 0x01, 0x00, 0x00,
}