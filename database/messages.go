package database

import "database/sql"

type Message struct {
	ID        int
	Username  string
	Body      string
	CreatedAt string
}

func GetMessages(db *sql.DB) ([]Message, error) {
	rows, err := db.Query(
		"SELECT id, username, body, created_at FROM messages ORDER BY id DESC",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var msgs []Message
	for rows.Next() {
		var m Message
		rows.Scan(&m.ID, &m.Username, &m.Body, &m.CreatedAt)
		msgs = append(msgs, m)
	}
	
	return msgs, nil
}