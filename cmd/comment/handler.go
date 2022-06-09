package main

import (
	"context"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/comment"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// PublishComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) PublishComment(ctx context.Context, req *comment.PublishCommentRequest) (resp *comment.CommentActionResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) DeleteComment(ctx context.Context, req *comment.DeleteCommentRequest) (resp *comment.DeleteCommentResponse, err error) {
	// TODO: Your code here...
	return
}

// GetCommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) GetCommentList(ctx context.Context, req *comment.GetCommentListRequest) (resp *comment.GetCommentListResponse, err error) {
	// TODO: Your code here...
	return
}
