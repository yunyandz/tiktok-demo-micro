package main

import (
	"context"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/publish_service"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/response"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// Publish implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) Publish(ctx context.Context, req *publish_service.PublishRequest) (resp *response.Response, err error) {
	// TODO: Your code here...
	return
}
