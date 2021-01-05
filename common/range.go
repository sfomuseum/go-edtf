package common

import (
	"fmt"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/calendar"
	"github.com/whosonfirst/go-edtf/re"
	"strconv"
	"strings"
	//"time"
)

type Qualifier struct {
	Value string
	Type  string
}

// StringWhatever is a bad naming convention - please make me better
// (20210105/thisisaaronland)

type StringDate struct {
	Year  string
	Month string
	Day   string
}

func (d *StringDate) String() string {
	return fmt.Sprintf("[[%T] Y: '%s' M: '%s' D: '%s']", d, d.Year, d.Month, d.Day)
}

func (d *StringDate) Equals(other_d *StringDate) bool {

	if d.Year != other_d.Year {
		return false
	}

	if d.Month != other_d.Month {
		return false
	}

	if d.Day != other_d.Day {
		return false
	}

	return true
}

type StringRange struct {
	Start       *StringDate
	End         *StringDate
	Precision   edtf.Precision
	Uncertain   edtf.Precision
	Approximate edtf.Precision
	EDTF        string
}

func (r *StringRange) String() string {
	return fmt.Sprintf("[[%T] Start: '%s' End: '%s']", r, r.Start, r.End)
}

func StringRangeFromEDTF(edtf_str string) (*StringRange, error) {

	precision := edtf.NONE
	uncertain := edtf.NONE
	approximate := edtf.NONE

	parts := re.YMD.FindStringSubmatch(edtf_str)
	count := len(parts)

	if count != 4 {
		return nil, edtf.Invalid("date", edtf_str)
	}

	yyyy := parts[1]
	mm := parts[2]
	dd := parts[3]

	// fmt.Printf("DATE Y: '%s' M: '%s' D: '%s'\n", yyyy, mm, dd)

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

	if dd_q != nil && dd_q.Type == "Group" {
		precision.AddFlag(edtf.ANNUAL)
		precision.AddFlag(edtf.MONTHLY)
		precision.AddFlag(edtf.DAILY)
	}

	if mm_q != nil && mm_q.Type == "Group" {
		precision.AddFlag(edtf.ANNUAL)
		precision.AddFlag(edtf.MONTHLY)
	}

	if yyyy_q != nil && yyyy_q.Type == "Group" {
		precision.AddFlag(edtf.ANNUAL)
	}

	if yyyy_q != nil && yyyy_q.Type == "Individual" {

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

	if mm_q != nil && mm_q.Type == "Individual" {

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

	if dd_q != nil && dd_q.Type == "Individual" {

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

	if start_mm == "" {
		start_mm = "01"
	}

	if start_dd == "" {
		start_dd = "01"
	}

	if end_mm == "" {
		end_mm = "12"
	}

	if end_dd == "" {

		yyyymm := fmt.Sprintf("%s-%s", end_yyyy, end_mm)

		dd, err := calendar.DaysInMonthWithString(yyyymm)

		if err != nil {
			return nil, err
		}

		end_dd = strconv.Itoa(int(dd))
	}

	start := &StringDate{
		Year:  start_yyyy,
		Month: start_mm,
		Day:   start_dd,
	}

	end := &StringDate{
		Year:  end_yyyy,
		Month: end_mm,
		Day:   end_dd,
	}

	r := &StringRange{
		Start:       start,
		End:         end,
		Precision:   precision,
		Uncertain:   uncertain,
		Approximate: approximate,
		EDTF:        edtf_str,
	}

	return r, nil
}

func DateRangeWithStringRange(r *StringRange) (*edtf.DateRange, error) {

	start_yyyy := r.Start.Year
	start_mm := r.Start.Month
	start_dd := r.Start.Day

	end_yyyy := r.End.Year
	end_mm := r.End.Month
	end_dd := r.End.Day

	edtf_str := r.EDTF

	// fmt.Println("LOWER", edtf_str, start_yyyy, start_mm, start_dd, edtf.HMS_LOWER)
	// fmt.Println("UPPER", edtf_str, end_yyyy, end_mm, end_dd, edtf.HMS_UPPER)

	start_ymd, err := YMDFromStrings(start_yyyy, start_mm, start_dd)

	if err != nil {
		return nil, err
	}

	fmt.Println("DATERANGE START", start_ymd, r.Start)

	end_ymd, err := YMDFromStrings(end_yyyy, end_mm, end_dd)

	if err != nil {
		return nil, err
	}

	lower_t, err := TimeWithYMD(start_ymd, edtf.HMS_LOWER)

	if err != nil {
		return nil, err
	}

	upper_t, err := TimeWithYMD(end_ymd, edtf.HMS_UPPER)

	if err != nil {
		return nil, err
	}

	// fmt.Println("LOWER", edtf_str, lower_t.Format(time.RFC3339))
	// fmt.Println("UPPER", edtf_str, upper_t.Format(time.RFC3339))

	lower_d := &edtf.Date{
		Time: lower_t,
		YMD:  start_ymd,
	}

	upper_d := &edtf.Date{
		Time: upper_t,
		YMD:  end_ymd,
	}

	dr := &edtf.DateRange{
		EDTF:  edtf_str,
		Lower: lower_d,
		Upper: upper_d,
	}

	if r.Uncertain != edtf.NONE {
		dr.Lower.Uncertain = r.Uncertain
		dr.Upper.Uncertain = r.Uncertain
	}

	if r.Approximate != edtf.NONE {
		dr.Lower.Approximate = r.Approximate
		dr.Upper.Approximate = r.Approximate
	}

	if r.Precision != edtf.NONE {
		dr.Lower.Unspecified = r.Precision
		dr.Upper.Unspecified = r.Precision
	}

	return dr, nil
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

	m := re.QualifiedIndividual.FindStringSubmatch(date)

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

	m = re.QualifiedGroup.FindStringSubmatch(date)

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
