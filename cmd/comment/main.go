package main

import (
	"log"

	comment_service "github.com/yunyandz/tiktok-demo-micro/kitex_gen/comment_service/commentservice"
)

func main() {
	svr := comment_service.NewServer(new(CommentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
