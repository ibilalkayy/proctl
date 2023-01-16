package cmd

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/spf13/cobra"
)

var choice string

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Confirm it to logout",
	Run: func(cmd *cobra.Command, args []string) {
		loginToken := redis.GetAccountInfo("LoginToken")
		if len(loginToken) != 0 {
			fmt.Printf("Want to logout [y/n]: ")
			fmt.Scanln(&choice)

			if choice == "Y" || choice == "y" {
				redis.DelToken("LoginToken")
				fmt.Println("You're successfully logged out.")
			} else if choice == "N" || choice == "n" {
				fmt.Println("You're not logged out.")
			} else {
				fmt.Println(errors.New("invalid choice: enter the correct one."))
			}
		} else {
			fmt.Println(errors.New("You're already logged out."))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(logoutCmd)
}
