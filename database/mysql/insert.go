package mysql

import (
	"log"
	"time"
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
