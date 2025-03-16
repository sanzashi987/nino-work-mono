package model

type GroupModel struct {
	BaseModel
}

func (p GroupModel) TableName() string {
	return "groups"
}
