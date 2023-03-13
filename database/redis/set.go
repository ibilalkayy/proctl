package redis

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

type MyInfo struct {
	MyKey string
}

func RedisConnect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client
}

func SetCredentials(value [4]string) {
	client := RedisConnect()
	insertEmails, err := client.LPush("Emails", value[0]).Result()
	insertPasswords, err := client.LPush("Passwords", value[1]).Result()
	insertFullName, err := client.LPush("Full Names", value[2]).Result()
	insertAccountName, err := client.LPush("Account Names", value[3]).Result()

	if err != nil {
		fmt.Println(insertEmails)
		fmt.Println(insertPasswords)
		fmt.Println(insertFullName)
		fmt.Println(insertAccountName)
	}
}

func SetAccountInfo(id, MyValue string) {
	client := RedisConnect()
	json, err := json.Marshal(MyInfo{MyKey: MyValue})
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Set(id, json, 0).Err(); err != nil {
		log.Fatal(err)
	}
}
