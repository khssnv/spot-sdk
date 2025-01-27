// Copyright (c) 2022 Boston Dynamics, Inc.  All rights reserved.
//
// Downloading, reproducing, distributing or otherwise using the SDK Software
// is subject to the terms and conditions of the Boston Dynamics Software
// Development Kit License (20191101-BDSDK-SL).

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: bosdyn/api/network_compute_bridge_service.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_bosdyn_api_network_compute_bridge_service_proto protoreflect.FileDescriptor

var file_bosdyn_api_network_compute_bridge_service_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x62, 0x72, 0x69,
	0x64, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x27, 0x62,
	0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xdb, 0x01, 0x0a, 0x14, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x12,
	0x59, 0x0a, 0x0e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x12, 0x21, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x13, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x12, 0x26, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x62, 0x6f, 0x73, 0x64,
	0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x32, 0xe1, 0x01, 0x0a, 0x1a, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x12, 0x59, 0x0a, 0x0e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x26, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x68, 0x73, 0x73, 0x6e, 0x76, 0x2f, 0x73, 0x70,
	0x6f, 0x74, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x65, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x2f,
	0x67, 0x6f, 0x2f, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_bosdyn_api_network_compute_bridge_service_proto_goTypes = []interface{}{
	(*NetworkComputeRequest)(nil),       // 0: bosdyn.api.NetworkComputeRequest
	(*ListAvailableModelsRequest)(nil),  // 1: bosdyn.api.ListAvailableModelsRequest
	(*NetworkComputeResponse)(nil),      // 2: bosdyn.api.NetworkComputeResponse
	(*ListAvailableModelsResponse)(nil), // 3: bosdyn.api.ListAvailableModelsResponse
}
var file_bosdyn_api_network_compute_bridge_service_proto_depIdxs = []int32{
	0, // 0: bosdyn.api.NetworkComputeBridge.NetworkCompute:input_type -> bosdyn.api.NetworkComputeRequest
	1, // 1: bosdyn.api.NetworkComputeBridge.ListAvailableModels:input_type -> bosdyn.api.ListAvailableModelsRequest
	0, // 2: bosdyn.api.NetworkComputeBridgeWorker.NetworkCompute:input_type -> bosdyn.api.NetworkComputeRequest
	1, // 3: bosdyn.api.NetworkComputeBridgeWorker.ListAvailableModels:input_type -> bosdyn.api.ListAvailableModelsRequest
	2, // 4: bosdyn.api.NetworkComputeBridge.NetworkCompute:output_type -> bosdyn.api.NetworkComputeResponse
	3, // 5: bosdyn.api.NetworkComputeBridge.ListAvailableModels:output_type -> bosdyn.api.ListAvailableModelsResponse
	2, // 6: bosdyn.api.NetworkComputeBridgeWorker.NetworkCompute:output_type -> bosdyn.api.NetworkComputeResponse
	3, // 7: bosdyn.api.NetworkComputeBridgeWorker.ListAvailableModels:output_type -> bosdyn.api.ListAvailableModelsResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bosdyn_api_network_compute_bridge_service_proto_init() }
func file_bosdyn_api_network_compute_bridge_service_proto_init() {
	if File_bosdyn_api_network_compute_bridge_service_proto != nil {
		return
	}
	file_bosdyn_api_network_compute_bridge_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bosdyn_api_network_compute_bridge_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_bosdyn_api_network_compute_bridge_service_proto_goTypes,
		DependencyIndexes: file_bosdyn_api_network_compute_bridge_service_proto_depIdxs,
	}.Build()
	File_bosdyn_api_network_compute_bridge_service_proto = out.File
	file_bosdyn_api_network_compute_bridge_service_proto_rawDesc = nil
	file_bosdyn_api_network_compute_bridge_service_proto_goTypes = nil
	file_bosdyn_api_network_compute_bridge_service_proto_depIdxs = nil
}
