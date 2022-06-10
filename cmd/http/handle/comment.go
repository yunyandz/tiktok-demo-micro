package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/comment"
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

type CommentListRequest struct {
	Token   string `form:"token" binding:"required"`
	VideoId uint64 `form:"video_id" binding:"required"`
}

const (
	CommentActionTypeCreate = 1
	CommentActionTypeDelete = 2
)

func (h *Handle) CommentAction(c *gin.Context) {
	var req CommentActionRequest
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
	case CommentActionTypeCreate:
		r := comment.PublishCommentRequest{
			VideoId: req.VideoId,
			Content: req.CommentText,
			UserId:  uc.UserId,
		}
		resp, e := h.comment.PublishComment(c, &r)
		if e != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  e.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, resp)
	case CommentActionTypeDelete:
		r := comment.DeleteCommentRequest{
			CommentId: req.CommentID,
			UserId:    uc.UserId,
		}
		resp, err := h.comment.DeleteComment(c, &r)
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

func (h *Handle) CommentList(c *gin.Context) {
	var req CommentListRequest
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
	r := comment.GetCommentListRequest{
		VideoId: req.VideoId,
		SelfId:  selfid,
	}
	resp, err := h.comment.GetCommentList(c, &r)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
