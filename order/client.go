package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/service"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-plugins/registry/consul"
	order "shopping/order/proto/order"
)

func main() {
	clientService := grpc.NewService(
		service.Name("go.micro.srv.order"),
		service.Registry(consul.NewRegistry(func(options *registry.Options) {
			options.Addrs = []string{"0.0.0.0:8500"}
		})),
		service.Version("latest"),
	)

	clientService.Init()

	client := order.NewOrderService("go.micro.srv.order", clientService.Client())
	response, err := client.OrderDetail(context.TODO(), &order.OrderDetailRequest{OrderId: "1"})
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(response.Code)
	fmt.Println(response.Msg)
}
