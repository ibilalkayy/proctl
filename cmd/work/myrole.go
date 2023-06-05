package work

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
)

// myroleCmd represents the myrole command
var myroleCmd = &cobra.Command{
	Use:   "myrole",
	Short: "Show you the role",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the user's account email and login token from Redis
		accountEmail := redis.GetAccountInfo("AccountEmail")
		loginToken := redis.GetAccountInfo("LoginToken")

		// Get the member's account email and login token from Redis
		memberAccountEmail := redis.GetAccountInfo("MemberAccountEmail")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		// Find the role information for the user and member
		userCredentials := mysql.FindRole(accountEmail)
		memberCredentials := mysql.FindRole(memberAccountEmail)

		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			// Check if the user has role credentials
			if len(userCredentials[0]) != 0 && len(userCredentials[1]) != 0 {
				fmt.Printf("My role: %s\n", userCredentials[1])
				return
			} else {
				fmt.Println(errors.New("This user has no role"))
			}
		} else if len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			// Check if the member has role credentials
			if len(memberCredentials[0]) != 0 && len(memberCredentials[1]) != 0 {
				fmt.Printf("My role: %s\n", memberCredentials[1])
				return
			} else {
				fmt.Println(errors.New("This member has no role"))
			}
		} else {
			fmt.Println(errors.New("First login to show you the role"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(myroleCmd)
}
