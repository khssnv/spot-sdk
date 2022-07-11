// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bosdyn/api/data_buffer_service.proto

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

// DataBufferServiceClient is the client API for DataBufferService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataBufferServiceClient interface {
	// Add text messages to the log.
	RecordTextMessages(ctx context.Context, in *RecordTextMessagesRequest, opts ...grpc.CallOption) (*RecordTextMessagesResponse, error)
	// Add a set of operator messages to the log.
	RecordOperatorComments(ctx context.Context, in *RecordOperatorCommentsRequest, opts ...grpc.CallOption) (*RecordOperatorCommentsResponse, error)
	// Add message-style data to the log.
	RecordDataBlobs(ctx context.Context, in *RecordDataBlobsRequest, opts ...grpc.CallOption) (*RecordDataBlobsResponse, error)
	// Add event data to the log.
	RecordEvents(ctx context.Context, in *RecordEventsRequest, opts ...grpc.CallOption) (*RecordEventsResponse, error)
	// Register a log tick schema, allowing client to later log tick data.
	RegisterSignalSchema(ctx context.Context, in *RegisterSignalSchemaRequest, opts ...grpc.CallOption) (*RegisterSignalSchemaResponse, error)
	// Add signal data for registered signal schema to the log.
	RecordSignalTicks(ctx context.Context, in *RecordSignalTicksRequest, opts ...grpc.CallOption) (*RecordSignalTicksResponse, error)
}

type dataBufferServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataBufferServiceClient(cc grpc.ClientConnInterface) DataBufferServiceClient {
	return &dataBufferServiceClient{cc}
}

