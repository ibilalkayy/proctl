package member

import (
	"errors"
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
	Short: "Setup the role of a user",
	Run: func(cmd *cobra.Command, args []string) {
		role, _ := cmd.Flags().GetString("role")

		accountEmail := redis.GetAccountInfo("AccountEmail")
		loginToken := redis.GetAccountInfo("LoginToken")

		memberAccountEmail := redis.GetAccountInfo("MemberAccountEmail")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		userCredentials := mysql.FindRole(accountEmail)
		memberCredentials := mysql.FindRole(memberAccountEmail)

		if len(userCredentials[0]) == 0 && len(userCredentials[1]) == 0 {
			if len(loginToken) != 0 && jwt.RefreshToken("user") {
				redis.DelToken("MemberLoginToken")
				mysql.InsertRole(accountEmail, role)
				fmt.Println("Your have successfully inserted the role")
				return
			}
		} else {
			fmt.Println(errors.New("You have already inserted the role. Please update it"))
		}

		if len(memberCredentials[0]) == 0 && len(memberCredentials[1]) == 0 {
			if len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
				redis.DelToken("LoginToken")
				mysql.InsertRole(memberAccountEmail, role)
				fmt.Println("Your have successfully inserted the role")
				return
			}
		} else {
			fmt.Println(errors.New("You have already inserted the role. Please update it"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(rolesCmd)
	rolesCmd.Flags().StringP("role", "r", "", "Specify the role to setup")
}
