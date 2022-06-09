package main

import (
	"context"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/response"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/video"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// LikeVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) LikeVideo(ctx context.Context, req *video.LikeVideoRequest) (resp *response.Response, err error) {
	// TODO: Your code here...
	return
}

// DislikeVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) DislikeVideo(ctx context.Context, req *video.DislikeVideoRequest) (resp *response.Response, err error) {
	// TODO: Your code here...
	return
}

// GetUserVideos implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetUserVideos(ctx context.Context, req *video.GetUserVideosRequest) (resp *video.VideoListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserLikedVideos implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetUserLikedVideos(ctx context.Context, req *video.GetUserLikedVideosRequest) (resp *video.VideoListResponse, err error) {
	// TODO: Your code here...
	return
}
