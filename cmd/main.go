package main

import (
	"log"
	"os"

	"example.com/workhours/internal/commands"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Provide command. Use workhours help for more information.")
	}

	var command string = os.Args[1]

	switch command {
	case "help":
		commands.CmdHelp()

	case "begin":
		commands.CmdBeginWorkDay()
	case "break":
		commands.CmdBreakWorkDay()
	case "resume":
		commands.CmdResumeWorkDay()
	case "end":
		commands.CmdEndWorkDay()

	case "start":
		commands.CmdStartTask()
	case "pause":
		commands.CmdPauseTask()
	case "finnish":
		commands.CmdFinnishTask()
	case "list":
		commands.CmdListTasks()
	}
}
