package mysql

import (
	"database/sql"
	"fmt"
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

func FindAccount(email, password string) (string, string, string, bool) {
	db := Connect()
	var uc UserCredentials
	q := "SELECT emails, passwords, is_active FROM Signup WHERE emails=? and passwords=?"
	if err := db.QueryRow(q, email, password).Scan(&uc.Email, &uc.Password, &uc.Status); err != nil {
		return "", "", "", false
	}

	return uc.Email, uc.Password, uc.Status, true
}

func FindProfile(email string) bool {
	db := Connect()
	var pc ProfileCredentials
	q := "SELECT titles, phones, locations, working_statuses FROM Profiles WHERE emails=?"
	if err := db.QueryRow(q, email).Scan(&pc.Title, &pc.Phone, &pc.Location, &pc.Working_status); err != nil {
		return false
	}
	return true
}

func queryUser(db *sql.DB, q string, args ...interface{}) UserCredentials {
	var uc UserCredentials
	row := db.QueryRow(q, args...)
	err := row.Scan(&uc.Email, &uc.FullName, &uc.AccountName)
	if err != nil && err != sql.ErrNoRows {
		fmt.Printf("")
	}

	return uc
}

func queryProfile(db *sql.DB, q string, args ...interface{}) ProfileCredentials {
	var pc ProfileCredentials
	row := db.QueryRow(q, args...)
	err := row.Scan(&pc.Title, &pc.Phone, &pc.Location)
	if err != nil && err != sql.ErrNoRows {
		fmt.Printf("")
	}

	return pc
}

func ListUserInfo(email, password string) string {
	db := Connect()
	q := "SELECT emails, fullnames, accountnames FROM Signup WHERE emails=? AND passwords=?"
	uc := queryUser(db, q, email, password)
	fmt.Printf("Email Address: %s\nFull Name: %s\nAccount Name: %s\n", uc.Email, uc.FullName, uc.AccountName)
	return ""
}

func ListProfileInfo(email string) string {
	db := Connect()
	q := "SELECT titles, phones, locations FROM Profiles WHERE emails=?"
	pc := queryProfile(db, q, email)
	fmt.Printf("Title: %s\nPhone Number: %s\nLocation: %s\n", pc.Title, pc.Phone, pc.Location)
	return ""
}

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

func FindWorkspaceName(email, name string) string {
	db := Connect()
	var Name string
	q := "SELECT names FROM Workspaces WHERE emails=? AND names=?"
	if err := db.QueryRow(q, email, name).Scan(&Name); err != nil {
		return ""
	}
	return Name
}

func FindMember(email, password string) ([3]string, bool) {
	db := Connect()
	var uc UserCredentials
	q := "SELECT emails, passwords, is_active FROM Members WHERE emails=?"
	args := []interface{}{email}
	if password != "" {
		q += " AND passwords=?"
		args = append(args, password)
	}
	if err := db.QueryRow(q, args...).Scan(&uc.Email, &uc.Password, &uc.Status); err != nil {
		return [3]string{}, false
	}

	memberCredentials := [3]string{uc.Email, "", uc.Status}
	if password != "" {
		memberCredentials[1] = uc.Password
	}
	return memberCredentials, true
}
