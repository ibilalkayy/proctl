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
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		accountEmail := redis.GetAccountInfo("AccountEmail")
		loginToken := redis.GetAccountInfo("LoginToken")

		memberAccountEmail := redis.GetAccountInfo("MemberAccountEmail")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		userCredentials := mysql.FindRole(accountEmail)
		memberCredentials := mysql.FindRole(memberAccountEmail)

		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			if len(userCredentials[0]) != 0 && len(userCredentials[1]) != 0 {
				fmt.Println("user role")
				return
			} else {
				fmt.Println(errors.New("not a user role"))
			}
		} else if len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			if len(memberCredentials[0]) != 0 && len(memberCredentials[1]) != 0 {
				fmt.Println("member role")
				return
			} else {
				fmt.Println(errors.New("not a member role"))
			}
		} else {
			fmt.Println(errors.New("First login to show you the role"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(myroleCmd)
}
