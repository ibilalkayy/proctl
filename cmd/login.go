package cmd

import (
	"errors"
	"fmt"

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

		redisSignupEmail, redisSignupPassword, redisSignupFound := redis.GetCredentials()
		tokenString, jwtTokenGenerated := jwt.GenerateJWT()
		for i := 0; i < 2; i++ {
			mysqlEmail, mysqlPassword, mysqlFound := mysql.FindAccount(loginEmail, redisSignupPassword[i])
			for ComparePasswords(redisSignupPassword[i], []byte(loginPassword)) && ComparePasswords(mysqlPassword, []byte(loginPassword)) {
				if redisSignupFound && jwtTokenGenerated && mysqlFound && loginEmail == mysqlEmail && loginEmail == redisSignupEmail[i] {
					redis.SetToken("LoginToken", tokenString)
					fmt.Println("You're successfully logged in")
					break
				} else {
					fmt.Println(errors.New("invalid or no credentials: try out again"))
					break
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("email", "e", "", "Specify an email address to login")
	loginCmd.Flags().StringP("password", "p", "", "Specify the password to login")
}
