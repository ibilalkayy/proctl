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

// setrolesCmd represents the setroles command
var setrolesCmd = &cobra.Command{
	Use:   "setroles",
	Short: "Setup the role of a user",
	Run: func(cmd *cobra.Command, args []string) {
		role, _ := cmd.Flags().GetString("role")

		accountEmail := redis.GetAccountInfo("AccountEmail")
		loginToken := redis.GetAccountInfo("LoginToken")

		memberAccountEmail := redis.GetAccountInfo("MemberAccountEmail")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		userCredentials := mysql.FindRole(accountEmail)
		memberCredentials := mysql.FindRole(memberAccountEmail)

		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			if len(userCredentials[0]) == 0 && len(userCredentials[1]) == 0 {
				redis.DelToken("MemberLoginToken")
				mysql.InsertRole(accountEmail, role)
				fmt.Println("Your have successfully inserted the role")
				return
			} else {
				fmt.Println(errors.New("You have already inserted the role. Please update it"))
			}
		} else if len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			if len(memberCredentials[0]) == 0 && len(memberCredentials[1]) == 0 {
				redis.DelToken("LoginToken")
				mysql.InsertRole(memberAccountEmail, role)
				fmt.Println("Your have successfully inserted the role")
				return
			} else {
				fmt.Println(errors.New("You have already inserted the role. Please update it"))
			}
		} else {
			fmt.Println(errors.New("First login to setup the role"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(setrolesCmd)
	setrolesCmd.Flags().StringP("role", "r", "", "Specify the role to setup")
}