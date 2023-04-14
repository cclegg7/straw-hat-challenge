package models

import "time"

type Climb struct {
	UserID      int `json:"user_id"`
	Category    ClimbCategory `json:"category"`
	Rating      int `json:"rating"`
	Date        time.Time `json:"date"`
	IsChallenge bool `json:"is_challenge"`
	CreatedAt   time.Time `json:"created_at"`
}
