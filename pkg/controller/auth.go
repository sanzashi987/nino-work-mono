package controller

import (
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var signer = []byte("signer.nino.work")

type AuthClaims struct {
	UserID   uint64
	Username string
	jwt.RegisteredClaims
}

func GenerateToken(username string, userId uint64, expires ...time.Duration) (string, error) {
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

const (
	UserID     = "UserID"
	Username   = "Username"
	CookieName = "login_token"
)

func ValidateFromCtx(ctx *gin.Context) (*AuthClaims, error) {
	jwtToken, err := ctx.Cookie(CookieName)
	if err != nil {
		return nil, err
	}

	claim, err := ValidateToken(jwtToken)
	if err != nil {
		return nil, err
	}

	return claim, nil
}

func ValidateMiddleware(loginPageUrl string) func(*gin.Context) {

	return func(ctx *gin.Context) {
		claim, err := ValidateFromCtx(ctx)
		if err != nil {
			redirectURL := loginPageUrl + "?redirect=" + url.QueryEscape(ctx.Request.Referer())
			ctx.Redirect(http.StatusFound, redirectURL)
			return
		}
		ctx.Set(UserID, claim.UserID)
		ctx.Set(Username, claim.Username)
		ctx.Next()
	}
}
