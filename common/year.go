package common

import (
	"time"
)

func FlipYear(yyyy int) int {
	return yyyy - (yyyy * 2)
}

func TimeToBCE(t time.Time) time.Time {
	return t.AddDate(-2*t.Year(), 0, 0)
}
