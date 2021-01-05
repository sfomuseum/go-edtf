package common

import (
	"errors"
	_ "fmt"
	"github.com/whosonfirst/go-edtf"
	"regexp"
	"strconv"
	"strings"
)

type Qualifier struct {
	Value string
	Type  string
}

// move these in to the re package

var re_ymd_string *regexp.Regexp

var re_qualifier_individual *regexp.Regexp
var re_qualifier_group *regexp.Regexp

func init() {

	// move this in to the re package

	qualifiers := []string{
		`\` + edtf.UNCERTAIN,
		edtf.APPROXIMATE,
		edtf.UNCERTAIN_AND_APPROXIMATE,
	}

	pattern_qualifier := `[` + strings.Join(qualifiers, "") + `]`

	pattern_year := `\-?[0-9X]{4}`
	pattern_month := `(?:[0X][1-9X]|[1X][0-2X])`
	pattern_day := `(?:[012X][1-9X]|[3X][01X])`

	pattern_yyyy := `(` + pattern_qualifier + `?` + pattern_year + `|` + pattern_year + pattern_qualifier + `?)`
	pattern_mm := `(` + pattern_qualifier + `?` + pattern_month + `|` + pattern_month + pattern_qualifier + `?)`
	pattern_dd := `(` + pattern_qualifier + `?` + pattern_day + `|` + pattern_day + pattern_qualifier + `?)`

	pattern_ymd := `^` + pattern_yyyy + `(?:\-` + pattern_mm + `(?:\-` + pattern_dd + `)?)?$`

	pattern_date := `(` + pattern_year + `|(?:` + pattern_month + `)|(?:` + pattern_day + `))`

	re_ymd_string = regexp.MustCompile(pattern_ymd)

	re_qualifier_individual = regexp.MustCompile(`^(` + pattern_qualifier + `)?` + pattern_date + `$`)
	re_qualifier_group = regexp.MustCompile(`^` + pattern_date + `(` + pattern_qualifier + `)?$`)
}

// PLEASE RENAME ME TO BE DateRangeWithYMDString

func DateRangeWithString(edtf_str string) (*edtf.DateRange, error) {

	precision := edtf.NONE
	uncertain := edtf.NONE
	approximate := edtf.NONE

	parts := re_ymd_string.FindStringSubmatch(edtf_str)
	count := len(parts)

	if count != 4 {
		return nil, edtf.Invalid("date", edtf_str)
	}

	yyyy := parts[1]
	mm := parts[2]
	dd := parts[3]

	var yyyy_q *Qualifier
	var mm_q *Qualifier
	var dd_q *Qualifier

	if yyyy != "" {

		y, q, err := parseYMDComponent(yyyy)

		if err != nil {
			return nil, err
		}

		yyyy = y
		yyyy_q = q
	}

	if mm != "" {

		m, q, err := parseYMDComponent(mm)

		if err != nil {
			return nil, err
		}

		mm = m
		mm_q = q
	}

	if dd != "" {

		d, q, err := parseYMDComponent(dd)

		if err != nil {
			return nil, err
		}

		dd = d
		dd_q = q
	}

	start_yyyy := yyyy
	start_mm := mm
	start_dd := dd

	end_yyyy := start_yyyy
	end_mm := start_mm
	end_dd := start_dd

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

	dr := &edtf.DateRange{
		Lower: lower_d,
		Upper: upper_d,
	}

	if yyyy_q != nil {

		switch yyyy_q.Value {
		case edtf.UNCERTAIN:
			uncertain.AddFlag(edtf.ANNUAL)
		case edtf.APPROXIMATE:
			approximate.AddFlag(edtf.ANNUAL)
		case edtf.UNCERTAIN_AND_APPROXIMATE:
			uncertain.AddFlag(edtf.ANNUAL)
			approximate.AddFlag(edtf.ANNUAL)
		default:
			// pass
		}
	}

	if mm_q != nil {

		switch mm_q.Value {
		case edtf.UNCERTAIN:
			uncertain.AddFlag(edtf.MONTHLY)
		case edtf.APPROXIMATE:
			approximate.AddFlag(edtf.MONTHLY)
		case edtf.UNCERTAIN_AND_APPROXIMATE:
			uncertain.AddFlag(edtf.MONTHLY)
			approximate.AddFlag(edtf.MONTHLY)
		default:
			// pass
		}
	}

	if dd_q != nil {

		switch dd_q.Value {
		case edtf.UNCERTAIN:
			uncertain.AddFlag(edtf.DAILY)
		case edtf.APPROXIMATE:
			approximate.AddFlag(edtf.DAILY)
		case edtf.UNCERTAIN_AND_APPROXIMATE:
			uncertain.AddFlag(edtf.DAILY)
			approximate.AddFlag(edtf.DAILY)
		default:
			// pass
		}
	}

	if uncertain != edtf.NONE {
		dr.Lower.Uncertain = uncertain
		dr.Upper.Uncertain = uncertain
	}

	if approximate != edtf.NONE {
		dr.Lower.Approximate = approximate
		dr.Upper.Approximate = approximate
	}

	if precision != edtf.NONE {
		dr.Lower.Unspecified = precision
		dr.Upper.Unspecified = precision
	}

	return dr, nil
}

// PLEASE RENAME ME (20210104/thisisaaronland)

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

func parseYMDComponent(date string) (string, *Qualifier, error) {

	m := re_qualifier_individual.FindStringSubmatch(date)

	if len(m) == 3 {

		var q *Qualifier

		if m[1] != "" {

			q = &Qualifier{
				Type:  "Individual",
				Value: m[1],
			}
		}

		return m[2], q, nil
	}

	m = re_qualifier_group.FindStringSubmatch(date)

	if len(m) == 3 {

		var q *Qualifier

		if m[2] != "" {

			q = &Qualifier{
				Type:  "Individual",
				Value: m[1],
			}
		}

		return m[1], q, nil
	}

	return "", nil, edtf.Invalid("date", date)
}
