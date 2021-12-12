// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bosdyn/api/estop_service.proto

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

// EstopServiceClient is the client API for EstopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EstopServiceClient interface {
	// Register an Estop "originator" or "endpoint".
	// This may be a replacement for another active endpoint.
	RegisterEstopEndpoint(ctx context.Context, in *RegisterEstopEndpointRequest, opts ...grpc.CallOption) (*RegisterEstopEndpointResponse, error)
	// Deregister the requested estop endpoint.
	DeregisterEstopEndpoint(ctx context.Context, in *DeregisterEstopEndpointRequest, opts ...grpc.CallOption) (*DeregisterEstopEndpointResponse, error)
	// Answer challenge from previous response (unless this is the first call), and request
	// a stop level.
	EstopCheckIn(ctx context.Context, in *EstopCheckInRequest, opts ...grpc.CallOption) (*EstopCheckInResponse, error)
	// Request the current EstopConfig, describing the expected set of endpoints.
	GetEstopConfig(ctx context.Context, in *GetEstopConfigRequest, opts ...grpc.CallOption) (*GetEstopConfigResponse, error)
	// Set a new active EstopConfig.
	SetEstopConfig(ctx context.Context, in *SetEstopConfigRequest, opts ...grpc.CallOption) (*SetEstopConfigResponse, error)
	// Ask for the current status of the estop system.
	GetEstopSystemStatus(ctx context.Context, in *GetEstopSystemStatusRequest, opts ...grpc.CallOption) (*GetEstopSystemStatusResponse, error)
}

type estopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEstopServiceClient(cc grpc.ClientConnInterface) EstopServiceClient {
	return &estopServiceClient{cc}
}

