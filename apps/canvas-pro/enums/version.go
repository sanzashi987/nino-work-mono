package enums

import (
	"regexp"

	"github.com/Masterminds/semver/v3"
)

var DefaultVersion = semver.New(0, 1, 0, "", "").String()

const legalNameRegex = `^[\u4E00-\u9FA5\uF900-\uFA2D\w][\u4E00-\u9FA5\uF900-\uFA2D\w-_]*[\u4E00-\u9FA5\uF900-\uFA2D\w]*$`

var LegalNameReg, _ = regexp.Compile(legalNameRegex)