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
var updatememCmd = &cobra.Command{
	Use:   "updatemem",
	Short: "Update the member credentials",
	Run: func(cmd *cobra.Command, args []string) {
		memberTitle, _ := cmd.Flags().GetString("title")
		memberPhone, _ := cmd.Flags().GetString("phone")
		memberLocation, _ := cmd.Flags().GetString("location")
		memberWorkingStatus, _ := cmd.Flags().GetString("working status")

		loginToken := redis.GetAccountInfo("LoginToken")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")
		memberAccountEmail := redis.GetAccountInfo("MemberAccountEmail")
		memberAccountPassword := redis.GetAccountInfo("MemberAccountPassword")
		_, memberFound := mysql.FindMember(memberAccountEmail, memberAccountPassword)

		if len(loginToken) != 0 && jwt.RefreshToken("user") && memberFound {
			fmt.Println(errors.New("Can't update the member credentials as an admin user"))
		} else if len(memberLoginToken) != 0 && jwt.RefreshToken("member") && memberFound {
			memberValues := [4]string{memberTitle, memberPhone, memberLocation, memberWorkingStatus}
			mysql.UpdateMember(memberValues, memberAccountEmail, memberAccountPassword)
			fmt.Println("You have successfully updated the member credentials")
		} else {
			fmt.Println(errors.New("First login to update the member credentials"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(updatememCmd)
	updatememCmd.Flags().StringP("title", "t", "", "Specify a member account title to update")
	updatememCmd.Flags().StringP("phone", "p", "", "Specify a member account phone number to update")
	updatememCmd.Flags().StringP("location", "l", "", "Specify a member account location to update")
	updatememCmd.Flags().StringP("working status", "w", "", "Specify a member account working status to update")
}
