package cmd

import (
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/ibilalkayy/proctl/middleware"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

func Verify(toEmail, accountName string) {
	mail := gomail.NewMessage()
	myEmail := middleware.LoadEnvVariable("APP_EMAIL")
	myPassword := middleware.LoadEnvVariable("APP_PASSWORD")
	mail.SetHeader("From", myEmail)
	mail.SetHeader("To", toEmail)
	mail.SetHeader("Reply-To", myEmail)
	mail.SetHeader("Subject", "[proctl] Confirm your email address")
	mail.SetBody("text/html", fmt.Sprintf(`
	<center><h1>We're glad you're here, %s.</h1></center>
	<center>We just want to confirm it's you.<br><br></center>
	<center>
		<button style="background-color:#008CBA; border-color:#008CBA; border-radius: 4px; color:white; height: 50px; width: 300px;">
			Click to confirm your email address
		</button>
	<center>
	<center><br>If you didn't create a proctl account, just delete this email.</center>`, accountName))
	a := gomail.NewDialer("smtp.gmail.com", 587, myEmail, myPassword)
	if err := a.DialAndSend(mail); err != nil {
		log.Fatal(err)
	}
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
					_, _, redisSignupAccountName, _ := redis.GetCredentials()
					redis.SetAccountInfo("LoginToken", tokenString)
					redis.SetAccountInfo("AccountName", redisSignupAccountName[0])
					accountName := redis.GetAccountInfo("AccountName")
					Verify(signupEmail, accountName)
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
