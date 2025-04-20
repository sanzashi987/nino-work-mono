package project

import "github.com/sanzashi987/nino-work/pkg/shared"

type ProjectDetail struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	Thumbnail  string `json:"thumbnail"`
	RootConfig string `json:"root_config"`
	shared.DBTime
}

type ProjectInfo struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	shared.DBTimestamp
}
