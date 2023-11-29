package utils

import (
	"database/sql"

    _ "github.com/mattn/go-sqlite3"
)

// DBConn return MySQL stable connection
func DBConn() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "./local.db")
	if err != nil {
		panic(err.Error())
	}
    return db
}
