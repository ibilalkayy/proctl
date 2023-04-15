package workspace

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
)

// deletespaceCmd represents the deletespace command
var deletespaceCmd = &cobra.Command{
	Use:   "deletespace",
	Short: "Delete a workspace",
	Run: func(cmd *cobra.Command, args []string) {
		deleteWorkspaceName, _ := cmd.Flags().GetString("name")
		accountEmail := redis.GetAccountInfo("AccountEmail")
		loginToken := redis.GetAccountInfo("LoginToken")
		values := [2]string{accountEmail, deleteWorkspaceName}
		oldWorkspaceName := mysql.FindWorkspaceName(accountEmail, deleteWorkspaceName)

		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			if deleteWorkspaceName == oldWorkspaceName {
				mysql.DeleteWorkspace(values)
				fmt.Println("Your workspace is successfully deleted")
			} else {
				fmt.Println(errors.New("A workspace is not present by this name"))
			}
		} else {
			fmt.Println(errors.New("First login to delete a workspace"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(deletespaceCmd)
	deletespaceCmd.Flags().StringP("name", "n", "", "Specify the name of the workspace to delete it")
}
