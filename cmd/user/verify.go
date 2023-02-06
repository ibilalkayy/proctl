package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
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

	AccountDetails := GetDetails()

	getAccountName := AccountInfo{
		GetAccountName: accountName,
		GetEncodedText: AccountDetails[2],
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
			AccountFullName := redis.GetAccountInfo("AccountFullName")
			AccountEmail := redis.GetAccountInfo("AccountEmail")
			AccountPassword := redis.GetAccountInfo("AccountPassword")
			getVerificationCode := redis.GetAccountInfo("VerificationCode")

			_, _, mysqlStatus, _ := mysql.FindAccount(AccountEmail, AccountPassword)
			if mysqlStatus == "0" && len(getVerificationCode) != 0 {

				Verify(AccountEmail, AccountName)
				var verificationCode string
				fmt.Printf("Enter the verification code: ")
				fmt.Scanln(&verificationCode)

				if getVerificationCode == verificationCode {
					redis.DelToken("VerificationCode")
					userData := [3]string{AccountFullName, AccountName, "1"}
					mysql.UpdateUser(userData, AccountEmail, AccountPassword)
					fmt.Println("You're successfully verified")
				} else {
					fmt.Println(errors.New("Error in verification. Please try again!!"))
				}
			} else {
				fmt.Println(errors.New("Your account is already verified"))
			}
		} else {
			fmt.Println(errors.New("First login to verify an account"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(verifyCmd)
}
