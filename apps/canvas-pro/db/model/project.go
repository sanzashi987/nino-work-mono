package model

const (
	ProjectPublishFlagNotPublish = 0
	ProjectPublishFlagNormal     = 1
	ProjectPublishFlagToken      = 2
	ProjectPublishFlagSecret     = 3
)

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
	PublishToken  string
	PublishSecret string
	PublishFlag   int8
	/** In type of `ProjectSettingsJson`*/
	// Settings string `gorm:"type:blob"`
}

func (p ProjectModel) TableName() string {
	return "projects"
}
