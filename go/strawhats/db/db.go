package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/cclegg7/straw-hat-challenge/models"
	"github.com/go-sql-driver/mysql"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "localhost",
		DBName:               "straw_hat_challenge",
		AllowNativePasswords: true,
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

func (d *Database) GetUsers() ([]*models.User, error) {
	var users []*models.User
	rows, err := d.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("error querying for users: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, fmt.Errorf("error reading a user: %w", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error on users result: %w", err)
	}
	return users, nil
}
