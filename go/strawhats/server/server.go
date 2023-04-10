package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cclegg7/straw-hat-challenge/db"
	"github.com/cclegg7/straw-hat-challenge/models"
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
	fmt.Println("Serving!")
	return http.ListenAndServe(":8080", nil)
}

type GetUsersResponse struct {
	Users []*models.User `json:"users"`
}

func (s *Server) getUsersHandler(w http.ResponseWriter, _ *http.Request) {
	users, err := s.database.GetUsers()
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := &GetUsersResponse{
		Users: users,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}
