package main

import (
	"github.com/sanzashi987/nino-work/apps/user/db/dao"
	"github.com/sanzashi987/nino-work/apps/user/http"
	"github.com/sanzashi987/nino-work/apps/user/service"
	"github.com/sanzashi987/nino-work/pkg/bootstrap"
	"github.com/sanzashi987/nino-work/proto/user"
	"go-micro.dev/v4"
	"go-micro.dev/v4/web"
)

func main() {
	conf, etcdRegistry := bootstrap.CommonBootstrap()
	dao.ConnectDB()

	userServiceConf, ok := conf.Service["userService"]
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

	user.RegisterUserServiceHandler(rpcService.Server(), service.GetUserServiceRpc())
	webService.Init()
	rpcService.Init()

	go func() {
		webService.Run()
	}()

	rpcService.Run()

}
