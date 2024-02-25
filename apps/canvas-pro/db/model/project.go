package model

type ProjectGroupModel struct {
	BaseModel
	Projects []ProjectModel
}

func (p ProjectGroupModel) TableName() string {
	return "project_groups"
}

type SystemConfigJson struct{}

type ProjectModel struct {
	TemplateModel
	SystemConfig string `gorm:"type:blob"`
}

func (p ProjectModel) TableName() string {
	return "projects"
}
