package level1

import (
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/common"
	"github.com/whosonfirst/go-edtf/re"
)

/*

 Negative calendar year

    Example 1       ‘-1985’

Note: ISO 8601 Part 1 does not support negative year.

*/

func IsNegativeCalendarYear(edtf_str string) bool {
	return re.NegativeYear.MatchString(edtf_str)
}

func ParseNegativeCalendarYear(edtf_str string) (*edtf.EDTFDate, error) {

	m := re.NegativeYear.FindStringSubmatch(edtf_str)

	if len(m) != 2 {
		return nil, edtf.Invalid(NEGATIVE_CALENDAR_YEAR, edtf_str)
	}

	start_yyyy := m[1]
	start_mm := ""
	start_dd := ""

	start, err := common.DateRangeWithYMDString(start_yyyy, start_mm, start_dd)

	if err != nil {
		return nil, err
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
