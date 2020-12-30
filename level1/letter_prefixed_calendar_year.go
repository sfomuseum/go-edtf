package level1

import (
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/calendar"
	"github.com/whosonfirst/go-edtf/common"
	"github.com/whosonfirst/go-edtf/re"
)

/*

'Y' may be used at the beginning of the date string to signify that the date is a year, when (and only when) the year exceeds four digits, i.e. for years later than 9999 or earlier than -9999.

    Example 1             'Y170000002' is the year 170000002
    Example 2             'Y-170000002' is the year -170000002

*/

func IsLetterPrefixedCalendarYear(edtf_str string) bool {
	return re.LetterPrefixedCalendarYear.MatchString(edtf_str)
}

// Years must be in the range 0000..9999.
// https://golang.org/pkg/time/#Parse

// sigh....
// fmt.Printf("DEBUG %v\n", start.Add(time.Hour * 8760 * 1000))
// ./prog.go:21:54: constant 31536000000000000000 overflows time.Duration

func ParseLetterPrefixedCalendarYear(edtf_str string) (*edtf.EDTFDate, error) {

	m := re.LetterPrefixedCalendarYear.FindStringSubmatch(edtf_str)

	if len(m) != 3 {
		return nil, edtf.Invalid(LETTER_PREFIXED_CALENDAR_YEAR, edtf_str)
	}

	prefix := m[1]

	start_yyyy := m[2]
	start_mm := ""
	start_dd := ""

	if len(start_yyyy) > 4 {
		return nil, edtf.Unsupported(LETTER_PREFIXED_CALENDAR_YEAR, edtf_str)
	}

	start, err := common.DateRangeWithYMDString(start_yyyy, start_mm, start_dd)

	if err != nil {
		return nil, err
	}

	if prefix == edtf.NEGATIVE {
		start.Lower.Time = calendar.ToBCE(start.Lower.Time)
		start.Upper.Time = calendar.ToBCE(start.Upper.Time)
		start.Lower.BCE = true
		start.Upper.BCE = true
	}

	end := start

	d := &edtf.EDTFDate{
		Start: start,
		End:   end,
		EDTF:  edtf_str,
		Level: LEVEL,
	}

	return d, nil
}
