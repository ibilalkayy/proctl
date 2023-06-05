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

// setboardCmd represents the setboard command.
// This is the command that will be used to set up a board.
var setboardCmd = &cobra.Command{
	Use:   "setboard",        // This is the command name.
	Short: "Setup the board", // A short description of the command.

	// This function is run when the command is called.
	Run: func(cmd *cobra.Command, args []string) {
		// Get board name flag.
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
		// If user board with the same name doesn't exist, create it.
		// Otherwise, output an error.
		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			if len(userBoard) == 0 && userBoard != boardName {
				mysql.InsertBoard(accountEmail, boardName)
				fmt.Println("The board is successfully inserted")
			} else {
				fmt.Println(errors.New("The board by this name is already present"))
			}
			// If user is not logged in, check for member login and refresh the token.
			// If member board with the same name doesn't exist, create it.
			// Otherwise, output an error.
		} else if len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			if len(memberBoard) == 0 && memberBoard != boardName {
				mysql.InsertBoard(memberAccountEmail, boardName)
				fmt.Println("The board is successfully inserted")
			} else {
				fmt.Println(errors.New("The board by this name is already present"))
			}
			// If neither user nor member are logged in, output an error.
		} else {
			fmt.Println(errors.New("First login to setup the board"))
		}
	},
}

// init is a special Go function that is automatically executed at the start of the program.
func init() {
	// This function adds setboardCmd to the list of commands in the root command.
	cmd.RootCmd.AddCommand(setboardCmd)

	// This function adds a flag to setboardCmd for specifying the board name.
	setboardCmd.Flags().StringP("board", "b", "", "Specify the board name to setup")
}
