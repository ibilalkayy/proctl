package cmd

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the user profile information",
	Run: func(cmd *cobra.Command, args []string) {
		profileTitle, _ := cmd.Flags().GetString("title")
		profilePhone, _ := cmd.Flags().GetString("phone")
		profileLocation, _ := cmd.Flags().GetString("location")
		profileWorkingStatus, _ := cmd.Flags().GetString("working status")

		if len(profileTitle) != 0 || len(profilePhone) != 0 || len(profileLocation) != 0 || len(profileWorkingStatus) != 0 {
			loginToken := redis.GetAccountInfo("LoginToken")
			if len(loginToken) != 0 && jwt.RefreshToken() {
				profileData := [4]string{profileTitle, profilePhone, profileLocation, profileWorkingStatus}
				AccountEmail := redis.GetAccountInfo("AccountEmail")
				mysql.UpdateProfile(profileData, AccountEmail)
				fmt.Println("Your profile data is successfully updated.")
			} else {
				fmt.Println(errors.New("First login to add the profile information."))
			}
		} else {
			fmt.Println(errors.New("Give the flags to update the profile information."))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("title", "t", "", "Specify an account title to update")
	updateCmd.Flags().StringP("phone", "p", "", "Specify an account phone number to update")
	updateCmd.Flags().StringP("location", "l", "", "Specify an account location to update")
	updateCmd.Flags().StringP("working status", "w", "", "Specify an account working status to update")
	updateCmd.Flags().StringP("full name", "f", "", "Specify the full name to update")
	updateCmd.Flags().StringP("account name", "a", "", "Specify the account name to update")
}
