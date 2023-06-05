package redis

import "log"

// DelToken deletes a token from Redis
func DelToken(id string) {
	client := RedisConnect()
	_, err := client.Del(id).Result()
	if err != nil {
		log.Fatal(err)
	}
}
