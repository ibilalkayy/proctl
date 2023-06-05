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
		// Get the workspace name from the command-line flag
		deleteWorkspaceName, _ := cmd.Flags().GetString("name")

		// Retrieve account information from Redis
		accountEmail := redis.GetAccountInfo("AccountEmail")
		memberEmail := redis.GetAccountInfo("MemberEmail")
		loginToken := redis.GetAccountInfo("LoginToken")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		var email string
		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			email = accountEmail
		} else if len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			email = memberEmail
		} else {
			// If no valid login token is found, print an error message
			fmt.Println(errors.New("First login to delete a workspace"))
			return
		}

		// Create an array with email and workspace name to pass to the database delete function
		values := [2]string{email, deleteWorkspaceName}

		// Find the actual workspace name from the database
		oldWorkspaceName := mysql.FindWorkspaceName(email, deleteWorkspaceName)

		if deleteWorkspaceName == oldWorkspaceName {
			// Delete the workspace from the database
			mysql.DeleteWorkspace(values)
			fmt.Println("Your workspace is successfully deleted")
		} else {
			// If the provided workspace name does not match the actual workspace name, print an error message
			fmt.Println(errors.New("A workspace is not present by this name"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(deletespaceCmd)
	// Add a command-line flag to specify the name of the workspace to delete
	deletespaceCmd.Flags().StringP("name", "n", "", "Specify the name of the workspace to delete it")
}
