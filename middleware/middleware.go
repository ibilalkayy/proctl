package middleware

import (
	"log"

	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

// LoadEnvVariable loads an environment variable value from the .env file
func LoadEnvVariable(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatal(err)
	}

	return value
}

// HandleError logs and terminates the program if an error occurs
func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// HashPassword hashes the given password using bcrypt
func HashPassword(value []byte) string {
	hash, err := bcrypt.GenerateFromPassword(value, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}
