package model

import "github.com/golang-jwt/jwt"

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//Claim va a tener el payload del token (osea el cuerpo, la data.)
type Claim struct {
	Email string `json:"email"`
	//StandardClaims nos agrega el Expired At, Issure, etc.
	jwt.StandardClaims
}
