package paths

import (
	"os"
	"path/filepath"
)

const AppName = "workhours"

// GetDataDir returns the OS-appropriate directory to store app data
func GetDataDir() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	dataDir := filepath.Join(base, AppName)

	// Create the folder if it doesn't exist
	err = os.MkdirAll(dataDir, 0755)
	if err != nil {
		return "", err
	}

	return dataDir, nil
}

// GetDBPath returns the full path to the SQLite DB file
func GetDBPath() (string, error) {
	dataDir, err := GetDataDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dataDir, "workhours.db"), nil
}
