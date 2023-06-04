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

// updateboardCmd represents the updateboard command
var updateboardCmd = &cobra.Command{
	Use:   "updateboard",
	Short: "Update the board",
	Run: func(cmd *cobra.Command, args []string) {
		oldBoardName, _ := cmd.Flags().GetString("oldBoard")
		newBoardName, _ := cmd.Flags().GetString("newBoard")

		accountEmail := redis.GetAccountInfo("AccountEmail")
		loginToken := redis.GetAccountInfo("LoginToken")

		memberAccountEmail := redis.GetAccountInfo("MemberAccountEmail")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		userBoard := mysql.FindBoard(accountEmail, oldBoardName)
		memberBoard := mysql.FindBoard(memberAccountEmail, oldBoardName)

		if len(loginToken) != 0 && jwt.RefreshToken("user") || len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			if len(userBoard) != 0 {
				mysql.UpdateBoard(accountEmail, oldBoardName, newBoardName)
				fmt.Println("The board is successfully updated")
			} else if len(memberBoard) != 0 {
				mysql.UpdateBoard(memberAccountEmail, oldBoardName, newBoardName)
				fmt.Println("The board is successfully updated")
			} else {
				fmt.Println(errors.New("The board is not present by this name. Please setup the board, first"))
			}
		} else {
			fmt.Println(errors.New("First login to update the board"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(updateboardCmd)
	updateboardCmd.Flags().StringP("oldBoard", "o", "", "Specify the old baord name to update")
	updateboardCmd.Flags().StringP("newBoard", "n", "", "Specify the new baord name to update")
}
