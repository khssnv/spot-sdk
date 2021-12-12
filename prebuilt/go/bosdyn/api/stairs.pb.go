// Copyright (c) 2021 Boston Dynamics, Inc.  All rights reserved.
//
// Downloading, reproducing, distributing or otherwise using the SDK Software
// is subject to the terms and conditions of the Boston Dynamics Software
// Development Kit License (20191101-BDSDK-SL).

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: bosdyn/api/stairs.proto

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

type StairTransform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The staircase origin is the bottom-center of the first rise.
	FrameTformStairs *SE3Pose `protobuf:"bytes,1,opt,name=frame_tform_stairs,json=frameTformStairs,proto3" json:"frame_tform_stairs,omitempty"`
	FrameName        string   `protobuf:"bytes,2,opt,name=frame_name,json=frameName,proto3" json:"frame_name,omitempty"`
}

func (x *StairTransform) Reset() {
	*x = StairTransform{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_stairs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StairTransform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StairTransform) ProtoMessage() {}

func (x *StairTransform) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_stairs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StairTransform.ProtoReflect.Descriptor instead.
func (*StairTransform) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_stairs_proto_rawDescGZIP(), []int{0}
}

func (x *StairTransform) GetFrameTformStairs() *SE3Pose {
	if x != nil {
		return x.FrameTformStairs
	}
	return nil
}

func (x *StairTransform) GetFrameName() string {
	if x != nil {
		return x.FrameName
	}
	return ""
}

type StraightStaircase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The staircase origin is the bottom-center of the first rise.
	//
	// Types that are assignable to Location:
	//	*StraightStaircase_FromKoTformStairs
	//	*StraightStaircase_Tform
	Location isStraightStaircase_Location `protobuf_oneof:"location"`
	// Each stair should be rise followed by run. The last stair will have zero run.
	Stairs []*StraightStaircase_Stair `protobuf:"bytes,2,rep,name=stairs,proto3" json:"stairs,omitempty"`
	// The lowermost landing of the stairs. The robot will try to
	// align itself to the stairs while on this landing.
	BottomLanding *StraightStaircase_Landing `protobuf:"bytes,3,opt,name=bottom_landing,json=bottomLanding,proto3" json:"bottom_landing,omitempty"`
	// The uppermost landing of the stairs.
	TopLanding *StraightStaircase_Landing `protobuf:"bytes,4,opt,name=top_landing,json=topLanding,proto3" json:"top_landing,omitempty"`
}

func (x *StraightStaircase) Reset() {
	*x = StraightStaircase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_stairs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StraightStaircase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StraightStaircase) ProtoMessage() {}

func (x *StraightStaircase) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_stairs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StraightStaircase.ProtoReflect.Descriptor instead.
func (*StraightStaircase) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_stairs_proto_rawDescGZIP(), []int{1}
}

func (m *StraightStaircase) GetLocation() isStraightStaircase_Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (x *StraightStaircase) GetFromKoTformStairs() *SE3Pose {
	if x, ok := x.GetLocation().(*StraightStaircase_FromKoTformStairs); ok {
		return x.FromKoTformStairs
	}
	return nil
}

func (x *StraightStaircase) GetTform() *StairTransform {
	if x, ok := x.GetLocation().(*StraightStaircase_Tform); ok {
		return x.Tform
	}
	return nil
}

func (x *StraightStaircase) GetStairs() []*StraightStaircase_Stair {
	if x != nil {
		return x.Stairs
	}
	return nil
}

func (x *StraightStaircase) GetBottomLanding() *StraightStaircase_Landing {
	if x != nil {
		return x.BottomLanding
	}
	return nil
}

func (x *StraightStaircase) GetTopLanding() *StraightStaircase_Landing {
	if x != nil {
		return x.TopLanding
	}
	return nil
}

type isStraightStaircase_Location interface {
	isStraightStaircase_Location()
}

type StraightStaircase_FromKoTformStairs struct {
	// It is expressed in ko frame of the from_waypoint.
	// This field is only used in GraphNav.
	FromKoTformStairs *SE3Pose `protobuf:"bytes,1,opt,name=from_ko_tform_stairs,json=fromKoTformStairs,proto3,oneof"`
}

type StraightStaircase_Tform struct {
	// Outside GraphNav, this field specifies the stair origin.
	Tform *StairTransform `protobuf:"bytes,5,opt,name=tform,proto3,oneof"`
}

func (*StraightStaircase_FromKoTformStairs) isStraightStaircase_Location() {}

func (*StraightStaircase_Tform) isStraightStaircase_Location() {}

