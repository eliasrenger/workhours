package utils

import "time"

func IsSameDate(t1 time.Time, t2 time.Time) bool {
	ay, am, ad := t1.Date()
	by, bm, bd := t2.Date()
	return ay == by && am == bm && ad == bd
}

func IsDateCorrectFormat(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}
