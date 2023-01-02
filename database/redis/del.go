package redis

import "log"

func DelToken(id string) {
	client := RedisConnect()
	_, err := client.Del(id).Result()
	if err != nil {
		log.Fatal(err)
	}
}
