package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/video"
)

type FavouriteActionRequest struct {
	Token      string `form:"token" binding:"required"`
	VideoId    uint64 `form:"video_id" binding:"required"`
	ActionType int8   `form:"action_type" binding:"required"`
}

type FavouriteListRequest struct {
	Token  string `form:"token" binding:"required"`
	UserId uint64 `form:"user_id" binding:"required"`
}

const (
	FavouriteActionTypeCreate = 1
	FavouriteActionTypeDelete = 2
)

func (h *Handle) FavouriteAction(c *gin.Context) {
	var req FavouriteActionRequest
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
	case FavouriteActionTypeCreate:
		r := video.LikeVideoRequest{
			VideoId: req.VideoId,
			UserId:  uc.UserId,
		}

		resp, err := h.video.LikeVideo(c, &r)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, resp)
	case FavouriteActionTypeDelete:
		r := video.DislikeVideoRequest{
			VideoId: req.VideoId,
			UserId:  uc.UserId,
		}
		resp, err := h.video.DislikeVideo(c, &r)
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

func (h *Handle) FavouriteList(c *gin.Context) {
	var req FavouriteListRequest
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
	r := video.GetUserLikedVideosRequest{
		SelfId: selfid,
		UserId: req.UserId,
	}
	resp, err := h.video.GetUserLikedVideos(c, &r)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
