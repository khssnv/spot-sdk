// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bosdyn/api/payload_registration_service.proto

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

// PayloadRegistrationServiceClient is the client API for PayloadRegistrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PayloadRegistrationServiceClient interface {
	// Register a payload with the directory.
	RegisterPayload(ctx context.Context, in *RegisterPayloadRequest, opts ...grpc.CallOption) (*RegisterPayloadResponse, error)
	// Update the version for the registered payload.
	UpdatePayloadVersion(ctx context.Context, in *UpdatePayloadVersionRequest, opts ...grpc.CallOption) (*UpdatePayloadVersionResponse, error)
	// Get the authentication token information associated with a given payload.
	GetPayloadAuthToken(ctx context.Context, in *GetPayloadAuthTokenRequest, opts ...grpc.CallOption) (*GetPayloadAuthTokenResponse, error)
	// Tell the robot whether the specified payload is attached..
	UpdatePayloadAttached(ctx context.Context, in *UpdatePayloadAttachedRequest, opts ...grpc.CallOption) (*UpdatePayloadAttachedResponse, error)
}

type payloadRegistrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPayloadRegistrationServiceClient(cc grpc.ClientConnInterface) PayloadRegistrationServiceClient {
	return &payloadRegistrationServiceClient{cc}
}

func (c *payloadRegistrationServiceClient) RegisterPayload(ctx context.Context, in *RegisterPayloadRequest, opts ...grpc.CallOption) (*RegisterPayloadResponse, error) {
	out := new(RegisterPayloadResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.PayloadRegistrationService/RegisterPayload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payloadRegistrationServiceClient) UpdatePayloadVersion(ctx context.Context, in *UpdatePayloadVersionRequest, opts ...grpc.CallOption) (*UpdatePayloadVersionResponse, error) {
	out := new(UpdatePayloadVersionResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.PayloadRegistrationService/UpdatePayloadVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payloadRegistrationServiceClient) GetPayloadAuthToken(ctx context.Context, in *GetPayloadAuthTokenRequest, opts ...grpc.CallOption) (*GetPayloadAuthTokenResponse, error) {
	out := new(GetPayloadAuthTokenResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.PayloadRegistrationService/GetPayloadAuthToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payloadRegistrationServiceClient) UpdatePayloadAttached(ctx context.Context, in *UpdatePayloadAttachedRequest, opts ...grpc.CallOption) (*UpdatePayloadAttachedResponse, error) {
	out := new(UpdatePayloadAttachedResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.PayloadRegistrationService/UpdatePayloadAttached", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayloadRegistrationServiceServer is the server API for PayloadRegistrationService service.
// All implementations must embed UnimplementedPayloadRegistrationServiceServer
// for forward compatibility
type PayloadRegistrationServiceServer interface {
	// Register a payload with the directory.
	RegisterPayload(context.Context, *RegisterPayloadRequest) (*RegisterPayloadResponse, error)
	// Update the version for the registered payload.
	UpdatePayloadVersion(context.Context, *UpdatePayloadVersionRequest) (*UpdatePayloadVersionResponse, error)
	// Get the authentication token information associated with a given payload.
	GetPayloadAuthToken(context.Context, *GetPayloadAuthTokenRequest) (*GetPayloadAuthTokenResponse, error)
	// Tell the robot whether the specified payload is attached..
	UpdatePayloadAttached(context.Context, *UpdatePayloadAttachedRequest) (*UpdatePayloadAttachedResponse, error)
	mustEmbedUnimplementedPayloadRegistrationServiceServer()
}

// UnimplementedPayloadRegistrationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPayloadRegistrationServiceServer struct {
}

func (UnimplementedPayloadRegistrationServiceServer) RegisterPayload(context.Context, *RegisterPayloadRequest) (*RegisterPayloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterPayload not implemented")
}
func (UnimplementedPayloadRegistrationServiceServer) UpdatePayloadVersion(context.Context, *UpdatePayloadVersionRequest) (*UpdatePayloadVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePayloadVersion not implemented")
}
func (UnimplementedPayloadRegistrationServiceServer) GetPayloadAuthToken(context.Context, *GetPayloadAuthTokenRequest) (*GetPayloadAuthTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayloadAuthToken not implemented")
}
func (UnimplementedPayloadRegistrationServiceServer) UpdatePayloadAttached(context.Context, *UpdatePayloadAttachedRequest) (*UpdatePayloadAttachedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePayloadAttached not implemented")
}
func (UnimplementedPayloadRegistrationServiceServer) mustEmbedUnimplementedPayloadRegistrationServiceServer() {
}

// UnsafePayloadRegistrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PayloadRegistrationServiceServer will
// result in compilation errors.
type UnsafePayloadRegistrationServiceServer interface {
	mustEmbedUnimplementedPayloadRegistrationServiceServer()
}

func RegisterPayloadRegistrationServiceServer(s grpc.ServiceRegistrar, srv PayloadRegistrationServiceServer) {
	s.RegisterService(&PayloadRegistrationService_ServiceDesc, srv)
}

func _PayloadRegistrationService_RegisterPayload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterPayloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayloadRegistrationServiceServer).RegisterPayload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.PayloadRegistrationService/RegisterPayload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayloadRegistrationServiceServer).RegisterPayload(ctx, req.(*RegisterPayloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayloadRegistrationService_UpdatePayloadVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePayloadVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayloadRegistrationServiceServer).UpdatePayloadVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.PayloadRegistrationService/UpdatePayloadVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayloadRegistrationServiceServer).UpdatePayloadVersion(ctx, req.(*UpdatePayloadVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayloadRegistrationService_GetPayloadAuthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPayloadAuthTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayloadRegistrationServiceServer).GetPayloadAuthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.PayloadRegistrationService/GetPayloadAuthToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayloadRegistrationServiceServer).GetPayloadAuthToken(ctx, req.(*GetPayloadAuthTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayloadRegistrationService_UpdatePayloadAttached_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePayloadAttachedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayloadRegistrationServiceServer).UpdatePayloadAttached(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.PayloadRegistrationService/UpdatePayloadAttached",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayloadRegistrationServiceServer).UpdatePayloadAttached(ctx, req.(*UpdatePayloadAttachedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PayloadRegistrationService_ServiceDesc is the grpc.ServiceDesc for PayloadRegistrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PayloadRegistrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bosdyn.api.PayloadRegistrationService",
	HandlerType: (*PayloadRegistrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterPayload",
			Handler:    _PayloadRegistrationService_RegisterPayload_Handler,
		},
		{
			MethodName: "UpdatePayloadVersion",
			Handler:    _PayloadRegistrationService_UpdatePayloadVersion_Handler,
		},
		{
			MethodName: "GetPayloadAuthToken",
			Handler:    _PayloadRegistrationService_GetPayloadAuthToken_Handler,
		},
		{
			MethodName: "UpdatePayloadAttached",
			Handler:    _PayloadRegistrationService_UpdatePayloadAttached_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bosdyn/api/payload_registration_service.proto",
}
