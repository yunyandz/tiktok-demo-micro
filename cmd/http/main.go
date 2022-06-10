package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-micro/cmd/http/handle"
	"github.com/yunyandz/tiktok-demo-micro/internal/config"
	"github.com/yunyandz/tiktok-demo-micro/internal/jwtx"
	"github.com/yunyandz/tiktok-demo-micro/internal/logger"
	"github.com/yunyandz/tiktok-demo-micro/internal/rpc"
)

func main() {
	cfg := config.New()
	cc := rpc.NewCommentClient(cfg)
	fd := rpc.NewFeedClient(cfg)
	pu := rpc.NewPublishClient(cfg)
	us := rpc.NewUserClient(cfg)
	vi := rpc.NewVideoClient(cfg)
	rl := rpc.NewRelationClient(cfg)
	lg := logger.New(cfg)
	ps := jwtx.NewParser(cfg)
	hl := handle.New(cc, fd, pu, us, vi, rl, lg)

	if cfg.GetBool("debug") {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.Default()

	InitRouter(lg, r, hl, ps)
	host := strings.Join([]string{cfg.GetString("http.host"), cfg.GetString("http.port")}, ":")
	if err := r.Run(host); err != nil {
		panic(err)
	}
}
