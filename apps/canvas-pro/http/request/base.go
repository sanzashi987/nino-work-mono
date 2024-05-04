package request

type PaginationRequest struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type BaseResponse[T any] struct {
	Data          T      `json:"data"`
	ResultCode    int    `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
}
