package model

import (
	"errors"
	"time"

	"github.com/sanzashi987/nino-work/apps/canvas-pro/consts"
	"github.com/sanzashi987/nino-work/apps/canvas-pro/utils"
	"github.com/sanzashi987/nino-work/pkg/db"
	"github.com/sanzashi987/nino-work/pkg/utils"
	"gorm.io/gorm"
)

type BaseModel struct {
	db.BaseModel
	TypeTag   string `gorm:"index"`
	Name      string
	Workspace uint64 `gorm:"default:0;index"`
	Creator   uint64
	Code      string `gorm:"index"`
}

var ErrorNegativeSnowflakeId = errors.New("a negative id is generated")
var ErrorTypeTagIsNotSupported = errors.New("canvas typeTag is not supported")

type GetTypeTag interface {
	GetTypeTag() string
}

func (b BaseModel) GetTypeTag() string {
	return b.TypeTag
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	tempId := utils.GenerateId()

	if consts.IsSupportedTypeTag(b.TypeTag) {
		err = ErrorTypeTagIsNotSupported
		return
	}

	if tempId < 0 {
		err = ErrorNegativeSnowflakeId
		return
	}
	b.Id = uint64(tempId)
	b.Code = consts.GetCodeFromId(b.TypeTag, b.Id)
	b.CreateTime = time.Now()
	return
}

func (b *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	b.UpdateTime = time.Now()
	return
}

func FilterRecordsInUse[T db.GetDeleteTime](records []T) []T {
	return utils.Filter(records, func(e T) bool {
		return e.GetDeleteTime() == nil
	})
}

func FilterRecordsByTypeTag[T GetTypeTag](records []T, typeTag string) []T {
	return utils.Filter(records, func(e T) bool {
		return e.GetTypeTag() == typeTag
	})
}
