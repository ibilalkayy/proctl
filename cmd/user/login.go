package user

// Import necessary packages
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

// ComparePasswords compares a hashed password with a plaintext one
func ComparePasswords(hashPass string, plainPass []byte) bool {
	hashByte := []byte(hashPass)

	// Use bcrypt to compare the hashed password with the plaintext password
	if err := bcrypt.CompareHashAndPassword(hashByte, plainPass); err != nil {
		return false
	}
	return true
}

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Provide an email address and password in order to login",
	// The Run function is the action performed when the login command is called
	Run: func(cmd *cobra.Command, args []string) {
		// Get email and password from command line flags
		loginEmail, _ := cmd.Flags().GetString("email")
		loginPassword, _ := cmd.Flags().GetString("password")

		// Check if there are existing tokens in Redis cache
		loginToken := redis.GetAccountInfo("LoginToken")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")

		// Handle logic of user login based on existing tokens and validity of user credentials
		if len(loginToken) != 0 || len(memberLoginToken) != 0 {
			fmt.Println(errors.New("You're already logged in."))
		} else if MemberLogin(loginEmail, loginPassword) {
			fmt.Println("You're successfully logged in.")
		} else if UserLogin(loginEmail, loginPassword) {
			fmt.Println("You're successfully logged in.")
		} else {
			fmt.Println(errors.New("Invalid email or password"))
		}
	},
}

// UserLogin handles the user login process
func UserLogin(email, password string) bool {
	// Count the columns in the Signup table in MySQL
	totalColumns := mysql.CountTableColumns("Signup")
	// Get user credentials from Redis cache
	redisSignupEmail, redisSignupPassword, redisSignupFullName, redisSignupAccountName, redisSignupFound := redis.GetUserCredentials(totalColumns)
	// Generate a JWT token
	tokenString, jwtTokenGenerated := jwt.GenerateJWT()

	// Loop through all the entries in the Signup table and attempt to match entered credentials
	for i := 0; i < totalColumns; i++ {
		mysqlEmail, mysqlPassword, mysqlStatus, mysqlFound := mysql.FindAccount(email, redisSignupPassword[i])

		// Continue if the entered password matches both the one in Redis and MySQL
		for ComparePasswords(redisSignupPassword[i], []byte(password)) && ComparePasswords(mysqlPassword, []byte(password)) {
			// If the user is found in both Redis and MySQL, and the JWT token was successfully generated
			if redisSignupFound && jwtTokenGenerated && mysqlFound && email == mysqlEmail && email == redisSignupEmail[i] {
				// Set the account info in Redis cache
				redis.SetAccountInfo("LoginToken", tokenString)
				redis.SetAccountInfo("AccountFullName", redisSignupFullName[i])
				redis.SetAccountInfo("AccountName", redisSignupAccountName[i])
				redis.SetAccountInfo("AccountEmail", redisSignupEmail[i])
				redis.SetAccountInfo("AccountPassword", redisSignupPassword[i])
				accountEmail := redis.GetAccountInfo("AccountEmail")
				accountPassword := redis.GetAccountInfo("AccountPassword")

				// Generate a random code for the account, and set it in Redis if the MySQL status is "0"
				accountCode := GetRandomCode(accountEmail, accountPassword)
				if mysqlStatus == "0" {
					redis.SetAccountInfo("VerificationCode", accountCode)
				}
				return true
			}
		}
	}
	// Return false if no match found
	return false
}

// MemberLogin handles the member login process
func MemberLogin(email, password string) bool {
	// Count the columns in the Members table in MySQL
	totalColumns := mysql.CountTableColumns("Members")
	// Get member credentials from Redis cache
	redisMemberEmail, redisMemberPassword, redisMemberAccountName, redisMemberFound := redis.GetMemberCredentials(totalColumns)
	// Generate a JWT token
	tokenString, jwtTokenGenerated := jwt.GenerateJWT()

	// Loop through all the entries in the Members table and attempt to match entered credentials
	for i := 0; i < totalColumns; i++ {
		mysqlCred, mysqlFound := mysql.FindMember(email, redisMemberPassword[i])

		// Continue if the entered password matches both the one in Redis and MySQL
		for ComparePasswords(redisMemberPassword[i], []byte(password)) && ComparePasswords(mysqlCred[1], []byte(password)) {
			// If the member is found in both Redis and MySQL, and the JWT token was successfully generated
			if redisMemberFound && jwtTokenGenerated && mysqlFound && email == mysqlCred[0] && email == redisMemberEmail[i] {
				// Set the member info in Redis cache
				redis.SetAccountInfo("MemberLoginToken", tokenString)
				redis.SetAccountInfo("MemberAccountEmail", redisMemberEmail[i])
				redis.SetAccountInfo("MemberAccountPassword", redisMemberPassword[i])
				redis.SetAccountInfo("MemberAccountName", redisMemberAccountName[i])
				return true
			}
		}
	}
	// Return false if no match found
	return false
}

// Add the login command to the root command
func init() {
	cmd.RootCmd.AddCommand(loginCmd)
	// Set the flags for the login command
	loginCmd.Flags().StringP("email", "e", "", "Specify an email address to login")
	loginCmd.Flags().StringP("password", "p", "", "Specify the password to login")
}
