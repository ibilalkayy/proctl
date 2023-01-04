package redis

import (
	"encoding/json"

	"github.com/ibilalkayy/proctl/database/mysql"
)

func GetCredentials() ([]string, []string, []string, bool) {
	client := RedisConnect()
	totalColumns := mysql.CountTableColumns("Signup")
	getEmails, err := client.LRange("Emails", 0, int64(totalColumns)-1).Result()
	getPasswords, err := client.LRange("Passwords", 0, int64(totalColumns)-1).Result()
	getAccountName, err := client.LRange("Account Names", 0, int64(totalColumns)-1).Result()

	if err != nil {
		return []string{}, []string{}, []string{}, false
	}

	return getEmails, getPasswords, getAccountName, true
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
