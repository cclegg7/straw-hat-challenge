package server

import (
	"encoding/json"
	"net/http"

	"github.com/cclegg7/straw-hat-challenge/models"
)

type GetUsersResponse struct {
	Users []*models.User `json:"users"`
}

func (s *Server) getUsersHandler(w http.ResponseWriter, _ *http.Request) {
	users, err := s.database.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response := &GetUsersResponse{
		Users: users,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(jsonResponse)
}
