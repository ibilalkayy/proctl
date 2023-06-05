package user

// Import necessary packages
import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/spf13/cobra"
)

// Global variable to store user choice
var choice string

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Confirm it to logout",
	// The Run function is the action performed when the logout command is called
	Run: func(cmd *cobra.Command, args []string) {
		// Get the login tokens from Redis cache
		loginToken := redis.GetAccountInfo("LoginToken")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		// Check if any of the tokens exist
		if len(loginToken) != 0 || len(memberLoginToken) != 0 {
			// Prompt user to confirm logout
			fmt.Printf("Want to logout [y/n]: ")
			// Scan user input into the choice variable
			fmt.Scanln(&choice)

			// If user chose to logout, delete tokens and account info from Redis
			if choice == "Y" || choice == "y" {
				redis.DelToken("LoginToken")
				redis.DelToken("AccountEmail")
				redis.DelToken("AccountPassword")
				redis.DelToken("AccountName")

				redis.DelToken("MemberLoginToken")
				redis.DelToken("MemberAccountEmail")
				redis.DelToken("MemberAccountPassword")
				redis.DelToken("MemberAccountName")
				fmt.Println("You're successfully logged out.")
			} else if choice == "N" || choice == "n" {
				// If user chose not to logout, print a message
				fmt.Println("You're not logged out.")
			} else {
				// If user entered invalid input, print an error
				fmt.Println(errors.New("invalid choice: enter the correct one."))
			}
		} else {
			// If there were no tokens (user wasn't logged in), print an error
			fmt.Println(errors.New("You're already logged out."))
		}
	},
}

// Add the logout command to the root command
func init() {
	cmd.RootCmd.AddCommand(logoutCmd)
}
