package main

import (
	"database/sql"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"github.com/tmestery/gochat/database"
)

type API struct {
	DB *sql.DB
}

// get all messages
func (a *API) GetMessages(_ struct{}, reply *[]database.Message) error {
	msgs, err := database.GetMessages(a.DB)
	if err != nil {
		return err
	}

	*reply = msgs
	return nil
}

// add a message
func (a *API) AddMessage(msg database.Message, reply *database.Message) error {
	err := database.InsertMessage(a.DB, msg.Username, msg.Body)
	if err != nil {
		return err
	}

	*reply = msg
	return nil
}

func main() {
	db, err := database.Open()
	if err != nil {
		log.Fatal(err)
	}
	
	defer db.Close()

	if err := database.Init(db); err != nil {
		log.Fatal(err)
	}

	api := &API{DB: db}

	if err := rpc.Register(api); err != nil {
		log.Fatal(err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("RPC server running on :4040")
	log.Fatal(http.Serve(listener, nil))
}