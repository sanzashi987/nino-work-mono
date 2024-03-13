package model

const (
	ProjectPublished    = 1
	ProjectUnPublished  = 0
	PojectPublishNormal = "1"
	PojectPublishToken  = "2"
	PojectPublishSecret = "3"
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
