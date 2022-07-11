// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bosdyn/api/docking/docking_service.proto

package docking

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

// DockingServiceClient is the client API for DockingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DockingServiceClient interface {
	// Starts a docking command on the robot.
	DockingCommand(ctx context.Context, in *DockingCommandRequest, opts ...grpc.CallOption) (*DockingCommandResponse, error)
	// Check the status of a docking command.
	DockingCommandFeedback(ctx context.Context, in *DockingCommandFeedbackRequest, opts ...grpc.CallOption) (*DockingCommandFeedbackResponse, error)
	// Get the configured dock ID ranges.
	GetDockingConfig(ctx context.Context, in *GetDockingConfigRequest, opts ...grpc.CallOption) (*GetDockingConfigResponse, error)
	// Get the robot's docking state
	GetDockingState(ctx context.Context, in *GetDockingStateRequest, opts ...grpc.CallOption) (*GetDockingStateResponse, error)
}

type dockingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDockingServiceClient(cc grpc.ClientConnInterface) DockingServiceClient {
	return &dockingServiceClient{cc}
}

func (c *dockingServiceClient) DockingCommand(ctx context.Context, in *DockingCommandRequest, opts ...grpc.CallOption) (*DockingCommandResponse, error) {
	out := new(DockingCommandResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.docking.DockingService/DockingCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockingServiceClient) DockingCommandFeedback(ctx context.Context, in *DockingCommandFeedbackRequest, opts ...grpc.CallOption) (*DockingCommandFeedbackResponse, error) {
	out := new(DockingCommandFeedbackResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.docking.DockingService/DockingCommandFeedback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockingServiceClient) GetDockingConfig(ctx context.Context, in *GetDockingConfigRequest, opts ...grpc.CallOption) (*GetDockingConfigResponse, error) {
	out := new(GetDockingConfigResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.docking.DockingService/GetDockingConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockingServiceClient) GetDockingState(ctx context.Context, in *GetDockingStateRequest, opts ...grpc.CallOption) (*GetDockingStateResponse, error) {
	out := new(GetDockingStateResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.docking.DockingService/GetDockingState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DockingServiceServer is the server API for DockingService service.
// All implementations must embed UnimplementedDockingServiceServer
// for forward compatibility
type DockingServiceServer interface {
	// Starts a docking command on the robot.
	DockingCommand(context.Context, *DockingCommandRequest) (*DockingCommandResponse, error)
	// Check the status of a docking command.
	DockingCommandFeedback(context.Context, *DockingCommandFeedbackRequest) (*DockingCommandFeedbackResponse, error)
	// Get the configured dock ID ranges.
	GetDockingConfig(context.Context, *GetDockingConfigRequest) (*GetDockingConfigResponse, error)
	// Get the robot's docking state
	GetDockingState(context.Context, *GetDockingStateRequest) (*GetDockingStateResponse, error)
	mustEmbedUnimplementedDockingServiceServer()
}

// UnimplementedDockingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDockingServiceServer struct {
}

func (UnimplementedDockingServiceServer) DockingCommand(context.Context, *DockingCommandRequest) (*DockingCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DockingCommand not implemented")
}
func (UnimplementedDockingServiceServer) DockingCommandFeedback(context.Context, *DockingCommandFeedbackRequest) (*DockingCommandFeedbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DockingCommandFeedback not implemented")
}
func (UnimplementedDockingServiceServer) GetDockingConfig(context.Context, *GetDockingConfigRequest) (*GetDockingConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDockingConfig not implemented")
}
func (UnimplementedDockingServiceServer) GetDockingState(context.Context, *GetDockingStateRequest) (*GetDockingStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDockingState not implemented")
}
func (UnimplementedDockingServiceServer) mustEmbedUnimplementedDockingServiceServer() {}

// UnsafeDockingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DockingServiceServer will
// result in compilation errors.
type UnsafeDockingServiceServer interface {
	mustEmbedUnimplementedDockingServiceServer()
}

func RegisterDockingServiceServer(s grpc.ServiceRegistrar, srv DockingServiceServer) {
	s.RegisterService(&DockingService_ServiceDesc, srv)
}

func _DockingService_DockingCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DockingCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockingServiceServer).DockingCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.docking.DockingService/DockingCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockingServiceServer).DockingCommand(ctx, req.(*DockingCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockingService_DockingCommandFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DockingCommandFeedbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockingServiceServer).DockingCommandFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.docking.DockingService/DockingCommandFeedback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockingServiceServer).DockingCommandFeedback(ctx, req.(*DockingCommandFeedbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockingService_GetDockingConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDockingConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockingServiceServer).GetDockingConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.docking.DockingService/GetDockingConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockingServiceServer).GetDockingConfig(ctx, req.(*GetDockingConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockingService_GetDockingState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDockingStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockingServiceServer).GetDockingState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.docking.DockingService/GetDockingState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockingServiceServer).GetDockingState(ctx, req.(*GetDockingStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DockingService_ServiceDesc is the grpc.ServiceDesc for DockingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DockingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bosdyn.api.docking.DockingService",
	HandlerType: (*DockingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DockingCommand",
			Handler:    _DockingService_DockingCommand_Handler,
		},
		{
			MethodName: "DockingCommandFeedback",
			Handler:    _DockingService_DockingCommandFeedback_Handler,
		},
		{
			MethodName: "GetDockingConfig",
			Handler:    _DockingService_GetDockingConfig_Handler,
		},
		{
			MethodName: "GetDockingState",
			Handler:    _DockingService_GetDockingState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bosdyn/api/docking/docking_service.proto",
}