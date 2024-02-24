package model

import (
	"errors"
	"time"

	"github.com/cza14h/nino-work/apps/canvas-pro/utils"
	"github.com/cza14h/nino-work/pkg/db"
	"gorm.io/gorm"
)

const (
	NotDeleted = 0
	Deleted    = 1
)

type BaseModel struct {
	db.BaseModel
	Name      string
	Workspace string
	Creator   string
	Deleted   uint8 `gorm:"deleted:tinyint(8)"`
}

var ErrorNegativeSnowflakeId = errors.New("a negative id is generated")

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	tempId := utils.GenerateId()
	if tempId < 0 {
		err = ErrorNegativeSnowflakeId
		return
	}
	b.Id = uint64(tempId)
	b.CreateTime = time.Now()
	return
}
