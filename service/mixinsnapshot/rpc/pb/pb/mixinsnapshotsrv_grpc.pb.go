// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: mixinsnapshotsrv.proto

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

// MixinsnapshotsrvClient is the client API for Mixinsnapshotsrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MixinsnapshotsrvClient interface {
	// -----------------------mixinsnapshot-----------------------
	AddMixinsnapshot(ctx context.Context, in *AddMixinsnapshotReq, opts ...grpc.CallOption) (*AddMixinsnapshotResp, error)
	UpdateMixinsnapshot(ctx context.Context, in *UpdateMixinsnapshotReq, opts ...grpc.CallOption) (*UpdateMixinsnapshotResp, error)
	DelMixinsnapshot(ctx context.Context, in *DelMixinsnapshotReq, opts ...grpc.CallOption) (*DelMixinsnapshotResp, error)
	GetMixinsnapshotById(ctx context.Context, in *GetMixinsnapshotByIdReq, opts ...grpc.CallOption) (*GetMixinsnapshotByIdResp, error)
	SearchMixinsnapshot(ctx context.Context, in *SearchMixinsnapshotReq, opts ...grpc.CallOption) (*SearchMixinsnapshotResp, error)
}

type mixinsnapshotsrvClient struct {
	cc grpc.ClientConnInterface
}

func NewMixinsnapshotsrvClient(cc grpc.ClientConnInterface) MixinsnapshotsrvClient {
	return &mixinsnapshotsrvClient{cc}
}

