package model

type TemplateModel struct {
	ProjectModel
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
