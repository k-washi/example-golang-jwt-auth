package utils

type jwtParseInterface interface{}

var (
	//JwtParse return jwt payload infomation
	JwtParse jwtParseInterface
)

type jwtParse struct{}

func init() {
	JwtParse = &jwtParse{}
}

func (s *jwtParse) JwtWithFBparse(jwt string) {

}
