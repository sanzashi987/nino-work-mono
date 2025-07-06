package model

const (
	ProjectPublishFlagNotPublish = 0
	ProjectPublishFlagNormal     = 1
	ProjectPublishFlagSecret     = 2
	ProjectPublishFlagToken      = 3
)

type ProjectSettingsJson struct {
}

type ProjectModel struct {
	TemplateModel
	Version string
	// PublishToken  string
	PublishStatus int8   `gorm:"column:publish_status;default:0"`
	PublishSecret string `gorm:"column:publish_secret"`
	/** In type of `ProjectSettingsJson`*/
	// Settings string `gorm:"type:blob"`
}

func (p ProjectModel) TableName() string {
	return "projects"
}
