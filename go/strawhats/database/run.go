package database

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

type RunsForWeek struct {
	Weekday     int
	Weekend     int
	MaxDistance int
}

func (d *Database) GetRunsByWeekForUserID(userID int) (map[int]*RunsForWeek, error) {
	rows, err := d.db.Query(selectUserWeeklyRunsQuery, userID)
	if err != nil {
		return nil, fmt.Errorf("error querying for runs for a user: %w", err)
	}
	defer rows.Close()

	runsByWeek := make(map[int]*RunsForWeek)
	for rows.Next() {
		var week, is_weekend, runCount, maxDistance int
		if err := rows.Scan(&week, &is_weekend, &runCount, &maxDistance); err != nil {
			return nil, fmt.Errorf("error reading a row: %w", err)
		}

		if runsByWeek[week] == nil {
			runsByWeek[week] = &RunsForWeek{}
		}

		if is_weekend == 1 {
			runsByWeek[week].Weekend = runCount
		} else {
			runsByWeek[week].Weekday = runCount
		}

		if runsByWeek[week].MaxDistance < maxDistance {
			runsByWeek[week].MaxDistance = maxDistance
		}
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error on user runs result: %w", err)
	}

	return runsByWeek, nil
}

const selectUserWeeklyRunsQuery = `
SELECT week_num, is_weekend, COUNT(1), MAX(distance)
from runs_with_week_info
WHERE user_id = ?
GROUP BY week_num, is_weekend`
