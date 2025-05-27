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
	CREATE TABLE IF NOT EXISTS work_session (
		id TEXT PRIMARY KEY,
		date TEXT NOT NULL,
		started_at INTEGER,
		ended_at INTEGER,
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
		created_at INTEGER NOT NULL,
		completed_at INTEGER,
		status TEXT DEFAULT 'paused' CHECK (status IN ('active', 'paused', 'completed', 'cancelled')),
		estimated_duration INTEGER,
		notes TEXT
	);
	CREATE TABLE IF NOT EXISTS work_task_session (
	    work_session_id TEXT NOT NULL,
	    task_session_id TEXT NOT NULL,
	    PRIMARY KEY (work_session_id, task_session_id),
	    FOREIGN KEY(work_session_id) REFERENCES work_session(id) ON DELETE CASCADE,
	    FOREIGN KEY(task_session_id) REFERENCES task_session(id) ON DELETE CASCADE
	);
	CREATE TRIGGER IF NOT EXISTS prevent_multiple_active_work_sessions
	BEFORE INSERT ON work_session
	WHEN NEW.ended_at IS NULL
	BEGIN
	    SELECT
	        RAISE(ABORT, 'Only one active work_session is allowed')
	    WHERE EXISTS (
	        SELECT 1 FROM work_session WHERE ended_at IS NULL
	    );
	END;
	CREATE TRIGGER IF NOT EXISTS prevent_update_to_multiple_active_work_sessions
	BEFORE UPDATE ON work_session
	WHEN NEW.ended_at IS NULL
	BEGIN
	    SELECT
	        RAISE(ABORT, 'Only one active work_session is allowed')
	    WHERE EXISTS (
	        SELECT 1 FROM work_session WHERE ended_at IS NULL AND id != NEW.id
	    );
	END;
	`

	if _, err = db.Exec(createTableSQL); err != nil {
		return err
	}

	return nil
}
