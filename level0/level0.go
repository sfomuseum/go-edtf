package level0

import (
	"errors"
	"fmt"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/calendar"
	"regexp"
	"strings"
	"time"
)

const LEVEL int = 0

const PATTERN_DATE string = `^(\d{4})(?:-(\d{2})(?:-(\d{2}))?)?$`
const PATTERN_DATE_TIME string = `^(\d{4})-(\d{2})-(\d{2})T(\d{2}):(\d{2}):(\d{2})(Z|(\+|-)(\d{2})(\:(\d{2}))?)?$`
const PATTERN_TIME_INTERVAL string = `^(\d{4})(?:-(\d{2})(?:-(\d{2}))?)?\/(\d{4})(?:-(\d{2})(?:-(\d{2}))?)?$`

var re_date *regexp.Regexp
var re_date_time *regexp.Regexp
var re_time_interval *regexp.Regexp
var re_level0 *regexp.Regexp

func init() {

	re_date = regexp.MustCompile(PATTERN_DATE)

	re_date_time = regexp.MustCompile(PATTERN_DATE_TIME)

	re_time_interval = regexp.MustCompile(PATTERN_TIME_INTERVAL)

	level0_patterns := []string{
		PATTERN_DATE,
		PATTERN_DATE_TIME,
		PATTERN_TIME_INTERVAL,
	}

	re_level0 = regexp.MustCompile(`(` + strings.Join(level0_patterns, "|") + `)`)
}

func IsLevel0(edtf_str string) bool {
	return re_level0.MatchString(edtf_str)
}

func ParseString(edtf_str string) (*edtf.EDTFDate, error) {

	if IsDate(edtf_str) {
		return ParseDate(edtf_str)
	}

	if IsDateTime(edtf_str) {
		return ParseDateTime(edtf_str)
	}

	if IsTimeInterval(edtf_str) {
		return ParseTimeInterval(edtf_str)
	}

	return nil, errors.New("Invalid Level 0 string")
}

/*

Date

    complete representation:            [year][“-”][month][“-”][day]
    Example 1          ‘1985-04-12’ refers to the calendar date 1985 April 12th with day precision.
    reduced precision for year and month:   [year][“-”][month]
    Example 2          ‘1985-04’ refers to the calendar month 1985 April with month precision.
    reduced precision for year:  [year]
    Example 3          ‘1985’ refers to the calendar year 1985 with year precision.

*/

func IsDate(edtf_str string) bool {
	return re_date.MatchString(edtf_str)
}

func ParseDate(edtf_str string) (*edtf.EDTFDate, error) {

	m := re_date.FindStringSubmatch(edtf_str)

	if len(m) != 4 {
		return nil, errors.New("Invalid Level 0 date string")
	}

	yyyy := m[1]
	mm := m[2]
	dd := m[3]

	lower_range, err := dateRangeWithYMD(yyyy, mm, dd)

	if err != nil {
		return nil, err
	}

	upper_range := lower_range

	if err != nil {
		return nil, err
	}

	d := &edtf.EDTFDate{
		Lower: lower_range,
		Upper: upper_range,
		Raw:   edtf_str,
		Level: LEVEL,
	}

	return d, nil
}

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

func IsDateTime(edtf_str string) bool {
	return re_date_time.MatchString(edtf_str)
}

func ParseDateTime(edtf_str string) (*edtf.EDTFDate, error) {

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

	upper_range := &edtf.DateRange{
		Upper: upper_date,
		Lower: upper_date,
	}

	lower_range := &edtf.DateRange{
		Upper: lower_date,
		Lower: lower_date,
	}

	d := &edtf.EDTFDate{
		Upper: upper_range,
		Lower: lower_range,
		Raw:   edtf_str,
		Level: LEVEL,
	}

	return d, nil
}

