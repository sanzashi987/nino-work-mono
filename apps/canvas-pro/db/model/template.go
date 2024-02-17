package model

type TemplateModel struct {
	AssetModel
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
