package level1

import (
	"errors"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/common"
)

/*

 Negative calendar year

    Example 1       ‘-1985’

Note: ISO 8601 Part 1 does not support negative year.

*/

func IsNegativeCalendarYear(edtf_str string) bool {
	return re_negative_year.MatchString(edtf_str)
}

func ParseNegativeCalendarYear(edtf_str string) (*edtf.EDTFDate, error) {

	m := re_negative_year.FindStringSubmatch(edtf_str)

	if len(m) != 2 {
		return nil, errors.New("Invalid Level 1 negative year string")
	}

	start_yyyy := m[1]
	start_mm := ""
	start_dd := ""

	start, err := common.DateRangeWithYMDString(start_yyyy, start_mm, start_dd)

	if err != nil {
		return nil, err
	}

	start.Lower.BCE = true
	start.Upper.BCE = true

	end := start

	d := &edtf.EDTFDate{
		Start: start,
		End:   end,
		EDTF:  edtf_str,
		Level: LEVEL,
	}

	return d, nil
}
