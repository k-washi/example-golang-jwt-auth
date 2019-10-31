package jwtauthdomain

//JwtCollectInfo key userid return jwt signature
//this need expire, so redis need setting expire
type JwtCollectInfo struct {
	Signature map[string]string
}

//RegisterJwt register jwt header, payload and signature
type RegisterJwt struct {
	Header  string
	Payload string
	Sign    string
}
