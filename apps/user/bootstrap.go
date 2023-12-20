package main

import (
	"github.com/cza14h/nino-work/apps/user/db/dao"
	"github.com/cza14h/nino-work/apps/user/service"
	"github.com/cza14h/nino-work/pkg/bootstrap"
	"github.com/cza14h/nino-work/proto/user"
	"go-micro.dev/v4"
)

func main() {
	config, etcdRegistry := bootstrap.CommonBootstrap("userSevice.client")
	dao.ConnectDB()

	microService := micro.NewService(
		micro.Name(config.UserServiceName),
		micro.Address(bootstrap.GetAddress(config.UserServiceHost, config.UserServicePort)),
		micro.Registry(etcdRegistry),
	)

	microService.Init()
	user.RegisterUserSerivceHandler(microService.Server(), service.GetUserServiceRpc())
	microService.Run()

}
