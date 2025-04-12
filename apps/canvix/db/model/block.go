package model

type BlockModel struct {
	BaseModel
	Version string
	Config  string `gorm:"type:blob"`
	GroupId uint64 `gorm:"default:0"`
}

func (m BlockModel) TableName() string {
	return "blocks"
}
