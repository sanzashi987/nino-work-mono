package model

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/sanzashi987/nino-work/pkg/db"
)

// JSONMap 用于存储动态JSON数据
type JSONMap map[string]interface{}

// Value 实现driver.Valuer接口
func (j JSONMap) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan 实现sql.Scanner接口
func (j *JSONMap) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, &j)
}

type ConfigModel struct {
	db.BaseModel
	NamespaceID uint64      `gorm:"column:namespace_id;index"`
	Key         string      `gorm:"column:key;type:varchar(255)"`
	Value       JSONMap     `gorm:"column:value;type:json"`
	Version     int64       `gorm:"column:version"`
	Description string      `gorm:"column:description"`
	Status      int         `gorm:"column:status;default:0"` // 0:启用 1:禁用
}

func (c ConfigModel) TableName() string {
	return "configs"
} 