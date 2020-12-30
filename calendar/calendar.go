package calendar

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func DaysInMonthWithString(yyyymm string) (uint, error) {

	ym := strings.Split(yyyymm, "-")

	if len(ym) != 2 {
		return 0, errors.New("Invalid YYYYMM string")
	}

	yyyy, err := strconv.Atoi(ym[0])

	if err != nil {
		return 0, err
	}

	mm, err := strconv.Atoi(ym[1])

	if err != nil {
		return 0, err
	}

	return DaysInMonth(uint(yyyy), uint(mm))
}

func DaysInMonth(yyyy uint, mm uint) (uint, error) {

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

	return uint(dd), nil
}

// https://stackoverflow.com/questions/51578482/parsing-dates-with-negative-year-in-go
// This needs to be updated for YYYY-MM-DD strings (20201229/thisisaaronland)

func ParseCEDate(value string) (time.Time, error) {

	const layout = "_2 Jan 2006"

	date, err := time.Parse(layout, value)

	if err == nil {
		return date, err
	}

	perr, ok := err.(*time.ParseError)

	if !ok {
		return time.Time{}, err
	}

	if perr.LayoutElem != "2006" {
		return time.Time{}, err
	}

	if !strings.HasPrefix(perr.ValueElem, "-") {
		return time.Time{}, err
	}

	value = strings.Replace(value, perr.ValueElem, perr.ValueElem[1:], 1)

	date, derr := time.Parse(layout, value)

	if derr != nil {
		return time.Time{}, err
	}

	return date.AddDate(-2*date.Year(), 0, 0), derr
}

func ToBCE(d *time.Time) *time.Time {
	new_d := d.AddDate(-2*d.Year(), 0, 0)
	return &new_d
}
