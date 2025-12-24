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

func main() {
	var reply Item
	var db []Item
	var choice int

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection Error: ", err)
	}

	a := Item{"First", "A first item"}
	b := Item{"Second", "A second item"}

	fmt.Println("Commands:\n1. Add Item\n\n")
	fmt.Scan(&choice)

	if choice == 1 {
		client.Call("API.AddItem", a, &reply)
		client.Call("API.AddItem", b, &reply)
		client.Call("API.GetDB", "", &db)

		fmt.Println("Database: ", db)
	}
}