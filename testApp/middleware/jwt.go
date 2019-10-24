package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
認可
*/

//JwtMiddleware middleware of jwt check
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		/**/
		ok, register, err := true, false, nil
		if err != nil {
			log.Printf("JWT application err")
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "JWT application err",
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

		if ok {
			c.Next()
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "jwt authorization error",
			})
		}

	}
}
