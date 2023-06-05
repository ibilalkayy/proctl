// Package profile contains functionality related to managing user profiles.
package profile

// Import necessary packages and libraries.
import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
)

// addProfileCmd represents the addProfile command
// This command is used to add user profile information.
var addCmd = &cobra.Command{
	Use:   "add",                                   // This is the command name.
	Short: "Add the profile information of a user", // A short description of the command.

	Run: func(cmd *cobra.Command, args []string) { // This function is run when the command is called.
		// Get the values of the flags (title, phone, location, working status).
		profileTitle, _ := cmd.Flags().GetString("title")
		profilePhone, _ := cmd.Flags().GetString("phone")
		profileLocation, _ := cmd.Flags().GetString("location")
		profileWorkingStatus, _ := cmd.Flags().GetString("working status")

		// Get login tokens and account info from Redis, and find the profile in MySQL.
		loginToken := redis.GetAccountInfo("LoginToken")
		AccountEmail := redis.GetAccountInfo("AccountEmail")
		profileFound := mysql.FindProfile(AccountEmail)

		// Check login status and refresh JWT if necessary.
		// If logged in and profile info is not empty, insert the profile data.
		// If profile data is already inserted, show an error message.
		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			if (len(profileTitle) != 0 || len(profilePhone) != 0 || len(profileLocation) != 0 || len(profileWorkingStatus) != 0) && !profileFound {
				profileData := [5]string{AccountEmail, profileTitle, profilePhone, profileLocation, profileWorkingStatus}
				mysql.InsertProfileData(profileData)
				fmt.Println("Your profile data is successfully inserted.")
			} else if profileFound {
				fmt.Println("Your profile data is already inserted. Type 'proctl update [flags]'")
			} else {
				fmt.Println(errors.New("Give the flags to insert the profile information."))
			}
		} else {
			fmt.Println(errors.New("First login to add the profile information."))
		}
	},
}

// init is a special Go function that is automatically executed at the start of the program.
func init() {
	// This function adds addCmd to the list of commands in the root command.
	cmd.RootCmd.AddCommand(addCmd)

	// This function adds flags to addCmd for specifying profile information.
	addCmd.Flags().StringP("title", "t", "", "Specify an account title")
	addCmd.Flags().StringP("phone", "p", "", "Specify an account phone number")
	addCmd.Flags().StringP("location", "l", "", "Specify an account location")
	addCmd.Flags().StringP("working status", "w", "", "Specify an account working status")
}
