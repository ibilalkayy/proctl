package workspace

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/cmd/user"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/email"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
)

type MemberInfo struct {
	GetAccountName string
}

// invitespaceCmd represents the invitespace command
var invitespaceCmd = &cobra.Command{
	Use:   "invitespace",
	Short: "Invite other members in the workspace",
	Run: func(cmd *cobra.Command, args []string) {
		loginToken := redis.GetAccountInfo("LoginToken")
		if len(loginToken) != 0 && jwt.RefreshToken() {
			inviteWorkspaceEmail, _ := cmd.Flags().GetString("email")
			accountName := redis.GetAccountInfo("AccountName")
			accountEmail := redis.GetAccountInfo("AccountEmail")

			getVerificationCode := user.GetRandomCode(accountEmail, inviteWorkspaceEmail)
			var verificationCode string

			values := [5]string{"member-template", accountName, getVerificationCode, inviteWorkspaceEmail, accountName + " has invited you to collaborate on the proctl project"}
			email.Verify(values)

			fmt.Printf("Enter the verification code: ")
			fmt.Scanln(&verificationCode)

			if len(verificationCode) != 0 && getVerificationCode == verificationCode {
				tokenString, _ := jwt.GenerateJWT()
				redis.SetAccountInfo("LoginToken", tokenString)
				mysql.InsertMemberData(inviteWorkspaceEmail)
				fmt.Println("Your account is successfully verified")
			} else {
				fmt.Println(errors.New("Please enter the correct verification code"))
			}
		} else {
			fmt.Println(errors.New("First login to invite a member"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(invitespaceCmd)
	invitespaceCmd.Flags().StringP("email", "e", "", "Specify an email address to invite people")
}
