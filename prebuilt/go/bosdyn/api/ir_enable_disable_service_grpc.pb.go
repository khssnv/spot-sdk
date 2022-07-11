// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bosdyn/api/ir_enable_disable_service.proto

package api

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

// IREnableDisableServiceClient is the client API for IREnableDisableService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IREnableDisableServiceClient interface {
	IREnableDisable(ctx context.Context, in *IREnableDisableRequest, opts ...grpc.CallOption) (*IREnableDisableResponse, error)
}

type iREnableDisableServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIREnableDisableServiceClient(cc grpc.ClientConnInterface) IREnableDisableServiceClient {
	return &iREnableDisableServiceClient{cc}
}

func (c *iREnableDisableServiceClient) IREnableDisable(ctx context.Context, in *IREnableDisableRequest, opts ...grpc.CallOption) (*IREnableDisableResponse, error) {
	out := new(IREnableDisableResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.IREnableDisableService/IREnableDisable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IREnableDisableServiceServer is the server API for IREnableDisableService service.
// All implementations must embed UnimplementedIREnableDisableServiceServer
// for forward compatibility
type IREnableDisableServiceServer interface {
	IREnableDisable(context.Context, *IREnableDisableRequest) (*IREnableDisableResponse, error)
	mustEmbedUnimplementedIREnableDisableServiceServer()
}

// UnimplementedIREnableDisableServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIREnableDisableServiceServer struct {
}

func (UnimplementedIREnableDisableServiceServer) IREnableDisable(context.Context, *IREnableDisableRequest) (*IREnableDisableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IREnableDisable not implemented")
}
func (UnimplementedIREnableDisableServiceServer) mustEmbedUnimplementedIREnableDisableServiceServer() {
}

// UnsafeIREnableDisableServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IREnableDisableServiceServer will
// result in compilation errors.
type UnsafeIREnableDisableServiceServer interface {
	mustEmbedUnimplementedIREnableDisableServiceServer()
}

func RegisterIREnableDisableServiceServer(s grpc.ServiceRegistrar, srv IREnableDisableServiceServer) {
	s.RegisterService(&IREnableDisableService_ServiceDesc, srv)
}

func _IREnableDisableService_IREnableDisable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IREnableDisableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IREnableDisableServiceServer).IREnableDisable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.IREnableDisableService/IREnableDisable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IREnableDisableServiceServer).IREnableDisable(ctx, req.(*IREnableDisableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IREnableDisableService_ServiceDesc is the grpc.ServiceDesc for IREnableDisableService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IREnableDisableService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bosdyn.api.IREnableDisableService",
	HandlerType: (*IREnableDisableServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IREnableDisable",
			Handler:    _IREnableDisableService_IREnableDisable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bosdyn/api/ir_enable_disable_service.proto",
}