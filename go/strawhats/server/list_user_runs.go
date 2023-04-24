package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cclegg7/straw-hat-challenge/models"
)

type ListUserRunsResponse struct {
	Runs []*models.Run `json:"runs"`
}

func (s *Server) listUserRunsHandler(w http.ResponseWriter, httpReq *http.Request) {
	userIdString := httpReq.URL.Query().Get("user_id")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	runs, err := s.database.ListUserRunsWithFiles(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response := &ListUserRunsResponse{
		Runs: runs,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(jsonResponse)
}
