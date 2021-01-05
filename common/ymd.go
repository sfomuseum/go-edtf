package common

import (
	"errors"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/calendar"
	"strconv"
	"strings"
	"fmt"
)

func YMDFromStringDate(d *StringDate) (*edtf.YMD, error) {
	return YMDFromStrings(d.Year, d.Month, d.Day)
}

func YMDFromString(str_ymd string) (*edtf.YMD, error) {

	yyyy := ""
	mm := ""
	dd := ""
	
	parts := strings.Split(str_ymd, "-")

	switch len(parts) {
	case 4:
		yyyy = "-" + parts[1]
		mm = parts[2]
		dd = parts[3]		
	case 3:
		yyyy = parts[0]
		mm = parts[1]
		dd = parts[2]
	case 2:
		yyyy = parts[0]
		mm = parts[1]
	case 1:
		yyyy = parts[0]
	default:
		return nil, errors.New("Invalid YMD string")
	}

	fmt.Println("YEAR", yyyy)
	
	return YMDFromStrings(yyyy, mm, dd)
}

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

	is_bce := false

	if yyyy < 0 {
		is_bce = true
		yyyy = FlipYear(yyyy)
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
	}

	if is_bce {
		yyyy = FlipYear(yyyy)
	}
	
	ymd := &edtf.YMD{
		Year:  yyyy,
		Month: mm,
		Day:   dd,
	}

	return ymd, nil
}