// A single stair from a staircase.
type StraightStaircase_Stair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Height of each stair.
	Rise float32 `protobuf:"fixed32,1,opt,name=rise,proto3" json:"rise,omitempty"`
	// Depth of each stair.
	Run float32 `protobuf:"fixed32,2,opt,name=run,proto3" json:"run,omitempty"`
}

func (x *StraightStaircase_Stair) Reset() {
	*x = StraightStaircase_Stair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_stairs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StraightStaircase_Stair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StraightStaircase_Stair) ProtoMessage() {}

func (x *StraightStaircase_Stair) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_stairs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StraightStaircase_Stair.ProtoReflect.Descriptor instead.
func (*StraightStaircase_Stair) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_stairs_proto_rawDescGZIP(), []int{1, 0}
}

func (x *StraightStaircase_Stair) GetRise() float32 {
	if x != nil {
		return x.Rise
	}
	return 0
}

func (x *StraightStaircase_Stair) GetRun() float32 {
	if x != nil {
		return x.Run
	}
	return 0
}

// Straight staircases have two landings, one at the top and one at the bottom.
// Landings are areas of free space before and after the stairs, and are represented
// as oriented bounding boxes.
type StraightStaircase_Landing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Pose of the landing's center relative to the stairs frame.
	StairsTformLandingCenter *SE3Pose `protobuf:"bytes,1,opt,name=stairs_tform_landing_center,json=stairsTformLandingCenter,proto3" json:"stairs_tform_landing_center,omitempty"`
	// The half-size of the box representing the landing in the x axis.
	LandingExtentX float64 `protobuf:"fixed64,2,opt,name=landing_extent_x,json=landingExtentX,proto3" json:"landing_extent_x,omitempty"`
	// The half-size of the box representing the landing in the y axis.
	LandingExtentY float64 `protobuf:"fixed64,3,opt,name=landing_extent_y,json=landingExtentY,proto3" json:"landing_extent_y,omitempty"`
}

func (x *StraightStaircase_Landing) Reset() {
	*x = StraightStaircase_Landing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_stairs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StraightStaircase_Landing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StraightStaircase_Landing) ProtoMessage() {}

func (x *StraightStaircase_Landing) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_stairs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StraightStaircase_Landing.ProtoReflect.Descriptor instead.
func (*StraightStaircase_Landing) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_stairs_proto_rawDescGZIP(), []int{1, 1}
}

func (x *StraightStaircase_Landing) GetStairsTformLandingCenter() *SE3Pose {
	if x != nil {
		return x.StairsTformLandingCenter
	}
	return nil
}

func (x *StraightStaircase_Landing) GetLandingExtentX() float64 {
	if x != nil {
		return x.LandingExtentX
	}
	return 0
}

func (x *StraightStaircase_Landing) GetLandingExtentY() float64 {
	if x != nil {
		return x.LandingExtentY
	}
	return 0
}

var File_bosdyn_api_stairs_proto protoreflect.FileDescriptor

var file_bosdyn_api_stairs_proto_rawDesc = []byte{
	0x0a, 0x17, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x61,
	0x69, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62, 0x6f, 0x73, 0x64, 0x79,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x19, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x72, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x69, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x12, 0x41, 0x0a, 0x12, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x5f, 0x73, 0x74, 0x61, 0x69, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x45, 0x33, 0x50,
	0x6f, 0x73, 0x65, 0x52, 0x10, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x54, 0x66, 0x6f, 0x72, 0x6d, 0x53,
	0x74, 0x61, 0x69, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0xd1, 0x04, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x61, 0x69, 0x67, 0x68,
	0x74, 0x53, 0x74, 0x61, 0x69, 0x72, 0x63, 0x61, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x14, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x6b, 0x6f, 0x5f, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x73, 0x74, 0x61, 0x69,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x45, 0x33, 0x50, 0x6f, 0x73, 0x65, 0x48, 0x00, 0x52,
	0x11, 0x66, 0x72, 0x6f, 0x6d, 0x4b, 0x6f, 0x54, 0x66, 0x6f, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x69,
	0x72, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x74, 0x61, 0x69, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x48, 0x00, 0x52,
	0x05, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x3b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x69, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x69,
	0x72, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x69, 0x72, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x69, 0x72, 0x73, 0x12, 0x4c, 0x0a, 0x0e, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x5f, 0x6c, 0x61,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x6f,
	0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x69, 0x67, 0x68,
	0x74, 0x53, 0x74, 0x61, 0x69, 0x72, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x4c, 0x61, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x0d, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x12, 0x46, 0x0a, 0x0b, 0x74, 0x6f, 0x70, 0x5f, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x69,
	0x72, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x4c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x74,
	0x6f, 0x70, 0x4c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x1a, 0x2d, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x69, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x69, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x04, 0x72, 0x69, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x75, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x03, 0x72, 0x75, 0x6e, 0x1a, 0xb1, 0x01, 0x0a, 0x07, 0x4c, 0x61, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x12, 0x52, 0x0a, 0x1b, 0x73, 0x74, 0x61, 0x69, 0x72, 0x73, 0x5f, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x6f, 0x73, 0x64,
	0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x45, 0x33, 0x50, 0x6f, 0x73, 0x65, 0x52, 0x18,
	0x73, 0x74, 0x61, 0x69, 0x72, 0x73, 0x54, 0x66, 0x6f, 0x72, 0x6d, 0x4c, 0x61, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x61, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0e, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x74, 0x58, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x74, 0x5f, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x6c, 0x61,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x74, 0x59, 0x42, 0x0a, 0x0a, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x40, 0x42, 0x0b, 0x53, 0x74, 0x61, 0x69,
	0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x68, 0x73, 0x73, 0x6e, 0x76, 0x2f, 0x73, 0x70, 0x6f, 0x74, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x65, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x2f, 0x67, 0x6f, 0x2f,
	0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_bosdyn_api_stairs_proto_rawDescOnce sync.Once
	file_bosdyn_api_stairs_proto_rawDescData = file_bosdyn_api_stairs_proto_rawDesc
)

