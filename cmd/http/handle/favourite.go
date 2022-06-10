package handle

import "github.com/gin-gonic/gin"

type FavouriteActionRequest struct {
	Token      string `form:"token" binding:"required"`
	VideoId    uint64 `form:"video_id" binding:"required"`
	ActionType int8   `form:"action_type" binding:"required"`
}

func (h *Handle) FavouriteAction(c *gin.Context) {

}

func (h *Handle) FavouriteList(c *gin.Context) {

}
