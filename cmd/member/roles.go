package member

import (
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
)

// rolesCmd represents the roles command
var rolesCmd = &cobra.Command{
	Use:   "roles",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		role, _ := cmd.Flags().GetString("role")

		accountEmail := redis.GetAccountInfo("AccountEmail")
		loginToken := redis.GetAccountInfo("LoginToken")
		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			redis.DelToken("MemberLoginToken")
			mysql.InsertRole(accountEmail, role)
			fmt.Println("Your role is successfully added")
			return
		}

		memberAccountEmail := redis.GetAccountInfo("MemberAccountEmail")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")
		if len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			redis.DelToken("LoginToken")
			mysql.InsertRole(memberAccountEmail, role)
			fmt.Println("Your role is successfully added")
			return
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(rolesCmd)
	rolesCmd.Flags().StringP("role", "r", "", "Specify the role to setup")
}
