package handle

import "github.com/gin-gonic/gin"

type RealationRequest struct {
	Token      string `form:"token" binding:"required"`
	ToUserId   uint64 `form:"to_user_id" binding:"required"`
	ActionType int8   `form:"action_type" binding:"required"`
}

func (h *Handle) RelationAction(c *gin.Context) {

}

type FollowRequest struct {
	UserId uint64 `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

func (h *Handle) FollowerList(c *gin.Context) {

}

func (h *Handle) FollowingList(c *gin.Context) {

}
