// Copyright (c) 2022 Boston Dynamics, Inc.  All rights reserved.
//
// Downloading, reproducing, distributing or otherwise using the SDK Software
// is subject to the terms and conditions of the Boston Dynamics Software
// Development Kit License (20191101-BDSDK-SL).

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: bosdyn/api/ray_cast.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RayIntersection_Type int32

const (
	// TYPE_UNKNOWN should not be used.
	RayIntersection_TYPE_UNKNOWN RayIntersection_Type = 0
	// Intersected against estimated ground plane.
	RayIntersection_TYPE_GROUND_PLANE RayIntersection_Type = 1
	// Intersected against the terrain map.
	RayIntersection_TYPE_TERRAIN_MAP RayIntersection_Type = 2
	// Intersected against the full 3D voxel map.
	RayIntersection_TYPE_VOXEL_MAP RayIntersection_Type = 3
	// Intersected against the hand depth data.
	RayIntersection_TYPE_HAND_DEPTH RayIntersection_Type = 4
)

// Enum value maps for RayIntersection_Type.
var (
	RayIntersection_Type_name = map[int32]string{
		0: "TYPE_UNKNOWN",
		1: "TYPE_GROUND_PLANE",
		2: "TYPE_TERRAIN_MAP",
		3: "TYPE_VOXEL_MAP",
		4: "TYPE_HAND_DEPTH",
	}
	RayIntersection_Type_value = map[string]int32{
		"TYPE_UNKNOWN":      0,
		"TYPE_GROUND_PLANE": 1,
		"TYPE_TERRAIN_MAP":  2,
		"TYPE_VOXEL_MAP":    3,
		"TYPE_HAND_DEPTH":   4,
	}
)

func (x RayIntersection_Type) Enum() *RayIntersection_Type {
	p := new(RayIntersection_Type)
	*p = x
	return p
}

func (x RayIntersection_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RayIntersection_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_bosdyn_api_ray_cast_proto_enumTypes[0].Descriptor()
}

func (RayIntersection_Type) Type() protoreflect.EnumType {
	return &file_bosdyn_api_ray_cast_proto_enumTypes[0]
}

func (x RayIntersection_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RayIntersection_Type.Descriptor instead.
func (RayIntersection_Type) EnumDescriptor() ([]byte, []int) {
	return file_bosdyn_api_ray_cast_proto_rawDescGZIP(), []int{1, 0}
}

type RaycastResponse_Status int32

const (
	// An unknown / unexpected error occurred.
	RaycastResponse_STATUS_UNKNOWN RaycastResponse_Status = 0
	// Request was accepted.
	RaycastResponse_STATUS_OK RaycastResponse_Status = 1
	// [Programming Error] Request was invalid / malformed in some way.
	RaycastResponse_STATUS_INVALID_REQUEST RaycastResponse_Status = 2
	// [Programming Error] Requested source not valid for current robot configuration.
	RaycastResponse_STATUS_INVALID_INTERSECTION_TYPE RaycastResponse_Status = 3
	// [Frame Error] The frame_name for a command was not a known frame.
	RaycastResponse_STATUS_UNKNOWN_FRAME RaycastResponse_Status = 4
)

// Enum value maps for RaycastResponse_Status.
var (
	RaycastResponse_Status_name = map[int32]string{
		0: "STATUS_UNKNOWN",
		1: "STATUS_OK",
		2: "STATUS_INVALID_REQUEST",
		3: "STATUS_INVALID_INTERSECTION_TYPE",
		4: "STATUS_UNKNOWN_FRAME",
	}
	RaycastResponse_Status_value = map[string]int32{
		"STATUS_UNKNOWN":                   0,
		"STATUS_OK":                        1,
		"STATUS_INVALID_REQUEST":           2,
		"STATUS_INVALID_INTERSECTION_TYPE": 3,
		"STATUS_UNKNOWN_FRAME":             4,
	}
)

func (x RaycastResponse_Status) Enum() *RaycastResponse_Status {
	p := new(RaycastResponse_Status)
	*p = x
	return p
}

func (x RaycastResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RaycastResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_bosdyn_api_ray_cast_proto_enumTypes[1].Descriptor()
}

func (RaycastResponse_Status) Type() protoreflect.EnumType {
	return &file_bosdyn_api_ray_cast_proto_enumTypes[1]
}

func (x RaycastResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RaycastResponse_Status.Descriptor instead.
func (RaycastResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_bosdyn_api_ray_cast_proto_rawDescGZIP(), []int{2, 0}
}

type RaycastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common request header.
	Header *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The ray's coordinate frame. When unset, this will default to vision frame.
	RayFrameName string `protobuf:"bytes,5,opt,name=ray_frame_name,json=rayFrameName,proto3" json:"ray_frame_name,omitempty"`
	// The ray, containing and origin and an direction.
	Ray *Ray `protobuf:"bytes,2,opt,name=ray,proto3" json:"ray,omitempty"`
	// Ignore intersections closer than this location on the ray.
	// Defaults to 0 if not provided.
	MinIntersectionDistance float32 `protobuf:"fixed32,4,opt,name=min_intersection_distance,json=minIntersectionDistance,proto3" json:"min_intersection_distance,omitempty"`
	// Type of the raycast you want to perform.  If multiple are set, the result will wait until
	// all raycasts are complete and return a single result proto.
	//
	// If this field is left empty, all available sources are used.
	IntersectionTypes []RayIntersection_Type `protobuf:"varint,7,rep,packed,name=intersection_types,json=intersectionTypes,proto3,enum=bosdyn.api.RayIntersection_Type" json:"intersection_types,omitempty"`
}

func (x *RaycastRequest) Reset() {
	*x = RaycastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_ray_cast_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaycastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaycastRequest) ProtoMessage() {}

func (x *RaycastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_ray_cast_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaycastRequest.ProtoReflect.Descriptor instead.
func (*RaycastRequest) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_ray_cast_proto_rawDescGZIP(), []int{0}
}

func (x *RaycastRequest) GetHeader() *RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *RaycastRequest) GetRayFrameName() string {
	if x != nil {
		return x.RayFrameName
	}
	return ""
}

