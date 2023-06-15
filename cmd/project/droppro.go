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

// dropproCmd represents the droppro command
var dropproCmd = &cobra.Command{
	Use:   "droppro",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		boardName, _ := cmd.Flags().GetString("board")
		projectName, _ := cmd.Flags().GetString("project")

		accountEmail := redis.GetAccountInfo("AccountEmail")
		loginToken := redis.GetAccountInfo("LoginToken")

		memberAccountEmail := redis.GetAccountInfo("MemberAccountEmail")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		userProject := mysql.FindProject(accountEmail, boardName, projectName)
		memberProject := mysql.FindProject(memberAccountEmail, boardName, projectName)

		userBoard := mysql.FindBoard(accountEmail, boardName)
		memberBoard := mysql.FindBoard(memberAccountEmail, boardName)

		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			if projectName == userProject {
				if userBoard == boardName {
					mysql.DeleteProject(accountEmail, boardName, projectName)
					fmt.Println("The project is successfully deleted")
				} else {
					fmt.Println(errors.New("The board name is wrong"))
				}
			} else {
				fmt.Println(errors.New("The project by this name is not present"))
			}
		} else if len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			if projectName == memberProject {
				if memberBoard == boardName {
					mysql.DeleteProject(accountEmail, boardName, projectName)
					fmt.Println("The project is successfully deleted")
				} else {
					fmt.Println(errors.New("The board name is wrong"))
				}
			} else {
				fmt.Println(errors.New("TThe project by this name is not present"))
			}
		} else {
			fmt.Println(errors.New("First login to add the project"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(dropproCmd)
	dropproCmd.Flags().StringP("board", "b", "", "Specify the board name to delete")
	dropproCmd.Flags().StringP("project", "p", "", "Specify the project name to delete")
}
