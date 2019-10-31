package jwtauthserver

import (
	"context"
	"log"

	jwtauthpb "github.com/k-washi/example-golang-jwt-auth/src/jwtAuthpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Failed to jwt decode: "+err.Error(),
		)
	}

	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "the client canceld the request")
	}

	log.Printf("Success jwtCheck: " + res.GetJwtCheckResult().GetUser())
	return res, nil
}

func confirmJwtWithFB(jwt string) (*jwtauthpb.JwtResponse, error) {
	//JWT検証
	if _, err := jwtFBAuthClient.VerifyIDToken(context.Background(), jwt); err != nil {
		log.Printf("error: %v\n", err)
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Failed to verify jwt token by firebase: "+err.Error(),
		)
	}

	//jwt decode
	user, email, err := getUserInfoFromJwt(jwt)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Failed to jwt decode: "+err.Error(),
		)
	}

	//Setting user and email info to response
	res := &jwtauthpb.JwtResponse{
		JwtCheckResult: &jwtauthpb.JwtCheckResult{
			User:  user,
			Email: email,
		}}

	return res, nil

}
