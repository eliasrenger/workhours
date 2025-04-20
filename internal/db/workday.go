package db

import (

	"github.com/eliasrenger/workhours/internal/models"
	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func InsertWorkday(w models.Workday) error {
	IsActiveInt := convertBoolToInt(w.IsActive)
	_, err := DB.Exec(
		`INSERT INTO workday (id, date, is_active, number_of_quick_breaks, last_quick_break, notes)
	VALUES (?, ?, ?, ?, ?, ?)`,
		w.Id,
		w.Date,
		IsActiveInt,
		w.NumberOfQuickBreaks,
		w.LastQuickBreak,
		w.Notes,
	)
	return err
}

func GetWorkdayByID(id string) (models.Workday, error) {
	var w models.Workday
	var isActiveInt int
	row := DB.QueryRow(`SELECT id, date, is_active, number_of_quick_breaks, last_quick_break, notes FROM workday WHERE id = ?`, id)

	err := row.Scan(&w.Id, &w.Date, &isActiveInt, &w.NumberOfQuickBreaks, &w.LastQuickBreak, &w.Notes)
	w.IsActive = convertIntToBool(isActiveInt)
	return w, err
}

func GetActiveWorkday() (models.Workday, error) {
	var w models.Workday
	var isActiveInt int
	row := DB.QueryRow(`SELECT id, date, is_active, number_of_quick_breaks, last_quick_break, notes FROM workday WHERE is_active = 1`)

	err := row.Scan(&w.Id, &w.Date, &isActiveInt, &w.NumberOfQuickBreaks, &w.LastQuickBreak, &w.Notes)
	w.IsActive = convertIntToBool(isActiveInt)
	return w, err
}

func UpdateWorkdayById(w models.Workday) error {
	isActiveInt := convertBoolToInt(w.IsActive)

	_, err := DB.Exec(
		`UPDATE workday 
		SET is_active = ?, 
		    number_of_quick_breaks = ?, 
		    last_quick_break = ?, 
		    notes = ?
		WHERE id = ?`,
		isActiveInt,
		w.NumberOfQuickBreaks,
		w.LastQuickBreak,
		w.Notes,
		w.Id,
	)

	return err
}
