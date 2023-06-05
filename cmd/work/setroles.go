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

		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			// Check if the user has not already inserted a role
			if len(userCredentials[0]) == 0 && len(userCredentials[1]) == 0 {
				redis.DelToken("MemberLoginToken")
				mysql.InsertRole(accountEmail, role)
				fmt.Println("Your have successfully inserted the role")
				return
			} else {
				fmt.Println(errors.New("You have already inserted the role. Please update it"))
			}
		} else if len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			// Check if the member has not already inserted a role
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
