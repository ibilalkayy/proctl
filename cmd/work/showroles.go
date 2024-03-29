package work

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
)

// showrolesCmd represents the showroles command
var showrolesCmd = &cobra.Command{
	Use:   "showroles",
	Short: "Show the roles of a member and the admin",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the login tokens from Redis
		loginToken := redis.GetAccountInfo("LoginToken")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		// Check if the user or member is logged in and has valid login tokens
		if (len(loginToken) != 0 && jwt.RefreshToken("user")) || (len(memberLoginToken) != 0 && jwt.RefreshToken("member")) {
			// Define the available roles
			roles := [4]string{
				"Business Owner",
				"Team Leader",
				"Team Member",
				"Freelancer",
			}

			fmt.Println("Following are the roles:")

			// Print the roles with index numbers
			index := 1
			for _, value := range roles {
				fmt.Printf("%d. %s\n", index, value)
				index++
			}
		} else {
			fmt.Println(errors.New("First login to show the roles of a member"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(showrolesCmd)
}
