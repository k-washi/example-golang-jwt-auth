package main

import (
	"log"

	jwtauthserver "github.com/k-washi/example-golang-jwt-auth/src/jwtAuthServer"
)

func main() {
	log.Printf("[jwtAuthServer] Start application")
	jwtauthserver.StartApp()
}
