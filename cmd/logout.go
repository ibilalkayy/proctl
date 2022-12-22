package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var choice string

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Confirm it to logout",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Want to logout [y/n]: ")
		fmt.Scanln(&choice)

		if choice == "Y" || choice == "y" {
			fmt.Println("You're successfully logged out")
		} else if choice == "N" || choice == "n" {
			fmt.Println("You're not logged out")
		} else {
			fmt.Println(errors.New("invalid choice: enter the correct one"))
		}
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
