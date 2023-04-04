package member

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/cmd"
	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/jwt"
	"github.com/spf13/cobra"
)

// memrolesCmd represents the memroles command
var memrolesCmd = &cobra.Command{
	Use:   "memroles",
	Short: "Show the roles of a member",
	Run: func(cmd *cobra.Command, args []string) {
		loginToken := redis.GetAccountInfo("LoginToken")
		if len(loginToken) != 0 && jwt.RefreshToken() {
			roles := [4]string{"Business Owner", "Team Leader", "Team Member", "Freelancer"}
			fmt.Println("Following are the roles")

			index := 1
			for _, value := range roles {
				fmt.Printf("%d. %s\n", index, value)
				index++
			}
		} else {
			fmt.Println(errors.New("First login to show the roles of a member"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(memrolesCmd)
}
