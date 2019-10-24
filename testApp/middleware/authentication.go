package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
		ok, err := true, nil
		if err != nil {
			log.Printf("Authenticate application err")
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Authenticate application err",
			})
		}
		if ok {
			c.Next()
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Authentication error",
			})
		}
	}
}
