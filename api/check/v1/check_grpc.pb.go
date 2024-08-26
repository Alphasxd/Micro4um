// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: api/check/v1/check.proto

package check

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Check_CreateCheck_FullMethodName  = "/api.check.v1.Check/CreateCheck"
	Check_DeleteCheck_FullMethodName  = "/api.check.v1.Check/DeleteCheck"
	Check_GetCheckById_FullMethodName = "/api.check.v1.Check/GetCheckById"
	Check_ListChecks_FullMethodName   = "/api.check.v1.Check/ListChecks"
	Check_SubmitCheck_FullMethodName  = "/api.check.v1.Check/SubmitCheck"
)

// CheckClient is the client API for Check service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckClient interface {
	CreateCheck(ctx context.Context, in *CreateCheckRequest, opts ...grpc.CallOption) (*CreateCheckReply, error)
	DeleteCheck(ctx context.Context, in *DeleteCheckRequest, opts ...grpc.CallOption) (*DeleteCheckReply, error)
	GetCheckById(ctx context.Context, in *GetCheckByIdRequest, opts ...grpc.CallOption) (*GetCheckByIdReply, error)
	ListChecks(ctx context.Context, in *ListChecksRequest, opts ...grpc.CallOption) (*ListChecksReply, error)
	SubmitCheck(ctx context.Context, in *SubmitCheckRequest, opts ...grpc.CallOption) (*SubmitCheckReply, error)
}

type checkClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckClient(cc grpc.ClientConnInterface) CheckClient {
	return &checkClient{cc}
}

func (c *checkClient) CreateCheck(ctx context.Context, in *CreateCheckRequest, opts ...grpc.CallOption) (*CreateCheckReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCheckReply)
	err := c.cc.Invoke(ctx, Check_CreateCheck_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkClient) DeleteCheck(ctx context.Context, in *DeleteCheckRequest, opts ...grpc.CallOption) (*DeleteCheckReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCheckReply)
	err := c.cc.Invoke(ctx, Check_DeleteCheck_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkClient) GetCheckById(ctx context.Context, in *GetCheckByIdRequest, opts ...grpc.CallOption) (*GetCheckByIdReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCheckByIdReply)
	err := c.cc.Invoke(ctx, Check_GetCheckById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkClient) ListChecks(ctx context.Context, in *ListChecksRequest, opts ...grpc.CallOption) (*ListChecksReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListChecksReply)
	err := c.cc.Invoke(ctx, Check_ListChecks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkClient) SubmitCheck(ctx context.Context, in *SubmitCheckRequest, opts ...grpc.CallOption) (*SubmitCheckReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubmitCheckReply)
	err := c.cc.Invoke(ctx, Check_SubmitCheck_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckServer is the server API for Check service.
// All implementations must embed UnimplementedCheckServer
// for forward compatibility
type CheckServer interface {
	CreateCheck(context.Context, *CreateCheckRequest) (*CreateCheckReply, error)
	DeleteCheck(context.Context, *DeleteCheckRequest) (*DeleteCheckReply, error)
	GetCheckById(context.Context, *GetCheckByIdRequest) (*GetCheckByIdReply, error)
	ListChecks(context.Context, *ListChecksRequest) (*ListChecksReply, error)
	SubmitCheck(context.Context, *SubmitCheckRequest) (*SubmitCheckReply, error)
	mustEmbedUnimplementedCheckServer()
}

// UnimplementedCheckServer must be embedded to have forward compatible implementations.
type UnimplementedCheckServer struct {
}

func (UnimplementedCheckServer) CreateCheck(context.Context, *CreateCheckRequest) (*CreateCheckReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCheck not implemented")
}
func (UnimplementedCheckServer) DeleteCheck(context.Context, *DeleteCheckRequest) (*DeleteCheckReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCheck not implemented")
}
func (UnimplementedCheckServer) GetCheckById(context.Context, *GetCheckByIdRequest) (*GetCheckByIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCheckById not implemented")
}
func (UnimplementedCheckServer) ListChecks(context.Context, *ListChecksRequest) (*ListChecksReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChecks not implemented")
}
func (UnimplementedCheckServer) SubmitCheck(context.Context, *SubmitCheckRequest) (*SubmitCheckReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitCheck not implemented")
}
func (UnimplementedCheckServer) mustEmbedUnimplementedCheckServer() {}

// UnsafeCheckServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CheckServer will
// result in compilation errors.
type UnsafeCheckServer interface {
	mustEmbedUnimplementedCheckServer()
}

func RegisterCheckServer(s grpc.ServiceRegistrar, srv CheckServer) {
	s.RegisterService(&Check_ServiceDesc, srv)
}

func _Check_CreateCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServer).CreateCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Check_CreateCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServer).CreateCheck(ctx, req.(*CreateCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Check_DeleteCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServer).DeleteCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Check_DeleteCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServer).DeleteCheck(ctx, req.(*DeleteCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Check_GetCheckById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCheckByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServer).GetCheckById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Check_GetCheckById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServer).GetCheckById(ctx, req.(*GetCheckByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Check_ListChecks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListChecksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServer).ListChecks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Check_ListChecks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServer).ListChecks(ctx, req.(*ListChecksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Check_SubmitCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServer).SubmitCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Check_SubmitCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServer).SubmitCheck(ctx, req.(*SubmitCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Check_ServiceDesc is the grpc.ServiceDesc for Check service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Check_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.check.v1.Check",
	HandlerType: (*CheckServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCheck",
			Handler:    _Check_CreateCheck_Handler,
		},
		{
			MethodName: "DeleteCheck",
			Handler:    _Check_DeleteCheck_Handler,
		},
		{
			MethodName: "GetCheckById",
			Handler:    _Check_GetCheckById_Handler,
		},
		{
			MethodName: "ListChecks",
			Handler:    _Check_ListChecks_Handler,
		},
		{
			MethodName: "SubmitCheck",
			Handler:    _Check_SubmitCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/check/v1/check.proto",
}
