package shared

type PaginationRequest struct {
	Page int `json:"page" form:"page" binding:"required"`
	Size int `json:"size" form:"size" binding:"required"`
}

type PaginationResponse struct {
	PageIndex   int `json:"page"`
	PageSize    int `json:"size"`
	RecordTotal int `json:"total"`
}

type DBTime struct {
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

type DBTimestamp struct {
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
}

type EnumMeta struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}
