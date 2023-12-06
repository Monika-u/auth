package middleware

import (
	"demo/constants"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims represents the structure of JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJwtToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims[ClaimAuthorized] = true
	claims[ClaimUsername] = username
	claims[ClaimExp] = time.Now().Add(TokenExpiryDuration).Unix()

	tokenString, err := token.SignedString([]byte(constants.SerectKey))
	if err != nil {
		log.Printf("Error generating JWT token: %s", err)
		return "", err
	}
	return tokenString, nil
}

const (
	ClaimAuthorized     = "authorized"
	ClaimUsername       = "username"
	ClaimExp            = "exp"
	TokenExpiryDuration = time.Minute * 30
)

func ExtractUserFromToken(r *http.Request) error {
	tokenString := r.Header.Get("token")
	if tokenString == "" {
		return fmt.Errorf("missing token")
	}

	_, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.SerectKey), nil
	})
	if err != nil {
		log.Printf("Token Error: %s", err)
		return fmt.Errorf("invalid token")
	}

	return nil
}
