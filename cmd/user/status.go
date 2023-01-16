package cmd

import (
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Status of the logged in or the logged out user.",
	Run: func(cmd *cobra.Command, args []string) {
		loginToken := redis.GetAccountInfo("LoginToken")
		if len(loginToken) != 0 && jwt.RefreshToken() {
			accountName := redis.GetAccountInfo("AccountName")
			fmt.Printf("%s is logged in.\n", accountName)
		} else {
			fmt.Println("User is logged out.")
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(statusCmd)
}
