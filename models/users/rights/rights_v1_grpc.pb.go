// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rights

import (
	context "context"
	common "github.com/badhouseplants/envspotting-go-proto/models/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RightsClient is the client API for Rights service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RightsClient interface {
	Init(ctx context.Context, in *AccessRuleWithoutId, opts ...grpc.CallOption) (*AccessRuleInfo, error)
	Create(ctx context.Context, in *AccessRuleWithoutId, opts ...grpc.CallOption) (*AccessRuleInfo, error)
	Update(ctx context.Context, in *AccessRuleIdAndRight, opts ...grpc.CallOption) (*AccessRuleIdAndRight, error)
	Delete(ctx context.Context, in *AccessRuleId, opts ...grpc.CallOption) (*common.EmptyMessage, error)
	Get(ctx context.Context, in *AccessRuleId, opts ...grpc.CallOption) (*AccessRuleInfo, error)
	List(ctx context.Context, in *RightsListOptions, opts ...grpc.CallOption) (Rights_ListClient, error)
}

type rightsClient struct {
	cc grpc.ClientConnInterface
}

func NewRightsClient(cc grpc.ClientConnInterface) RightsClient {
	return &rightsClient{cc}
}

func (c *rightsClient) Init(ctx context.Context, in *AccessRuleWithoutId, opts ...grpc.CallOption) (*AccessRuleInfo, error) {
	out := new(AccessRuleInfo)
	err := c.cc.Invoke(ctx, "/users.Rights/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rightsClient) Create(ctx context.Context, in *AccessRuleWithoutId, opts ...grpc.CallOption) (*AccessRuleInfo, error) {
	out := new(AccessRuleInfo)
	err := c.cc.Invoke(ctx, "/users.Rights/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rightsClient) Update(ctx context.Context, in *AccessRuleIdAndRight, opts ...grpc.CallOption) (*AccessRuleIdAndRight, error) {
	out := new(AccessRuleIdAndRight)
	err := c.cc.Invoke(ctx, "/users.Rights/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rightsClient) Delete(ctx context.Context, in *AccessRuleId, opts ...grpc.CallOption) (*common.EmptyMessage, error) {
	out := new(common.EmptyMessage)
	err := c.cc.Invoke(ctx, "/users.Rights/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rightsClient) Get(ctx context.Context, in *AccessRuleId, opts ...grpc.CallOption) (*AccessRuleInfo, error) {
	out := new(AccessRuleInfo)
	err := c.cc.Invoke(ctx, "/users.Rights/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rightsClient) List(ctx context.Context, in *RightsListOptions, opts ...grpc.CallOption) (Rights_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &Rights_ServiceDesc.Streams[0], "/users.Rights/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &rightsListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Rights_ListClient interface {
	Recv() (*AccessRuleInfo, error)
	grpc.ClientStream
}

type rightsListClient struct {
	grpc.ClientStream
}

func (x *rightsListClient) Recv() (*AccessRuleInfo, error) {
	m := new(AccessRuleInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RightsServer is the server API for Rights service.
// All implementations must embed UnimplementedRightsServer
// for forward compatibility
type RightsServer interface {
	Init(context.Context, *AccessRuleWithoutId) (*AccessRuleInfo, error)
	Create(context.Context, *AccessRuleWithoutId) (*AccessRuleInfo, error)
	Update(context.Context, *AccessRuleIdAndRight) (*AccessRuleIdAndRight, error)
	Delete(context.Context, *AccessRuleId) (*common.EmptyMessage, error)
	Get(context.Context, *AccessRuleId) (*AccessRuleInfo, error)
	List(*RightsListOptions, Rights_ListServer) error
	mustEmbedUnimplementedRightsServer()
}

// UnimplementedRightsServer must be embedded to have forward compatible implementations.
type UnimplementedRightsServer struct {
}

func (UnimplementedRightsServer) Init(context.Context, *AccessRuleWithoutId) (*AccessRuleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedRightsServer) Create(context.Context, *AccessRuleWithoutId) (*AccessRuleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRightsServer) Update(context.Context, *AccessRuleIdAndRight) (*AccessRuleIdAndRight, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRightsServer) Delete(context.Context, *AccessRuleId) (*common.EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRightsServer) Get(context.Context, *AccessRuleId) (*AccessRuleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRightsServer) List(*RightsListOptions, Rights_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRightsServer) mustEmbedUnimplementedRightsServer() {}

// UnsafeRightsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RightsServer will
// result in compilation errors.
type UnsafeRightsServer interface {
	mustEmbedUnimplementedRightsServer()
}

func RegisterRightsServer(s grpc.ServiceRegistrar, srv RightsServer) {
	s.RegisterService(&Rights_ServiceDesc, srv)
}

func _Rights_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessRuleWithoutId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RightsServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Rights/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RightsServer).Init(ctx, req.(*AccessRuleWithoutId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rights_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessRuleWithoutId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RightsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Rights/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RightsServer).Create(ctx, req.(*AccessRuleWithoutId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rights_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessRuleIdAndRight)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RightsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Rights/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RightsServer).Update(ctx, req.(*AccessRuleIdAndRight))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rights_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessRuleId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RightsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Rights/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RightsServer).Delete(ctx, req.(*AccessRuleId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rights_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessRuleId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RightsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Rights/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RightsServer).Get(ctx, req.(*AccessRuleId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rights_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RightsListOptions)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RightsServer).List(m, &rightsListServer{stream})
}

type Rights_ListServer interface {
	Send(*AccessRuleInfo) error
	grpc.ServerStream
}

type rightsListServer struct {
	grpc.ServerStream
}

func (x *rightsListServer) Send(m *AccessRuleInfo) error {
	return x.ServerStream.SendMsg(m)
}

// Rights_ServiceDesc is the grpc.ServiceDesc for Rights service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rights_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.Rights",
	HandlerType: (*RightsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _Rights_Init_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Rights_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Rights_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Rights_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Rights_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _Rights_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "users/rights/rights_v1.proto",
}
