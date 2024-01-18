package http

import "github.com/gin-gonic/gin"

func NewRouter(loginPageUrl string) *gin.Engine {
	router := gin.Default()

	root := router.Group("enc-oss-canvas/V1")

	loginGroup := root.Group(LOGIN_GROUP)
	{
		loginController := &LoginController{}
		loginGroup.POST("login", loginController.login)
		loginGroup.GET("login-verification/get-uuidkey", loginController.getUuid)
		loginGroup.GET("login-verification/get-verification-code", loginController.getVerifyCode)
		loginGroup.GET("logout", loginController.logout)
		loginGroup.GET("common/userInfo", loginController.getUserInfo)
	}

	return router
}
