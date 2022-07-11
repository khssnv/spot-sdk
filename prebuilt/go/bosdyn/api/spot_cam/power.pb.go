// Copyright (c) 2021 Boston Dynamics, Inc.  All rights reserved.
//
// Downloading, reproducing, distributing or otherwise using the SDK Software
// is subject to the terms and conditions of the Boston Dynamics Software
// Development Kit License (20191101-BDSDK-SL).

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: bosdyn/api/spot_cam/power.proto

package spot_cam

import (
	api "github.com/khssnv/spot-sdk/prebuilt/go/bosdyn/api"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Power on or off of components of the SpotCam.
type PowerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//these switches are 'true' for power-on, 'false' for power-off
	Ptz         *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=ptz,proto3" json:"ptz,omitempty"`
	Aux1        *wrapperspb.BoolValue `protobuf:"bytes,3,opt,name=aux1,proto3" json:"aux1,omitempty"`
	Aux2        *wrapperspb.BoolValue `protobuf:"bytes,4,opt,name=aux2,proto3" json:"aux2,omitempty"`
	ExternalMic *wrapperspb.BoolValue `protobuf:"bytes,5,opt,name=external_mic,json=externalMic,proto3" json:"external_mic,omitempty"`
}

func (x *PowerStatus) Reset() {
	*x = PowerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_spot_cam_power_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PowerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PowerStatus) ProtoMessage() {}

func (x *PowerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_spot_cam_power_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PowerStatus.ProtoReflect.Descriptor instead.
func (*PowerStatus) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_spot_cam_power_proto_rawDescGZIP(), []int{0}
}

func (x *PowerStatus) GetPtz() *wrapperspb.BoolValue {
	if x != nil {
		return x.Ptz
	}
	return nil
}

func (x *PowerStatus) GetAux1() *wrapperspb.BoolValue {
	if x != nil {
		return x.Aux1
	}
	return nil
}

func (x *PowerStatus) GetAux2() *wrapperspb.BoolValue {
	if x != nil {
		return x.Aux2
	}
	return nil
}

func (x *PowerStatus) GetExternalMic() *wrapperspb.BoolValue {
	if x != nil {
		return x.ExternalMic
	}
	return nil
}

// Request component power status.
type GetPowerStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common request header.
	Header *api.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *GetPowerStatusRequest) Reset() {
	*x = GetPowerStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_spot_cam_power_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPowerStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPowerStatusRequest) ProtoMessage() {}

func (x *GetPowerStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_spot_cam_power_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPowerStatusRequest.ProtoReflect.Descriptor instead.
func (*GetPowerStatusRequest) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_spot_cam_power_proto_rawDescGZIP(), []int{1}
}

func (x *GetPowerStatusRequest) GetHeader() *api.RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

// Provides the power status of all components.
type GetPowerStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common response header.
	Header *api.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	//status indicates the power status of the controllable devices
	//'true' for power-on, 'false' for power-off
	Status *PowerStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetPowerStatusResponse) Reset() {
	*x = GetPowerStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_spot_cam_power_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPowerStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPowerStatusResponse) ProtoMessage() {}

func (x *GetPowerStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_spot_cam_power_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPowerStatusResponse.ProtoReflect.Descriptor instead.
func (*GetPowerStatusResponse) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_spot_cam_power_proto_rawDescGZIP(), []int{2}
}

func (x *GetPowerStatusResponse) GetHeader() *api.ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GetPowerStatusResponse) GetStatus() *PowerStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// Turn components on or off.
type SetPowerStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common request header.
	Header *api.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	//status indicates the requested power status of the controllable devices
	//'true' for power-on, 'false' for power-off
	Status *PowerStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SetPowerStatusRequest) Reset() {
	*x = SetPowerStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_spot_cam_power_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPowerStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPowerStatusRequest) ProtoMessage() {}

func (x *SetPowerStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_spot_cam_power_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPowerStatusRequest.ProtoReflect.Descriptor instead.
func (*SetPowerStatusRequest) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_spot_cam_power_proto_rawDescGZIP(), []int{3}
}

func (x *SetPowerStatusRequest) GetHeader() *api.RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *SetPowerStatusRequest) GetStatus() *PowerStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// Result of turning components on or off.
type SetPowerStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common response header.
	Header *api.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	//status indicates the requested changes upon success
	//'true' for power-on, 'false' for power-off
	Status *PowerStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SetPowerStatusResponse) Reset() {
	*x = SetPowerStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_spot_cam_power_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPowerStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPowerStatusResponse) ProtoMessage() {}

func (x *SetPowerStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_spot_cam_power_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPowerStatusResponse.ProtoReflect.Descriptor instead.
func (*SetPowerStatusResponse) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_spot_cam_power_proto_rawDescGZIP(), []int{4}
}

