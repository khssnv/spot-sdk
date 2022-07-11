// Copyright (c) 2022 Boston Dynamics, Inc.  All rights reserved.
//
// Downloading, reproducing, distributing or otherwise using the SDK Software
// is subject to the terms and conditions of the Boston Dynamics Software
// Development Kit License (20191101-BDSDK-SL).

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: bosdyn/api/manipulation_api_service.proto

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

var File_bosdyn_api_manipulation_api_service_proto protoreflect.FileDescriptor

var file_bosdyn_api_manipulation_api_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e,
	0x69, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62, 0x6f, 0x73,
	0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x21, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x69, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xca, 0x02, 0x0a, 0x16, 0x4d,
	0x61, 0x6e, 0x69, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x70, 0x69, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x0f, 0x4d, 0x61, 0x6e, 0x69, 0x70, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x70, 0x69, 0x12, 0x22, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x6e, 0x69, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x62,
	0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x6e, 0x69, 0x70, 0x75,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x74, 0x0a, 0x17, 0x4d, 0x61, 0x6e, 0x69, 0x70, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x70, 0x69, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x2a,
	0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x6e, 0x69,
	0x70, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x70, 0x69, 0x46, 0x65, 0x65, 0x64, 0x62,
	0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x62, 0x6f, 0x73,
	0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x6e, 0x69, 0x70, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x70, 0x69, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0d, 0x4f, 0x76, 0x65,
	0x72, 0x72, 0x69, 0x64, 0x65, 0x47, 0x72, 0x61, 0x73, 0x70, 0x12, 0x23, 0x2e, 0x62, 0x6f, 0x73,
	0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x70, 0x69, 0x47, 0x72, 0x61, 0x73, 0x70,
	0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x70, 0x69,
	0x47, 0x72, 0x61, 0x73, 0x70, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x68, 0x73, 0x73, 0x6e, 0x76, 0x2f, 0x73, 0x70, 0x6f,
	0x74, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x65, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x2f, 0x67,
	0x6f, 0x2f, 0x62, 0x6f, 0x73, 0x64, 0x79, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_bosdyn_api_manipulation_api_service_proto_goTypes = []interface{}{
	(*ManipulationApiRequest)(nil),          // 0: bosdyn.api.ManipulationApiRequest
	(*ManipulationApiFeedbackRequest)(nil),  // 1: bosdyn.api.ManipulationApiFeedbackRequest
	(*ApiGraspOverrideRequest)(nil),         // 2: bosdyn.api.ApiGraspOverrideRequest
	(*ManipulationApiResponse)(nil),         // 3: bosdyn.api.ManipulationApiResponse
	(*ManipulationApiFeedbackResponse)(nil), // 4: bosdyn.api.ManipulationApiFeedbackResponse
	(*ApiGraspOverrideResponse)(nil),        // 5: bosdyn.api.ApiGraspOverrideResponse
}
var file_bosdyn_api_manipulation_api_service_proto_depIdxs = []int32{
	0, // 0: bosdyn.api.ManipulationApiService.ManipulationApi:input_type -> bosdyn.api.ManipulationApiRequest
	1, // 1: bosdyn.api.ManipulationApiService.ManipulationApiFeedback:input_type -> bosdyn.api.ManipulationApiFeedbackRequest
	2, // 2: bosdyn.api.ManipulationApiService.OverrideGrasp:input_type -> bosdyn.api.ApiGraspOverrideRequest
	3, // 3: bosdyn.api.ManipulationApiService.ManipulationApi:output_type -> bosdyn.api.ManipulationApiResponse
	4, // 4: bosdyn.api.ManipulationApiService.ManipulationApiFeedback:output_type -> bosdyn.api.ManipulationApiFeedbackResponse
	5, // 5: bosdyn.api.ManipulationApiService.OverrideGrasp:output_type -> bosdyn.api.ApiGraspOverrideResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bosdyn_api_manipulation_api_service_proto_init() }
func file_bosdyn_api_manipulation_api_service_proto_init() {
	if File_bosdyn_api_manipulation_api_service_proto != nil {
		return
	}
	file_bosdyn_api_manipulation_api_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bosdyn_api_manipulation_api_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bosdyn_api_manipulation_api_service_proto_goTypes,
		DependencyIndexes: file_bosdyn_api_manipulation_api_service_proto_depIdxs,
	}.Build()
	File_bosdyn_api_manipulation_api_service_proto = out.File
	file_bosdyn_api_manipulation_api_service_proto_rawDesc = nil
	file_bosdyn_api_manipulation_api_service_proto_goTypes = nil
	file_bosdyn_api_manipulation_api_service_proto_depIdxs = nil
}
