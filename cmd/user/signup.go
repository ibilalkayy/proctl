package cmd

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

type AccountInfo struct {
	GetAccountName string
	GetEncodedText string
}

func emailValid(email string) bool {
	regexEmail := regexp.MustCompile(`^[a-zA-Z0-9-_]+@[a-z]+\.[a-z]{1,3}$`)
	return regexEmail.MatchString(email)
}

func HashPassword(value []byte) string {
	hash, err := bcrypt.GenerateFromPassword(value, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}

func Encode(s string) string {
	body := base64.StdEncoding.EncodeToString([]byte(s))
	return string(body)
}

// signupCmd represents the signup command
var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "Provide an email address, password, full name and account name in order to signup",
	Run: func(cmd *cobra.Command, args []string) {
		signupEmail, _ := cmd.Flags().GetString("email")
		signupPassword, _ := cmd.Flags().GetString("password")
		signupFullName, _ := cmd.Flags().GetString("full name")
		signupAccountName, _ := cmd.Flags().GetString("account name")
		hashPass := HashPassword([]byte(signupPassword))

		loginToken := redis.GetAccountInfo("LoginToken")
		if len(loginToken) == 0 {
			if len(signupEmail) != 0 && emailValid(signupEmail) && len(signupPassword) != 0 && len(signupFullName) != 0 && len(signupAccountName) != 0 {
				signupCredentials := [4]string{signupEmail, hashPass, signupFullName, signupAccountName}
				mysql.InsertSignupData(signupCredentials)

				tokenString, jwtTokenGenerated := jwt.GenerateJWT()
				if jwtTokenGenerated {
					redis.SetCredentials(signupEmail, hashPass, signupAccountName)
					redisSignupEmail, redisSignupPassword, redisSignupAccountName, _ := redis.GetCredentials()
					redis.SetAccountInfo("LoginToken", tokenString)
					redis.SetAccountInfo("AccountName", redisSignupAccountName[0])
					redis.SetAccountInfo("AccountEmail", redisSignupEmail[0])
					redis.SetAccountInfo("AccountPassword", redisSignupPassword[0])

					AccountEmail := redis.GetAccountInfo("AccountEmail")
					AccountPassword := redis.GetAccountInfo("AccountPassword")
					combinedText := AccountEmail + AccountPassword
					encodedText := Encode(combinedText)
					redis.SetAccountInfo("VerificationCode", encodedText)
					fmt.Println("You have successfully created an account.")
				} else {
					fmt.Println("Signup failure.")
				}
			} else {
				fmt.Println(errors.New("Give the correct or full credentials."))
			}
		} else {
			fmt.Println(errors.New("First logout and then signup."))
		}
	},
}

func init() {
	rootCmd.AddCommand(signupCmd)
	signupCmd.Flags().StringP("email", "e", "", "Specify an email address to signup")
	signupCmd.Flags().StringP("password", "p", "", "Specify the password to signup")
	signupCmd.Flags().StringP("full name", "f", "", "Specify the full name to signup")
	signupCmd.Flags().StringP("account name", "a", "", "Specify the account name to signup")
}