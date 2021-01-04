package common

import (
	"errors"
	"fmt"
	"github.com/whosonfirst/go-edtf"
	"strconv"
	"strings"
	"regexp"
)

// move these in to the re package

var re_qualifier_prefix *regexp.Regexp
var re_qualifier_suffix *regexp.Regexp

func init () {

	pattern_date := `(\-?[0-9X]{4}|[0-3X][0-9X])`
	pattern_qualifier := `(\[|\{)`
	
	re_qualifier_prefix = regexp.MustCompile(`^` + pattern_qualifier + `?` + pattern_date + `$`)
	re_qualifier_suffix = regexp.MustCompile(`^` + pattern_date + pattern_qualifier + `?$`)	
}

func DateRangeWithString(edtf_str string) (*edtf.DateRange, error) {

	yyyy := ""
	mm := ""
	dd := ""

	yyyy_q := ""
	mm_q := ""
	dd_q := ""	

	start_yyyy := ""
	start_mm := ""
	start_dd := ""

	end_yyyy := ""
	end_mm := ""
	end_dd := ""

	precision := edtf.NONE
	
	parts := strings.Split(edtf_str, "-")
	count := len(parts)

	switch count {
	case 3:
		
		y, q, err := parseDate(parts[0])

		if err != nil {
			return nil, err
		}

		yyyy = y
		yyyy_q = q

		m, q, err := parseDate(parts[1])

		if err != nil {
			return nil, err
		}

		mm = m
		mm_q = q

		d, q, err := parseDate(parts[2])

		if err != nil {
			return nil, err
		}
		
		dd = d
		dd_q = q
		
	case 2:

		y, q, err := parseDate(parts[0])

		if err != nil {
			return nil, err
		}

		yyyy = y
		yyyy_q = q

		m, q, err := parseDate(parts[1])

		if err != nil {
			return nil, err
		}

		mm = m
		mm_q = q
		
	case 1:

		y, q, err := parseDate(parts[0])

		if err != nil {
			return nil, err
		}

		yyyy = y
		yyyy_q = q
		
	default:
		return nil, edtf.Invalid("date string", edtf_str)
	}	

	if strings.HasSuffix(yyyy, "X") {

		start_m := int64(0)
		end_m := int64(0)

		start_c := int64(0)
		end_c := int64(900)

		start_d := int64(0)
		end_d := int64(90)

		start_y := int64(0)
		end_y := int64(9)

		if string(yyyy[0]) == "X" {
			return nil, edtf.NotImplemented("date", edtf_str)
		} else {

			m, err := strconv.ParseInt(string(yyyy[0]), 10, 32)

			if err != nil {
				return nil, err
			}

			start_m = m * 1000
			end_m = start_m
		}

		if string(yyyy[1]) != "X" {

			c, err := strconv.ParseInt(string(yyyy[1]), 10, 32)

			if err != nil {
				return nil, err
			}

			start_c = c * 100
			end_c = start_c
		}

		if string(yyyy[2]) != "X" {

			d, err := strconv.ParseInt(string(yyyy[2]), 10, 32)

			if err != nil {
				return nil, err
			}

			start_d = d * 10
			end_d = start_d
		}

		if string(yyyy[3]) != "X" {

			y, err := strconv.ParseInt(string(yyyy[3]), 10, 32)

			if err != nil {
				return nil, err
			}

			start_y = y * 1
			end_y = start_y
		}

		start_ymd := start_m + start_c + start_d + start_y
		end_ymd := end_m + end_c + end_d + end_y

		// fmt.Printf("OMG '%s' '%d' '%d' '%d' '%d' '%d'\n", yyyy, start_m, start_c, start_d, start_y, start_ymd)
		// fmt.Printf("WTF '%s' '%d' '%d' '%d' '%d' '%d'\n", yyyy, end_m, end_c, end_d, end_y, end_ymd)

		start_yyyy = strconv.FormatInt(start_ymd, 10)
		end_yyyy = strconv.FormatInt(end_ymd, 10)

		precision.AddFlag(edtf.ANNUAL)
	}

	if strings.HasSuffix(mm, "X") {

		// this does not account for 1985-24, etc.

		if strings.HasPrefix(mm, "X") {

			start_mm = "01"
			end_mm = "12"
		} else {

			start_mm = "10"
			end_mm = "12"
		}

		precision.AddFlag(edtf.MONTHLY)
	}

	if strings.HasSuffix(dd, "X") {

		switch string(dd[0]) {
		case "X":
			start_dd = "01"
			end_dd = ""
		case "1":
			start_dd = "10"
			end_dd = "19"
		case "2":
			start_dd = "20"
			end_dd = "29"
		case "3":
			start_dd = "30"
			end_dd = ""
		default:
			return nil, edtf.Invalid("date", edtf_str)
		}

		precision.AddFlag(edtf.DAILY)
	}

	lower_t, err := TimeWithYMDString(start_yyyy, start_mm, start_dd, edtf.HMS_LOWER)

	if err != nil {
		return nil, err
	}

	upper_t, err := TimeWithYMDString(end_yyyy, end_mm, end_dd, edtf.HMS_UPPER)

	if err != nil {
		return nil, err
	}	

	lower_d := &edtf.Date{
		Time: lower_t,
	}

	upper_d := &edtf.Date{
		Time: upper_t,
	}

	dt := &edtf.DateRange{
		Lower: lower_d,
		Upper: upper_d,
	}

	fmt.Println(yyyy_q, mm_q, dd_q)
	return dt, nil
}

// DEPRECATED

func DateRangeWithYMDStringCombined(ymd string) (*edtf.DateRange, error) {

	return DateRangeWithString(ymd)
}

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

	lower_t, err := TimeWithYMD(yyyy, mm, dd, edtf.HMS_LOWER)

	if err != nil {
		return nil, err
	}

	upper_t, err := TimeWithYMD(yyyy, mm, dd, edtf.HMS_UPPER)

	if err != nil {
		return nil, err
	}	
	

	lower_d := &edtf.Date{
		Time: lower_t,
	}

	upper_d := &edtf.Date{
		Time: upper_t,
	}

	dt := &edtf.DateRange{
		Lower: lower_d,
		Upper: upper_d,
	}

	return dt, nil
}


func EmptyDateRange() *edtf.DateRange {

	lower_d := &edtf.Date{}
	upper_d := &edtf.Date{}

	dt := &edtf.DateRange{
		Lower: lower_d,
		Upper: upper_d,
	}

	return dt
}

func parseDate(date string) (string, string, error) {

	m := re_qualifier_prefix.FindStringSubmatch(date)

	if len(m) == 2 {
		return m[1], m[0], nil
	}

	m = re_qualifier_suffix.FindStringSubmatch(date)

	if len(m) == 2 {
		return m[0], m[1], nil
	}

	return "", "", edtf.Invalid("date", date)
}

