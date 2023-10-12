package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type ServiceAuth interface {
	GenerateToken(userId int) (string, error)
	ValidateToken(token string)(*jwt.Token, error)
}

type jwtService struct {
}

func NewJwtService() ServiceAuth {
	return &jwtService{}
}


var SECRET_KEY = []byte("TEST123")

func (j *jwtService) GenerateToken(userId int) (string, error) {
	//	claim atau payload sama aja

	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)

	signedToken,err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken,err
	}

	return signedToken,nil
}



func(j *jwtService) ValidateToken(encodedToken string)(*jwt.Token, error) {
	token,err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_,ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid Token")
		}

		return []byte(SECRET_KEY),nil
	})

	if err != nil {
		return token, err
	}

	return token,nil

}