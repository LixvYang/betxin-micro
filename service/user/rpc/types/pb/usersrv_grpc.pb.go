// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: usersrv.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UsersrvClient is the client API for Usersrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersrvClient interface {
	// -----------------------user-----------------------
	AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserResp, error)
	UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error)
	DelUser(ctx context.Context, in *DelUserReq, opts ...grpc.CallOption) (*DelUserResp, error)
	GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdResp, error)
	GetUserByUid(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdResp, error)
	SearchUser(ctx context.Context, in *SearchUserReq, opts ...grpc.CallOption) (*SearchUserResp, error)
}

type usersrvClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersrvClient(cc grpc.ClientConnInterface) UsersrvClient {
	return &usersrvClient{cc}
}

func (c *usersrvClient) AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserResp, error) {
	out := new(AddUserResp)
	err := c.cc.Invoke(ctx, "/pb.usersrv/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersrvClient) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error) {
	out := new(UpdateUserResp)
	err := c.cc.Invoke(ctx, "/pb.usersrv/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersrvClient) DelUser(ctx context.Context, in *DelUserReq, opts ...grpc.CallOption) (*DelUserResp, error) {
	out := new(DelUserResp)
	err := c.cc.Invoke(ctx, "/pb.usersrv/DelUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersrvClient) GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdResp, error) {
	out := new(GetUserByIdResp)
	err := c.cc.Invoke(ctx, "/pb.usersrv/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersrvClient) GetUserByUid(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdResp, error) {
	out := new(GetUserByIdResp)
	err := c.cc.Invoke(ctx, "/pb.usersrv/GetUserByUid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersrvClient) SearchUser(ctx context.Context, in *SearchUserReq, opts ...grpc.CallOption) (*SearchUserResp, error) {
	out := new(SearchUserResp)
	err := c.cc.Invoke(ctx, "/pb.usersrv/SearchUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersrvServer is the server API for Usersrv service.
// All implementations must embed UnimplementedUsersrvServer
// for forward compatibility
type UsersrvServer interface {
	// -----------------------user-----------------------
	AddUser(context.Context, *AddUserReq) (*AddUserResp, error)
	UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserResp, error)
	DelUser(context.Context, *DelUserReq) (*DelUserResp, error)
	GetUserById(context.Context, *GetUserByIdReq) (*GetUserByIdResp, error)
	GetUserByUid(context.Context, *GetUserByIdReq) (*GetUserByIdResp, error)
	SearchUser(context.Context, *SearchUserReq) (*SearchUserResp, error)
	mustEmbedUnimplementedUsersrvServer()
}

// UnimplementedUsersrvServer must be embedded to have forward compatible implementations.
type UnimplementedUsersrvServer struct {
}

func (UnimplementedUsersrvServer) AddUser(context.Context, *AddUserReq) (*AddUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedUsersrvServer) UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUsersrvServer) DelUser(context.Context, *DelUserReq) (*DelUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelUser not implemented")
}
func (UnimplementedUsersrvServer) GetUserById(context.Context, *GetUserByIdReq) (*GetUserByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUsersrvServer) GetUserByUid(context.Context, *GetUserByIdReq) (*GetUserByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByUid not implemented")
}
func (UnimplementedUsersrvServer) SearchUser(context.Context, *SearchUserReq) (*SearchUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUser not implemented")
}
func (UnimplementedUsersrvServer) mustEmbedUnimplementedUsersrvServer() {}

// UnsafeUsersrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersrvServer will
// result in compilation errors.
type UnsafeUsersrvServer interface {
	mustEmbedUnimplementedUsersrvServer()
}

func RegisterUsersrvServer(s grpc.ServiceRegistrar, srv UsersrvServer) {
	s.RegisterService(&Usersrv_ServiceDesc, srv)
}

func _Usersrv_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersrvServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usersrv/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersrvServer).AddUser(ctx, req.(*AddUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usersrv_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersrvServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usersrv/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersrvServer).UpdateUser(ctx, req.(*UpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usersrv_DelUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersrvServer).DelUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usersrv/DelUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersrvServer).DelUser(ctx, req.(*DelUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usersrv_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersrvServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usersrv/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersrvServer).GetUserById(ctx, req.(*GetUserByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usersrv_GetUserByUid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersrvServer).GetUserByUid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usersrv/GetUserByUid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersrvServer).GetUserByUid(ctx, req.(*GetUserByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usersrv_SearchUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersrvServer).SearchUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usersrv/SearchUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersrvServer).SearchUser(ctx, req.(*SearchUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Usersrv_ServiceDesc is the grpc.ServiceDesc for Usersrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Usersrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.usersrv",
	HandlerType: (*UsersrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _Usersrv_AddUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Usersrv_UpdateUser_Handler,
		},
		{
			MethodName: "DelUser",
			Handler:    _Usersrv_DelUser_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _Usersrv_GetUserById_Handler,
		},
		{
			MethodName: "GetUserByUid",
			Handler:    _Usersrv_GetUserByUid_Handler,
		},
		{
			MethodName: "SearchUser",
			Handler:    _Usersrv_SearchUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usersrv.proto",
}
