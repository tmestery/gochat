package main

import (
	"log"
	"net/rpc"
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Item struct {
	Title string
	Body string
}

func menuOptions() {
	fmt.Println("-------------gochat command options-------------\n1. Get Database\n2. Get Item By Name\n3. Add Item\n4. Edit Item\n5. Delete Item\n6. Exit")
}

func addItemOptions() (string, string) {
	var title, body string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("-------------adding an item-------------\nEnter item name:")
	title, _ = reader.ReadString('\n')
	title = strings.TrimSpace(title)
	
	fmt.Println("\nEnter item body:")
	body, _ = reader.ReadString('\n')
	body = strings.TrimSpace(body)

	return title, body
}

func getItemByNameOptions() (string, string) {
	var title, body string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("-------------get item by name-------------\nEnter item name:")
	title, _ = reader.ReadString('\n')
	title = strings.TrimSpace(title)

	return title, body
}

func deleteItemOptions() (string, string) {
	var title, body string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("-------------delete item by name-------------\nEnter item name you're deleting:")
	title, _ = reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Println("\nEnter item body you're deleting:")
	body, _ = reader.ReadString('\n')
	body = strings.TrimSpace(body)

	return title, body
}

func editItemOptions() (string, string) {
	var title, body string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("-------------edit item by name-------------\nEnter item name you're editing:")
	title, _ = reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Println("\nEnter edited body you want:")
	body, _ = reader.ReadString('\n')
	body = strings.TrimSpace(body)

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
			title, body = getItemByNameOptions()
			client.Call("API.GetByName", Item{title, body}, &reply)
		case 3:
			title, body = addItemOptions()
			client.Call("API.AddItem", Item{title, body}, &reply)
		case 4:
			title, body = editItemOptions()
			client.Call("API.EditItem", Item{title, body}, &reply)
		case 5:
			title, body = deleteItemOptions()
			client.Call("API.DeleteItem", Item{title, body}, &reply)
		default:
			fmt.Println("\nExiting gochat, thank you!")
			running = false
			break
		}
	}
}