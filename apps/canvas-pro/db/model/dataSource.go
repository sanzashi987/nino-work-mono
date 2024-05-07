package model

const (
	STATIC = 0
	API    = 1
	FILE   = 2
)

type DataSourceModel struct {
	BaseModel
	Version      string
	SourceType   uint8  `gorm:"index"`
	SourceConfig string `gorm:"type:blob"`
}

func (m DataSourceModel) TableName() string {
	return "data_sources"
}
