package contoroller

import (
	"net/http"

	testutils "github.com/k-washi/example-golang-jwt-auth/testApp/utils"

	jwtauthclient "github.com/k-washi/example-golang-jwt-auth/src/client/jwtAuthClient"

	"github.com/gin-gonic/gin"
)

//GetJWTSuccessStatus jwt success pattern
func GetJWTSuccessStatus(c *gin.Context) {
	payload, err := jwtauthclient.JwtFBgRPCclient.GetJwtPayloadHeader(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "JWT error: " + err.Error(),
		})
	}

	res := testutils.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Authorization: " + payload.User + " /" + payload.Email,
	}
	c.JSON(http.StatusOK, res)
}

//GetAuthSuccessStatus auth success pattern
func GetAuthSuccessStatus(c *gin.Context) {
	payload, err := jwtauthclient.JwtFBgRPCclient.GetJwtPayloadHeader(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "JWT error: " + err.Error(),
		})
	}

	res := testutils.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success: Authorization: " + payload.User + " /" + payload.Email,
	}
	c.JSON(http.StatusOK, res)
}
