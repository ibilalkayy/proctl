package redis

func GetCredentials() ([]string, []string, bool) {
	client := RedisConnect()

	getEmails, err := client.LRange("Emails", 0, 1).Result()
	getPasswords, err := client.LRange("Passwords", 0, 1).Result()

	if err != nil {
		return []string{}, []string{}, false
	}

	return getEmails, getPasswords, true
}
