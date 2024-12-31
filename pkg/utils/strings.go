package utils

import "strings"

func Capitialize(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}
