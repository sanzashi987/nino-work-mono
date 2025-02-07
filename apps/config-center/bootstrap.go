package main

import (
	"fmt"

	"github.com/sanzashi987/nino-work/apps/config-center/db"
	"github.com/sanzashi987/nino-work/apps/config-center/http"
	"github.com/sanzashi987/nino-work/pkg/bootstrap"
)

func main() {
	runAsIndependentService()
}

func runAsMicroService() {
	bootstraper := bootstrap.CommonBootstrap("tcc.nino.work")
	db.ConnectDB(bootstraper.Config.System.DbName)

	rpcService := bootstraper.InitRpcService()

	webService := bootstraper.InitWebService(http.NewRouter(bootstraper.Config.System.LoginPage))

	// user.RegisterUserServiceHandler(rpcService.Server(), service.GetUserServiceRpc())
	webService.Init()
	rpcService.Init()

	go func() {
		webService.Run()
	}()

	rpcService.Run()
}

func runAsIndependentService() {
	conf, psmConf := bootstrap.ParseConfig("tcc.nino.work")
	db.ConnectDB(psmConf.DbName)

	router := http.NewRouter(conf.System.LoginPage)
	router.Run(fmt.Sprintf(":%s", psmConf.WebPort))
}
