package main

import (
	relation_service "github.com/yunyandz/tiktok-demo-micro/kitex_gen/relation_service/relationservice"
	"log"
)

func main() {
	svr := relation_service.NewServer(new(RelationServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