func (c *mixinsnapshotsrvClient) AddMixinsnapshot(ctx context.Context, in *AddMixinsnapshotReq, opts ...grpc.CallOption) (*AddMixinsnapshotResp, error) {
	out := new(AddMixinsnapshotResp)
	err := c.cc.Invoke(ctx, "/pb.mixinsnapshotsrv/AddMixinsnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixinsnapshotsrvClient) UpdateMixinsnapshot(ctx context.Context, in *UpdateMixinsnapshotReq, opts ...grpc.CallOption) (*UpdateMixinsnapshotResp, error) {
	out := new(UpdateMixinsnapshotResp)
	err := c.cc.Invoke(ctx, "/pb.mixinsnapshotsrv/UpdateMixinsnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixinsnapshotsrvClient) DelMixinsnapshot(ctx context.Context, in *DelMixinsnapshotReq, opts ...grpc.CallOption) (*DelMixinsnapshotResp, error) {
	out := new(DelMixinsnapshotResp)
	err := c.cc.Invoke(ctx, "/pb.mixinsnapshotsrv/DelMixinsnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixinsnapshotsrvClient) GetMixinsnapshotById(ctx context.Context, in *GetMixinsnapshotByIdReq, opts ...grpc.CallOption) (*GetMixinsnapshotByIdResp, error) {
	out := new(GetMixinsnapshotByIdResp)
	err := c.cc.Invoke(ctx, "/pb.mixinsnapshotsrv/GetMixinsnapshotById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixinsnapshotsrvClient) SearchMixinsnapshot(ctx context.Context, in *SearchMixinsnapshotReq, opts ...grpc.CallOption) (*SearchMixinsnapshotResp, error) {
	out := new(SearchMixinsnapshotResp)
	err := c.cc.Invoke(ctx, "/pb.mixinsnapshotsrv/SearchMixinsnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MixinsnapshotsrvServer is the server API for Mixinsnapshotsrv service.
// All implementations must embed UnimplementedMixinsnapshotsrvServer
// for forward compatibility
type MixinsnapshotsrvServer interface {
	// -----------------------mixinsnapshot-----------------------
	AddMixinsnapshot(context.Context, *AddMixinsnapshotReq) (*AddMixinsnapshotResp, error)
	UpdateMixinsnapshot(context.Context, *UpdateMixinsnapshotReq) (*UpdateMixinsnapshotResp, error)
	DelMixinsnapshot(context.Context, *DelMixinsnapshotReq) (*DelMixinsnapshotResp, error)
	GetMixinsnapshotById(context.Context, *GetMixinsnapshotByIdReq) (*GetMixinsnapshotByIdResp, error)
	SearchMixinsnapshot(context.Context, *SearchMixinsnapshotReq) (*SearchMixinsnapshotResp, error)
	mustEmbedUnimplementedMixinsnapshotsrvServer()
}

// UnimplementedMixinsnapshotsrvServer must be embedded to have forward compatible implementations.
type UnimplementedMixinsnapshotsrvServer struct {
}

func (UnimplementedMixinsnapshotsrvServer) AddMixinsnapshot(context.Context, *AddMixinsnapshotReq) (*AddMixinsnapshotResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMixinsnapshot not implemented")
}
func (UnimplementedMixinsnapshotsrvServer) UpdateMixinsnapshot(context.Context, *UpdateMixinsnapshotReq) (*UpdateMixinsnapshotResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMixinsnapshot not implemented")
}
func (UnimplementedMixinsnapshotsrvServer) DelMixinsnapshot(context.Context, *DelMixinsnapshotReq) (*DelMixinsnapshotResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelMixinsnapshot not implemented")
}
func (UnimplementedMixinsnapshotsrvServer) GetMixinsnapshotById(context.Context, *GetMixinsnapshotByIdReq) (*GetMixinsnapshotByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMixinsnapshotById not implemented")
}
func (UnimplementedMixinsnapshotsrvServer) SearchMixinsnapshot(context.Context, *SearchMixinsnapshotReq) (*SearchMixinsnapshotResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMixinsnapshot not implemented")
}
func (UnimplementedMixinsnapshotsrvServer) mustEmbedUnimplementedMixinsnapshotsrvServer() {}

// UnsafeMixinsnapshotsrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MixinsnapshotsrvServer will
// result in compilation errors.
type UnsafeMixinsnapshotsrvServer interface {
	mustEmbedUnimplementedMixinsnapshotsrvServer()
}

func RegisterMixinsnapshotsrvServer(s grpc.ServiceRegistrar, srv MixinsnapshotsrvServer) {
	s.RegisterService(&Mixinsnapshotsrv_ServiceDesc, srv)
}

func _Mixinsnapshotsrv_AddMixinsnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMixinsnapshotReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixinsnapshotsrvServer).AddMixinsnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.mixinsnapshotsrv/AddMixinsnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixinsnapshotsrvServer).AddMixinsnapshot(ctx, req.(*AddMixinsnapshotReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixinsnapshotsrv_UpdateMixinsnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMixinsnapshotReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixinsnapshotsrvServer).UpdateMixinsnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.mixinsnapshotsrv/UpdateMixinsnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixinsnapshotsrvServer).UpdateMixinsnapshot(ctx, req.(*UpdateMixinsnapshotReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixinsnapshotsrv_DelMixinsnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelMixinsnapshotReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixinsnapshotsrvServer).DelMixinsnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.mixinsnapshotsrv/DelMixinsnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixinsnapshotsrvServer).DelMixinsnapshot(ctx, req.(*DelMixinsnapshotReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixinsnapshotsrv_GetMixinsnapshotById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMixinsnapshotByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixinsnapshotsrvServer).GetMixinsnapshotById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.mixinsnapshotsrv/GetMixinsnapshotById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixinsnapshotsrvServer).GetMixinsnapshotById(ctx, req.(*GetMixinsnapshotByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixinsnapshotsrv_SearchMixinsnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchMixinsnapshotReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixinsnapshotsrvServer).SearchMixinsnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.mixinsnapshotsrv/SearchMixinsnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixinsnapshotsrvServer).SearchMixinsnapshot(ctx, req.(*SearchMixinsnapshotReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Mixinsnapshotsrv_ServiceDesc is the grpc.ServiceDesc for Mixinsnapshotsrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mixinsnapshotsrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.mixinsnapshotsrv",
	HandlerType: (*MixinsnapshotsrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMixinsnapshot",
			Handler:    _Mixinsnapshotsrv_AddMixinsnapshot_Handler,
		},
		{
			MethodName: "UpdateMixinsnapshot",
			Handler:    _Mixinsnapshotsrv_UpdateMixinsnapshot_Handler,
		},
		{
			MethodName: "DelMixinsnapshot",
			Handler:    _Mixinsnapshotsrv_DelMixinsnapshot_Handler,
		},
		{
			MethodName: "GetMixinsnapshotById",
			Handler:    _Mixinsnapshotsrv_GetMixinsnapshotById_Handler,
		},
		{
			MethodName: "SearchMixinsnapshot",
			Handler:    _Mixinsnapshotsrv_SearchMixinsnapshot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mixinsnapshotsrv.proto",
}
