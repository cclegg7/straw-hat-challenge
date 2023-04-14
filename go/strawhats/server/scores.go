package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cclegg7/straw-hat-challenge/domain/scoreboard"
)

type scoreEntry struct {
	Rank           int    `json:"rank"`
	UserName       string `json:"user_name"`
	UserID         int    `json:"user_id"`
	CharacterToken string `json:"character_token"`
	Score          int    `json:"score"`
}

type getScoresResponse struct {
	Scores []*scoreEntry `json:"scores"`
}

func (res *getScoresResponse) fromScores(scores []*scoreboard.Score) {
	for i, score := range scores {
		res.Scores = append(res.Scores, &scoreEntry{
			Rank:           i + 1,
			UserName:       score.User.Name,
			UserID:         score.User.ID,
			CharacterToken: score.User.CharacterToken,
			Score:          score.Points,
		})
	}
}

func (s *Server) getScoresHandler(w http.ResponseWriter, _ *http.Request) {
	startTime, _ := time.Parse("2006-01-02", "2023-04-10")
	endTime, _ := time.Parse("2006-01-02", "2023-05-10")
	scoreCalculator := scoreboard.New(s.database, startTime, endTime)

	scores, err := scoreCalculator.AllScoresSorted()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response := &getScoresResponse{}
	response.fromScores(scores)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(jsonResponse)
}
