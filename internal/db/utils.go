package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

var DB *sql.DB

func convertBoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func convertIntToBool(i int) bool {
	return i == 1
}
