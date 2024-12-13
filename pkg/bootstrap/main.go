package bootstrap

import (
	"fmt"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"github.com/sanzashi987/nino-work/config"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
)

type ServiceClient struct {
	name   string
	client client.Client
}

func (c ServiceClient) Name() string {
	return c.name
}
func (c ServiceClient) Client() client.Client {
	return c.client
}

func InitClient(name string, reg registry.Registry) *ServiceClient {
	mySelector := selector.NewSelector(
		selector.Registry(reg),
		selector.SetStrategy(selector.RoundRobin),
	)

	serviceClient := &ServiceClient{
		name: name,
		client: client.NewClient(
			client.Selector(mySelector),
		),
	}

	return serviceClient
}

func GetAddress(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}

type Bootstraper struct {
	psm          string
	config       *config.Config
	etcdRegistry *registry.Registry
}

func CommonBootstrap(psm string) Bootstraper {
	conf := config.GetConfig()
	etcdRegistry := etcd.NewRegistry(
		registry.Addrs(GetAddress(conf.System.EtcdHost, conf.System.EtcdPort)),
	)

	return Bootstraper{
		psm:          psm,
		config:       conf,
		etcdRegistry: &etcdRegistry,
	}
}

func (b Bootstraper) InitRpcService() micro.Service {
	psmConf, ok := b.config.Service[b.psm]
	if !ok {
		panic(b.psm + " is not configured")
	}

	rpcService := micro.NewService(
		micro.Name(psmConf.Name),
	)

	return rpcService

}
