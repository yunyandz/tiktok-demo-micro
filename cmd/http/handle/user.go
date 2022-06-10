package handle

import "github.com/gin-gonic/gin"

type User struct {
	ID            uint64 `json:"id"`
	Username      string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type RegisterRequest struct {
	Username string `form:"username" binding:"required,min=3,max=32"`
	Password string `form:"password" binding:"required,max=32"`
}

type RegisterResponse struct {
	Response
	UserID uint64 `json:"user_id"`
	Token  string `json:"token"`
}

func (h *Handle) UserRegister(c *gin.Context) {

}

type LoginRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	Response
	UserID uint64 `json:"user_id"`
	Token  string `json:"token"`
}

func (h *Handle) UserLogin(c *gin.Context) {

}

type UserInfoRequest struct {
	UserID uint64 `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

type UserInfoResponse struct {
	Response
	User User `json:"user"`
}

func (h *Handle) UserInfo(c *gin.Context) {

}
