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

	if yyyy == 0 {
		return nil, errors.New("Missing year")
	}

	if mm == 0 && dd != 0 {
		return nil, errors.New("Missing month")
	}

	if dd == 0 {

		days, err := calendar.DaysInMonth(uint(yyyy), uint(mm))

		if err != nil {
			return nil, err
		}

		dd = int(days)

	} else {

		days, err := calendar.DaysInMonth(uint(yyyy), uint(mm))

		if err != nil {
			return nil, err
		}

		if uint(dd) > days {
			return nil, errors.New("Invalid days for month")
		}
	}

	lower_hms := "00:00:00"
	upper_hms := "23:59:59"

	lower_str := fmt.Sprintf("%04d-%02d-%02dT%s", yyyy, mm, dd, lower_hms)
	upper_str := fmt.Sprintf("%04d-%02d-%02dT%s", yyyy, mm, dd, upper_hms)

	lower_t, err := time.Parse("2006-01-02T15:04:05", lower_str)

	if err != nil {
		return nil, err
	}

	upper_t, err := time.Parse("2006-01-02T15:04:05", upper_str)

	if err != nil {
		return nil, err
	}

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
