package common

import (
	"errors"
	"fmt"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/calendar"
	"strconv"
	"time"
)

func TimeWithYMDString(str_yyyy string, str_mm string, str_dd string, hms string) (*time.Time, error) {

	yyyy, mm, dd, err := YMDFromStrings(str_yyyy, str_mm, str_dd)

	if err != nil {
		return nil, err
	}
	
	return TimeWithYMD(yyyy, mm, dd, hms)
}

func TimeWithYMD(yyyy int, mm int, dd int, hms string) (*time.Time, error) {

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
		yyyy = FlipYear(yyyy)
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

	t_str := fmt.Sprintf("%04d-%02d-%02dT%s", yyyy, mm, dd, hms)

	t, err := time.Parse("2006-01-02T15:04:05", t_str)

	if err != nil {
		return nil, err
	}

	if is_bce {
		t = TimeToBCE(t)
	}

	return &t, nil
}
