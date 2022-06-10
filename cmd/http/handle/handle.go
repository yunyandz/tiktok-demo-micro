package handle

import (
	"sync"

	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/comment_service/commentservice"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/feed_service/feedservice"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/publish_service/publishservice"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/user_service/userservice"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/video_service/videoservice"
)

type Handle struct {
	comment commentservice.Client
	feed    feedservice.Client
	publish publishservice.Client
	user    userservice.Client
	video   videoservice.Client
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
) *Handle {
	once.Do(func() {
		handle = &Handle{
			comment: comment,
			feed:    feed,
			publish: publish,
			user:    user,
			video:   video,
		}
	})
	return handle
}
