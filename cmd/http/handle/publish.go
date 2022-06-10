package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	publish "github.com/yunyandz/tiktok-demo-micro/kitex_gen/publish_service"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/video"
)

type PublishRequest struct {
	Token string `form:"token" binding:"required"`
	Title string `form:"title" binding:"required"`
	Data  []byte `form:"data" binding:"required"`
}

func (h *Handle) Publish(c *gin.Context) {
	var req PublishRequest
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
	r := publish.PublishRequest{
		UserId: uc.UserId,
		Title:  req.Title,
		Data:   req.Data,
	}
	resp, err := h.publish.Publish(c, &r)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

type PublishListRequest struct {
	Token  string `form:"token" binding:"required"`
	UserID uint64 `form:"user_id" binding:"required"`
}

func (h *Handle) PublishList(c *gin.Context) {
	var req PublishListRequest
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
	r := video.GetUserVideosRequest{
		SelfId: selfid,
		UserId: req.UserID,
	}
	resp, err := h.video.GetUserVideos(c, &r)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
