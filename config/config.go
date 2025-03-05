package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	TasksFilePath    string `json:"tasks_file_path"`
	WorkDaysFilePath string `json:"work_days_file_path"`
}

func LoadConfig() Config {
	// Read as a specific struct
	cwd, _ := os.Getwd()
	log.Printf("Current working directory: %s", cwd)

	// Read the file
	var config Config
	data, err := os.ReadFile("./config/config.json")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return config
	}

	// Unmarshal the JSON
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return config
	}

	return config
}
