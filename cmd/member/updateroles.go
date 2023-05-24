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

// updaterolesCmd represents the updateroles command
var updaterolesCmd = &cobra.Command{
	Use:   "updateroles",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		role, _ := cmd.Flags().GetString("role")

		accountEmail := redis.GetAccountInfo("AccountEmail")
		loginToken := redis.GetAccountInfo("LoginToken")

		memberAccountEmail := redis.GetAccountInfo("MemberAccountEmail")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		userCredentials := mysql.FindRole(accountEmail)
		memberCredentials := mysql.FindRole(memberAccountEmail)

		if len(loginToken) != 0 && jwt.RefreshToken("user") || len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			if len(userCredentials[0]) != 0 && len(userCredentials[1]) != 0 {
				mysql.UpdateRole(accountEmail, role)
				fmt.Println("Your have successfully updated the role")
				return
			} else if len(memberCredentials[0]) != 0 && len(memberCredentials[1]) != 0 {
				mysql.UpdateRole(memberAccountEmail, role)
				fmt.Println("Your have successfully updated the role")
				return
			} else {
				fmt.Println(errors.New("You have not inserted the role. How you can update it?"))
			}
		} else {
			fmt.Println(errors.New("First login to update the role"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(updaterolesCmd)
	updaterolesCmd.Flags().StringP("role", "r", "", "Specify the role to update")
}
