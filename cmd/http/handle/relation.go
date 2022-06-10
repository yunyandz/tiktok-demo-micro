package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/user"
)

type RealationRequest struct {
	Token      string `form:"token" binding:"required"`
	ToUserId   uint64 `form:"to_user_id" binding:"required"`
	ActionType int8   `form:"action_type" binding:"required"`
}

const (
	RelationActionTypeFollow   = 1
	RelationActionTypeUnfollow = 2
)

func (h *Handle) RelationAction(c *gin.Context) {
	var req RealationRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	uc, e := h.getUserClaims(c)
	if !e {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "token error",
		})
		return
	}
	switch req.ActionType {
	case RelationActionTypeFollow:
		r := user.FollowUserRequest{
			UserId:   uc.UserId,
			FollowId: req.ToUserId,
		}
		resp, err := h.relation.FollowUser(c, &r)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, resp)
	case RelationActionTypeUnfollow:
		r := user.UnFollowUserRequest{
			UserId:   uc.UserId,
			FollowId: req.ToUserId,
		}
		resp, err := h.relation.UnFollowUser(c, &r)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, resp)
	default:
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "action type error",
		})
		return
	}

}

type FollowListRequest struct {
	UserId uint64 `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

func (h *Handle) FollowerList(c *gin.Context) {
	var req FollowListRequest
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
	r := user.GetFollowUserListRequest{
		SelfId: selfid,
		UserId: req.UserId,
	}
	resp, err := h.relation.GetFollowUserList(c, &r)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handle) FollowingList(c *gin.Context) {
	var req FollowListRequest
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
	r := user.GetFollowerUserListRequest{
		SelfId: selfid,
		UserId: req.UserId,
	}
	resp, err := h.relation.GetFollowerUserList(c, &r)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
