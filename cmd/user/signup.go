package user

import (
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/ibilalkayy/proctl/middleware"
	"github.com/spf13/cobra"
)

// emailValid checks if the provided email address is valid using a regular expression pattern.
// It returns true if the email is valid, otherwise false.
func emailValid(email string) bool {
	regexEmail := regexp.MustCompile(`^[a-zA-Z0-9-_]+@[a-z]+\.[a-z]{1,3}$`)
	return regexEmail.MatchString(email)
}

// Encode encodes the provided string using base64 encoding.
// It returns the encoded string.
func Encode(s string) string {
	body := base64.StdEncoding.EncodeToString([]byte(s))
	return string(body)
}

// GetRandomCode concatenates the first and second strings and encodes the result using base64 encoding.
// It returns the encoded string.
func GetRandomCode(first, second string) string {
	var AccountDetails string
	combinedText := first + second
	AccountDetails = Encode(combinedText)
	return AccountDetails
}

// signupCmd represents the signup command
var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "Provide an email address, password, full name, and account name in order to signup",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the values of the command-line flags
		signupEmail, _ := cmd.Flags().GetString("email")
		signupPassword, _ := cmd.Flags().GetString("password")
		signupFullName, _ := cmd.Flags().GetString("full name")
		signupAccountName, _ := cmd.Flags().GetString("account name")
		hashPass := middleware.HashPassword([]byte(signupPassword))

		// Check if a login token already exists
		loginToken := redis.GetAccountInfo("LoginToken")
		if len(loginToken) == 0 {
			// Validate the provided email, password, full name, and account name
			if len(signupEmail) != 0 && emailValid(signupEmail) && len(signupPassword) != 0 && len(signupFullName) != 0 && len(signupAccountName) != 0 {
				signupCredentials := [4]string{signupEmail, hashPass, signupFullName, signupAccountName}
				mysql.InsertSignupData(signupCredentials)

				// Generate a JWT token
				tokenString, jwtTokenGenerated := jwt.GenerateJWT()
				if jwtTokenGenerated {
					// Get user credentials from Redis
					totalColumns := mysql.CountTableColumns("Signup")
					redis.SetUserCredentials(signupCredentials)
					redisSignupEmail, redisSignupPassword, redisSignupFullName, redisSignupAccountName, _ := redis.GetUserCredentials(totalColumns)

					// Store account information in Redis
					redis.SetAccountInfo("LoginToken", tokenString)
					redis.SetAccountInfo("AccountName", redisSignupAccountName[0])
					redis.SetAccountInfo("AccountFullName", redisSignupFullName[0])
					redis.SetAccountInfo("AccountEmail", redisSignupEmail[0])
					redis.SetAccountInfo("AccountPassword", redisSignupPassword[0])

					accountEmail := redis.GetAccountInfo("AccountEmail")
					accountPassword := redis.GetAccountInfo("AccountPassword")

					// Generate a random account code
					accountCode := GetRandomCode(accountEmail, accountPassword)

					// Check if the account exists in the database
					_, _, mysqlStatus, _ := mysql.FindAccount(accountEmail, accountPassword)
					if mysqlStatus == "0" {
						redis.SetAccountInfo("VerificationCode", accountCode)
					}
					fmt.Println("You have successfully created an account.")
				} else {
					fmt.Println("Signup failure.")
				}
			} else {
				fmt.Println(errors.New("Provide the correct or full credentials."))
			}
		} else {
			fmt.Println(errors.New("First logout and then signup."))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(signupCmd)
	signupCmd.Flags().StringP("email", "e", "", "Specify an email address to signup")
	signupCmd.Flags().StringP("password", "p", "", "Specify the password to signup")
	signupCmd.Flags().StringP("full name", "f", "", "Specify the full name to signup")
	signupCmd.Flags().StringP("account name", "a", "", "Specify the account name to signup")
}
