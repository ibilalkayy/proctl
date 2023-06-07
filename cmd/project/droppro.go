package project

import (
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/spf13/cobra"
)

// dropproCmd represents the droppro command
var dropproCmd = &cobra.Command{
	Use:   "droppro",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("droppro called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(dropproCmd)
}
