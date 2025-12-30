package cli

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "users.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Init(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(query)
	return err
}

func CreateUser(db *sql.DB, username string, password string) error {
	_, err := db.Exec(
		"INSERT INTO users (username, password) VALUES (?, ?)",
		username, password,
	)
	return err
}