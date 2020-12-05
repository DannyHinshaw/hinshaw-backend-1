package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"time"
)

type JWTPayload struct {
	AccessToken string `json:"access_token"`
}

var jwtSecret = os.Getenv("JWT_SECRET")
var Secret = []byte(jwtSecret)

// Handles generating new JWTs for sessions tokens.
func GenerateJWT(userId string) (*JWTPayload, error) {
	signingMethod := jwt.SigningMethodHS512
	token := jwt.New(signingMethod)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["userId"] = userId

	// Generate the signed token.
	access, err := token.SignedString(Secret)
	if err != nil {
		log.Println("error signing token::", err)
		return nil, err
	}

	return &JWTPayload{AccessToken: access}, nil
}

// Handles validating a JWTs by alg in header and signature with secret.
func ValidateJWT(accessToken string) (bool, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return Secret, nil
	})
	if err != nil {
		log.Println("error parsing accessToken::", err)
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["userId"] == "" {
			log.Println("jwt claims missing userId")
			return false, nil
		}

		return true, nil
	}

	return false, nil
}

// Configuration for JWT middleware.
var JWTConf = middleware.JWTConfig{
	SigningKey:    Secret,
	SigningMethod: jwt.SigningMethodHS512.Name,
	ErrorHandler: func(err error) error {
		log.Println("error validating JWT in middleware::", err)
		return err
	},
}
