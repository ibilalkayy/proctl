package member

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
var setmemCmd = &cobra.Command{
	Use:   "setmem",
	Short: "Setup the member credentials",
	Run: func(cmd *cobra.Command, args []string) {
		memberEmail, _ := cmd.Flags().GetString("email")
		memberPassword, _ := cmd.Flags().GetString("password")
		memberFullName, _ := cmd.Flags().GetString("full name")
		memberAccountName, _ := cmd.Flags().GetString("account name")

		loginToken := redis.GetAccountInfo("LoginToken")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")
		memberCred, memberFound := mysql.FindMember(memberEmail, "")
		hashPass := middleware.HashPassword([]byte(memberPassword))
		memberCredentials := [3]string{memberEmail, hashPass, memberAccountName}

		if len(loginToken) == 0 && len(memberLoginToken) == 0 {
			tokenString, jwtTokenGenerated := jwt.GenerateJWT()
			if len(memberEmail) != 0 && len(memberPassword) != 0 && len(memberFullName) != 0 && len(memberAccountName) != 0 {
				if len(memberCred[1]) == 0 && len(memberCred[2]) == 0 && len(memberCred[3]) == 0 {
					if jwtTokenGenerated {
						redis.SetMemberCredentials(memberCredentials)
						totalColumns := mysql.CountTableColumns("Members")
						redisMemberEmail, redisMemberPassword, _, redisMemberAccountName, _ := redis.GetMemberCredentials(totalColumns)
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

func init() {
	cmd.RootCmd.AddCommand(setmemCmd)
	setmemCmd.Flags().StringP("email", "e", "", "Specify an email address to setup the credentials")
	setmemCmd.Flags().StringP("password", "p", "", "Specify a password to setup the credentials")
	setmemCmd.Flags().StringP("full name", "f", "", "Specify a full name to setup the credentials")
	setmemCmd.Flags().StringP("account name", "a", "", "Specify an account name to setup the credentials")
}
