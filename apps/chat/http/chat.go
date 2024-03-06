package http

import (
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
		c.AbortClientError(ctx, "[http] chat: fail to get required field, "+err.Error())
		return
	}
	rpcInstance := service.GetChatServiceRpc()

	if err := rpcInstance.Chat(ctx, &req, &res); err != nil {
		c.AbortJson(ctx, int(res.Reason), "[rpc] chat service: fail to get required field, "+err.Error())
		return
	}
	c.ResponseJson(ctx, &res)
}
