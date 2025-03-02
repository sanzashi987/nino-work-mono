package model

// const TEMPLATE = 6

type AssetModel struct {
	BaseModel
	Version string
	FileId  string
	// FilePath string
	GroupId     uint64 `gorm:"default:0"`
	AssetConfig string `gorm:"type:blob"`
}

func (m AssetModel) TableName() string {
	return "assets"
}
