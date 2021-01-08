package edtf

import (
	"fmt"
)

type YMD struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

func (ymd *YMD) String() string {
	return fmt.Sprintf("[%T] Y: '%d' M: '%d' D: '%d'", ymd, ymd.Year, ymd.Month, ymd.Day)
}

func (ymd *YMD) Equals(other_ymd *YMD) bool {

	if ymd.Year != other_ymd.Year {
		return false
	}

	if ymd.Month != other_ymd.Month {
		return false
	}

	if ymd.Day != other_ymd.Day {
		return false
	}

	return true
}
