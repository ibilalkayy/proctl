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

// addproCmd represents the addpro command
var addproCmd = &cobra.Command{
	Use:   "addpro",
	Short: "Add a project",
	Run: func(cmd *cobra.Command, args []string) {
		boardName, _ := cmd.Flags().GetString("board")
		projectName, _ := cmd.Flags().GetString("project")
		personName, _ := cmd.Flags().GetString("person")
		status, _ := cmd.Flags().GetString("status")
		date, _ := cmd.Flags().GetString("date")

		accountEmail := redis.GetAccountInfo("AccountEmail")
		loginToken := redis.GetAccountInfo("LoginToken")

		memberAccountEmail := redis.GetAccountInfo("MemberAccountEmail")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		userProject := mysql.FindProject(accountEmail, boardName, projectName)
		memberProject := mysql.FindProject(memberAccountEmail, boardName, projectName)

		userBoard := mysql.FindBoard(accountEmail, boardName)
		memberBoard := mysql.FindBoard(memberAccountEmail, boardName)

		accountName := redis.GetAccountInfo("AccountName")
		memberAccountName := redis.GetAccountInfo("MemberAccountName")

		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			if len(userProject) == 0 {
				if userBoard == boardName && accountName == personName {
					mysql.InsertProject(accountEmail, boardName, projectName, personName, status, date)
					fmt.Println("The project is successfully inserted")
				} else {
					fmt.Println(errors.New("The board name or the person name is wrong"))
				}
			} else {
				fmt.Println(errors.New("The project by this name is already present"))
			}
		} else if len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			if len(memberProject) == 0 {
				if memberBoard == boardName && memberAccountName == personName {
					mysql.InsertProject(memberAccountEmail, boardName, projectName, personName, status, date)
					fmt.Println("The project is successfully inserted")
				} else {
					fmt.Println(errors.New("The board name or the person name is wrong"))
				}
			} else {
				fmt.Println(errors.New("The project by this name is already present"))
			}
		} else {
			fmt.Println(errors.New("First login to add the project"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(addproCmd)
	addproCmd.Flags().StringP("board", "b", "", "Specify the board name to add")
	addproCmd.Flags().StringP("project", "p", "", "Specify the project name to add")
	addproCmd.Flags().StringP("person", "n", "", "Specify the person name to add")
	addproCmd.Flags().StringP("status", "s", "", "Specify the status to add")
	addproCmd.Flags().StringP("date", "d", "", "Specify the date to add")
}
