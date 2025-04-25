package commands

import (
	"time"
)

func CmdStartWork(currentTime time.Time) {
	// foundWorkSession, err := db.GetActiveWorkSession()
	// if err != nil {
	// 	log.Fatalln("failed to get active work session:", err)
	// }

	// if foundWorkSession == nil {
	// 	fmt.Println("there is already a workday ongoing started at:", foundWorkSession.StartedAt)
	// 	return
	// }

	// err := services.StartWork(currentTime)
}

func CmdStopWork(currentTime time.Time) {
	// foundWorkDay, ok := work_day_utils.GetActiveWorkDay()
	// if !ok {
	// 	log.Fatalln("no workday is logged - use command 'start' to log your first workday")
	// }

	// if !work_day_utils.IsLastSessionActive(foundWorkDay) {
	// 	fmt.Println("the workday is already paused.")
	// 	return
	// }

	// foundWorkDay.TimeSessions[len(foundWorkDay.TimeSessions)-1].FinishedAt = currentTime
	// work_day_utils.EditWorkDay(foundWorkDay)

	// activeTask, ok := task_utils.GetActiveTask()
	// if ok {
	// 	task_utils.PauseActiveTask(activeTask.Name, currentTime)
	// }
}

func CmdQuickieWork(currentTime time.Time) {
	// foundWorkDay, ok := work_day_utils.GetActiveWorkDay()
	// if !ok {
	// 	log.Fatalln("no workday is logged - use command 'start' to log your first workday")
	// }

	// if foundWorkDay.LastQuickBreak.Year() == 1 {
	// 	fmt.Println("first quickie of the day. good job!")
	// } else { // TODO round output
	// 	fmt.Printf("last quickie was %v min ago\n", currentTime.Sub(foundWorkDay.LastQuickBreak).Minutes())
	// }

	// foundWorkDay.NumberOfQuickBreaks++
	// foundWorkDay.LastQuickBreak = currentTime
	// work_day_utils.EditWorkDay(foundWorkDay)

	// fmt.Println("quick break registered")
}

func CmdHoursWork() {
	// foundWorkDay, ok := work_day_utils.GetActiveWorkDay()
	// if !ok {
	// 	log.Fatalln("no workday is active - use command 'start' to log your first workday")
	// }

	// // calculate hours
	// updatedWorkDay := work_day_utils.UpdateWorkDay(foundWorkDay)
	// work_day_utils.EditWorkDay(updatedWorkDay)

	// fmt.Println(textformatting.EndOfWorkDayFormat(updatedWorkDay))
}
