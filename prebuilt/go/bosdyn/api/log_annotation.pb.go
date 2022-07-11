// Copyright (c) 2021 Boston Dynamics, Inc.  All rights reserved.
//
// Downloading, reproducing, distributing or otherwise using the SDK Software
// is subject to the terms and conditions of the Boston Dynamics Software
// Development Kit License (20191101-BDSDK-SL).

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: bosdyn/api/log_annotation.proto

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

type LogAnnotationTextMessage_Level int32

const (
	// Invalid, do not use.
	LogAnnotationTextMessage_LEVEL_UNKNOWN LogAnnotationTextMessage_Level = 0
	// Events likely of interest only in a debugging context.
	LogAnnotationTextMessage_LEVEL_DEBUG LogAnnotationTextMessage_Level = 1
	// Informational message during normal operation.
	LogAnnotationTextMessage_LEVEL_INFO LogAnnotationTextMessage_Level = 2
	// Information about an unexpected but recoverable condition.
	LogAnnotationTextMessage_LEVEL_WARN LogAnnotationTextMessage_Level = 3
	// Information about an operation which did not succeed.
	LogAnnotationTextMessage_LEVEL_ERROR LogAnnotationTextMessage_Level = 4
)

// Enum value maps for LogAnnotationTextMessage_Level.
var (
	LogAnnotationTextMessage_Level_name = map[int32]string{
		0: "LEVEL_UNKNOWN",
		1: "LEVEL_DEBUG",
		2: "LEVEL_INFO",
		3: "LEVEL_WARN",
		4: "LEVEL_ERROR",
	}
	LogAnnotationTextMessage_Level_value = map[string]int32{
		"LEVEL_UNKNOWN": 0,
		"LEVEL_DEBUG":   1,
		"LEVEL_INFO":    2,
		"LEVEL_WARN":    3,
		"LEVEL_ERROR":   4,
	}
)

func (x LogAnnotationTextMessage_Level) Enum() *LogAnnotationTextMessage_Level {
	p := new(LogAnnotationTextMessage_Level)
	*p = x
	return p
}

func (x LogAnnotationTextMessage_Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogAnnotationTextMessage_Level) Descriptor() protoreflect.EnumDescriptor {
	return file_bosdyn_api_log_annotation_proto_enumTypes[0].Descriptor()
}

func (LogAnnotationTextMessage_Level) Type() protoreflect.EnumType {
	return &file_bosdyn_api_log_annotation_proto_enumTypes[0]
}

func (x LogAnnotationTextMessage_Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogAnnotationTextMessage_Level.Descriptor instead.
func (LogAnnotationTextMessage_Level) EnumDescriptor() ([]byte, []int) {
	return file_bosdyn_api_log_annotation_proto_rawDescGZIP(), []int{2, 0}
}

// DEPRECATED as of 2.1.0: Please use the DataBufferService instead of the LogAnnotationService.
// The AddLogAnnotation request sends the information that should be added into the log.
type AddLogAnnotationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common request/response header.
	Header *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The annotations to be aded into the log (can be text messages, blobs or robot operator messages).
	Annotations *LogAnnotations `protobuf:"bytes,2,opt,name=annotations,proto3" json:"annotations,omitempty"`
}

func (x *AddLogAnnotationRequest) Reset() {
	*x = AddLogAnnotationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_log_annotation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddLogAnnotationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddLogAnnotationRequest) ProtoMessage() {}

func (x *AddLogAnnotationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_log_annotation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddLogAnnotationRequest.ProtoReflect.Descriptor instead.
func (*AddLogAnnotationRequest) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_log_annotation_proto_rawDescGZIP(), []int{0}
}

func (x *AddLogAnnotationRequest) GetHeader() *RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *AddLogAnnotationRequest) GetAnnotations() *LogAnnotations {
	if x != nil {
		return x.Annotations
	}
	return nil
}

// DEPRECATED as of 2.1.0: Please use the DataBufferService instead of the LogAnnotationService.
// A container for elements to be added to the robot's logs.
type LogAnnotations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Text messages to be added to the log.
	TextMessages []*LogAnnotationTextMessage `protobuf:"bytes,1,rep,name=text_messages,json=textMessages,proto3" json:"text_messages,omitempty"`
	// Messages from the robot operator to be added to the log.
	OperatorMessages []*LogAnnotationOperatorMessage `protobuf:"bytes,2,rep,name=operator_messages,json=operatorMessages,proto3" json:"operator_messages,omitempty"`
	// One or more binary blobs to add to the log.
	BlobData []*LogAnnotationLogBlob `protobuf:"bytes,3,rep,name=blob_data,json=blobData,proto3" json:"blob_data,omitempty"`
}

