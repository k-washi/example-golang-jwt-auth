package utils

import (
	"errors"
	"log"
	"os"
)

//AmbassadorHostAndPort return host and port for ambassador
type AmbassadorHostAndPort struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

// GetAmbassadorHostAndPort return
//$AMBASSADORHOST:$PORT
func GetAmbassadorHostAndPort() (*AmbassadorHostAndPort, error) {
	funcName := "GetAmbassadorHostAndPort"

	AmbassadorHost := os.Getenv("AMBASSADORHOST")
	if AmbassadorHost == "" {
		return nil, errors.New(funcName + ": Empty host")
	}

	Port := os.Getenv("PORT")
	if Port == "" {
		return nil, errors.New(funcName + ": Empty port")
	}

	log.Printf("GetAmbassadorHostAndPort: " + AmbassadorHost + " / " + Port)
	return &AmbassadorHostAndPort{
		Host: AmbassadorHost,
		Port: Port,
	}, nil

}

//GetGoogleAppCredentialsFilePath get firebase api key file path
func GetGoogleAppCredentialsFilePath() (string, error) {
	funcName := "GetGoogleAppCredentialsFilePath"

	GoogleAppCredFPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if GoogleAppCredFPath == "" {
		return "", errors.New(funcName + ": Empty env valiable")
	}
	if !FileExists(GoogleAppCredFPath) {
		return "", errors.New(funcName + ": Bad file path")
	}
	log.Printf("GetGoogleAppCredentialsFilePath: " + GoogleAppCredFPath)
	return GoogleAppCredFPath, nil
}

//FileExists check file exist
//if file exists, return true
func FileExists(filename string) bool {
	f, err := os.Stat(filename)
	if os.IsNotExist(err) {
		log.Printf("File is not exist: " + filename)
		return false
	}
	if f.IsDir() {
		log.Printf("File is Dir: " + filename)
		return false
	}

	return true

}
