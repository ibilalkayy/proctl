package member

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
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check the invitation of the member",
	Run: func(cmd *cobra.Command, args []string) {
		checkEmail, _ := cmd.Flags().GetString("email")
		loginToken := redis.GetAccountInfo("LoginToken")
		var verificationCode string
		memberStatus, _ := mysql.FindMember(checkEmail, "")

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
			fmt.Println(errors.New("First logout the check the invitation"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(checkCmd)
	checkCmd.Flags().StringP("email", "e", "", "Specify an email address to check that you're invited or not")
}
