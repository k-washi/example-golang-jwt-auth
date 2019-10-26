package middleware

import (
	"log"
	"net/http"

	"github.com/k-washi/example-golang-jwt-auth/src/utils"

	jwtauthclient "github.com/k-washi/example-golang-jwt-auth/src/client/jwtAuthClient"

	"github.com/gin-gonic/gin"
)

/*
認可
*/

//JwtMiddleware middleware of jwt check
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		/**/
		jwtPayload, err := jwtauthclient.JwtFBgRPCclient.ConfirmJwt(c)
		if err != nil {
			log.Printf("Authenticate application err")
			c.JSON(http.StatusInternalServerError,
				utils.NewAPIError(http.StatusInternalServerError, "Authenticate application: "+err.Error()),
			)
			c.Abort()
			return

		}
		if jwtPayload != nil {
			if jwtPayload.User != "" {

				_, err := jwtauthclient.JwtFBgRPCclient.SetJwtPayloadHeader(c, jwtPayload)
				if err != nil {
					c.JSON(http.StatusInternalServerError,
						utils.NewAPIError(http.StatusInternalServerError, "Authentication error: "+err.Error()),
					)
					c.Abort()
					return
				}
				log.Printf("Authenticate application ok")
				c.Next()
				return

			} else {
				c.JSON(http.StatusInternalServerError,
					utils.NewAPIError(http.StatusInternalServerError, "Authentication error: "+err.Error()),
				)
				c.Abort()
				return

			}
		}
		c.JSON(http.StatusInternalServerError,
			utils.NewAPIError(http.StatusInternalServerError, "Authentication error: Can not catch jwt authorization"),
		)
		c.Abort()
		return

	}
}
