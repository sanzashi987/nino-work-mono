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
		c.AbortJson(ctx, http.StatusBadRequest, "Fail to get required field")
		return
	}
	rpcInstance := service.GetChatServiceRpc()

	if err := rpcInstance.Chat(ctx, &req, &res); err != nil {
		c.AbortJson(ctx, int(res.Reason), "Fail to get required field")
		return
	}
	c.ResponseJson(ctx, &res)
}
