package main

import (
	"context"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.GetUserRequest) (resp *user.UserResponse, err error) {
	// TODO: Your code here...
	return
}

// FollowUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) FollowUser(ctx context.Context, req *user.FollowUserRequest) (resp *user.UserResponse, err error) {
	// TODO: Your code here...
	return
}

// UnFollowUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UnFollowUser(ctx context.Context, req *user.UnFollowUserRequest) (resp *user.UserResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFollowUserList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFollowUserList(ctx context.Context, req *user.GetFollowUserListRequest) (resp *user.UserListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFollowerUserList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFollowerUserList(ctx context.Context, req *user.GetFollowerUserListRequest) (resp *user.UserListResponse, err error) {
	// TODO: Your code here...
	return
}
