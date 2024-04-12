package service

import (
	"context"
	"encoding/json"
	"errors"

	// "sync"

	"github.com/cza14h/nino-work/apps/chat/consts"
	"github.com/cza14h/nino-work/apps/chat/db/dao"
	"github.com/cza14h/nino-work/proto/chat"
	"github.com/sashabaranov/go-openai"
)

type ChatServiceRpcImpl struct{}

// var once sync.Once
var chatService = &ChatServiceRpcImpl{}

func GetChatServiceRpc() chat.ChatServiceHandler {
	// once.Do(func() {
	// 	chatService = &ChatServiceRpcImpl{}
	// })
	return chatService
}

func (c *ChatServiceRpcImpl) Chat(ctx context.Context, in *chat.ChatRequest, out *chat.ChatResponse) (err error) {
	dbSession := dao.NewChatDao(ctx)
	var gptRequest openai.ChatCompletionRequest
	messages := []openai.ChatCompletionMessage{}
	for _, history := range in.History {
		convert := openai.ChatCompletionMessage{
			Name:    history.Name,
			Role:    history.Role,
			Content: history.Content,
		}
		messages = append(messages, convert)
	}

	json.Unmarshal([]byte(in.Preference), &gptRequest)

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    "user",
		Content: in.Content,
	})

	gptRequest.Messages = messages
	gptRequest.Stream = false

	gptConfig := openai.DefaultConfig("")
	client := openai.NewClientWithConfig(gptConfig)

	if ok := consts.SupportModels[gptRequest.Model]; ok {
		var response openai.ChatCompletionResponse
		response, err = client.CreateChatCompletion(ctx, gptRequest)
		if err != nil {
			out.Reason = consts.FailToCreateCompletion
			return
		}

		content := response.Choices[0].Message.Content
		var userMessageId uint64
		userMessageId, _, err = dbSession.CreateMessagePair(content, in.Content, uint64(in.DialogId))
		if err != nil {
			out.Reason = consts.FailToInsertMessagePair
			return
		}
		out.Content = content
		out.Reason = consts.Success
		out.Id = userMessageId
		return
	}

	err = errors.New("Unknown edge case in chat service")
	return
}
