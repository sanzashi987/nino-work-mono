package main

import (
	"fmt"

	"github.com/sanzashi987/nino-work/apps/canvix/db/dao"
	"github.com/sanzashi987/nino-work/apps/canvix/http"
	"github.com/sanzashi987/nino-work/pkg/bootstrap"
	"github.com/sanzashi987/nino-work/proto/storage"
)

func main() {
	runAsIndependentService()
}

func runAsMicroService() {
	bootstraper := bootstrap.CommonBootstrap("canvix.nino.work")
	dao.ConnectDB()

	client := bootstraper.InitRpcClient()

	uploadClient := storage.NewStorageService(client.Name, client.Client)

	webService := bootstraper.InitWebService(http.NewRouter(bootstraper.Config.System.LoginPage, map[string]any{
		"storage": uploadClient,
		// "storage": nil,
	}))

	webService.Init()
	webService.Run()
}

func runAsIndependentService() {
	conf, psmConf := bootstrap.ParseConfig("canvix.nino.work")
	dao.ConnectDB(psmConf.DbName)

	router := http.NewRouter(conf.System.LoginPage, map[string]any{"storage": nil})
	router.Run(fmt.Sprintf(":%s", psmConf.WebPort))
}
