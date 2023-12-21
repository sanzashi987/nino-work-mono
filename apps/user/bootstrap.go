package main

import (
	"github.com/cza14h/nino-work/apps/user/db/dao"
	"github.com/cza14h/nino-work/apps/user/http"
	"github.com/cza14h/nino-work/apps/user/service"
	"github.com/cza14h/nino-work/pkg/bootstrap"
	"github.com/cza14h/nino-work/proto/user"
	"go-micro.dev/v4"
	"go-micro.dev/v4/web"
)

func main() {
	config, etcdRegistry := bootstrap.CommonBootstrap("userSevice.client")
	dao.ConnectDB()

	rpcService := micro.NewService(
		micro.Name(config.UserServiceName),
		micro.Address(bootstrap.GetAddress(config.UserServiceHost, config.UserServicePort)),
		micro.Registry(etcdRegistry),
	)

	webService := web.NewService(
		web.Name(config.UserServiceName+".web"),
		web.Address(bootstrap.GetAddress(config.UserServiceHost, config.UserServiceWebPort)),
		web.Handler(http.NewRouter()),
		web.Registry(etcdRegistry),
	)

	webService.Init()
	rpcService.Init()
	user.RegisterUserSerivceHandler(rpcService.Server(), service.GetUserServiceRpc())

	go func() {
		webService.Run()
	}()
	
	rpcService.Run()

}
