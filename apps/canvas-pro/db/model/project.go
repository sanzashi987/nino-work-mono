package model

type ProjectGroup struct {
	BaseModel
}

func (p ProjectGroup) TableName() string {
	return "project_groups"
}

type ProjectModel struct {
	ProjectGroup
	Code    string
	Version string
	Config  string `gorm:"type:blob"`
}

func (p ProjectModel) TableName() string {
	return "projects"
}
