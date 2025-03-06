package main

import (
	"github.com/sanzashi987/nino-work/apps/canvix/db/dao"
	"github.com/sanzashi987/nino-work/apps/canvix/http"
	"github.com/sanzashi987/nino-work/pkg/bootstrap"
	"github.com/sanzashi987/nino-work/proto/storage"
)

func main() {
	bootstraper := bootstrap.CommonBootstrap("canvas.nino.work")
	dao.ConnectDB()

	client := bootstraper.InitRpcClient()

	uploadClient := storage.NewStorageService(client.Name, client.Client)

	webService := bootstraper.InitWebService(http.NewRouter(bootstraper.Config.System.LoginPage, map[string]any{
		"storage": uploadClient,
	}))

	webService.Init()
	webService.Run()
}
