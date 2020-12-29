package level0

import (
	"errors"
	"github.com/whosonfirst/go-edtf"
	"strings"
	"time"
)

/*

Date and Time

    [date][“T”][time]
    Complete representations for calendar date and (local) time of day
    Example 1          ‘1985-04-12T23:20:30’ refers to the date 1985 April 12th at 23:20:30 local time.
     [dateI][“T”][time][“Z”]
    Complete representations for calendar date and UTC time of day
    Example 2       ‘1985-04-12T23:20:30Z’ refers to the date 1985 April 12th at 23:20:30 UTC time.
    [dateI][“T”][time][shiftHour]
    Date and time with timeshift in hours (only)
    Example 3       ‘1985-04-12T23:20:30-04’ refers to the date 1985 April 12th time of day 23:20:30 with time shift of 4 hours behind UTC.
    [dateI][“T”][time][shiftHourMinute]
    Date and time with timeshift in hours and minutes
    Example 4       ‘1985-04-12T23:20:30+04:30’ refers to the date 1985 April 12th,  time of day  23:20:30 with time shift of 4 hours and 30 minutes ahead of UTC.

*/

func IsDateAndTime(edtf_str string) bool {
	return re_date_time.MatchString(edtf_str)
}

func ParseDateAndTime(edtf_str string) (*edtf.EDTFDate, error) {

	m := re_date_time.FindStringSubmatch(edtf_str)

	if len(m) != 12 {
		return nil, errors.New("Invalid Level 0 date and time string")
	}

	t_fmt := "2006-01-02T15:04:05"

	if m[7] == "Z" {
		t_fmt = "2006-01-02T15:04:05Z"
	}

	if m[8] == "-" || m[8] == "+" {

		if strings.HasPrefix(m[10], ":") {
			t_fmt = "2006-01-02T15:04:05-07:00"
		} else {
			t_fmt = "2006-01-02T15:04:05-07"
		}
	}

	t, err := time.Parse(t_fmt, edtf_str)

	if err != nil {
		return nil, err
	}

	upper_date := &edtf.Date{
		Time: &t,
	}

	lower_date := &edtf.Date{
		Time: &t,
	}

	start := &edtf.DateRange{
		Lower: lower_date,
		Upper: lower_date,
	}

	end := &edtf.DateRange{
		Lower: upper_date,
		Upper: upper_date,
	}

	d := &edtf.EDTFDate{
		Start: start,
		End:   end,
		EDTF:  edtf_str,
		Level: LEVEL,
	}

	return d, nil
}