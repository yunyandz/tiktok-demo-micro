package handle

import "github.com/gin-gonic/gin"

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
	Token       string `form:"token" binding:"required"`
	LastestTime int64  `form:"lastest_time" binding:"required"`
}

type FeedResponse struct {
	VideoListResponse
	NextTime int64 `json:"next_time,omitempty"`
}

func (h *Handle) Feed(c *gin.Context) {

}
