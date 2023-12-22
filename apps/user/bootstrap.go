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
	conf, etcdRegistry := bootstrap.CommonBootstrap("userSevice.client")
	dao.ConnectDB()

	userServiceConf, ok := conf.Service["userSevice"]
	if !ok {
		panic("UserService is not configured")
	}

	rpcService := micro.NewService(
		micro.Name(userServiceConf.Name),
		micro.Address(bootstrap.GetAddress(userServiceConf.Host, userServiceConf.Port)),
		micro.Registry(etcdRegistry),
	)

	webService := web.NewService(
		web.Name(userServiceConf.Name+".web"),
		web.Address(bootstrap.GetAddress(userServiceConf.Host, userServiceConf.WebPort)),
		web.Handler(http.NewRouter()),
		web.Registry(etcdRegistry),
	)

	user.RegisterUserSerivceHandler(rpcService.Server(), service.GetUserServiceRpc())
	webService.Init()
	rpcService.Init()

	go func() {
		webService.Run()
	}()

	rpcService.Run()

}
