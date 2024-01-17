package response

type BaseResponse[T any] struct {
	Data          T      `json:"data"`
	ResultCode    int    `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
}