func (x *SetPowerStatusResponse) GetHeader() *api.ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *SetPowerStatusResponse) GetStatus() *PowerStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// Turn components off and then back on without needing two separate requests.
type CyclePowerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common request header.
	Header *api.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	//status indicates the devices for which cycle-power is requested
	//'true' for cycle-power, else no effect
	//power cycle will not be performed on a given device if its state is power-off prior to this call
	Status *PowerStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CyclePowerRequest) Reset() {
	*x = CyclePowerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_spot_cam_power_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CyclePowerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CyclePowerRequest) ProtoMessage() {}

func (x *CyclePowerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_spot_cam_power_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CyclePowerRequest.ProtoReflect.Descriptor instead.
func (*CyclePowerRequest) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_spot_cam_power_proto_rawDescGZIP(), []int{5}
}

func (x *CyclePowerRequest) GetHeader() *api.RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *CyclePowerRequest) GetStatus() *PowerStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// Result of power cycling components.
type CyclePowerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common response header.
	Header *api.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	//status indicates the power status of the controllable devices after a successful power cycle
	//'true' for power-on, 'false' for power-off
	Status *PowerStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CyclePowerResponse) Reset() {
	*x = CyclePowerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_spot_cam_power_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CyclePowerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CyclePowerResponse) ProtoMessage() {}

func (x *CyclePowerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_spot_cam_power_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CyclePowerResponse.ProtoReflect.Descriptor instead.
func (*CyclePowerResponse) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_spot_cam_power_proto_rawDescGZIP(), []int{6}
}

func (x *CyclePowerResponse) GetHeader() *api.ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *CyclePowerResponse) GetStatus() *PowerStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_bosdyn_api_spot_cam_power_proto protoreflect.FileDescriptor

var file_bosdyn_api_spot_cam_power_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x70, 0x6f,
	0x74, 0x5f, 0x63, 0x61, 0x6d, 0x2f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x70,
	0x6f, 0x74, 0x5f, 0x63, 0x61, 0x6d, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xda, 0x01, 0x0a, 0x0b, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x2c, 0x0a, 0x03, 0x70, 0x74, 0x7a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x70, 0x74, 0x7a, 0x12, 0x2e, 0x0a,
	0x04, 0x61, 0x75, 0x78, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x61, 0x75, 0x78, 0x31, 0x12, 0x2e, 0x0a,
	0x04, 0x61, 0x75, 0x78, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x61, 0x75, 0x78, 0x32, 0x12, 0x3d, 0x0a,
	0x0c, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x69, 0x63, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d, 0x69, 0x63, 0x22, 0x4a, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x86, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x5f, 0x63, 0x61, 0x6d, 0x2e, 0x50, 0x6f,
	0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x84, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x6f,
	0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x38,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x70, 0x6f, 0x74,
	0x5f, 0x63, 0x61, 0x6d, 0x2e, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x86, 0x01, 0x0a, 0x16, 0x53, 0x65, 0x74,
	0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x5f, 0x63, 0x61, 0x6d, 0x2e, 0x50, 0x6f,
	0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x80, 0x01, 0x0a, 0x11, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x50, 0x6f, 0x77, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x6f, 0x73,
	0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x5f, 0x63, 0x61, 0x6d,
	0x2e, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x12, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x50, 0x6f,
	0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6f,
	0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x70, 0x6f,
	0x74, 0x5f, 0x63, 0x61, 0x6d, 0x2e, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x48, 0x42, 0x0a, 0x50, 0x6f, 0x77,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x68, 0x73, 0x73, 0x6e, 0x76, 0x2f, 0x73, 0x70, 0x6f, 0x74, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x65, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x2f, 0x67, 0x6f, 0x2f,
	0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x70, 0x6f, 0x74, 0x5f,
	0x63, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bosdyn_api_spot_cam_power_proto_rawDescOnce sync.Once
	file_bosdyn_api_spot_cam_power_proto_rawDescData = file_bosdyn_api_spot_cam_power_proto_rawDesc
)

func file_bosdyn_api_spot_cam_power_proto_rawDescGZIP() []byte {
	file_bosdyn_api_spot_cam_power_proto_rawDescOnce.Do(func() {
		file_bosdyn_api_spot_cam_power_proto_rawDescData = protoimpl.X.CompressGZIP(file_bosdyn_api_spot_cam_power_proto_rawDescData)
	})
	return file_bosdyn_api_spot_cam_power_proto_rawDescData
}

