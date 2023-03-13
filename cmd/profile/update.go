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

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the user profile information",
	Run: func(cmd *cobra.Command, args []string) {
		profileTitle, _ := cmd.Flags().GetString("title")
		profilePhone, _ := cmd.Flags().GetString("phone")
		profileLocation, _ := cmd.Flags().GetString("location")
		profileWorkingStatus, _ := cmd.Flags().GetString("working status")
		profileFullName, _ := cmd.Flags().GetString("full name")
		profileAccountName, _ := cmd.Flags().GetString("account name")
		profilePassword, _ := cmd.Flags().GetString("password")

		loginToken := redis.GetAccountInfo("LoginToken")
		accountEmail := redis.GetAccountInfo("AccountEmail")
		accountPassword := redis.GetAccountInfo("AccountPassword")
		verificationCode := redis.GetAccountInfo("VerificationCode")

		_, _, mysqlStatus, _ := mysql.FindAccount(accountEmail, accountPassword)

		if len(loginToken) != 0 && jwt.RefreshToken() {
			updateProfile := len(profileTitle) != 0 || len(profilePhone) != 0 || len(profileLocation) != 0 || len(profileWorkingStatus) != 0
			updateUser := len(profileFullName) != 0 || len(profileAccountName) != 0 || len(profilePassword) != 0

			if !updateProfile && !updateUser {
				fmt.Println(errors.New("Give the flags to update the profile information."))
			} else {
				if updateProfile {
					profileData := [4]string{profileTitle, profilePhone, profileLocation, profileWorkingStatus}
					mysql.UpdateProfile(profileData, accountEmail)
					fmt.Println("Your profile data is successfully updated.")
				}

				if updateUser {
					if mysqlStatus == "0" && len(verificationCode) != 0 {
						userData := [4]string{profileFullName, profileAccountName, profilePassword, "0"}
						mysql.UpdateUser(userData, accountEmail, accountPassword)
					} else if mysqlStatus == "1" && len(verificationCode) == 0 {
						userData := [4]string{profileFullName, profileAccountName, profilePassword, "1"}
						mysql.UpdateUser(userData, accountEmail, accountPassword)
					}
					fmt.Println("Your profile data is successfully updated.")
				}
			}
		} else {
			fmt.Println(errors.New("First login to add the profile information."))
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
	updateCmd.Flags().StringP("password", "s", "", "Specify the password to update")
}
