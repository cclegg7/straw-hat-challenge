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
