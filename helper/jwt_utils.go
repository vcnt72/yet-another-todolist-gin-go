package helper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/yet-another-todo-list-golang/config"
	"github.com/yet-another-todo-list-golang/model/entity"
)

var secret = []byte(config.GetEnvConfig("server.secret"))

func GenerateJWT(user entity.User) (error, string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        user.ID,
		Audience:  "YATD",
		ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    user.Email,
		Subject:   user.Email,
	})

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return err, tokenString
	}
	return nil, tokenString
}

func VerifyJWT(token string) (error, *jwt.Token) {

	tkn, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return err, nil
	}
	return nil, tkn
}
