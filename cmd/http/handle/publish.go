package handle

import "github.com/gin-gonic/gin"

type PublishRequest struct {
	Token string `form:"token" binding:"required"`
	Title string `form:"title" binding:"required"`
	Data  []byte `form:"data" binding:"required"`
}

func (h *Handle) Publish(c *gin.Context) {

}

type PublishListRequest struct {
	Token  string `form:"token" binding:"required"`
	UserID uint64 `form:"user_id" binding:"required"`
}

func (h *Handle) PublishList(c *gin.Context) {

}
