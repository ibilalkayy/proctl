package project

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
)

// updateproCmd represents the updatepro command
var updateproCmd = &cobra.Command{
	Use:   "updatepro",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		boardName, _ := cmd.Flags().GetString("board")
		oldProjectName, _ := cmd.Flags().GetString("oldProject")
		newProjectName, _ := cmd.Flags().GetString("newProject")
		status, _ := cmd.Flags().GetString("status")
		date, _ := cmd.Flags().GetString("date")

		accountEmail := redis.GetAccountInfo("AccountEmail")
		loginToken := redis.GetAccountInfo("LoginToken")

		memberAccountEmail := redis.GetAccountInfo("MemberAccountEmail")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		userBoard := mysql.FindBoard(accountEmail, boardName)
		memberBoard := mysql.FindBoard(memberAccountEmail, boardName)

		values := [3]string{newProjectName, status, date}

		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			if userBoard == boardName {
				mysql.UpdateProject(values, boardName, oldProjectName)
				fmt.Println("The project data is successfully updated")
			} else {
				fmt.Println(errors.New("The board name is wrong"))
			}
		} else if len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			if memberBoard == boardName {
				mysql.UpdateProject(values, boardName, oldProjectName)
				fmt.Println("The project data is successfully updated")
			} else {
				fmt.Println(errors.New("The board name is wrong"))
			}
		} else {
			fmt.Println(errors.New("First login to add the project"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(updateproCmd)
	updateproCmd.Flags().StringP("board", "b", "", "Specify the board name to update")
	updateproCmd.Flags().StringP("oldProject", "o", "", "Specify the old project name to update")
	updateproCmd.Flags().StringP("newProject", "n", "", "Specify the new project name to update")
	updateproCmd.Flags().StringP("status", "s", "", "Specify the status to update")
	updateproCmd.Flags().StringP("date", "d", "", "Specify the date to update")
}
