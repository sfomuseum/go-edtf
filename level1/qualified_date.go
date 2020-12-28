package level1

import (
	"errors"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/common"
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

	m := re_qualified.FindStringSubmatch(edtf_str)

	if len(m) != 5 {
		return nil, errors.New("Invalid Level 1 qualified date string")
	}

	yyyy := m[1]
	mm := m[2]
	dd := m[3]
	q := m[4]

	start, err := common.DateRangeWithYMDString(yyyy, mm, dd)

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

	switch q {
	case edtf.UNCERTAIN:
		d.Start.Lower.Uncertain = true
		d.Start.Upper.Uncertain = true
		d.End.Lower.Uncertain = true
		d.End.Upper.Uncertain = true
	case edtf.APPROXIMATE:
		d.Start.Lower.Approximate = true
		d.Start.Upper.Approximate = true
		d.End.Lower.Approximate = true
		d.End.Upper.Approximate = true
	case edtf.UNCERTAIN_AND_APPROXIMATE:
		d.Start.Lower.Uncertain = true
		d.Start.Upper.Uncertain = true
		d.End.Lower.Uncertain = true
		d.End.Upper.Uncertain = true
		d.Start.Lower.Approximate = true
		d.Start.Upper.Approximate = true
		d.End.Lower.Approximate = true
		d.End.Upper.Approximate = true
	default:
		return nil, errors.New("Invalid or unrecognized qualifier")
	}

	return d, nil
}
