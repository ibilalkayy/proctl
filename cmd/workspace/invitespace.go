package workspace

import (
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/email"
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
		inviteWorkspaceEmail, _ := cmd.Flags().GetString("email")
		accountName := redis.GetAccountInfo("AccountName")
		var verificationCode string

		// VerifyMember(inviteWorkspaceEmail, accountName)
		values := [5]string{"member-template", accountName, "", inviteWorkspaceEmail, accountName + " has invited you to collaborate on the proctl project"}
		email.Verify(values)
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
