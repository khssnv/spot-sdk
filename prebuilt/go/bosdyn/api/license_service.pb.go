// Copyright (c) 2022 Boston Dynamics, Inc.  All rights reserved.
//
// Downloading, reproducing, distributing or otherwise using the SDK Software
// is subject to the terms and conditions of the Boston Dynamics Software
// Development Kit License (20191101-BDSDK-SL).

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: bosdyn/api/license_service.proto

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

var File_bosdyn_api_license_service_proto protoreflect.FileDescriptor

var file_bosdyn_api_license_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x18,
	0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xcf, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x2e,
	0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x24, 0x2e, 0x62, 0x6f,
	0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x48, 0x42, 0x13, 0x4c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x68,
	0x73, 0x73, 0x6e, 0x76, 0x2f, 0x73, 0x70, 0x6f, 0x74, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72,
	0x65, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_bosdyn_api_license_service_proto_goTypes = []interface{}{
	(*GetLicenseInfoRequest)(nil),     // 0: bosdyn.api.GetLicenseInfoRequest
	(*GetFeatureEnabledRequest)(nil),  // 1: bosdyn.api.GetFeatureEnabledRequest
	(*GetLicenseInfoResponse)(nil),    // 2: bosdyn.api.GetLicenseInfoResponse
	(*GetFeatureEnabledResponse)(nil), // 3: bosdyn.api.GetFeatureEnabledResponse
}
var file_bosdyn_api_license_service_proto_depIdxs = []int32{
	0, // 0: bosdyn.api.LicenseService.GetLicenseInfo:input_type -> bosdyn.api.GetLicenseInfoRequest
	1, // 1: bosdyn.api.LicenseService.GetFeatureEnabled:input_type -> bosdyn.api.GetFeatureEnabledRequest
	2, // 2: bosdyn.api.LicenseService.GetLicenseInfo:output_type -> bosdyn.api.GetLicenseInfoResponse
	3, // 3: bosdyn.api.LicenseService.GetFeatureEnabled:output_type -> bosdyn.api.GetFeatureEnabledResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bosdyn_api_license_service_proto_init() }
func file_bosdyn_api_license_service_proto_init() {
	if File_bosdyn_api_license_service_proto != nil {
		return
	}
	file_bosdyn_api_license_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bosdyn_api_license_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bosdyn_api_license_service_proto_goTypes,
		DependencyIndexes: file_bosdyn_api_license_service_proto_depIdxs,
	}.Build()
	File_bosdyn_api_license_service_proto = out.File
	file_bosdyn_api_license_service_proto_rawDesc = nil
	file_bosdyn_api_license_service_proto_goTypes = nil
	file_bosdyn_api_license_service_proto_depIdxs = nil
}
