package main

import (
	"context"
	"github.com/yunyandz/tiktok-demo-micro/kitex_gen/feed_service"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// GetFeed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) GetFeed(ctx context.Context, req *feed_service.FeedRequest) (resp *feed_service.FeedResPonse, err error) {
	// TODO: Your code here...
	return
}
