package cmd

import (
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a workspace",
	Run: func(cmd *cobra.Command, args []string) {
		deleteWorkspace, _ := cmd.Flags().GetString("name")
		fmt.Println(deleteWorkspace)
	},
}

func init() {
	cmd.RootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("name", "n", "", "Specify the name of workspace to delete it")
}
