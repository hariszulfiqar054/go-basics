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
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`

	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`

	_, err := DB.Exec(createEventTable)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(createUserTable)
	if err != nil {
		panic(err)
	}
}
