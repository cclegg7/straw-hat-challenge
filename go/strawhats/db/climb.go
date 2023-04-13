package db

import (
	"context"
	"fmt"
	"time"

	"github.com/cclegg7/straw-hat-challenge/models"
)

const insertClimbStatement = "INSERT INTO `climbs` (`user_id`, `category`, `rating`, `is_challenge`, `date`) VALUES (?, ?, ?, ?, ?)"

func (d *Database) CreateClimb(climb *models.Climb) error {
	_, err := d.db.ExecContext(context.Background(), insertClimbStatement, climb.UserID, climb.Category, climb.Rating, climb.IsChallenge, formatDate(climb.Date))
	if err != nil {
		return fmt.Errorf("error inserting climb: %w", err)
	}
	return nil
}

type AggregateClimbs struct {
	Boulders map[models.BoulderDifficulty]int
	TopRopes map[models.TopRopeDifficulty]int
}

const topClimbsQuery = `SELECT user_id, category, rating, date 
FROM climbs 
WHERE 
    user_id = ? 
	AND date >= ? 
    AND date <= ?
	AND category = ?
ORDER BY rating DESC
LIMIT 5`

func (d *Database) GetTopClimbsInCategoryForUserID(userID int, start, end time.Time, category models.ClimbCategory) ([]*models.Climb, error) {
	rows, err := d.db.Query(topClimbsQuery, userID, formatDate(start), formatDate(end), category)
	if err != nil {
		return nil, fmt.Errorf("error querying for top climbs: %w", err)
	}
	defer rows.Close()

	var climbs []*models.Climb
	for rows.Next() {
		climb := &models.Climb{}
		if err := rows.Scan(&climb.UserID, &climb.Category, &climb.Rating, &climb.Date); err != nil {
			return nil, fmt.Errorf("error reading a user: %w", err)
		}
		climbs = append(climbs, climb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error on top climbs result: %w", err)
	}
	return climbs, nil
}

func (d *Database) GetTopClimbsForUsers(users []*models.User, start, end time.Time) (map[*models.User]*AggregateClimbs, error) {
	results := make(map[*models.User]*AggregateClimbs)
	for _, user := range users {
		aggregateClimbs := &AggregateClimbs{
			Boulders: make(map[models.BoulderDifficulty]int),
			TopRopes: make(map[models.TopRopeDifficulty]int),
		}

		topBoulders, err := d.GetTopClimbsInCategoryForUserID(user.ID, start, end, models.ClimbCategory_Boulder)
		if err != nil {
			return nil, fmt.Errorf("error fetching boulders for user %d: %w", user.ID, err)
		}

		for _, boulder := range topBoulders {
			aggregateClimbs.Boulders[models.BoulderDifficulty(boulder.Rating)]++
		}

		topTopRopes, err := d.GetTopClimbsInCategoryForUserID(user.ID, start, end, models.ClimbCategory_TopRope)
		if err != nil {
			return nil, fmt.Errorf("error fetching top ropes for user %d: %w", user.ID, err)
		}

		for _, topRope := range topTopRopes {
			aggregateClimbs.TopRopes[models.TopRopeDifficulty(topRope.Rating)]++
		}

		results[user] = aggregateClimbs
	}

	return results, nil
}
