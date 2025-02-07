package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
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
	method string,
	path string,
	query url.Values,
	contentType string,
	timestamp string,
	secretKey string,
) string {
	// 1. 规范化参数
	params := make(url.Values)
	for k, v := range query {
		params[k] = v
	}

	// 2. 参数排序
	var paramKeys []string
	for k := range params {
		paramKeys = append(paramKeys, k)
	}
	sort.Strings(paramKeys)

	// 3. 构建待签名字符串
	var paramParts []string
	for _, k := range paramKeys {
		paramParts = append(paramParts,
			fmt.Sprintf("%s=%s", url.QueryEscape(k), url.QueryEscape(params.Get(k))))
	}

	payload := strings.Join([]string{
		method,
		path,
		strings.Join(paramParts, "&"),
		contentType,
		timestamp,
	}, "\n")

	// 4. 计算HMAC-SHA256
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(payload))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return signature
}