func file_bosdyn_api_stairs_proto_rawDescGZIP() []byte {
	file_bosdyn_api_stairs_proto_rawDescOnce.Do(func() {
		file_bosdyn_api_stairs_proto_rawDescData = protoimpl.X.CompressGZIP(file_bosdyn_api_stairs_proto_rawDescData)
	})
	return file_bosdyn_api_stairs_proto_rawDescData
}

var file_bosdyn_api_stairs_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_bosdyn_api_stairs_proto_goTypes = []interface{}{
	(*StairTransform)(nil),            // 0: bosdyn.api.StairTransform
	(*StraightStaircase)(nil),         // 1: bosdyn.api.StraightStaircase
	(*StraightStaircase_Stair)(nil),   // 2: bosdyn.api.StraightStaircase.Stair
	(*StraightStaircase_Landing)(nil), // 3: bosdyn.api.StraightStaircase.Landing
	(*SE3Pose)(nil),                   // 4: bosdyn.api.SE3Pose
}
var file_bosdyn_api_stairs_proto_depIdxs = []int32{
	4, // 0: bosdyn.api.StairTransform.frame_tform_stairs:type_name -> bosdyn.api.SE3Pose
	4, // 1: bosdyn.api.StraightStaircase.from_ko_tform_stairs:type_name -> bosdyn.api.SE3Pose
	0, // 2: bosdyn.api.StraightStaircase.tform:type_name -> bosdyn.api.StairTransform
	2, // 3: bosdyn.api.StraightStaircase.stairs:type_name -> bosdyn.api.StraightStaircase.Stair
	3, // 4: bosdyn.api.StraightStaircase.bottom_landing:type_name -> bosdyn.api.StraightStaircase.Landing
	3, // 5: bosdyn.api.StraightStaircase.top_landing:type_name -> bosdyn.api.StraightStaircase.Landing
	4, // 6: bosdyn.api.StraightStaircase.Landing.stairs_tform_landing_center:type_name -> bosdyn.api.SE3Pose
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_bosdyn_api_stairs_proto_init() }
func file_bosdyn_api_stairs_proto_init() {
	if File_bosdyn_api_stairs_proto != nil {
		return
	}
	file_bosdyn_api_geometry_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bosdyn_api_stairs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StairTransform); i {
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
		file_bosdyn_api_stairs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StraightStaircase); i {
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
		file_bosdyn_api_stairs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StraightStaircase_Stair); i {
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
		file_bosdyn_api_stairs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StraightStaircase_Landing); i {
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
	file_bosdyn_api_stairs_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*StraightStaircase_FromKoTformStairs)(nil),
		(*StraightStaircase_Tform)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bosdyn_api_stairs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bosdyn_api_stairs_proto_goTypes,
		DependencyIndexes: file_bosdyn_api_stairs_proto_depIdxs,
		MessageInfos:      file_bosdyn_api_stairs_proto_msgTypes,
	}.Build()
	File_bosdyn_api_stairs_proto = out.File
	file_bosdyn_api_stairs_proto_rawDesc = nil
	file_bosdyn_api_stairs_proto_goTypes = nil
	file_bosdyn_api_stairs_proto_depIdxs = nil
}