func (x *RaycastRequest) GetRay() *Ray {
	if x != nil {
		return x.Ray
	}
	return nil
}

func (x *RaycastRequest) GetMinIntersectionDistance() float32 {
	if x != nil {
		return x.MinIntersectionDistance
	}
	return 0
}

func (x *RaycastRequest) GetIntersectionTypes() []RayIntersection_Type {
	if x != nil {
		return x.IntersectionTypes
	}
	return nil
}

type RayIntersection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of the raycast intersection that was performed.
	Type RayIntersection_Type `protobuf:"varint,1,opt,name=type,proto3,enum=bosdyn.api.RayIntersection_Type" json:"type,omitempty"`
	// Position of ray cast hit in the RaycastResponse hit_frame.
	HitPositionInHitFrame *Vec3 `protobuf:"bytes,2,opt,name=hit_position_in_hit_frame,json=hitPositionInHitFrame,proto3" json:"hit_position_in_hit_frame,omitempty"`
	// Distance of hit from ray origin.
	DistanceMeters float64 `protobuf:"fixed64,3,opt,name=distance_meters,json=distanceMeters,proto3" json:"distance_meters,omitempty"`
}

func (x *RayIntersection) Reset() {
	*x = RayIntersection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_ray_cast_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RayIntersection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RayIntersection) ProtoMessage() {}

func (x *RayIntersection) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_ray_cast_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RayIntersection.ProtoReflect.Descriptor instead.
func (*RayIntersection) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_ray_cast_proto_rawDescGZIP(), []int{1}
}

func (x *RayIntersection) GetType() RayIntersection_Type {
	if x != nil {
		return x.Type
	}
	return RayIntersection_TYPE_UNKNOWN
}

func (x *RayIntersection) GetHitPositionInHitFrame() *Vec3 {
	if x != nil {
		return x.HitPositionInHitFrame
	}
	return nil
}

func (x *RayIntersection) GetDistanceMeters() float64 {
	if x != nil {
		return x.DistanceMeters
	}
	return 0
}

type RaycastResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common response header.
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// Return status for a request.
	Status RaycastResponse_Status `protobuf:"varint,5,opt,name=status,proto3,enum=bosdyn.api.RaycastResponse_Status" json:"status,omitempty"`
	// Human-readable error description.  Not for programmatic analysis.
	Message string `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	// The frame raycast hits are returned in. Generally this should be the same frame the client
	// initially requested in.
	HitFrameName string `protobuf:"bytes,3,opt,name=hit_frame_name,json=hitFrameName,proto3" json:"hit_frame_name,omitempty"`
	// Ray cast hits, sorted with the closest hit first along the ray's extent.
	Hits []*RayIntersection `protobuf:"bytes,2,rep,name=hits,proto3" json:"hits,omitempty"`
	// A tree-based collection of transformations, which will include the
	// transformations to each of the returned world objects in addition to
	// transformations to the common frames ("vision", "body", "odom").  All
	// transforms within the snapshot are taken at the time when the request is received.
	//
	// Note that each object's frame names are defined within the properties
	// submessage e.g. "frame_name".
	TransformsSnapshot *FrameTreeSnapshot `protobuf:"bytes,4,opt,name=transforms_snapshot,json=transformsSnapshot,proto3" json:"transforms_snapshot,omitempty"`
}

func (x *RaycastResponse) Reset() {
	*x = RaycastResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_ray_cast_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaycastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaycastResponse) ProtoMessage() {}

func (x *RaycastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_ray_cast_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaycastResponse.ProtoReflect.Descriptor instead.
func (*RaycastResponse) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_ray_cast_proto_rawDescGZIP(), []int{2}
}

func (x *RaycastResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *RaycastResponse) GetStatus() RaycastResponse_Status {
	if x != nil {
		return x.Status
	}
	return RaycastResponse_STATUS_UNKNOWN
}

func (x *RaycastResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RaycastResponse) GetHitFrameName() string {
	if x != nil {
		return x.HitFrameName
	}
	return ""
}

func (x *RaycastResponse) GetHits() []*RayIntersection {
	if x != nil {
		return x.Hits
	}
	return nil
}

func (x *RaycastResponse) GetTransformsSnapshot() *FrameTreeSnapshot {
	if x != nil {
		return x.TransformsSnapshot
	}
	return nil
}

var File_bosdyn_api_ray_cast_proto protoreflect.FileDescriptor

var file_bosdyn_api_ray_cast_proto_rawDesc = []byte{
	0x0a, 0x19, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x61, 0x79,
	0x5f, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62, 0x6f, 0x73,
	0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x17, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6f,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x02, 0x0a, 0x0e,
	0x52, 0x61, 0x79, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31,
	0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x61, 0x79, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x61, 0x79, 0x46, 0x72,
	0x61, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x03, 0x72, 0x61, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x61, 0x79, 0x52, 0x03, 0x72, 0x61, 0x79, 0x12, 0x3a, 0x0a, 0x19, 0x6d, 0x69,
	0x6e, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x17, 0x6d,
	0x69, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x20, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x61, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0xac, 0x02, 0x0a, 0x0f, 0x52, 0x61, 0x79, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x62, 0x6f, 0x73, 0x64,
	0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x61, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x73,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x4a, 0x0a, 0x19, 0x68, 0x69, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x6e, 0x5f, 0x68, 0x69, 0x74, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x56, 0x65, 0x63, 0x33, 0x52, 0x15, 0x68, 0x69, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x48, 0x69, 0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a,
	0x0f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x4d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0x6e, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10,
	0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x15, 0x0a, 0x11, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x5f,
	0x50, 0x4c, 0x41, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x54, 0x45, 0x52, 0x52, 0x41, 0x49, 0x4e, 0x5f, 0x4d, 0x41, 0x50, 0x10, 0x02, 0x12, 0x12, 0x0a,
	0x0e, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x4f, 0x58, 0x45, 0x4c, 0x5f, 0x4d, 0x41, 0x50, 0x10,
	0x03, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x41, 0x4e, 0x44, 0x5f, 0x44,
	0x45, 0x50, 0x54, 0x48, 0x10, 0x04, 0x22, 0xcc, 0x03, 0x0a, 0x0f, 0x52, 0x61, 0x79, 0x63, 0x61,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6f, 0x73,
	0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3a,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22,
	0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x61, 0x79, 0x63,
	0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x68, 0x69, 0x74, 0x5f, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x69,
	0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x68, 0x69,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x61, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x68, 0x69, 0x74, 0x73, 0x12, 0x4e, 0x0a, 0x13, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x54, 0x72, 0x65, 0x65, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x73, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x22, 0x87, 0x01, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x53, 0x45, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x46, 0x52,
	0x41, 0x4d, 0x45, 0x10, 0x04, 0x42, 0x41, 0x42, 0x0c, 0x52, 0x61, 0x79, 0x43, 0x61, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6b, 0x68, 0x73, 0x73, 0x6e, 0x76, 0x2f, 0x73, 0x70, 0x6f, 0x74, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x70, 0x72, 0x65, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x6f,
	0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bosdyn_api_ray_cast_proto_rawDescOnce sync.Once
	file_bosdyn_api_ray_cast_proto_rawDescData = file_bosdyn_api_ray_cast_proto_rawDesc
)

func file_bosdyn_api_ray_cast_proto_rawDescGZIP() []byte {
	file_bosdyn_api_ray_cast_proto_rawDescOnce.Do(func() {
		file_bosdyn_api_ray_cast_proto_rawDescData = protoimpl.X.CompressGZIP(file_bosdyn_api_ray_cast_proto_rawDescData)
	})
	return file_bosdyn_api_ray_cast_proto_rawDescData
}

var file_bosdyn_api_ray_cast_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_bosdyn_api_ray_cast_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bosdyn_api_ray_cast_proto_goTypes = []interface{}{
	(RayIntersection_Type)(0),   // 0: bosdyn.api.RayIntersection.Type
	(RaycastResponse_Status)(0), // 1: bosdyn.api.RaycastResponse.Status
	(*RaycastRequest)(nil),      // 2: bosdyn.api.RaycastRequest
	(*RayIntersection)(nil),     // 3: bosdyn.api.RayIntersection
	(*RaycastResponse)(nil),     // 4: bosdyn.api.RaycastResponse
	(*RequestHeader)(nil),       // 5: bosdyn.api.RequestHeader
	(*Ray)(nil),                 // 6: bosdyn.api.Ray
	(*Vec3)(nil),                // 7: bosdyn.api.Vec3
	(*ResponseHeader)(nil),      // 8: bosdyn.api.ResponseHeader
	(*FrameTreeSnapshot)(nil),   // 9: bosdyn.api.FrameTreeSnapshot
}
var file_bosdyn_api_ray_cast_proto_depIdxs = []int32{
	5, // 0: bosdyn.api.RaycastRequest.header:type_name -> bosdyn.api.RequestHeader
	6, // 1: bosdyn.api.RaycastRequest.ray:type_name -> bosdyn.api.Ray
	0, // 2: bosdyn.api.RaycastRequest.intersection_types:type_name -> bosdyn.api.RayIntersection.Type
	0, // 3: bosdyn.api.RayIntersection.type:type_name -> bosdyn.api.RayIntersection.Type
	7, // 4: bosdyn.api.RayIntersection.hit_position_in_hit_frame:type_name -> bosdyn.api.Vec3
	8, // 5: bosdyn.api.RaycastResponse.header:type_name -> bosdyn.api.ResponseHeader
	1, // 6: bosdyn.api.RaycastResponse.status:type_name -> bosdyn.api.RaycastResponse.Status
	3, // 7: bosdyn.api.RaycastResponse.hits:type_name -> bosdyn.api.RayIntersection
	9, // 8: bosdyn.api.RaycastResponse.transforms_snapshot:type_name -> bosdyn.api.FrameTreeSnapshot
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_bosdyn_api_ray_cast_proto_init() }
func file_bosdyn_api_ray_cast_proto_init() {
	if File_bosdyn_api_ray_cast_proto != nil {
		return
	}
	file_bosdyn_api_header_proto_init()
	file_bosdyn_api_geometry_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bosdyn_api_ray_cast_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaycastRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bosdyn_api_ray_cast_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RayIntersection); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bosdyn_api_ray_cast_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaycastResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bosdyn_api_ray_cast_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bosdyn_api_ray_cast_proto_goTypes,
		DependencyIndexes: file_bosdyn_api_ray_cast_proto_depIdxs,
		EnumInfos:         file_bosdyn_api_ray_cast_proto_enumTypes,
		MessageInfos:      file_bosdyn_api_ray_cast_proto_msgTypes,
	}.Build()
	File_bosdyn_api_ray_cast_proto = out.File
	file_bosdyn_api_ray_cast_proto_rawDesc = nil
	file_bosdyn_api_ray_cast_proto_goTypes = nil
	file_bosdyn_api_ray_cast_proto_depIdxs = nil
}