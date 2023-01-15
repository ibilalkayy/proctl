package mysql

import "log"

func UpdateStatus(status, email, password string) {
	db := Connect()
	q := "UPDATE Signup SET is_active=? WHERE emails=? AND passwords=?"
	update, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}

	_, err = update.Exec(status, email, password)
	if err != nil {
		log.Fatal(err)
	}
}
