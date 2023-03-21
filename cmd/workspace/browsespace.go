package workspace

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
)

// browsespaceCmd represents the browsespace command
var browsespaceCmd = &cobra.Command{
	Use:   "browsespace",
	Short: "Browse all the workspaces",
	Run: func(cmd *cobra.Command, args []string) {
		loginToken := redis.GetAccountInfo("LoginToken")
		accountEmail := redis.GetAccountInfo("AccountEmail")

		if len(loginToken) != 0 && jwt.RefreshToken() {
			name := mysql.FindWorkspace(accountEmail)
			fmt.Println(name)
		} else {
			fmt.Println(errors.New("First login to add a new workspace"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(browsespaceCmd)
}
