// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bosdyn/api/data_acquisition_plugin_service.proto

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

// DataAcquisitionPluginServiceClient is the client API for DataAcquisitionPluginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataAcquisitionPluginServiceClient interface {
	// Trigger a data acquisition to save metadata and non-image data to the data buffer.
	// Sent by the main DAQ as a result of a data acquisition request from the tablet or a client.
	AcquirePluginData(ctx context.Context, in *AcquirePluginDataRequest, opts ...grpc.CallOption) (*AcquirePluginDataResponse, error)
	// Query the status of a data acquisition.
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error)
	// Get information from a DAQ service; lists acquisition capabilities.
	GetServiceInfo(ctx context.Context, in *GetServiceInfoRequest, opts ...grpc.CallOption) (*GetServiceInfoResponse, error)
	// Cancel an in-progress data acquisition.
	CancelAcquisition(ctx context.Context, in *CancelAcquisitionRequest, opts ...grpc.CallOption) (*CancelAcquisitionResponse, error)
}

type dataAcquisitionPluginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataAcquisitionPluginServiceClient(cc grpc.ClientConnInterface) DataAcquisitionPluginServiceClient {
	return &dataAcquisitionPluginServiceClient{cc}
}

func (c *dataAcquisitionPluginServiceClient) AcquirePluginData(ctx context.Context, in *AcquirePluginDataRequest, opts ...grpc.CallOption) (*AcquirePluginDataResponse, error) {
	out := new(AcquirePluginDataResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.DataAcquisitionPluginService/AcquirePluginData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataAcquisitionPluginServiceClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.DataAcquisitionPluginService/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataAcquisitionPluginServiceClient) GetServiceInfo(ctx context.Context, in *GetServiceInfoRequest, opts ...grpc.CallOption) (*GetServiceInfoResponse, error) {
	out := new(GetServiceInfoResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.DataAcquisitionPluginService/GetServiceInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataAcquisitionPluginServiceClient) CancelAcquisition(ctx context.Context, in *CancelAcquisitionRequest, opts ...grpc.CallOption) (*CancelAcquisitionResponse, error) {
	out := new(CancelAcquisitionResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.DataAcquisitionPluginService/CancelAcquisition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataAcquisitionPluginServiceServer is the server API for DataAcquisitionPluginService service.
// All implementations must embed UnimplementedDataAcquisitionPluginServiceServer
// for forward compatibility
type DataAcquisitionPluginServiceServer interface {
	// Trigger a data acquisition to save metadata and non-image data to the data buffer.
	// Sent by the main DAQ as a result of a data acquisition request from the tablet or a client.
	AcquirePluginData(context.Context, *AcquirePluginDataRequest) (*AcquirePluginDataResponse, error)
	// Query the status of a data acquisition.
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error)
	// Get information from a DAQ service; lists acquisition capabilities.
	GetServiceInfo(context.Context, *GetServiceInfoRequest) (*GetServiceInfoResponse, error)
	// Cancel an in-progress data acquisition.
	CancelAcquisition(context.Context, *CancelAcquisitionRequest) (*CancelAcquisitionResponse, error)
	mustEmbedUnimplementedDataAcquisitionPluginServiceServer()
}

// UnimplementedDataAcquisitionPluginServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataAcquisitionPluginServiceServer struct {
}

func (UnimplementedDataAcquisitionPluginServiceServer) AcquirePluginData(context.Context, *AcquirePluginDataRequest) (*AcquirePluginDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcquirePluginData not implemented")
}
func (UnimplementedDataAcquisitionPluginServiceServer) GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedDataAcquisitionPluginServiceServer) GetServiceInfo(context.Context, *GetServiceInfoRequest) (*GetServiceInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceInfo not implemented")
}
func (UnimplementedDataAcquisitionPluginServiceServer) CancelAcquisition(context.Context, *CancelAcquisitionRequest) (*CancelAcquisitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelAcquisition not implemented")
}
func (UnimplementedDataAcquisitionPluginServiceServer) mustEmbedUnimplementedDataAcquisitionPluginServiceServer() {
}

// UnsafeDataAcquisitionPluginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataAcquisitionPluginServiceServer will
// result in compilation errors.
type UnsafeDataAcquisitionPluginServiceServer interface {
	mustEmbedUnimplementedDataAcquisitionPluginServiceServer()
}

func RegisterDataAcquisitionPluginServiceServer(s grpc.ServiceRegistrar, srv DataAcquisitionPluginServiceServer) {
	s.RegisterService(&DataAcquisitionPluginService_ServiceDesc, srv)
}

func _DataAcquisitionPluginService_AcquirePluginData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcquirePluginDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataAcquisitionPluginServiceServer).AcquirePluginData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.DataAcquisitionPluginService/AcquirePluginData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataAcquisitionPluginServiceServer).AcquirePluginData(ctx, req.(*AcquirePluginDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataAcquisitionPluginService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataAcquisitionPluginServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.DataAcquisitionPluginService/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataAcquisitionPluginServiceServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataAcquisitionPluginService_GetServiceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataAcquisitionPluginServiceServer).GetServiceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.DataAcquisitionPluginService/GetServiceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataAcquisitionPluginServiceServer).GetServiceInfo(ctx, req.(*GetServiceInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataAcquisitionPluginService_CancelAcquisition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelAcquisitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataAcquisitionPluginServiceServer).CancelAcquisition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.DataAcquisitionPluginService/CancelAcquisition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataAcquisitionPluginServiceServer).CancelAcquisition(ctx, req.(*CancelAcquisitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataAcquisitionPluginService_ServiceDesc is the grpc.ServiceDesc for DataAcquisitionPluginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataAcquisitionPluginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bosdyn.api.DataAcquisitionPluginService",
	HandlerType: (*DataAcquisitionPluginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AcquirePluginData",
			Handler:    _DataAcquisitionPluginService_AcquirePluginData_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _DataAcquisitionPluginService_GetStatus_Handler,
		},
		{
			MethodName: "GetServiceInfo",
			Handler:    _DataAcquisitionPluginService_GetServiceInfo_Handler,
		},
		{
			MethodName: "CancelAcquisition",
			Handler:    _DataAcquisitionPluginService_CancelAcquisition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bosdyn/api/data_acquisition_plugin_service.proto",
}
