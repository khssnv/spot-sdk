// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bosdyn/api/auto_return/auto_return_service.proto

package auto_return

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

// AutoReturnServiceClient is the client API for AutoReturnService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AutoReturnServiceClient interface {
	// Configure the service.
	Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*ConfigureResponse, error)
	// Get the current configuration.
	GetConfiguration(ctx context.Context, in *GetConfigurationRequest, opts ...grpc.CallOption) (*GetConfigurationResponse, error)
	// Start AutoReturn now.
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
}

type autoReturnServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAutoReturnServiceClient(cc grpc.ClientConnInterface) AutoReturnServiceClient {
	return &autoReturnServiceClient{cc}
}

func (c *autoReturnServiceClient) Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*ConfigureResponse, error) {
	out := new(ConfigureResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.auto_return.AutoReturnService/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoReturnServiceClient) GetConfiguration(ctx context.Context, in *GetConfigurationRequest, opts ...grpc.CallOption) (*GetConfigurationResponse, error) {
	out := new(GetConfigurationResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.auto_return.AutoReturnService/GetConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoReturnServiceClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.auto_return.AutoReturnService/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AutoReturnServiceServer is the server API for AutoReturnService service.
// All implementations must embed UnimplementedAutoReturnServiceServer
// for forward compatibility
type AutoReturnServiceServer interface {
	// Configure the service.
	Configure(context.Context, *ConfigureRequest) (*ConfigureResponse, error)
	// Get the current configuration.
	GetConfiguration(context.Context, *GetConfigurationRequest) (*GetConfigurationResponse, error)
	// Start AutoReturn now.
	Start(context.Context, *StartRequest) (*StartResponse, error)
	mustEmbedUnimplementedAutoReturnServiceServer()
}

// UnimplementedAutoReturnServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAutoReturnServiceServer struct {
}

func (UnimplementedAutoReturnServiceServer) Configure(context.Context, *ConfigureRequest) (*ConfigureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (UnimplementedAutoReturnServiceServer) GetConfiguration(context.Context, *GetConfigurationRequest) (*GetConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfiguration not implemented")
}
func (UnimplementedAutoReturnServiceServer) Start(context.Context, *StartRequest) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedAutoReturnServiceServer) mustEmbedUnimplementedAutoReturnServiceServer() {}

// UnsafeAutoReturnServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AutoReturnServiceServer will
// result in compilation errors.
type UnsafeAutoReturnServiceServer interface {
	mustEmbedUnimplementedAutoReturnServiceServer()
}

func RegisterAutoReturnServiceServer(s grpc.ServiceRegistrar, srv AutoReturnServiceServer) {
	s.RegisterService(&AutoReturnService_ServiceDesc, srv)
}

func _AutoReturnService_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoReturnServiceServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.auto_return.AutoReturnService/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoReturnServiceServer).Configure(ctx, req.(*ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoReturnService_GetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoReturnServiceServer).GetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.auto_return.AutoReturnService/GetConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoReturnServiceServer).GetConfiguration(ctx, req.(*GetConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoReturnService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoReturnServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.auto_return.AutoReturnService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoReturnServiceServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AutoReturnService_ServiceDesc is the grpc.ServiceDesc for AutoReturnService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AutoReturnService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bosdyn.api.auto_return.AutoReturnService",
	HandlerType: (*AutoReturnServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _AutoReturnService_Configure_Handler,
		},
		{
			MethodName: "GetConfiguration",
			Handler:    _AutoReturnService_GetConfiguration_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _AutoReturnService_Start_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bosdyn/api/auto_return/auto_return_service.proto",
}