package dao

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/chat/db/model"
	"gorm.io/gorm"
)

type ChatDao struct {
	*gorm.DB
}

func NewChatDao(ctx context.Context) *ChatDao {
	return &ChatDao{DB: newDBSession(ctx)}
}

func (c *ChatDao) CreateMessagePair(gptAnswer, userInput string, dialogID uint64) (uint64, uint64, error) {
	userMessage := model.MessageModel{
		DialogID: dialogID,
		Content:  userInput,
	}

	if err := c.Create(&userMessage).Error; err != nil {
		return 0, 0, err
	}

	gptAnwserMessage := model.MessageModel{
		ReplyTo:  userMessage.Id,
		DialogID: dialogID,
		Content:  gptAnswer,
	}

	if err := c.Create(&gptAnwserMessage).Error; err != nil {
		return 0, 0, err
	}

	return userMessage.Id, gptAnwserMessage.Id, nil
}
