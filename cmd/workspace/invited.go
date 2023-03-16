package workspace

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/cmd/user"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/spf13/cobra"
)

// invitedCmd represents the invited command
var invitedCmd = &cobra.Command{
	Use:   "invited",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		invitedEmail, _ := cmd.Flags().GetString("email")
		loginToken := redis.GetAccountInfo("LoginToken")
		var verificationCode string

		if len(loginToken) == 0 {
			if len(invitedEmail) != 0 {
				fmt.Printf("Enter the verification code: ")
				fmt.Scanln(&verificationCode)

				getVerificationCode := user.GetRandomCode(invitedEmail, invitedEmail)
				if len(verificationCode) != 0 && getVerificationCode == verificationCode {
					// tokenString, _ := jwt.GenerateJWT()
					// redis.SetAccountInfo("LoginToken", tokenString)
					values := [4]string{invitedEmail, "", "", ""}
					mysql.InsertMemberData(values)
					fmt.Println("Your account is successfully verified")
				} else {
					fmt.Println(errors.New("Please enter the correct verification code"))
				}
			} else {
				fmt.Println(errors.New("Please put the email address or type 'proctl invited --help' for help"))
			}
		} else {
			fmt.Println(errors.New("First logout the check the invitation"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(invitedCmd)
	invitedCmd.Flags().StringP("email", "e", "", "Specify an email address to check that you're invited or not")
}
