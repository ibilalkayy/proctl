package mysql

import "log"

func CountTableColumns(table string) int {
	db := Connect()
	var counted int
	q := "SELECT COUNT(*) FROM " + table
	if err := db.QueryRow(q).Scan(&counted); err != nil {
		log.Fatal(err)
	}
	return counted
}
