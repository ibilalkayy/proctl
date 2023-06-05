package profile

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
)

// infoCmd represents the 'info' command
// It is a Cobra Command object responsible for running the command to display profile information.
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show the profile information",
	Run: func(cmd *cobra.Command, args []string) {
		// Retrieving user information from the Redis database
		loginToken := redis.GetAccountInfo("LoginToken")
		accountEmail := redis.GetAccountInfo("AccountEmail")
		accountPassword := redis.GetAccountInfo("AccountPassword")

		// Checking if the user is logged in and has a valid token
		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			// If the user is logged in, the user and profile information is displayed
			mysql.ListUserInfo(accountEmail, accountPassword)
			mysql.ListProfileInfo(accountEmail)
		} else {
			// If the user is not logged in, an error message is returned
			fmt.Println(errors.New("First login to show the profile information."))
		}
	},
}

func init() {
	// Adding the infoCmd to the root command of the Cobra application
	cmd.RootCmd.AddCommand(infoCmd)
}
