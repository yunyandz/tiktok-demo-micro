package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-micro/cmd/http/handle"
	"github.com/yunyandz/tiktok-demo-micro/internal/jwtx"
	"go.uber.org/zap"
)

func InitRouter(logger *zap.Logger, r *gin.Engine, h *handle.Handle, ps *jwtx.Parser) {
	apiRouter := r.Group("/douyin")

	apiRouter.GET("/feed/", JWTAuth(logger, false, ps), h.Feed)

	// using jwt auth
	apiRouter.GET("/user/", JWTAuth(logger, true, ps), h.UserInfo)

	apiRouter.POST("/user/register/", h.Register)
	apiRouter.POST("/user/login/", h.Login)

	apiRouter.POST("/publish/action/", JWTAuth(logger, true, ps), h.Publish)
	apiRouter.GET("/publish/list/", JWTAuth(logger, false, ps), h.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", JWTAuth(logger, true, ps), h.FavouriteAction)
	apiRouter.GET("/favorite/list/", JWTAuth(logger, false, ps), h.FavouriteList)

	apiRouter.POST("/comment/action/", JWTAuth(logger, true, ps), h.CommentAction)
	apiRouter.GET("/comment/list/", JWTAuth(logger, false, ps), h.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", JWTAuth(logger, true, ps), h.RelationAction)
	apiRouter.GET("/relation/follow/list/", JWTAuth(logger, false, ps), h.FollowingList)
	apiRouter.GET("/relation/follower/list/", JWTAuth(logger, false, ps), h.FollowerList)
}
