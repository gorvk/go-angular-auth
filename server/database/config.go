package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Setup() {
	connStr := "user=gorvk dbname=test_db password=1212"
	dbInstance, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	db = dbInstance
}

func GetDbInstance() *sql.DB {
	return db
}
