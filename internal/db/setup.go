package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // SQLite driver

	"github.com/eliasrenger/workhours/internal/paths"
)

func SetupDB() error {
	db_path, err := paths.GetDBPath()
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		return err
	}

	// Create tables if they do not exist
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS workday_session (
		id TEXT PRIMARY KEY,
		workday_id TEXT,
		started_at INTEGER,
		ended_at INTEGER,
		notes TEXT,
		FOREIGN KEY(workday_id) REFERENCES workday(id)
		ON DELETE CASCADE
	);
	CREATE TABLE IF NOT EXISTS workday (
		id TEXT PRIMARY KEY,
		date TEXT NOT NULL,
		is_active INTEGER NOT NULL DEFAULT 1 CHECK (is_active IN (0, 1)),
		number_of_quick_breaks INTEGER,
		last_quick_break INTEGER,
		notes TEXT
	);
	CREATE TABLE IF NOT EXISTS task_session (
		id TEXT PRIMARY KEY,
		task_id TEXT NOT NULL,
		started_at INTEGER,
		ended_at INTEGER,
		notes TEXT,
		FOREIGN KEY(task_id) REFERENCES task(id) ON DELETE CASCADE
	);
	CREATE TABLE IF NOT EXISTS task (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		priority INTEGER,
		category TEXT,
		created_at INTEGER NOT NULL ,
		completed_at INTEGER,
		status TEXT DEFAULT 'paused' CHECK (status IN ('active', 'paused', 'completed', 'cancelled')),
		estimated_duration INTEGER,
		notes TEXT
	);
	CREATE TABLE workday_task_session (
	    workday_id TEXT NOT NULL,
	    task_session_id TEXT NOT NULL,
	    PRIMARY KEY (workday_id, task_session_id),
	    FOREIGN KEY(workday_id) REFERENCES workdays(id) ON DELETE CASCADE,
	    FOREIGN KEY(task_session_id) REFERENCES task_sessions(id) ON DELETE CASCADE
	);
	`

	if _, err = db.Exec(createTableSQL); err != nil {
		return err
	}

	return nil
}
