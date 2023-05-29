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

// setdepsCmd represents the setdeps command
var setdepsCmd = &cobra.Command{
	Use:   "setdeps",
	Short: "Setup the department of a user",
	Run: func(cmd *cobra.Command, args []string) {
		department, _ := cmd.Flags().GetString("department")

		accountEmail := redis.GetAccountInfo("AccountEmail")
		loginToken := redis.GetAccountInfo("LoginToken")

		memberAccountEmail := redis.GetAccountInfo("MemberAccountEmail")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		userCredentials := mysql.FindDepartment(accountEmail)
		memberCredentials := mysql.FindDepartment(memberAccountEmail)

		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			if len(userCredentials[0]) == 0 && len(userCredentials[1]) == 0 {
				redis.DelToken("MemberLoginToken")
				mysql.InsertDepartment(accountEmail, department)
				fmt.Println("Your have successfully inserted the department")
				return
			} else {
				fmt.Println(errors.New("You have already inserted the department. Please update it"))
			}
		} else if len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			if len(memberCredentials[0]) == 0 && len(memberCredentials[1]) == 0 {
				redis.DelToken("LoginToken")
				mysql.InsertDepartment(memberAccountEmail, department)
				fmt.Println("Your have successfully inserted the department")
				return
			} else {
				fmt.Println(errors.New("You have already inserted the department. Please update it"))
			}
		} else {
			fmt.Println(errors.New("First login to setup the department"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(setdepsCmd)
	setdepsCmd.Flags().StringP("department", "d", "", "Specify the department to setup")
}
