package http

import (
	"net/http"

	"github.com/cza14h/nino-work/apps/chat/service"
	"github.com/cza14h/nino-work/pkg/controller"
	"github.com/cza14h/nino-work/proto/chat"
	"github.com/gin-gonic/gin"
)

type ChatController struct {
	controller.BaseController
}

func (c *ChatController) Chat(ctx *gin.Context) {
	req := chat.ChatRequest{}
	res := chat.ChatResponse{}
	if err := ctx.BindJSON(&req); err != nil {
		c.ResponseJson(ctx, http.StatusBadRequest, "Fail to get required field", nil)
		ctx.Abort()
		return
	}
	rpcInstance := service.GetChatServiceRpc()

	if err := rpcInstance.Chat(ctx, &req, &res); err != nil {
		ctx.Abort()
		c.ResponseJson(ctx, int(res.Reason), "Error from chat instance", nil)
		return
	}
	c.ResponseJson(ctx, int(res.Reason), "", &res)
}
