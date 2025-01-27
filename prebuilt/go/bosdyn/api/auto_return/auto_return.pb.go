// Copyright (c) 2022 Boston Dynamics, Inc.  All rights reserved.
//
// Downloading, reproducing, distributing or otherwise using the SDK Software
// is subject to the terms and conditions of the Boston Dynamics Software
// Development Kit License (20191101-BDSDK-SL).

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: bosdyn/api/auto_return/auto_return.proto

package auto_return

import (
	api "github.com/khssnv/spot-sdk/prebuilt/go/bosdyn/api"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConfigureResponse_Status int32

const (
	ConfigureResponse_STATUS_UNKNOWN        ConfigureResponse_Status = 0
	ConfigureResponse_STATUS_OK             ConfigureResponse_Status = 1
	ConfigureResponse_STATUS_INVALID_PARAMS ConfigureResponse_Status = 2
)

// Enum value maps for ConfigureResponse_Status.
var (
	ConfigureResponse_Status_name = map[int32]string{
		0: "STATUS_UNKNOWN",
		1: "STATUS_OK",
		2: "STATUS_INVALID_PARAMS",
	}
	ConfigureResponse_Status_value = map[string]int32{
		"STATUS_UNKNOWN":        0,
		"STATUS_OK":             1,
		"STATUS_INVALID_PARAMS": 2,
	}
)

func (x ConfigureResponse_Status) Enum() *ConfigureResponse_Status {
	p := new(ConfigureResponse_Status)
	*p = x
	return p
}

func (x ConfigureResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConfigureResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_bosdyn_api_auto_return_auto_return_proto_enumTypes[0].Descriptor()
}

func (ConfigureResponse_Status) Type() protoreflect.EnumType {
	return &file_bosdyn_api_auto_return_auto_return_proto_enumTypes[0]
}

func (x ConfigureResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConfigureResponse_Status.Descriptor instead.
func (ConfigureResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_bosdyn_api_auto_return_auto_return_proto_rawDescGZIP(), []int{2, 0}
}

// Parameters to AutoReturn actions.
type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Robot-specific mobility parameters to use.
	// For example, see bosdyn.api.spot.MobilityParams.
	MobilityParams *anypb.Any `protobuf:"bytes,1,opt,name=mobility_params,json=mobilityParams,proto3" json:"mobility_params,omitempty"`
	// Allow AutoReturn to move the robot this far in the XY plane from the location where
	// AutoReturn started. Travel along the Z axis (which is gravity-aligned) does not count.
	// Must be > 0.
	MaxDisplacement float32 `protobuf:"fixed32,2,opt,name=max_displacement,json=maxDisplacement,proto3" json:"max_displacement,omitempty"` // meters
	// Optionally specify the maximum amount of time AutoReturn can take.
	// If this duration is exceeded, AutoReturn will stop trying to move the robot.
	// Omit to try indefinitely. Robot may become stuck and never take other comms loss actions!
	MaxDuration *durationpb.Duration `protobuf:"bytes,3,opt,name=max_duration,json=maxDuration,proto3" json:"max_duration,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_auto_return_auto_return_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_auto_return_auto_return_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_auto_return_auto_return_proto_rawDescGZIP(), []int{0}
}

func (x *Params) GetMobilityParams() *anypb.Any {
	if x != nil {
		return x.MobilityParams
	}
	return nil
}

func (x *Params) GetMaxDisplacement() float32 {
	if x != nil {
		return x.MaxDisplacement
	}
	return 0
}

func (x *Params) GetMaxDuration() *durationpb.Duration {
	if x != nil {
		return x.MaxDuration
	}
	return nil
}

// Configure the AutoReturn system.
type ConfigureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common request header.
	Header *api.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// Leases that AutoReturn may use to accomplish its goals when AutoReturn automatically
	// triggers. If left empty, AutoReturn will not automatically trigger.
	Leases []*api.Lease `protobuf:"bytes,2,rep,name=leases,proto3" json:"leases,omitempty"`
	// Parameters to use when AutoReturn automatically triggers.
	Params *Params `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
	// Forget any buffered locations the robot may be remembering. Defaults to false.
	// Set to true if the robot has just crossed an area it should not traverse in AutoReturn.
	ClearBuffer bool `protobuf:"varint,4,opt,name=clear_buffer,json=clearBuffer,proto3" json:"clear_buffer,omitempty"`
}

func (x *ConfigureRequest) Reset() {
	*x = ConfigureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_auto_return_auto_return_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureRequest) ProtoMessage() {}

func (x *ConfigureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_auto_return_auto_return_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureRequest.ProtoReflect.Descriptor instead.
func (*ConfigureRequest) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_auto_return_auto_return_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigureRequest) GetHeader() *api.RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ConfigureRequest) GetLeases() []*api.Lease {
	if x != nil {
		return x.Leases
	}
	return nil
}

func (x *ConfigureRequest) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *ConfigureRequest) GetClearBuffer() bool {
	if x != nil {
		return x.ClearBuffer
	}
	return false
}

