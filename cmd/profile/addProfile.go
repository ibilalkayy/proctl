package cmd

import (
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/spf13/cobra"
)

// addProfileCmd represents the addProfile command
var addProfileCmd = &cobra.Command{
	Use:   "add",
	Short: "Add the profile information of a user",
	Run: func(cmd *cobra.Command, args []string) {
		profileTitle, _ := cmd.Flags().GetString("title")
		profilePhone, _ := cmd.Flags().GetString("phone")
		profileLocation, _ := cmd.Flags().GetString("location")
		profileWorkingStatus, _ := cmd.Flags().GetString("working status")
		fmt.Println(profileTitle)
		fmt.Println(profilePhone)
		fmt.Println(profileLocation)
		fmt.Println(profileWorkingStatus)
	},
}

func init() {
	cmd.RootCmd.AddCommand(addProfileCmd)
	addProfileCmd.Flags().StringP("title", "t", "", "Specify an account title")
	addProfileCmd.Flags().StringP("phone", "p", "", "Specify an account phone number")
	addProfileCmd.Flags().StringP("location", "l", "", "Specify an account location")
	addProfileCmd.Flags().StringP("working status", "w", "", "Specify an account working status")
}
