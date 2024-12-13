package main

import (
	"github.com/sanzashi987/nino-work/apps/chat/db/dao"
	"github.com/sanzashi987/nino-work/apps/chat/http"
	"github.com/sanzashi987/nino-work/pkg/bootstrap"
	"go-micro.dev/v4/web"
)

func main() {
	conf, etcdRegistry := bootstrap.CommonBootstrap("chatService")
	dao.ConnectDB()

	chatService, ok := conf.Service["chatService"]
	if !ok {
		panic("ChatService is not configured")
	}

	webService := web.NewService(
		web.Name(chatService.Name+".web"),
		web.Address(bootstrap.GetAddress(chatService.Host, chatService.Port)),
		web.Handler(http.NewRouter(conf.System.LoginPage)),
		web.Registry(etcdRegistry),
	)

	webService.Init()
	webService.Run()
}
