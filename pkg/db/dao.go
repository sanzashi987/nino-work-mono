package db

import (
	"context"

	"github.com/glebarez/sqlite"
	"github.com/sanzashi987/nino-work/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var instance *gorm.DB

type BaseDao[Model any] struct {
	db          *gorm.DB
	transaction *gorm.DB
	ctx         context.Context
}

func ConnectDB(names ...string) *gorm.DB {
	fallbackName := ""
	if len(names) >= 1 {
		fallbackName = names[0]
	} else {
		conf := config.GetConfig()
		fallbackName = conf.System.DbName
	}

	db, err := gorm.Open(sqlite.Open(fallbackName+".db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Fail to connect database")
	}
	instance = db
	return db
}

func NewDBSession(ctx context.Context) *gorm.DB {
	return instance.WithContext(ctx)
}

func InitBaseDao[Model any](ctx context.Context) BaseDao[Model] {
	return BaseDao[Model]{
		db: NewDBSession(ctx),
		ctx:ctx,
	}
}

func Paginate(page, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageNumber, pageSize := page, size
		if pageNumber == 0 {
			pageNumber = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
