package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"

	"example.com/workhours/internal/models"
)

func ReadWorkDays(filePath string) ([]models.WorkDay, error) {
	var result []models.WorkDay

	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return result, err
	}

	// Unmarshal the JSON
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func AppendWorkDays(filePath string, newWorkDay models.WorkDay) error {
	oldData, err := ReadWorkDays(filePath)
	if err != nil {
		return err
	}
	rawData := append(oldData, newWorkDay)
	return SaveWorkDays(filePath, rawData)
}

func SaveWorkDays(filePath string, rawWorkDays []models.WorkDay) error {
	data, err := json.MarshalIndent(rawWorkDays, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

func EditWorkDay(filePath string, newWorkDay models.WorkDay) error {
	workDays, err := ReadWorkDays(filePath)
	if err != nil {
		return err
	}

	var cleanedWorkDays []models.WorkDay
	var foundTargetTask bool
	for idx, workDay := range workDays {
		fmt.Println(idx, workDay, newWorkDay)
		if workDay.Id == newWorkDay.Id {
			cleanedWorkDays = slices.Delete(workDays, idx, idx+1)
			fmt.Println("found workDay", cleanedWorkDays)
			foundTargetTask = true
			break
		}
	}
	if !foundTargetTask {
		var Error error
		return Error
	}
	editedWorkDays := append(cleanedWorkDays, newWorkDay)
	return SaveWorkDays(filePath, editedWorkDays)
}
