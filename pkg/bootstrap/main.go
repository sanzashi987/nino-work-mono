package bootstrap

import (
	"fmt"
	"net/http"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"github.com/sanzashi987/nino-work/config"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
	"go-micro.dev/v4/web"
)

type ServiceClient struct {
	Name   string
	Client client.Client
}

func GetAddress(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}

type Bootstraper struct {
	psm          string
	Config       *config.Config
	PsmConf      *config.ServiceConfig
	EtcdRegistry registry.Registry
}

func CommonBootstrap(psm string) Bootstraper {
	conf := config.GetConfig()
	etcdRegistry := etcd.NewRegistry(
		registry.Addrs(GetAddress(conf.System.EtcdHost, conf.System.EtcdPort)),
	)
	psmConf, ok := conf.Service[psm]
	if !ok {
		panic(psm + " is not configured")
	}

	return Bootstraper{
		psm:          psm,
		Config:       conf,
		PsmConf:      psmConf,
		EtcdRegistry: etcdRegistry,
	}
}

func (b Bootstraper) InitRpcService() micro.Service {

	rpcService := micro.NewService(
		micro.Name(b.PsmConf.Name),
		micro.Address(GetAddress(b.PsmConf.Host, b.PsmConf.Port)),
		micro.Registry(b.EtcdRegistry),
	)

	return rpcService
}

func (b Bootstraper) InitWebService(h http.Handler) web.Service {

	webService := web.NewService(
		web.Name(b.PsmConf.Name+".web"),
		web.Address(GetAddress(b.PsmConf.Host, b.PsmConf.WebPort)),
		web.Handler(h),
		web.Registry(b.EtcdRegistry),
	)
	return webService
}

func (b Bootstraper) InitRpcClient() *ServiceClient {
	mySelector := selector.NewSelector(
		selector.Registry(b.EtcdRegistry),
		selector.SetStrategy(selector.RoundRobin),
	)

	serviceClient := &ServiceClient{
		Name: b.PsmConf.Name + ".client",
		Client: client.NewClient(
			client.Selector(mySelector),
		),
	}

	return serviceClient
}
