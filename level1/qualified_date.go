package level1

import (
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/common"
	"github.com/whosonfirst/go-edtf/re"
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

	m := re.QualifiedDate.FindStringSubmatch(edtf_str)

	if len(m) != 5 {
		return nil, edtf.Invalid(QUALIFIED_DATE, edtf_str)
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
		d.Start.Lower.Uncertain = edtf.DAILY
		d.Start.Upper.Uncertain = edtf.DAILY
		d.End.Lower.Uncertain = edtf.DAILY
		d.End.Upper.Uncertain = edtf.DAILY
	case edtf.APPROXIMATE:
		d.Start.Lower.Approximate = edtf.DAILY
		d.Start.Upper.Approximate = edtf.DAILY
		d.End.Lower.Approximate = edtf.DAILY
		d.End.Upper.Approximate = edtf.DAILY
	case edtf.UNCERTAIN_AND_APPROXIMATE:
		d.Start.Lower.Uncertain = edtf.DAILY
		d.Start.Upper.Uncertain = edtf.DAILY
		d.End.Lower.Uncertain = edtf.DAILY
		d.End.Upper.Uncertain = edtf.DAILY
		d.Start.Lower.Approximate = edtf.DAILY
		d.Start.Upper.Approximate = edtf.DAILY
		d.End.Lower.Approximate = edtf.DAILY
		d.End.Upper.Approximate = edtf.DAILY
	default:
		return nil, edtf.Invalid(QUALIFIED_DATE, edtf_str)
	}

	return d, nil
}
