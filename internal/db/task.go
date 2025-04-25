package db

import (
	"database/sql"

	"github.com/eliasrenger/workhours/internal/models"
	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func InsertTask(t models.Task) error {
	_, err := DB.Exec(`
		INSERT INTO task (id, title, description, priority, category, created_at, completed_at, status, estimated_duration, notes)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		t.Id, t.Title, t.Description, t.Priority, t.Category,
		t.CreatedAt, t.CompletedAt, t.Status, t.EstimatedDuration, t.Notes,
	)
	return err
}

func GetTaskByID(id string) (models.Task, error) {
	var t models.Task
	row := DB.QueryRow(`SELECT id, title, description, priority, category, created_at, completed_at, status, estimated_duration, notes FROM task WHERE id = ?`, id)

	err := row.Scan(
		&t.Id,
		&t.Title,
		&t.Description,
		&t.Priority,
		&t.Category,
		&t.CreatedAt,
		&t.CompletedAt,
		&t.Status,
		&t.EstimatedDuration,
		&t.Notes,
	)
	if err == sql.ErrNoRows {
		return t, nil // No active session found
	}
	return t, err
}

func GetActiveTask() (models.Task, error) {
	var t models.Task
	row := DB.QueryRow(`SELECT id, title, description, priority, category, created_at, completed_at, status, estimated_duration, notes FROM task WHERE status = 'active'`)

	err := row.Scan(&t.Id, &t.Title, &t.Description, &t.Priority, &t.Category, &t.CreatedAt, &t.CompletedAt, &t.Status, &t.EstimatedDuration, &t.Notes)
	if err == sql.ErrNoRows {
		return t, nil // No active session found
	}
	return t, err
}

func UpdateTaskById(t models.Task) error {
	_, err := DB.Exec(
		`UPDATE task 
		SET description = ?, 
		    priority = ?, 
		    category = ?, 
		    created_at = ?,
			completed_at = ?,
			status = ?,
			estimated_duration = ?,
			notes = ?,
		WHERE id = ?`,
		t.Description,
		t.Priority,
		t.Category,
		t.CreatedAt,
		t.CompletedAt,
		t.Status,
		t.EstimatedDuration,
		t.Notes,
		t.Id,
	)

	return err
}
