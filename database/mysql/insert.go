package mysql

import (
	"errors"
	"fmt"
	"time"

	"github.com/ibilalkayy/proctl/middleware"
)

func InsertSignupData(value [4]string) {
	db := CreateTable(0)
	q := "INSERT INTO Signup(emails, passwords, fullnames, accountnames, is_active, created_at) VALUES(?, ?, ?, ?, ?, ?)"
	insert, err := db.Prepare(q)
	middleware.HandleError(err)

	defer insert.Close()

	if len(value[0]) != 0 || len(value[1]) != 0 {
		currentTime := time.Now()
		_, err := insert.Exec(value[0], value[1], value[2], value[3], 0, currentTime)
		middleware.HandleError(err)
	}
}

func InsertProfileData(value [5]string) {
	db := CreateTable(1)
	q := "INSERT INTO Profiles(emails, titles, phones, locations, working_statuses) VALUES(?, ?, ?, ?, ?)"
	insert, err := db.Prepare(q)
	middleware.HandleError(err)

	defer insert.Close()

	profileFound := FindProfile(value[0])
	if profileFound {
		fmt.Println("Your profile data is already inserted. Type 'proctl update [flags]'")
	} else {
		if len(value[0]) == 0 && len(value[1]) == 0 && len(value[2]) == 0 && len(value[3]) == 0 && len(value[4]) == 0 {
			fmt.Println(errors.New("Enter the profile information first."))
		} else if len(value[1]) == 0 {
			_, err := insert.Exec(value[0], "", value[2], value[3], value[4])
			middleware.HandleError(err)
		} else if len(value[2]) == 0 {
			_, err := insert.Exec(value[0], value[1], "", value[3], value[4])
			middleware.HandleError(err)
		} else if len(value[3]) == 0 {
			_, err := insert.Exec(value[0], value[1], value[2], "", value[4])
			middleware.HandleError(err)
		} else if len(value[4]) == 0 {
			_, err := insert.Exec(value[0], value[1], value[2], value[3], "")
			middleware.HandleError(err)
		} else {
			_, err := insert.Exec(value[0], value[1], value[2], value[3], value[4])
			middleware.HandleError(err)
		}
	}
}

func InsertWorkspaceData(value [2]string) {
	db := CreateTable(2)
	q := "INSERT INTO Workspaces(emails, names) VALUES(?, ?)"
	insert, err := db.Prepare(q)
	middleware.HandleError(err)

	defer insert.Close()

	if len(value[0]) != 0 && len(value[1]) != 0 {
		_, err := insert.Exec(value[0], value[1])
		middleware.HandleError(err)
	}
}

func InsertMemberData(email string) {
	db := CreateTable(3)
	q := "INSERT INTO Members(emails, passwords, fullnames, accountnames, is_active, created_at) VALUES(?, ?, ?, ?, ?, ?)"
	insert, err := db.Prepare(q)
	middleware.HandleError(err)

	defer insert.Close()

	if len(email) != 0 {
		_, err := insert.Exec(email, "", "", "", "0", time.Now())
		middleware.HandleError(err)
	}
}
