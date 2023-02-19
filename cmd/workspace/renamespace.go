package cmd

import (
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/spf13/cobra"
)

// renamespaceCmd represents the renamespace command
var renamespaceCmd = &cobra.Command{
	Use:   "renamespace",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		renameWorkspace, _ := cmd.Flags().GetString("name")
		fmt.Println(renameWorkspace)
	},
}

func init() {
	cmd.RootCmd.AddCommand(renamespaceCmd)
	renamespaceCmd.Flags().StringP("name", "n", "", "Specify the name to add a workspace")
}