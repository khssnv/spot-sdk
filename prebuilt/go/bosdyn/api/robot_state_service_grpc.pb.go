// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bosdyn/api/robot_state_service.proto

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

// RobotStateServiceClient is the client API for RobotStateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RobotStateServiceClient interface {
	// Get robot state information (such as kinematic state, power state, or faults).
	GetRobotState(ctx context.Context, in *RobotStateRequest, opts ...grpc.CallOption) (*RobotStateResponse, error)
	// Get different robot metrics and parameters from the robot.
	GetRobotMetrics(ctx context.Context, in *RobotMetricsRequest, opts ...grpc.CallOption) (*RobotMetricsResponse, error)
	// Get the hardware configuration of the robot, which describes the robot skeleton and urdf.
	GetRobotHardwareConfiguration(ctx context.Context, in *RobotHardwareConfigurationRequest, opts ...grpc.CallOption) (*RobotHardwareConfigurationResponse, error)
	// Returns the OBJ file for a specifc robot link. Intended to be called after
	// GetRobotHardwareConfiguration, using the link names returned by that call.
	GetRobotLinkModel(ctx context.Context, in *RobotLinkModelRequest, opts ...grpc.CallOption) (*RobotLinkModelResponse, error)
}

type robotStateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRobotStateServiceClient(cc grpc.ClientConnInterface) RobotStateServiceClient {
	return &robotStateServiceClient{cc}
}

func (c *robotStateServiceClient) GetRobotState(ctx context.Context, in *RobotStateRequest, opts ...grpc.CallOption) (*RobotStateResponse, error) {
	out := new(RobotStateResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.RobotStateService/GetRobotState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotStateServiceClient) GetRobotMetrics(ctx context.Context, in *RobotMetricsRequest, opts ...grpc.CallOption) (*RobotMetricsResponse, error) {
	out := new(RobotMetricsResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.RobotStateService/GetRobotMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotStateServiceClient) GetRobotHardwareConfiguration(ctx context.Context, in *RobotHardwareConfigurationRequest, opts ...grpc.CallOption) (*RobotHardwareConfigurationResponse, error) {
	out := new(RobotHardwareConfigurationResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.RobotStateService/GetRobotHardwareConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotStateServiceClient) GetRobotLinkModel(ctx context.Context, in *RobotLinkModelRequest, opts ...grpc.CallOption) (*RobotLinkModelResponse, error) {
	out := new(RobotLinkModelResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.RobotStateService/GetRobotLinkModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RobotStateServiceServer is the server API for RobotStateService service.
// All implementations must embed UnimplementedRobotStateServiceServer
// for forward compatibility
type RobotStateServiceServer interface {
	// Get robot state information (such as kinematic state, power state, or faults).
	GetRobotState(context.Context, *RobotStateRequest) (*RobotStateResponse, error)
	// Get different robot metrics and parameters from the robot.
	GetRobotMetrics(context.Context, *RobotMetricsRequest) (*RobotMetricsResponse, error)
	// Get the hardware configuration of the robot, which describes the robot skeleton and urdf.
	GetRobotHardwareConfiguration(context.Context, *RobotHardwareConfigurationRequest) (*RobotHardwareConfigurationResponse, error)
	// Returns the OBJ file for a specifc robot link. Intended to be called after
	// GetRobotHardwareConfiguration, using the link names returned by that call.
	GetRobotLinkModel(context.Context, *RobotLinkModelRequest) (*RobotLinkModelResponse, error)
	mustEmbedUnimplementedRobotStateServiceServer()
}

// UnimplementedRobotStateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRobotStateServiceServer struct {
}

func (UnimplementedRobotStateServiceServer) GetRobotState(context.Context, *RobotStateRequest) (*RobotStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRobotState not implemented")
}
func (UnimplementedRobotStateServiceServer) GetRobotMetrics(context.Context, *RobotMetricsRequest) (*RobotMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRobotMetrics not implemented")
}
func (UnimplementedRobotStateServiceServer) GetRobotHardwareConfiguration(context.Context, *RobotHardwareConfigurationRequest) (*RobotHardwareConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRobotHardwareConfiguration not implemented")
}
func (UnimplementedRobotStateServiceServer) GetRobotLinkModel(context.Context, *RobotLinkModelRequest) (*RobotLinkModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRobotLinkModel not implemented")
}
func (UnimplementedRobotStateServiceServer) mustEmbedUnimplementedRobotStateServiceServer() {}

// UnsafeRobotStateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RobotStateServiceServer will
// result in compilation errors.
type UnsafeRobotStateServiceServer interface {
	mustEmbedUnimplementedRobotStateServiceServer()
}

func RegisterRobotStateServiceServer(s grpc.ServiceRegistrar, srv RobotStateServiceServer) {
	s.RegisterService(&RobotStateService_ServiceDesc, srv)
}

func _RobotStateService_GetRobotState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RobotStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotStateServiceServer).GetRobotState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.RobotStateService/GetRobotState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotStateServiceServer).GetRobotState(ctx, req.(*RobotStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotStateService_GetRobotMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RobotMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotStateServiceServer).GetRobotMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.RobotStateService/GetRobotMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotStateServiceServer).GetRobotMetrics(ctx, req.(*RobotMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotStateService_GetRobotHardwareConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RobotHardwareConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotStateServiceServer).GetRobotHardwareConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.RobotStateService/GetRobotHardwareConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotStateServiceServer).GetRobotHardwareConfiguration(ctx, req.(*RobotHardwareConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotStateService_GetRobotLinkModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RobotLinkModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotStateServiceServer).GetRobotLinkModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.RobotStateService/GetRobotLinkModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotStateServiceServer).GetRobotLinkModel(ctx, req.(*RobotLinkModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RobotStateService_ServiceDesc is the grpc.ServiceDesc for RobotStateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RobotStateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bosdyn.api.RobotStateService",
	HandlerType: (*RobotStateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRobotState",
			Handler:    _RobotStateService_GetRobotState_Handler,
		},
		{
			MethodName: "GetRobotMetrics",
			Handler:    _RobotStateService_GetRobotMetrics_Handler,
		},
		{
			MethodName: "GetRobotHardwareConfiguration",
			Handler:    _RobotStateService_GetRobotHardwareConfiguration_Handler,
		},
		{
			MethodName: "GetRobotLinkModel",
			Handler:    _RobotStateService_GetRobotLinkModel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bosdyn/api/robot_state_service.proto",
}
