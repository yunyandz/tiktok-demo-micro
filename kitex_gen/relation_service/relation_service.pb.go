// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: relation_service.proto

package relation_service

import (
	context "context"
	user "github.com/yunyandz/tiktok-demo-micro/kitex_gen/user"
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

var File_relation_service_proto protoreflect.FileDescriptor

var file_relation_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa9, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x55, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x6e, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x79, 0x75, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x7a, 0x2f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b,
	0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65,
	0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_relation_service_proto_goTypes = []interface{}{
	(*user.FollowUserRequest)(nil),          // 0: user.FollowUserRequest
	(*user.UnFollowUserRequest)(nil),        // 1: user.UnFollowUserRequest
	(*user.GetFollowUserListRequest)(nil),   // 2: user.GetFollowUserListRequest
	(*user.GetFollowerUserListRequest)(nil), // 3: user.GetFollowerUserListRequest
	(*user.UserResponse)(nil),               // 4: user.UserResponse
	(*user.UserListResponse)(nil),           // 5: user.UserListResponse
}
var file_relation_service_proto_depIdxs = []int32{
	0, // 0: relation_service.RelationService.FollowUser:input_type -> user.FollowUserRequest
	1, // 1: relation_service.RelationService.UnFollowUser:input_type -> user.UnFollowUserRequest
	2, // 2: relation_service.RelationService.GetFollowUserList:input_type -> user.GetFollowUserListRequest
	3, // 3: relation_service.RelationService.GetFollowerUserList:input_type -> user.GetFollowerUserListRequest
	4, // 4: relation_service.RelationService.FollowUser:output_type -> user.UserResponse
	4, // 5: relation_service.RelationService.UnFollowUser:output_type -> user.UserResponse
	5, // 6: relation_service.RelationService.GetFollowUserList:output_type -> user.UserListResponse
	5, // 7: relation_service.RelationService.GetFollowerUserList:output_type -> user.UserListResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_relation_service_proto_init() }
func file_relation_service_proto_init() {
	if File_relation_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_relation_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_relation_service_proto_goTypes,
		DependencyIndexes: file_relation_service_proto_depIdxs,
	}.Build()
	File_relation_service_proto = out.File
	file_relation_service_proto_rawDesc = nil
	file_relation_service_proto_goTypes = nil
	file_relation_service_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.3.2. DO NOT EDIT.

type RelationService interface {
	FollowUser(ctx context.Context, req *user.FollowUserRequest) (res *user.UserResponse, err error)
	UnFollowUser(ctx context.Context, req *user.UnFollowUserRequest) (res *user.UserResponse, err error)
	GetFollowUserList(ctx context.Context, req *user.GetFollowUserListRequest) (res *user.UserListResponse, err error)
	GetFollowerUserList(ctx context.Context, req *user.GetFollowerUserListRequest) (res *user.UserListResponse, err error)
}
