package level0

import (
	"errors"
	// "fmt"
	"github.com/whosonfirst/go-edtf"
	// "github.com/whosonfirst/go-edtf/calendar"
	"regexp"
	// "strconv"
	"strings"
	// "time"
)

const LEVEL int = 0

const PATTERN_DATE string = `^(\d{4})(?:-(\d{2})(?:-(\d{2}))?)?$`
const PATTERN_DATE_TIME string = `^(\d{4})-(\d{2})-(\d{2})T(\d{2}):(\d{2}):(\d{2})(Z|(\+|-)(\d{2})(\:(\d{2}))?)?$`
const PATTERN_TIME_INTERVAL string = `^(\d{4})(?:-(\d{2})(?:-(\d{2}))?)?\/(\d{4})(?:-(\d{2})(?:-(\d{2}))?)?$`

var re_date *regexp.Regexp
var re_date_time *regexp.Regexp
var re_time_interval *regexp.Regexp
var re_level0 *regexp.Regexp

var Tests map[string][]string = map[string][]string{
	"date": []string{
		"1985-04-12",
		"1985-04",
		"1985",
	},
	"date_and_time": []string{
		"1985-04-12T23:20:30",
		"1985-04-12T23:20:30Z",
		"1985-04-12T23:20:30-04",
		"1985-04-12T23:20:30+04:30",
	},
	"time_interval": []string{
		"1964/2008",
		"2004-06/2006-08",
		"2004-02-01/2005-02-08",
		"2004-02-01/2005-02",
		"2004-02-01/2005",
		"2005/2006-02",
	},
}

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

	if IsDateAndTime(edtf_str) {
		return ParseDateAndTime(edtf_str)
	}

	if IsTimeInterval(edtf_str) {
		return ParseTimeInterval(edtf_str)
	}

	return nil, errors.New("Invalid Level 0 string")
}

/*
func dateRangeWithYMD(yyyy string, mm string, dd string) (*edtf.DateRange, error) {

	lower, upper, err := timeRangeWithYMD(yyyy, mm, dd)

	if err != nil {
		return nil, err
	}

	edtf_str, err := ymdToEDTFWithStrings(yyyy, mm, dd)

	if err != nil {
		return nil, err
	}

	dr := &edtf.DateRange{
		Lower: &edtf.Date{
			Time: lower,
			EDTF: edtf_str,
		},
		Upper: &edtf.Date{
			Time: upper,
			EDTF: edtf_str,
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

func ymdToEDTFWithStrings(str_yyyy string, str_mm string, str_dd string) (string, error) {

	if str_yyyy == "" {
		return "", errors.New("Missing year")
	}

	if str_mm == "" && str_dd != "" {
		return "", errors.New("Missing month")
	}

	yyyy, err := strconv.Atoi(str_yyyy)

	if err != nil {
		return "", err
	}

	mm := 0
	dd := 0

	if str_mm != "" {

		m, err := strconv.Atoi(str_mm)

		if err != nil {
			return "", err
		}

		mm = m
	}

	if str_dd != "" {

		d, err := strconv.Atoi(str_dd)

		if err != nil {
			return "", err
		}

		dd = d
	}

	return ymdToEDTF(uint(yyyy), uint(mm), uint(dd))
}

func ymdToEDTF(yyyy uint, mm uint, dd uint) (string, error) {

	if yyyy == 0 {
		return "", errors.New("Missing year")
	}

	if mm == 0 && dd != 0 {
		return "", errors.New("Missing month")
	}

	edtf_parts := []string{
		fmt.Sprintf("%04d", yyyy),
	}

	if mm > 0 {

		if mm > 12 {
			return "", errors.New("Invalid month for year")
		}

		edtf_parts = append(edtf_parts, fmt.Sprintf("%02d", mm))
	}

	if dd > 0 {

		days, err := calendar.DaysInMonth(yyyy, mm)

		if err != nil {
			return "", err
		}

		if dd > days {
			return "", errors.New("Invalid days for month")
		}

		edtf_parts = append(edtf_parts, fmt.Sprintf("%02d", dd))
	}

	edtf_str := strings.Join(edtf_parts, "-")
	return edtf_str, nil
}
*/
