package model

const ProjectUnPublished = 0
const ProjectPublished = 1

type ProjectGroupModel struct {
	BaseModel
	// Projects []ProjectModel `gorm:"foreignkey:GroupId;references:Id"`
}

func (p ProjectGroupModel) TableName() string {
	return "project_groups"
}

type ProjectSettingsJson struct {
}

type ProjectModel struct {
	TemplateModel
	GroupId uint64
	Publish int

	/** In type of `ProjectSettingsJson`*/
	// Settings string `gorm:"type:blob"`
}

func (p ProjectModel) TableName() string {
	return "projects"
}
