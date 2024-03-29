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

// RedisConnect establishes a connection to Redis and returns a Redis client
func RedisConnect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client
}

// SetUserCredentials stores user credentials in Redis
func SetUserCredentials(value [4]string) {
	client := RedisConnect()
	insertEmails, err := client.LPush("UserEmails", value[0]).Result()
	insertPasswords, err := client.LPush("UserPasswords", value[1]).Result()
	insertFullName, err := client.LPush("UserFullNames", value[2]).Result()
	insertAccountName, err := client.LPush("UserAccountNames", value[3]).Result()

	if err != nil {
		fmt.Println(insertEmails)
		fmt.Println(insertPasswords)
		fmt.Println(insertFullName)
		fmt.Println(insertAccountName)
	}
}

// SetMemberCredentials stores member credentials in Redis
func SetMemberCredentials(value [3]string) {
	client := RedisConnect()
	insertEmails, err := client.LPush("MemberEmails", value[0]).Result()
	insertPasswords, err := client.LPush("MemberPasswords", value[1]).Result()
	insertAccountName, err := client.LPush("MemberAccountNames", value[2]).Result()

	if err != nil {
		fmt.Println(insertEmails)
		fmt.Println(insertPasswords)
		fmt.Println(insertAccountName)
	}
}

// SetAccountInfo stores account information in Redis for the given ID
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
