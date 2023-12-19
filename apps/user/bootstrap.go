package main

import (
	"fmt"

	"github.com/cza14h/nino-work/apps/user/db/dao"
	"github.com/cza14h/nino-work/apps/user/service"
	"github.com/cza14h/nino-work/config"
	"github.com/cza14h/nino-work/pkg/clientService"
	"github.com/cza14h/nino-work/proto/user"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func getAddress(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}

func main() {
	config.LoadConfig()
	dao.ConnectDB()

	etcdRegistry := etcd.NewRegistry(
		registry.Addrs(getAddress(config.EtcdHost, config.EtcdPort)),
	)

	microService := micro.NewService(
		micro.Name(config.UserServiceName),
		micro.Address(getAddress(config.UserServiceHost, config.UserServicePort)),
		micro.Registry(etcdRegistry),
	)

	clientService.InitClient(config.UserServiceName, etcdRegistry)

	microService.Init()
	user.RegisterUserSerivceHandler(microService.Server(), service.GetUserServiceRpc())
	microService.Run()

}