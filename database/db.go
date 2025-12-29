package database

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "messages.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Init(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		body TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(query)
	return err
}

func InsertMessage(db *sql.DB, user, body string) error {
	_, err := db.Exec(
		"INSERT INTO messages (username, body) VALUES (?, ?)",
		user, body,
	)
	return err
}