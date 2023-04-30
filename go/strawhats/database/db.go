package database

import (
	"database/sql"
	"github.com/cclegg7/straw-hat-challenge/configs"
	"log"

	"github.com/go-sql-driver/mysql"
)

type Database struct {
	db *sql.DB
}

func New(configs *configs.Database) (*Database, error) {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 configs.User,
		Passwd:               configs.Password,
		Net:                  "tcp",
		Addr:                 configs.Hostname,
		DBName:               configs.Name,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	// Get a database handle.
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return &Database{
		db: db,
	}, nil
}
