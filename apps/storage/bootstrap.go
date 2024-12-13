package main

import (
	"github.com/sanzashi987/nino-work/apps/storage/db"
	"github.com/sanzashi987/nino-work/apps/storage/service"
	"github.com/sanzashi987/nino-work/pkg/bootstrap"
	"github.com/sanzashi987/nino-work/proto/storage"
	"go-micro.dev/v4"
)

func main() {

	conf, etcdRegistry := bootstrap.CommonBootstrap("storageService")

	storageServiceConf, ok := conf.Service["storageService"]
	if !ok {
		panic("File Service is not configured")
	}

	db.ConnectDB()
	rpcService := micro.NewService(
		micro.Name(storageServiceConf.Name),
		micro.Address(bootstrap.GetAddress(storageServiceConf.Host, storageServiceConf.Port)),
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

	storage.RegisterStorageServiceHandler(rpcService.Server(), service.GetUploadServiceRpc())

	rpcService.Init()
	rpcService.Run()

}
