package main

import (
	"github.com/cza14h/nino-work/apps/chat/db/dao"
	"github.com/cza14h/nino-work/pkg/bootstrap"
	"go-micro.dev/v4/web"
)

func main() {
	conf, etcdRegistry := bootstrap.CommonBootstrap("chatService.client")
	dao.ConnectDB()

	chatService, ok := conf.Service["chatService"]
	if !ok {
		panic("ChatService is not configured")
	}

	webService := web.NewService(
		web.Name(chatService.Name+".web"),
		web.Address(bootstrap.GetAddress(chatService.Host, chatService.Port)),
		web.Registry(etcdRegistry),
	)

	webService.Init()
	webService.Run()
}
