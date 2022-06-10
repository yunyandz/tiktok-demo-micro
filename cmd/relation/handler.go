package main

import (
	"context"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/user"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// FollowUser implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FollowUser(ctx context.Context, req *user.FollowUserRequest) (resp *user.UserResponse, err error) {
	// TODO: Your code here...
	return
}

// UnFollowUser implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) UnFollowUser(ctx context.Context, req *user.UnFollowUserRequest) (resp *user.UserResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFollowUserList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowUserList(ctx context.Context, req *user.GetFollowUserListRequest) (resp *user.UserListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFollowerUserList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowerUserList(ctx context.Context, req *user.GetFollowerUserListRequest) (resp *user.UserListResponse, err error) {
	// TODO: Your code here...
	return
}
