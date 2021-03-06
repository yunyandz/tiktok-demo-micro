// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: video_service.proto

package video_service

import (
	context "context"
	response "github.com/yunyandz/tiktok-demo-micro/kitex_gen/response"
	video "github.com/yunyandz/tiktok-demo-micro/kitex_gen/video"
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

var File_video_service_proto protoreflect.FileDescriptor

var file_video_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x1a, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xa2, 0x02, 0x0a, 0x0c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x4c, 0x69, 0x6b, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12,
	0x17, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0c,
	0x44, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x1a, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x44, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x1b, 0x2e,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x69, 0x6b, 0x65, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x20, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x7a, 0x2f, 0x74, 0x69,
	0x6b, 0x74, 0x6f, 0x6b, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f,
	0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_video_service_proto_goTypes = []interface{}{
	(*video.LikeVideoRequest)(nil),          // 0: video.LikeVideoRequest
	(*video.DislikeVideoRequest)(nil),       // 1: video.DislikeVideoRequest
	(*video.GetUserVideosRequest)(nil),      // 2: video.GetUserVideosRequest
	(*video.GetUserLikedVideosRequest)(nil), // 3: video.GetUserLikedVideosRequest
	(*response.Response)(nil),               // 4: response.Response
	(*video.VideoListResponse)(nil),         // 5: video.VideoListResponse
}
var file_video_service_proto_depIdxs = []int32{
	0, // 0: video_service.VideoService.LikeVideo:input_type -> video.LikeVideoRequest
	1, // 1: video_service.VideoService.DislikeVideo:input_type -> video.DislikeVideoRequest
	2, // 2: video_service.VideoService.GetUserVideos:input_type -> video.GetUserVideosRequest
	3, // 3: video_service.VideoService.GetUserLikedVideos:input_type -> video.GetUserLikedVideosRequest
	4, // 4: video_service.VideoService.LikeVideo:output_type -> response.Response
	4, // 5: video_service.VideoService.DislikeVideo:output_type -> response.Response
	5, // 6: video_service.VideoService.GetUserVideos:output_type -> video.VideoListResponse
	5, // 7: video_service.VideoService.GetUserLikedVideos:output_type -> video.VideoListResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_video_service_proto_init() }
func file_video_service_proto_init() {
	if File_video_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_video_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_video_service_proto_goTypes,
		DependencyIndexes: file_video_service_proto_depIdxs,
	}.Build()
	File_video_service_proto = out.File
	file_video_service_proto_rawDesc = nil
	file_video_service_proto_goTypes = nil
	file_video_service_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.3.2. DO NOT EDIT.

type VideoService interface {
	LikeVideo(ctx context.Context, req *video.LikeVideoRequest) (res *response.Response, err error)
	DislikeVideo(ctx context.Context, req *video.DislikeVideoRequest) (res *response.Response, err error)
	GetUserVideos(ctx context.Context, req *video.GetUserVideosRequest) (res *video.VideoListResponse, err error)
	GetUserLikedVideos(ctx context.Context, req *video.GetUserLikedVideosRequest) (res *video.VideoListResponse, err error)
}
