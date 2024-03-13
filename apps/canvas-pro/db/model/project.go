package model

const (
	ProjectPublished = 1

	ProjectUnPublished  = 0
	PojectPublishNormal = "0"
	PojectPublishToken  = "1"
	PojectPublishSecret = "2"
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
	GroupId     uint64
	Publish     int
	Token       string
	Secret      string
	PublishFlag string
	/** In type of `ProjectSettingsJson`*/
	// Settings string `gorm:"type:blob"`
}

func (p ProjectModel) TableName() string {
	return "projects"
}
