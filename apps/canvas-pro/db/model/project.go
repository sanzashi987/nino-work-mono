package model

type ProjectGroupModel struct {
	BaseModel
	Projects []ProjectModel
}

func (p ProjectGroupModel) TableName() string {
	return "project_groups"
}

type ProjectModel struct {
	BaseModel
	Code    string
	Version string
	Config  string `gorm:"type:blob"`
}

func (p ProjectModel) TableName() string {
	return "projects"
}
