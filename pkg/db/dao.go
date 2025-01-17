package db

import (
	"context"

	"github.com/glebarez/sqlite"
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
	fallbackName := "nino-mono"
	if len(names) >= 1 {
		fallbackName = names[0]
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

func NewDao[T any](ctx context.Context, dao ...*BaseDao[T]) BaseDao[T] {
	var baseDao BaseDao[T]
	if len(dao) > 0 {
		baseDao = *dao[0]
	} else {
		baseDao = BaseDao[T]{
			db:  instance.WithContext(ctx),
			ctx: ctx,
		}
	}
	return baseDao
}

func NewTx(ctx context.Context) *gorm.DB {
	return instance.WithContext(ctx)
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
