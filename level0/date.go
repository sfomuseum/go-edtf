package level0

import (
	"errors"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/common"
)

/*

Date

    complete representation:            [year][“-”][month][“-”][day]
    Example 1          ‘1985-04-12’ refers to the calendar date 1985 April 12th with day precision.
    reduced precision for year and month:   [year][“-”][month]
    Example 2          ‘1985-04’ refers to the calendar month 1985 April with month precision.
    reduced precision for year:  [year]
    Example 3          ‘1985’ refers to the calendar year 1985 with year precision.

*/

func IsDate(edtf_str string) bool {
	return re_date.MatchString(edtf_str)
}

func ParseDate(edtf_str string) (*edtf.EDTFDate, error) {

	m := re_date.FindStringSubmatch(edtf_str)

	if len(m) != 4 {
		return nil, errors.New("Invalid Level 0 date string")
	}

	yyyy := m[1]
	mm := m[2]
	dd := m[3]

	start, err := common.DateRangeWithYMDString(yyyy, mm, dd)

	if err != nil {
		return nil, err
	}

	end := start

	if err != nil {
		return nil, err
	}

	d := &edtf.EDTFDate{
		Start: start,
		End:   end,
		EDTF:  edtf_str,
		Level: LEVEL,
	}

	return d, nil
}
