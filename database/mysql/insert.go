package mysql

import (
	"errors"
	"fmt"
	"time"

	"github.com/ibilalkayy/proctl/middleware"
)

// InsertSignupData inserts signup data into the database.
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

// InsertProfileData inserts profile data into the database.
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

// InsertWorkspaceData inserts workspace data into the database.
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

// InsertMemberData inserts member data into the database.
func InsertMemberData(email string) {
	db := CreateTable(3)
	q := "INSERT INTO Members(emails, passwords, fullnames, accountnames, titles, phones, locations, working_statuses, is_active, created_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	insert, err := db.Prepare(q)
	middleware.HandleError(err)

	defer insert.Close()

	if len(email) != 0 {
		_, err := insert.Exec(email, "", "", "", "", "", "", "", "0", time.Now())
		middleware.HandleError(err)
	}
}

// InsertDepartment inserts department data into the database.
func InsertDepartment(email, department string) {
	db := CreateTable(4)
	q := "INSERT INTO Departments(emails, departments) VALUES(?, ?)"
	insert, err := db.Prepare(q)
	middleware.HandleError(err)

	defer insert.Close()

	if len(email) != 0 {
		_, err := insert.Exec(email, department)
		middleware.HandleError(err)
	}
}

// InsertRole inserts role data into the database.
func InsertRole(email, role string) {
	db := CreateTable(5)
	q := "INSERT INTO Roles(emails, roles) VALUES(?, ?)"
	insert, err := db.Prepare(q)
	middleware.HandleError(err)

	defer insert.Close()

	if len(email) != 0 {
		_, err := insert.Exec(email, role)
		middleware.HandleError(err)
	}
}

// InsertBoard inserts board data into the database.
func InsertBoard(email, board string) {
	db := CreateTable(6)
	q := "INSERT INTO Boards(emails, boards) VALUES(?, ?)"
	insert, err := db.Prepare(q)
	middleware.HandleError(err)

	defer insert.Close()

	if len(email) != 0 {
		_, err := insert.Exec(email, board)
		middleware.HandleError(err)
	}
}
