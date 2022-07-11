// Copyright (c) 2022 Boston Dynamics, Inc.  All rights reserved.
//
// Downloading, reproducing, distributing or otherwise using the SDK Software
// is subject to the terms and conditions of the Boston Dynamics Software
// Development Kit License (20191101-BDSDK-SL).

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: bosdyn/api/network_stats.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type WifiDevice_Type int32

const (
	WifiDevice_UNKNOWN WifiDevice_Type = 0
	WifiDevice_AP      WifiDevice_Type = 1
	WifiDevice_CLIENT  WifiDevice_Type = 2
)

// Enum value maps for WifiDevice_Type.
var (
	WifiDevice_Type_name = map[int32]string{
		0: "UNKNOWN",
		1: "AP",
		2: "CLIENT",
	}
	WifiDevice_Type_value = map[string]int32{
		"UNKNOWN": 0,
		"AP":      1,
		"CLIENT":  2,
	}
)

func (x WifiDevice_Type) Enum() *WifiDevice_Type {
	p := new(WifiDevice_Type)
	*p = x
	return p
}

func (x WifiDevice_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WifiDevice_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_bosdyn_api_network_stats_proto_enumTypes[0].Descriptor()
}

func (WifiDevice_Type) Type() protoreflect.EnumType {
	return &file_bosdyn_api_network_stats_proto_enumTypes[0]
}

func (x WifiDevice_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WifiDevice_Type.Descriptor instead.
func (WifiDevice_Type) EnumDescriptor() ([]byte, []int) {
	return file_bosdyn_api_network_stats_proto_rawDescGZIP(), []int{1, 0}
}

type Association struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// MAC address of the associated station
	MacAddress string `protobuf:"bytes,1,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	// Time duration since the station last connected.
	ConnectedTime *durationpb.Duration `protobuf:"bytes,2,opt,name=connected_time,json=connectedTime,proto3" json:"connected_time,omitempty"`
	// Signal strength of last received packet
	RxSignalDbm int32 `protobuf:"varint,3,opt,name=rx_signal_dbm,json=rxSignalDbm,proto3" json:"rx_signal_dbm,omitempty"`
	// Signal strength average
	RxSignalAvgDbm int32 `protobuf:"varint,4,opt,name=rx_signal_avg_dbm,json=rxSignalAvgDbm,proto3" json:"rx_signal_avg_dbm,omitempty"`
	// Signal strength average for beacons only.
	RxBeaconSignalAvgDbm int32 `protobuf:"varint,5,opt,name=rx_beacon_signal_avg_dbm,json=rxBeaconSignalAvgDbm,proto3" json:"rx_beacon_signal_avg_dbm,omitempty"`
	// Expected throughput
	ExpectedBitsPerSecond int64 `protobuf:"varint,6,opt,name=expected_bits_per_second,json=expectedBitsPerSecond,proto3" json:"expected_bits_per_second,omitempty"`
	// Total received bytes
	RxBytes int64 `protobuf:"varint,7,opt,name=rx_bytes,json=rxBytes,proto3" json:"rx_bytes,omitempty"`
	// Total received packets from the associated station
	RxPackets int64 `protobuf:"varint,8,opt,name=rx_packets,json=rxPackets,proto3" json:"rx_packets,omitempty"`
	// Last unicast receive rate
	RxBitsPerSecond int64 `protobuf:"varint,9,opt,name=rx_bits_per_second,json=rxBitsPerSecond,proto3" json:"rx_bits_per_second,omitempty"`
	// Total transmitted bytes
	TxBytes int64 `protobuf:"varint,10,opt,name=tx_bytes,json=txBytes,proto3" json:"tx_bytes,omitempty"`
	// Total transmitted packets to the associated station
	TxPackets int64 `protobuf:"varint,11,opt,name=tx_packets,json=txPackets,proto3" json:"tx_packets,omitempty"`
	// Current unicast transmit rate
	TxBitsPerSecond int64 `protobuf:"varint,12,opt,name=tx_bits_per_second,json=txBitsPerSecond,proto3" json:"tx_bits_per_second,omitempty"`
	// Cumulative retry count to this station, within connected time
	TxRetries int64 `protobuf:"varint,13,opt,name=tx_retries,json=txRetries,proto3" json:"tx_retries,omitempty"`
	// Cumulative failed tx packet count to this station, within connected time
	TxFailed int64 `protobuf:"varint,14,opt,name=tx_failed,json=txFailed,proto3" json:"tx_failed,omitempty"`
	// Number of beacons received from this peer
	BeaconsReceived int64 `protobuf:"varint,15,opt,name=beacons_received,json=beaconsReceived,proto3" json:"beacons_received,omitempty"`
	// Number of times beacon loss was detected
	BeaconLossCount int64 `protobuf:"varint,16,opt,name=beacon_loss_count,json=beaconLossCount,proto3" json:"beacon_loss_count,omitempty"`
}

func (x *Association) Reset() {
	*x = Association{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_network_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Association) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Association) ProtoMessage() {}

func (x *Association) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_network_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Association.ProtoReflect.Descriptor instead.
func (*Association) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_network_stats_proto_rawDescGZIP(), []int{0}
}

func (x *Association) GetMacAddress() string {
	if x != nil {
		return x.MacAddress
	}
	return ""
}

func (x *Association) GetConnectedTime() *durationpb.Duration {
	if x != nil {
		return x.ConnectedTime
	}
	return nil
}

func (x *Association) GetRxSignalDbm() int32 {
	if x != nil {
		return x.RxSignalDbm
	}
	return 0
}

func (x *Association) GetRxSignalAvgDbm() int32 {
	if x != nil {
		return x.RxSignalAvgDbm
	}
	return 0
}

func (x *Association) GetRxBeaconSignalAvgDbm() int32 {
	if x != nil {
		return x.RxBeaconSignalAvgDbm
	}
	return 0
}

func (x *Association) GetExpectedBitsPerSecond() int64 {
	if x != nil {
		return x.ExpectedBitsPerSecond
	}
	return 0
}

func (x *Association) GetRxBytes() int64 {
	if x != nil {
		return x.RxBytes
	}
	return 0
}

func (x *Association) GetRxPackets() int64 {
	if x != nil {
		return x.RxPackets
	}
	return 0
}

func (x *Association) GetRxBitsPerSecond() int64 {
	if x != nil {
		return x.RxBitsPerSecond
	}
	return 0
}

func (x *Association) GetTxBytes() int64 {
	if x != nil {
		return x.TxBytes
	}
	return 0
}

func (x *Association) GetTxPackets() int64 {
	if x != nil {
		return x.TxPackets
	}
	return 0
}

func (x *Association) GetTxBitsPerSecond() int64 {
	if x != nil {
		return x.TxBitsPerSecond
	}
	return 0
}

func (x *Association) GetTxRetries() int64 {
	if x != nil {
		return x.TxRetries
	}
	return 0
}

func (x *Association) GetTxFailed() int64 {
	if x != nil {
		return x.TxFailed
	}
	return 0
}

func (x *Association) GetBeaconsReceived() int64 {
	if x != nil {
		return x.BeaconsReceived
	}
	return 0
}

func (x *Association) GetBeaconLossCount() int64 {
	if x != nil {
		return x.BeaconLossCount
	}
	return 0
}

type WifiDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         WifiDevice_Type `protobuf:"varint,1,opt,name=type,proto3,enum=bosdyn.api.WifiDevice_Type" json:"type,omitempty"`
	Name         string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MacAddress   string          `protobuf:"bytes,3,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	Ssid         string          `protobuf:"bytes,4,opt,name=ssid,proto3" json:"ssid,omitempty"`
	TxPowerDbm   int32           `protobuf:"varint,5,opt,name=tx_power_dbm,json=txPowerDbm,proto3" json:"tx_power_dbm,omitempty"`
	Associations []*Association  `protobuf:"bytes,6,rep,name=associations,proto3" json:"associations,omitempty"`
}

