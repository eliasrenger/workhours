package commands

import (
	"time"
)

func CmdStartTask(currentTime time.Time, args []string) {
	// find target task
	// targetTask, ok := task_utils.GetTaskByName(args[0])
	// targetTaskStatus := "finished"
	// if !task_utils.IsTaskFinished(targetTask) {
	// 	targetTaskStatus = "ongoing"
	// }
	// if ok {
	// 	fmt.Printf("there is already a task with the name %v and status %v", targetTask.Name, targetTaskStatus)
	// 	return
	// }

	// // check for active task
	// activeTask, ok := task_utils.GetActiveTask()

	// if ok { // if there is an active task, finish last time session
	// 	activeTask.TimeSessions[len(activeTask.TimeSessions)-1].FinishedAt = currentTime
	// 	task_utils.EditTask(activeTask)
	// 	hour, minute, second := currentTime.Clock()
	// 	fmt.Printf("task %v was paused at %v:%v:%v", activeTask.Name, hour, minute, second)
	// }

	// // create new task
	// var task models.Task = models.Task{
	// 	Id:           uuid.New(),
	// 	Name:         args[0],
	// 	StartedAt:    currentTime,
	// 	TimeSessions: []models.TimeSession{},
	// 	WorkCount:    1,
	// }
	// current_time_session := models.TimeSession{StartedAt: currentTime}
	// task.TimeSessions = append(task.TimeSessions, current_time_session)
	// task_utils.AppendTask(task)

	// // begin or resume workday
	// activeWorkDay, ok := work_day_utils.GetActiveWorkDay()
	// if !ok {
	// 	work_day_utils.ResumeWorkDay(activeWorkDay, currentTime)
	// } else {
	// 	CmdStartWork(currentTime)
	// 	activeWorkDay, _ = work_day_utils.GetActiveWorkDay() // line above creates it
	// }

	// // save task to work day
	// foundTask := false
	// for _, taskName := range activeWorkDay.TasksWorkedOn {
	// 	if taskName == task.Name {
	// 		foundTask = true
	// 	}
	// }
	// if !foundTask {
	// 	activeWorkDay.TasksWorkedOn = append(activeWorkDay.TasksWorkedOn, task.Name)
	// 	work_day_utils.EditWorkDay(activeWorkDay)
	// }
}

func CmdStopTask(currentTime time.Time, args []string) {
	// check for active task
	// targetTask, ok := task_utils.GetTaskByName(args[0])
	// if !ok {
	// 	fmt.Print("there is no task with the name:", targetTask.Name)
	// 	return
	// }

	// // find active task
	// activeTask, ok := task_utils.GetActiveTask()
	// if !ok { // if there is an active task
	// 	log.Fatalln("there is no active task.")
	// }
	// if activeTask.Name != targetTask.Name {
	// 	log.Fatalln("there is already an active task with name:", activeTask.Name)
	// }

	// // assume activeTask == targetTask
	// // finish last timesession
	// activeTask.TimeSessions[len(activeTask.TimeSessions)-1].FinishedAt = currentTime
	// updatedTask := task_utils.UpdateTask(activeTask)
	// task_utils.EditTask(updatedTask)
	// fmt.Printf("task %v was paused at %v", updatedTask.Name, currentTime)
}

// func CmdResumeTask(currentTime time.Time, args []string) {
// 	// check for active tasks
// 	activeTask, ok := task_utils.GetActiveTask()
// 	if ok { // if there is an active task
// 		CmdPauseTask(currentTime, []string{activeTask.Name})
// 		log.Println("paused", activeTask.Name)
// 	}

// 	// find target task
// 	targetTask, ok := task_utils.GetTaskByName(args[0])

// 	if !ok {
// 		log.Fatalln("there is no task with the name:", targetTask.Name)
// 	}
// 	if task_utils.IsTaskFinished(targetTask) {
// 		log.Fatalf("the task %v is finished, can't resume task\n", targetTask.Name)
// 	}

// 	// start new timesession
// 	current_time_session := models.TimeSession{StartedAt: currentTime}
// 	targetTask.TimeSessions = append(targetTask.TimeSessions, current_time_session)
// 	task_utils.EditTask(targetTask)
// 	fmt.Printf("clock for task %v was resumed at %v", activeTask.Name, currentTime)

// 	// start or resume workday
// 	activeWorkDay, ok := work_day_utils.GetActiveWorkDay()
// 	if ok {
// 		CmdResumeWorkDay(currentTime)
// 	} else {
// 		CmdStartWork(currentTime)
// 	}

// 	// save task to work day
// 	foundTask := false
// 	for _, taskName := range activeWorkDay.TasksWorkedOn {
// 		if taskName == targetTask.Name {
// 			foundTask = true
// 		}
// 	}
// 	if !foundTask {
// 		activeWorkDay.TasksWorkedOn = append(activeWorkDay.TasksWorkedOn, targetTask.Name)
// 		work_day_utils.EditWorkDay(activeWorkDay)
// 	}
// }

func CmdFinishTask(currentTime time.Time, args []string) {
	// find target task
	// targetTask, ok := task_utils.GetTaskByName(args[0])
	// if !ok {
	// 	log.Fatalln("there is no task with the name:", targetTask.Name)
	// }
	// if task_utils.IsTaskFinished(targetTask) {
	// 	log.Fatalln("task is already finished")
	// }

	// // finish last timesession and finish task
	// targetTask.TimeSessions[len(targetTask.TimeSessions)-1].FinishedAt = currentTime
	// targetTask.FinishedAt = currentTime
	// updatedTask := task_utils.UpdateTask(targetTask)
	// task_utils.EditTask(updatedTask)
	// fmt.Printf("task %v finished, took %v", updatedTask.Name, updatedTask.Duration)
}

func CmdListTasks() {
	// tasks := task_utils.ReadTasks()

	// // extract ongoing "unfinished" tasks
	// var ongoingTasks []models.Task
	// for _, task := range tasks {
	// 	if task_utils.IsTaskFinished(task) {
	// 		continue
	// 	}
	// 	updatedTask := task_utils.UpdateTask(task)
	// 	ongoingTasks = append(ongoingTasks, updatedTask)
	// }
	// if len(ongoingTasks) == 0 {
	// 	fmt.Println("No ongoing tasks")
	// 	return
	// }

	// fmt.Println("Currently ongoing tasks")
	// fmt.Printf("| id | task name | task duration [h] |\n")
	// for id, ongoingTasks := range ongoingTasks {
	// 	fmt.Printf("| %v | %v |   %v   |\n", id, ongoingTasks.Name, ongoingTasks.Duration.Hours())
	// }
}
