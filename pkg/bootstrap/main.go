package bootstrap

import (
	"fmt"

	"github.com/cza14h/nino-work/config"
	"github.com/go-micro/plugins/v4/registry/etcd"
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

	serviceClient = &ServiceClient{
		name: name,
		client: client.NewClient(
			client.Selector(mySelector),
		),
	}
}

func GetAddress(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}

func CommonBootstrap(name string) (*config.Config, registry.Registry) {
	configInstance := config.InitConfig()
	etcdRegistry := etcd.NewRegistry(
		registry.Addrs(GetAddress(configInstance.EtcdHost, configInstance.EtcdPort)),
	)

	InitClient(name, etcdRegistry)
	return configInstance, etcdRegistry
}