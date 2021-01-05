package level0

import (
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/common"
	"github.com/whosonfirst/go-edtf/re"
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
	return re.Date.MatchString(edtf_str)
}

func ParseDate(edtf_str string) (*edtf.EDTFDate, error) {

	if !re.Date.MatchString(edtf_str) {
		return nil, edtf.Invalid(DATE, edtf_str)
	}

	start, err := common.DateRangeWithString(edtf_str)

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
