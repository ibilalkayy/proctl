// Package board contains functionality related to managing boards.
package board

// Import necessary packages and libraries.
import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
)

// dropboardCmd represents the dropboard command.
// This is the command that will be used to delete a board.
var dropboardCmd = &cobra.Command{
	Use:   "dropboard",        // This is the command name.
	Short: "Delete the board", // A short description of the command.

	// This function is run when the command is called.
	Run: func(cmd *cobra.Command, args []string) {
		// Get board name flag
		boardName, _ := cmd.Flags().GetString("board")

		// Get account information from Redis.
		accountEmail := redis.GetAccountInfo("AccountEmail")
		loginToken := redis.GetAccountInfo("LoginToken")

		// Get member account information from Redis.
		memberAccountEmail := redis.GetAccountInfo("MemberAccountEmail")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		// Get board details from MySQL.
		userBoard := mysql.FindBoard(accountEmail, boardName)
		memberBoard := mysql.FindBoard(memberAccountEmail, boardName)

		// Check if user is logged in and refresh the token.
		// If user or member is logged in and their board exists, delete it. Otherwise, output an error.
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

// init is a special Go function that is automatically executed at the start of the program.
func init() {
	// This function adds dropboardCmd to the list of commands in the root command.
	cmd.RootCmd.AddCommand(dropboardCmd)

	// This function adds a flag to dropboardCmd for specifying the board name.
	dropboardCmd.Flags().StringP("board", "b", "", "Specify the baord name to delete")
}
