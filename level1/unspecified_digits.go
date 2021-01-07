package level1

import (
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/common"
	"github.com/sfomuseum/go-edtf/re"
)

/*

Unspecified digit(s) from the right

The character 'X' may be used in place of one or more rightmost digits to indicate that the value of that digit is unspecified, for the following cases:

    A year with one or two (rightmost) unspecified digits in a year-only expression (year precision)
    Example 1       ‘201X’
    Example 2       ‘20XX’
    Year specified, month unspecified in a year-month expression (month precision)
    Example 3       ‘2004-XX’
    Year and month specified, day unspecified in a year-month-day expression (day precision)
    Example 4       ‘1985-04-XX’
    Year specified, day and month unspecified in a year-month-day expression  (day precision)
    Example 5       ‘1985-XX-XX’


*/

func IsUnspecifiedDigits(edtf_str string) bool {
	return re.UnspecifiedDigits.MatchString(edtf_str)
}

func ParseUnspecifiedDigits(edtf_str string) (*edtf.EDTFDate, error) {

	if !re.UnspecifiedDigits.MatchString(edtf_str) {
		return nil, edtf.Invalid(UNSPECIFIED_DIGITS, edtf_str)
	}

	sp, err := common.DateSpanFromEDTF(edtf_str)

	if err != nil {
		return nil, err
	}

	d := &edtf.EDTFDate{
		Start:   sp.Start,
		End:     sp.End,
		EDTF:    edtf_str,
		Level:   LEVEL,
		Feature: UNSPECIFIED_DIGITS,
	}

	return d, nil
}
