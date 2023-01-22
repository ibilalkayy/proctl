package mysql

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ibilalkayy/proctl/middleware"
)

func InsertSignupData(value [4]string) {
	db := CreateTable(0)
	q := "INSERT INTO Signup(emails, passwords, fullnames, accountnames, is_active, created_at) VALUES(?, ?, ?, ?, ?, ?)"
	insert, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}

	defer insert.Close()

	if len(value[0]) != 0 || len(value[1]) != 0 {
		currentTime := time.Now()
		_, err := insert.Exec(value[0], value[1], value[2], value[3], 0, currentTime)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func InsertProfileData(value [4]string) {
	db := CreateTable(1)
	q := "INSERT INTO Profiles(titles, phones, locations, working_statuses) VALUES(?, ?, ?, ?)"
	insert, err := db.Prepare(q)
	middleware.HandleError(err)

	defer insert.Close()

	insertMessage := "Your profile data is successfully inserted."

	profileFound := FindProfile(value[0], value[1])
	if profileFound {
		fmt.Println("Your profile data is already inserted. Type 'proctl update [flags]'")
	} else {
		if len(value[0]) == 0 && len(value[1]) == 0 && len(value[2]) == 0 && len(value[3]) == 0 {
			fmt.Println(errors.New("Enter the profile information first."))
		} else if len(value[0]) == 0 {
			_, err := insert.Exec("", value[1], value[2], value[3])
			middleware.HandleError(err)
			fmt.Println(insertMessage)
		} else if len(value[1]) == 0 {
			_, err := insert.Exec(value[0], "", value[2], value[3])
			middleware.HandleError(err)
			fmt.Println(insertMessage)
		} else if len(value[2]) == 0 {
			_, err := insert.Exec(value[0], value[1], "", value[3])
			middleware.HandleError(err)
			fmt.Println(insertMessage)
		} else if len(value[3]) == 0 {
			_, err := insert.Exec(value[0], value[1], value[2], "")
			middleware.HandleError(err)
			fmt.Println(insertMessage)
		} else {
			_, err := insert.Exec(value[0], value[1], value[2], value[3])
			middleware.HandleError(err)
			fmt.Println(insertMessage)
		}
	}
}
