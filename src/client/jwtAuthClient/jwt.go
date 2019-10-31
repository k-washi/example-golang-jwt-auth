package jwtauthclient

import (
	"context"
	"errors"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	jwtauthpb "github.com/k-washi/example-golang-jwt-auth/src/jwtAuthpb"

	"github.com/k-washi/example-golang-jwt-auth/src/utils"

	"github.com/gin-gonic/gin"
)

func (s *jwtFBgRPCclient) ConfirmJwt(c *gin.Context) (*utils.JwtPayload, error) {
	//Authorization: Bearer e7?aXaGEGkKLK...
	idToken, conn, clientConn, err := confirmConnInitialize(c)
	if err != nil {
		log.Printf("Error ConfirmJwt init:" + err.Error())
		return nil, err
	}
	defer conn.Close()

	res, err := doUnaryJwt(clientConn, idToken)
	if err != nil {
		log.Printf("Error ConfirmJwt duUnary: " + err.Error())
		return nil, err
	}
	return res, nil
}

func doUnaryJwt(c jwtauthpb.JwtServiceClient, jwt string) (*utils.JwtPayload, error) {
	req := &jwtauthpb.JwtRequest{
		JwtRequest: &jwtauthpb.Jwt{
			Jwt: jwt,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := c.JwtCheck(ctx, req)

	//res, err := c.JwtCheck(context.Background(), req)

	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// err from gRPC
			if respErr.Code() == codes.InvalidArgument {
				log.Printf("Error ConfirmJwt invalid: " + err.Error())
				return nil, err
			} else if respErr.Code() == codes.DeadlineExceeded {
				log.Printf("Error ConfirmJwt deadline: " + err.Error())
				return nil, errors.New("Error: Timeout was hit! Deadline was exceeded")
			} else {
				log.Printf("Error ConfirmJwt else: " + err.Error())
				return nil, errors.New("Error: " + err.Error())
			}
		} else {
			log.Printf("Error ConfirmJwt : " + err.Error())
			return nil, err
		}
	}

	payloadRes := utils.JwtPayload{
		User:     res.GetJwtCheckResult().GetUser(),
		Email:    res.GetJwtCheckResult().GetEmail(),
		Register: false,
	}
	return &payloadRes, nil

}
