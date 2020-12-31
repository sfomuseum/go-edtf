package common

import (
	"errors"
	"fmt"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/calendar"
	"strconv"
	"time"
)

func DateRangeWithYMDString(str_yyyy string, str_mm string, str_dd string) (*edtf.DateRange, error) {

	if str_yyyy == "" {
		return nil, errors.New("Missing year")
	}

	if str_mm == "" && str_dd != "" {
		return nil, errors.New("Missing month")
	}

	yyyy, err := strconv.Atoi(str_yyyy)

	if err != nil {
		return nil, err
	}

	mm := 0
	dd := 0

	if str_mm != "" {

		m, err := strconv.Atoi(str_mm)

		if err != nil {
			return nil, err
		}

		mm = m
	}

	if str_dd != "" {

		d, err := strconv.Atoi(str_dd)

		if err != nil {
			return nil, err
		}

		dd = d
	}

	return DateRangeWithYMD(yyyy, mm, dd)
}

func DateRangeWithYMD(yyyy int, mm int, dd int) (*edtf.DateRange, error) {

	// See this? If yyyy < 0 then we are dealing with a BCE year
	// which can't be parsed by the time.Parse() function so we're
	// going to set a flag and convert yyyy to a positive number.
	// After we've created time.Time instances below, we'll check to see
	// whether the flag is set and if it is then we'll update the
	// year to be BCE again. One possible gotcha in this approach is
	// that the calendar.DaysInMonth method may return wonky results
	// since it will calculating things on a CE year rather than a BCE
	// year. (20201230/thisisaaronland)

	is_bce := false

	if yyyy < 0 {
		is_bce = true
		yyyy = yyyy - (yyyy * 2)
	}

	if yyyy == 0 {
		return nil, errors.New("Missing year")
	}

	if yyyy > edtf.MAX_YEARS {
		return nil, errors.New("Year exceeds Golang limit")
	}

	if mm == 0 && dd != 0 {
		return nil, errors.New("Missing month")
	}

	lower_yyyy := yyyy
	upper_yyyy := yyyy

	lower_mm := 1
	upper_mm := 12

	lower_dd := 1
	upper_dd := -1

	if mm != 0 {

		if mm > 12 {
			return nil, errors.New("Invalid month")
		}

		lower_mm = mm
		upper_mm = mm
	}

	if dd == 0 {

		lower_dd = 1

		days, err := calendar.DaysInMonth(uint(yyyy), uint(mm))

		if err != nil {
			return nil, err
		}

		upper_dd = int(days)

	} else {

		days, err := calendar.DaysInMonth(uint(yyyy), uint(mm))

		if err != nil {
			return nil, err
		}

		if uint(dd) > days {
			return nil, errors.New("Invalid days for month")
		}

		lower_dd = int(dd)
		upper_dd = lower_dd
	}

	lower_hms := "00:00:00"
	upper_hms := "23:59:59"

	lower_str := fmt.Sprintf("%04d-%02d-%02dT%s", lower_yyyy, lower_mm, lower_dd, lower_hms)
	upper_str := fmt.Sprintf("%04d-%02d-%02dT%s", upper_yyyy, upper_mm, upper_dd, upper_hms)

	// fmt.Println("LOWER", lower_str)
	// fmt.Println("UPPER", upper_str)

	lower_t, err := time.Parse("2006-01-02T15:04:05", lower_str)

	if err != nil {
		return nil, err
	}

	upper_t, err := time.Parse("2006-01-02T15:04:05", upper_str)

	if err != nil {
		return nil, err
	}

	if is_bce {
		lower_t = toBCE(lower_t)
		upper_t = toBCE(upper_t)
	}

	// fmt.Printf("lower %v\n", lower_t)
	// fmt.Printf("upper %v\n", upper_t)

	lower_d := &edtf.Date{
		Time: &lower_t,
	}

	upper_d := &edtf.Date{
		Time: &upper_t,
	}

	dt := &edtf.DateRange{
		Lower: lower_d,
		Upper: upper_d,
	}

	return dt, nil
}

func toBCE(d time.Time) time.Time {
	return d.AddDate(-2*d.Year(), 0, 0)
}
