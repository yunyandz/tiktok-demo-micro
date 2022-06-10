package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/user"
)

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

func (h *Handle) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	r := user.UserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	}
	resp, err := h.user.UserRegister(c, &r)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
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

func (h *Handle) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	r := user.UserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	}
	resp, err := h.user.UserLogin(c, &r)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
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
	var req UserInfoRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	selfid := uint64(0)
	uc, e := h.getUserClaims(c)
	if e {
		selfid = uc.UserId
	}
	r := user.GetUserRequest{
		UserId: req.UserID,
		SelfId: selfid,
	}
	resp, err := h.user.GetUser(c, &r)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
