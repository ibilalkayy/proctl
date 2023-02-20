package mysql

import (
	"fmt"

	"github.com/ibilalkayy/proctl/middleware"
)

type UserCredentials struct {
	Email    string
	Password string
	Status   string
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

func FindWorkspace(email string) string {
	db := Connect()
	q := "SELECT names FROM Workspace WHERE emails=?"
	rows, err := db.Query(q, email)
	middleware.HandleError(err)

	defer rows.Close()

	for rows.Next() {
		var Names string
		err := rows.Scan(&Names)
		middleware.HandleError(err)

		fmt.Println(Names)
	}
	return ""
}
