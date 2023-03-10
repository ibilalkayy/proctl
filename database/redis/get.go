package redis

import (
	"encoding/json"
)

func GetCredentials(totalColumns int) ([]string, []string, []string, []string, bool) {
	client := RedisConnect()
	getEmails, err := client.LRange("Emails", 0, int64(totalColumns)-1).Result()
	getPasswords, err := client.LRange("Passwords", 0, int64(totalColumns)-1).Result()
	getFullName, err := client.LRange("Full Names", 0, int64(totalColumns)-1).Result()
	getAccountName, err := client.LRange("Account Names", 0, int64(totalColumns)-1).Result()

	if err != nil {
		return []string{}, []string{}, []string{}, []string{}, false
	}

	return getEmails, getPasswords, getFullName, getAccountName, true
}

func GetAccountInfo(id string) string {
	client := RedisConnect()
	val, err := client.Get(id).Result()
	if err != nil {
		return ""
	}

	cred := MyInfo{}
	err = json.Unmarshal([]byte(val), &cred)
	if err != nil {
		return ""
	}

	return cred.MyKey
}
