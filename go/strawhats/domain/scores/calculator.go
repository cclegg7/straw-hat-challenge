package scores

import (
	"github.com/cclegg7/straw-hat-challenge/database"
	"sort"

	"github.com/cclegg7/straw-hat-challenge/models"
)

type Calculator struct {
	database *database.Database
}

func NewCalculator(database *database.Database) *Calculator {
	return &Calculator{
		database: database,
	}
}

func (c *Calculator) AllTotalScores() ([]*models.Score, error) {
	users, err := c.database.GetUsers()
	if err != nil {
		return nil, err
	}

	var scores []*models.Score
	for _, user := range users {
		points, err := c.totalScore(user)
		if err != nil {
			return nil, err
		}

		scores = append(scores, &models.Score{
			User:   user,
			Points: points,
		})
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Points > scores[j].Points
	})

	return scores, nil
}

func (c *Calculator) totalScore(user *models.User) (int, error) {
	climbingScore, err := c.climbingScore(user)
	if err != nil {
		return 0, err
	}

	runningScore, err := c.runningScore(user)
	if err != nil {
		return 0, err
	}

	return climbingScore + runningScore, nil
}
