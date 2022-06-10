package handle

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-micro/internal/jwtx"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/comment_service/commentservice"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/feed_service/feedservice"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/publish_service/publishservice"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/relation_service/relationservice"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/user_service/userservice"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/video_service/videoservice"
	"go.uber.org/zap"
)

type Handle struct {
	comment  commentservice.Client
	feed     feedservice.Client
	publish  publishservice.Client
	user     userservice.Client
	video    videoservice.Client
	relation relationservice.Client
	logger   *zap.Logger
}

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type VideoListResponse struct {
	Response
	VideoList []*Video `json:"video_list"`
}

var (
	handle *Handle
	once   sync.Once
)

func New(
	comment commentservice.Client,
	feed feedservice.Client,
	publish publishservice.Client,
	user userservice.Client,
	video videoservice.Client,
	relation relationservice.Client,
	logger *zap.Logger,
) *Handle {
	once.Do(func() {
		handle = &Handle{
			comment:  comment,
			feed:     feed,
			publish:  publish,
			user:     user,
			video:    video,
			relation: relation,
			logger:   logger,
		}
	})
	return handle
}

func (hdl *Handle) getUserClaims(c *gin.Context) (*jwtx.UserClaims, bool) {
	uc, e := c.Get("claims")
	if !e {
		hdl.logger.Sugar().Errorf("Get claims error: %v", e)
		return nil, false
	}
	return uc.(*jwtx.UserClaims), true
}
