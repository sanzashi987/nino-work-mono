package main

import (
	"github.com/cza14h/nino-work/apps/upload/service"
	"github.com/cza14h/nino-work/pkg/bootstrap"
	"github.com/cza14h/nino-work/proto/upload"
	"go-micro.dev/v4"
)

func main() {

	conf, etcdRegistry := bootstrap.CommonBootstrap("uploadService.client")

	uploadServiceConf, ok := conf.Service["uploadService"]
	if !ok {
		panic("File Service is not configured")
	}

	rpcService := micro.NewService(
		micro.Name(uploadServiceConf.Name),
		micro.Address(bootstrap.GetAddress(uploadServiceConf.Host, uploadServiceConf.Port)),
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

	upload.RegisterFileUploadServiceHandler(rpcService.Server(), service.GetUploadServiceRpc())

	rpcService.Init()
	rpcService.Run()

}
