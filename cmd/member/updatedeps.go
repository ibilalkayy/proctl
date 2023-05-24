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

// updatedepsCmd represents the updatedeps command
var updatedepsCmd = &cobra.Command{
	Use:   "updatedeps",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		department, _ := cmd.Flags().GetString("department")

		accountEmail := redis.GetAccountInfo("AccountEmail")
		loginToken := redis.GetAccountInfo("LoginToken")

		memberAccountEmail := redis.GetAccountInfo("MemberAccountEmail")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		userCredentials := mysql.FindDepartment(accountEmail)
		memberCredentials := mysql.FindDepartment(memberAccountEmail)

		if len(loginToken) != 0 && jwt.RefreshToken("user") || len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			if len(userCredentials[0]) != 0 && len(userCredentials[1]) != 0 {
				mysql.UpdateDepartment(accountEmail, department)
				fmt.Println("Your have successfully updated the department")
				return
			} else if len(memberCredentials[0]) != 0 && len(memberCredentials[1]) != 0 {
				mysql.UpdateDepartment(memberAccountEmail, department)
				fmt.Println("Your have successfully updated the department")
				return
			} else {
				fmt.Println(errors.New("You have not inserted the department. How you can update it?"))
			}
		} else {
			fmt.Println(errors.New("First login to update the department"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(updatedepsCmd)
	updatedepsCmd.Flags().StringP("department", "d", "", "Specify the department to update")
}
