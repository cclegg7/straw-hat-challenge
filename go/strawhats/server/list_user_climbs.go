package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cclegg7/straw-hat-challenge/models"
)

type ListUserClimbsResponse struct {
	Climbs []*models.Climb `json:"climbs"`
}

func (s *Server) listUserClimbsHandler(w http.ResponseWriter, httpReq *http.Request) {
	userIdString := httpReq.URL.Query().Get("user_id")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return

	}

	categoryString := httpReq.URL.Query().Get("category")
	category, err := strconv.Atoi(categoryString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return

	}

	climbs, err := s.database.ListUserClimbs(userId, category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response := &ListUserClimbsResponse{
		Climbs: climbs,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(jsonResponse)
}



