package cmd

import (
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/spf13/cobra"
)

// browsespaceCmd represents the browsespace command
var browsespaceCmd = &cobra.Command{
	Use:   "browsespace",
	Short: "Browse all the workspaces",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("workspaces list")
	},
}

func init() {
	cmd.RootCmd.AddCommand(browsespaceCmd)
}
