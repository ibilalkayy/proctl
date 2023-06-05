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
// It is a Cobra Command object which is responsible for running the command to update a board.
var updateboardCmd = &cobra.Command{
	Use:   "updateboard",      // This is the command name.
	Short: "Update the board", // A short description of the command.

	Run: func(cmd *cobra.Command, args []string) {
		// Getting the old and new board names from the command flags
		oldBoardName, _ := cmd.Flags().GetString("oldBoard")
		newBoardName, _ := cmd.Flags().GetString("newBoard")

		// Retrieving user information from the Redis database
		accountEmail := redis.GetAccountInfo("AccountEmail")
		loginToken := redis.GetAccountInfo("LoginToken")

		// Retrieving member information from the Redis database
		memberAccountEmail := redis.GetAccountInfo("MemberAccountEmail")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		// Looking for a board with the specified old name for both the user and the member
		userBoard := mysql.FindBoard(accountEmail, oldBoardName)
		memberBoard := mysql.FindBoard(memberAccountEmail, oldBoardName)

		// Checking if the user or member is logged in and has a valid token
		if len(loginToken) != 0 && jwt.RefreshToken("user") || len(memberLoginToken) != 0 && jwt.RefreshToken("member") {
			// If the user or member owns a board with the old name, it will be updated
			if len(userBoard) != 0 {
				mysql.UpdateBoard(accountEmail, oldBoardName, newBoardName)
				fmt.Println("The board is successfully updated")
			} else if len(memberBoard) != 0 {
				mysql.UpdateBoard(memberAccountEmail, oldBoardName, newBoardName)
				fmt.Println("The board is successfully updated")
			} else {
				// If no board with the old name is found, an error is returned
				fmt.Println(errors.New("The board is not present by this name. Please setup the board, first"))
			}
		} else {
			// If the user or member is not logged in, an error is returned
			fmt.Println(errors.New("First login to update the board"))
		}
	},
}

func init() {
	// Adding the updateboardCmd to the root command of the Cobra application
	cmd.RootCmd.AddCommand(updateboardCmd)
	// Defining flags for the updateboardCmd command. These allow the user to specify the old and new board names
	updateboardCmd.Flags().StringP("oldBoard", "o", "", "Specify the old baord name to update")
	updateboardCmd.Flags().StringP("newBoard", "n", "", "Specify the new baord name to update")
}
