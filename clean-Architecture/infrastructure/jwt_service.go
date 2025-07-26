package infrastructure

import (
	"errors"
	"task-management/domain"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret_token_generator")

func CreateToken(user domain.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"user_id": user.ID,
		"email": user.Email})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if !token.Valid {
		return errors.New("invalid token")
	}
	return err
}
