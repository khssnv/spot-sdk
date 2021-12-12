// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bosdyn/api/spot/spot_check_service.proto

package spot

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

// SpotCheckServiceClient is the client API for SpotCheckService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpotCheckServiceClient interface {
	// Send a command to the SpotCheck service. The spotcheck service is responsible to both
	// recalibrating actuation sensors and checking camera health.
	SpotCheckCommand(ctx context.Context, in *SpotCheckCommandRequest, opts ...grpc.CallOption) (*SpotCheckCommandResponse, error)
	// Check the status of the spot check procedure. After procedure completes, this reports back
	// results for specific joints and cameras.
	SpotCheckFeedback(ctx context.Context, in *SpotCheckFeedbackRequest, opts ...grpc.CallOption) (*SpotCheckFeedbackResponse, error)
	// Send a camera calibration command to the robot. Used to start or abort a calibration routine.
	CameraCalibrationCommand(ctx context.Context, in *CameraCalibrationCommandRequest, opts ...grpc.CallOption) (*CameraCalibrationCommandResponse, error)
	// Check the status of the camera calibration procedure.
	CameraCalibrationFeedback(ctx context.Context, in *CameraCalibrationFeedbackRequest, opts ...grpc.CallOption) (*CameraCalibrationFeedbackResponse, error)
}

type spotCheckServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpotCheckServiceClient(cc grpc.ClientConnInterface) SpotCheckServiceClient {
	return &spotCheckServiceClient{cc}
}

func (c *spotCheckServiceClient) SpotCheckCommand(ctx context.Context, in *SpotCheckCommandRequest, opts ...grpc.CallOption) (*SpotCheckCommandResponse, error) {
	out := new(SpotCheckCommandResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.spot.SpotCheckService/SpotCheckCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spotCheckServiceClient) SpotCheckFeedback(ctx context.Context, in *SpotCheckFeedbackRequest, opts ...grpc.CallOption) (*SpotCheckFeedbackResponse, error) {
	out := new(SpotCheckFeedbackResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.spot.SpotCheckService/SpotCheckFeedback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spotCheckServiceClient) CameraCalibrationCommand(ctx context.Context, in *CameraCalibrationCommandRequest, opts ...grpc.CallOption) (*CameraCalibrationCommandResponse, error) {
	out := new(CameraCalibrationCommandResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.spot.SpotCheckService/CameraCalibrationCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spotCheckServiceClient) CameraCalibrationFeedback(ctx context.Context, in *CameraCalibrationFeedbackRequest, opts ...grpc.CallOption) (*CameraCalibrationFeedbackResponse, error) {
	out := new(CameraCalibrationFeedbackResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.spot.SpotCheckService/CameraCalibrationFeedback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpotCheckServiceServer is the server API for SpotCheckService service.
// All implementations must embed UnimplementedSpotCheckServiceServer
// for forward compatibility
type SpotCheckServiceServer interface {
	// Send a command to the SpotCheck service. The spotcheck service is responsible to both
	// recalibrating actuation sensors and checking camera health.
	SpotCheckCommand(context.Context, *SpotCheckCommandRequest) (*SpotCheckCommandResponse, error)
	// Check the status of the spot check procedure. After procedure completes, this reports back
	// results for specific joints and cameras.
	SpotCheckFeedback(context.Context, *SpotCheckFeedbackRequest) (*SpotCheckFeedbackResponse, error)
	// Send a camera calibration command to the robot. Used to start or abort a calibration routine.
	CameraCalibrationCommand(context.Context, *CameraCalibrationCommandRequest) (*CameraCalibrationCommandResponse, error)
	// Check the status of the camera calibration procedure.
	CameraCalibrationFeedback(context.Context, *CameraCalibrationFeedbackRequest) (*CameraCalibrationFeedbackResponse, error)
	mustEmbedUnimplementedSpotCheckServiceServer()
}

// UnimplementedSpotCheckServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSpotCheckServiceServer struct {
}

func (UnimplementedSpotCheckServiceServer) SpotCheckCommand(context.Context, *SpotCheckCommandRequest) (*SpotCheckCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SpotCheckCommand not implemented")
}
func (UnimplementedSpotCheckServiceServer) SpotCheckFeedback(context.Context, *SpotCheckFeedbackRequest) (*SpotCheckFeedbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SpotCheckFeedback not implemented")
}
func (UnimplementedSpotCheckServiceServer) CameraCalibrationCommand(context.Context, *CameraCalibrationCommandRequest) (*CameraCalibrationCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CameraCalibrationCommand not implemented")
}
func (UnimplementedSpotCheckServiceServer) CameraCalibrationFeedback(context.Context, *CameraCalibrationFeedbackRequest) (*CameraCalibrationFeedbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CameraCalibrationFeedback not implemented")
}
func (UnimplementedSpotCheckServiceServer) mustEmbedUnimplementedSpotCheckServiceServer() {}

// UnsafeSpotCheckServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpotCheckServiceServer will
// result in compilation errors.
type UnsafeSpotCheckServiceServer interface {
	mustEmbedUnimplementedSpotCheckServiceServer()
}

func RegisterSpotCheckServiceServer(s grpc.ServiceRegistrar, srv SpotCheckServiceServer) {
	s.RegisterService(&SpotCheckService_ServiceDesc, srv)
}

func _SpotCheckService_SpotCheckCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpotCheckCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpotCheckServiceServer).SpotCheckCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.spot.SpotCheckService/SpotCheckCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpotCheckServiceServer).SpotCheckCommand(ctx, req.(*SpotCheckCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpotCheckService_SpotCheckFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpotCheckFeedbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpotCheckServiceServer).SpotCheckFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.spot.SpotCheckService/SpotCheckFeedback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpotCheckServiceServer).SpotCheckFeedback(ctx, req.(*SpotCheckFeedbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpotCheckService_CameraCalibrationCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CameraCalibrationCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpotCheckServiceServer).CameraCalibrationCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.spot.SpotCheckService/CameraCalibrationCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpotCheckServiceServer).CameraCalibrationCommand(ctx, req.(*CameraCalibrationCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpotCheckService_CameraCalibrationFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CameraCalibrationFeedbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpotCheckServiceServer).CameraCalibrationFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.spot.SpotCheckService/CameraCalibrationFeedback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpotCheckServiceServer).CameraCalibrationFeedback(ctx, req.(*CameraCalibrationFeedbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SpotCheckService_ServiceDesc is the grpc.ServiceDesc for SpotCheckService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpotCheckService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bosdyn.api.spot.SpotCheckService",
	HandlerType: (*SpotCheckServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SpotCheckCommand",
			Handler:    _SpotCheckService_SpotCheckCommand_Handler,
		},
		{
			MethodName: "SpotCheckFeedback",
			Handler:    _SpotCheckService_SpotCheckFeedback_Handler,
		},
		{
			MethodName: "CameraCalibrationCommand",
			Handler:    _SpotCheckService_CameraCalibrationCommand_Handler,
		},
		{
			MethodName: "CameraCalibrationFeedback",
			Handler:    _SpotCheckService_CameraCalibrationFeedback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bosdyn/api/spot/spot_check_service.proto",
}
