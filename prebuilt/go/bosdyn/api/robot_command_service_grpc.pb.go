// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bosdyn/api/robot_command_service.proto

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

// RobotCommandServiceClient is the client API for RobotCommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RobotCommandServiceClient interface {
	// Starts a behavior command on the robot. Issuing a new command overrides the active command.
	// Each command is issued a UID for feedback retrieval.
	RobotCommand(ctx context.Context, in *RobotCommandRequest, opts ...grpc.CallOption) (*RobotCommandResponse, error)
	// A client queries this RPC to determine a robot's progress towards completion of a command.
	// This updates the client with metrics like "distance to goal."
	// The client should use this feedback to determine whether the current command has
	// succeeeded or failed, and thus send the next command.
	RobotCommandFeedback(ctx context.Context, in *RobotCommandFeedbackRequest, opts ...grpc.CallOption) (*RobotCommandFeedbackResponse, error)
	// Clear robot behavior fault.
	ClearBehaviorFault(ctx context.Context, in *ClearBehaviorFaultRequest, opts ...grpc.CallOption) (*ClearBehaviorFaultResponse, error)
}

type robotCommandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRobotCommandServiceClient(cc grpc.ClientConnInterface) RobotCommandServiceClient {
	return &robotCommandServiceClient{cc}
}

func (c *robotCommandServiceClient) RobotCommand(ctx context.Context, in *RobotCommandRequest, opts ...grpc.CallOption) (*RobotCommandResponse, error) {
	out := new(RobotCommandResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.RobotCommandService/RobotCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotCommandServiceClient) RobotCommandFeedback(ctx context.Context, in *RobotCommandFeedbackRequest, opts ...grpc.CallOption) (*RobotCommandFeedbackResponse, error) {
	out := new(RobotCommandFeedbackResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.RobotCommandService/RobotCommandFeedback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotCommandServiceClient) ClearBehaviorFault(ctx context.Context, in *ClearBehaviorFaultRequest, opts ...grpc.CallOption) (*ClearBehaviorFaultResponse, error) {
	out := new(ClearBehaviorFaultResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.RobotCommandService/ClearBehaviorFault", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RobotCommandServiceServer is the server API for RobotCommandService service.
// All implementations must embed UnimplementedRobotCommandServiceServer
// for forward compatibility
type RobotCommandServiceServer interface {
	// Starts a behavior command on the robot. Issuing a new command overrides the active command.
	// Each command is issued a UID for feedback retrieval.
	RobotCommand(context.Context, *RobotCommandRequest) (*RobotCommandResponse, error)
	// A client queries this RPC to determine a robot's progress towards completion of a command.
	// This updates the client with metrics like "distance to goal."
	// The client should use this feedback to determine whether the current command has
	// succeeeded or failed, and thus send the next command.
	RobotCommandFeedback(context.Context, *RobotCommandFeedbackRequest) (*RobotCommandFeedbackResponse, error)
	// Clear robot behavior fault.
	ClearBehaviorFault(context.Context, *ClearBehaviorFaultRequest) (*ClearBehaviorFaultResponse, error)
	mustEmbedUnimplementedRobotCommandServiceServer()
}

// UnimplementedRobotCommandServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRobotCommandServiceServer struct {
}

func (UnimplementedRobotCommandServiceServer) RobotCommand(context.Context, *RobotCommandRequest) (*RobotCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RobotCommand not implemented")
}
func (UnimplementedRobotCommandServiceServer) RobotCommandFeedback(context.Context, *RobotCommandFeedbackRequest) (*RobotCommandFeedbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RobotCommandFeedback not implemented")
}
func (UnimplementedRobotCommandServiceServer) ClearBehaviorFault(context.Context, *ClearBehaviorFaultRequest) (*ClearBehaviorFaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearBehaviorFault not implemented")
}
func (UnimplementedRobotCommandServiceServer) mustEmbedUnimplementedRobotCommandServiceServer() {}

// UnsafeRobotCommandServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RobotCommandServiceServer will
// result in compilation errors.
type UnsafeRobotCommandServiceServer interface {
	mustEmbedUnimplementedRobotCommandServiceServer()
}

func RegisterRobotCommandServiceServer(s grpc.ServiceRegistrar, srv RobotCommandServiceServer) {
	s.RegisterService(&RobotCommandService_ServiceDesc, srv)
}

func _RobotCommandService_RobotCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RobotCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotCommandServiceServer).RobotCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.RobotCommandService/RobotCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotCommandServiceServer).RobotCommand(ctx, req.(*RobotCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotCommandService_RobotCommandFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RobotCommandFeedbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotCommandServiceServer).RobotCommandFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.RobotCommandService/RobotCommandFeedback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotCommandServiceServer).RobotCommandFeedback(ctx, req.(*RobotCommandFeedbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotCommandService_ClearBehaviorFault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearBehaviorFaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotCommandServiceServer).ClearBehaviorFault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.RobotCommandService/ClearBehaviorFault",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotCommandServiceServer).ClearBehaviorFault(ctx, req.(*ClearBehaviorFaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RobotCommandService_ServiceDesc is the grpc.ServiceDesc for RobotCommandService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RobotCommandService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bosdyn.api.RobotCommandService",
	HandlerType: (*RobotCommandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RobotCommand",
			Handler:    _RobotCommandService_RobotCommand_Handler,
		},
		{
			MethodName: "RobotCommandFeedback",
			Handler:    _RobotCommandService_RobotCommandFeedback_Handler,
		},
		{
			MethodName: "ClearBehaviorFault",
			Handler:    _RobotCommandService_ClearBehaviorFault_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bosdyn/api/robot_command_service.proto",
}