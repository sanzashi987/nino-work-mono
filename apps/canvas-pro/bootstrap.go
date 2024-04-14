package main

import (
	"github.com/cza14h/nino-work/apps/canvas-pro/db/dao"
	"github.com/cza14h/nino-work/apps/canvas-pro/http"
	"github.com/cza14h/nino-work/pkg/bootstrap"
	"go-micro.dev/v4/web"
)

func main() {
	conf, etcdRegistry := bootstrap.CommonBootstrap("canvasService.client")
	dao.ConnectDB()

	canvasService, ok := conf.Service["canvasService"]
	if !ok {
		panic("Canvas Service is not configured")
	}

	webService := web.NewService(
		web.Name(canvasService.Name+".web"),
		web.Address(bootstrap.GetAddress(canvasService.Host, canvasService.Port)),
		web.Handler(http.NewRouter(conf.System.LoginPage)),
		web.Registry(etcdRegistry),
	)

	webService.Init()
	webService.Run()
}
