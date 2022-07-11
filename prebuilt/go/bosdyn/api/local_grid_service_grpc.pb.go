// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bosdyn/api/local_grid_service.proto

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

// LocalGridServiceClient is the client API for LocalGridService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocalGridServiceClient interface {
	// Obtain the list of available map types.
	// The name field keys access to individual local grids when calling GetLocalGrids.
	GetLocalGridTypes(ctx context.Context, in *GetLocalGridTypesRequest, opts ...grpc.CallOption) (*GetLocalGridTypesResponse, error)
	// Request a set of local grids by type name.
	GetLocalGrids(ctx context.Context, in *GetLocalGridsRequest, opts ...grpc.CallOption) (*GetLocalGridsResponse, error)
}

type localGridServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocalGridServiceClient(cc grpc.ClientConnInterface) LocalGridServiceClient {
	return &localGridServiceClient{cc}
}

func (c *localGridServiceClient) GetLocalGridTypes(ctx context.Context, in *GetLocalGridTypesRequest, opts ...grpc.CallOption) (*GetLocalGridTypesResponse, error) {
	out := new(GetLocalGridTypesResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.LocalGridService/GetLocalGridTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localGridServiceClient) GetLocalGrids(ctx context.Context, in *GetLocalGridsRequest, opts ...grpc.CallOption) (*GetLocalGridsResponse, error) {
	out := new(GetLocalGridsResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.LocalGridService/GetLocalGrids", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocalGridServiceServer is the server API for LocalGridService service.
// All implementations must embed UnimplementedLocalGridServiceServer
// for forward compatibility
type LocalGridServiceServer interface {
	// Obtain the list of available map types.
	// The name field keys access to individual local grids when calling GetLocalGrids.
	GetLocalGridTypes(context.Context, *GetLocalGridTypesRequest) (*GetLocalGridTypesResponse, error)
	// Request a set of local grids by type name.
	GetLocalGrids(context.Context, *GetLocalGridsRequest) (*GetLocalGridsResponse, error)
	mustEmbedUnimplementedLocalGridServiceServer()
}

// UnimplementedLocalGridServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLocalGridServiceServer struct {
}

func (UnimplementedLocalGridServiceServer) GetLocalGridTypes(context.Context, *GetLocalGridTypesRequest) (*GetLocalGridTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocalGridTypes not implemented")
}
func (UnimplementedLocalGridServiceServer) GetLocalGrids(context.Context, *GetLocalGridsRequest) (*GetLocalGridsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocalGrids not implemented")
}
func (UnimplementedLocalGridServiceServer) mustEmbedUnimplementedLocalGridServiceServer() {}

// UnsafeLocalGridServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocalGridServiceServer will
// result in compilation errors.
type UnsafeLocalGridServiceServer interface {
	mustEmbedUnimplementedLocalGridServiceServer()
}

func RegisterLocalGridServiceServer(s grpc.ServiceRegistrar, srv LocalGridServiceServer) {
	s.RegisterService(&LocalGridService_ServiceDesc, srv)
}

func _LocalGridService_GetLocalGridTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocalGridTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalGridServiceServer).GetLocalGridTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.LocalGridService/GetLocalGridTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalGridServiceServer).GetLocalGridTypes(ctx, req.(*GetLocalGridTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocalGridService_GetLocalGrids_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocalGridsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalGridServiceServer).GetLocalGrids(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.LocalGridService/GetLocalGrids",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalGridServiceServer).GetLocalGrids(ctx, req.(*GetLocalGridsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LocalGridService_ServiceDesc is the grpc.ServiceDesc for LocalGridService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocalGridService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bosdyn.api.LocalGridService",
	HandlerType: (*LocalGridServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLocalGridTypes",
			Handler:    _LocalGridService_GetLocalGridTypes_Handler,
		},
		{
			MethodName: "GetLocalGrids",
			Handler:    _LocalGridService_GetLocalGrids_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bosdyn/api/local_grid_service.proto",
}