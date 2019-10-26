package jwtauthserver

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	jwtauthpb "github.com/k-washi/example-golang-jwt-auth/src/jwtAuthpb"
	"github.com/k-washi/example-golang-jwt-auth/src/utils"

	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type jwtFBgRPCserverServiceInterface interface {
	JwtCheck(context.Context, *jwtauthpb.JwtRequest) (*jwtauthpb.JwtResponse, error)
	AuthCheck(context.Context, *jwtauthpb.JwtRequest) (*jwtauthpb.AuthResponse, error)
}

var (
	//JwtFBgRPCserver jwt authorization server
	JwtFBgRPCserver jwtFBgRPCserverServiceInterface
)

type jwtFBgRPCserver struct{}

func init() {

	JwtFBgRPCserver = &jwtFBgRPCserver{}
}

//JwtCheck jwt check by firebase and return user infomation
func (*jwtFBgRPCserver) JwtCheck(ctx context.Context, req *jwtauthpb.JwtRequest) (*jwtauthpb.JwtResponse, error) {

	jwt := req.GetJwtRequest().GetJwt()
	if jwt == "" {

		return nil, status.Errorf(
			codes.InvalidArgument,
			"Recived a empty string",
		)
	}

	res, err := confirmJwtWithFB(jwt)
	if err != nil {
		log.Printf("Error jwtCheck with firebase: " + err.Error())
		return nil, err
	}

	if ctx.Err() == context.Canceled {
		return nil, status.Error(codes.Canceled, "the client canceld the request")
	}

	return res, nil
}

func (*jwtFBgRPCserver) AuthCheck(ctx context.Context, req *jwtauthpb.JwtRequest) (*jwtauthpb.AuthResponse, error) {
	return &jwtauthpb.AuthResponse{
		AuthCheckResult: &jwtauthpb.AuthCheckResult{
			User:     "",
			Email:    "",
			Register: true,
		},
	}, nil
}

func confirmJwtWithFB(jwt string) (*jwtauthpb.JwtResponse, error) {
	//JWT検証
	GoogleAppCred, err := utils.GetGoogleAppCredentialsFilePath()
	if err != nil {
		return nil, err
	}
	opt := option.WithCredentialsFile(GoogleAppCred)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("error: %v\n", err)
		return nil, err
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		log.Printf("error: %v\n", err)
		return nil, err
	}

	token, err := auth.VerifyIDToken(context.Background(), jwt)
	if err != nil {
		log.Printf("error: %v\n", err)
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Failed to verify jwt token by firebase",
		)
	}

	res := &jwtauthpb.JwtResponse{JwtCheckResult: &jwtauthpb.JwtCheckResult{}}
	user := token.Claims["user_id"]
	if userStr, ok := user.(string); ok {
		res.JwtCheckResult.User = userStr
	} else {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Failed to get user name from firebase",
		)
	}

	email := token.Claims["email"]
	if emailStr, ok := email.(string); ok {
		res.JwtCheckResult.Email = emailStr
	} else {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Failed to get email from firebase",
		)
	}

	log.Printf("JwtCheck Success: " + res.JwtCheckResult.User)
	return res, nil

}
