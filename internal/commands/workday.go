package commands

import (
	"fmt"
	"log"
	"time"

	"example.com/workhours/config"
	"example.com/workhours/internal/models"
	"example.com/workhours/utils"
)

var cfg config.Config = config.LoadConfig()

func CmdBeginWorkDay() {
	fmt.Println("LFG!")
	foundWorkDay, err := utils.GetOngoingWorkDay()
	if err != nil {
		log.Fatalln("failed to get ongoing workday:", err)
	}
	// check if uninitilized fix this
	if foundWorkDay.FinnishedAt.Year() == 1 && foundWorkDay.StartedAt.Year() != 1 {
		fmt.Println("There is already a workday ongoing. Started at:", foundWorkDay.StartedAt)
		return
	}
	startedAt := time.Now()
	var workDay models.WorkDay = models.WorkDay{
		Id:            utils.GetFakeUUID(),
		StartedAt:     startedAt,
		TasksWorkedAt: []string{},
	}
	utils.AppendWorkDays(cfg.WorkDaysFilePath, workDay)
}

func CmdBreakWorkDay() {
	fmt.Println("I say hammer on but sure...")
}

// Handle case when day ends on a different date then it begun?
func CmdEndWorkDay() {
	finnishedAt := time.Now()
	fmt.Println("Then go to bed Bitch!")
	foundWorkDay, err := utils.GetOngoingWorkDay()
	if err != nil {
		log.Fatalln("failed to get ongoing workday:", err)
	}
	// check if uninitilized fix this
	if foundWorkDay.FinnishedAt.Year() == 1 && foundWorkDay.StartedAt.Year() == 1 {
		fmt.Println("There is no ongoing workday")
		return
	}
	foundWorkDay.FinnishedAt = finnishedAt
	editErr := utils.EditWorkDay(cfg.WorkDaysFilePath, foundWorkDay)
	if editErr != nil {
		log.Fatalln("failed to edit ongoing workday:", err)
	}
}
