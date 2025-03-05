package config

import (
	"log"
	"os"

	"example.com/workhours/utils"
)

type Config struct {
	TasksFilePath    string `json:"tasks_file_path"`
	WorkDaysFilePath string `json:"work_days_file_path"`
}

func LoadConfig() Config {
	// Read as a specific struct
	cwd, _ := os.Getwd()
	log.Printf("Current working directory: %s", cwd)
	config, err := utils.ReadJSON[Config]("./config/config.json")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	return config
}
