package model

type ProjectGroupModel struct {
	BaseModel
	Projects []ProjectModel
}

func (p ProjectGroupModel) TableName() string {
	return "project_groups"
}

type SystemConfigJson struct {
	Name      string `json:"name" binding:"required"`
	Thumbnail string `json:"thumbnail" binding:"required"`
}

type ProjectModel struct {
	TemplateModel
	SystemConfig string `gorm:"type:blob"`
}

func (p ProjectModel) TableName() string {
	return "projects"
}
