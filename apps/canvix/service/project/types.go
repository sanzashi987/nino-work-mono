package project

import "github.com/sanzashi987/nino-work/pkg/shared"

type ProjectDetail struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	shared.DBTime
}

type ProjectInfo struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	shared.DBTimestamp
}
