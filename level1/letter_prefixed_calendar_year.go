package level1

import (
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/common"
	"github.com/sfomuseum/go-edtf/re"
	"strings"
)

/*

'Y' may be used at the beginning of the date string to signify that the date is a year, when (and only when) the year exceeds four digits, i.e. for years later than 9999 or earlier than -9999.

    Example 1             'Y170000002' is the year 170000002
    Example 2             'Y-170000002' is the year -170000002

*/

func IsLetterPrefixedCalendarYear(edtf_str string) bool {
	return re.LetterPrefixedCalendarYear.MatchString(edtf_str)
}

func ParseLetterPrefixedCalendarYear(edtf_str string) (*edtf.EDTFDate, error) {

	m := re.LetterPrefixedCalendarYear.FindStringSubmatch(edtf_str)

	if len(m) != 2 {
		return nil, edtf.Invalid(LETTER_PREFIXED_CALENDAR_YEAR, edtf_str)
	}

	// Years must be in the range 0000..9999.
	// https://golang.org/pkg/time/#Parse

	// sigh....
	// fmt.Printf("DEBUG %v\n", start.Add(time.Hour * 8760 * 1000))
	// ./prog.go:21:54: constant 31536000000000000000 overflows time.Duration

	// common.DateSpanFromEDTF needs to be updated to simply assign a valid
	// *edtf.YMD element and leave *time.Time blank when creating *edtf.Date
	// instances (20210105/thisisaaronland)

	yyyy := m[1]

	max_length := 4

	if strings.HasPrefix(yyyy, "-") {
		max_length = 5
	}

	if len(yyyy) > max_length {
		return nil, edtf.Unsupported(LETTER_PREFIXED_CALENDAR_YEAR, edtf_str)
	}

	sp, err := common.DateSpanFromEDTF(yyyy)

	if err != nil {
		return nil, err
	}

	d := &edtf.EDTFDate{
		Start:   sp.Start,
		End:     sp.End,
		EDTF:    edtf_str,
		Level:   LEVEL,
		Feature: LETTER_PREFIXED_CALENDAR_YEAR,
	}

	return d, nil
}
