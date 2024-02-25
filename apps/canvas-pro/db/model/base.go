package model

import (
	"errors"
	"time"

	"github.com/cza14h/nino-work/apps/canvas-pro/enums"
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
	TypeTag   string `gorm:"-"`
	Name      string
	Workspace string
	Creator   string
	Code      string
	Deleted   uint8 `gorm:"deleted:tinyint(8)"`
}

var ErrorNegativeSnowflakeId = errors.New("a negative id is generated")
var ErrorTypeTagIsNotSupported = errors.New("canvas typeTag is not supported")

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	tempId := utils.GenerateId()

	if enums.IsSupportedTypeTag(b.TypeTag) {
		err = ErrorTypeTagIsNotSupported
		return
	}

	if tempId < 0 {
		err = ErrorNegativeSnowflakeId
		return
	}
	b.Id = uint64(tempId)
	b.Code = enums.GetCodeFromId(b.TypeTag, b.Id)
	b.CreateTime = time.Now()
	return
}
