package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/cclegg7/straw-hat-challenge/models"
)

type postRunRequest struct {
	UserID    int    `json:"user_id"`
	Distance  int    `json:"distance"`
	Date      string `json:"date"`
	FileToken string `json:"file_token"`
}

func (req *postRunRequest) toModel() *models.Run {
	time, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		fmt.Printf("error parsing time: %+v", err)
	}
	return &models.Run{
		UserID:   req.UserID,
		Distance: req.Distance,
		Date:     time,
	}
}

func (req *postRunRequest) fromHTTPRequest(httpReq *http.Request) error {
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

func (s *Server) postRunHandler(w http.ResponseWriter, httpReq *http.Request) {
	req := &postRunRequest{}
	if err := req.fromHTTPRequest(httpReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	runModel := req.toModel()
	runID, err := s.database.CreateRun(runModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if len(req.FileToken) > 0 {
		if err := s.database.LinkFileToRun(req.FileToken, runID); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}

	w.Write([]byte("success"))
}
