package work

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
)

// mydepCmd represents the mydep command
var mydepCmd = &cobra.Command{
	Use:   "mydep",
	Short: "Show you the department",
	Run: func(cmd *cobra.Command, args []string) {
		accountEmail := redis.GetAccountInfo("AccountEmail")
		loginToken := redis.GetAccountInfo("LoginToken")

		memberAccountEmail := redis.GetAccountInfo("MemberAccountEmail")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		userCredentials := mysql.FindDepartment(accountEmail)
		memberCredentials := mysql.FindDepartment(memberAccountEmail)

		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			if len(userCredentials[0]) != 0 && len(userCredentials[1]) != 0 {
				fmt.Printf("My department: %s\n", userCredentials[1])
				return
			} else {
				fmt.Println(errors.New("This user has no department"))
			}
		} else if len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			if len(memberCredentials[0]) != 0 && len(memberCredentials[1]) != 0 {
				fmt.Printf("My department: %s\n", memberCredentials[1])
				return
			} else {
				fmt.Println(errors.New("This member has no department"))
			}
		} else {
			fmt.Println(errors.New("First login to show you the department"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(mydepCmd)
}
