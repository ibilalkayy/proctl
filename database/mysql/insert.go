package mysql

import "log"

func InsertSignupData(value [4]string) {
	db := CreateTable(0)
	q := "INSERT INTO Signup(emails, passwords, fullnames, accountnames) VALUES(?, ?, ?, ?)"
	insert, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}

	defer insert.Close()

	if len(value[0]) != 0 || len(value[1]) != 0 {
		_, err := insert.Exec(value[0], value[1], value[2], value[3])
		if err != nil {
			log.Fatal(err)
		}
	}
}