func (x *LogAnnotations) Reset() {
	*x = LogAnnotations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_log_annotation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogAnnotations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogAnnotations) ProtoMessage() {}

func (x *LogAnnotations) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_log_annotation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogAnnotations.ProtoReflect.Descriptor instead.
func (*LogAnnotations) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_log_annotation_proto_rawDescGZIP(), []int{1}
}

func (x *LogAnnotations) GetTextMessages() []*LogAnnotationTextMessage {
	if x != nil {
		return x.TextMessages
	}
	return nil
}

func (x *LogAnnotations) GetOperatorMessages() []*LogAnnotationOperatorMessage {
	if x != nil {
		return x.OperatorMessages
	}
	return nil
}

func (x *LogAnnotations) GetBlobData() []*LogAnnotationLogBlob {
	if x != nil {
		return x.BlobData
	}
	return nil
}

// DEPRECATED as of 2.1.0: Please use the DataBufferService instead of the LogAnnotationService.
// A text message to add to the robot's logs.
// These could be internal text-log messages from a client for use in debugging, for
// example.
type LogAnnotationTextMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// String annotation message to add to the log.
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// Required timestamp of data in robot clock time.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The service responsible for the annotation. May be omitted.
	Service string `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`
	// Level of significance of the text message.
	Level LogAnnotationTextMessage_Level `protobuf:"varint,4,opt,name=level,proto3,enum=bosdyn.api.LogAnnotationTextMessage_Level" json:"level,omitempty"`
	// Optional tag to identify from what code/module this message originated from.
	Tag string `protobuf:"bytes,5,opt,name=tag,proto3" json:"tag,omitempty"`
	// Optional source file name originating the log message.
	Filename string `protobuf:"bytes,6,opt,name=filename,proto3" json:"filename,omitempty"`
	// Optional source file line number originating the log message.
	LineNumber int32 `protobuf:"varint,7,opt,name=line_number,json=lineNumber,proto3" json:"line_number,omitempty"`
	// Optional timestamp of data in client clock time.
	TimestampClient *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=timestamp_client,json=timestampClient,proto3" json:"timestamp_client,omitempty"`
}

func (x *LogAnnotationTextMessage) Reset() {
	*x = LogAnnotationTextMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_log_annotation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogAnnotationTextMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogAnnotationTextMessage) ProtoMessage() {}

func (x *LogAnnotationTextMessage) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_log_annotation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogAnnotationTextMessage.ProtoReflect.Descriptor instead.
func (*LogAnnotationTextMessage) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_log_annotation_proto_rawDescGZIP(), []int{2}
}

func (x *LogAnnotationTextMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LogAnnotationTextMessage) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *LogAnnotationTextMessage) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *LogAnnotationTextMessage) GetLevel() LogAnnotationTextMessage_Level {
	if x != nil {
		return x.Level
	}
	return LogAnnotationTextMessage_LEVEL_UNKNOWN
}

func (x *LogAnnotationTextMessage) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *LogAnnotationTextMessage) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *LogAnnotationTextMessage) GetLineNumber() int32 {
	if x != nil {
		return x.LineNumber
	}
	return 0
}

func (x *LogAnnotationTextMessage) GetTimestampClient() *timestamppb.Timestamp {
	if x != nil {
		return x.TimestampClient
	}
	return nil
}

// DEPRECATED as of 2.1.0: Please use the DataBufferService instead of the LogAnnotationService.
// An operator message to be added to the robot's logs.
// These are notes especially intended to mark when logs should be preserved and reviewed
// to ensure that robot hardware and/or software is working as intended.
type LogAnnotationOperatorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// String annotation message to add to the log.
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// Required timestamp of data in robot clock time.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Optional timestamp of data in client clock time.
	TimestampClient *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp_client,json=timestampClient,proto3" json:"timestamp_client,omitempty"`
}

func (x *LogAnnotationOperatorMessage) Reset() {
	*x = LogAnnotationOperatorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_log_annotation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogAnnotationOperatorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogAnnotationOperatorMessage) ProtoMessage() {}

func (x *LogAnnotationOperatorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_log_annotation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogAnnotationOperatorMessage.ProtoReflect.Descriptor instead.
func (*LogAnnotationOperatorMessage) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_log_annotation_proto_rawDescGZIP(), []int{3}
}

func (x *LogAnnotationOperatorMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LogAnnotationOperatorMessage) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *LogAnnotationOperatorMessage) GetTimestampClient() *timestamppb.Timestamp {
	if x != nil {
		return x.TimestampClient
	}
	return nil
}

