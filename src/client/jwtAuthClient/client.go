package jwtauthclient

import (
	"errors"
	"fmt"
	"log"
	"strings"

	jwtauthpb "github.com/k-washi/example-golang-jwt-auth/src/jwtAuthpb"

	"google.golang.org/grpc"

	"github.com/k-washi/example-golang-jwt-auth/src/utils"

	"github.com/gin-gonic/gin"
)

type jwtFBgRPCclientServiceInterface interface {
	ConfirmJwt(*gin.Context) (*utils.JwtPayload, error)
	ConfirmAuth(c *gin.Context) (*utils.JwtPayload, error)
	SetJwtPayloadHeader(*gin.Context, *utils.JwtPayload) (*gin.Context, error)
	GetJwtPayloadHeader(*gin.Context) (*utils.JwtPayload, error)
}

var (
	//JwtFBgRPCclient jwt auth client
	JwtFBgRPCclient jwtFBgRPCclientServiceInterface
)

type jwtFBgRPCclient struct{}

func init() {
	JwtFBgRPCclient = &jwtFBgRPCclient{}
}

//SetJwtPayloadHeader set payload
func (s *jwtFBgRPCclient) SetJwtPayloadHeader(c *gin.Context, payload *utils.JwtPayload) (*gin.Context, error) {
	c.Set("AuthorizedUser", payload.User)
	c.Set("AuthorizedEmail", payload.Email)
	return c, nil
}

//GetJwtPayloadHeader set payload
func (s *jwtFBgRPCclient) GetJwtPayloadHeader(c *gin.Context) (*utils.JwtPayload, error) {
	user, ok := c.Get("AuthorizedUser")
	if !ok {
		log.Printf("Error: jwt payload can not get")
		return nil, errors.New("Error: jwt payload can not get")
	}
	email, ok := c.Get("AuthorizedEmail")
	if !ok {
		log.Printf("Error: jwt payload can not get")
		return nil, errors.New("Error: jwt payload can not get")
	}

	res := &utils.JwtPayload{}
	if userStr, ok := user.(string); ok {
		res.User = userStr
	} else {
		log.Printf("Error: jwt payload type assersion")
		return nil, errors.New("Error: jwt payload type assersion")
	}

	if emailStr, ok := email.(string); ok {
		res.Email = emailStr
	} else {
		log.Printf("Error: jwt payload type assersion")
		return nil, errors.New("Error: jwt payload type assersion")
	}
	return res, nil
}

//ConfirmConnInitialize initialize confirm process
func confirmConnInitialize(c *gin.Context) (string, *grpc.ClientConn, jwtauthpb.JwtServiceClient, error) {
	//Authorization: Bearer e7?aXaGEGkKLK...
	authHeader := c.GetHeader("Authorization")
	idToken := strings.Replace(authHeader, "Bearer ", "", 1)

	//get env value of host and port
	ambassadorHostAndPort, err := utils.GetAmbassadorHostAndPort()
	if err != nil {
		log.Printf("Error ConfirmJwt Init: " + err.Error())
		return "", nil, nil, err
	}

	url := ambassadorHostAndPort.Host + ":" + ambassadorHostAndPort.Port
	//url := ":" + ambassadorHostAndPort.Port
	log.Printf("url:" + url)

	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Printf("Error ConfirmJwt Init: " + err.Error())
		return "", nil, nil, fmt.Errorf("JwtFBgRPCclient: Could not connect: %v", err)
	}

	clientConn := jwtauthpb.NewJwtServiceClient(conn)

	return idToken, conn, clientConn, nil
}
