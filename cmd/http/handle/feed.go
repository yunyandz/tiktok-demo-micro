package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	feed "github.com/yunyandz/tiktok-demo-micro/kitex_gen/feed_service"
)

type Video struct {
	Id            uint64 `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount uint64 `json:"favorite_count,omitempty"`
	CommentCount  uint64 `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty"`
}

type FeedRequest struct {
	Token      string `form:"token" binding:"required"`
	LatestTime int64  `form:"latest_time" binding:"required"`
}

type FeedResponse struct {
	VideoListResponse
	NextTime int64 `json:"next_time,omitempty"`
}

func (h *Handle) Feed(c *gin.Context) {
	var req FeedRequest
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
	r := feed.FeedRequest{
		UserId:     selfid,
		LatestTime: req.LatestTime,
	}
	resp, err := h.feed.GetFeed(c, &r)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
