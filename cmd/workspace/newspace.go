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

// newspaceCmd represents the newspace command
var newspaceCmd = &cobra.Command{
	Use:   "newspace",
	Short: "Add a new workspace",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the new workspace name from the command-line flag
		newWorkspaceName, _ := cmd.Flags().GetString("name")

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
			fmt.Println(errors.New("First login to add a new workspace"))
			return
		}

		// Create an array with email and new workspace name to pass to the database insert function
		values := [2]string{email, newWorkspaceName}

		// Check if the new workspace name already exists
		oldWorkspaceName := mysql.FindWorkspaceName(email, newWorkspaceName)

		if oldWorkspaceName == newWorkspaceName {
			// If the new workspace name already exists, print an error message
			fmt.Println(errors.New("The workspace name is already present. Please try another one"))
		} else {
			// Insert the new workspace into the database
			mysql.InsertWorkspaceData(values)
			fmt.Println("New workspace is successfully added")
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(newspaceCmd)
	// Add a command-line flag to specify the name of the new workspace
	newspaceCmd.Flags().StringP("name", "n", "", "Specify the name to add a new workspace")
}
