package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/cclegg7/straw-hat-challenge/models"
)

type postClimbRequest struct {
	UserID      int    `json:"user_id"`
	Category    int    `json:"category"`
	Rating      int    `json:"rating"`
	Date        string `json:"date"`
	IsChallenge bool   `json:"is_challenge"`
	Notes       string `json:"notes"`
}

func (req *postClimbRequest) toModel() *models.Climb {
	time, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		fmt.Printf("error parsing time: %+v", err)
	}
	return &models.Climb{
		UserID:      req.UserID,
		Category:    models.ClimbCategory(req.Category),
		Rating:      req.Rating,
		IsChallenge: req.IsChallenge,
		Date:        time,
	}
}

func (req *postClimbRequest) fromHTTPRequest(httpReq *http.Request) error {
	reqBody, err := io.ReadAll(httpReq.Body)
	if err != nil {
		return fmt.Errorf("error creating reading request body: %w", err)
	}

	err = json.Unmarshal(reqBody, req)
	if err != nil {
		return fmt.Errorf("error unmarshaling request: %w", err)
	}

	return nil
}

func (s *Server) postClimbHandler(w http.ResponseWriter, httpReq *http.Request) {
	req := &postClimbRequest{}
	if err := req.fromHTTPRequest(httpReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	climbModel := req.toModel()
	fmt.Printf("model: %+v", climbModel)
	if err := s.database.CreateClimb(climbModel); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("success"))
}
