package main

import (
	"log"

	video_service "github.com/yunyandz/tiktok-demo-micro/kitex_gen/video_service/videoservice"
)

func main() {
	svr := video_service.NewServer(new(VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
