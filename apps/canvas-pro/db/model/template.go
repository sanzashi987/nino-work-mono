package model

type ConfigMetaJson struct{}

type TemplateModel struct {
	BaseModel
	Version string
	Config  string `gorm:"type:blob"`
}

func (m TemplateModel) TableName() string {
	return "templates"
}

type TemplateGroupModel struct {
	AssetGroupModel
}

func (m TemplateGroupModel) TableName() string {
	return "template_groups"
}
