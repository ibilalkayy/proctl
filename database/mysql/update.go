package mysql

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/middleware"
)

func UpdateUser(value [4]string, email, password string) {
	db := Connect()
	q := "UPDATE Signup SET fullnames=?, accountnames=?, passwords=?, is_active=? WHERE emails=? AND passwords=?"
	update, err := db.Prepare(q)
	middleware.HandleError(err)

	defer update.Close()

	if len(value[0]) != 0 {
		redis.SetAccountInfo("AccountFullName", value[0])
		getAccountName := redis.GetAccountInfo("AccountName")
		getPassword := redis.GetAccountInfo("AccountPassword")
		_, err = update.Exec(value[0], getAccountName, getPassword, value[3], email, password)
		middleware.HandleError(err)
	}

	if len(value[1]) != 0 {
		redis.SetAccountInfo("AccountName", value[1])
		getFullName := redis.GetAccountInfo("AccountFullName")
		getPassword := redis.GetAccountInfo("AccountPassword")
		_, err = update.Exec(getFullName, value[1], getPassword, value[3], email, password)
		middleware.HandleError(err)
	}

	if len(value[2]) != 0 {
		redis.SetAccountInfo("AccountPassword", value[2])
		getFullName := redis.GetAccountInfo("AccountFullName")
		getAccountName := redis.GetAccountInfo("AccountName")
		getAccountEmail := redis.GetAccountInfo("AccountEmail")
		hashPass := middleware.HashPassword([]byte(value[2]))
		redis.DelToken("LoginToken")

		values := [4]string{getAccountEmail, hashPass, getFullName, getAccountName}
		redis.SetCredentials(values)
		_, err = update.Exec(getFullName, getAccountName, hashPass, value[3], email, password)
		middleware.HandleError(err)
	}
}

func UpdateProfile(value [4]string, email string) {
	if len(value) == 0 && len(email) == 0 {
		return
	}

	db := Connect()
	q := "UPDATE Profiles SET "
	var updateValues []interface{}

	for i := 0; i < len(value); i++ {
		if len(value[i]) != 0 {
			switch i {
			case 0:
				q += "titles=?, "
			case 1:
				q += "phones=?, "
			case 2:
				q += "locations=?, "
			case 3:
				q += "working_statuses=?, "
			}
			updateValues = append(updateValues, value[i])
		}
	}

	q = strings.TrimSuffix(q, ", ")
	q += " WHERE emails=?"
	updateValues = append(updateValues, email)

	update, err := db.Prepare(q)
	middleware.HandleError(err)

	defer update.Close()

	_, err = update.Exec(updateValues...)
	middleware.HandleError(err)
}

func UpdateWorkspace(value [3]string) {
	db := Connect()
	q := "UPDATE Workspaces SET names=? WHERE emails=? AND names=?"
	update, err := db.Prepare(q)
	middleware.HandleError(err)

	defer update.Close()

	if len(value[0]) != 0 && len(value[1]) != 0 && len(value[2]) != 0 {
		_, err = update.Exec(value[0], value[1], value[2])
		middleware.HandleError(err)
	} else {
		fmt.Println(errors.New("More flags are required to update the workspace"))
	}
}

func UpdateMember(value [7]string, email string) {
	if len(value) == 0 && len(email) == 0 {
		return
	}

	db := Connect()
	q := "UPDATE Members SET "
	var updateValues []interface{}

	for i := 0; i < len(value); i++ {
		if len(value[i]) != 0 {
			switch i {
			case 0:
				q += "passwords=?, "
			case 1:
				q += "fullnames=?, "
			case 2:
				q += "accountnames=?, "
			case 3:
				q += "titles=?, "
			case 4:
				q += "phones=?, "
			case 5:
				q += "locations=?, "
			case 6:
				q += "working_statuses=?, "
			}
			updateValues = append(updateValues, value[i])
		}
	}

	q += "is_active=? WHERE emails=?"
	updateValues = append(updateValues, "1", email)

	update, err := db.Prepare(q)
	middleware.HandleError(err)

	defer update.Close()

	_, err = update.Exec(updateValues...)
	middleware.HandleError(err)
}
