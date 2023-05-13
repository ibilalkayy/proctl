package user

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/mysql"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

func ComparePasswords(hashPass string, plainPass []byte) bool {
	hashByte := []byte(hashPass)

	if err := bcrypt.CompareHashAndPassword(hashByte, plainPass); err != nil {
		return false
	}
	return true
}

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Provide an email address and password in order to login",
	Run: func(cmd *cobra.Command, args []string) {
		loginEmail, _ := cmd.Flags().GetString("email")
		loginPassword, _ := cmd.Flags().GetString("password")

		loginToken := redis.GetAccountInfo("LoginToken")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		if len(loginToken) != 0 || len(memberLoginToken) != 0 {
			fmt.Println(errors.New("You're already logged in."))
		}

		if UserLogin(loginEmail, loginPassword) || MemberLogin(loginEmail, loginPassword) {
			fmt.Println("You're successfully logged in.")
		} else {
			fmt.Println(errors.New("Invalid email or password"))
		}
	},
}

func UserLogin(email, password string) bool {
	totalColumns := mysql.CountTableColumns("Signup")
	redisSignupEmail, redisSignupPassword, redisSignupFullName, redisSignupAccountName, redisSignupFound := redis.GetUserCredentials(totalColumns)
	tokenString, jwtTokenGenerated := jwt.GenerateJWT()
	for i := 0; i < totalColumns; i++ {
		mysqlEmail, mysqlPassword, mysqlStatus, mysqlFound := mysql.FindAccount(email, redisSignupPassword[i])
		for ComparePasswords(redisSignupPassword[i], []byte(password)) && ComparePasswords(mysqlPassword, []byte(password)) {
			if redisSignupFound && jwtTokenGenerated && mysqlFound && email == mysqlEmail && email == redisSignupEmail[i] {
				redis.SetAccountInfo("LoginToken", tokenString)
				redis.SetAccountInfo("AccountFullName", redisSignupFullName[i])
				redis.SetAccountInfo("AccountName", redisSignupAccountName[i])
				redis.SetAccountInfo("AccountEmail", redisSignupEmail[i])
				redis.SetAccountInfo("AccountPassword", redisSignupPassword[i])
				accountEmail := redis.GetAccountInfo("AccountEmail")
				accountPassword := redis.GetAccountInfo("AccountPassword")

				accountCode := GetRandomCode(accountEmail, accountPassword)
				if mysqlStatus == "0" {
					redis.SetAccountInfo("VerificationCode", accountCode)
				}
				return true
			}
		}
	}
	return false
}

func MemberLogin(email, password string) bool {
	totalColumns := mysql.CountTableColumns("Members")
	redisMemberEmail, redisMemberPassword, redisMemberAccountName, redisMemberFound := redis.GetMemberCredentials(totalColumns)
	tokenString, jwtTokenGenerated := jwt.GenerateJWT()
	for i := 0; i < totalColumns; i++ {
		mysqlCred, mysqlFound := mysql.FindMember(email, redisMemberPassword[i])
		for ComparePasswords(redisMemberPassword[i], []byte(password)) && ComparePasswords(mysqlCred[1], []byte(password)) {
			if redisMemberFound && jwtTokenGenerated && mysqlFound && email == mysqlCred[0] && email == redisMemberEmail[i] {
				redis.SetAccountInfo("MemberLoginToken", tokenString)
				redis.SetAccountInfo("MemberAccountEmail", redisMemberEmail[i])
				redis.SetAccountInfo("MemberAccountPassword", redisMemberPassword[i])
				redis.SetAccountInfo("MemberAccountName", redisMemberAccountName[i])
				return true
			}
		}
	}
	return false
}

func init() {
	cmd.RootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("email", "e", "", "Specify an email address to login")
	loginCmd.Flags().StringP("password", "p", "", "Specify the password to login")
}
