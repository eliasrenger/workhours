package work_day_utils

import (
	"encoding/json"
	"log"
	"os"
	"slices"

	"github.com/eliasrenger/workhours/internal/models"
)

func ReadWorkDays() []models.WorkDay {
	var result []models.WorkDay

	// Read the file
	data, err := os.ReadFile(cfg.WorkDaysFilePath)
	if err != nil {
		log.Fatalln("failed to read file", cfg.WorkDaysFilePath)
	}

	// Unmarshal the JSON
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Fatalln("failed to unmarshal json data")
	}

	return result
}

func AppendWorkDay(newWorkDay models.WorkDay) {
	oldData := ReadWorkDays()
	rawData := append(oldData, newWorkDay)
	SaveWorkDays(rawData)
}

func SaveWorkDays(rawWorkDays []models.WorkDay) {
	data, err := json.MarshalIndent(rawWorkDays, "", "  ")
	if err != nil {
		log.Fatalln("failed to marshal json data")
	}
	err = os.WriteFile(cfg.WorkDaysFilePath, data, 0644)
	if err != nil {
		log.Fatalln("failed to write data to path", cfg.WorkDaysFilePath)
	}
}

func EditWorkDay(newWorkDay models.WorkDay) {
	workDays := ReadWorkDays()

	var cleanedWorkDays []models.WorkDay
	var foundTargetWorkDay bool
	for idx, workDay := range workDays {
		if workDay.Id == newWorkDay.Id {
			cleanedWorkDays = slices.Delete(workDays, idx, idx+1)
			foundTargetWorkDay = true
			break
		}
	}
	if !foundTargetWorkDay {
		log.Fatalln("workday to edit wasn't found in saved data")
	}
	editedWorkDays := append(cleanedWorkDays, newWorkDay)
	SaveWorkDays(editedWorkDays)
}
