package main

import (
	"github.com/sanzashi987/nino-work/apps/user/db/dao"
	"github.com/sanzashi987/nino-work/apps/user/http"
	"github.com/sanzashi987/nino-work/apps/user/service"
	"github.com/sanzashi987/nino-work/pkg/bootstrap"
	"github.com/sanzashi987/nino-work/proto/user"
)

func main() {
	bootstraper := bootstrap.CommonBootstrap("userService")
	dao.ConnectDB()

	rpcService := bootstraper.InitRpcService()

	webService := bootstraper.InitWebService(http.NewRouter())

	user.RegisterUserServiceHandler(rpcService.Server(), service.GetUserServiceRpc())
	webService.Init()
	rpcService.Init()

	go func() {
		webService.Run()
	}()

	rpcService.Run()

}
