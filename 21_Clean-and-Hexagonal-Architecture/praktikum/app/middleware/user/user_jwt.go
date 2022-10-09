package user

import (
	"belajar-go-echo/app/middleware/constants"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateUserToken(userId uint, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return token.SignedString([]byte(constants.SECRET_JWT))
}