func (x *WifiDevice) Reset() {
	*x = WifiDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_network_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WifiDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WifiDevice) ProtoMessage() {}

func (x *WifiDevice) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_network_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WifiDevice.ProtoReflect.Descriptor instead.
func (*WifiDevice) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_network_stats_proto_rawDescGZIP(), []int{1}
}

func (x *WifiDevice) GetType() WifiDevice_Type {
	if x != nil {
		return x.Type
	}
	return WifiDevice_UNKNOWN
}

func (x *WifiDevice) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WifiDevice) GetMacAddress() string {
	if x != nil {
		return x.MacAddress
	}
	return ""
}

func (x *WifiDevice) GetSsid() string {
	if x != nil {
		return x.Ssid
	}
	return ""
}

func (x *WifiDevice) GetTxPowerDbm() int32 {
	if x != nil {
		return x.TxPowerDbm
	}
	return 0
}

func (x *WifiDevice) GetAssociations() []*Association {
	if x != nil {
		return x.Associations
	}
	return nil
}

type WifiStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string        `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Devices  []*WifiDevice `protobuf:"bytes,2,rep,name=devices,proto3" json:"devices,omitempty"`
}

func (x *WifiStats) Reset() {
	*x = WifiStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosdyn_api_network_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WifiStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WifiStats) ProtoMessage() {}

func (x *WifiStats) ProtoReflect() protoreflect.Message {
	mi := &file_bosdyn_api_network_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WifiStats.ProtoReflect.Descriptor instead.
func (*WifiStats) Descriptor() ([]byte, []int) {
	return file_bosdyn_api_network_stats_proto_rawDescGZIP(), []int{2}
}

func (x *WifiStats) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *WifiStats) GetDevices() []*WifiDevice {
	if x != nil {
		return x.Devices
	}
	return nil
}

var File_bosdyn_api_network_stats_proto protoreflect.FileDescriptor

var file_bosdyn_api_network_stats_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x05, 0x0a,
	0x0b, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x40, 0x0a,
	0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x22, 0x0a, 0x0d, 0x72, 0x78, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x62, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x78, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x44, 0x62, 0x6d, 0x12, 0x29, 0x0a, 0x11, 0x72, 0x78, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x5f, 0x61, 0x76, 0x67, 0x5f, 0x64, 0x62, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e,
	0x72, 0x78, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x41, 0x76, 0x67, 0x44, 0x62, 0x6d, 0x12, 0x36,
	0x0a, 0x18, 0x72, 0x78, 0x5f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x5f, 0x61, 0x76, 0x67, 0x5f, 0x64, 0x62, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x14, 0x72, 0x78, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x41, 0x76, 0x67, 0x44, 0x62, 0x6d, 0x12, 0x37, 0x0a, 0x18, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x42, 0x69, 0x74, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x72, 0x78, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x72, 0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x78,
	0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x72, 0x78, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x12, 0x72, 0x78, 0x5f,
	0x62, 0x69, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x72, 0x78, 0x42, 0x69, 0x74, 0x73, 0x50, 0x65, 0x72,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x78, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x78, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x78, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x12, 0x2b, 0x0a, 0x12, 0x74, 0x78, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x74, 0x78,
	0x42, 0x69, 0x74, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x78, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x74, 0x78, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x78, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x74, 0x78, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x62, 0x65, 0x61,
	0x63, 0x6f, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x6c,
	0x6f, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x4c, 0x6f, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x8e, 0x02, 0x0a, 0x0a, 0x57, 0x69, 0x66, 0x69, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e,
	0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x69, 0x66, 0x69, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x63, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x73, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x73, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x74, 0x78, 0x5f,
	0x70, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x64, 0x62, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x74, 0x78, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x44, 0x62, 0x6d, 0x12, 0x3b, 0x0a, 0x0c, 0x61,
	0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41,
	0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x61, 0x73, 0x73, 0x6f,
	0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x27, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x06, 0x0a,
	0x02, 0x41, 0x50, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10,
	0x02, 0x22, 0x59, 0x0a, 0x09, 0x57, 0x69, 0x66, 0x69, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x6f,
	0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x69, 0x66, 0x69, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x46, 0x42, 0x11,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x68,
	0x73, 0x73, 0x6e, 0x76, 0x2f, 0x73, 0x70, 0x6f, 0x74, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72,
	0x65, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bosdyn_api_network_stats_proto_rawDescOnce sync.Once
	file_bosdyn_api_network_stats_proto_rawDescData = file_bosdyn_api_network_stats_proto_rawDesc
)

func file_bosdyn_api_network_stats_proto_rawDescGZIP() []byte {
	file_bosdyn_api_network_stats_proto_rawDescOnce.Do(func() {
		file_bosdyn_api_network_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_bosdyn_api_network_stats_proto_rawDescData)
	})
	return file_bosdyn_api_network_stats_proto_rawDescData
}

var file_bosdyn_api_network_stats_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bosdyn_api_network_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bosdyn_api_network_stats_proto_goTypes = []interface{}{
	(WifiDevice_Type)(0),        // 0: bosdyn.api.WifiDevice.Type
	(*Association)(nil),         // 1: bosdyn.api.Association
	(*WifiDevice)(nil),          // 2: bosdyn.api.WifiDevice
	(*WifiStats)(nil),           // 3: bosdyn.api.WifiStats
	(*durationpb.Duration)(nil), // 4: google.protobuf.Duration
}
var file_bosdyn_api_network_stats_proto_depIdxs = []int32{
	4, // 0: bosdyn.api.Association.connected_time:type_name -> google.protobuf.Duration
	0, // 1: bosdyn.api.WifiDevice.type:type_name -> bosdyn.api.WifiDevice.Type
	1, // 2: bosdyn.api.WifiDevice.associations:type_name -> bosdyn.api.Association
	2, // 3: bosdyn.api.WifiStats.devices:type_name -> bosdyn.api.WifiDevice
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_bosdyn_api_network_stats_proto_init() }
func file_bosdyn_api_network_stats_proto_init() {
	if File_bosdyn_api_network_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bosdyn_api_network_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Association); i {
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
		file_bosdyn_api_network_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WifiDevice); i {
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
		file_bosdyn_api_network_stats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WifiStats); i {
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
			RawDescriptor: file_bosdyn_api_network_stats_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bosdyn_api_network_stats_proto_goTypes,
		DependencyIndexes: file_bosdyn_api_network_stats_proto_depIdxs,
		EnumInfos:         file_bosdyn_api_network_stats_proto_enumTypes,
		MessageInfos:      file_bosdyn_api_network_stats_proto_msgTypes,
	}.Build()
	File_bosdyn_api_network_stats_proto = out.File
	file_bosdyn_api_network_stats_proto_rawDesc = nil
	file_bosdyn_api_network_stats_proto_goTypes = nil
	file_bosdyn_api_network_stats_proto_depIdxs = nil
}
