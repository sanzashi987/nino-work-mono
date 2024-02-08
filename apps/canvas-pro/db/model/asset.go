package model

// const
const DESGIN = 1

// const FILE = 2
const DATASOURCE = 2
const FONT = 3
const BLOCK = 4
const COMPONENT = 5

type AssetModel struct {
	ProjectModel
	Type     uint8
	FileId   string
	FilePath string
}

func (m AssetModel) TableName() string {
	return "assets"
}

type AssetGroupModel struct {
	ProjectGroup
}

func (m AssetGroupModel) TableName() string {
	return "asset_groups"
}
