package model

type ThemeModel struct {
	BaseModel
	Flag   int8   `gorm:"column:flag"` //0 is built-in
	Config string `gorm:"type:blob"`
}

func (theme ThemeModel) TableName() string {
	return "themes"
}
