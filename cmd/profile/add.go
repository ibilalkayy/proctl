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

// addProfileCmd represents the addProfile command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add the profile information of a user",
	Run: func(cmd *cobra.Command, args []string) {
		profileTitle, _ := cmd.Flags().GetString("title")
		profilePhone, _ := cmd.Flags().GetString("phone")
		profileLocation, _ := cmd.Flags().GetString("location")
		profileWorkingStatus, _ := cmd.Flags().GetString("working status")

		loginToken := redis.GetAccountInfo("LoginToken")
		AccountEmail := redis.GetAccountInfo("AccountEmail")
		profileFound := mysql.FindProfile(AccountEmail)
		if len(loginToken) != 0 && jwt.RefreshToken() {
			if (len(profileTitle) != 0 || len(profilePhone) != 0 || len(profileLocation) != 0 || len(profileWorkingStatus) != 0) && !profileFound {
				profileData := [5]string{AccountEmail, profileTitle, profilePhone, profileLocation, profileWorkingStatus}
				redis.SetAccountInfo("ProfileTitle", profileTitle)
				redis.SetAccountInfo("ProfilePhone", profilePhone)
				redis.SetAccountInfo("ProfileLocation", profileLocation)
				redis.SetAccountInfo("ProfileWorkingStatus", profileWorkingStatus)
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

func init() {
	cmd.RootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("title", "t", "", "Specify an account title")
	addCmd.Flags().StringP("phone", "p", "", "Specify an account phone number")
	addCmd.Flags().StringP("location", "l", "", "Specify an account location")
	addCmd.Flags().StringP("working status", "w", "", "Specify an account working status")
}
