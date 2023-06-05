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

// updatedepsCmd represents the updatedeps command
var updatedepsCmd = &cobra.Command{
	Use:   "updatedeps",
	Short: "Update the department of a user",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the department flag value from the command
		department, _ := cmd.Flags().GetString("department")

		// Get the user's account email and login token from Redis
		accountEmail := redis.GetAccountInfo("AccountEmail")
		loginToken := redis.GetAccountInfo("LoginToken")

		// Get the member's account email and login token from Redis
		memberAccountEmail := redis.GetAccountInfo("MemberAccountEmail")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		// Find the department information for the user and member
		userCredentials := mysql.FindDepartment(accountEmail)
		memberCredentials := mysql.FindDepartment(memberAccountEmail)

		if len(loginToken) != 0 && jwt.RefreshToken("user") || len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			// Check if the user has already inserted a department
			if len(userCredentials[0]) != 0 && len(userCredentials[1]) != 0 {
				mysql.UpdateDepartment(accountEmail, department)
				fmt.Println("Your have successfully updated the department")
				return
			} else if len(memberCredentials[0]) != 0 && len(memberCredentials[1]) != 0 {
				mysql.UpdateDepartment(memberAccountEmail, department)
				fmt.Println("Your have successfully updated the department")
				return
			} else {
				fmt.Println(errors.New("You have not inserted the department. How can you update it?"))
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
