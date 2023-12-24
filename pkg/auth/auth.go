package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var signer = []byte("signer.nino.work")

type AuthClaims struct {
	UserID   uint
	Username string
	jwt.RegisteredClaims
}

func GenerateToken(username string, userId uint, expires ...time.Duration) (string, error) {
	now := time.Now()
	var expiry time.Duration
	if len(expires) == 0 {
		expiry = 24 * time.Hour
	} else {
		expiry = expires[0]
	}
	claims := AuthClaims{
		UserID:   userId,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add((expiry))),
			Issuer:    "nino.work",
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(signer)

}

func ValidateToken(inputToken string) (*AuthClaims, error) {

	token, err := jwt.ParseWithClaims(inputToken, &AuthClaims{}, func(t *jwt.Token) (interface{}, error) {
		return signer, nil
	})

	if token != nil {
		if claim, ok := token.Claims.(*AuthClaims); ok {
			return claim, nil
		}
	}

	return nil, err

}
