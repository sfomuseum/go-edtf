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

	if !re.QualifiedDate.MatchString(edtf_str) {
		return nil, edtf.Invalid(QUALIFIED_DATE, edtf_str)
	}

	sp, err := common.DateSpanFromEDTF(edtf_str)

	if err != nil {
		return nil, err
	}

	d := &edtf.EDTFDate{
		Start: sp.Start,
		End:   sp.End,
		EDTF:  edtf_str,
		Level: LEVEL,
	}

	return d, nil

	/*
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

		precision := edtf.NONE
		precision.AddFlag(edtf.DAILY)

		switch q {
		case edtf.UNCERTAIN:
			d.Start.Lower.Uncertain = precision
			d.Start.Upper.Uncertain = precision
			d.End.Lower.Uncertain = precision
			d.End.Upper.Uncertain = precision
		case edtf.APPROXIMATE:
			d.Start.Lower.Approximate = precision
			d.Start.Upper.Approximate = precision
			d.End.Lower.Approximate = precision
			d.End.Upper.Approximate = precision
		case edtf.UNCERTAIN_AND_APPROXIMATE:
			d.Start.Lower.Uncertain = precision
			d.Start.Upper.Uncertain = precision
			d.End.Lower.Uncertain = precision
			d.End.Upper.Uncertain = precision
			d.Start.Lower.Approximate = precision
			d.Start.Upper.Approximate = precision
			d.End.Lower.Approximate = precision
			d.End.Upper.Approximate = precision
		default:
			return nil, edtf.Invalid(QUALIFIED_DATE, edtf_str)
		}

		return d, nil
	*/
}
