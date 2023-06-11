package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/ibilalkayy/proctl/middleware"
)

type UserCredentials struct {
	Email       string
	Password    string
	Status      string
	FullName    string
	AccountName string
}

type ProfileCredentials struct {
	Title          string
	Phone          string
	Location       string
	Working_status string
}

// FindAccount finds an account with the given email and password.
// It returns the email, password, status, and a boolean indicating if the account was found.
func FindAccount(email, password string) (string, string, string, bool) {
	db := Connect()
	var uc UserCredentials
	q := "SELECT emails, passwords, is_active FROM Signup WHERE emails=? and passwords=?"
	if err := db.QueryRow(q, email, password).Scan(&uc.Email, &uc.Password, &uc.Status); err != nil {
		return "", "", "", false
	}

	return uc.Email, uc.Password, uc.Status, true
}

// FindProfile checks if a profile exists for the given email.
// It returns a boolean indicating if the profile exists.
func FindProfile(email string) bool {
	db := Connect()
	var pc ProfileCredentials
	q := "SELECT titles, phones, locations, working_statuses FROM Profiles WHERE emails=?"
	if err := db.QueryRow(q, email).Scan(&pc.Title, &pc.Phone, &pc.Location, &pc.Working_status); err != nil {
		return false
	}
	return true
}

// queryUser performs a database query and returns user credentials.
func queryUser(db *sql.DB, q string, args ...interface{}) UserCredentials {
	var uc UserCredentials
	row := db.QueryRow(q, args...)
	err := row.Scan(&uc.Email, &uc.FullName, &uc.AccountName)
	if err != nil && err != sql.ErrNoRows {
		fmt.Printf("")
	}

	return uc
}

// queryProfile performs a database query and returns profile credentials.
func queryProfile(db *sql.DB, q string, args ...interface{}) ProfileCredentials {
	var pc ProfileCredentials
	row := db.QueryRow(q, args...)
	err := row.Scan(&pc.Title, &pc.Phone, &pc.Location)
	if err != nil && err != sql.ErrNoRows {
		fmt.Printf("")
	}

	return pc
}

// ListUserInfo lists user information for the given email and password.
// It returns an empty string.
func ListUserInfo(email, password string) string {
	db := Connect()
	q := "SELECT emails, fullnames, accountnames FROM Signup WHERE emails=? AND passwords=?"
	uc := queryUser(db, q, email, password)
	fmt.Printf("Email Address: %s\nFull Name: %s\nAccount Name: %s\n", uc.Email, uc.FullName, uc.AccountName)
	return ""
}

// ListProfileInfo lists profile information for the given email.
// It returns an empty string.
func ListProfileInfo(email string) string {
	db := Connect()
	q := "SELECT titles, phones, locations FROM Profiles WHERE emails=?"
	pc := queryProfile(db, q, email)
	fmt.Printf("Title: %s\nPhone Number: %s\nLocation: %s\n", pc.Title, pc.Phone, pc.Location)
	return ""
}

// FindWorkspace finds workspaces associated with the given email.
// It returns a string containing the list of workspace names.
func FindWorkspace(email string) string {
	db := Connect()
	q := "SELECT names FROM Workspaces WHERE emails=?"
	rows, err := db.Query(q, email)
	middleware.HandleError(err)

	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		middleware.HandleError(err)

		names = append(names, name)
	}
	listOfName := strings.Join(names, "\n")
	if len(listOfName) != 0 {
		return listOfName
	} else {
		return "This account has no workspaces"
	}
}

// FindWorkspaceName finds the workspace with the given email and name.
// It returns the workspace name if found, otherwise an empty string.
func FindWorkspaceName(email, name string) string {
	db := Connect()
	var Name string
	q := "SELECT names FROM Workspaces WHERE emails=? AND names=?"
	if err := db.QueryRow(q, email, name).Scan(&Name); err != nil {
		return ""
	}
	return Name
}

// FindMember finds a member with the given email and password.
// It returns member credentials and a boolean indicating if the member was found.
func FindMember(email, password string) ([4]string, bool) {
	db := Connect()
	var uc UserCredentials
	q := "SELECT emails, passwords, fullnames, accountnames FROM Members WHERE emails=?"
	args := []interface{}{email}
	if password != "" {
		q += " AND passwords=?"
		args = append(args, password)
	}
	if err := db.QueryRow(q, args...).Scan(&uc.Email, &uc.Password, &uc.FullName, &uc.AccountName); err != nil {
		if err == sql.ErrNoRows {
			return [4]string{}, false
		} else {
			log.Fatal(err)
		}
	}

	memberCredentials := [4]string{uc.Email, "", uc.FullName, uc.AccountName}
	if password != "" {
		memberCredentials[1] = uc.Password
	}
	return memberCredentials, true
}

// FindDepartment finds the department associated with the given email.
// It returns the email and department as a string array.
func FindDepartment(email string) [2]string {
	db := Connect()
	var Email, Department string
	q := "SELECT emails, departments FROM Departments WHERE emails=?"
	if err := db.QueryRow(q, email).Scan(&Email, &Department); err != nil {
		return [2]string{}
	}
	credentials := [2]string{Email, Department}
	return credentials
}

// FindRole finds the role associated with the given email.
// It returns the email and role as a string array.
func FindRole(email string) [2]string {
	db := Connect()
	var Email, Role string
	q := "SELECT emails, roles FROM Roles WHERE emails=?"
	if err := db.QueryRow(q, email).Scan(&Email, &Role); err != nil {
		return [2]string{}
	}
	credentials := [2]string{Email, Role}
	return credentials
}

// FindBoard finds a board with the given email and board name.
// It returns the board name if found, otherwise an empty string.
func FindBoard(email, board string) string {
	db := Connect()
	var Board string
	q := "SELECT boards FROM Boards WHERE emails=? AND boards=?"
	if err := db.QueryRow(q, email, board).Scan(&Board); err != nil {
		return ""
	}
	return Board
}

func FindProject(email, board, project string) string {
	db := CreateTable(7)
	var Project string
	q := "SELECT projects FROM Projects WHERE emails=? AND boards=? AND projects=?"
	if err := db.QueryRow(q, email, board, project).Scan(&Project); err != nil {
		if err == sql.ErrNoRows {
			return ""
		} else {
			return ""
		}
	}
	return Project
}
