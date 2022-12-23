package cmd

import (
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/spf13/cobra"
)

// signupCmd represents the signup command
var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "Provide an email address, password, full name and account name in order to signup",
	Run: func(cmd *cobra.Command, args []string) {
		signupEmail, _ := cmd.Flags().GetString("email")
		signupPassword, _ := cmd.Flags().GetString("password")
		signupFullName, _ := cmd.Flags().GetString("full name")
		signupAccountName, _ := cmd.Flags().GetString("account name")

		signupCredentials := [4]string{signupEmail, signupPassword, signupFullName, signupAccountName}
		mysql.InsertSignupData(signupCredentials)
	},
}

func init() {
	rootCmd.AddCommand(signupCmd)
	signupCmd.Flags().StringP("email", "e", "", "Specify an email address to signup")
	signupCmd.Flags().StringP("password", "p", "", "Specify the password to signup")
	signupCmd.Flags().StringP("full name", "f", "", "Specify the full name to signup")
	signupCmd.Flags().StringP("account name", "a", "", "Specify the account name to signup")
}
