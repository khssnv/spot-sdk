// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bosdyn/api/data_acquisition_store_service.proto

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

// DataAcquisitionStoreServiceClient is the client API for DataAcquisitionStoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataAcquisitionStoreServiceClient interface {
	// List all CaptureActionIds (which identify an entire AcquireData RPC's data captures)
	// that match the query parameters provided in the request.
	ListCaptureActions(ctx context.Context, in *ListCaptureActionsRequest, opts ...grpc.CallOption) (*ListCaptureActionsResponse, error)
	// List data identifiers (which identify specific pieces of data from
	// an action) for stored data that satisfy the query parameters in the request.
	ListStoredData(ctx context.Context, in *ListStoredDataRequest, opts ...grpc.CallOption) (*ListStoredDataResponse, error)
	// Store arbitrary data associated with a DataIdentifier.
	StoreData(ctx context.Context, in *StoreDataRequest, opts ...grpc.CallOption) (*StoreDataResponse, error)
	// Type-safe to images: list data identifiers (which identify specific images
	// from an action) for stored images that satisfy the
	// query parameters in the request.
	ListStoredImages(ctx context.Context, in *ListStoredImagesRequest, opts ...grpc.CallOption) (*ListStoredImagesResponse, error)
	// Type-safe to images: store image data associated with a DataIdentifier.
	StoreImage(ctx context.Context, in *StoreImageRequest, opts ...grpc.CallOption) (*StoreImageResponse, error)
	// Type-safe to JSON metadata: list data identifiers (which identify specific metadata from
	// an action) for stored metadata that satisfy the query parameters in the request.
	ListStoredMetadata(ctx context.Context, in *ListStoredMetadataRequest, opts ...grpc.CallOption) (*ListStoredMetadataResponse, error)
	// Type-safe to JSON metadata: store metadata associated with a DataIdentifier.
	StoreMetadata(ctx context.Context, in *StoreMetadataRequest, opts ...grpc.CallOption) (*StoreMetadataResponse, error)
}

type dataAcquisitionStoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataAcquisitionStoreServiceClient(cc grpc.ClientConnInterface) DataAcquisitionStoreServiceClient {
	return &dataAcquisitionStoreServiceClient{cc}
}

