package clientService

import (
	"context"

	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
)

type ServiceClient struct {
	name   string
	client client.Client
}

var serviceClient *ServiceClient

func InitClient(name string, reg registry.Registry) {
	mySelector := selector.NewSelector(
		selector.Registry(reg),
		selector.SetStrategy(selector.RoundRobin),
	)

	c := client.NewClient(
		client.Selector(mySelector),
	)

	serviceClient = &ServiceClient{
		name:   name,
		client: c,
	}

}

func GetClientSession() *ServiceClient {
	return serviceClient
}

func CallRpc(ctx context.Context, endpoint string, req *interface{}, res *interface{}, reqOpts []client.RequestOption, resOpts []client.CallOption) error {
	request := client.NewRequest(serviceClient.name, endpoint, req, reqOpts...)
	return serviceClient.client.Call(ctx, request, res, resOpts...)
}