// Response to a configuration request.
type ConfigureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common response header.
	Header *api.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// Return status for the request.
	Status ConfigureResponse_Status `protobuf:"varint,2,opt,name=status,proto3,enum=bosdyn.api.auto_return.ConfigureResponse_Status" json:"status,omitempty"`
	// If status is STATUS_INVALID_PARAMS, this contains the settings that were invalid.
	InvalidParams *Params `protobuf:"bytes,3,opt,name=invalid_params,json=invalidParams,proto3" json:"invalid_params,omitempty"`
}

func (x *ConfigureResponse) Reset() {
	*x = ConfigureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_auto_return_auto_return_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureResponse) ProtoMessage() {}

func (x *ConfigureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_auto_return_auto_return_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureResponse.ProtoReflect.Descriptor instead.
func (*ConfigureResponse) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_auto_return_auto_return_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigureResponse) GetHeader() *api.ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ConfigureResponse) GetStatus() ConfigureResponse_Status {
	if x != nil {
		return x.Status
	}
	return ConfigureResponse_STATUS_UNKNOWN
}

func (x *ConfigureResponse) GetInvalidParams() *Params {
	if x != nil {
		return x.InvalidParams
	}
	return nil
}

// Ask for the current configuration.
type GetConfigurationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common request header.
	Header *api.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *GetConfigurationRequest) Reset() {
	*x = GetConfigurationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_auto_return_auto_return_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigurationRequest) ProtoMessage() {}

func (x *GetConfigurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_auto_return_auto_return_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigurationRequest.ProtoReflect.Descriptor instead.
func (*GetConfigurationRequest) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_auto_return_auto_return_proto_rawDescGZIP(), []int{3}
}

func (x *GetConfigurationRequest) GetHeader() *api.RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

// Response to a "get configuration" request.
type GetConfigurationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common response header.
	Header *api.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// A simple yes/no, will AutoReturn automatically trigger.
	Enabled bool `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The most recent successful ConfigureRequest.
	// Will be empty if service has not successfully been configured.
	Request *ConfigureRequest `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *GetConfigurationResponse) Reset() {
	*x = GetConfigurationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_auto_return_auto_return_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigurationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigurationResponse) ProtoMessage() {}

func (x *GetConfigurationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_auto_return_auto_return_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigurationResponse.ProtoReflect.Descriptor instead.
func (*GetConfigurationResponse) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_auto_return_auto_return_proto_rawDescGZIP(), []int{4}
}

func (x *GetConfigurationResponse) GetHeader() *api.ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GetConfigurationResponse) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *GetConfigurationResponse) GetRequest() *ConfigureRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

// Start AutoReturn behavior now.
type StartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common request header.
	Header *api.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_auto_return_auto_return_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_auto_return_auto_return_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRequest.ProtoReflect.Descriptor instead.
func (*StartRequest) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_auto_return_auto_return_proto_rawDescGZIP(), []int{5}
}

func (x *StartRequest) GetHeader() *api.RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

type StartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common response header.
	Header *api.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *StartResponse) Reset() {
	*x = StartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_auto_return_auto_return_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartResponse) ProtoMessage() {}

func (x *StartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_auto_return_auto_return_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartResponse.ProtoReflect.Descriptor instead.
func (*StartResponse) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_auto_return_auto_return_proto_rawDescGZIP(), []int{6}
}

func (x *StartResponse) GetHeader() *api.ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

var File_bosdyn_api_auto_return_auto_return_proto protoreflect.FileDescriptor

var file_bosdyn_api_auto_return_auto_return_proto_rawDesc = []byte{
	0x0a, 0x28, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74,
	0x6f, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x62, 0x6f, 0x73, 0x64,
	0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x62,
	0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0,
	0x01, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x3d, 0x0a, 0x0f, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xcb, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x06, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6f, 0x73, 0x64,
	0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x06, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x2e, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6c, 0x65, 0x61, 0x72, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x22,
	0xa0, 0x02, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x62, 0x6f, 0x73, 0x64,
	0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x45, 0x0a, 0x0e, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x6f,
	0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0d, 0x69, 0x6e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x46, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x53,
	0x10, 0x02, 0x22, 0x4c, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x22, 0xac, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x42, 0x0a, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x62,
	0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x41, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x31, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x22, 0x43, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x50, 0x42, 0x0f, 0x41, 0x75, 0x74, 0x6f, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x68, 0x73, 0x73, 0x6e, 0x76, 0x2f, 0x73, 0x70,
	0x6f, 0x74, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x65, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x2f,
	0x67, 0x6f, 0x2f, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75,
	0x74, 0x6f, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_bosdyn_api_auto_return_auto_return_proto_rawDescOnce sync.Once
	file_bosdyn_api_auto_return_auto_return_proto_rawDescData = file_bosdyn_api_auto_return_auto_return_proto_rawDesc
)

