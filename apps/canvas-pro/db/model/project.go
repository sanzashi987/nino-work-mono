package model

type ProjectGroupModel struct {
	BaseModel
	Projects []ProjectModel
}

func (p ProjectGroupModel) TableName() string {
	return "project_groups"
}

type ProjectSettingsJson struct {
	Name      string `json:"name" binding:"required"`
	Thumbnail string `json:"thumbnail"`
}

type ProjectModel struct {
	TemplateModel
	Settings string `gorm:"type:blob"`
}

func (p ProjectModel) TableName() string {
	return "projects"
}
