package model

import (
	"errors"
	"time"

	"github.com/cza14h/nino-work/apps/canvas-pro/consts"
	"github.com/cza14h/nino-work/apps/canvas-pro/utils"
	"github.com/cza14h/nino-work/pkg/db"
	"github.com/cza14h/nino-work/pkg/filter"
	"gorm.io/gorm"
)

type BaseModel struct {
	db.BaseModel
	TypeTag   string `gorm:"-"`
	Name      string
	Workspace uint64 `gorm:"default:0;index"`
	Creator   uint64
	Code      string
}

var ErrorNegativeSnowflakeId = errors.New("a negative id is generated")
var ErrorTypeTagIsNotSupported = errors.New("canvas typeTag is not supported")

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
	return filter.Filter(records, func(e T) bool {
		return e.GetDeleteTime() == nil
	})
}
