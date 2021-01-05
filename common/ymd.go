package common

import (
	"errors"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/calendar"
	"strconv"
)

func YMDFromStrings(str_yyyy string, str_mm string, str_dd string) (*edtf.YMD, error) {

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

	mm := 1
	dd := 1

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

	if yyyy == 0 {
		return nil, errors.New("Missing year")
	}

	if yyyy > edtf.MAX_YEARS {
		return nil, edtf.Unsupported("year", strconv.Itoa(yyyy))
	}

	if mm == 0 && dd != 0 {
		return nil, errors.New("Missing month")
	}

	if mm == 0 {
		mm = 1
	} else {

		if mm > 12 {
			return nil, errors.New("Invalid month")
		}
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

		dd = int(days)
	}

	ymd := &edtf.YMD{
		Year:  yyyy,
		Month: mm,
		Day:   dd,
	}

	return ymd, nil
}
