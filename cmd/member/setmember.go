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

// setmemCmd represents the setmem command
var setmemCmd = &cobra.Command{
	Use:   "setmem",
	Short: "Setup the member credentials",
	Run: func(cmd *cobra.Command, args []string) {
		loginToken := redis.GetAccountInfo("LoginToken")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")
		memberEmail, _ := cmd.Flags().GetString("email")
		memberPassword, _ := cmd.Flags().GetString("password")
		memberFullName, _ := cmd.Flags().GetString("full name")
		memberAccountName, _ := cmd.Flags().GetString("account name")
		memberTitle, _ := cmd.Flags().GetString("title")
		memberPhone, _ := cmd.Flags().GetString("phone")
		memberLocation, _ := cmd.Flags().GetString("location")
		memberWorkingStatus, _ := cmd.Flags().GetString("working status")
		memberFound := mysql.FindMember(memberEmail)

		if len(loginToken) == 0 && len(memberLoginToken) == 0 {
			tokenString, jwtTokenGenerated := jwt.GenerateJWT()
			if jwtTokenGenerated {
				redis.SetAccountInfo("MemberLoginToken", tokenString)
				redis.SetAccountInfo("MemberAccountName", memberAccountName)
				if memberFound && len(memberEmail) != 0 {
					values := [7]string{memberPassword, memberFullName, memberAccountName, memberTitle, memberPhone, memberLocation, memberWorkingStatus}
					mysql.UpdateMember(values, memberEmail)
					fmt.Println("The member credentials are successfully updated")
				} else {
					fmt.Println(errors.New("Please enter the email address or type 'proctl setmem --help'"))
				}
			} else {
				fmt.Println(errors.New("Failure in setting up a member"))
			}
		} else {
			fmt.Println(errors.New("First logout to setup the member credentials"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(setmemCmd)
	setmemCmd.Flags().StringP("email", "e", "", "Specify an email address to setup the credentials")
	setmemCmd.Flags().StringP("password", "p", "", "Specify a password to setup the credentials")
	setmemCmd.Flags().StringP("full name", "f", "", "Specify a full name to setup the credentials")
	setmemCmd.Flags().StringP("account name", "a", "", "Specify an account name to setup the credentials")
	setmemCmd.Flags().StringP("title", "t", "", "Specify an account title to setup the credentials")
	setmemCmd.Flags().StringP("phone", "n", "", "Specify an account phone number to setup the credentials")
	setmemCmd.Flags().StringP("location", "l", "", "Specify an account location to setup the credentials")
	setmemCmd.Flags().StringP("working status", "w", "", "Specify an account working status to setup the credentials")
}
