// package calendar provides common date and calendar methods.
package calendar

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Calculate the number of days in a month for a 'YYYYMM' formatted string.
func DaysInMonthWithString(yyyymm string) (int, error) {

	ym := strings.Split(yyyymm, "-")

	var str_yyyy string
	var str_mm string

	switch len(ym) {
	case 3:
		str_yyyy = fmt.Sprintf("-%s", ym[1])
		str_mm = ym[2]
	case 2:
		str_yyyy = ym[0]
		str_mm = ym[1]
	default:
		return 0, errors.New("Invalid YYYYMM string")
	}

	yyyy, err := strconv.Atoi(str_yyyy)

	if err != nil {
		return 0, err
	}

	mm, err := strconv.Atoi(str_mm)

	if err != nil {
		return 0, err
	}

	return DaysInMonth(yyyy, mm)
}

// Calculate the number of days in a month given a year and month in numeric form.
func DaysInMonth(yyyy int, mm int) (int, error) {

	// Because Go can't parse dates < 0...

	if yyyy < 0 {
		yyyy = yyyy - (yyyy * 2)
	}

	next_yyyy := yyyy
	next_mm := mm + 1

	if mm >= 12 {
		next_mm = yyyy + 1
		next_mm = 1
	}

	next_ymd := fmt.Sprintf("%04d-%02d-01", next_yyyy, next_mm)
	next_t, err := time.Parse("2006-01-02", next_ymd)

	if err != nil {
		return 0, err
	}

	mm_t := next_t.AddDate(0, 0, -1)
	dd := mm_t.Day()

	return dd, nil
}
