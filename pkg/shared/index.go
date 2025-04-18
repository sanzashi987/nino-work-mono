package shared

type PaginationRequest struct {
	Page int `json:"page" form:"page" binding:"required"`
	Size int `json:"size" form:"size" binding:"required"`
}

func (r *PaginationRequest) CalibratePage(total int) int {
	rest := total % r.Size
	maxPage := total / r.Size
	if rest > 0 {
		return maxPage + 1
	}
	return maxPage
}

// type PaginationResponse struct {
// 	PageIndex   int `json:"page"`
// 	RecordTotal int `json:"total"`
// }

type ResponseWithPagination[T any] struct {
	Data        T   `json:"data"`
	PageIndex   int `json:"page"`
	RecordTotal int `json:"total"`
}

func (r *ResponseWithPagination[T]) Init(data T, pageIndex, recordTotal int) {
	r.Data = data
	r.PageIndex = pageIndex
	r.RecordTotal = recordTotal
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
