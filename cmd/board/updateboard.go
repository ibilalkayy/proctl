package board

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
)

// updateboardCmd represents the updateboard command
var updateboardCmd = &cobra.Command{
	Use:   "updateboard",
	Short: "Update the board",
	Run: func(cmd *cobra.Command, args []string) {
		loginToken := redis.GetAccountInfo("LoginToken")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			fmt.Println("user is logged in")
		} else if len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			fmt.Println("member is logged in")
		} else {
			fmt.Println(errors.New("First login to update the board"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(updateboardCmd)
}
