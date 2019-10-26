package app

import (
	"github.com/k-washi/example-golang-jwt-auth/testApp/contoroller"
	"github.com/k-washi/example-golang-jwt-auth/testApp/middleware"
)

func mapUrls() {
	jwt := router.Group("/jwt")
	jwt.Use(middleware.JwtMiddleware())
	jwt.GET("/ex-jwt-auth", contoroller.GetJWTSuccessStatus)

	auth := router.Group("/auth")
	auth.Use(middleware.AuthMiddleware())
	auth.GET("/ex-authentication", contoroller.GetAuthSuccessStatus)

}
