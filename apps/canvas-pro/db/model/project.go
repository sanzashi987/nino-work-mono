package model

type ProjectGroupModel struct {
	BaseModel
	Projects []ProjectModel
}

func (p ProjectGroupModel) TableName() string {
	return "project_groups"
}

type ProjectSettingsJson struct {
}

type ProjectModel struct {
	TemplateModel
	/** In type of `ProjectSettingsJson`*/
	// Settings string `gorm:"type:blob"`
}

func (p ProjectModel) TableName() string {
	return "projects"
}
