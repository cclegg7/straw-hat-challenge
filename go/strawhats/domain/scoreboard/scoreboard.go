package scoreboard

import (
	"fmt"
	"sort"
	"time"

	"github.com/cclegg7/straw-hat-challenge/db"
	"github.com/cclegg7/straw-hat-challenge/models"
)

type Scoreboard struct {
	database *db.Database
	start    time.Time
	end      time.Time
}

func New(database *db.Database, start, end time.Time) *Scoreboard {
	return &Scoreboard{
		database: database,
		start:    start,
		end:      end,
	}
}

type Score struct {
	User   *models.User
	Points int
}

func (s *Scoreboard) AllScoresSorted() ([]*Score, error) {
	users, err := s.database.GetUsers()
	if err != nil {
		return nil, fmt.Errorf("error fetching users: %w", err)
	}

	var results []*Score
	for _, user := range users {
		results = append(results, &Score{
			User:   user,
			Points: 0,
		})
	}

	//climbingScores, err := s.climbingScores(users)
	//if err != nil {
	//	return nil, fmt.Errorf("error getting climbing scores: %w", err)
	//}
	//
	//var results []*Score
	//for user, climbingScore := range climbingScores {
	//	results = append(results, &Score{
	//		User:   user,
	//		Points: climbingScore,
	//	})
	//}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Points < results[j].Points
	})

	return results, nil
}
