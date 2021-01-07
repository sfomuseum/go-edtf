package level1

import (
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/common"
	"github.com/sfomuseum/go-edtf/re"
)

/*

Qualification of a date (complete)

The characters '?', '~' and '%' are used to mean "uncertain", "approximate", and "uncertain" as well as "approximate", respectively. These characters may occur only at the end of the date string and apply to the entire date.

    Example 1             '1984?'             year uncertain (possibly the year 1984, but not definitely)
    Example 2              '2004-06~''       year-month approximate
    Example 3        '2004-06-11%'          entire date (year-month-day) uncertain and approximate

*/

func IsQualifiedDate(edtf_str string) bool {
	return re.QualifiedDate.MatchString(edtf_str)
}

func ParseQualifiedDate(edtf_str string) (*edtf.EDTFDate, error) {

	if !re.QualifiedDate.MatchString(edtf_str) {
		return nil, edtf.Invalid(QUALIFIED_DATE, edtf_str)
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
		Feature: QUALIFIED_DATE,
	}

	return d, nil
}
