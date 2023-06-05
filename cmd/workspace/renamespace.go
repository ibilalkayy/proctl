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

// renamespaceCmd represents the renamespace command
var renamespaceCmd = &cobra.Command{
	Use:   "renamespace",
	Short: "Rename an existing workspace",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the old and new workspace names from the command-line flags
		oldWorkspaceName, _ := cmd.Flags().GetString("oldname")
		newWorkspaceName, _ := cmd.Flags().GetString("newname")

		// Retrieve account information from Redis
		accountEmail := redis.GetAccountInfo("AccountEmail")
		loginToken := redis.GetAccountInfo("LoginToken")
		memberEmail := redis.GetAccountInfo("MemberEmail")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		var email string
		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			email = accountEmail
		} else if len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			email = memberEmail
		} else {
			// If no valid login token is found, print an error message
			fmt.Println(errors.New("First login to rename a workspace"))
			return
		}

		// Find the actual workspace name from the database
		foundWorkspaceName := mysql.FindWorkspaceName(email, oldWorkspaceName)

		// Create an array with the new workspace name, email, and old workspace name to pass to the database update function
		values := [3]string{newWorkspaceName, email, oldWorkspaceName}

		if oldWorkspaceName == foundWorkspaceName {
			// If the old workspace name matches the actual workspace name, update the workspace name in the database
			mysql.UpdateWorkspace(values)
			fmt.Println("Your workspace is successfully renamed")
		} else {
			// If no matching workspace is found, print an error message
			fmt.Println(errors.New("A workspace is not present by this name"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(renamespaceCmd)
	// Add command-line flags to specify the old and new names of the workspace
	renamespaceCmd.Flags().StringP("oldname", "o", "", "Specify the name of an existing workspace")
	renamespaceCmd.Flags().StringP("newname", "n", "", "Specify the new name of a workspace")
}
