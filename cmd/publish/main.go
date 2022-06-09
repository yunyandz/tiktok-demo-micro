package main

import (
	"log"

	publish_service "github.com/yunyandz/tiktok-demo-micro/kitex_gen/publish_service/publishservice"
)

func main() {
	svr := publish_service.NewServer(new(PublishServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
