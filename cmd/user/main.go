package main

import (
	"log"

	user_service "github.com/yunyandz/tiktok-demo-micro/kitex_gen/user_service/userservice"
)

func main() {
	svr := user_service.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
