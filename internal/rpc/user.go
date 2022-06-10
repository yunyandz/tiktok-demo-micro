package rpc

import (
	"sync"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/spf13/viper"
	"github.com/yunyandz/tiktok-demo-micro/internal/constants"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/user_service/userservice"
)

var (
	userClient userservice.Client
	userOnce   sync.Once
)

func NewUserClient(cfg *viper.Viper) userservice.Client {
	userOnce.Do(func() {
		r, err := etcd.NewEtcdResolver([]string{cfg.GetString("etcd.addr")})
		if err != nil {
			panic(err)
		}
		userClient, err = userservice.NewClient(
			constants.CommentServiceName,
			client.WithResolver(r),
			client.WithLongConnection(connpool.IdleConfig{
				MaxIdlePerAddress: 10,
				MaxIdleGlobal:     1000,
				MaxIdleTimeout:    time.Minute}),
			client.WithMuxConnection(1),
			client.WithRPCTimeout(3*time.Second),
			client.WithSuite(trace.NewDefaultClientSuite()),
			client.WithConnectTimeout(50*time.Millisecond),
		)
	})
	return userClient
}
