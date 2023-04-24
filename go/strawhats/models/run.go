package models

import "time"

type Run struct {
	UserID    int       `json:"user_id"`
	Distance  int       `json:"distance"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	Files     []*File   `json:"files"`
}
