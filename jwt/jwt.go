package jwt

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/middleware"
)

func GenerateJWT() (string, bool) {
	expirationTime := time.Now().Add(20 * time.Second)
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

func RefreshToken() bool {
	if time.Until(time.Now().Add(time.Duration(jwt.StandardClaims{}.ExpiresAt))) > 5*time.Second {
		return false
	} else {
		tokenString, ok := GenerateJWT()
		if ok {
			redis.SetAccountInfo("LoginToken", tokenString)
			return true
		} else {
			return false
		}
	}
}
