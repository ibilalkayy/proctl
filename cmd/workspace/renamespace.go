package cmd

import (
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/spf13/cobra"
)

// renamespaceCmd represents the renamespace command
var renamespaceCmd = &cobra.Command{
	Use:   "renamespace",
	Short: "Rename an existing workspace",
	Run: func(cmd *cobra.Command, args []string) {
		renameWorkspace, _ := cmd.Flags().GetString("name")
		fmt.Println(renameWorkspace)
	},
}

func init() {
	cmd.RootCmd.AddCommand(renamespaceCmd)
	renamespaceCmd.Flags().StringP("name", "n", "", "Specify the name to rename an existing workspace")
}
