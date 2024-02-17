package model

type ProjectGroupModel struct {
	BaseModel
}

func (p ProjectGroupModel) TableName() string {
	return "project_groups"
}

type ProjectModel struct {
	ProjectGroupModel
	Name    string
	Code    string
	Version string
	Config  string `gorm:"type:blob"`
}

func (p ProjectModel) TableName() string {
	return "projects"
}
