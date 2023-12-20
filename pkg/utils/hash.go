package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const cost = 14

func MakeHash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return ""
	}

	return string(bytes)
}

func CheckHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(err)

	return err == nil
}

func IsHashed(str string) bool {
	return len(str) == 60
}

func CompareHash(first, second string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(first), []byte(second))
	return err == nil
}
