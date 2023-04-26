package models

type Score struct {
	User   *User `json:"user"`
	Points int   `json:"points"`
}
