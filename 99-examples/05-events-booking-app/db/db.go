package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "events.db") // DB file name is events.db
	if err != nil {
		panic("Could not initialize database: " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable()
	createEventsTable()
	createRegistrationTable()
}

func createUsersTable() {
	createUsersTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)`

	_, err := DB.Exec(createUsersTableQuery)

	if err != nil {
		panic("Could not create users tables: " + err.Error())
	}

}

func createEventsTable() {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		date DATETIME NOT NULL,
		location TEXT not NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events tables: " + err.Error())
	}
}

func createRegistrationTable() {
	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		event_id INTEGER NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id),
		FOREIGN KEY(event_id) REFERENCES events(id)
	)`

	_, err := DB.Exec(createRegistrationTable)

	if err != nil {
		panic("Could not create registration tables: " + err.Error())
	}
}
