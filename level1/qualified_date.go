package level1

import (
	"errors"
	"github.com/whosonfirst/go-edtf"
)

/*

Qualification of a date (complete)

The characters '?', '~' and '%' are used to mean "uncertain", "approximate", and "uncertain" as well as "approximate", respectively. These characters may occur only at the end of the date string and apply to the entire date.

    Example 1             '1984?'             year uncertain (possibly the year 1984, but not definitely)
    Example 2              '2004-06~''       year-month approximate
    Example 3        '2004-06-11%'          entire date (year-month-day) uncertain and approximate

*/

func IsQualifiedDate(edtf_str string) bool {
	return re_qualified.MatchString(edtf_str)
}

func ParseQualifiedDate(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_qualified.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 1 qualified date string")
	}

	return nil, nil
}
