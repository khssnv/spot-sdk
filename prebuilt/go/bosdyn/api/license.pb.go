// Copyright (c) 2022 Boston Dynamics, Inc.  All rights reserved.
//
// Downloading, reproducing, distributing or otherwise using the SDK Software
// is subject to the terms and conditions of the Boston Dynamics Software
// Development Kit License (20191101-BDSDK-SL).

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: bosdyn/api/license.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LicenseInfo_Status int32

const (
	LicenseInfo_STATUS_UNKNOWN         LicenseInfo_Status = 0
	LicenseInfo_STATUS_VALID           LicenseInfo_Status = 1
	LicenseInfo_STATUS_EXPIRED         LicenseInfo_Status = 2
	LicenseInfo_STATUS_NOT_YET_VALID   LicenseInfo_Status = 3
	LicenseInfo_STATUS_MALFORMED       LicenseInfo_Status = 4
	LicenseInfo_STATUS_SERIAL_MISMATCH LicenseInfo_Status = 5
	LicenseInfo_STATUS_NO_LICENSE      LicenseInfo_Status = 6
)

// Enum value maps for LicenseInfo_Status.
var (
	LicenseInfo_Status_name = map[int32]string{
		0: "STATUS_UNKNOWN",
		1: "STATUS_VALID",
		2: "STATUS_EXPIRED",
		3: "STATUS_NOT_YET_VALID",
		4: "STATUS_MALFORMED",
		5: "STATUS_SERIAL_MISMATCH",
		6: "STATUS_NO_LICENSE",
	}
	LicenseInfo_Status_value = map[string]int32{
		"STATUS_UNKNOWN":         0,
		"STATUS_VALID":           1,
		"STATUS_EXPIRED":         2,
		"STATUS_NOT_YET_VALID":   3,
		"STATUS_MALFORMED":       4,
		"STATUS_SERIAL_MISMATCH": 5,
		"STATUS_NO_LICENSE":      6,
	}
)

func (x LicenseInfo_Status) Enum() *LicenseInfo_Status {
	p := new(LicenseInfo_Status)
	*p = x
	return p
}

func (x LicenseInfo_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LicenseInfo_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_bosdyn_api_license_proto_enumTypes[0].Descriptor()
}

func (LicenseInfo_Status) Type() protoreflect.EnumType {
	return &file_bosdyn_api_license_proto_enumTypes[0]
}

func (x LicenseInfo_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LicenseInfo_Status.Descriptor instead.
func (LicenseInfo_Status) EnumDescriptor() ([]byte, []int) {
	return file_bosdyn_api_license_proto_rawDescGZIP(), []int{0, 0}
}

type LicenseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status of the uploaded license for this robot.
	Status LicenseInfo_Status `protobuf:"varint,1,opt,name=status,proto3,enum=bosdyn.api.LicenseInfo_Status" json:"status,omitempty"`
	// Unique license number.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Serial number of the robot this license covers.
	RobotSerial string `protobuf:"bytes,3,opt,name=robot_serial,json=robotSerial,proto3" json:"robot_serial,omitempty"`
	// The license is not valid for use for any dates before this timestamp.
	NotValidBefore *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=not_valid_before,json=notValidBefore,proto3" json:"not_valid_before,omitempty"`
	// The license is not valid for use for any dates after this timestamp.
	NotValidAfter *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=not_valid_after,json=notValidAfter,proto3" json:"not_valid_after,omitempty"`
	/// Human readable list of licensed features included for this license.
	LicensedFeatures []string `protobuf:"bytes,6,rep,name=licensed_features,json=licensedFeatures,proto3" json:"licensed_features,omitempty"`
}

func (x *LicenseInfo) Reset() {
	*x = LicenseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_license_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LicenseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LicenseInfo) ProtoMessage() {}

func (x *LicenseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_license_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LicenseInfo.ProtoReflect.Descriptor instead.
func (*LicenseInfo) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_license_proto_rawDescGZIP(), []int{0}
}

func (x *LicenseInfo) GetStatus() LicenseInfo_Status {
	if x != nil {
		return x.Status
	}
	return LicenseInfo_STATUS_UNKNOWN
}

func (x *LicenseInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LicenseInfo) GetRobotSerial() string {
	if x != nil {
		return x.RobotSerial
	}
	return ""
}

func (x *LicenseInfo) GetNotValidBefore() *timestamppb.Timestamp {
	if x != nil {
		return x.NotValidBefore
	}
	return nil
}

func (x *LicenseInfo) GetNotValidAfter() *timestamppb.Timestamp {
	if x != nil {
		return x.NotValidAfter
	}
	return nil
}

