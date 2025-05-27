package db

import (
	"database/sql"
	"errors"
	"time"

	"github.com/eliasrenger/workhours/internal/models"
	"github.com/eliasrenger/workhours/utils"
	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

var (
	ErrDateFormat = errors.New("date format is incorrect expected YYYY-MM-DD")
)

func InsertWorkSession(w models.WorkSession) error {
	_, err := DB.Exec(
		`INSERT INTO work_session (id, date, started_at, ended_at, number_of_quick_breaks, last_quick_break, notes)
	VALUES (?, ?, ?, ?, ?, ?, ?)`,
		w.Id,
		w.Date,
		w.StartedAt,
		w.EndedAt,
		w.NumberOfQuickBreaks,
		w.LastQuickBreak,
		w.Notes,
	)

	return err
}

func GetWorkSessionByID(id string) (models.WorkSession, error) {
	var w models.WorkSession
	row := DB.QueryRow(`SELECT * FROM workday_session WHERE id = ?`, id)

	err := row.Scan(
		&w.Id,
		&w.Date,
		&w.StartedAt,
		&w.EndedAt,
		&w.NumberOfQuickBreaks,
		&w.LastQuickBreak,
		&w.Notes,
	)
	if err == sql.ErrNoRows {
		return w, nil // No active session found
	}
	return w, err
}

func GetActiveWorkSession() (models.WorkSession, error) {
	var w models.WorkSession
	err := InitDB()
	if err != nil {
		return w, err
	}
	row := DB.QueryRow(`SELECT * FROM workday_session WHERE ended_at = 0`)
	err = row.Scan(
		&w.Id,
		&w.Date,
		&w.StartedAt,
		&w.EndedAt,
		&w.NumberOfQuickBreaks,
		&w.LastQuickBreak,
		&w.Notes,
	)
	if err == sql.ErrNoRows {
		return w, nil // No active session found
	}
	return w, err
}

func HasActiveWorkSession() (bool, error) {
	var count int
	err := InitDB()
	if err != nil {
		return false, err
	}
	row := DB.QueryRow(`SELECT COUNT(*) FROM work_session WHERE ended_at = 0`)
	err = row.Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func UpdateWorkSessionById(w models.WorkSession) error {
	_, err := DB.Exec(
		`UPDATE work_session 
		SET started_at = ?, 
		    ended_at = ?,
			number_of_quick_breaks = ?,
			last_quick_break = ?,
		    notes = ?
		WHERE id = ?`,
		w.StartedAt,
		w.EndedAt,
		w.NumberOfQuickBreaks,
		w.LastQuickBreak,
		w.Notes,
		w.Id,
	)

	return err
}

// returns seconds worked on the given date
// if no workday is found, it returns 0 and nil error
// if an error occurs, it returns 0 and the error
func GetWorktimeByDate(date string) (int64, error) {
	isDateCorrectFormat := utils.IsDateCorrectFormat(date)
	if !isDateCorrectFormat {
		return 0, ErrDateFormat
	}
	var secondsWorked int64
	rows, err := DB.Query(`SELECT id, date, started_at, ended_at, number_of_quick_breaks, last_quick_break, notes FROM work_session WHERE date = ?`, date)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	for rows.Next() {
		var w models.WorkSession
		err := rows.Scan(
			&w.Id,
			&w.Date,
			&w.StartedAt,
			&w.EndedAt,
			&w.NumberOfQuickBreaks,
			&w.LastQuickBreak,
			&w.Notes,
		)
		if err != nil {
			return 0, err
		}
		if w.EndedAt == 0 {
			w.EndedAt = time.Now().Unix()
		}
		secondsWorked += w.EndedAt - w.StartedAt
	}

	return secondsWorked, nil
}
