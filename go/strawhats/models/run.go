package models

import "time"

type Run struct {
	UserID   int
	Distance int
	Date     time.Time
}
