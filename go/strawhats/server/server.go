package server

import (
	"fmt"
	"net/http"

	"github.com/cclegg7/straw-hat-challenge/db"
)

type Server struct {
	database *db.Database
}

func NewServer(database *db.Database) *Server {
	return &Server{
		database: database,
	}
}

func (s *Server) Serve() error {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/users", s.getUsersHandler)
	http.HandleFunc("/climb", s.postClimbHandler)
	http.HandleFunc("/run", s.postRunHandler)
	http.HandleFunc("/scores", s.getScoresHandler)
	http.HandleFunc("/runs", s.listUserRunsHandler)
	http.HandleFunc("/climbs", s.listUserClimbsHandler)
	fmt.Println("Serving!")
	return http.ListenAndServe(":8080", nil)
}
