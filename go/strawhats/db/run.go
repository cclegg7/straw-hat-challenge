package db

import (
	"context"
	"fmt"

	"github.com/cclegg7/straw-hat-challenge/models"
)

const insertRunStatement = "INSERT INTO `runs` (`user_id`, `distance`, `date`) VALUES (?, ?, ?)"

func (d *Database) CreateRun(run *models.Run) (int, error) {
	result, err := d.db.ExecContext(context.Background(), insertRunStatement, run.UserID, run.Distance, formatDate(run.Date))
	if err != nil {
		return 0, fmt.Errorf("error inserting run: %w", err)
	}

	runId, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting run id: %w", err)
	}

	return int(runId), nil
}

const selectUserRunsWithFilesQuery = "SELECT date, distance, created_at, f.token, f.url, f.content_type FROM runs r LEFT OUTER JOIN files f ON (r.id = f.run_id) WHERE user_id = ? ORDER BY date DESC"

type selectUserRunsWithFilesRow struct {
	run             *models.Run
	fileToken       *string
	fileURL         *string
	fileContentType *string
}

func (d *Database) ListUserRunsWithFiles(userID int) ([]*models.Run, error) {
	rows, err := d.db.Query(selectUserRunsWithFilesQuery, userID)
	if err != nil {
		return nil, fmt.Errorf("error querying for runs for a user: %w", err)
	}
	defer rows.Close()

	var runs []*models.Run
	for rows.Next() {
		row := &selectUserRunsWithFilesRow{
			run: &models.Run{},
		}
		if err := rows.Scan(&row.run.Date, &row.run.Distance, &row.run.CreatedAt, &row.fileToken, &row.fileURL, &row.fileContentType); err != nil {
			return nil, fmt.Errorf("error reading a run: %w", err)
		}
		if row.fileToken != nil && len(*row.fileToken) > 0 {
			row.run.Files = append(row.run.Files, &models.File{
				Token:       *row.fileToken,
				URL:         *row.fileURL,
				ContentType: *row.fileContentType,
			})
		}
		runs = append(runs, row.run)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error on user runs result: %w", err)
	}
	return runs, nil
}