/*

Time Interval

EDTF Level 0 adopts representations of a time interval where both the start and end are dates: start and end date only; that is, both start and duration, and duration and end, are excluded. Time of day is excluded.

    Example 1          ‘1964/2008’ is a time interval with calendar year precision, beginning sometime in 1964 and ending sometime in 2008.
    Example 2          ‘2004-06/2006-08’ is a time interval with calendar month precision, beginning sometime in June 2004 and ending sometime in August of 2006.
    Example 3          ‘2004-02-01/2005-02-08’ is a time interval with calendar day precision, beginning sometime on February 1, 2004 and ending sometime on February 8, 2005.
    Example 4          ‘2004-02-01/2005-02’ is a time interval beginning sometime on February 1, 2004 and ending sometime in February 2005. Since the start endpoint precision (day) is different than that of the end endpoint (month) the precision of the time interval at large is undefined.
    Example 5          ‘2004-02-01/2005’ is a time interval beginning sometime on February 1, 2004 and ending sometime in 2005. The start endpoint has calendar day precision and the end endpoint has calendar year precision. Similar to the previous example, the precision of the time interval at large is undefined.
    Example 6          ‘2005/2006-02’ is a time interval beginning sometime in 2005 and ending sometime in February 2006.

*/

func IsTimeInterval(edtf_str string) bool {
	return re_time_interval.MatchString(edtf_str)
}

func ParseTimeInterval(edtf_str string) (*edtf.EDTFDate, error) {

	m := re_time_interval.FindStringSubmatch(edtf_str)

	if len(m) != 7 {
		return nil, errors.New("Invalid Level 0 time interval string")
	}

	start_yyyy := m[1]
	start_mm := m[2]
	start_dd := m[3]

	end_yyyy := m[4]
	end_mm := m[5]
	end_dd := m[6]

	lower_range, err := dateRangeWithYMD(start_yyyy, start_mm, start_dd)

	if err != nil {
		return nil, err
	}

	upper_range, err := dateRangeWithYMD(end_yyyy, end_mm, end_dd)

	if err != nil {
		return nil, err
	}

	d := &edtf.EDTFDate{
		Lower: lower_range,
		Upper: upper_range,
		Raw:   edtf_str,
		Level: LEVEL,
	}

	return d, nil
}

func dateRangeWithYMD(yyyy string, mm string, dd string) (*edtf.DateRange, error) {

	lower, upper, err := timeRangeWithYMD(yyyy, mm, dd)

	if err != nil {
		return nil, err
	}

	dr := &edtf.DateRange{
		Lower: &edtf.Date{
			Time: lower,
		},
		Upper: &edtf.Date{
			Time: upper,
		},
	}

	return dr, nil
}

func timeRangeWithYMD(yyyy string, mm string, dd string) (*time.Time, *time.Time, error) {

	var lower_yyyy string
	var lower_mm string
	var lower_dd string

	var upper_yyyy string
	var upper_mm string
	var upper_dd string

	if yyyy == "" {
		return nil, nil, errors.New("Missing yyyy")
	}

	if mm == "" && dd == "" {

		lower_yyyy = yyyy
		lower_mm = "01"
		lower_dd = "01"

		upper_yyyy = yyyy
		upper_mm = "12"
		upper_dd = "31"

	} else if dd == "" {

		lower_yyyy = yyyy
		lower_mm = mm
		lower_dd = "01"

		upper_yyyy = yyyy
		upper_mm = mm

		upper_ym := fmt.Sprintf("%s-%s", upper_yyyy, upper_mm)

		dd, err := calendar.DaysInMonthWithString(upper_ym)

		if err != nil {
			return nil, nil, err
		}

		upper_dd = fmt.Sprintf("%02d", dd)

	} else {
		upper_yyyy = yyyy
		upper_mm = mm
		upper_dd = dd

		lower_yyyy = yyyy
		lower_mm = mm
		lower_dd = dd
	}

	lower_hms := "00:00:00"
	upper_hms := "23:59:59"

	upper_str := fmt.Sprintf("%s-%s-%sT%s", upper_yyyy, upper_mm, upper_dd, upper_hms)
	lower_str := fmt.Sprintf("%s-%s-%sT%s", lower_yyyy, lower_mm, lower_dd, lower_hms)

	upper_t, err := time.Parse("2006-01-02T15:04:05", upper_str)

	if err != nil {
		return nil, nil, err
	}

	lower_t, err := time.Parse("2006-01-02T15:04:05", lower_str)

	if err != nil {
		return nil, nil, err
	}

	return &lower_t, &upper_t, nil
}
