// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bosdyn/api/graph_nav/map_processing_service.proto

package graph_nav

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

// MapProcessingServiceClient is the client API for MapProcessingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MapProcessingServiceClient interface {
	// Processes a GraphNav map by creating additional edges or waypoints. After processing,
	// a new subgraph is created containing additional waypoints or edges to add to the map.
	ProcessTopology(ctx context.Context, in *ProcessTopologyRequest, opts ...grpc.CallOption) (MapProcessingService_ProcessTopologyClient, error)
	// Processes a GraphNav map by modifying the anchoring of waypoints and world objects in the map
	// with respect to a seed frame. After processing, a new anchoring is streamed back.
	ProcessAnchoring(ctx context.Context, in *ProcessAnchoringRequest, opts ...grpc.CallOption) (MapProcessingService_ProcessAnchoringClient, error)
}

type mapProcessingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMapProcessingServiceClient(cc grpc.ClientConnInterface) MapProcessingServiceClient {
	return &mapProcessingServiceClient{cc}
}

func (c *mapProcessingServiceClient) ProcessTopology(ctx context.Context, in *ProcessTopologyRequest, opts ...grpc.CallOption) (MapProcessingService_ProcessTopologyClient, error) {
	stream, err := c.cc.NewStream(ctx, &MapProcessingService_ServiceDesc.Streams[0], "/bosdyn.api.graph_nav.MapProcessingService/ProcessTopology", opts...)
	if err != nil {
		return nil, err
	}
	x := &mapProcessingServiceProcessTopologyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MapProcessingService_ProcessTopologyClient interface {
	Recv() (*ProcessTopologyResponse, error)
	grpc.ClientStream
}

type mapProcessingServiceProcessTopologyClient struct {
	grpc.ClientStream
}

func (x *mapProcessingServiceProcessTopologyClient) Recv() (*ProcessTopologyResponse, error) {
	m := new(ProcessTopologyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mapProcessingServiceClient) ProcessAnchoring(ctx context.Context, in *ProcessAnchoringRequest, opts ...grpc.CallOption) (MapProcessingService_ProcessAnchoringClient, error) {
	stream, err := c.cc.NewStream(ctx, &MapProcessingService_ServiceDesc.Streams[1], "/bosdyn.api.graph_nav.MapProcessingService/ProcessAnchoring", opts...)
	if err != nil {
		return nil, err
	}
	x := &mapProcessingServiceProcessAnchoringClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MapProcessingService_ProcessAnchoringClient interface {
	Recv() (*ProcessAnchoringResponse, error)
	grpc.ClientStream
}

type mapProcessingServiceProcessAnchoringClient struct {
	grpc.ClientStream
}

func (x *mapProcessingServiceProcessAnchoringClient) Recv() (*ProcessAnchoringResponse, error) {
	m := new(ProcessAnchoringResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MapProcessingServiceServer is the server API for MapProcessingService service.
// All implementations must embed UnimplementedMapProcessingServiceServer
// for forward compatibility
type MapProcessingServiceServer interface {
	// Processes a GraphNav map by creating additional edges or waypoints. After processing,
	// a new subgraph is created containing additional waypoints or edges to add to the map.
	ProcessTopology(*ProcessTopologyRequest, MapProcessingService_ProcessTopologyServer) error
	// Processes a GraphNav map by modifying the anchoring of waypoints and world objects in the map
	// with respect to a seed frame. After processing, a new anchoring is streamed back.
	ProcessAnchoring(*ProcessAnchoringRequest, MapProcessingService_ProcessAnchoringServer) error
	mustEmbedUnimplementedMapProcessingServiceServer()
}

// UnimplementedMapProcessingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMapProcessingServiceServer struct {
}

func (UnimplementedMapProcessingServiceServer) ProcessTopology(*ProcessTopologyRequest, MapProcessingService_ProcessTopologyServer) error {
	return status.Errorf(codes.Unimplemented, "method ProcessTopology not implemented")
}
func (UnimplementedMapProcessingServiceServer) ProcessAnchoring(*ProcessAnchoringRequest, MapProcessingService_ProcessAnchoringServer) error {
	return status.Errorf(codes.Unimplemented, "method ProcessAnchoring not implemented")
}
func (UnimplementedMapProcessingServiceServer) mustEmbedUnimplementedMapProcessingServiceServer() {}

// UnsafeMapProcessingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MapProcessingServiceServer will
// result in compilation errors.
type UnsafeMapProcessingServiceServer interface {
	mustEmbedUnimplementedMapProcessingServiceServer()
}

func RegisterMapProcessingServiceServer(s grpc.ServiceRegistrar, srv MapProcessingServiceServer) {
	s.RegisterService(&MapProcessingService_ServiceDesc, srv)
}

func _MapProcessingService_ProcessTopology_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProcessTopologyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MapProcessingServiceServer).ProcessTopology(m, &mapProcessingServiceProcessTopologyServer{stream})
}

type MapProcessingService_ProcessTopologyServer interface {
	Send(*ProcessTopologyResponse) error
	grpc.ServerStream
}

type mapProcessingServiceProcessTopologyServer struct {
	grpc.ServerStream
}

func (x *mapProcessingServiceProcessTopologyServer) Send(m *ProcessTopologyResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MapProcessingService_ProcessAnchoring_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProcessAnchoringRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MapProcessingServiceServer).ProcessAnchoring(m, &mapProcessingServiceProcessAnchoringServer{stream})
}

type MapProcessingService_ProcessAnchoringServer interface {
	Send(*ProcessAnchoringResponse) error
	grpc.ServerStream
}

type mapProcessingServiceProcessAnchoringServer struct {
	grpc.ServerStream
}

func (x *mapProcessingServiceProcessAnchoringServer) Send(m *ProcessAnchoringResponse) error {
	return x.ServerStream.SendMsg(m)
}

// MapProcessingService_ServiceDesc is the grpc.ServiceDesc for MapProcessingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MapProcessingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bosdyn.api.graph_nav.MapProcessingService",
	HandlerType: (*MapProcessingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProcessTopology",
			Handler:       _MapProcessingService_ProcessTopology_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ProcessAnchoring",
			Handler:       _MapProcessingService_ProcessAnchoring_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "bosdyn/api/graph_nav/map_processing_service.proto",
}