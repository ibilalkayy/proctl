package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Provide an email address and password in order to login",
	Run: func(cmd *cobra.Command, args []string) {
		loginEmail, _ := cmd.Flags().GetString("email")
		loginPassword, _ := cmd.Flags().GetString("password")
		fmt.Println(loginEmail)
		fmt.Println(loginPassword)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("email", "e", "", "Specify an email address to login")
	loginCmd.Flags().StringP("password", "p", "", "Specify the password to login")
}
