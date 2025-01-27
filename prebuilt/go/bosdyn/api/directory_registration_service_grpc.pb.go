// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bosdyn/api/directory_registration_service.proto

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

// DirectoryRegistrationServiceClient is the client API for DirectoryRegistrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DirectoryRegistrationServiceClient interface {
	// Called by a producer to register as a provider with the application.  Returns the
	// record for that provider.  Requires unique name and correctly filled out service
	// record in request.
	RegisterService(ctx context.Context, in *RegisterServiceRequest, opts ...grpc.CallOption) (*RegisterServiceResponse, error)
	// Called by a producer to remove its registration from the DirectoryManager.
	UnregisterService(ctx context.Context, in *UnregisterServiceRequest, opts ...grpc.CallOption) (*UnregisterServiceResponse, error)
	// Update the ServiceEntry for a producer on the server.
	UpdateService(ctx context.Context, in *UpdateServiceRequest, opts ...grpc.CallOption) (*UpdateServiceResponse, error)
}

type directoryRegistrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDirectoryRegistrationServiceClient(cc grpc.ClientConnInterface) DirectoryRegistrationServiceClient {
	return &directoryRegistrationServiceClient{cc}
}

func (c *directoryRegistrationServiceClient) RegisterService(ctx context.Context, in *RegisterServiceRequest, opts ...grpc.CallOption) (*RegisterServiceResponse, error) {
	out := new(RegisterServiceResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.DirectoryRegistrationService/RegisterService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryRegistrationServiceClient) UnregisterService(ctx context.Context, in *UnregisterServiceRequest, opts ...grpc.CallOption) (*UnregisterServiceResponse, error) {
	out := new(UnregisterServiceResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.DirectoryRegistrationService/UnregisterService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryRegistrationServiceClient) UpdateService(ctx context.Context, in *UpdateServiceRequest, opts ...grpc.CallOption) (*UpdateServiceResponse, error) {
	out := new(UpdateServiceResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.DirectoryRegistrationService/UpdateService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DirectoryRegistrationServiceServer is the server API for DirectoryRegistrationService service.
// All implementations must embed UnimplementedDirectoryRegistrationServiceServer
// for forward compatibility
type DirectoryRegistrationServiceServer interface {
	// Called by a producer to register as a provider with the application.  Returns the
	// record for that provider.  Requires unique name and correctly filled out service
	// record in request.
	RegisterService(context.Context, *RegisterServiceRequest) (*RegisterServiceResponse, error)
	// Called by a producer to remove its registration from the DirectoryManager.
	UnregisterService(context.Context, *UnregisterServiceRequest) (*UnregisterServiceResponse, error)
	// Update the ServiceEntry for a producer on the server.
	UpdateService(context.Context, *UpdateServiceRequest) (*UpdateServiceResponse, error)
	mustEmbedUnimplementedDirectoryRegistrationServiceServer()
}

// UnimplementedDirectoryRegistrationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDirectoryRegistrationServiceServer struct {
}

func (UnimplementedDirectoryRegistrationServiceServer) RegisterService(context.Context, *RegisterServiceRequest) (*RegisterServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterService not implemented")
}
func (UnimplementedDirectoryRegistrationServiceServer) UnregisterService(context.Context, *UnregisterServiceRequest) (*UnregisterServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterService not implemented")
}
func (UnimplementedDirectoryRegistrationServiceServer) UpdateService(context.Context, *UpdateServiceRequest) (*UpdateServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateService not implemented")
}
func (UnimplementedDirectoryRegistrationServiceServer) mustEmbedUnimplementedDirectoryRegistrationServiceServer() {
}

// UnsafeDirectoryRegistrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DirectoryRegistrationServiceServer will
// result in compilation errors.
type UnsafeDirectoryRegistrationServiceServer interface {
	mustEmbedUnimplementedDirectoryRegistrationServiceServer()
}

func RegisterDirectoryRegistrationServiceServer(s grpc.ServiceRegistrar, srv DirectoryRegistrationServiceServer) {
	s.RegisterService(&DirectoryRegistrationService_ServiceDesc, srv)
}

func _DirectoryRegistrationService_RegisterService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryRegistrationServiceServer).RegisterService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.DirectoryRegistrationService/RegisterService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryRegistrationServiceServer).RegisterService(ctx, req.(*RegisterServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryRegistrationService_UnregisterService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryRegistrationServiceServer).UnregisterService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.DirectoryRegistrationService/UnregisterService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryRegistrationServiceServer).UnregisterService(ctx, req.(*UnregisterServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryRegistrationService_UpdateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryRegistrationServiceServer).UpdateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.DirectoryRegistrationService/UpdateService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryRegistrationServiceServer).UpdateService(ctx, req.(*UpdateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DirectoryRegistrationService_ServiceDesc is the grpc.ServiceDesc for DirectoryRegistrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DirectoryRegistrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bosdyn.api.DirectoryRegistrationService",
	HandlerType: (*DirectoryRegistrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterService",
			Handler:    _DirectoryRegistrationService_RegisterService_Handler,
		},
		{
			MethodName: "UnregisterService",
			Handler:    _DirectoryRegistrationService_UnregisterService_Handler,
		},
		{
			MethodName: "UpdateService",
			Handler:    _DirectoryRegistrationService_UpdateService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bosdyn/api/directory_registration_service.proto",
}
