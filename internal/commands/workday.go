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
		TimeSessions:  []models.TimeSession{},
	}
	current_time_session := models.TimeSession{StartedAt: startedAt}
	workDay.TimeSessions = append(workDay.TimeSessions, current_time_session)
	utils.AppendWorkDays(cfg.WorkDaysFilePath, workDay)
}

// TODO: break and resume should check length of TimeSessions
func CmdBreakWorkDay() {
	finnishedAt := time.Now()
	fmt.Println("I say hammer on but sure...")
	foundWorkDay, err := utils.GetOngoingWorkDay()
	if err != nil {
		log.Fatalln("failed to get ongoing workday:", err)
	}

	lastSessionId := len(foundWorkDay.TimeSessions) - 1
	lastSession := foundWorkDay.TimeSessions[lastSessionId]
	if lastSession.FinnishedAt.Year() != 1 {
		fmt.Println("The workday was paused already.")
		return
	} else {
		foundWorkDay.TimeSessions[lastSessionId].FinnishedAt = finnishedAt
	}
	editErr := utils.EditWorkDay(cfg.WorkDaysFilePath, foundWorkDay)
	if editErr != nil {
		log.Fatalln("failed to edit ongoing workday:", err)
	}
}

func CmdResumeWorkDay() {
	startedAt := time.Now()
	foundWorkDay, err := utils.GetOngoingWorkDay()
	if err != nil {
		log.Fatalln("failed to get ongoing workday:", err)
	}
	// check if last session is finnished
	lastSessionId := len(foundWorkDay.TimeSessions) - 1
	lastSession := foundWorkDay.TimeSessions[lastSessionId]
	if lastSession.FinnishedAt.Year() == 1 {
		fmt.Println("The workday wasn't paused.")
		return
	} else {
		foundWorkDay.TimeSessions = append(
			foundWorkDay.TimeSessions,
			models.TimeSession{
				StartedAt: startedAt,
			},
		)
	}
	editErr := utils.EditWorkDay(cfg.WorkDaysFilePath, foundWorkDay)
	if editErr != nil {
		log.Fatalln("failed to edit ongoing workday:", err)
	}
}

// Handle case when day ends on a different date then it begun?
func CmdEndWorkDay() {
	finnishedAt := time.Now()
	fmt.Println("Then go to bed Bitch!")
	foundWorkDay, err := utils.GetOngoingWorkDay()
	if err != nil {
		log.Fatalln("failed to get ongoing workday:", err)
	}
	// check if we have unfinnished session
	lastSessionId := len(foundWorkDay.TimeSessions) - 1
	lastSession := foundWorkDay.TimeSessions[lastSessionId]
	if lastSession.FinnishedAt.Year() != 1 {
		fmt.Println("The workday was paused already. Saving from pause time stamp.")
		foundWorkDay.FinnishedAt = lastSession.FinnishedAt
	} else {
		foundWorkDay.FinnishedAt = finnishedAt
		foundWorkDay.TimeSessions[lastSessionId].FinnishedAt = finnishedAt
	}
	editErr := utils.EditWorkDay(cfg.WorkDaysFilePath, foundWorkDay)
	if editErr != nil {
		log.Fatalln("failed to edit ongoing workday:", err)
	}
}
