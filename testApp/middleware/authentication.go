package middleware

import (
	"log"
	"net/http"

	testutils "github.com/k-washi/example-golang-jwt-auth/testApp/utils"

	"github.com/gin-gonic/gin"
	jwtauthclient "github.com/k-washi/example-golang-jwt-auth/src/client/jwtAuthClient"
	"github.com/k-washi/example-golang-jwt-auth/src/utils"
)

/*
認証 Authentication (AuthN)
*/

//AuthMiddleware middleware of check registration of jwt and differance jwt from register jwt
//stage1 no register => jwt register
//new jwt => differanc from register jwt
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		/**/
		jwtPayload, err := jwtauthclient.JwtFBgRPCclient.ConfirmAuth(c)
		if err != nil {
			log.Printf("Error AuthMiddleware: " + err.Error())
			c.JSON(http.StatusInternalServerError,
				utils.NewAPIError(http.StatusInternalServerError, "Error AuthMiddleware: "+err.Error()),
			)
			c.Abort()
			return
		}

		if jwtPayload.Register {
			res := testutils.SuccessResponse{
				Status:  http.StatusAccepted,
				Message: "Success: Authentification jwt registerd",
			}
			c.JSON(http.StatusAccepted, res)
			c.Abort()
			return
		}

		//Register true => return status:200
		//when next jwt of request is authorized and different from register jwt, authentification is true
		if jwtPayload.User == "" || jwtPayload.Email == "" {
			log.Printf("Error AuthMiddleware: pyload empty")
			c.JSON(http.StatusInternalServerError,
				utils.NewAPIError(http.StatusInternalServerError, "Error AuthMiddleware: pyload empty"),
			)
			c.Abort()
			return
		}

		//Set user info to header
		if _, err = jwtauthclient.JwtFBgRPCclient.SetJwtPayloadHeader(c, jwtPayload); err != nil {
			c.JSON(http.StatusInternalServerError,
				utils.NewAPIError(http.StatusInternalServerError, "Authentication error: "+err.Error()),
			)
			c.Abort()
			return
		}

		log.Printf("Authenticate application ok")
		c.Next()
		return

	}
}
