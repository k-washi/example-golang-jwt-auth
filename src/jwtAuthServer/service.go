package jwtauthserver

import (
	"context"
	"errors"
	"log"

	"firebase.google.com/go/auth"

	firebase "firebase.google.com/go"
	jwtauthpb "github.com/k-washi/example-golang-jwt-auth/src/jwtAuthpb"
	"github.com/k-washi/example-golang-jwt-auth/src/utils"
	"github.com/k-washi/jwt-decode/jwtdecode"
	"google.golang.org/api/option"
)

type jwtFBgRPCserverServiceInterface interface {
	//createAuthCreate(*auth.Client) error
	JwtCheck(context.Context, *jwtauthpb.JwtRequest) (*jwtauthpb.JwtResponse, error)
	AuthCheck(context.Context, *jwtauthpb.JwtRequest) (*jwtauthpb.AuthResponse, error)
}

var (
	//JwtFBgRPCserver jwt authorization server
	JwtFBgRPCserver jwtFBgRPCserverServiceInterface

	//jwtFBAuthClient auth client credented by google app credentials file
	jwtFBAuthClient *auth.Client
)

type jwtFBgRPCserver struct{}

func init() {

	JwtFBgRPCserver = &jwtFBgRPCserver{}

	c, err := authConfig()
	if err != nil {
		log.Fatalf("Error: Firebase auth can't credent")
	}
	jwtFBAuthClient = c
	/*
		if err = JwtFBgRPCserver.createAuthCreate(c); err != nil {
			log.Fatalf("Error: Firebase auth can't set")
		}
	*/

}

func authConfig() (*auth.Client, error) {
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
	return auth, nil
}

//jwtUserGet return user info
func getUserInfoFromJwt(jwt string) (string, string, error) {
	//jwt decompose
	hCS, err := jwtdecode.JwtDecode.DecomposeFB(jwt)
	if err != nil {
		return "", "", errors.New("Error jwt decompose: " + err.Error())
	}
	//jwt deconde
	payload, err := jwtdecode.JwtDecode.DecodeClaimFB(hCS[1])
	if err != nil {
		return "", "", errors.New("Error jwt decode: " + err.Error())
	}

	return payload.Subject, payload.Email, nil
}
