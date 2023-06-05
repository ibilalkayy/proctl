// Package member contains functionality related to managing members.
package member

// Import necessary packages and libraries.
import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/cmd/user"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
// This is the command that will be used to check a member's invitation status.
var checkCmd = &cobra.Command{
	Use:   "check",                              // This is the command name.
	Short: "Check the invitation of the member", // A short description of the command.
	// This function is run when the command is called.
	Run: func(cmd *cobra.Command, args []string) {
		// Get email flag.
		checkEmail, _ := cmd.Flags().GetString("email")

		// Get login token from Redis.
		loginToken := redis.GetAccountInfo("LoginToken")
		var verificationCode string

		// Check if member is present in MySQL database.
		memberStatus, _ := mysql.FindMember(checkEmail, "")

		// If not logged in, check if the email is not empty and if the member is not yet verified.
		// If true, verify the member. Otherwise, output an error.
		if len(loginToken) == 0 {
			if len(checkEmail) != 0 {
				if len(memberStatus[2]) == 0 {
					fmt.Printf("Enter the verification code: ")
					fmt.Scanln(&verificationCode)

					getVerificationCode := user.GetRandomCode(checkEmail, checkEmail)
					if len(verificationCode) != 0 && getVerificationCode == verificationCode {
						mysql.InsertMemberData(checkEmail)
						fmt.Println("Your account is successfully verified")
					} else {
						fmt.Println(errors.New("Please enter the correct verification code"))
					}
				} else {
					fmt.Println(errors.New("Your account is already verified"))
				}
			} else {
				fmt.Println(errors.New("Please put the email address or type 'proctl check --help' for help"))
			}
		} else {
			fmt.Println(errors.New("First logout to check the invitation"))
		}
	},
}

// init is a special Go function that is automatically executed at the start of the program.
func init() {
	// This function adds checkCmd to the list of commands in the root command.
	cmd.RootCmd.AddCommand(checkCmd)

	// This function adds a flag to checkCmd for specifying the email.
	checkCmd.Flags().StringP("email", "e", "", "Specify an email address to check that you're invited or not")
}
