package clientService

import (
	"context"

	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
)

var microClient client.Client

func InitClient(reg registry.Registry) {
	mySelector := selector.NewSelector(
		selector.Registry(reg),
		selector.SetStrategy(selector.RoundRobin),
	)

	microClient = client.NewClient(
		client.Selector(mySelector),
	)

}

func GetClientSession() *client.Client {
	return &microClient
}

func CallRpc(ctx context.Context, service, endpoint string, req *interface{}, res *interface{}, reqOpts []client.RequestOption, resOpts []client.CallOption) error {
	request := client.NewRequest(service, endpoint, req, reqOpts...)
	return microClient.Call(ctx, request, res, resOpts...)
}
