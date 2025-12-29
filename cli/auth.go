package cli

import (
  "fmt"
)

type User struct {
	Username string
	Password string
}

var userDatabase []User

func login() string {
	fmt.Println("Login")

	var username string
	var password string
	var running bool = true
	
	for running {
		fmt.Print("Enter username: ")
		fmt.Scanln(&username)
		fmt.Print("Enter password: ")
		fmt.Scanln(&password)

		var loginUser User = User{Username: username, Password: password}

		found := false
		for _, user := range userDatabase {
			if loginUser == user {
				fmt.Println("Logging Successful.")
				running = false
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Invalid username/password.")
		}
	}

	return username
}

func signup() {
	fmt.Println("Signup")

	var username string
	var password string

	fmt.Print("Enter username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	var signedUpUser User = User{Username: username, Password: password}
	userDatabase = append(userDatabase, signedUpUser)

	fmt.Println("Signup Successful.")
}