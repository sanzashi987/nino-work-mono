package main

import (
	"github.com/cza14h/nino-work/pkg/bootstrap"
	"go-micro.dev/v4"
)

func main() {

	conf, etcdRegistry := bootstrap.CommonBootstrap("fileService.client")

	fileServiceConf, ok := conf.Service["fileService"]
	if !ok {
		panic("File Service is not configured")
	}

	rpcService := micro.NewService(
		micro.Name(fileServiceConf.Name),
		micro.Address(bootstrap.GetAddress(fileServiceConf.Host, fileServiceConf.Port)),
		micro.Registry(etcdRegistry),
	)

	// webService := web.NewService(
	// 	web.Name(fileServiceConf.Name+".web"),
	// 	web.Address(bootstrap.GetAddress(fileServiceConf.Host, fileServiceConf.Port)),
	// 	web.Handler(http.NewRouter(conf.System.LoginPage)),
	// 	web.Registry(etcdRegistry),
	// )

	// user.RegisterUserSerivceHandler(rpcService.Server(), service.GetUserServiceRpc())

	// go func() {
	// 	webService.Run()
	// }()

	rpcService.Init()
	rpcService.Run()

}
