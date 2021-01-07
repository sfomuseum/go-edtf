package level1

import (
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/common"
	"github.com/sfomuseum/go-edtf/re"
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

	if !re.NegativeYear.MatchString(edtf_str) {
		return nil, edtf.Invalid(NEGATIVE_CALENDAR_YEAR, edtf_str)
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
		Feature: NEGATIVE_CALENDAR_YEAR,
	}

	return d, nil
}
