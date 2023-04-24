package db

import "fmt"

const (
	insertFileStatement      = "INSERT INTO `files` (`token`, `url`, `content_type`) VALUES (?, ?, ?)"
	linkFileToRunStatement   = "UPDATE `files` SET `type` = ?, `run_id` = ? WHERE `token` = ?"
	linkFileToClimbStatement = "UPDATE `files` SET `type` = ?, `climb_id` = ? WHERE `token` = ?"
)

type fileReferenceType uint8

const (
	fileReferenceType_unknown fileReferenceType = 0
	fileReferenceType_run     fileReferenceType = 1
	fileReferenceType_climb   fileReferenceType = 2
)

func (d *Database) InsertFile(token, url, contentType string) error {
	if _, err := d.db.Exec(insertFileStatement, token, url, contentType); err != nil {
		return fmt.Errorf("error inserting file for run: %w", err)
	}

	return nil
}

func (d *Database) LinkFileToRun(token string, runID int) error {
	if _, err := d.db.Exec(linkFileToRunStatement, fileReferenceType_run, runID, token); err != nil {
		return fmt.Errorf("error linking file to run: %w", err)
	}

	return nil
}

func (d *Database) LinkFileToClimb(token string, climbID int) error {
	if _, err := d.db.Exec(linkFileToClimbStatement, fileReferenceType_climb, climbID, token); err != nil {
		return fmt.Errorf("error linking file to run: %w", err)
	}

	return nil
}
