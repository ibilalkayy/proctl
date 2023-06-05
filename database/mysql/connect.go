package mysql

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ibilalkayy/proctl/middleware"
)

// Connect establishes a connection to the MySQL database
func Connect() (db *sql.DB) {
	// Load database connection details from environment variables
	dbUser := middleware.LoadEnvVariable("DB_USER")
	dbPassword := middleware.LoadEnvVariable("DB_PASSWORD")
	dbAddress := middleware.LoadEnvVariable("DB_ADDRESS")
	dbDB := middleware.LoadEnvVariable("DB_DB")

	// Create the connection string
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbAddress, dbDB)

	// Open a connection to the MySQL database
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// CreateTable creates a table in the MySQL database using SQL queries from a file
func CreateTable(number int) (db *sql.DB) {
	// Connect to the MySQL database
	db = Connect()

	// Read the SQL queries from the file
	query, err := ioutil.ReadFile("database/mysql/db.SQL")
	if err != nil {
		log.Fatal(err)
	}

	// Split the queries into separate requests
	requests := strings.Split(string(query), ";")[number]

	// Prepare and execute each request
	stmt, err := db.Prepare(requests)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
