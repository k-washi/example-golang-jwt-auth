package contoroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//SuccessResponse define response
type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func GetAuthSuccessStatus(c *gin.Context) {
	res := SuccessResponse{
		Status:  http.StatusOK,
		Message: "Authenticated",
	}
	c.JSON(http.StatusOK, res)
}

func GetJWTSuccessStatus(c *gin.Context) {
	res := SuccessResponse{
		Status:  http.StatusOK,
		Message: "Authorized",
	}
	c.JSON(http.StatusOK, res)
}
