package commands

import (
	"log"
	"time"

	"github.com/eliasrenger/workhours/internal/services"
)

func CmdStartWork(currentTime time.Time) {
	err := services.StartWork(currentTime)
	if err == services.ErrWorkSessionAlreadyActive {
		log.Fatalf("Work session already active.")
	}
	if err != nil {
		log.Fatalf("Error starting work session: %v", err)
	} else {
		log.Println("Work session started.")
	}
}

func CmdStopWork(currentTime time.Time) {
	err := services.StopWork(currentTime)
	if err == services.ErrNoWorkSessionActive {
		log.Fatalf("No active work session.")
	}
	if err != nil {
		log.Fatalf("Error stopping work session: %v", err)
	} else {
		log.Println("Work session stopped.")
	}
}

func CmdQuickieWork(currentTime time.Time) {
	err := services.AddQuickBreak(currentTime)
	if err == services.ErrNoWorkSessionActive {
		log.Fatalf("No active work session.")
	}
	if err != nil {
		log.Fatalf("Error saving quick break: %v", err)
	} else {
		log.Println("Quick break saved.")
	}
}

func CmdHoursWork(currentTime time.Time) {
	secondsWorked, err := services.GetSecondsWorkedToday()
	if err != nil {
		log.Fatalf("Error getting time worked today: %v", err)
	} else {
		wholeHours := secondsWorked % 3600
		wholeMinutes := (secondsWorked - wholeHours*3600) % 60
		log.Printf("You've worked %d hours and %d minutes today.\n", wholeHours, wholeMinutes)
	}
}
