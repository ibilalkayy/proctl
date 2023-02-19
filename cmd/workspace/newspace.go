package cmd

import (
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/spf13/cobra"
)

// newspaceCmd represents the newspace command
var newspaceCmd = &cobra.Command{
	Use:   "newspace",
	Short: "Add a new workspace",
	Run: func(cmd *cobra.Command, args []string) {
		newWorkspace, _ := cmd.Flags().GetString("name")
		fmt.Println(newWorkspace)
	},
}

func init() {
	cmd.RootCmd.AddCommand(newspaceCmd)
	newspaceCmd.Flags().StringP("name", "n", "", "Specify the name to add a new workspace")
}
