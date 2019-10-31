package utils

//JwtPayload client return
type JwtPayload struct {
	User     string `json:"user"`
	Email    string `json:"email"`
	Register bool   `json:"register"`
}
