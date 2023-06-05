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

// updaterolesCmd represents the updateroles command
var updaterolesCmd = &cobra.Command{
	Use:   "updateroles",
	Short: "Update the role of a user",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the role flag value from the command
		role, _ := cmd.Flags().GetString("role")

		// Get the user's account email and login token from Redis
		accountEmail := redis.GetAccountInfo("AccountEmail")
		loginToken := redis.GetAccountInfo("LoginToken")

		// Get the member's account email and login token from Redis
		memberAccountEmail := redis.GetAccountInfo("MemberAccountEmail")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		// Find the role information for the user and member
		userCredentials := mysql.FindRole(accountEmail)
		memberCredentials := mysql.FindRole(memberAccountEmail)

		if len(loginToken) != 0 && jwt.RefreshToken("user") || len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			// Check if the user has already inserted a role
			if len(userCredentials[0]) != 0 && len(userCredentials[1]) != 0 {
				mysql.UpdateRole(accountEmail, role)
				fmt.Println("Your have successfully updated the role")
				return
			} else if len(memberCredentials[0]) != 0 && len(memberCredentials[1]) != 0 {
				mysql.UpdateRole(memberAccountEmail, role)
				fmt.Println("Your have successfully updated the role")
				return
			} else {
				fmt.Println(errors.New("You have not inserted the role. How can you update it?"))
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
