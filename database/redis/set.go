package redis

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

type NewToken struct {
	MyToken string
}

func RedisConnect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client
}

func SetCredentials(email, password string) {
	client := RedisConnect()
	insertEmails, err := client.LPush("Emails", email).Result()
	insertPasswords, err := client.LPush("Passwords", password).Result()

	if err != nil {
		fmt.Println(insertEmails)
		fmt.Println(insertPasswords)
	}
}

func SetToken(id, tokenString string) {
	client := RedisConnect()
	json, err := json.Marshal(NewToken{MyToken: tokenString})
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Set(id, json, 0).Err(); err != nil {
		log.Fatal(err)
	}
}
