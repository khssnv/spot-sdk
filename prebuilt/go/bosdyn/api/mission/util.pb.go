// Copyright (c) 2022 Boston Dynamics, Inc.  All rights reserved.
//
// Downloading, reproducing, distributing or otherwise using the SDK Software
// is subject to the terms and conditions of the Boston Dynamics Software
// Development Kit License (20191101-BDSDK-SL).

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: bosdyn/api/mission/util.proto

package mission

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Results from executing / ticking / running a single node.
type Result int32

const (
	// Invalid, should not be used.
	Result_RESULT_UNKNOWN Result = 0
	// The node completed running, but failed.
	Result_RESULT_FAILURE Result = 1
	// The node is still in process and has not completed.
	Result_RESULT_RUNNING Result = 2
	// The node completed, and succeeded.
	Result_RESULT_SUCCESS Result = 3
	// The node encountered an operational error while trying to execute.
	Result_RESULT_ERROR Result = 4
)

// Enum value maps for Result.
var (
	Result_name = map[int32]string{
		0: "RESULT_UNKNOWN",
		1: "RESULT_FAILURE",
		2: "RESULT_RUNNING",
		3: "RESULT_SUCCESS",
		4: "RESULT_ERROR",
	}
	Result_value = map[string]int32{
		"RESULT_UNKNOWN": 0,
		"RESULT_FAILURE": 1,
		"RESULT_RUNNING": 2,
		"RESULT_SUCCESS": 3,
		"RESULT_ERROR":   4,
	}
)

func (x Result) Enum() *Result {
	p := new(Result)
	*p = x
	return p
}

func (x Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Result) Descriptor() protoreflect.EnumDescriptor {
	return file_bosdyn_api_mission_util_proto_enumTypes[0].Descriptor()
}

func (Result) Type() protoreflect.EnumType {
	return &file_bosdyn_api_mission_util_proto_enumTypes[0]
}

func (x Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Result.Descriptor instead.
func (Result) EnumDescriptor() ([]byte, []int) {
	return file_bosdyn_api_mission_util_proto_rawDescGZIP(), []int{0}
}

// Supported types for blackboard or parameter values.
type VariableDeclaration_Type int32

const (
	VariableDeclaration_TYPE_UNKNOWN VariableDeclaration_Type = 0
	VariableDeclaration_TYPE_FLOAT   VariableDeclaration_Type = 1
	VariableDeclaration_TYPE_STRING  VariableDeclaration_Type = 2
	VariableDeclaration_TYPE_INT     VariableDeclaration_Type = 3
	VariableDeclaration_TYPE_BOOL    VariableDeclaration_Type = 4
	VariableDeclaration_TYPE_MESSAGE VariableDeclaration_Type = 5
)

// Enum value maps for VariableDeclaration_Type.
var (
	VariableDeclaration_Type_name = map[int32]string{
		0: "TYPE_UNKNOWN",
		1: "TYPE_FLOAT",
		2: "TYPE_STRING",
		3: "TYPE_INT",
		4: "TYPE_BOOL",
		5: "TYPE_MESSAGE",
	}
	VariableDeclaration_Type_value = map[string]int32{
		"TYPE_UNKNOWN": 0,
		"TYPE_FLOAT":   1,
		"TYPE_STRING":  2,
		"TYPE_INT":     3,
		"TYPE_BOOL":    4,
		"TYPE_MESSAGE": 5,
	}
)

func (x VariableDeclaration_Type) Enum() *VariableDeclaration_Type {
	p := new(VariableDeclaration_Type)
	*p = x
	return p
}

func (x VariableDeclaration_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VariableDeclaration_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_bosdyn_api_mission_util_proto_enumTypes[1].Descriptor()
}

func (VariableDeclaration_Type) Type() protoreflect.EnumType {
	return &file_bosdyn_api_mission_util_proto_enumTypes[1]
}

func (x VariableDeclaration_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VariableDeclaration_Type.Descriptor instead.
func (VariableDeclaration_Type) EnumDescriptor() ([]byte, []int) {
	return file_bosdyn_api_mission_util_proto_rawDescGZIP(), []int{2, 0}
}

// Key/Value pair, used in other messages.
type KeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value *Value `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_mission_util_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_mission_util_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValue.ProtoReflect.Descriptor instead.
func (*KeyValue) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_mission_util_proto_rawDescGZIP(), []int{0}
}

func (x *KeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValue) GetValue() *Value {
	if x != nil {
		return x.Value
	}
	return nil
}

// A value of a run-time or compile-time variable.
type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Source:
	//	*Value_Constant
	//	*Value_RuntimeVar
	//	*Value_Parameter
	Source isValue_Source `protobuf_oneof:"source"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_mission_util_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_mission_util_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_mission_util_proto_rawDescGZIP(), []int{1}
}

func (m *Value) GetSource() isValue_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *Value) GetConstant() *ConstantValue {
	if x, ok := x.GetSource().(*Value_Constant); ok {
		return x.Constant
	}
	return nil
}

