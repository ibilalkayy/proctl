package mysql

import (
	"log"

	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/middleware"
)

func UpdateUser(value [3]string, email, password string) {
	db := Connect()
	q := "UPDATE Signup SET fullnames=?, accountnames=?, is_active=? WHERE emails=? AND passwords=?"
	update, err := db.Prepare(q)
	middleware.HandleError(err)

	if len(value[0]) != 0 {
		redis.SetAccountInfo("AccountFullName", value[0])
		getAccountName := redis.GetAccountInfo("AccountName")
		_, err = update.Exec(value[0], getAccountName, value[2], email, password)
		middleware.HandleError(err)
	}

	if len(value[1]) != 0 {
		redis.SetAccountInfo("AccountName", value[1])
		getFullName := redis.GetAccountInfo("AccountFullName")
		_, err = update.Exec(getFullName, value[1], value[2], email, password)
		middleware.HandleError(err)
	}
}

func UpdateProfile(value [4]string, email string) {
	db := Connect()
	q := "UPDATE Profiles SET titles=?, phones=?, locations=?, working_statuses=? WHERE emails=?"
	update, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}

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
