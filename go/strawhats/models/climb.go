package models

import "time"

type Climb struct {
	UserID      int
	Category    ClimbCategory
	Rating      int
	Date        time.Time
	IsChallenge bool
	Notes       string
}
