package service

import (
	"context"
	"sync"

	"github.com/cza14h/nino-work/proto/chat"
)

type ChatServiceRpcImpl struct{}

var once sync.Once
var chatService *ChatServiceRpcImpl

func GetChatServiceRpc() *ChatServiceRpcImpl {
	once.Do(func() {
		chatService = &ChatServiceRpcImpl{}
	})
	return chatService
}

func Chat(ctx context.Context, in *chat.ChatRequest, out *chat.ChatResponse) {

}
