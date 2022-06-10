package handle

import (
	"github.com/gin-gonic/gin"
)

type Comment struct {
	Id         uint64 `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type CommentActionRequest struct {
	Token       string `form:"token" binding:"required"`
	VideoId     uint64 `form:"video_id" binding:"required"`
	ActionType  int8   `form:"action_type" binding:"required"`
	CommentText string `form:"comment_text" binding:"required"`
	CommentID   uint64 `form:"comment_id" binding:"required"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment"`
}

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

func (h *Handle) CommentAction(c *gin.Context) {

}

func (h *Handle) CommentList(c *gin.Context) {

}
