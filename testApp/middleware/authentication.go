package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	jwtauthclient "github.com/k-washi/example-golang-jwt-auth/src/client/jwtAuthClient"
)

/*
認証
*/

//AuthMiddleware middleware of check registration of jwt and differance jwt from register jwt
//stage1 no register => jwt register
//new jwt => differanc from register jwt
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		/**/
		payload, register, err := jwtauthclient.JwtFBgRPCclient.ConfirmAuth(c)
		if err != nil {
			log.Printf("JWT application err")
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "JWT application: " + err.Error(),
			})
		}
		if register {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "jwt Register",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "jwt authorization error",
			})
		}

		if payload.User != "" {
			c.Next()
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "jwt authorization error",
			})
		}

	}
}
