package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"

	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/ibilalkayy/proctl/middleware"
	"github.com/spf13/cobra"
	"gopkg.in/gomail.v2"
)

func Verify(toEmail, accountName string) {
	mail := gomail.NewMessage()
	myEmail := middleware.LoadEnvVariable("APP_EMAIL")
	myPassword := middleware.LoadEnvVariable("APP_PASSWORD")

	body := new(bytes.Buffer)
	temp, err := template.ParseFiles("cmd/views/template.html")
	if err != nil {
		log.Fatal(err)
	}

	AccountEmail := redis.GetAccountInfo("AccountEmail")
	AccountPassword := redis.GetAccountInfo("AccountPassword")
	combinedText := AccountEmail + AccountPassword
	encodedText := Encode(combinedText)

	getAccountName := AccountInfo{
		GetAccountName: accountName,
		GetEncodedText: encodedText,
	}

	if err := temp.Execute(body, getAccountName); err != nil {
		fmt.Println(errors.New("Count load the template"))
	}

	mail.SetHeader("From", myEmail)
	mail.SetHeader("To", toEmail)
	mail.SetHeader("Reply-To", myEmail)
	mail.SetHeader("Subject", "[proctl] Confirm your email address")
	mail.SetBody("text/html", body.String())
	a := gomail.NewDialer("smtp.gmail.com", 587, myEmail, myPassword)
	if err := a.DialAndSend(mail); err != nil {
		log.Fatal(err)
	}
}

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Check verification of a user",
	Run: func(cmd *cobra.Command, args []string) {
		loginToken := redis.GetAccountInfo("LoginToken")
		if len(loginToken) != 0 && jwt.RefreshToken() {
			AccountName := redis.GetAccountInfo("AccountName")
			AccountEmail := redis.GetAccountInfo("AccountEmail")
			Verify(AccountEmail, AccountName)

			var verificationCode string
			fmt.Printf("Enter the verification code: ")
			fmt.Scanln(&verificationCode)

			getVerificationCode := redis.GetAccountInfo("VerificationCode")
			if getVerificationCode == verificationCode {
				fmt.Println("You're successfully verified")
			} else {
				fmt.Println(errors.New("You're not verified"))
			}
		} else {
			fmt.Println(errors.New("First login to verify an account"))
		}
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)
}
