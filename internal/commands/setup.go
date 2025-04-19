package commands

import (
	"fmt"
	"github.com/eliasrenger/workhours/internal/db"
)

func CmdDBSetup() {
	fmt.Println("setting up database...")
	err := db.SetupDB()
	if err != nil {
		fmt.Println("Error setting up database:", err)
		return
	}
	fmt.Println("Database setup complete.")
}