// DEPRECATED as of 2.1.0: Please use the DataBufferService instead of the LogAnnotationService.
// A unit of binary data to be entered in a log.
type LogAnnotationLogBlob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required timestamp of data in robot clock time.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// A general label for this blob.
	// This is distinct from type_id, which identifies how the blob is to be parsed.
	Channel string `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	// A description of the data's content and its encoding.
	// This should be sufficient for deciding how to deserialize the data.
	// For example, this could be the full name of a protobuf message type.
	TypeId string `protobuf:"bytes,3,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	// Raw data to be included as the blob log.
	Data []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	// Optional timestamp of data in client clock time.
	TimestampClient *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp_client,json=timestampClient,proto3" json:"timestamp_client,omitempty"`
}

func (x *LogAnnotationLogBlob) Reset() {
	*x = LogAnnotationLogBlob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_log_annotation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogAnnotationLogBlob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogAnnotationLogBlob) ProtoMessage() {}

func (x *LogAnnotationLogBlob) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_log_annotation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogAnnotationLogBlob.ProtoReflect.Descriptor instead.
func (*LogAnnotationLogBlob) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_log_annotation_proto_rawDescGZIP(), []int{4}
}

func (x *LogAnnotationLogBlob) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *LogAnnotationLogBlob) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *LogAnnotationLogBlob) GetTypeId() string {
	if x != nil {
		return x.TypeId
	}
	return ""
}

func (x *LogAnnotationLogBlob) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *LogAnnotationLogBlob) GetTimestampClient() *timestamppb.Timestamp {
	if x != nil {
		return x.TimestampClient
	}
	return nil
}

// DEPRECATED as of 2.1.0: Please use the DataBufferService instead of the LogAnnotationService.
// The AddLogAnnotation response message, which is empty except for any potential header errors/warnings.
type AddLogAnnotationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common request/response header.
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *AddLogAnnotationResponse) Reset() {
	*x = AddLogAnnotationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_log_annotation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddLogAnnotationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddLogAnnotationResponse) ProtoMessage() {}

func (x *AddLogAnnotationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_log_annotation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddLogAnnotationResponse.ProtoReflect.Descriptor instead.
func (*AddLogAnnotationResponse) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_log_annotation_proto_rawDescGZIP(), []int{5}
}

func (x *AddLogAnnotationResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

var File_bosdyn_api_log_annotation_proto protoreflect.FileDescriptor

var file_bosdyn_api_log_annotation_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x67,
	0x5f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x17, 0x62,
	0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x4c,
	0x6f, 0x67, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6f,
	0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0xf1, 0x01, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x49, 0x0a, 0x0d, 0x74, 0x65, 0x78, 0x74, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67, 0x41,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x0c, 0x74, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x55, 0x0a, 0x11, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x09, 0x62, 0x6c, 0x6f,
	0x62, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62,
	0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x08,
	0x62, 0x6c, 0x6f, 0x62, 0x44, 0x61, 0x74, 0x61, 0x22, 0xbe, 0x03, 0x0a, 0x18, 0x4c, 0x6f, 0x67,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4c, 0x6f, 0x67, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x5c, 0x0a, 0x05, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x45, 0x56, 0x45, 0x4c,
	0x5f, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x45, 0x56, 0x45,
	0x4c, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x45, 0x56, 0x45,
	0x4c, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x45, 0x56, 0x45,
	0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x22, 0xb9, 0x01, 0x0a, 0x1c, 0x4c, 0x6f,
	0x67, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x45,
	0x0a, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0xde, 0x01, 0x0a, 0x14, 0x4c, 0x6f, 0x67, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x38,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x45, 0x0a, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x54, 0x0a, 0x18, 0x41, 0x64, 0x64, 0x4c, 0x6f, 0x67,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x42, 0x47, 0x42, 0x12,
	0x4c, 0x6f, 0x67, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x68, 0x73, 0x73, 0x6e, 0x76, 0x2f, 0x73, 0x70, 0x6f, 0x74, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70,
	0x72, 0x65, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x6f, 0x73, 0x64, 0x79,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bosdyn_api_log_annotation_proto_rawDescOnce sync.Once
	file_bosdyn_api_log_annotation_proto_rawDescData = file_bosdyn_api_log_annotation_proto_rawDesc
)

func file_bosdyn_api_log_annotation_proto_rawDescGZIP() []byte {
	file_bosdyn_api_log_annotation_proto_rawDescOnce.Do(func() {
		file_bosdyn_api_log_annotation_proto_rawDescData = protoimpl.X.CompressGZIP(file_bosdyn_api_log_annotation_proto_rawDescData)
	})
	return file_bosdyn_api_log_annotation_proto_rawDescData
}