func (c *dataAcquisitionStoreServiceClient) ListCaptureActions(ctx context.Context, in *ListCaptureActionsRequest, opts ...grpc.CallOption) (*ListCaptureActionsResponse, error) {
	out := new(ListCaptureActionsResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.DataAcquisitionStoreService/ListCaptureActions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataAcquisitionStoreServiceClient) ListStoredData(ctx context.Context, in *ListStoredDataRequest, opts ...grpc.CallOption) (*ListStoredDataResponse, error) {
	out := new(ListStoredDataResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.DataAcquisitionStoreService/ListStoredData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataAcquisitionStoreServiceClient) StoreData(ctx context.Context, in *StoreDataRequest, opts ...grpc.CallOption) (*StoreDataResponse, error) {
	out := new(StoreDataResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.DataAcquisitionStoreService/StoreData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataAcquisitionStoreServiceClient) ListStoredImages(ctx context.Context, in *ListStoredImagesRequest, opts ...grpc.CallOption) (*ListStoredImagesResponse, error) {
	out := new(ListStoredImagesResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.DataAcquisitionStoreService/ListStoredImages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataAcquisitionStoreServiceClient) StoreImage(ctx context.Context, in *StoreImageRequest, opts ...grpc.CallOption) (*StoreImageResponse, error) {
	out := new(StoreImageResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.DataAcquisitionStoreService/StoreImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataAcquisitionStoreServiceClient) ListStoredMetadata(ctx context.Context, in *ListStoredMetadataRequest, opts ...grpc.CallOption) (*ListStoredMetadataResponse, error) {
	out := new(ListStoredMetadataResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.DataAcquisitionStoreService/ListStoredMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataAcquisitionStoreServiceClient) StoreMetadata(ctx context.Context, in *StoreMetadataRequest, opts ...grpc.CallOption) (*StoreMetadataResponse, error) {
	out := new(StoreMetadataResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.DataAcquisitionStoreService/StoreMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataAcquisitionStoreServiceServer is the server API for DataAcquisitionStoreService service.
// All implementations must embed UnimplementedDataAcquisitionStoreServiceServer
// for forward compatibility
type DataAcquisitionStoreServiceServer interface {
	// List all CaptureActionIds (which identify an entire AcquireData RPC's data captures)
	// that match the query parameters provided in the request.
	ListCaptureActions(context.Context, *ListCaptureActionsRequest) (*ListCaptureActionsResponse, error)
	// List data identifiers (which identify specific pieces of data from
	// an action) for stored data that satisfy the query parameters in the request.
	ListStoredData(context.Context, *ListStoredDataRequest) (*ListStoredDataResponse, error)
	// Store arbitrary data associated with a DataIdentifier.
	StoreData(context.Context, *StoreDataRequest) (*StoreDataResponse, error)
	// Type-safe to images: list data identifiers (which identify specific images
	// from an action) for stored images that satisfy the
	// query parameters in the request.
	ListStoredImages(context.Context, *ListStoredImagesRequest) (*ListStoredImagesResponse, error)
	// Type-safe to images: store image data associated with a DataIdentifier.
	StoreImage(context.Context, *StoreImageRequest) (*StoreImageResponse, error)
	// Type-safe to JSON metadata: list data identifiers (which identify specific metadata from
	// an action) for stored metadata that satisfy the query parameters in the request.
	ListStoredMetadata(context.Context, *ListStoredMetadataRequest) (*ListStoredMetadataResponse, error)
	// Type-safe to JSON metadata: store metadata associated with a DataIdentifier.
	StoreMetadata(context.Context, *StoreMetadataRequest) (*StoreMetadataResponse, error)
	mustEmbedUnimplementedDataAcquisitionStoreServiceServer()
}

// UnimplementedDataAcquisitionStoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataAcquisitionStoreServiceServer struct {
}

func (UnimplementedDataAcquisitionStoreServiceServer) ListCaptureActions(context.Context, *ListCaptureActionsRequest) (*ListCaptureActionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCaptureActions not implemented")
}
func (UnimplementedDataAcquisitionStoreServiceServer) ListStoredData(context.Context, *ListStoredDataRequest) (*ListStoredDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStoredData not implemented")
}
func (UnimplementedDataAcquisitionStoreServiceServer) StoreData(context.Context, *StoreDataRequest) (*StoreDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreData not implemented")
}
func (UnimplementedDataAcquisitionStoreServiceServer) ListStoredImages(context.Context, *ListStoredImagesRequest) (*ListStoredImagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStoredImages not implemented")
}
func (UnimplementedDataAcquisitionStoreServiceServer) StoreImage(context.Context, *StoreImageRequest) (*StoreImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreImage not implemented")
}
func (UnimplementedDataAcquisitionStoreServiceServer) ListStoredMetadata(context.Context, *ListStoredMetadataRequest) (*ListStoredMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStoredMetadata not implemented")
}
func (UnimplementedDataAcquisitionStoreServiceServer) StoreMetadata(context.Context, *StoreMetadataRequest) (*StoreMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreMetadata not implemented")
}
func (UnimplementedDataAcquisitionStoreServiceServer) mustEmbedUnimplementedDataAcquisitionStoreServiceServer() {
}

// UnsafeDataAcquisitionStoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataAcquisitionStoreServiceServer will
// result in compilation errors.
type UnsafeDataAcquisitionStoreServiceServer interface {
	mustEmbedUnimplementedDataAcquisitionStoreServiceServer()
}

func RegisterDataAcquisitionStoreServiceServer(s grpc.ServiceRegistrar, srv DataAcquisitionStoreServiceServer) {
	s.RegisterService(&DataAcquisitionStoreService_ServiceDesc, srv)
}

func _DataAcquisitionStoreService_ListCaptureActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCaptureActionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataAcquisitionStoreServiceServer).ListCaptureActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.DataAcquisitionStoreService/ListCaptureActions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataAcquisitionStoreServiceServer).ListCaptureActions(ctx, req.(*ListCaptureActionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataAcquisitionStoreService_ListStoredData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStoredDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataAcquisitionStoreServiceServer).ListStoredData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.DataAcquisitionStoreService/ListStoredData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataAcquisitionStoreServiceServer).ListStoredData(ctx, req.(*ListStoredDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataAcquisitionStoreService_StoreData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataAcquisitionStoreServiceServer).StoreData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.DataAcquisitionStoreService/StoreData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataAcquisitionStoreServiceServer).StoreData(ctx, req.(*StoreDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataAcquisitionStoreService_ListStoredImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStoredImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataAcquisitionStoreServiceServer).ListStoredImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.DataAcquisitionStoreService/ListStoredImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataAcquisitionStoreServiceServer).ListStoredImages(ctx, req.(*ListStoredImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataAcquisitionStoreService_StoreImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataAcquisitionStoreServiceServer).StoreImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.DataAcquisitionStoreService/StoreImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataAcquisitionStoreServiceServer).StoreImage(ctx, req.(*StoreImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataAcquisitionStoreService_ListStoredMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStoredMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataAcquisitionStoreServiceServer).ListStoredMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.DataAcquisitionStoreService/ListStoredMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataAcquisitionStoreServiceServer).ListStoredMetadata(ctx, req.(*ListStoredMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataAcquisitionStoreService_StoreMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataAcquisitionStoreServiceServer).StoreMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.DataAcquisitionStoreService/StoreMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataAcquisitionStoreServiceServer).StoreMetadata(ctx, req.(*StoreMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataAcquisitionStoreService_ServiceDesc is the grpc.ServiceDesc for DataAcquisitionStoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataAcquisitionStoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bosdyn.api.DataAcquisitionStoreService",
	HandlerType: (*DataAcquisitionStoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCaptureActions",
			Handler:    _DataAcquisitionStoreService_ListCaptureActions_Handler,
		},
		{
			MethodName: "ListStoredData",
			Handler:    _DataAcquisitionStoreService_ListStoredData_Handler,
		},
		{
			MethodName: "StoreData",
			Handler:    _DataAcquisitionStoreService_StoreData_Handler,
		},
		{
			MethodName: "ListStoredImages",
			Handler:    _DataAcquisitionStoreService_ListStoredImages_Handler,
		},
		{
			MethodName: "StoreImage",
			Handler:    _DataAcquisitionStoreService_StoreImage_Handler,
		},
		{
			MethodName: "ListStoredMetadata",
			Handler:    _DataAcquisitionStoreService_ListStoredMetadata_Handler,
		},
		{
			MethodName: "StoreMetadata",
			Handler:    _DataAcquisitionStoreService_StoreMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bosdyn/api/data_acquisition_store_service.proto",
}
