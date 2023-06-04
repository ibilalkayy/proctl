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

// dropboardCmd represents the dropboard command
var dropboardCmd = &cobra.Command{
	Use:   "dropboard",
	Short: "Delete the board",
	Run: func(cmd *cobra.Command, args []string) {
		boardName, _ := cmd.Flags().GetString("board")

		accountEmail := redis.GetAccountInfo("AccountEmail")
		loginToken := redis.GetAccountInfo("LoginToken")

		memberAccountEmail := redis.GetAccountInfo("MemberAccountEmail")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		userBoard := mysql.FindBoard(accountEmail, boardName)
		memberBoard := mysql.FindBoard(memberAccountEmail, boardName)

		if len(loginToken) != 0 && jwt.RefreshToken("user") || len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			if len(userBoard) != 0 {
				mysql.DeleteBoard(accountEmail, boardName)
				fmt.Println("The board is successfully deleted")
			} else if len(memberBoard) != 0 {
				mysql.DeleteBoard(memberAccountEmail, boardName)
				fmt.Println("The board is successfully deleted")
			} else {
				fmt.Println(errors.New("The board is not present by this name"))
			}
		} else {
			fmt.Println(errors.New("First login to delete the board"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(dropboardCmd)
	dropboardCmd.Flags().StringP("board", "b", "", "Specify the baord name to delete")
}
