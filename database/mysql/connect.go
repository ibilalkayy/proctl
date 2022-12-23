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

func Connect() (db *sql.DB) {
	db_user := middleware.LoadEnvVariable("DB_USER")
	db_password := middleware.LoadEnvVariable("DB_PASSWORD")
	db_address := middleware.LoadEnvVariable("DB_ADDRESS")
	db_db := middleware.LoadEnvVariable("DB_DB")
	s := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_password, db_address, db_db)
	db, err := sql.Open("mysql", s)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CreateTable(number int) (db *sql.DB) {
	db = Connect()
	query, err := ioutil.ReadFile("database/mysql/db.SQL")
	if err != nil {
		log.Fatal(err)
	}

	requests := strings.Split(string(query), ";")[number]

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
