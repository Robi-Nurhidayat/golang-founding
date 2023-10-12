package auth

import "github.com/golang-jwt/jwt/v5"

type ServiceAuth interface {
	GenerateToken(userId int) (string, error)
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