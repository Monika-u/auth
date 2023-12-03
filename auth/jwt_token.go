package auth

import (
	"demo/constants"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJwtToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(constants.SerectKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
