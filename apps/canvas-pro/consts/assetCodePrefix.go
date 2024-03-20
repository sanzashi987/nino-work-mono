package consts

import (
	"errors"
	"fmt"
	"strings"
)

type AssetPrefix struct{}

const (
	PREFIX        = "cVs"
	PROJECT       = "A"
	BLOCK         = "B"
	DESIGN        = "C"
	FONT          = "D"
	COMPONENT     = "E"
	DATASOURCE    = "F"
	STATIC_SOURCE = "H"
	WORKSPACE     = "P"
	GROUP         = "R"
)

var supportedTags = [7]string{
	PROJECT,
	BLOCK,
	DESIGN,
	FONT,
	COMPONENT,
	DATASOURCE,
	STATIC_SOURCE,
}

const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const length = len(charset)

var runeToIndex = map[rune]int{}
var supportedTagMap = map[string]bool{}

func init() {
	for i, c := range charset {
		runeToIndex[c] = i
	}
	for _, c := range supportedTags {
		supportedTagMap[c] = true
	}
}

func IsSupportedTypeTag(tag string) bool {
	val, exist := supportedTagMap[tag]
	return exist && val
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
	decoded, e := Decode(canvasCode)
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
