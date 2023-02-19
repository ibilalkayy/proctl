package cmd

import (
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/spf13/cobra"
)

// deletespaceCmd represents the deletespace command
var deletespaceCmd = &cobra.Command{
	Use:   "deletespace",
	Short: "Delete a workspace",
	Run: func(cmd *cobra.Command, args []string) {
		deleteWorkspace, _ := cmd.Flags().GetString("name")
		fmt.Println(deleteWorkspace)
	},
}

func init() {
	cmd.RootCmd.AddCommand(deletespaceCmd)
	deletespaceCmd.Flags().StringP("name", "n", "", "Specify the name of the workspace to delete it")
}
