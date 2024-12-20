package main

import (
	"fmt"

	"github.com/sanzashi987/nino-work/apps/user/db/dao"
	"github.com/sanzashi987/nino-work/apps/user/http"
	"github.com/sanzashi987/nino-work/apps/user/service"
	"github.com/sanzashi987/nino-work/pkg/bootstrap"
	"github.com/sanzashi987/nino-work/proto/user"
)

func main() {
	runAsIndependentService()
}

func runAsMicroService() {
	bootstraper := bootstrap.CommonBootstrap("sso.nino.work")
	dao.ConnectDB(bootstraper.Config.System.DbName)

	rpcService := bootstraper.InitRpcService()

	webService := bootstraper.InitWebService(http.NewRouter(bootstraper.Config.System.LoginPage))

	user.RegisterUserServiceHandler(rpcService.Server(), service.GetUserServiceRpc())
	webService.Init()
	rpcService.Init()

	go func() {
		webService.Run()
	}()

	rpcService.Run()
}

func runAsIndependentService() {
	conf, psmConf := bootstrap.ParseConfig("sso.nino.work")
	dao.ConnectDB(psmConf.DbName)

	router := http.NewRouter(conf.System.LoginPage)
	router.Run(fmt.Sprintf(":%s", psmConf.WebPort))
}
