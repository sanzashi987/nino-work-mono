package model

type ThemeModel struct {
	BaseModel
	Type   int8   `gorm:"column:type"` //0 is built-in
	Config string `gorm:"type:blob"`
}

func (theme ThemeModel) TableName() string {
	return "themes"
}
