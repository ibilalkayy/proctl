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

// invitememCmd represents the invitemem command
var invitememCmd = &cobra.Command{
	Use:   "invitemem",
	Short: "Invite other members in the workspace",
	Run: func(cmd *cobra.Command, args []string) {
		loginToken := redis.GetAccountInfo("LoginToken")
		inviteWorkspaceEmail, _ := cmd.Flags().GetString("email")
		accountName := redis.GetAccountInfo("AccountName")
		accountEmail := redis.GetAccountInfo("AccountEmail")
		getVerificationCode := user.GetRandomCode(inviteWorkspaceEmail, inviteWorkspaceEmail)
		workspaceName := mysql.FindWorkspace(accountEmail)

		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			if len(workspaceName) != 0 {
				// Prepare values for the email invitation
				values := [5]string{"member-template", accountName, getVerificationCode, inviteWorkspaceEmail, accountName + " has invited you to collaborate on the proctl project"}
				// Send the invitation email
				email.Verify(values)
				fmt.Println("You have successfully invited a member")
			} else {
				// If no workspace exists, print an error message
				fmt.Println(errors.New("Please create a workspace first"))
			}
		} else {
			// If the user is not logged in, print an error message
			fmt.Println(errors.New("First login to invite a member"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(invitememCmd)
	// Add a command-line flag to specify the email address to invite
	invitememCmd.Flags().StringP("email", "e", "", "Specify an email address to invite people")
}
