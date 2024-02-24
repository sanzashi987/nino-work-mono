package model

// const
const DESIGN = 1

// const FILE = 2
const DATASOURCE = 2
const FONT = 3
const BLOCK = 4
const COMPONENT = 5
const TEMPLATE = 6

type AssetModel struct {
	BaseModel
	Version  string
	Type     uint8
	FileId   string
	FilePath string
}

func (m AssetModel) TableName() string {
	return "assets"
}

type AssetGroupModel struct {
	ProjectGroupModel
}

func (m AssetGroupModel) TableName() string {
	return "asset_groups"
}
