package request

type BasicPageResponse struct {
	Id       string
	Code     string
	Group    string `json:"groupCode"`
	CreateAt string `json:"createTime"`
	UpdateAt string `json:"updateTime"`
}
