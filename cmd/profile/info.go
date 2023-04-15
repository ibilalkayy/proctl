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

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show the profile information",
	Run: func(cmd *cobra.Command, args []string) {
		loginToken := redis.GetAccountInfo("LoginToken")
		accountEmail := redis.GetAccountInfo("AccountEmail")
		accountPassword := redis.GetAccountInfo("AccountPassword")
		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			mysql.ListUserInfo(accountEmail, accountPassword)
			mysql.ListProfileInfo(accountEmail)
		} else {
			fmt.Println(errors.New("First login to show the profile information."))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(infoCmd)
}
