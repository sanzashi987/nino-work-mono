package request

type PaginationRequest struct {
	Page int `json:"page" form:"page" binding:"required"`
	Size int `json:"size" form:"size" binding:"required"`
}

type PaginationResponse struct {
	PageIndex   int `json:"pageIndex"`
	PageSize    int `json:"pageSize"`
	PageTotal   int `json:"pageTotal"`
	RecordTotal int `json:"recordTotal"`
}

type BaseResponse[T any] struct {
	Data          T      `json:"data"`
	ResultCode    int    `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
}

type DBTime struct {
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}