func (c *dataBufferServiceClient) RecordTextMessages(ctx context.Context, in *RecordTextMessagesRequest, opts ...grpc.CallOption) (*RecordTextMessagesResponse, error) {
	out := new(RecordTextMessagesResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.DataBufferService/RecordTextMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataBufferServiceClient) RecordOperatorComments(ctx context.Context, in *RecordOperatorCommentsRequest, opts ...grpc.CallOption) (*RecordOperatorCommentsResponse, error) {
	out := new(RecordOperatorCommentsResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.DataBufferService/RecordOperatorComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataBufferServiceClient) RecordDataBlobs(ctx context.Context, in *RecordDataBlobsRequest, opts ...grpc.CallOption) (*RecordDataBlobsResponse, error) {
	out := new(RecordDataBlobsResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.DataBufferService/RecordDataBlobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataBufferServiceClient) RecordEvents(ctx context.Context, in *RecordEventsRequest, opts ...grpc.CallOption) (*RecordEventsResponse, error) {
	out := new(RecordEventsResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.DataBufferService/RecordEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataBufferServiceClient) RegisterSignalSchema(ctx context.Context, in *RegisterSignalSchemaRequest, opts ...grpc.CallOption) (*RegisterSignalSchemaResponse, error) {
	out := new(RegisterSignalSchemaResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.DataBufferService/RegisterSignalSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataBufferServiceClient) RecordSignalTicks(ctx context.Context, in *RecordSignalTicksRequest, opts ...grpc.CallOption) (*RecordSignalTicksResponse, error) {
	out := new(RecordSignalTicksResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.DataBufferService/RecordSignalTicks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataBufferServiceServer is the server API for DataBufferService service.
// All implementations must embed UnimplementedDataBufferServiceServer
// for forward compatibility
type DataBufferServiceServer interface {
	// Add text messages to the log.
	RecordTextMessages(context.Context, *RecordTextMessagesRequest) (*RecordTextMessagesResponse, error)
	// Add a set of operator messages to the log.
	RecordOperatorComments(context.Context, *RecordOperatorCommentsRequest) (*RecordOperatorCommentsResponse, error)
	// Add message-style data to the log.
	RecordDataBlobs(context.Context, *RecordDataBlobsRequest) (*RecordDataBlobsResponse, error)
	// Add event data to the log.
	RecordEvents(context.Context, *RecordEventsRequest) (*RecordEventsResponse, error)
	// Register a log tick schema, allowing client to later log tick data.
	RegisterSignalSchema(context.Context, *RegisterSignalSchemaRequest) (*RegisterSignalSchemaResponse, error)
	// Add signal data for registered signal schema to the log.
	RecordSignalTicks(context.Context, *RecordSignalTicksRequest) (*RecordSignalTicksResponse, error)
	mustEmbedUnimplementedDataBufferServiceServer()
}

// UnimplementedDataBufferServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataBufferServiceServer struct {
}

func (UnimplementedDataBufferServiceServer) RecordTextMessages(context.Context, *RecordTextMessagesRequest) (*RecordTextMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordTextMessages not implemented")
}
func (UnimplementedDataBufferServiceServer) RecordOperatorComments(context.Context, *RecordOperatorCommentsRequest) (*RecordOperatorCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordOperatorComments not implemented")
}
func (UnimplementedDataBufferServiceServer) RecordDataBlobs(context.Context, *RecordDataBlobsRequest) (*RecordDataBlobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordDataBlobs not implemented")
}
func (UnimplementedDataBufferServiceServer) RecordEvents(context.Context, *RecordEventsRequest) (*RecordEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordEvents not implemented")
}
func (UnimplementedDataBufferServiceServer) RegisterSignalSchema(context.Context, *RegisterSignalSchemaRequest) (*RegisterSignalSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterSignalSchema not implemented")
}
func (UnimplementedDataBufferServiceServer) RecordSignalTicks(context.Context, *RecordSignalTicksRequest) (*RecordSignalTicksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordSignalTicks not implemented")
}
func (UnimplementedDataBufferServiceServer) mustEmbedUnimplementedDataBufferServiceServer() {}

// UnsafeDataBufferServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataBufferServiceServer will
// result in compilation errors.
type UnsafeDataBufferServiceServer interface {
	mustEmbedUnimplementedDataBufferServiceServer()
}

func RegisterDataBufferServiceServer(s grpc.ServiceRegistrar, srv DataBufferServiceServer) {
	s.RegisterService(&DataBufferService_ServiceDesc, srv)
}

func _DataBufferService_RecordTextMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordTextMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataBufferServiceServer).RecordTextMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.DataBufferService/RecordTextMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataBufferServiceServer).RecordTextMessages(ctx, req.(*RecordTextMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataBufferService_RecordOperatorComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordOperatorCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataBufferServiceServer).RecordOperatorComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.DataBufferService/RecordOperatorComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataBufferServiceServer).RecordOperatorComments(ctx, req.(*RecordOperatorCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataBufferService_RecordDataBlobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordDataBlobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataBufferServiceServer).RecordDataBlobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.DataBufferService/RecordDataBlobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataBufferServiceServer).RecordDataBlobs(ctx, req.(*RecordDataBlobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataBufferService_RecordEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataBufferServiceServer).RecordEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.DataBufferService/RecordEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataBufferServiceServer).RecordEvents(ctx, req.(*RecordEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataBufferService_RegisterSignalSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterSignalSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataBufferServiceServer).RegisterSignalSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.DataBufferService/RegisterSignalSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataBufferServiceServer).RegisterSignalSchema(ctx, req.(*RegisterSignalSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataBufferService_RecordSignalTicks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordSignalTicksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataBufferServiceServer).RecordSignalTicks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.DataBufferService/RecordSignalTicks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataBufferServiceServer).RecordSignalTicks(ctx, req.(*RecordSignalTicksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataBufferService_ServiceDesc is the grpc.ServiceDesc for DataBufferService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataBufferService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bosdyn.api.DataBufferService",
	HandlerType: (*DataBufferServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecordTextMessages",
			Handler:    _DataBufferService_RecordTextMessages_Handler,
		},
		{
			MethodName: "RecordOperatorComments",
			Handler:    _DataBufferService_RecordOperatorComments_Handler,
		},
		{
			MethodName: "RecordDataBlobs",
			Handler:    _DataBufferService_RecordDataBlobs_Handler,
		},
		{
			MethodName: "RecordEvents",
			Handler:    _DataBufferService_RecordEvents_Handler,
		},
		{
			MethodName: "RegisterSignalSchema",
			Handler:    _DataBufferService_RegisterSignalSchema_Handler,
		},
		{
			MethodName: "RecordSignalTicks",
			Handler:    _DataBufferService_RecordSignalTicks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bosdyn/api/data_buffer_service.proto",
}