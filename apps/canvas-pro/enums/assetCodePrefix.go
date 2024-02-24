package enums

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cza14h/nino-work/apps/canvas-pro/utils"
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
)

const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const length = len(charset)

var runeToIndex = map[rune]int{}

func init() {
	for i, c := range charset {
		runeToIndex[c] = i
	}
}

func Encode[T ~int64](id T) string {
	var (
		result []byte
		abs    T = id
	)
	if id < 0 {
		abs = -1 * id
	}
	for abs > 0 {
		rem := abs % T(length)
		abs = abs / T(length)
		result = append([]byte{charset[rem]}, result...)
	}
	return string(result)
}

var ErrorDecodeIllegalChar = errors.New("contains an illegal coded char")

func Decode(code string) (result int64, err error) {
	for _, c := range code {
		rem, exist := runeToIndex[c]
		if !exist {
			err = ErrorDecodeIllegalChar
			result = -1
			return
		}
		result = result*int64(length) + int64(rem)
	}
	return
}

func CreateCode(typeTag string) string {
	return fmt.Sprintf("%s%s%s", PREFIX, typeTag, Encode(utils.GenerateId()))
}

var ErrorNotCanvasCode = errors.New("not a canvas code string")

func GetIdFromCode(canvasCode string) (id int64, typeTag string, err error) {
	id, typeTag = int64(-1), ""
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

func GetCodeFromId(typeTag string, id int64) string {
	return fmt.Sprintf("%s%s%s", PREFIX, typeTag, Encode(id))

}
