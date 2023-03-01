package cmd

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
)

// renamespaceCmd represents the renamespace command
var renamespaceCmd = &cobra.Command{
	Use:   "renamespace",
	Short: "Rename an existing workspace",
	Run: func(cmd *cobra.Command, args []string) {
		oldWorkspaceName, _ := cmd.Flags().GetString("oldname")
		newWorkspaceName, _ := cmd.Flags().GetString("newname")
		accountEmail := redis.GetAccountInfo("AccountEmail")
		loginToken := redis.GetAccountInfo("LoginToken")
		foundWorkspaceName := mysql.FindWorkspaceName(accountEmail, oldWorkspaceName)
		values := [3]string{newWorkspaceName, accountEmail, oldWorkspaceName}

		if len(loginToken) != 0 && jwt.RefreshToken() {
			if oldWorkspaceName == foundWorkspaceName {
				mysql.UpdateWorkspace(values)
				fmt.Println("Your workspace is successfully renamed")
			} else {
				fmt.Println(errors.New("A workspace is not present by this name"))
			}
		} else {
			fmt.Println(errors.New("First login to rename a workspace"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(renamespaceCmd)
	renamespaceCmd.Flags().StringP("oldname", "o", "", "Specify the name of an existing workspace")
	renamespaceCmd.Flags().StringP("newname", "n", "", "Specify the new name of a workspace")
}
