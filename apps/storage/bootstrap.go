package main

import (
	"fmt"

	"github.com/sanzashi987/nino-work/apps/storage/consts"
	"github.com/sanzashi987/nino-work/apps/storage/db"
	"github.com/sanzashi987/nino-work/apps/storage/http"
	"github.com/sanzashi987/nino-work/apps/storage/service"
	"github.com/sanzashi987/nino-work/pkg/bootstrap"
	"github.com/sanzashi987/nino-work/proto/storage"
)

func main() {
	runAsIndependentService()
}

func runAsMicroService() {
	bootstraper := bootstrap.CommonBootstrap("storage.nino.work")

	db.ConnectDB(bootstraper.PsmConf.DbName)
	rpcService := bootstraper.InitRpcService()

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

func runAsIndependentService() {
	conf, psmConf := bootstrap.ParseConfig("storage.nino.work")

	bucketPath, buckPathExsit := psmConf.Raw[consts.BucketPath]
	tmpPath, TmpPathExsit := psmConf.Raw[consts.TmpPath]
	if !TmpPathExsit || !buckPathExsit {
		panic("Tmp Path or Bucket Path is not configured")
	}

	db.ConnectDB(psmConf.DbName)

	router := http.NewRouter(conf.System.LoginPage, bucketPath, tmpPath)
	router.Run(fmt.Sprintf(":%s", psmConf.WebPort))
}
