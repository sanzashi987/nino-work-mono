package main

import (
	"github.com/sanzashi987/nino-work/apps/canvas-pro/db/dao"
	"github.com/sanzashi987/nino-work/apps/canvas-pro/http"
	"github.com/sanzashi987/nino-work/pkg/bootstrap"
	"github.com/sanzashi987/nino-work/proto/storage"
	"go-micro.dev/v4/web"
)

func main() {
	conf, etcdRegistry := bootstrap.CommonBootstrap("canvasService")
	dao.ConnectDB()

	client := bootstrap.InitClient("canvasService.client", etcdRegistry)

	// client := bootstrap.GetClient()
	// client.GetClientInstance().Init()
	uploadClient := storage.NewStorageService(client.Name(), client.Client())

	canvasService, ok := conf.Service["canvasService"]
	if !ok {
		panic("Canvas Service is not configured")
	}

	webService := web.NewService(
		web.Name(canvasService.Name+".web"),
		web.Address(bootstrap.GetAddress(canvasService.Host, canvasService.Port)),
		web.Handler(http.NewRouter(conf.System.LoginPage, map[string]any{
			"storage": uploadClient,
		})),
		web.Registry(etcdRegistry),
	)

	webService.Init()
	webService.Run()
}
