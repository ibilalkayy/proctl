package redis

import (
	"encoding/json"
)

func GetUserCredentials(totalColumns int) ([]string, []string, []string, []string, bool) {
	client := RedisConnect()
	getEmails, err := client.LRange("UserEmails", 0, int64(totalColumns)-1).Result()
	getPasswords, err := client.LRange("UserPasswords", 0, int64(totalColumns)-1).Result()
	getFullName, err := client.LRange("UserFullNames", 0, int64(totalColumns)-1).Result()
	getAccountName, err := client.LRange("UserAccountNames", 0, int64(totalColumns)-1).Result()

	if err != nil {
		return []string{}, []string{}, []string{}, []string{}, false
	}

	return getEmails, getPasswords, getFullName, getAccountName, true
}

func GetMemberCredentials(totalColumns int) ([]string, []string, []string, []string, bool) {
	client := RedisConnect()
	getEmails, err := client.LRange("MemberEmails", 0, int64(totalColumns)-1).Result()
	getPasswords, err := client.LRange("MemberPasswords", 0, int64(totalColumns)-1).Result()
	getFullName, err := client.LRange("MemberFullNames", 0, int64(totalColumns)-1).Result()
	getAccountName, err := client.LRange("MemberAccountNames", 0, int64(totalColumns)-1).Result()

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
