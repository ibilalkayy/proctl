package member

import (
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
)

// depsCmd represents the deps command
var depsCmd = &cobra.Command{
	Use:   "deps",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		department, _ := cmd.Flags().GetString("department")

		accountEmail := redis.GetAccountInfo("AccountEmail")
		loginToken := redis.GetAccountInfo("LoginToken")
		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			redis.DelToken("MemberLoginToken")
			mysql.InsertDepartment(accountEmail, department)
			fmt.Println("Your department is successfully added")
			return
		}

		memberAccountEmail := redis.GetAccountInfo("MemberAccountEmail")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")
		if len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			redis.DelToken("LoginToken")
			mysql.InsertDepartment(memberAccountEmail, department)
			fmt.Println("Your department is successfully added")
			return
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(depsCmd)
	depsCmd.Flags().StringP("department", "d", "", "Specify the department to setup")
}
