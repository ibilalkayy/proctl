package jwt

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/middleware"
)

// GenerateJWT generates a new JSON Web Token (JWT) with a 5-minute expiration time
func GenerateJWT() (string, bool) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}

	// Load the access secret from environment variables
	byteEnv := []byte(middleware.LoadEnvVariable("ACCESS_SECRET"))

	// Create a new token with the claims and signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate the signed string token
	tokenString, err := token.SignedString(byteEnv)
	if err != nil {
		log.Fatal(err)
	}

	return tokenString, true
}

// RefreshToken refreshes the JWT by generating a new token and updating it in Redis
func RefreshToken(tokenType string) bool {
	var tokenKey string
	switch tokenType {
	case "user":
		tokenKey = "LoginToken"
	case "member":
		tokenKey = "MemberLoginToken"
	default:
		return false
	}

	// Check if the token expiration is within 30 seconds
	if time.Until(time.Now().Add(time.Duration(jwt.StandardClaims{}.ExpiresAt))) > 30*time.Second {
		return false
	} else {
		// Generate a new JWT token
		tokenString, ok := GenerateJWT()
		if ok {
			// Update the token in Redis
			redis.SetAccountInfo(tokenKey, tokenString)
			return true
		} else {
			return false
		}
	}
}
