package contoroller

import (
	"net/http"

	jwtauthclient "github.com/k-washi/example-golang-jwt-auth/src/client/jwtAuthClient"

	"github.com/gin-gonic/gin"
)

//SuccessResponse define response
type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

//GetJWTSuccessStatus jwt success pattern
func GetJWTSuccessStatus(c *gin.Context) {
	payload, err := jwtauthclient.JwtFBgRPCclient.GetJwtPayloadHeader(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "JWT error: " + err.Error(),
		})
	}

	res := SuccessResponse{
		Status:  http.StatusOK,
		Message: "Authenticated: " + payload.User + " /" + payload.Email,
	}
	c.JSON(http.StatusOK, res)
}

//GetAuthSuccessStatus auth success pattern
func GetAuthSuccessStatus(c *gin.Context) {
	res := SuccessResponse{
		Status:  http.StatusOK,
		Message: "Authorized",
	}
	c.JSON(http.StatusOK, res)
}
