package middleware

import (
	"log"
	"net/http"

	"github.com/k-washi/example-golang-jwt-auth/src/utils"

	jwtauthclient "github.com/k-washi/example-golang-jwt-auth/src/client/jwtAuthClient"

	"github.com/gin-gonic/gin"
)

/*
認可 Authorizatoin (AuthZ)
*/

//JwtMiddleware middleware of jwt check
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		/**/
		jwtPayload, err := jwtauthclient.JwtFBgRPCclient.ConfirmJwt(c)
		if err != nil {
			log.Printf("Error JwtMiddleware: " + err.Error())
			c.JSON(http.StatusInternalServerError,
				utils.NewAPIError(http.StatusInternalServerError, "Error JwtMiddleware:: "+err.Error()),
			)
			c.Abort()
			return
		}
		if jwtPayload.User == "" || jwtPayload.Email == "" {
			log.Printf("Error JwtMiddleware: pyload empty")
			c.JSON(http.StatusInternalServerError,
				utils.NewAPIError(http.StatusInternalServerError, "Error JwtMiddleware: pyload empty"),
			)
			c.Abort()
			return
		}

		//Set user info to header
		_, err = jwtauthclient.JwtFBgRPCclient.SetJwtPayloadHeader(c, jwtPayload)
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				utils.NewAPIError(http.StatusInternalServerError, "Error JwtMiddleware: "+err.Error()),
			)
			c.Abort()
			return
		}

		log.Printf("Authorizetion application ok")
		c.Next()
		return

	}
}
