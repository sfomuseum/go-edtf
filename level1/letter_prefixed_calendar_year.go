package level1

import (
	"errors"
	"github.com/whosonfirst/go-edtf"
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

	return nil, errors.New("Not implemented")

}
