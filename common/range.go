package common

import (
	"fmt"
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/calendar"
	"github.com/sfomuseum/go-edtf/re"
	"strconv"
	"strings"
)

type Qualifier struct {
	Value string
	Type  string
}

func (q *Qualifier) String() string {
	return fmt.Sprintf("[%T] Value: '%s' Type: '%s'", q, q.Value, q.Type)
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

func StringRangeFromYMD(edtf_str string) (*StringRange, error) {

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

	if yyyy != "" && mm != "" && dd != "" {
		precision.AddFlag(edtf.DAY)
	} else if yyyy != "" && mm != "" {
		precision.AddFlag(edtf.MONTH)
	} else if yyyy != "" {
		precision.AddFlag(edtf.YEAR)
	}

	// fmt.Println("PRECISION -", edtf_str, precision)

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

	// fmt.Println("YYYY", yyyy_q)
	// fmt.Println("MM", mm_q)
	// fmt.Println("DD", dd_q)

	if dd_q != nil && dd_q.Type == "Group" {

		// precision.AddFlag(edtf.YEAR)
		// precision.AddFlag(edtf.MONTH)
		// precision.AddFlag(edtf.DAY)

		switch dd_q.Value {
		case edtf.UNCERTAIN:
			uncertain.AddFlag(edtf.YEAR)
			uncertain.AddFlag(edtf.MONTH)
			uncertain.AddFlag(edtf.DAY)
		case edtf.APPROXIMATE:
			approximate.AddFlag(edtf.YEAR)
			approximate.AddFlag(edtf.MONTH)
			approximate.AddFlag(edtf.DAY)
		case edtf.UNCERTAIN_AND_APPROXIMATE:
			uncertain.AddFlag(edtf.YEAR)
			uncertain.AddFlag(edtf.MONTH)
			uncertain.AddFlag(edtf.DAY)
			approximate.AddFlag(edtf.YEAR)
			approximate.AddFlag(edtf.MONTH)
			approximate.AddFlag(edtf.DAY)
		default:
			// pass
		}

	}

	if mm_q != nil && mm_q.Type == "Group" {

		// precision.AddFlag(edtf.YEAR)
		// precision.AddFlag(edtf.MONTH)

		switch mm_q.Value {
		case edtf.UNCERTAIN:
			uncertain.AddFlag(edtf.YEAR)
			uncertain.AddFlag(edtf.MONTH)
		case edtf.APPROXIMATE:
			approximate.AddFlag(edtf.YEAR)
			approximate.AddFlag(edtf.MONTH)
		case edtf.UNCERTAIN_AND_APPROXIMATE:
			uncertain.AddFlag(edtf.YEAR)
			uncertain.AddFlag(edtf.MONTH)
			approximate.AddFlag(edtf.YEAR)
			approximate.AddFlag(edtf.MONTH)
		default:
			// pass
		}

	}

	if yyyy_q != nil && yyyy_q.Type == "Group" {

		// precision.AddFlag(edtf.YEAR)

		switch yyyy_q.Value {
		case edtf.UNCERTAIN:
			uncertain.AddFlag(edtf.YEAR)
		case edtf.APPROXIMATE:
			approximate.AddFlag(edtf.YEAR)
		case edtf.UNCERTAIN_AND_APPROXIMATE:
			uncertain.AddFlag(edtf.YEAR)
			approximate.AddFlag(edtf.YEAR)
		default:
			// pass
		}

	}

	if yyyy_q != nil && yyyy_q.Type == "Individual" {

		switch yyyy_q.Value {
		case edtf.UNCERTAIN:
			uncertain.AddFlag(edtf.YEAR)
		case edtf.APPROXIMATE:
			approximate.AddFlag(edtf.YEAR)
		case edtf.UNCERTAIN_AND_APPROXIMATE:
			uncertain.AddFlag(edtf.YEAR)
			approximate.AddFlag(edtf.YEAR)
		default:
			// pass
		}
	}

	if mm_q != nil && mm_q.Type == "Individual" {

		switch mm_q.Value {
		case edtf.UNCERTAIN:
			uncertain.AddFlag(edtf.MONTH)
		case edtf.APPROXIMATE:
			approximate.AddFlag(edtf.MONTH)
		case edtf.UNCERTAIN_AND_APPROXIMATE:
			uncertain.AddFlag(edtf.MONTH)
			approximate.AddFlag(edtf.MONTH)
		default:
			// pass
		}
	}

	if dd_q != nil && dd_q.Type == "Individual" {

		switch dd_q.Value {
		case edtf.UNCERTAIN:
			uncertain.AddFlag(edtf.DAY)
		case edtf.APPROXIMATE:
			approximate.AddFlag(edtf.DAY)
		case edtf.UNCERTAIN_AND_APPROXIMATE:
			uncertain.AddFlag(edtf.DAY)
			approximate.AddFlag(edtf.DAY)
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

	// fmt.Println("PRECISION 0", edtf_str, precision)

	if !strings.HasSuffix(yyyy, "X") {

		precision = edtf.NONE
		precision.AddFlag(edtf.YEAR)

	} else {

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

			precision = edtf.NONE
			precision.AddFlag(edtf.MILLENIUM)
		}

		if string(yyyy[1]) != "X" {

			c, err := strconv.ParseInt(string(yyyy[1]), 10, 32)

			if err != nil {
				return nil, err
			}

			start_c = c * 100
			end_c = start_c

			precision = edtf.NONE
			precision.AddFlag(edtf.CENTURY)
		}

		if string(yyyy[2]) != "X" {

			d, err := strconv.ParseInt(string(yyyy[2]), 10, 32)

			if err != nil {
				return nil, err
			}

			start_d = d * 10
			end_d = start_d

			precision = edtf.NONE
			precision.AddFlag(edtf.DECADE)
		}

		if string(yyyy[3]) != "X" {

			y, err := strconv.ParseInt(string(yyyy[3]), 10, 32)

			if err != nil {
				return nil, err
			}

			start_y = y * 1
			end_y = start_y

			precision = edtf.NONE
			precision.AddFlag(edtf.YEAR)
		}

		start_ymd := start_m + start_c + start_d + start_y
		end_ymd := end_m + end_c + end_d + end_y

		// fmt.Printf("OMG '%s' '%d' '%d' '%d' '%d' '%d'\n", yyyy, start_m, start_c, start_d, start_y, start_ymd)
		// fmt.Printf("WTF '%s' '%d' '%d' '%d' '%d' '%d'\n", yyyy, end_m, end_c, end_d, end_y, end_ymd)

		start_yyyy = strconv.FormatInt(start_ymd, 10)
		end_yyyy = strconv.FormatInt(end_ymd, 10)

	}

	// fmt.Println("PRECISION 1", edtf_str, precision)

	if !strings.HasSuffix(mm, "X") {

		if mm != "" && precision == edtf.NONE {
			precision = edtf.NONE
			precision.AddFlag(edtf.MONTH)
		}

	} else {

		// this does not account for 1985-24, etc.

		if strings.HasPrefix(mm, "X") {
			start_mm = "01"
			end_mm = "12"

		} else {
			start_mm = "10"
			end_mm = "12"

			precision = edtf.NONE
			precision.AddFlag(edtf.MONTH)
		}
	}

	// fmt.Println("PRECISION 2", edtf_str, precision)

	if !strings.HasSuffix(dd, "X") {

		if dd != "" && precision == edtf.NONE {
			precision = edtf.NONE
			precision.AddFlag(edtf.DAY)
		}

	} else {

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
	}

	// the fact that I need to do this tells me that all of the precision
	// logic around significant digits needs to be refactored but this will
	// do for now... (20210106/thisisaaronland)

	if dd == "XX" && mm == "XX" {
		precision = edtf.NONE
		precision.AddFlag(edtf.YEAR)
	} else if dd == "XX" {
		precision = edtf.NONE
		precision.AddFlag(edtf.MONTH)
	} else {
	}

	// fmt.Println("PRECISION 3", edtf_str, precision)

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

func EmptyDateRange() *edtf.DateRange {

	lower_d := &edtf.Date{}
	upper_d := &edtf.Date{}

	dt := &edtf.DateRange{
		Lower: lower_d,
		Upper: upper_d,
	}

	return dt
}

func UnknownDateRange() *edtf.DateRange {

	dr := EmptyDateRange()
	dr.Lower.Unknown = true
	dr.Upper.Unknown = true
	return dr
}

func OpenDateRange() *edtf.DateRange {

	dr := EmptyDateRange()
	dr.Lower.Open = true
	dr.Upper.Open = true
	return dr
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
				Type:  "Group",
				Value: m[2],
			}
		}

		return m[1], q, nil
	}

	return "", nil, edtf.Invalid("date", date)
}
