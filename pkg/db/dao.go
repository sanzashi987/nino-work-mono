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

func NewTx(ctx context.Context) *gorm.DB {
	return instance.WithContext(ctx)
}

func Paginate(page, size, total int) func(db *gorm.DB) *gorm.DB {
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

		rest := total % size
		maxPage := total / size
		if rest > 0 {
			maxPage += 1
		}
		if page > maxPage {
			page = maxPage
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize).Order("update_time DESC")
	}
}

func calibratePage(page, size, total int) int {
	rest := total % size
	maxPage := total / size
	if rest > 0 {
		maxPage += 1
	}

	if page > maxPage {
		return maxPage
	}
	return page

}

type ListResponse[T any] struct {
	Total   int
	Records []*T
	Page    int
}

func QueryWithTotal[T any](condition *gorm.DB, page, size int) (*ListResponse[T], error) {
	var total *int64

	if err := condition.Select("id").Count(total).Error; err != nil {
		return nil, err
	}
	p := calibratePage(page, size, int(*total))

	records := []*T{}
	if err := condition.Scopes(Paginate(p, size, int(*total))).Find(&records).Error; err != nil {
		return nil, err
	}

	res := ListResponse[T]{
		Total:   int(*total),
		Records: records,
		Page:    p,
	}

	return &res, nil
}
