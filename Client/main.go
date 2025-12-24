package main

import (
	"log"
	"net/rpc"
	"fmt"
)

type Item struct {
	Title string
	Body string
}

func menuOptions() {
	fmt.Println("\n-------------gochat command options-------------\n1. Get Database\n2. Get Item By Name\n3. Add Item\n4. Edit Item\n5. Delete Item\n6. Exit")
}

func addItemOptions() (string, string) {
	var title, body string

	fmt.Println("\n-------------adding an item-------------\nEnter item name:\n")
	fmt.Scan(&title)
	fmt.Println("\nEnter item body:\n")
	fmt.Scan(&body)

	fmt.Println("Body: ", body, "Title: ", title)
	return title, body
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection Error: ", err)
	}

	var running bool = true
	var response int
	var title, body string

	for running {
		menuOptions()
		fmt.Scan(&response)

		switch response {
		case 1:
			client.Call("API.GetDB", "", &db)
			fmt.Println("Database: ", db)
		case 2:

		case 3:
			title, body = addItemOptions()
			client.Call("API.AddItem", Item{title, body}, &reply)
		case 4:

		case 5:
		
		default:
			fmt.Println("\nExiting gochat, thank you!")
			running = false
			break
		}
	}
}