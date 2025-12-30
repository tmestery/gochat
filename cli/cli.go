package cli

import (
  "fmt"
  "log"
  "bufio"
  "os"
  "strings"
  "github.com/tmestery/gochat/Client"
  "database/sql"
)

func Runner() {
	displayIntro()

	db, err := Open()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := Init(db); err != nil {
		log.Fatal(err)
	}

	username := loginSignup(db)

	c, err := client.New("localhost:4040")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	running := true

	for running {
		menuOptions()

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var msgs []client.Message
			err := c.GetMessages(&msgs)
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Println("\n--- Messages ---")
			for _, m := range msgs {
				fmt.Printf("[%s]: %s\n", m.Username, m.Body)
			}

		case 2:
			user := username

			fmt.Print("Message: ")
			body, _ := reader.ReadString('\n')
			body = strings.TrimSpace(body)

			var reply client.Message
			err := c.AddMessage(client.Message{
				Username: user,
				Body:     body,
			}, &reply)

			if err != nil {
				log.Println(err)
			} else {
				fmt.Println("Message sent.")
			}

		default:
			fmt.Println("Exiting gochat.")
			running = false
		}
	}
}

func loginSignup(db *sql.DB) string {
  var option int
  var username string
  running := true

  for running {
    fmt.Println("1. Login\n2. Signup")
    fmt.Scanln(&option)

    if option == 1 {
      signup(db)
    }

    username = login(db)
    break
  }

  return username
}

func menuOptions() {
	fmt.Println(`
------------- gochat -------------
1. View messages
2. Send message
3. Exit
`)
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