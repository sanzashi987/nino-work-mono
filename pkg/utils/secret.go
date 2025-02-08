package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

// 生成安全密钥（推荐方式）
func GenerateSecureKeys() (accessKey, secretKey string, err error) {
	// 生成AccessKey（示例：AK_前缀 + 24字节随机数）
	akBytes := make([]byte, 18) // 18*8=144 bits
	if _, err := rand.Read(akBytes); err != nil {
		return "", "", err
	}
	accessKey = "AK_" + base64.URLEncoding.EncodeToString(akBytes)

	// 生成SecretKey（32字节随机数）
	skBytes := make([]byte, 32) // 32*8=256 bits
	if _, err := rand.Read(skBytes); err != nil {
		return "", "", err
	}
	secretKey = base64.URLEncoding.EncodeToString(skBytes)

	return accessKey, secretKey, nil
}

func GenerateSignature(
	psm string,
	method string,
	path string,
	contentType string,
	timestamp string,
	secretKey string,
) string {

	payload := strings.Join([]string{
		psm,
		method,
		path,
		contentType,
		timestamp,
	}, "\n")

	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(payload))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return signature
}
