package member

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
)

// updatememCmd represents the updatemem command
// This is the command that will be used to update member credentials.
var updatememCmd = &cobra.Command{
	Use:   "updatemem",                     // This is the command name.
	Short: "Update the member credentials", // A short description of the command.

	// This function is run when the command is called.
	Run: func(cmd *cobra.Command, args []string) {
		memberEmail, _ := cmd.Flags().GetString("email")
		memberPassword, _ := cmd.Flags().GetString("password")
		memberFullName, _ := cmd.Flags().GetString("full name")
		memberAccountName, _ := cmd.Flags().GetString("account name")
		memberTitle, _ := cmd.Flags().GetString("title")
		memberPhone, _ := cmd.Flags().GetString("phone")
		memberLocation, _ := cmd.Flags().GetString("location")
		memberWorkingStatus, _ := cmd.Flags().GetString("working status")

		// Get login tokens and account info from Redis, and find the member in MySQL.
		loginToken := redis.GetAccountInfo("LoginToken")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")
		memberAccountEmail := redis.GetAccountInfo("MemberAccountEmail")
		memberAccountPassword := redis.GetAccountInfo("MemberAccountPassword")
		_, memberFound := mysql.FindMember(memberAccountEmail, memberAccountPassword)

		// Check login status and refresh JWT if necessary.
		// If logged in as a user, show an error. If logged in as a member, update member credentials.
		if len(loginToken) != 0 && jwt.RefreshToken("user") && memberFound {
			fmt.Println(errors.New("Can't update the member credentials as an admin user"))
		} else if len(memberLoginToken) != 0 && jwt.RefreshToken("member") && memberFound {
			memberValues := [8]string{memberEmail, memberPassword, memberFullName, memberAccountName, memberTitle, memberPhone, memberLocation, memberWorkingStatus}
			mysql.UpdateMember(memberValues, memberAccountEmail, memberAccountPassword, false)
			fmt.Println("You have successfully updated the member credentials")
		} else {
			fmt.Println(errors.New("First login to update the member credentials"))
		}
	},
}

// init is a special Go function that is automatically executed at the start of the program.
func init() {
	// This function adds updatememCmd to the list of commands in the root command.
	cmd.RootCmd.AddCommand(updatememCmd)
	updatememCmd.Flags().StringP("email", "e", "", "Specify an email address to update")
	updatememCmd.Flags().StringP("password", "p", "", "Specify a password to update")
	updatememCmd.Flags().StringP("full name", "f", "", "Specify a full name to update")
	updatememCmd.Flags().StringP("account name", "a", "", "Specify an account name to update")
	updatememCmd.Flags().StringP("title", "t", "", "Specify a member account title to update")
	updatememCmd.Flags().StringP("phone", "n", "", "Specify a member account phone number to update")
	updatememCmd.Flags().StringP("location", "l", "", "Specify a member account location to update")
	updatememCmd.Flags().StringP("working status", "w", "", "Specify a member account working status to update")
}
