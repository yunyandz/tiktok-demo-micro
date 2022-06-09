package main

import (
	"log"

	feed_service "github.com/yunyandz/tiktok-demo-micro/kitex_gen/feed_service/feedservice"
)

func main() {
	svr := feed_service.NewServer(new(FeedServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
