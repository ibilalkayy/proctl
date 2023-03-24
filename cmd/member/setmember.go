package member

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/spf13/cobra"
)

// setmemCmd represents the setmem command
var setmemCmd = &cobra.Command{
	Use:   "setmem",
	Short: "Setup the member credentials",
	Run: func(cmd *cobra.Command, args []string) {
		loginToken := redis.GetAccountInfo("LoginToken")
		memberEmail, _ := cmd.Flags().GetString("email")
		memberPassword, _ := cmd.Flags().GetString("password")
		memberFullName, _ := cmd.Flags().GetString("full name")
		memberAccountName, _ := cmd.Flags().GetString("account name")
		memberFound := mysql.FindMember(memberEmail)

		if len(loginToken) == 0 {
			if memberFound && len(memberEmail) != 0 {
				if len(memberPassword) != 0 {
					redis.SetAccountInfo("MemberPassword", memberPassword)
					fullName := redis.GetAccountInfo("MemberFullName")
					accountName := redis.GetAccountInfo("MemberAccountName")
					values := [3]string{memberPassword, fullName, accountName}
					mysql.UpdateMember(values, memberEmail)
				}
				if len(memberFullName) != 0 {
					redis.SetAccountInfo("MemberFullName", memberFullName)
					password := redis.GetAccountInfo("MemberPassword")
					accountName := redis.GetAccountInfo("MemberAccountName")
					values := [3]string{password, memberFullName, accountName}
					mysql.UpdateMember(values, memberEmail)
				}
				if len(memberAccountName) != 0 {
					redis.SetAccountInfo("MemberAccount", memberAccountName)
					password := redis.GetAccountInfo("MemberPassword")
					fullName := redis.GetAccountInfo("MemberFullName")
					values := [3]string{password, fullName, memberAccountName}
					mysql.UpdateMember(values, memberEmail)
				}
				fmt.Println("The member credentials are successfully updated")
			} else {
				fmt.Println(errors.New("Please enter the email address or type 'proctl setmem --help'"))
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
}
