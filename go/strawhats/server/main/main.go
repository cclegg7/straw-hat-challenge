package main

import (
	"fmt"
	"os"

	"github.com/cclegg7/straw-hat-challenge/db"
	"github.com/cclegg7/straw-hat-challenge/server"
)

func main() {
	database, err := db.NewDatabase()
	if err != nil {
		fmt.Printf("\nerror connecting to database: %v", err.Error())
		os.Exit(0)
	}

	server := server.NewServer(database)
	if err := server.Serve(); err != nil {
		fmt.Printf("\nerror starting server: %v", err.Error())
		os.Exit(0)
	}
}
