package model

type ConfigMetaJson struct{}

type TemplateModel struct {
	BaseModel
	Thumbnail string `json:"thumbnail"`
	Version   string
	Config    string `gorm:"type:blob"`
	GroupId   uint64
}

func (m TemplateModel) TableName() string {
	return "templates"
}
