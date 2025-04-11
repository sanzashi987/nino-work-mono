package consts

import (
	"errors"
	"fmt"
	"strings"
)

const (
	PREFIX        = "cVs"
	PROJECT       = "A"
	BLOCK         = "B"
	DESIGN        = "C"
	FONT          = "D"
	COMPONENT     = "E"
	DATASOURCE    = "F"
	STATIC_SOURCE = "H"
	THEME         = "J"
	WORKSPACE     = "P"
	GROUP         = "R"
)

type CanvixCodeEnum struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

var TagToName = map[string]string{
	PROJECT:    "project",
	BLOCK:      "block",
	DESIGN:     "design",
	FONT:       "font",
	COMPONENT:  "component",
	DATASOURCE: "datasource",
}

var supportedTags = []string{
	PROJECT,
	BLOCK,
	DESIGN,
	FONT,
	COMPONENT,
	DATASOURCE,
	// STATIC_SOURCE,
}

// may useful in the future
var supportedGroupTags = []string{}

const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const length = len(charset)

var runeToIndex = map[rune]int{}
var supportedTagMap = map[string]int{}
var basicToGroup = map[string]string{}
var groupToBasic = map[string]string{}

const groupStandard = rune('R')

func initGroupTypeTagFromBasic(typeTag string, index int) string {
	charList := []rune(typeTag)
	charList[0] += rune(index) + groupStandard
	groupTypeTag := string(charList)
	basicToGroup[typeTag] = groupTypeTag
	groupToBasic[groupTypeTag] = typeTag
	return groupTypeTag
}

func IsGroup(tag string) bool {
	_, exist := groupToBasic[tag]
	return exist
}

func init() {
	for i, typeTag := range charset {
		runeToIndex[typeTag] = i
	}
	basicTagListLength := len(supportedTags)
	for i, typeTag := range supportedTags {
		supportedTagMap[typeTag] = i
		groupTypeTag := initGroupTypeTagFromBasic(typeTag, i)
		supportedTagMap[groupTypeTag] = basicTagListLength + i
		supportedGroupTags = append(supportedGroupTags, groupTypeTag)
	}
}

func IsSupportedTypeTag(tag string) bool {
	_, exist := supportedTagMap[tag]
	return exist
}

func Encode[T ~uint64](id T) string {
	var (
		result []byte
		abs    T = id
	)
	for abs > 0 {
		rem := abs % T(length)
		abs = abs / T(length)
		result = append([]byte{charset[rem]}, result...)
	}
	return string(result)
}

var ErrorDecodeIllegalChar = errors.New("contains an illegal coded char")

func Decode(code string) (result uint64, err error) {
	for _, c := range code {
		rem, exist := runeToIndex[c]
		if !exist {
			err = ErrorDecodeIllegalChar
			result = 0
			return
		}
		result = result*uint64(length) + uint64(rem)
	}
	return
}

var ErrorNotCanvasCode = errors.New("not a canvas code string")

func GetIdFromCode(canvasCode string) (id uint64, typeTag string, err error) {
	id, typeTag = uint64(0), ""
	if !strings.HasPrefix(canvasCode, PREFIX) {
		err = ErrorNotCanvasCode
		return
	}
	decoded, e := Decode(canvasCode[4:])
	if e != nil {
		err = e
		return
	}
	id, typeTag = decoded, string(canvasCode[3])
	return
}

func GetCodeFromId(typeTag string, id uint64) string {
	return fmt.Sprintf("%s%s%s", PREFIX, typeTag, Encode(id))
}

func CodesIntoIds(codes []string) ([]uint64, error) {
	ids := make([]uint64, len(codes))
	for i, code := range codes {
		id, _, err := GetIdFromCode(code)
		if err != nil {
			return nil, err
		}
		ids[i] = id
	}
	return ids, nil
}
