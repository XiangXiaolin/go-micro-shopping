package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/service"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-plugins/registry/consul"

	product "shopping/product/proto/product"
)

func main() {
	clientService := grpc.NewService(
		service.Name("go.micro.srv.product"),
		service.Registry(consul.NewRegistry(func(options *registry.Options) {
			options.Addrs = []string{"0.0.0.0:8500"}
		})),
		service.Version("latest"),
	)

	clientService.Init()

	client := product.NewProductService("go.micro.srv.product", clientService.Client())
	response, err := client.Search(context.TODO(), &product.SearchRequest{Name: "dog"})
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(response.Code)
	fmt.Println(response.Msg)
	fmt.Println(response.Products)
}
