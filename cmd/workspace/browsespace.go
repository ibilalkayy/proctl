package workspace

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
)

// browsespaceCmd represents the browsespace command
var browsespaceCmd = &cobra.Command{
	Use:   "browsespace",
	Short: "Browse all the workspaces",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the login token, account email, member login token, and member email from Redis
		loginToken := redis.GetAccountInfo("LoginToken")
		accountEmail := redis.GetAccountInfo("AccountEmail")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")
		memberEmail := redis.GetAccountInfo("MemberEmail")

		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			// Retrieve and print the workspace name for the user
			name := mysql.FindWorkspace(accountEmail)
			fmt.Println(name)
		} else if len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			// Retrieve and print the workspace name for the member
			name := mysql.FindWorkspace(memberEmail)
			fmt.Println(name)
		} else {
			fmt.Println(errors.New("First login to browse workspaces"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(browsespaceCmd)
}
