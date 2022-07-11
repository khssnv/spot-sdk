// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bosdyn/api/network_compute_bridge_service.proto

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

// NetworkComputeBridgeClient is the client API for NetworkComputeBridge service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkComputeBridgeClient interface {
	NetworkCompute(ctx context.Context, in *NetworkComputeRequest, opts ...grpc.CallOption) (*NetworkComputeResponse, error)
	ListAvailableModels(ctx context.Context, in *ListAvailableModelsRequest, opts ...grpc.CallOption) (*ListAvailableModelsResponse, error)
}

type networkComputeBridgeClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkComputeBridgeClient(cc grpc.ClientConnInterface) NetworkComputeBridgeClient {
	return &networkComputeBridgeClient{cc}
}

func (c *networkComputeBridgeClient) NetworkCompute(ctx context.Context, in *NetworkComputeRequest, opts ...grpc.CallOption) (*NetworkComputeResponse, error) {
	out := new(NetworkComputeResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.NetworkComputeBridge/NetworkCompute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkComputeBridgeClient) ListAvailableModels(ctx context.Context, in *ListAvailableModelsRequest, opts ...grpc.CallOption) (*ListAvailableModelsResponse, error) {
	out := new(ListAvailableModelsResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.NetworkComputeBridge/ListAvailableModels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkComputeBridgeServer is the server API for NetworkComputeBridge service.
// All implementations must embed UnimplementedNetworkComputeBridgeServer
// for forward compatibility
type NetworkComputeBridgeServer interface {
	NetworkCompute(context.Context, *NetworkComputeRequest) (*NetworkComputeResponse, error)
	ListAvailableModels(context.Context, *ListAvailableModelsRequest) (*ListAvailableModelsResponse, error)
	mustEmbedUnimplementedNetworkComputeBridgeServer()
}

// UnimplementedNetworkComputeBridgeServer must be embedded to have forward compatible implementations.
type UnimplementedNetworkComputeBridgeServer struct {
}

func (UnimplementedNetworkComputeBridgeServer) NetworkCompute(context.Context, *NetworkComputeRequest) (*NetworkComputeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NetworkCompute not implemented")
}
func (UnimplementedNetworkComputeBridgeServer) ListAvailableModels(context.Context, *ListAvailableModelsRequest) (*ListAvailableModelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAvailableModels not implemented")
}
func (UnimplementedNetworkComputeBridgeServer) mustEmbedUnimplementedNetworkComputeBridgeServer() {}

// UnsafeNetworkComputeBridgeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkComputeBridgeServer will
// result in compilation errors.
type UnsafeNetworkComputeBridgeServer interface {
	mustEmbedUnimplementedNetworkComputeBridgeServer()
}

func RegisterNetworkComputeBridgeServer(s grpc.ServiceRegistrar, srv NetworkComputeBridgeServer) {
	s.RegisterService(&NetworkComputeBridge_ServiceDesc, srv)
}

func _NetworkComputeBridge_NetworkCompute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkComputeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkComputeBridgeServer).NetworkCompute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.NetworkComputeBridge/NetworkCompute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkComputeBridgeServer).NetworkCompute(ctx, req.(*NetworkComputeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkComputeBridge_ListAvailableModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAvailableModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkComputeBridgeServer).ListAvailableModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.NetworkComputeBridge/ListAvailableModels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkComputeBridgeServer).ListAvailableModels(ctx, req.(*ListAvailableModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NetworkComputeBridge_ServiceDesc is the grpc.ServiceDesc for NetworkComputeBridge service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkComputeBridge_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bosdyn.api.NetworkComputeBridge",
	HandlerType: (*NetworkComputeBridgeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NetworkCompute",
			Handler:    _NetworkComputeBridge_NetworkCompute_Handler,
		},
		{
			MethodName: "ListAvailableModels",
			Handler:    _NetworkComputeBridge_ListAvailableModels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bosdyn/api/network_compute_bridge_service.proto",
}

// NetworkComputeBridgeWorkerClient is the client API for NetworkComputeBridgeWorker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkComputeBridgeWorkerClient interface {
	NetworkCompute(ctx context.Context, in *NetworkComputeRequest, opts ...grpc.CallOption) (*NetworkComputeResponse, error)
	ListAvailableModels(ctx context.Context, in *ListAvailableModelsRequest, opts ...grpc.CallOption) (*ListAvailableModelsResponse, error)
}

type networkComputeBridgeWorkerClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkComputeBridgeWorkerClient(cc grpc.ClientConnInterface) NetworkComputeBridgeWorkerClient {
	return &networkComputeBridgeWorkerClient{cc}
}

func (c *networkComputeBridgeWorkerClient) NetworkCompute(ctx context.Context, in *NetworkComputeRequest, opts ...grpc.CallOption) (*NetworkComputeResponse, error) {
	out := new(NetworkComputeResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.NetworkComputeBridgeWorker/NetworkCompute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkComputeBridgeWorkerClient) ListAvailableModels(ctx context.Context, in *ListAvailableModelsRequest, opts ...grpc.CallOption) (*ListAvailableModelsResponse, error) {
	out := new(ListAvailableModelsResponse)
	err := c.cc.Invoke(ctx, "/bosdyn.api.NetworkComputeBridgeWorker/ListAvailableModels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkComputeBridgeWorkerServer is the server API for NetworkComputeBridgeWorker service.
// All implementations must embed UnimplementedNetworkComputeBridgeWorkerServer
// for forward compatibility
type NetworkComputeBridgeWorkerServer interface {
	NetworkCompute(context.Context, *NetworkComputeRequest) (*NetworkComputeResponse, error)
	ListAvailableModels(context.Context, *ListAvailableModelsRequest) (*ListAvailableModelsResponse, error)
	mustEmbedUnimplementedNetworkComputeBridgeWorkerServer()
}

// UnimplementedNetworkComputeBridgeWorkerServer must be embedded to have forward compatible implementations.
type UnimplementedNetworkComputeBridgeWorkerServer struct {
}

func (UnimplementedNetworkComputeBridgeWorkerServer) NetworkCompute(context.Context, *NetworkComputeRequest) (*NetworkComputeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NetworkCompute not implemented")
}
func (UnimplementedNetworkComputeBridgeWorkerServer) ListAvailableModels(context.Context, *ListAvailableModelsRequest) (*ListAvailableModelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAvailableModels not implemented")
}
func (UnimplementedNetworkComputeBridgeWorkerServer) mustEmbedUnimplementedNetworkComputeBridgeWorkerServer() {
}

// UnsafeNetworkComputeBridgeWorkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkComputeBridgeWorkerServer will
// result in compilation errors.
type UnsafeNetworkComputeBridgeWorkerServer interface {
	mustEmbedUnimplementedNetworkComputeBridgeWorkerServer()
}

func RegisterNetworkComputeBridgeWorkerServer(s grpc.ServiceRegistrar, srv NetworkComputeBridgeWorkerServer) {
	s.RegisterService(&NetworkComputeBridgeWorker_ServiceDesc, srv)
}

func _NetworkComputeBridgeWorker_NetworkCompute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkComputeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkComputeBridgeWorkerServer).NetworkCompute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.NetworkComputeBridgeWorker/NetworkCompute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkComputeBridgeWorkerServer).NetworkCompute(ctx, req.(*NetworkComputeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkComputeBridgeWorker_ListAvailableModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAvailableModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkComputeBridgeWorkerServer).ListAvailableModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosdyn.api.NetworkComputeBridgeWorker/ListAvailableModels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkComputeBridgeWorkerServer).ListAvailableModels(ctx, req.(*ListAvailableModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NetworkComputeBridgeWorker_ServiceDesc is the grpc.ServiceDesc for NetworkComputeBridgeWorker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkComputeBridgeWorker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bosdyn.api.NetworkComputeBridgeWorker",
	HandlerType: (*NetworkComputeBridgeWorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NetworkCompute",
			Handler:    _NetworkComputeBridgeWorker_NetworkCompute_Handler,
		},
		{
			MethodName: "ListAvailableModels",
			Handler:    _NetworkComputeBridgeWorker_ListAvailableModels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bosdyn/api/network_compute_bridge_service.proto",
}