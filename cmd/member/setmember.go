package member

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/spf13/cobra"
)

// setmemberCmd represents the setmember command
var setmemberCmd = &cobra.Command{
	Use:   "setmember",
	Short: "Setup the member credentials",
	Run: func(cmd *cobra.Command, args []string) {
		loginToken := redis.GetAccountInfo("LoginToken")
		memberEmail, _ := cmd.Flags().GetString("email")
		memberPassword, _ := cmd.Flags().GetString("password")
		memberFullName, _ := cmd.Flags().GetString("full name")
		memberAccountName, _ := cmd.Flags().GetString("account name")

		if len(loginToken) == 0 {
			foundMember := mysql.FindMember(memberEmail)
			if foundMember {
				values := [4]string{memberPassword, memberFullName, memberAccountName, memberEmail}
				mysql.UpdateMember(values)
				fmt.Println("The member credentials are successfully updated")
			} else {
				fmt.Println(errors.New("The member credentials are not updated"))
			}
		} else {
			fmt.Println(errors.New("First logout to setup the member credentials"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(setmemberCmd)
	setmemberCmd.Flags().StringP("email", "e", "", "Specify an email address to setup the credentials")
	setmemberCmd.Flags().StringP("password", "p", "", "Specify a password to setup the credentials")
	setmemberCmd.Flags().StringP("full name", "f", "", "Specify a full name to setup the credentials")
	setmemberCmd.Flags().StringP("account name", "a", "", "Specify an account name to setup the credentials")
}
