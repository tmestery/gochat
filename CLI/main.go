package main

import (
  "fmt"
  "log"
  "bufio"
  "os"
  "strings"
	"github.com/tmestery/gochat/client"
  "github.com/tmestery/gochat/cli"
)

type Item struct {
	Title string
	Body  string
}

func main() {
  displayIntro()
  loginSignup()
  
  c, err := client.New("localhost:4040")
  if err != nil {
	log.Fatal(err)
  }

  var reply client.Item
  var db []client.Item
  var title, body string
  running := true

  for running {
      menuOptions()
      var choice int
      fmt.Scanln(&choice)

      switch choice {
      case 1:
          c.GetDB(&db)
          fmt.Println("Database: ", db)
      case 2:
          title, body = getItemByNameOptions()
          c.GetByName(client.Item{Title: title, Body: body}, &reply)
          fmt.Println("Item: ", reply)
      case 3:
          title, body = addItemOptions()
          c.AddItem(client.Item{Title: title, Body: body}, &reply)
          fmt.Println("Added: ", reply)
      case 4:
          title, body = editItemOptions()
          c.EditItem(client.Item{Title: title, Body: body}, &reply)
          fmt.Println("Edited: ", reply)
      case 5:
          title, body = deleteItemOptions()
          c.DeleteItem(client.Item{Title: title, Body: body}, &reply)
          fmt.Println("Deleted: ", reply)
      default:
          fmt.Println("\nExiting gochat, thank you!")
          running = false
    }
  }
}

func loginSignup() {
  var option int
  var running bool

  for running {
    fmt.Println("1. Login\n2. Signup")
    fmt.Scanln(&option)

    if option == 1 {
      signup()
    }

    login()
    break
  }
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

func displayIntro() {
	fmt.Println(`   ██████╗  ██████╗  ██████╗██╗  ██╗ █████╗ ████████╗
  ██╔════╝ ██╔═══██╗██╔════╝██║  ██║██╔══██╗╚══██╔══╝
  ██║  ███╗██║   ██║██║     ███████║███████║   ██║   
  ██║   ██║██║   ██║██║     ██╔══██║██╔══██║   ██║   
  ╚██████╔╝╚██████╔╝╚██████╗██║  ██║██║  ██║   ██║   
   ╚═════╝  ╚═════╝  ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   
                    real-time CLI chat`)

	fmt.Println("Welcome to gochat!")
}