package enums

import "github.com/Masterminds/semver/v3"

var DefaultVersion = semver.New(0, 1, 0, "", "").String()
