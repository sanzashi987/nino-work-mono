package dao

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/chat/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type ChatDao struct {
	db.BaseDao[model.MessageModel]
}

func NewChatDao(ctx context.Context, dao ...*db.BaseDao[model.MessageModel]) *ChatDao {
	return &ChatDao{BaseDao: db.NewDao[model.MessageModel](ctx, dao...)}
}

func (c *ChatDao) CreateMessagePair(gptAnswer, userInput string, dialogID uint64) (uint64, uint64, error) {
	userMessage := model.MessageModel{
		DialogID: dialogID,
		Content:  userInput,
	}

	if err := c.GetOrm().Create(&userMessage).Error; err != nil {
		return 0, 0, err
	}

	gptAnwserMessage := model.MessageModel{
		ReplyTo:  userMessage.Id,
		DialogID: dialogID,
		Content:  gptAnswer,
	}

	if err := c.GetOrm().Create(&gptAnwserMessage).Error; err != nil {
		return 0, 0, err
	}

	return userMessage.Id, gptAnwserMessage.Id, nil
}
