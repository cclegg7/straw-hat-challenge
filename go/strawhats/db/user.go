package db

import (
	"fmt"

	"github.com/cclegg7/straw-hat-challenge/models"
)

func (d *Database) GetUsers() ([]*models.User, error) {
	var users []*models.User
	rows, err := d.db.Query("SELECT u.id, u.name, u.toprope, u.boulder, c.token FROM users u JOIN characters c ON (u.character_id = c.id)")
	if err != nil {
		return nil, fmt.Errorf("error querying for users: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.TopRope, &user.Boulder, &user.CharacterToken); err != nil {
			return nil, fmt.Errorf("error reading a user: %w", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error on users result: %w", err)
	}
	return users, nil
}
