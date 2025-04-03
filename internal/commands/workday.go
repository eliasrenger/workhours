package commands

import (
	"fmt"
	"log"
	"time"

	"github.com/eliasrenger/workhours/internal/models"
	textformatting "github.com/eliasrenger/workhours/internal/text_formatting"
	"github.com/eliasrenger/workhours/utils"
	task_utils "github.com/eliasrenger/workhours/utils/task"
	work_day_utils "github.com/eliasrenger/workhours/utils/work_day"
)

func CmdStartWorkDay(currentTime time.Time) {
	foundWorkDay, ok := work_day_utils.GetActiveWorkDay()

	if ok && work_day_utils.IsWorkDayActive(foundWorkDay) {
		fmt.Println("there is already a workday ongoing started at:", foundWorkDay.StartedAt)
		return
	}

	var workDay models.WorkDay = models.WorkDay{
		Id:            utils.GetFakeUUID(),
		StartedAt:     currentTime,
		TasksWorkedOn: []string{},
		TimeSessions:  []models.TimeSession{},
	}
	current_time_session := models.TimeSession{StartedAt: currentTime}
	workDay.TimeSessions = append(workDay.TimeSessions, current_time_session)
	work_day_utils.AppendWorkDay(workDay)
}

func CmdPauseWorkDay(currentTime time.Time) {
	foundWorkDay, ok := work_day_utils.GetActiveWorkDay()
	if !ok {
		log.Fatalln("no workday is logged - use command 'start' to log your first workday")
	}

	if !work_day_utils.IsLastSessionActive(foundWorkDay) {
		fmt.Println("the workday is already paused.")
		return
	}

	foundWorkDay.TimeSessions[len(foundWorkDay.TimeSessions)-1].FinishedAt = currentTime
	work_day_utils.EditWorkDay(foundWorkDay)

	activeTask, ok := task_utils.GetActiveTask()
	if ok {
		task_utils.PauseActiveTask(activeTask.Name, currentTime)
	}
}

func CmdResumeWorkDay(currentTime time.Time) {
	foundWorkDay, ok := work_day_utils.GetActiveWorkDay()
	if !ok {
		log.Fatalln("no active workday - use command 'start' to start your workday")
	}

	if work_day_utils.IsLastSessionActive(foundWorkDay) {
		fmt.Println("The workday isn't paused.")
		return
	}

	work_day_utils.ResumeWorkDay(foundWorkDay, currentTime)
}

func CmdQuickieWorkDay(currentTime time.Time) {
	foundWorkDay, ok := work_day_utils.GetActiveWorkDay()
	if !ok {
		log.Fatalln("no workday is logged - use command 'start' to log your first workday")
	}

	if foundWorkDay.LastQuickBreak.Year() == 1 {
		fmt.Println("first quickie of the day. good job!")
	} else { // TODO round output
		fmt.Printf("last quickie was %v min ago\n", currentTime.Sub(foundWorkDay.LastQuickBreak).Minutes())
	}

	foundWorkDay.NumberOfQuickBreaks++
	foundWorkDay.LastQuickBreak = currentTime
	work_day_utils.EditWorkDay(foundWorkDay)

	fmt.Println("quick break registered")
}

// Handle case when day finishes on a different date than it begun?
func CmdFinishWorkDay(currentTime time.Time) {
	foundWorkDay, ok := work_day_utils.GetActiveWorkDay()
	if !ok {
		log.Fatalln("no workday is active - use command 'start' to log your first workday")
	}

	if work_day_utils.IsLastSessionActive(foundWorkDay) {
		foundWorkDay.FinishedAt = currentTime
		foundWorkDay.TimeSessions[len(foundWorkDay.TimeSessions)-1].FinishedAt = currentTime
	} else {
		foundWorkDay.FinishedAt = currentTime
	}

	updatedWorkDay := work_day_utils.UpdateWorkDay(foundWorkDay)
	work_day_utils.EditWorkDay(updatedWorkDay)

	fmt.Println(textformatting.EndOfWorkDayFormat(updatedWorkDay))
}

func CmdHoursWorkDay() {
	foundWorkDay, ok := work_day_utils.GetActiveWorkDay()
	if !ok {
		log.Fatalln("no workday is active - use command 'start' to log your first workday")
	}

	// calculate hours
	updatedWorkDay := work_day_utils.UpdateWorkDay(foundWorkDay)
	work_day_utils.EditWorkDay(updatedWorkDay)

	fmt.Println(textformatting.EndOfWorkDayFormat(updatedWorkDay))
}
