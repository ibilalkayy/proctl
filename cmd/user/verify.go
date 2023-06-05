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
		// Check if a login token exists for the user
		loginToken := redis.GetAccountInfo("LoginToken")
		if len(loginToken) != 0 && jwt.RefreshToken("user") {
			AccountName := redis.GetAccountInfo("AccountName")
			AccountFullName := redis.GetAccountInfo("AccountFullName")
			AccountEmail := redis.GetAccountInfo("AccountEmail")
			AccountPassword := redis.GetAccountInfo("AccountPassword")
			getVerificationCode := redis.GetAccountInfo("VerificationCode")

			// Check if the account is not already verified and has a verification code
			_, _, mysqlStatus, _ := mysql.FindAccount(AccountEmail, AccountPassword)
			if mysqlStatus == "0" && len(getVerificationCode) != 0 {

				// Send verification email
				accountEmail := redis.GetAccountInfo("AccountEmail")
				accountPassword := redis.GetAccountInfo("AccountPassword")
				accountCode := GetRandomCode(accountEmail, accountPassword)
				values := [5]string{"account-template", AccountName, accountCode, AccountEmail, "[proctl] Confirm your email address"}
				email.Verify(values)

				var verificationCode string
				fmt.Printf("Enter the verification code: ")
				fmt.Scanln(&verificationCode)

				// Delete the verification code and update the user status in the database
				if getVerificationCode == verificationCode {
					redis.DelToken("VerificationCode")
					userData := [4]string{AccountFullName, AccountName, AccountPassword, "1"}
					mysql.UpdateUser(userData, AccountEmail, AccountPassword)

					// Generate a new login token
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
