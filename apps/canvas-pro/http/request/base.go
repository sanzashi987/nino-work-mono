package request

type PaginationRequest struct {
	Page int `json:"page"`
	Size int `json:"size"`
}
