package bootstrap

import (
	"fmt"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"github.com/sanzashi987/nino-work/config"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
)

type ServiceClient struct {
	name   string
	client client.Client
}

func (c ServiceClient) GetName() string {
	return c.name
}
func (c ServiceClient) GetClientInstance() client.Client {
	return c.client
}

var serviceClient *ServiceClient

func initClient(name string, reg registry.Registry) {
	mySelector := selector.NewSelector(
		selector.Registry(reg),
		selector.SetStrategy(selector.RoundRobin),
	)

	serviceClient = &ServiceClient{
		name: name,
		client: client.NewClient(
			client.Selector(mySelector),
		),
	}
}

func GetClient() *ServiceClient {
	return serviceClient
}

func GetAddress(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}

func CommonBootstrap(name string) (*config.Config, registry.Registry) {
	conf := config.GetConfig()
	etcdRegistry := etcd.NewRegistry(
		registry.Addrs(GetAddress(conf.System.EtcdHost, conf.System.EtcdPort)),
	)

	initClient(name, etcdRegistry)
	return conf, etcdRegistry
}
