package main

import (
	"log"
	"os"
	"time"

	"github.com/eliasrenger/workhours/internal/commands"
)

func main() {
	currentTime := time.Now()
	if len(os.Args) < 2 {
		log.Fatalln("Provide command. Use workhours help for more information.")
	}

	var command string = os.Args[1]
	var subCommand string
	if len(os.Args) > 2 {
		subCommand = os.Args[2]
	}
	var args []string
	if len(os.Args) > 3 {
		args = os.Args[3:]
	}

	switch command {
	case "help":
		commands.CmdHelp()

	case "setup":
		commands.CmdDBSetup()

	// type specific commands
	case "quickie":
		commands.CmdQuickieWork(currentTime)
	case "list":
		commands.CmdListTasks()
	case "hours":
		commands.CmdHoursWork()

	// start tracking
	case "start":
		switch subCommand {
		case "work":
			commands.CmdStartWork(currentTime)
		case "task":
			commands.CmdStartTask(currentTime, args)
		default:
			commands.CmdHelp()
		}

	// pause tracking
	case "stop":
		switch subCommand {
		case "work":
			commands.CmdStopWork(currentTime)
		case "task":
			if len(args) < 1 {
				log.Fatalln("provide task name")
			}
			commands.CmdStopTask(currentTime, args)
		default:
			commands.CmdHelp()
		}

	// finish tracking
	case "finish":
		switch subCommand {
		case "task":
			if len(args) < 1 {
				log.Fatalln("provide task name")
			}
			commands.CmdFinishTask(currentTime, args)
		default:
			commands.CmdHelp()
		}

	// no match
	default:
		commands.CmdHelp()
	}
}