var file_bosdyn_api_log_annotation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bosdyn_api_log_annotation_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_bosdyn_api_log_annotation_proto_goTypes = []interface{}{
	(LogAnnotationTextMessage_Level)(0),  // 0: bosdyn.api.LogAnnotationTextMessage.Level
	(*AddLogAnnotationRequest)(nil),      // 1: bosdyn.api.AddLogAnnotationRequest
	(*LogAnnotations)(nil),               // 2: bosdyn.api.LogAnnotations
	(*LogAnnotationTextMessage)(nil),     // 3: bosdyn.api.LogAnnotationTextMessage
	(*LogAnnotationOperatorMessage)(nil), // 4: bosdyn.api.LogAnnotationOperatorMessage
	(*LogAnnotationLogBlob)(nil),         // 5: bosdyn.api.LogAnnotationLogBlob
	(*AddLogAnnotationResponse)(nil),     // 6: bosdyn.api.AddLogAnnotationResponse
	(*RequestHeader)(nil),                // 7: bosdyn.api.RequestHeader
	(*timestamppb.Timestamp)(nil),        // 8: google.protobuf.Timestamp
	(*ResponseHeader)(nil),               // 9: bosdyn.api.ResponseHeader
}
var file_bosdyn_api_log_annotation_proto_depIdxs = []int32{
	7,  // 0: bosdyn.api.AddLogAnnotationRequest.header:type_name -> bosdyn.api.RequestHeader
	2,  // 1: bosdyn.api.AddLogAnnotationRequest.annotations:type_name -> bosdyn.api.LogAnnotations
	3,  // 2: bosdyn.api.LogAnnotations.text_messages:type_name -> bosdyn.api.LogAnnotationTextMessage
	4,  // 3: bosdyn.api.LogAnnotations.operator_messages:type_name -> bosdyn.api.LogAnnotationOperatorMessage
	5,  // 4: bosdyn.api.LogAnnotations.blob_data:type_name -> bosdyn.api.LogAnnotationLogBlob
	8,  // 5: bosdyn.api.LogAnnotationTextMessage.timestamp:type_name -> google.protobuf.Timestamp
	0,  // 6: bosdyn.api.LogAnnotationTextMessage.level:type_name -> bosdyn.api.LogAnnotationTextMessage.Level
	8,  // 7: bosdyn.api.LogAnnotationTextMessage.timestamp_client:type_name -> google.protobuf.Timestamp
	8,  // 8: bosdyn.api.LogAnnotationOperatorMessage.timestamp:type_name -> google.protobuf.Timestamp
	8,  // 9: bosdyn.api.LogAnnotationOperatorMessage.timestamp_client:type_name -> google.protobuf.Timestamp
	8,  // 10: bosdyn.api.LogAnnotationLogBlob.timestamp:type_name -> google.protobuf.Timestamp
	8,  // 11: bosdyn.api.LogAnnotationLogBlob.timestamp_client:type_name -> google.protobuf.Timestamp
	9,  // 12: bosdyn.api.AddLogAnnotationResponse.header:type_name -> bosdyn.api.ResponseHeader
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_bosdyn_api_log_annotation_proto_init() }
func file_bosdyn_api_log_annotation_proto_init() {
	if File_bosdyn_api_log_annotation_proto != nil {
		return
	}
	file_bosdyn_api_header_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bosdyn_api_log_annotation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddLogAnnotationRequest); i {
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
		file_bosdyn_api_log_annotation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogAnnotations); i {
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
		file_bosdyn_api_log_annotation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogAnnotationTextMessage); i {
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
		file_bosdyn_api_log_annotation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogAnnotationOperatorMessage); i {
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
		file_bosdyn_api_log_annotation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogAnnotationLogBlob); i {
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
		file_bosdyn_api_log_annotation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddLogAnnotationResponse); i {
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
			RawDescriptor: file_bosdyn_api_log_annotation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bosdyn_api_log_annotation_proto_goTypes,
		DependencyIndexes: file_bosdyn_api_log_annotation_proto_depIdxs,
		EnumInfos:         file_bosdyn_api_log_annotation_proto_enumTypes,
		MessageInfos:      file_bosdyn_api_log_annotation_proto_msgTypes,
	}.Build()
	File_bosdyn_api_log_annotation_proto = out.File
	file_bosdyn_api_log_annotation_proto_rawDesc = nil
	file_bosdyn_api_log_annotation_proto_goTypes = nil
	file_bosdyn_api_log_annotation_proto_depIdxs = nil
}