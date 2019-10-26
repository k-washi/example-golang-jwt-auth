package jwtauthserver

import (
	"log"
	"net"

	jwtauthpb "github.com/k-washi/example-golang-jwt-auth/src/jwtAuthpb"
	"github.com/k-washi/example-golang-jwt-auth/src/utils"
	"google.golang.org/grpc"
)

//StartApp gRPC controller
func StartApp() {

	ambassadorHostAndPort, err := utils.GetAmbassadorHostAndPort()
	if err != nil {
		log.Fatalf("Failed to get env valiable of host and port")
	}
	url := ambassadorHostAndPort.Host + ":" + ambassadorHostAndPort.Port
	lis, err := net.Listen("tcp", url)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	jwtauthpb.RegisterJwtServiceServer(s, JwtFBgRPCserver)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
