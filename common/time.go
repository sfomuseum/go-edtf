package common

import (
	"fmt"
	"github.com/sfomuseum/go-edtf"
	"time"
)

func TimeWithYMDString(str_yyyy string, str_mm string, str_dd string, hms string) (*time.Time, error) {

	ymd, err := YMDFromStrings(str_yyyy, str_mm, str_dd)

	if err != nil {
		return nil, err
	}

	return TimeWithYMD(ymd, hms)
}

func TimeWithYMD(ymd *edtf.YMD, hms string) (*time.Time, error) {

	// See this? If yyyy < 0 then we are dealing with a BCE year
	// which can't be parsed by the time.Parse() function so we're
	// going to set a flag and convert yyyy to a positive number.
	// After we've created time.Time instances below, we'll check to see
	// whether the flag is set and if it is then we'll update the
	// year to be BCE again. One possible gotcha in this approach is
	// that the calendar.DaysInMonth method may return wonky results
	// since it will calculating things on a CE year rather than a BCE
	// year. (20201230/thisisaaronland)

	yyyy := ymd.Year
	mm := ymd.Month
	dd := ymd.Day

	is_bce := false

	if yyyy < 0 {
		is_bce = true
		yyyy = FlipYear(yyyy)
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
