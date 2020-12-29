package level1

import (
	"errors"
	"fmt"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/common"
	"strings"
)

/*

'Y' may be used at the beginning of the date string to signify that the date is a year, when (and only when) the year exceeds four digits, i.e. for years later than 9999 or earlier than -9999.

    Example 1             'Y170000002' is the year 170000002
    Example 2             'Y-170000002' is the year -170000002

*/

func IsLetterPrefixedCalendarYear(edtf_str string) bool {
	return re_calendaryear.MatchString(edtf_str)
}

func ParseLetterPrefixedCalendarYear(edtf_str string) (*edtf.EDTFDate, error) {

	m := re_calendaryear.FindStringSubmatch(edtf_str)

	fmt.Println("CALENDAR", len(m), strings.Join(m, ","))

	if len(m) != 3 {
		return nil, errors.New("Invalid Level 1 letter prefixed calendar year string")
	}

	prefix := m[1]

	start_yyyy := m[2]
	start_mm := ""
	start_dd := ""

	start, err := common.DateRangeWithYMDString(start_yyyy, start_mm, start_dd)

	if err != nil {
		return nil, err
	}

	if prefix == edtf.NEGATIVE {
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
