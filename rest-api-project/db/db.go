package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var dbError error

func InitDB() {
	DB, dbError = sql.Open("sqlite3", "api.db")
	if dbError != nil {
		panic(dbError)
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		date DATETIME,
		user_id INTEGER
	)
	`
	_, err := DB.Exec(createEventTable)
	if err != nil {
		panic(err)
	}
}
