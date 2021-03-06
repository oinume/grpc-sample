// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/api/v1/users.proto

package api_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type ListUsersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_45cb83530147bc38, []int{0}
}
func (m *ListUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersRequest.Unmarshal(m, b)
}
func (m *ListUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersRequest.Marshal(b, m, deterministic)
}
func (dst *ListUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersRequest.Merge(dst, src)
}
func (m *ListUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ListUsersRequest.Size(m)
}
func (m *ListUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersRequest proto.InternalMessageInfo

type ListUsersResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersResponse) Reset()         { *m = ListUsersResponse{} }
func (m *ListUsersResponse) String() string { return proto.CompactTextString(m) }
func (*ListUsersResponse) ProtoMessage()    {}
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_45cb83530147bc38, []int{1}
}
func (m *ListUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersResponse.Unmarshal(m, b)
}
func (m *ListUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersResponse.Marshal(b, m, deterministic)
}
func (dst *ListUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersResponse.Merge(dst, src)
}
func (m *ListUsersResponse) XXX_Size() int {
	return xxx_messageInfo_ListUsersResponse.Size(m)
}
func (m *ListUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersResponse proto.InternalMessageInfo

func (m *ListUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type GetUserRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_45cb83530147bc38, []int{2}
}
func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (dst *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(dst, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	RealName             string   `protobuf:"bytes,2,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_45cb83530147bc38, []int{3}
}
func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (dst *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(dst, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateUserRequest) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

type UpdateUserRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_45cb83530147bc38, []int{4}
}
func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(dst, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UpdateUserResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserResponse) Reset()         { *m = UpdateUserResponse{} }
func (m *UpdateUserResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateUserResponse) ProtoMessage()    {}
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_45cb83530147bc38, []int{5}
}
func (m *UpdateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserResponse.Unmarshal(m, b)
}
func (m *UpdateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserResponse.Merge(dst, src)
}
func (m *UpdateUserResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateUserResponse.Size(m)
}
func (m *UpdateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserResponse proto.InternalMessageInfo

type DeleteUserRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_45cb83530147bc38, []int{6}
}
func (m *DeleteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserRequest.Unmarshal(m, b)
}
func (m *DeleteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserRequest.Merge(dst, src)
}
func (m *DeleteUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserRequest.Size(m)
}
func (m *DeleteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserRequest proto.InternalMessageInfo

func (m *DeleteUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteUserResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserResponse) Reset()         { *m = DeleteUserResponse{} }
func (m *DeleteUserResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteUserResponse) ProtoMessage()    {}
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_45cb83530147bc38, []int{7}
}
func (m *DeleteUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserResponse.Unmarshal(m, b)
}
func (m *DeleteUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserResponse.Merge(dst, src)
}
func (m *DeleteUserResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteUserResponse.Size(m)
}
func (m *DeleteUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserResponse proto.InternalMessageInfo

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RealName             string   `protobuf:"bytes,3,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_45cb83530147bc38, []int{8}
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

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func init() {
	proto.RegisterType((*ListUsersRequest)(nil), "api.v1.ListUsersRequest")
	proto.RegisterType((*ListUsersResponse)(nil), "api.v1.ListUsersResponse")
	proto.RegisterType((*GetUserRequest)(nil), "api.v1.GetUserRequest")
	proto.RegisterType((*CreateUserRequest)(nil), "api.v1.CreateUserRequest")
	proto.RegisterType((*UpdateUserRequest)(nil), "api.v1.UpdateUserRequest")
	proto.RegisterType((*UpdateUserResponse)(nil), "api.v1.UpdateUserResponse")
	proto.RegisterType((*DeleteUserRequest)(nil), "api.v1.DeleteUserRequest")
	proto.RegisterType((*DeleteUserResponse)(nil), "api.v1.DeleteUserResponse")
	proto.RegisterType((*User)(nil), "api.v1.User")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/api.v1.Users/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/api.v1.Users/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/api.v1.Users/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/api.v1.Users/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.v1.Users/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*empty.Empty, error)
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.Users/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.Users/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.Users/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.Users/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.Users/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUsers",
			Handler:    _Users_ListUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Users_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Users_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Users_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Users_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api/v1/users.proto",
}

func init() { proto.RegisterFile("proto/api/v1/users.proto", fileDescriptor_users_45cb83530147bc38) }

var fileDescriptor_users_45cb83530147bc38 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcb, 0x6a, 0xdb, 0x40,
	0x14, 0x86, 0x91, 0x7c, 0x69, 0x75, 0x5c, 0x84, 0x75, 0x6a, 0x8c, 0x3c, 0xee, 0xc2, 0x4c, 0x37,
	0xc6, 0x0b, 0x09, 0xbb, 0x8b, 0x42, 0xb7, 0x75, 0xf1, 0xa2, 0xa5, 0x0b, 0x17, 0xd3, 0x5d, 0xcb,
	0x18, 0x4d, 0x8c, 0xc0, 0x96, 0x14, 0x49, 0x36, 0x84, 0x90, 0x4d, 0x5e, 0x21, 0x8f, 0x91, 0xc7,
	0xc9, 0x2b, 0xe4, 0x41, 0xc2, 0x8c, 0x46, 0x57, 0x9b, 0x90, 0x9d, 0x66, 0xce, 0x7f, 0xbe, 0x73,
	0xf9, 0x35, 0x60, 0x47, 0x71, 0x98, 0x86, 0x2e, 0x8b, 0x7c, 0xf7, 0x34, 0x77, 0x8f, 0x09, 0x8f,
	0x13, 0x47, 0x5e, 0x61, 0x97, 0x45, 0xbe, 0x73, 0x9a, 0x93, 0x4f, 0xbb, 0x30, 0xdc, 0xed, 0xb9,
	0x94, 0xb0, 0x20, 0x08, 0x53, 0x96, 0xfa, 0x61, 0xa0, 0x54, 0x64, 0xac, 0xa2, 0xf2, 0xb4, 0x3d,
	0x5e, 0xb9, 0xfc, 0x10, 0xa5, 0x37, 0x59, 0x90, 0x22, 0xf4, 0x7f, 0xf9, 0x49, 0xba, 0x11, 0xd4,
	0x35, 0xbf, 0x3e, 0xf2, 0x24, 0xa5, 0x5f, 0xc1, 0xaa, 0xdc, 0x25, 0x51, 0x18, 0x24, 0x1c, 0x29,
	0x74, 0x64, 0x69, 0x5b, 0x9b, 0xb4, 0xa6, 0xbd, 0xc5, 0x07, 0x27, 0xab, 0xed, 0x08, 0xd5, 0x3a,
	0x0b, 0xd1, 0x09, 0x98, 0x2b, 0x2e, 0xf3, 0x14, 0x0a, 0x4d, 0xd0, 0x7d, 0xcf, 0xd6, 0x26, 0xda,
	0xd4, 0x58, 0xeb, 0xbe, 0x47, 0x97, 0x60, 0x7d, 0x8f, 0x39, 0x4b, 0x79, 0x55, 0x84, 0xd0, 0x0e,
	0xd8, 0x81, 0x2b, 0x99, 0xfc, 0xc6, 0x31, 0x18, 0x31, 0x67, 0xfb, 0xff, 0x32, 0xa0, 0xcb, 0xc0,
	0x7b, 0x71, 0xf1, 0x9b, 0x1d, 0xb8, 0x68, 0x70, 0x13, 0x79, 0x0d, 0x4a, 0xa3, 0x54, 0x41, 0xd5,
	0x4b, 0x2a, 0x1d, 0x00, 0x56, 0x13, 0xb3, 0xd1, 0xe8, 0x67, 0xb0, 0x96, 0x7c, 0xcf, 0x5f, 0xc5,
	0x89, 0xd4, 0xaa, 0x48, 0xa5, 0xae, 0xa0, 0x2d, 0xce, 0x6f, 0x29, 0x5e, 0x1f, 0xa9, 0x55, 0x1f,
	0x69, 0xf1, 0xd8, 0x82, 0x8e, 0x5c, 0x38, 0xfe, 0x01, 0xa3, 0xd8, 0x3e, 0xda, 0xf9, 0x9a, 0x9b,
	0x26, 0x91, 0xd1, 0x85, 0x88, 0x6a, 0xca, 0xba, 0x7f, 0x7a, 0x7e, 0xd0, 0x7b, 0x68, 0x14, 0xff,
	0x0b, 0xae, 0xe0, 0x9d, 0x72, 0x06, 0x87, 0x79, 0x62, 0xdd, 0x2a, 0x52, 0x73, 0x94, 0x0e, 0x25,
	0xa3, 0x8f, 0x66, 0xc1, 0x70, 0x6f, 0x7d, 0xef, 0x0e, 0x7f, 0x02, 0x94, 0x06, 0x62, 0xd1, 0xc4,
	0x99, 0xa9, 0x0d, 0xdc, 0x40, 0xe2, 0x4c, 0x5a, 0xb6, 0xf4, 0x4d, 0x9b, 0xe1, 0x3f, 0x80, 0xd2,
	0x8e, 0x12, 0x76, 0xe6, 0x2d, 0x21, 0x97, 0x42, 0x6a, 0xda, 0x91, 0x44, 0x7f, 0x24, 0x8d, 0x4e,
	0x05, 0xff, 0x2f, 0x40, 0xe9, 0x59, 0xc9, 0x3f, 0x33, 0x9b, 0x0c, 0x9d, 0xec, 0x8d, 0x38, 0xf9,
	0x1b, 0x71, 0x7e, 0x88, 0x37, 0x92, 0x6f, 0x61, 0xd6, 0x60, 0x6f, 0xbb, 0x52, 0xf7, 0xe5, 0x25,
	0x00, 0x00, 0xff, 0xff, 0x62, 0xd3, 0xd5, 0x74, 0x9b, 0x03, 0x00, 0x00,
}
