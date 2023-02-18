package cmd

import (
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/spf13/cobra"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename an existing workspace",
	Run: func(cmd *cobra.Command, args []string) {
		renameWorkspace, _ := cmd.Flags().GetString("rename")
		fmt.Println(renameWorkspace)
	},
}

func init() {
	cmd.RootCmd.AddCommand(renameCmd)
	renameCmd.Flags().StringP("rename", "r", "", "Specify the name to rename a workspace")
}
