package model

type WorkspaceModel struct {
	BaseModel
	Default    int8
	Owner      uint64
	Capacity   uint64
	Permission uint64
	Members    []*CanvixUserModel `gorm:"many2many:canvix_workspace_user;foreignKey:Id;joinForeginKey:WorkspaceId;joinReferences:CanvasUserId;References:Id"`
}

// func (s WorkspaceModel) TableName() string {
// 	return "workspaces"
// }
