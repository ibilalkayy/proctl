package middleware

import (
	"log"

	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

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

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func HashPassword(value []byte) string {
	hash, err := bcrypt.GenerateFromPassword(value, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}
