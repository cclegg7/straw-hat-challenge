package db

import (
	"context"
	"fmt"

	"github.com/cclegg7/straw-hat-challenge/models"
)

const insertRunStatement = "INSERT INTO `runs` (`user_id`, `distance`, `date`) VALUES (?, ?, ?)"

func (d *Database) CreateRun(run *models.Run) error {
	_, err := d.db.ExecContext(context.Background(), insertRunStatement, run.UserID, run.Distance, formatDate(run.Date))
	if err != nil {
		return fmt.Errorf("error inserting run: %w", err)
	}
	return nil
}

const selectUserRunsQuery = "SELECT date, distance, created_at FROM runs WHERE user_id = ? ORDER BY date DESC"

func (d *Database) ListUserRuns(userID int) ([]*models.Run, error) {
	rows, err := d.db.Query(selectUserRunsQuery, userID)
	if err != nil {
		return nil, fmt.Errorf("error querying for runs for a user: %w", err)
	}
	defer rows.Close()

	var runs []*models.Run
	for rows.Next() {
		run := &models.Run{}
		if err := rows.Scan(&run.Date, &run.Distance, &run.CreatedAt); err != nil {
			return nil, fmt.Errorf("error reading a run: %w", err)
		}
		runs = append(runs, run)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error on user runs result: %w", err)
	}
	return runs, nil
}
