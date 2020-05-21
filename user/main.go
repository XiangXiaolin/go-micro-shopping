package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"shopping/user/handler"
	"shopping/user/subscriber"

	user "shopping/user/proto/user"
)

// 开发步骤：
// 1.定义接口 ->
// 2.生成接口代码 ->
// 3.编写model层代码 ->
// 4.编写repository数据操作代码 ->
// 5.实现接口 ->
// 6.修改main.go

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.user", service.Server(), new(subscriber.User))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