func file_bosdyn_api_auto_return_auto_return_proto_rawDescGZIP() []byte {
	file_bosdyn_api_auto_return_auto_return_proto_rawDescOnce.Do(func() {
		file_bosdyn_api_auto_return_auto_return_proto_rawDescData = protoimpl.X.CompressGZIP(file_bosdyn_api_auto_return_auto_return_proto_rawDescData)
	})
	return file_bosdyn_api_auto_return_auto_return_proto_rawDescData
}

var file_bosdyn_api_auto_return_auto_return_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bosdyn_api_auto_return_auto_return_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_bosdyn_api_auto_return_auto_return_proto_goTypes = []interface{}{
	(ConfigureResponse_Status)(0),    // 0: bosdyn.api.auto_return.ConfigureResponse.Status
	(*Params)(nil),                   // 1: bosdyn.api.auto_return.Params
	(*ConfigureRequest)(nil),         // 2: bosdyn.api.auto_return.ConfigureRequest
	(*ConfigureResponse)(nil),        // 3: bosdyn.api.auto_return.ConfigureResponse
	(*GetConfigurationRequest)(nil),  // 4: bosdyn.api.auto_return.GetConfigurationRequest
	(*GetConfigurationResponse)(nil), // 5: bosdyn.api.auto_return.GetConfigurationResponse
	(*StartRequest)(nil),             // 6: bosdyn.api.auto_return.StartRequest
	(*StartResponse)(nil),            // 7: bosdyn.api.auto_return.StartResponse
	(*anypb.Any)(nil),                // 8: google.protobuf.Any
	(*durationpb.Duration)(nil),      // 9: google.protobuf.Duration
	(*api.RequestHeader)(nil),        // 10: bosdyn.api.RequestHeader
	(*api.Lease)(nil),                // 11: bosdyn.api.Lease
	(*api.ResponseHeader)(nil),       // 12: bosdyn.api.ResponseHeader
}
var file_bosdyn_api_auto_return_auto_return_proto_depIdxs = []int32{
	8,  // 0: bosdyn.api.auto_return.Params.mobility_params:type_name -> google.protobuf.Any
	9,  // 1: bosdyn.api.auto_return.Params.max_duration:type_name -> google.protobuf.Duration
	10, // 2: bosdyn.api.auto_return.ConfigureRequest.header:type_name -> bosdyn.api.RequestHeader
	11, // 3: bosdyn.api.auto_return.ConfigureRequest.leases:type_name -> bosdyn.api.Lease
	1,  // 4: bosdyn.api.auto_return.ConfigureRequest.params:type_name -> bosdyn.api.auto_return.Params
	12, // 5: bosdyn.api.auto_return.ConfigureResponse.header:type_name -> bosdyn.api.ResponseHeader
	0,  // 6: bosdyn.api.auto_return.ConfigureResponse.status:type_name -> bosdyn.api.auto_return.ConfigureResponse.Status
	1,  // 7: bosdyn.api.auto_return.ConfigureResponse.invalid_params:type_name -> bosdyn.api.auto_return.Params
	10, // 8: bosdyn.api.auto_return.GetConfigurationRequest.header:type_name -> bosdyn.api.RequestHeader
	12, // 9: bosdyn.api.auto_return.GetConfigurationResponse.header:type_name -> bosdyn.api.ResponseHeader
	2,  // 10: bosdyn.api.auto_return.GetConfigurationResponse.request:type_name -> bosdyn.api.auto_return.ConfigureRequest
	10, // 11: bosdyn.api.auto_return.StartRequest.header:type_name -> bosdyn.api.RequestHeader
	12, // 12: bosdyn.api.auto_return.StartResponse.header:type_name -> bosdyn.api.ResponseHeader
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_bosdyn_api_auto_return_auto_return_proto_init() }
func file_bosdyn_api_auto_return_auto_return_proto_init() {
	if File_bosdyn_api_auto_return_auto_return_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bosdyn_api_auto_return_auto_return_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
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
		file_bosdyn_api_auto_return_auto_return_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigureRequest); i {
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
		file_bosdyn_api_auto_return_auto_return_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigureResponse); i {
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
		file_bosdyn_api_auto_return_auto_return_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigurationRequest); i {
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
		file_bosdyn_api_auto_return_auto_return_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigurationResponse); i {
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
		file_bosdyn_api_auto_return_auto_return_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRequest); i {
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
		file_bosdyn_api_auto_return_auto_return_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartResponse); i {
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
			RawDescriptor: file_bosdyn_api_auto_return_auto_return_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bosdyn_api_auto_return_auto_return_proto_goTypes,
		DependencyIndexes: file_bosdyn_api_auto_return_auto_return_proto_depIdxs,
		EnumInfos:         file_bosdyn_api_auto_return_auto_return_proto_enumTypes,
		MessageInfos:      file_bosdyn_api_auto_return_auto_return_proto_msgTypes,
	}.Build()
	File_bosdyn_api_auto_return_auto_return_proto = out.File
	file_bosdyn_api_auto_return_auto_return_proto_rawDesc = nil
	file_bosdyn_api_auto_return_auto_return_proto_goTypes = nil
	file_bosdyn_api_auto_return_auto_return_proto_depIdxs = nil
}
