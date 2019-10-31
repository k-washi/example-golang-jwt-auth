package jwtauthserver

import (
	"context"
	"errors"
	"log"

	jwtauthdomain "github.com/k-washi/example-golang-jwt-auth/src/domain"
	"github.com/k-washi/jwt-decode/jwtdecode"

	jwtauthpb "github.com/k-washi/example-golang-jwt-auth/src/jwtAuthpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*jwtFBgRPCserver) AuthCheck(ctx context.Context, req *jwtauthpb.JwtRequest) (*jwtauthpb.AuthResponse, error) {
	jwt := req.GetJwtRequest().GetJwt()
	if jwt == "" {

		return nil, status.Errorf(
			codes.InvalidArgument,
			"Recived a empty string",
		)
	}

	//Verify jwt and get user info
	res, err := confirmAuthWithFB(jwt)
	if err != nil {
		//jwt decode
		user, _, e := getUserInfoFromJwt(jwt)
		if e != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"Failed to jwt decode: "+e.Error(),
			)
		}
		//古いjwtが残っている場合、削除
		if e = jwtRegisterDelete(user); e != nil {
			return nil, e
		}

		return nil, status.Errorf(
			codes.InvalidArgument,
			"Failed to jwt decode: "+err.Error(),
		)
	}

	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "the client canceld the request")
	}

	return res, nil
}

func confirmAuthWithFB(jwt string) (*jwtauthpb.AuthResponse, error) {
	//JWT検証
	if _, err := jwtFBAuthClient.VerifyIDToken(context.Background(), jwt); err != nil {
		log.Printf("error: %v\n", err)
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Failed to verify jwt token by firebase",
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
	//decompose jwt to header, claim and sign.
	hCS, err := jwtdecode.JwtDecode.DecomposeFB(jwt)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Failed to jwt decompose: "+err.Error(),
		)
	}
	signature := hCS[2]
	regsterdJwt, jwtRegisterd, err := jwtauthdomain.JwtRegister.Get(user)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Failed to jwt register check: "+err.Error(),
		)
	}

	if jwtRegisterd {
		//jwtRegisterd
		if signature != regsterdJwt.Sign {
			//OK
			//delete register jwt
			if e := jwtRegisterDelete(user); e != nil {
				return nil, e
			}
			log.Printf("Success: jwt Authentication")
			return &jwtauthpb.AuthResponse{
				AuthCheckResult: &jwtauthpb.AuthCheckResult{
					User:     user,
					Email:    email,
					Register: false,
				}}, nil
		}
		log.Print("Error: receved jwt was already registerd")
		return nil, errors.New("Error: receved jwt was already registerd")
	}

	//Register jwt
	if e := jwtRegisterCreate(user, signature); e != nil {
		return nil, e
	}

	log.Printf("Success: jwt registerd")
	return &jwtauthpb.AuthResponse{
		AuthCheckResult: &jwtauthpb.AuthCheckResult{
			User:     user,
			Email:    email,
			Register: true,
		}}, nil

}

func jwtRegisterCreate(user string, signature string) error {
	//user add
	if err := jwtauthdomain.JwtRegister.Create(user, signature); err != nil {
		return errors.New("Error jwtRegister: " + err.Error())
	}

	return nil
}

func jwtRegisterDelete(user string) error {
	//user delete
	if err := jwtauthdomain.JwtRegister.Delete(user); err != nil {
		return errors.New("Error jwtRegister: " + err.Error())
	}

	return nil
}
