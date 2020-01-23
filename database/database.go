package database

import (
	"database/sql"
	"log"
)
import _ "github.com/mattn/go-sqlite3"

func Connect() *sql.DB {
	db, err := sql.Open("sqlite3", "guide.db")
	if err != nil {
		log.Fatalf("Error in connection to DB: %s\n", err)
	}
	return db
}
