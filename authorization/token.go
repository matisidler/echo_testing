package authorization

import (
	"api_persons_echo/model"
	"errors"

	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(data *model.Login) (string, error) {
	claim := model.Claim{
		Email: data.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "Mati el más capo",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateToken(token string) (model.Claim, error) {
	tok, err := jwt.ParseWithClaims(token, &model.Claim{}, verifyFunction)
	if err != nil {
		return model.Claim{}, err
	}
	if !tok.Valid {
		return model.Claim{}, errors.New("not valid token")
	}
	claim, ok := tok.Claims.(*model.Claim)
	if !ok {
		return model.Claim{}, errors.New("can't obtain claims")
	}
	return *claim, nil
}

func verifyFunction(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
