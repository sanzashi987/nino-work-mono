package main

import (
	"github.com/cza14h/nino-work/apps/canvas-pro/db/dao"
	"github.com/cza14h/nino-work/apps/canvas-pro/http"
	"github.com/cza14h/nino-work/pkg/bootstrap"
	"github.com/cza14h/nino-work/proto/upload"
	"go-micro.dev/v4/web"
)

func main() {
	conf, etcdRegistry := bootstrap.CommonBootstrap("canvasService.client")
	dao.ConnectDB()

	client := bootstrap.GetClient()
	client.GetClientInstance().Init()
	uploadClient := upload.NewFileUploadService(client.GetName(), client.GetClientInstance())

	canvasService, ok := conf.Service["canvasService"]
	if !ok {
		panic("Canvas Service is not configured")
	}

	webService := web.NewService(
		web.Name(canvasService.Name+".web"),
		web.Address(bootstrap.GetAddress(canvasService.Host, canvasService.Port)),
		web.Handler(http.NewRouter(conf.System.LoginPage, map[string]any{
			"upload": uploadClient,
		})),
		web.Registry(etcdRegistry),
	)

	webService.Init()
	webService.Run()
}
