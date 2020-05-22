package main

import (
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/service"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/consul"
	"shopping/user/handler"
	"shopping/user/model"
	"shopping/user/repository"
	//"shopping/user/subscriber"

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
	db, err := CreateConnection()
	if err != nil {
		log.Fatalf("connection error: %v\n", err)
	}
	defer db.Close()
	db.AutoMigrate(&model.User{})

	repo := &repository.User{db}

	// New Service
	serverService := grpc.NewService(
		service.Name("go.micro.srv.user"),
		service.Registry(consul.NewRegistry(func(options *registry.Options) {
			options.Addrs = []string{"0.0.0.0:8500"}
		})),
		service.Version("latest"),
	)

	// Initialise service
	serverService.Init()

	// Register Handler
	user.RegisterUserServiceHandler(serverService.Server(), &handler.User{Repo: repo})

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("go.micro.srv.user", serverService.Server(), new(subscriber.User))

	// Register Function as Subscriber
	// micro.RegisterSubscriber("go.micro.srv.user", serverService.Server(), subscriber.Handler)

	// Run service
	if err := serverService.Run(); err != nil {
		log.Fatal(err)
	}
}
