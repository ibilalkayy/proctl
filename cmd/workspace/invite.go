package cmd

import (
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/spf13/cobra"
)

// inviteCmd represents the invite command
var inviteCmd = &cobra.Command{
	Use:   "invite",
	Short: "Invite other members in the workspace",
	Run: func(cmd *cobra.Command, args []string) {
		inviteWorkspace, _ := cmd.Flags().GetString("email")
		fmt.Println(inviteWorkspace)
	},
}

func init() {
	cmd.RootCmd.AddCommand(inviteCmd)
	inviteCmd.Flags().StringP("email", "e", "", "Specify an email address to invite people")
}
