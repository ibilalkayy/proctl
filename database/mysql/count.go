package mysql

import "log"

// CountTableColumns counts the number of columns in a table in the MySQL database
func CountTableColumns(table string) int {
	// Connect to the MySQL database
	db := Connect()

	var counted int
	q := "SELECT COUNT(*) FROM " + table

	// Execute the query and retrieve the count
	if err := db.QueryRow(q).Scan(&counted); err != nil {
		log.Fatal(err)
	}

	return counted
}