func (x *Value) GetRuntimeVar() *VariableDeclaration {
	if x, ok := x.GetSource().(*Value_RuntimeVar); ok {
		return x.RuntimeVar
	}
	return nil
}

func (x *Value) GetParameter() *VariableDeclaration {
	if x, ok := x.GetSource().(*Value_Parameter); ok {
		return x.Parameter
	}
	return nil
}

type isValue_Source interface {
	isValue_Source()
}

type Value_Constant struct {
	Constant *ConstantValue `protobuf:"bytes,2,opt,name=constant,proto3,oneof"` // A constant value.
}

type Value_RuntimeVar struct {
	RuntimeVar *VariableDeclaration `protobuf:"bytes,3,opt,name=runtime_var,json=runtimeVar,proto3,oneof"` // Look up a variable provided at run-time.
}

type Value_Parameter struct {
	Parameter *VariableDeclaration `protobuf:"bytes,4,opt,name=parameter,proto3,oneof"` // Look up a Node Parameter.
}

func (*Value_Constant) isValue_Source() {}

func (*Value_RuntimeVar) isValue_Source() {}

func (*Value_Parameter) isValue_Source() {}

// Declaration of a run-time or compile-time variable.
type VariableDeclaration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the variable, to be used as the key in KeyValue pairs.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Type that this variable is expected to have. Used to verify assignments and comparisons.
	Type VariableDeclaration_Type `protobuf:"varint,2,opt,name=type,proto3,enum=bosdyn.api.mission.VariableDeclaration_Type" json:"type,omitempty"`
}

func (x *VariableDeclaration) Reset() {
	*x = VariableDeclaration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_mission_util_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VariableDeclaration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariableDeclaration) ProtoMessage() {}

func (x *VariableDeclaration) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_mission_util_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VariableDeclaration.ProtoReflect.Descriptor instead.
func (*VariableDeclaration) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_mission_util_proto_rawDescGZIP(), []int{2}
}

func (x *VariableDeclaration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VariableDeclaration) GetType() VariableDeclaration_Type {
	if x != nil {
		return x.Type
	}
	return VariableDeclaration_TYPE_UNKNOWN
}

// A constant value. Corresponds to the VariableDeclaration Type enum.
type ConstantValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*ConstantValue_FloatValue
	//	*ConstantValue_StringValue
	//	*ConstantValue_IntValue
	//	*ConstantValue_BoolValue
	//	*ConstantValue_MsgValue
	Value isConstantValue_Value `protobuf_oneof:"value"`
}

func (x *ConstantValue) Reset() {
	*x = ConstantValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_mission_util_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConstantValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConstantValue) ProtoMessage() {}

func (x *ConstantValue) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_mission_util_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConstantValue.ProtoReflect.Descriptor instead.
func (*ConstantValue) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_mission_util_proto_rawDescGZIP(), []int{3}
}

