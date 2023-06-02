package board

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
)

// setboardCmd represents the setboard command
var setboardCmd = &cobra.Command{
	Use:   "setboard",
	Short: "Setup the board",
	Run: func(cmd *cobra.Command, args []string) {
		boardName, _ := cmd.Flags().GetString("board")

		accountEmail := redis.GetAccountInfo("AccountEmail")
		loginToken := redis.GetAccountInfo("LoginToken")

		memberAccountEmail := redis.GetAccountInfo("MemberAccountEmail")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		userBoard := mysql.FindBoard(accountEmail, boardName)
		memberBoard := mysql.FindBoard(memberAccountEmail, boardName)

		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			if len(userBoard) == 0 && userBoard != boardName {
				mysql.InsertBoard(accountEmail, boardName)
				fmt.Println("The board is successfully inserted")
			} else {
				fmt.Println(errors.New("The board by this name is already present"))
			}
		} else if len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			if len(memberBoard) == 0 && memberBoard != boardName {
				mysql.InsertBoard(memberAccountEmail, boardName)
				fmt.Println("The board is successfully inserted")
			} else {
				fmt.Println(errors.New("The board by this name is already present"))
			}
		} else {
			fmt.Println(errors.New("First login to setup the board"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(setboardCmd)
	setboardCmd.Flags().StringP("board", "b", "", "Specify the baord name to setup")
}
