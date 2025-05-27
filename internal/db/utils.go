package db

import (
	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func convertBoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func convertIntToBool(i int) bool {
	return i == 1
}
