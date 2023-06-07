package project

import (
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/spf13/cobra"
)

// updateproCmd represents the updatepro command
var updateproCmd = &cobra.Command{
	Use:   "updatepro",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updatepro called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(updateproCmd)
}
