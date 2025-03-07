package commands

import (
	"fmt"
	"log"
	"time"

	"example.com/workhours/config"
	"example.com/workhours/internal/models"
	textformatting "example.com/workhours/internal/text_formatting"
	"example.com/workhours/utils"
	work_day_utils "example.com/workhours/utils/work_day"
)

var cfg config.Config = config.LoadConfig()

func CmdBeginWorkDay() {
	startedAt := time.Now()
	foundWorkDay, ok, err := work_day_utils.GetActiveWorkDay()
	if err != nil {
		log.Fatalln("failed to get ongoing workday:", err)
	}

	if ok && work_day_utils.IsWorkDayActive(foundWorkDay) {
		fmt.Println("there is already a workday ongoing started at:", foundWorkDay.StartedAt)
		return
	}

	var workDay models.WorkDay = models.WorkDay{
		Id:            utils.GetFakeUUID(),
		StartedAt:     startedAt,
		TasksWorkedOn: []string{},
		TimeSessions:  []models.TimeSession{},
	}
	current_time_session := models.TimeSession{StartedAt: startedAt}
	workDay.TimeSessions = append(workDay.TimeSessions, current_time_session)
	work_day_utils.AppendWorkDays(cfg.WorkDaysFilePath, workDay)
}

// TODO: break and resume should check length of TimeSessions
func CmdBreakWorkDay() {
	finnishedAt := time.Now()
	foundWorkDay, ok, err := work_day_utils.GetActiveWorkDay()
	if err != nil {
		log.Fatalln("failed to get ongoing workday:", err)
	}
	if !ok {
		log.Fatalln("no workday is logged - use command 'begin' to log your first workday")
	}

	if !work_day_utils.IsLastSessionActive(foundWorkDay) {
		fmt.Println("the workday is already paused.")
		return
	}

	foundWorkDay.TimeSessions[len(foundWorkDay.TimeSessions)-1].FinnishedAt = finnishedAt
	editErr := work_day_utils.EditWorkDay(cfg.WorkDaysFilePath, foundWorkDay)
	if editErr != nil {
		log.Fatalln("failed to edit ongoing workday:", err)
	}
}

func CmdResumeWorkDay() {
	startedAt := time.Now()
	foundWorkDay, ok, err := work_day_utils.GetActiveWorkDay()
	if err != nil {
		log.Fatalln("failed to get ongoing workday:", err)
	}
	if !ok {
		log.Fatalln("no workday is logged - use command 'begin' to log your first workday")
	}

	if work_day_utils.IsLastSessionActive(foundWorkDay) {
		fmt.Println("The workday isn't paused.")
		return
	}

	foundWorkDay.TimeSessions = append(
		foundWorkDay.TimeSessions,
		models.TimeSession{
			StartedAt: startedAt,
		},
	)

	editErr := work_day_utils.EditWorkDay(cfg.WorkDaysFilePath, foundWorkDay)
	if editErr != nil {
		log.Fatalln("failed to edit ongoing workday:", err)
	}
}

func CmdQuickieWorkDay() {
	currentTime := time.Now()
	foundWorkDay, ok, err := work_day_utils.GetActiveWorkDay()
	if err != nil {
		log.Fatalln("failed to get ongoing workday:", err)
	}
	if !ok {
		log.Fatalln("no workday is logged - use command 'begin' to log your first workday")
	}

	if foundWorkDay.LastQuickBreak.Year() == 1 {
		fmt.Println("first quickie of the day. good job!")
	} else { // TODO round output
		fmt.Printf("last quickie was %v min ago\n", currentTime.Sub(foundWorkDay.LastQuickBreak).Minutes())
	}

	foundWorkDay.NumberOfQuickBreaks++
	foundWorkDay.LastQuickBreak = currentTime
	editErr := work_day_utils.EditWorkDay(cfg.WorkDaysFilePath, foundWorkDay)
	if editErr != nil {
		log.Fatalln("failed to edit ongoing workday:", err)
	}

	fmt.Println("quick break registered")
}

// Handle case when day ends on a different date then it begun?
func CmdEndWorkDay() {
	finnishedAt := time.Now()
	foundWorkDay, ok, err := work_day_utils.GetActiveWorkDay()
	if err != nil {
		log.Fatalln("failed to get ongoing workday:", err)
	}
	if !ok {
		log.Fatalln("no workday is logged - use command 'begin' to log your first workday")
	}

	if work_day_utils.IsLastSessionActive(foundWorkDay) {
		foundWorkDay.FinnishedAt = finnishedAt
		foundWorkDay.TimeSessions[len(foundWorkDay.TimeSessions)-1].FinnishedAt = finnishedAt
	} else {
		foundWorkDay.FinnishedAt = finnishedAt
	}

	editErr := work_day_utils.EditWorkDay(cfg.WorkDaysFilePath, foundWorkDay)
	if editErr != nil {
		log.Fatalln("failed to edit ongoing workday:", err)
	}
	fmt.Println(textformatting.EndOfWorkDayFormat(foundWorkDay))
}
