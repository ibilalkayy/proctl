package jwt

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ibilalkayy/proctl/middleware"
)

func GenerateJWT() (string, bool) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}

	byteEnv := []byte(middleware.LoadEnvVariable("ACCESS_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(byteEnv)
	if err != nil {
		log.Fatal(err)
	}

	return tokenString, true
}
