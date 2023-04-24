package models

type File struct {
	Token       string `json:"token"`
	URL         string `json:"url"`
	ContentType string `json:"content_type"`
}