func (c *estopServiceClient) RegisterEstopEndpoint(ctx context.Context, in *RegisterEstopEndpointRequest, opts ...grpc.CallOption) (*RegisterEstopEndpointResponse, error) {
	out := new(RegisterEstopEndpointResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.EstopService/RegisterEstopEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *estopServiceClient) DeregisterEstopEndpoint(ctx context.Context, in *DeregisterEstopEndpointRequest, opts ...grpc.CallOption) (*DeregisterEstopEndpointResponse, error) {
	out := new(DeregisterEstopEndpointResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.EstopService/DeregisterEstopEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *estopServiceClient) EstopCheckIn(ctx context.Context, in *EstopCheckInRequest, opts ...grpc.CallOption) (*EstopCheckInResponse, error) {
	out := new(EstopCheckInResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.EstopService/EstopCheckIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *estopServiceClient) GetEstopConfig(ctx context.Context, in *GetEstopConfigRequest, opts ...grpc.CallOption) (*GetEstopConfigResponse, error) {
	out := new(GetEstopConfigResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.EstopService/GetEstopConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *estopServiceClient) SetEstopConfig(ctx context.Context, in *SetEstopConfigRequest, opts ...grpc.CallOption) (*SetEstopConfigResponse, error) {
	out := new(SetEstopConfigResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.EstopService/SetEstopConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *estopServiceClient) GetEstopSystemStatus(ctx context.Context, in *GetEstopSystemStatusRequest, opts ...grpc.CallOption) (*GetEstopSystemStatusResponse, error) {
	out := new(GetEstopSystemStatusResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.EstopService/GetEstopSystemStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EstopServiceServer is the server API for EstopService service.
// All implementations must embed UnimplementedEstopServiceServer
// for forward compatibility
type EstopServiceServer interface {
	// Register an Estop "originator" or "endpoint".
	// This may be a replacement for another active endpoint.
	RegisterEstopEndpoint(context.Context, *RegisterEstopEndpointRequest) (*RegisterEstopEndpointResponse, error)
	// Deregister the requested estop endpoint.
	DeregisterEstopEndpoint(context.Context, *DeregisterEstopEndpointRequest) (*DeregisterEstopEndpointResponse, error)
	// Answer challenge from previous response (unless this is the first call), and request
	// a stop level.
	EstopCheckIn(context.Context, *EstopCheckInRequest) (*EstopCheckInResponse, error)
	// Request the current EstopConfig, describing the expected set of endpoints.
	GetEstopConfig(context.Context, *GetEstopConfigRequest) (*GetEstopConfigResponse, error)
	// Set a new active EstopConfig.
	SetEstopConfig(context.Context, *SetEstopConfigRequest) (*SetEstopConfigResponse, error)
	// Ask for the current status of the estop system.
	GetEstopSystemStatus(context.Context, *GetEstopSystemStatusRequest) (*GetEstopSystemStatusResponse, error)
	mustEmbedUnimplementedEstopServiceServer()
}

// UnimplementedEstopServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEstopServiceServer struct {
}

func (UnimplementedEstopServiceServer) RegisterEstopEndpoint(context.Context, *RegisterEstopEndpointRequest) (*RegisterEstopEndpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterEstopEndpoint not implemented")
}
func (UnimplementedEstopServiceServer) DeregisterEstopEndpoint(context.Context, *DeregisterEstopEndpointRequest) (*DeregisterEstopEndpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeregisterEstopEndpoint not implemented")
}
func (UnimplementedEstopServiceServer) EstopCheckIn(context.Context, *EstopCheckInRequest) (*EstopCheckInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstopCheckIn not implemented")
}
func (UnimplementedEstopServiceServer) GetEstopConfig(context.Context, *GetEstopConfigRequest) (*GetEstopConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEstopConfig not implemented")
}
func (UnimplementedEstopServiceServer) SetEstopConfig(context.Context, *SetEstopConfigRequest) (*SetEstopConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetEstopConfig not implemented")
}
func (UnimplementedEstopServiceServer) GetEstopSystemStatus(context.Context, *GetEstopSystemStatusRequest) (*GetEstopSystemStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEstopSystemStatus not implemented")
}
func (UnimplementedEstopServiceServer) mustEmbedUnimplementedEstopServiceServer() {}

// UnsafeEstopServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EstopServiceServer will
// result in compilation errors.
type UnsafeEstopServiceServer interface {
	mustEmbedUnimplementedEstopServiceServer()
}

func RegisterEstopServiceServer(s grpc.ServiceRegistrar, srv EstopServiceServer) {
	s.RegisterService(&EstopService_ServiceDesc, srv)
}

func _EstopService_RegisterEstopEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterEstopEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EstopServiceServer).RegisterEstopEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.EstopService/RegisterEstopEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EstopServiceServer).RegisterEstopEndpoint(ctx, req.(*RegisterEstopEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EstopService_DeregisterEstopEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeregisterEstopEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EstopServiceServer).DeregisterEstopEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.EstopService/DeregisterEstopEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EstopServiceServer).DeregisterEstopEndpoint(ctx, req.(*DeregisterEstopEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EstopService_EstopCheckIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EstopCheckInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EstopServiceServer).EstopCheckIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.EstopService/EstopCheckIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EstopServiceServer).EstopCheckIn(ctx, req.(*EstopCheckInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EstopService_GetEstopConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEstopConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EstopServiceServer).GetEstopConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.EstopService/GetEstopConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EstopServiceServer).GetEstopConfig(ctx, req.(*GetEstopConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EstopService_SetEstopConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetEstopConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EstopServiceServer).SetEstopConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.EstopService/SetEstopConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EstopServiceServer).SetEstopConfig(ctx, req.(*SetEstopConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EstopService_GetEstopSystemStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEstopSystemStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EstopServiceServer).GetEstopSystemStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.EstopService/GetEstopSystemStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EstopServiceServer).GetEstopSystemStatus(ctx, req.(*GetEstopSystemStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EstopService_ServiceDesc is the grpc.ServiceDesc for EstopService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EstopService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bosdyn.api.EstopService",
	HandlerType: (*EstopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterEstopEndpoint",
			Handler:    _EstopService_RegisterEstopEndpoint_Handler,
		},
		{
			MethodName: "DeregisterEstopEndpoint",
			Handler:    _EstopService_DeregisterEstopEndpoint_Handler,
		},
		{
			MethodName: "EstopCheckIn",
			Handler:    _EstopService_EstopCheckIn_Handler,
		},
		{
			MethodName: "GetEstopConfig",
			Handler:    _EstopService_GetEstopConfig_Handler,
		},
		{
			MethodName: "SetEstopConfig",
			Handler:    _EstopService_SetEstopConfig_Handler,
		},
		{
			MethodName: "GetEstopSystemStatus",
			Handler:    _EstopService_GetEstopSystemStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bosdyn/api/estop_service.proto",
}
