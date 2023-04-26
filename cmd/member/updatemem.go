package member

import (
	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/spf13/cobra"
)

// updatememCmd represents the updatemem command
var updatememCmd = &cobra.Command{
	Use:   "updatemem",
	Short: "Update the member credentials",
	Run: func(cmd *cobra.Command, args []string) {
		memberEmail, _ := cmd.Flags().GetString("email")
		// memberPassword, _ := cmd.Flags().GetString("password")
		memberTitle, _ := cmd.Flags().GetString("title")
		memberPhone, _ := cmd.Flags().GetString("phone")
		memberLocation, _ := cmd.Flags().GetString("location")
		memberWorkingStatus, _ := cmd.Flags().GetString("working status")
		memberValues := [4]string{memberTitle, memberPhone, memberLocation, memberWorkingStatus}
		mysql.UpdateMember(memberValues, memberEmail, "")
	},
}

func init() {
	cmd.RootCmd.AddCommand(updatememCmd)
	updatememCmd.Flags().StringP("email", "e", "", "Specify a member account email to verify")
	// updatememCmd.Flags().StringP("password", "p", "", "Specify a member account password to verify")
	updatememCmd.Flags().StringP("title", "t", "", "Specify a member account title to update")
	updatememCmd.Flags().StringP("phone", "n", "", "Specify a member account phone number to update")
	updatememCmd.Flags().StringP("location", "l", "", "Specify a member account location to update")
	updatememCmd.Flags().StringP("working status", "w", "", "Specify a member account working status to update")
}
