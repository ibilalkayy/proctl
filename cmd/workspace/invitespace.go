package cmd

import (
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/spf13/cobra"
)

// invitespaceCmd represents the invitespace command
var invitespaceCmd = &cobra.Command{
	Use:   "invitespace",
	Short: "Invite other members in the workspace",
	Run: func(cmd *cobra.Command, args []string) {
		inviteWorkspace, _ := cmd.Flags().GetString("email")
		fmt.Println(inviteWorkspace)
	},
}

func init() {
	cmd.RootCmd.AddCommand(invitespaceCmd)
	invitespaceCmd.Flags().StringP("email", "e", "", "Specify an email address to invite people")
}
