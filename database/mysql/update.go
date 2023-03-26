package mysql

import (
	"errors"
	"fmt"

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
	db := Connect()
	q := "UPDATE Profiles SET titles=?, phones=?, locations=?, working_statuses=? WHERE emails=?"
	update, err := db.Prepare(q)
	middleware.HandleError(err)

	defer update.Close()

	if len(value[0]) != 0 {
		redis.SetAccountInfo("ProfileTitle", value[0])
		getPhone := redis.GetAccountInfo("ProfilePhone")
		getLocation := redis.GetAccountInfo("ProfileLocation")
		getWorkingStatus := redis.GetAccountInfo("ProfileWorkingStatus")
		_, err = update.Exec(value[0], getPhone, getLocation, getWorkingStatus, email)
		middleware.HandleError(err)
	}

	if len(value[1]) != 0 {
		redis.SetAccountInfo("ProfilePhone", value[1])
		getTitle := redis.GetAccountInfo("ProfileTitle")
		getLocation := redis.GetAccountInfo("ProfileLocation")
		getWorkingStatus := redis.GetAccountInfo("ProfileWorkingStatus")
		_, err = update.Exec(getTitle, value[1], getLocation, getWorkingStatus, email)
		middleware.HandleError(err)
	}

	if len(value[2]) != 0 {
		redis.SetAccountInfo("ProfileLocation", value[2])
		getTitle := redis.GetAccountInfo("ProfileTitle")
		getPhone := redis.GetAccountInfo("ProfilePhone")
		getWorkingStatus := redis.GetAccountInfo("ProfileWorkingStatus")
		_, err = update.Exec(getTitle, getPhone, value[2], getWorkingStatus, email)
		middleware.HandleError(err)
	}

	if len(value[3]) != 0 {
		redis.SetAccountInfo("ProfileWorkingStatus", value[3])
		getTitle := redis.GetAccountInfo("ProfileTitle")
		getPhone := redis.GetAccountInfo("ProfilePhone")
		getLocation := redis.GetAccountInfo("ProfileLocation")
		_, err = update.Exec(getTitle, getPhone, getLocation, value[3], email)
		middleware.HandleError(err)
	}
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

func SetMember() (string, string, string) {
	password := redis.GetAccountInfo("MemberPassword")
	fullName := redis.GetAccountInfo("MemberFullName")
	accountName := redis.GetAccountInfo("MemberAccountName")

	return password, fullName, accountName
}

func UpdateMember(value [3]string, email string) {
	db := Connect()
	q := "UPDATE Members SET passwords=?, fullnames=?, accountnames=?, is_active=? WHERE emails=?"
	update, err := db.Prepare(q)
	middleware.HandleError(err)

	defer update.Close()

	if len(value[0]) != 0 && len(email) != 0 {
		redis.SetAccountInfo("MemberPassword", value[0])
		_, fullName, accountName := SetMember()
		_, err = update.Exec(value[0], fullName, accountName, "1", email)
		middleware.HandleError(err)
	}

	if len(value[1]) != 0 && len(email) != 0 {
		redis.SetAccountInfo("MemberFullName", value[1])
		password, _, accountName := SetMember()
		_, err = update.Exec(password, value[1], accountName, "1", email)
		middleware.HandleError(err)
	}

	if len(value[2]) != 0 && len(email) != 0 {
		redis.SetAccountInfo("MemberAccountName", value[2])
		password, fullName, _ := SetMember()
		_, err = update.Exec(password, fullName, value[2], "1", email)
		middleware.HandleError(err)
	}
}
