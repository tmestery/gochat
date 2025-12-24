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

func main() {
	var reply Item
	var db []Item
	var choice int

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection Error: ", err)
	}

	var running bool = true
	var response int

	for running {
		menuOptions()
		fmt.Scan(&response)

		switch response {
		case 1:
			fmt.Println("Database: ", db)
		case 2:

		case 3:

		case 4:

		case 5:
		
		default:
			fmt.Println("\nExiting gochat, thank you!")
			running = false
			break
		}
	}




	a := Item{"1st", ""}
	b := Item{"2nd", "A second item"}

	if choice == 1 {
		client.Call("API.AddItem", a, &reply)
		client.Call("API.AddItem", b, &reply)
		client.Call("API.GetDB", "", &db)
	}
}