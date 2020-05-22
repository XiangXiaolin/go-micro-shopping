package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/service"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-plugins/registry/consul"
	user "shopping/user/proto/user"
)

func main() {
	clientService := grpc.NewService(
		service.Name("go.micro.srv.user"),
		service.Registry(consul.NewRegistry(func(options *registry.Options) {
			options.Addrs = []string{"0.0.0.0:8500"}
		})),
		service.Version("latest"),
	)

	clientService.Init()

	regUser := &user.User{
		Name:     "allin",
		Phone:    "18102390678",
		Password: "123456",
	}

	client := user.NewUserService("go.micro.srv.user", clientService.Client())
	response, err := client.Register(context.TODO(), &user.RegisterRequest{
		User: regUser,
	})
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(response.Code)
	fmt.Println(response.Msg)
}
