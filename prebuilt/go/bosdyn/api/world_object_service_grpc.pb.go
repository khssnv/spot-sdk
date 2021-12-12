// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bosdyn/api/world_object_service.proto

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

// WorldObjectServiceClient is the client API for WorldObjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorldObjectServiceClient interface {
	// Request a list of all the world objects in the robot's perception scene.
	ListWorldObjects(ctx context.Context, in *ListWorldObjectRequest, opts ...grpc.CallOption) (*ListWorldObjectResponse, error)
	// Mutate (add, change, or delete) the world objects.
	MutateWorldObjects(ctx context.Context, in *MutateWorldObjectRequest, opts ...grpc.CallOption) (*MutateWorldObjectResponse, error)
}

type worldObjectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorldObjectServiceClient(cc grpc.ClientConnInterface) WorldObjectServiceClient {
	return &worldObjectServiceClient{cc}
}

func (c *worldObjectServiceClient) ListWorldObjects(ctx context.Context, in *ListWorldObjectRequest, opts ...grpc.CallOption) (*ListWorldObjectResponse, error) {
	out := new(ListWorldObjectResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.WorldObjectService/ListWorldObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *worldObjectServiceClient) MutateWorldObjects(ctx context.Context, in *MutateWorldObjectRequest, opts ...grpc.CallOption) (*MutateWorldObjectResponse, error) {
	out := new(MutateWorldObjectResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.WorldObjectService/MutateWorldObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorldObjectServiceServer is the server API for WorldObjectService service.
// All implementations must embed UnimplementedWorldObjectServiceServer
// for forward compatibility
type WorldObjectServiceServer interface {
	// Request a list of all the world objects in the robot's perception scene.
	ListWorldObjects(context.Context, *ListWorldObjectRequest) (*ListWorldObjectResponse, error)
	// Mutate (add, change, or delete) the world objects.
	MutateWorldObjects(context.Context, *MutateWorldObjectRequest) (*MutateWorldObjectResponse, error)
	mustEmbedUnimplementedWorldObjectServiceServer()
}

// UnimplementedWorldObjectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorldObjectServiceServer struct {
}

func (UnimplementedWorldObjectServiceServer) ListWorldObjects(context.Context, *ListWorldObjectRequest) (*ListWorldObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorldObjects not implemented")
}
func (UnimplementedWorldObjectServiceServer) MutateWorldObjects(context.Context, *MutateWorldObjectRequest) (*MutateWorldObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateWorldObjects not implemented")
}
func (UnimplementedWorldObjectServiceServer) mustEmbedUnimplementedWorldObjectServiceServer() {}

// UnsafeWorldObjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorldObjectServiceServer will
// result in compilation errors.
type UnsafeWorldObjectServiceServer interface {
	mustEmbedUnimplementedWorldObjectServiceServer()
}

func RegisterWorldObjectServiceServer(s grpc.ServiceRegistrar, srv WorldObjectServiceServer) {
	s.RegisterService(&WorldObjectService_ServiceDesc, srv)
}

func _WorldObjectService_ListWorldObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorldObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorldObjectServiceServer).ListWorldObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.WorldObjectService/ListWorldObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorldObjectServiceServer).ListWorldObjects(ctx, req.(*ListWorldObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorldObjectService_MutateWorldObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateWorldObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorldObjectServiceServer).MutateWorldObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.WorldObjectService/MutateWorldObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorldObjectServiceServer).MutateWorldObjects(ctx, req.(*MutateWorldObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorldObjectService_ServiceDesc is the grpc.ServiceDesc for WorldObjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorldObjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bosdyn.api.WorldObjectService",
	HandlerType: (*WorldObjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListWorldObjects",
			Handler:    _WorldObjectService_ListWorldObjects_Handler,
		},
		{
			MethodName: "MutateWorldObjects",
			Handler:    _WorldObjectService_MutateWorldObjects_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bosdyn/api/world_object_service.proto",
}
