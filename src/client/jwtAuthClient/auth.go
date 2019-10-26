package jwtauthclient

import (
	"github.com/gin-gonic/gin"
	"github.com/k-washi/example-golang-jwt-auth/src/utils"
)

func (s *jwtFBgRPCclient) ConfirmAuth(c *gin.Context) (*utils.JwtPayload, bool, error) {
	return nil, true, nil
}
