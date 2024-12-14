package main

import (
	"github.com/sanzashi987/nino-work/apps/chat/db/dao"
	"github.com/sanzashi987/nino-work/apps/chat/http"
	"github.com/sanzashi987/nino-work/pkg/bootstrap"
)

func main() {
	bootstraper := bootstrap.CommonBootstrap("chatService")
	dao.ConnectDB()

	webService := bootstraper.InitWebService(http.NewRouter(bootstraper.Config.System.LoginPage))

	webService.Init()
	webService.Run()
}
