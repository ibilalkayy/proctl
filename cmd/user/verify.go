package user

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/email"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Check verification of a user",
	Run: func(cmd *cobra.Command, args []string) {
		loginToken := redis.GetAccountInfo("LoginToken")
		if len(loginToken) != 0 && jwt.RefreshToken() {
			AccountName := redis.GetAccountInfo("AccountName")
			AccountFullName := redis.GetAccountInfo("AccountFullName")
			AccountEmail := redis.GetAccountInfo("AccountEmail")
			AccountPassword := redis.GetAccountInfo("AccountPassword")
			getVerificationCode := redis.GetAccountInfo("VerificationCode")

			_, _, mysqlStatus, _ := mysql.FindAccount(AccountEmail, AccountPassword)
			if mysqlStatus == "0" && len(getVerificationCode) != 0 {

				accountDetails := GetDetails()
				values := [5]string{"account-template", AccountName, accountDetails[2], AccountEmail, "[proctl] Confirm your email address"}
				email.Verify(values)
				var verificationCode string
				fmt.Printf("Enter the verification code: ")
				fmt.Scanln(&verificationCode)

				if getVerificationCode == verificationCode {
					redis.DelToken("VerificationCode")
					userData := [4]string{AccountFullName, AccountName, AccountPassword, "1"}
					mysql.UpdateUser(userData, AccountEmail, AccountPassword)
					tokenString, _ := jwt.GenerateJWT()
					redis.SetAccountInfo("LoginToken", tokenString)
					fmt.Println("You're successfully verified")
				} else {
					fmt.Println(errors.New("Error in verification. Please try again!!"))
				}
			} else {
				fmt.Println(errors.New("Your account is already verified"))
			}
		} else {
			fmt.Println(errors.New("First login to verify an account"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(verifyCmd)
}
