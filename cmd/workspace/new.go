package cmd

import (
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Add a new workspace",
	Run: func(cmd *cobra.Command, args []string) {
		addWorkspace, _ := cmd.Flags().GetString("name")
		fmt.Println(addWorkspace)
	},
}

func init() {
	cmd.RootCmd.AddCommand(newCmd)
	newCmd.Flags().StringP("name", "n", "", "Specify the name to add a workspace")
}
