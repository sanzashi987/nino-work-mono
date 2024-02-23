package enums

import (
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

func Decode(code string) (result int64) {
	for _, c := range code {
		rem, exist := runeToIndex[c]
		if !exist {
			panic("Fail to parse the code")
		}
		result = result*int64(length) + int64(rem)
	}
	return
}

func CreateCode(cat string) string {
	return fmt.Sprintf("%s%s%s", PREFIX, cat, Encode(utils.GenerateId()))
}

func GetIdFromCode(canvasCode string) (int64, string) {
	if !strings.HasPrefix(canvasCode, PREFIX) {
		panic("Not a legal canvas code!")
	}
	

	Decode(canvasCode)
}
