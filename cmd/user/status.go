package user

import (
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Status of the logged in or the logged out user.",
	Run: func(cmd *cobra.Command, args []string) {
		// Check if a login token exists for the user
		loginToken := redis.GetAccountInfo("LoginToken")
		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			redis.DelToken("MemberLoginToken")
			accountName := redis.GetAccountInfo("AccountName")
			fmt.Printf("%s is logged in.\n", accountName)
			return
		}

		// Check if a login token exists for the member
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")
		if len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			redis.DelToken("LoginToken")
			accountName := redis.GetAccountInfo("MemberAccountName")
			fmt.Printf("%s is logged in.\n", accountName)
			return
		}

		// If no login token exists, the user is logged out
		fmt.Println("User is logged out.")
	},
}

func init() {
	cmd.RootCmd.AddCommand(statusCmd)
}
