package cli

import (
  "fmt"
  "database/sql"
)

type User struct {
	ID		 int
	Username string
	Password string
	CreatedAt string
}

type API struct {
	DB *sql.DB
}

func login(db *sql.DB) string {
	fmt.Println("Login")

	var username, password string

	for {
		fmt.Print("Enter username: ")
		fmt.Scanln(&username)
		fmt.Print("Enter password: ")
		fmt.Scanln(&password)

		ok, err := FindUser(db, username, password)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if ok {
			fmt.Println("Login successful.")
			return username
		}

		fmt.Println("Invalid username/password.")
	}
}

func signup(db *sql.DB) {
	fmt.Println("Signup")

	var username, password string

	fmt.Print("Enter username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	err := CreateUser(db, username, password)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Signup successful.")
}

func FindUser(db *sql.DB, user string, pass string) (bool, error) {
	row := db.QueryRow(
		"SELECT COUNT(*) FROM users WHERE username = ? AND password = ?",
		user,
		pass,
	)

	var count int
	err := row.Scan(&count)

	if count > 0 {
		return true, err
	}

	return false, err
}