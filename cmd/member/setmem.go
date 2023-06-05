// Package member contains functionality related to managing members.
package member

// Import necessary packages and libraries.
import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/ibilalkayy/proctl/middleware"
	"github.com/spf13/cobra"
)

// setmemCmd represents the setmem command
// This is the command that will be used to setup member credentials.
var setmemCmd = &cobra.Command{
	Use:   "setmem",                       // This is the command name.
	Short: "Setup the member credentials", // A short description of the command.
	// This function is run when the command is called.
	Run: func(cmd *cobra.Command, args []string) {
		// Get the values of the flags (email, password, full name, account name).
		memberEmail, _ := cmd.Flags().GetString("email")
		memberPassword, _ := cmd.Flags().GetString("password")
		memberFullName, _ := cmd.Flags().GetString("full name")
		memberAccountName, _ := cmd.Flags().GetString("account name")

		// Get login token from Redis.
		loginToken := redis.GetAccountInfo("LoginToken")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")
		memberCred, memberFound := mysql.FindMember(memberEmail, "")
		hashPass := middleware.HashPassword([]byte(memberPassword))

		// Create member credentials array.
		memberCredentials := [3]string{memberEmail, hashPass, memberAccountName}

		// If the user is not logged in, generate a JWT, store credentials in Redis, and update or insert the member in MySQL.
		if len(loginToken) == 0 && len(memberLoginToken) == 0 {
			tokenString, jwtTokenGenerated := jwt.GenerateJWT()

			// Check if all required fields are present and the member is not yet set.
			if len(memberEmail) != 0 && len(memberPassword) != 0 && len(memberFullName) != 0 && len(memberAccountName) != 0 {
				if len(memberCred[1]) == 0 && len(memberCred[2]) == 0 && len(memberCred[3]) == 0 {
					if jwtTokenGenerated {
						redis.SetMemberCredentials(memberCredentials)
						totalColumns := mysql.CountTableColumns("Members")
						redisMemberEmail, redisMemberPassword, redisMemberAccountName, _ := redis.GetMemberCredentials(totalColumns)
						redis.SetAccountInfo("MemberLoginToken", tokenString)
						redis.SetAccountInfo("MemberAccountEmail", redisMemberEmail[0])
						redis.SetAccountInfo("MemberAccountPassword", redisMemberPassword[0])
						redis.SetAccountInfo("MemberAccountName", redisMemberAccountName[0])
						if memberFound {
							values := [8]string{hashPass, memberFullName, memberAccountName, "", "", "", "", ""}
							mysql.UpdateMember(values, memberEmail, "", true)
							fmt.Println("You have successfully setup the member credentials")
						} else {
							fmt.Println(errors.New("Please enter the email address or type 'proctl setmem --help'"))
						}
					} else {
						fmt.Println(errors.New("Failure in setting up a member"))
					}
				} else {
					fmt.Println(errors.New("The member is already set. Please update it"))
				}
			} else {
				fmt.Println(errors.New("The required credentials are not entered"))
			}
		} else {
			fmt.Println(errors.New("First logout to setup the member credentials"))
		}
	},
}

// init is a special Go function that is automatically executed at the start of the program.
func init() {
	// This function adds setmemCmd to the list of commands in the root command.
	cmd.RootCmd.AddCommand(setmemCmd)

	// This function adds flags to setmemCmd for specifying member credentials.
	setmemCmd.Flags().StringP("email", "e", "", "Specify an email address to setup the credentials")
	setmemCmd.Flags().StringP("password", "p", "", "Specify a password to setup the credentials")
	setmemCmd.Flags().StringP("full name", "f", "", "Specify a full name to setup the credentials")
	setmemCmd.Flags().StringP("account name", "a", "", "Specify an account name to setup the credentials")
}
