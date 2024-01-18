package model

type CanvasUserModel struct {
	ID         uint64           `gorm:"column:id;primaryKey;index:,unique"`
	Workspaces []WorkSpaceModel `gorm:"many2many:workspace_user;foreignKey:ID;References:Code"`
	Permission uint64
}

func (u CanvasUserModel) TableName() string {
	return "canvas_users"
}
