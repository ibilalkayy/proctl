package cmd

import (
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/spf13/cobra"
)

// browseCmd represents the browse command
var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "Give all the workspaces",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("workspaces list")
	},
}

func init() {
	cmd.RootCmd.AddCommand(browseCmd)
}
