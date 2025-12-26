package main

import (
  "fmt"
)

type User struct {
	Username string
	Password string
}

var userDatabase []User

func login() {
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

		for _, user := range userDatabase {
			if loginUser == user {
				fmt.Println("Logging Successful.")
				running = false
				break
			}
		}

		fmt.Println("Invalid username/password.")
	}
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