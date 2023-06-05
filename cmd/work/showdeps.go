package work

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
)

// showdepsCmd represents the showdeps command
var showdepsCmd = &cobra.Command{
	Use:   "showdeps",
	Short: "Show the deparments of a member and the admin",
	Run: func(cmd *cobra.Command, args []string) {
		loginToken := redis.GetAccountInfo("LoginToken")
		memberLoginToken := redis.GetAccountInfo("MemberLoginToken")
		if (len(loginToken) != 0 && jwt.RefreshToken("user")) || (len(memberLoginToken) != 0 && jwt.RefreshToken("member")) {
			roles := [12]string{" Sales & CRM", " Legal", " HR & Recruiting", " Marketing", " Software Development", " Finance", " Education", " Operations", " Product Management", "Construction", "Nonprofits", "IT"}
			fmt.Println("Following are the departments")

			index := 1
			for _, value := range roles {
				fmt.Printf("%d. %s\n", index, value)
				index++
			}
		} else {
			fmt.Println(errors.New("First login to show the departments of a member"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(showdepsCmd)
}