func (m *ConstantValue) GetValue() isConstantValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *ConstantValue) GetFloatValue() float64 {
	if x, ok := x.GetValue().(*ConstantValue_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

func (x *ConstantValue) GetStringValue() string {
	if x, ok := x.GetValue().(*ConstantValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *ConstantValue) GetIntValue() int64 {
	if x, ok := x.GetValue().(*ConstantValue_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (x *ConstantValue) GetBoolValue() bool {
	if x, ok := x.GetValue().(*ConstantValue_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (x *ConstantValue) GetMsgValue() *anypb.Any {
	if x, ok := x.GetValue().(*ConstantValue_MsgValue); ok {
		return x.MsgValue
	}
	return nil
}

type isConstantValue_Value interface {
	isConstantValue_Value()
}

type ConstantValue_FloatValue struct {
	FloatValue float64 `protobuf:"fixed64,1,opt,name=float_value,json=floatValue,proto3,oneof"`
}

type ConstantValue_StringValue struct {
	StringValue string `protobuf:"bytes,2,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type ConstantValue_IntValue struct {
	IntValue int64 `protobuf:"varint,3,opt,name=int_value,json=intValue,proto3,oneof"`
}

type ConstantValue_BoolValue struct {
	BoolValue bool `protobuf:"varint,4,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type ConstantValue_MsgValue struct {
	MsgValue *anypb.Any `protobuf:"bytes,5,opt,name=msg_value,json=msgValue,proto3,oneof"`
}

func (*ConstantValue_FloatValue) isConstantValue_Value() {}

func (*ConstantValue_StringValue) isConstantValue_Value() {}

func (*ConstantValue_IntValue) isConstantValue_Value() {}

func (*ConstantValue_BoolValue) isConstantValue_Value() {}

func (*ConstantValue_MsgValue) isConstantValue_Value() {}

// Data a user can associate with a node.
type UserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. Enables matching the Node uploaded to the MissionService with the NodeInfo
	// downloaded from the MissionService.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Arbitrary data. We recommend keeping it small, to avoid bloating the size of the mission.
	Bytestring []byte `protobuf:"bytes,3,opt,name=bytestring,proto3" json:"bytestring,omitempty"`
}

func (x *UserData) Reset() {
	*x = UserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_mission_util_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_mission_util_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_mission_util_proto_rawDescGZIP(), []int{4}
}

func (x *UserData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserData) GetBytestring() []byte {
	if x != nil {
		return x.Bytestring
	}
	return nil
}

var File_bosdyn_api_mission_util_proto protoreflect.FileDescriptor

var file_bosdyn_api_mission_util_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d,
	0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x6f,
	0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xe7, 0x01,
	0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x6f, 0x73, 0x64,
	0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x08,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x12, 0x4a, 0x0a, 0x0b, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x76, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x65, 0x63, 0x6c, 0x61,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0a, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x56, 0x61, 0x72, 0x12, 0x47, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x61, 0x72,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x42, 0x08, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0xd5, 0x01, 0x0a, 0x13, 0x56, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x44, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2c, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x44,
	0x65, 0x63, 0x6c, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x68, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a,
	0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10, 0x01, 0x12,
	0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x02,
	0x12, 0x0c, 0x0a, 0x08, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x0d,
	0x0a, 0x09, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x04, 0x12, 0x10, 0x0a,
	0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x05, 0x22,
	0xd5, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x21, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6e, 0x74,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08,
	0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09,
	0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x6d, 0x73, 0x67,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x73, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x2a, 0x6a, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a,
	0x0e, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x55, 0x52, 0x45, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f,
	0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x53,
	0x55, 0x4c, 0x54, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x03, 0x12, 0x10, 0x0a,
	0x0c, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x42,
	0x46, 0x42, 0x09, 0x55, 0x74, 0x69, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x39, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x68, 0x73, 0x73, 0x6e, 0x76, 0x2f,
	0x73, 0x70, 0x6f, 0x74, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x65, 0x62, 0x75, 0x69, 0x6c,
	0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bosdyn_api_mission_util_proto_rawDescOnce sync.Once
	file_bosdyn_api_mission_util_proto_rawDescData = file_bosdyn_api_mission_util_proto_rawDesc
)

func file_bosdyn_api_mission_util_proto_rawDescGZIP() []byte {
	file_bosdyn_api_mission_util_proto_rawDescOnce.Do(func() {
		file_bosdyn_api_mission_util_proto_rawDescData = protoimpl.X.CompressGZIP(file_bosdyn_api_mission_util_proto_rawDescData)
	})
	return file_bosdyn_api_mission_util_proto_rawDescData
}

var file_bosdyn_api_mission_util_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_bosdyn_api_mission_util_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_bosdyn_api_mission_util_proto_goTypes = []interface{}{
	(Result)(0),                   // 0: bosdyn.api.mission.Result
	(VariableDeclaration_Type)(0), // 1: bosdyn.api.mission.VariableDeclaration.Type
	(*KeyValue)(nil),              // 2: bosdyn.api.mission.KeyValue
	(*Value)(nil),                 // 3: bosdyn.api.mission.Value
	(*VariableDeclaration)(nil),   // 4: bosdyn.api.mission.VariableDeclaration
	(*ConstantValue)(nil),         // 5: bosdyn.api.mission.ConstantValue
	(*UserData)(nil),              // 6: bosdyn.api.mission.UserData
	(*anypb.Any)(nil),             // 7: google.protobuf.Any
}
var file_bosdyn_api_mission_util_proto_depIdxs = []int32{
	3, // 0: bosdyn.api.mission.KeyValue.value:type_name -> bosdyn.api.mission.Value
	5, // 1: bosdyn.api.mission.Value.constant:type_name -> bosdyn.api.mission.ConstantValue
	4, // 2: bosdyn.api.mission.Value.runtime_var:type_name -> bosdyn.api.mission.VariableDeclaration
	4, // 3: bosdyn.api.mission.Value.parameter:type_name -> bosdyn.api.mission.VariableDeclaration
	1, // 4: bosdyn.api.mission.VariableDeclaration.type:type_name -> bosdyn.api.mission.VariableDeclaration.Type
	7, // 5: bosdyn.api.mission.ConstantValue.msg_value:type_name -> google.protobuf.Any
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_bosdyn_api_mission_util_proto_init() }
func file_bosdyn_api_mission_util_proto_init() {
	if File_bosdyn_api_mission_util_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bosdyn_api_mission_util_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValue); i {
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
		file_bosdyn_api_mission_util_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_bosdyn_api_mission_util_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VariableDeclaration); i {
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
		file_bosdyn_api_mission_util_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConstantValue); i {
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
		file_bosdyn_api_mission_util_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserData); i {
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
	file_bosdyn_api_mission_util_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Value_Constant)(nil),
		(*Value_RuntimeVar)(nil),
		(*Value_Parameter)(nil),
	}
	file_bosdyn_api_mission_util_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ConstantValue_FloatValue)(nil),
		(*ConstantValue_StringValue)(nil),
		(*ConstantValue_IntValue)(nil),
		(*ConstantValue_BoolValue)(nil),
		(*ConstantValue_MsgValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bosdyn_api_mission_util_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bosdyn_api_mission_util_proto_goTypes,
		DependencyIndexes: file_bosdyn_api_mission_util_proto_depIdxs,
		EnumInfos:         file_bosdyn_api_mission_util_proto_enumTypes,
		MessageInfos:      file_bosdyn_api_mission_util_proto_msgTypes,
	}.Build()
	File_bosdyn_api_mission_util_proto = out.File
	file_bosdyn_api_mission_util_proto_rawDesc = nil
	file_bosdyn_api_mission_util_proto_goTypes = nil
	file_bosdyn_api_mission_util_proto_depIdxs = nil
}
