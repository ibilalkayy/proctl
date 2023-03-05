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
	"github.com/ibilalkayy/proctl/middleware"
	"github.com/spf13/cobra"
	"gopkg.in/gomail.v2"
)

type MemberInfo struct {
	GetAccountName string
}

func VerifyMember(toEmail, accountName string) {
	mail := gomail.NewMessage()
	myEmail := middleware.LoadEnvVariable("APP_EMAIL")
	myPassword := middleware.LoadEnvVariable("APP_PASSWORD")

	body := new(bytes.Buffer)
	temp, err := template.ParseFiles("cmd/views/member-template.html")
	if err != nil {
		log.Fatal(err)
	}

	getAccountName := MemberInfo{
		GetAccountName: accountName,
	}

	if err := temp.Execute(body, getAccountName); err != nil {
		fmt.Println(errors.New("Cannot load the template"))
	}

	mail.SetHeader("From", myEmail)
	mail.SetHeader("To", toEmail)
	mail.SetHeader("Reply-To", myEmail)
	mail.SetHeader("Subject", accountName+" has invited you to collaborate on the proctl project")
	mail.SetBody("text/html", body.String())
	a := gomail.NewDialer("smtp.gmail.com", 587, myEmail, myPassword)
	if err := a.DialAndSend(mail); err != nil {
		log.Fatal(err)
	}
}

// invitespaceCmd represents the invitespace command
var invitespaceCmd = &cobra.Command{
	Use:   "invitespace",
	Short: "Invite other members in the workspace",
	Run: func(cmd *cobra.Command, args []string) {
		inviteWorkspaceEmail, _ := cmd.Flags().GetString("email")
		accountName := redis.GetAccountInfo("AccountName")
		var verificationCode string

		VerifyMember(inviteWorkspaceEmail, accountName)
		fmt.Printf("Enter the verification code: ")
		fmt.Scanln(&verificationCode)

		if len(verificationCode) != 0 {
			mysql.InsertMemberData(inviteWorkspaceEmail)
		} else {
			fmt.Println("Please enter the verification code")
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(invitespaceCmd)
	invitespaceCmd.Flags().StringP("email", "e", "", "Specify an email address to invite people")
}