var file_bosdyn_api_spot_cam_power_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_bosdyn_api_spot_cam_power_proto_goTypes = []interface{}{
	(*PowerStatus)(nil),            // 0: bosdyn.api.spot_cam.PowerStatus
	(*GetPowerStatusRequest)(nil),  // 1: bosdyn.api.spot_cam.GetPowerStatusRequest
	(*GetPowerStatusResponse)(nil), // 2: bosdyn.api.spot_cam.GetPowerStatusResponse
	(*SetPowerStatusRequest)(nil),  // 3: bosdyn.api.spot_cam.SetPowerStatusRequest
	(*SetPowerStatusResponse)(nil), // 4: bosdyn.api.spot_cam.SetPowerStatusResponse
	(*CyclePowerRequest)(nil),      // 5: bosdyn.api.spot_cam.CyclePowerRequest
	(*CyclePowerResponse)(nil),     // 6: bosdyn.api.spot_cam.CyclePowerResponse
	(*wrapperspb.BoolValue)(nil),   // 7: google.protobuf.BoolValue
	(*api.RequestHeader)(nil),      // 8: bosdyn.api.RequestHeader
	(*api.ResponseHeader)(nil),     // 9: bosdyn.api.ResponseHeader
}
var file_bosdyn_api_spot_cam_power_proto_depIdxs = []int32{
	7,  // 0: bosdyn.api.spot_cam.PowerStatus.ptz:type_name -> google.protobuf.BoolValue
	7,  // 1: bosdyn.api.spot_cam.PowerStatus.aux1:type_name -> google.protobuf.BoolValue
	7,  // 2: bosdyn.api.spot_cam.PowerStatus.aux2:type_name -> google.protobuf.BoolValue
	7,  // 3: bosdyn.api.spot_cam.PowerStatus.external_mic:type_name -> google.protobuf.BoolValue
	8,  // 4: bosdyn.api.spot_cam.GetPowerStatusRequest.header:type_name -> bosdyn.api.RequestHeader
	9,  // 5: bosdyn.api.spot_cam.GetPowerStatusResponse.header:type_name -> bosdyn.api.ResponseHeader
	0,  // 6: bosdyn.api.spot_cam.GetPowerStatusResponse.status:type_name -> bosdyn.api.spot_cam.PowerStatus
	8,  // 7: bosdyn.api.spot_cam.SetPowerStatusRequest.header:type_name -> bosdyn.api.RequestHeader
	0,  // 8: bosdyn.api.spot_cam.SetPowerStatusRequest.status:type_name -> bosdyn.api.spot_cam.PowerStatus
	9,  // 9: bosdyn.api.spot_cam.SetPowerStatusResponse.header:type_name -> bosdyn.api.ResponseHeader
	0,  // 10: bosdyn.api.spot_cam.SetPowerStatusResponse.status:type_name -> bosdyn.api.spot_cam.PowerStatus
	8,  // 11: bosdyn.api.spot_cam.CyclePowerRequest.header:type_name -> bosdyn.api.RequestHeader
	0,  // 12: bosdyn.api.spot_cam.CyclePowerRequest.status:type_name -> bosdyn.api.spot_cam.PowerStatus
	9,  // 13: bosdyn.api.spot_cam.CyclePowerResponse.header:type_name -> bosdyn.api.ResponseHeader
	0,  // 14: bosdyn.api.spot_cam.CyclePowerResponse.status:type_name -> bosdyn.api.spot_cam.PowerStatus
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_bosdyn_api_spot_cam_power_proto_init() }
func file_bosdyn_api_spot_cam_power_proto_init() {
	if File_bosdyn_api_spot_cam_power_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bosdyn_api_spot_cam_power_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PowerStatus); i {
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
		file_bosdyn_api_spot_cam_power_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPowerStatusRequest); i {
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
		file_bosdyn_api_spot_cam_power_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPowerStatusResponse); i {
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
		file_bosdyn_api_spot_cam_power_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPowerStatusRequest); i {
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
		file_bosdyn_api_spot_cam_power_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPowerStatusResponse); i {
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
		file_bosdyn_api_spot_cam_power_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CyclePowerRequest); i {
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
		file_bosdyn_api_spot_cam_power_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CyclePowerResponse); i {
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
			RawDescriptor: file_bosdyn_api_spot_cam_power_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bosdyn_api_spot_cam_power_proto_goTypes,
		DependencyIndexes: file_bosdyn_api_spot_cam_power_proto_depIdxs,
		MessageInfos:      file_bosdyn_api_spot_cam_power_proto_msgTypes,
	}.Build()
	File_bosdyn_api_spot_cam_power_proto = out.File
	file_bosdyn_api_spot_cam_power_proto_rawDesc = nil
	file_bosdyn_api_spot_cam_power_proto_goTypes = nil
	file_bosdyn_api_spot_cam_power_proto_depIdxs = nil
}