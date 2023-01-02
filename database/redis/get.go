package redis

import (
	"regexp"

	"github.com/ibilalkayy/proctl/database/mysql"
)

func GetCredentials() ([]string, []string, bool) {
	client := RedisConnect()
	totalColumns := mysql.CountTableColumns("Signup")
	getEmails, err := client.LRange("Emails", 0, int64(totalColumns)).Result()
	getPasswords, err := client.LRange("Passwords", 0, int64(totalColumns)).Result()

	if err != nil {
		return []string{}, []string{}, false
	}

	return getEmails, getPasswords, true
}

func GetToken(id string) string {
	client := RedisConnect()
	val, err := client.Get(id).Result()
	if err != nil {
		return ""
	}

	re, err := regexp.Compile(`.*"|".*`)
	if err != nil {
		return ""
	}

	value := re.ReplaceAllString(val, "")
	return value
}
