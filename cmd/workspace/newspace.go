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
		newWorkspaceName, _ := cmd.Flags().GetString("name")
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
			fmt.Println(errors.New("First login to add a new workspace"))
		}

		values := [2]string{email, newWorkspaceName}
		oldWorkspaceName := mysql.FindWorkspaceName(email, newWorkspaceName)

		if oldWorkspaceName == newWorkspaceName {
			fmt.Println(errors.New("The workspace name is already present. Please try another one"))
		} else {
			mysql.InsertWorkspaceData(values)
			fmt.Println("New workspace is successfully added")
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(newspaceCmd)
	newspaceCmd.Flags().StringP("name", "n", "", "Specify the name to add a new workspace")
}
