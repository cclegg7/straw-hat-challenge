package main

import (
	"fmt"
	"github.com/cclegg7/straw-hat-challenge/configs"
	"github.com/cclegg7/straw-hat-challenge/database"
	"os"

	"github.com/cclegg7/straw-hat-challenge/clients/aws"

	"github.com/cclegg7/straw-hat-challenge/server"
)

func main() {
	configs, err := configs.New()
	if err != nil {
		fmt.Printf("\nerror reading configs: %v", err.Error())
		os.Exit(0)
	}

	database, err := database.New(configs.Database)
	if err != nil {
		fmt.Printf("\nerror connecting to database: %v", err.Error())
		os.Exit(0)
	}

	s3, err := aws.NewS3Client(configs.FileStorage)
	if err != nil {
		fmt.Printf("\nerror creating s3 client: %v", err.Error())
		os.Exit(0)
	}

	server := server.New(database, s3, configs.Server)
	if err := server.Serve(); err != nil {
		fmt.Printf("\nerror starting server: %v", err.Error())
		os.Exit(0)
	}
}
