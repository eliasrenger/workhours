package textformatting

import (
	"fmt"

	"github.com/eliasrenger/workhours/internal/models"
)

func EndOfWorkDayFormat(workDay models.WorkDay) string {
	y, m, d := workDay.StartedAt.Date()
	date := fmt.Sprintf("%v-%v-%v", d, m, y)
	return fmt.Sprintf(
		"| Date         | Work Duration | Short Breaks | Tasks \n| %v |        %v |         %v | %v ",
		date, workDay.WorkDuration, workDay.NumberOfQuickBreaks, workDay.TasksWorkedOn,
	)
}
