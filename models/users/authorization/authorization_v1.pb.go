/// This file has messages for describing authorization

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.7
// source: users/authorization/authorization_v1.proto

package authorization

import (
	common "github.com/badhouseplants/envspotting-go-proto/models/common"
	accounts "github.com/badhouseplants/envspotting-go-proto/models/users/accounts"
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

var File_users_authorization_authorization_v1_proto protoreflect.FileDescriptor

var file_users_authorization_authorization_v1_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x1a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc6, 0x01,
	0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x38, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0d, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x10, 0x50, 0x61, 0x72, 0x73,
	0x65, 0x49, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0x00, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x64, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x73, 0x2f, 0x65, 0x6e, 0x76, 0x73, 0x70, 0x6f, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2d,
	0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_users_authorization_authorization_v1_proto_goTypes = []interface{}{
	(*accounts.AccountId)(nil),  // 0: users.AccountId
	(*common.EmptyMessage)(nil), // 1: common.EmptyMessage
}
var file_users_authorization_authorization_v1_proto_depIdxs = []int32{
	0, // 0: users.Authorization.RefreshToken:input_type -> users.AccountId
	1, // 1: users.Authorization.ValidateToken:input_type -> common.EmptyMessage
	1, // 2: users.Authorization.ParseIdFromToken:input_type -> common.EmptyMessage
	1, // 3: users.Authorization.RefreshToken:output_type -> common.EmptyMessage
	1, // 4: users.Authorization.ValidateToken:output_type -> common.EmptyMessage
	0, // 5: users.Authorization.ParseIdFromToken:output_type -> users.AccountId
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_users_authorization_authorization_v1_proto_init() }
func file_users_authorization_authorization_v1_proto_init() {
	if File_users_authorization_authorization_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_users_authorization_authorization_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_users_authorization_authorization_v1_proto_goTypes,
		DependencyIndexes: file_users_authorization_authorization_v1_proto_depIdxs,
	}.Build()
	File_users_authorization_authorization_v1_proto = out.File
	file_users_authorization_authorization_v1_proto_rawDesc = nil
	file_users_authorization_authorization_v1_proto_goTypes = nil
	file_users_authorization_authorization_v1_proto_depIdxs = nil
}