func (x *LicenseInfo) GetLicensedFeatures() []string {
	if x != nil {
		return x.LicensedFeatures
	}
	return nil
}

//
type GetLicenseInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"` // Common request header.
}

func (x *GetLicenseInfoRequest) Reset() {
	*x = GetLicenseInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_license_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLicenseInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLicenseInfoRequest) ProtoMessage() {}

func (x *GetLicenseInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_license_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLicenseInfoRequest.ProtoReflect.Descriptor instead.
func (*GetLicenseInfoRequest) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_license_proto_rawDescGZIP(), []int{1}
}

func (x *GetLicenseInfoRequest) GetHeader() *RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

type GetLicenseInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"` // Common response header
	// The details about the current license that is uploaded to the robot.
	License *LicenseInfo `protobuf:"bytes,2,opt,name=license,proto3" json:"license,omitempty"`
}

func (x *GetLicenseInfoResponse) Reset() {
	*x = GetLicenseInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_license_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLicenseInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLicenseInfoResponse) ProtoMessage() {}

func (x *GetLicenseInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_license_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLicenseInfoResponse.ProtoReflect.Descriptor instead.
func (*GetLicenseInfoResponse) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_license_proto_rawDescGZIP(), []int{2}
}

func (x *GetLicenseInfoResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GetLicenseInfoResponse) GetLicense() *LicenseInfo {
	if x != nil {
		return x.License
	}
	return nil
}

type GetFeatureEnabledRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common request header.
	Header *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// Check if specific named features are enabled on the robot under the currently
	// loaded license.
	FeatureCodes []string `protobuf:"bytes,2,rep,name=feature_codes,json=featureCodes,proto3" json:"feature_codes,omitempty"`
}

func (x *GetFeatureEnabledRequest) Reset() {
	*x = GetFeatureEnabledRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_license_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeatureEnabledRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeatureEnabledRequest) ProtoMessage() {}

func (x *GetFeatureEnabledRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_license_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeatureEnabledRequest.ProtoReflect.Descriptor instead.
func (*GetFeatureEnabledRequest) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_license_proto_rawDescGZIP(), []int{3}
}

func (x *GetFeatureEnabledRequest) GetHeader() *RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GetFeatureEnabledRequest) GetFeatureCodes() []string {
	if x != nil {
		return x.FeatureCodes
	}
	return nil
}

type GetFeatureEnabledResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common response header.
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The resulting map showing the feature name (as the map key) and a boolean indicating
	// if the feature is enabled with this license (as the map value).
	FeatureEnabled map[string]bool `protobuf:"bytes,2,rep,name=feature_enabled,json=featureEnabled,proto3" json:"feature_enabled,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *GetFeatureEnabledResponse) Reset() {
	*x = GetFeatureEnabledResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_license_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeatureEnabledResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeatureEnabledResponse) ProtoMessage() {}

func (x *GetFeatureEnabledResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_license_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeatureEnabledResponse.ProtoReflect.Descriptor instead.
func (*GetFeatureEnabledResponse) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_license_proto_rawDescGZIP(), []int{4}
}

func (x *GetFeatureEnabledResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GetFeatureEnabledResponse) GetFeatureEnabled() map[string]bool {
	if x != nil {
		return x.FeatureEnabled
	}
	return nil
}

var File_bosdyn_api_license_proto protoreflect.FileDescriptor

var file_bosdyn_api_license_proto_rawDesc = []byte{
	0x0a, 0x18, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62, 0x6f, 0x73, 0x64,
	0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x17, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd7, 0x03, 0x0a, 0x0b, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1e, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x6f, 0x62, 0x6f,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x72, 0x6f, 0x62, 0x6f, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x44, 0x0a, 0x10, 0x6e,
	0x6f, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0e, 0x6e, 0x6f, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x42, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x12, 0x42, 0x0a, 0x0f, 0x6e, 0x6f, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6e, 0x6f, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x64, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x10, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x64, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x22, 0xa5, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a,
	0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x58,
	0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x59, 0x45, 0x54, 0x5f, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10,
	0x03, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4d, 0x41, 0x4c, 0x46,
	0x4f, 0x52, 0x4d, 0x45, 0x44, 0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x53, 0x45, 0x52, 0x49, 0x41, 0x4c, 0x5f, 0x4d, 0x49, 0x53, 0x4d, 0x41, 0x54, 0x43,
	0x48, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f,
	0x5f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x10, 0x06, 0x22, 0x4a, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x7f, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x32, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07,
	0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x22, 0x72, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x22, 0xf6, 0x01, 0x0a, 0x19,
	0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6f, 0x73, 0x64,
	0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x62, 0x0a,
	0x0f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x1a, 0x41, 0x0a, 0x13, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x41, 0x42, 0x0c, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x68, 0x73, 0x73, 0x6e, 0x76, 0x2f, 0x73, 0x70, 0x6f, 0x74, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x70, 0x72, 0x65, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x6f, 0x73,
	0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bosdyn_api_license_proto_rawDescOnce sync.Once
	file_bosdyn_api_license_proto_rawDescData = file_bosdyn_api_license_proto_rawDesc
)

func file_bosdyn_api_license_proto_rawDescGZIP() []byte {
	file_bosdyn_api_license_proto_rawDescOnce.Do(func() {
		file_bosdyn_api_license_proto_rawDescData = protoimpl.X.CompressGZIP(file_bosdyn_api_license_proto_rawDescData)
	})
	return file_bosdyn_api_license_proto_rawDescData
}

var file_bosdyn_api_license_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bosdyn_api_license_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_bosdyn_api_license_proto_goTypes = []interface{}{
	(LicenseInfo_Status)(0),           // 0: bosdyn.api.LicenseInfo.Status
	(*LicenseInfo)(nil),               // 1: bosdyn.api.LicenseInfo
	(*GetLicenseInfoRequest)(nil),     // 2: bosdyn.api.GetLicenseInfoRequest
	(*GetLicenseInfoResponse)(nil),    // 3: bosdyn.api.GetLicenseInfoResponse
	(*GetFeatureEnabledRequest)(nil),  // 4: bosdyn.api.GetFeatureEnabledRequest
	(*GetFeatureEnabledResponse)(nil), // 5: bosdyn.api.GetFeatureEnabledResponse
	nil,                               // 6: bosdyn.api.GetFeatureEnabledResponse.FeatureEnabledEntry
	(*timestamppb.Timestamp)(nil),     // 7: google.protobuf.Timestamp
	(*RequestHeader)(nil),             // 8: bosdyn.api.RequestHeader
	(*ResponseHeader)(nil),            // 9: bosdyn.api.ResponseHeader
}
var file_bosdyn_api_license_proto_depIdxs = []int32{
	0, // 0: bosdyn.api.LicenseInfo.status:type_name -> bosdyn.api.LicenseInfo.Status
	7, // 1: bosdyn.api.LicenseInfo.not_valid_before:type_name -> google.protobuf.Timestamp
	7, // 2: bosdyn.api.LicenseInfo.not_valid_after:type_name -> google.protobuf.Timestamp
	8, // 3: bosdyn.api.GetLicenseInfoRequest.header:type_name -> bosdyn.api.RequestHeader
	9, // 4: bosdyn.api.GetLicenseInfoResponse.header:type_name -> bosdyn.api.ResponseHeader
	1, // 5: bosdyn.api.GetLicenseInfoResponse.license:type_name -> bosdyn.api.LicenseInfo
	8, // 6: bosdyn.api.GetFeatureEnabledRequest.header:type_name -> bosdyn.api.RequestHeader
	9, // 7: bosdyn.api.GetFeatureEnabledResponse.header:type_name -> bosdyn.api.ResponseHeader
	6, // 8: bosdyn.api.GetFeatureEnabledResponse.feature_enabled:type_name -> bosdyn.api.GetFeatureEnabledResponse.FeatureEnabledEntry
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_bosdyn_api_license_proto_init() }
func file_bosdyn_api_license_proto_init() {
	if File_bosdyn_api_license_proto != nil {
		return
	}
	file_bosdyn_api_header_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bosdyn_api_license_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LicenseInfo); i {
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
		file_bosdyn_api_license_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLicenseInfoRequest); i {
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
		file_bosdyn_api_license_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLicenseInfoResponse); i {
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
		file_bosdyn_api_license_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeatureEnabledRequest); i {
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
		file_bosdyn_api_license_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeatureEnabledResponse); i {
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
			RawDescriptor: file_bosdyn_api_license_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bosdyn_api_license_proto_goTypes,
		DependencyIndexes: file_bosdyn_api_license_proto_depIdxs,
		EnumInfos:         file_bosdyn_api_license_proto_enumTypes,
		MessageInfos:      file_bosdyn_api_license_proto_msgTypes,
	}.Build()
	File_bosdyn_api_license_proto = out.File
	file_bosdyn_api_license_proto_rawDesc = nil
	file_bosdyn_api_license_proto_goTypes = nil
	file_bosdyn_api_license_proto_depIdxs = nil
}
