package level2

import (
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/common"
	"github.com/whosonfirst/go-edtf/re"
)

/*

Group Qualification

A qualification character to the immediate right of a component applies to that component as well as to all components to the left.

    Example 1                ‘2004-06-11%’
    year, month, and day uncertain and approximate
    Example 2                 ‘2004-06~-11’
    year and month approximate
    Example  3              ‘2004?-06-11’
    year uncertain
*/

func IsGroupQualification(edtf_str string) bool {
	return re.GroupQualification.MatchString(edtf_str)
}

func ParseGroupQualification(edtf_str string) (*edtf.EDTFDate, error) {

	/*

		GROUP 2004-06-11% 7 2004-06-11%,2004,,06,,11,%
		GROUP 2004-06~-11 7 2004-06~-11,2004,,06,~,11,
		GROUP 2004?-06-11 7 2004?-06-11,2004,?,06,,11,

	*/

	if !re.GroupQualification.MatchString(edtf_str) {
		return nil, edtf.Invalid(GROUP_QUALIFICATION, edtf_str)
	}

	m := re.GroupQualification.FindStringSubmatch(edtf_str)

	if len(m) != 7 {
		return nil, edtf.Invalid(GROUP_QUALIFICATION, edtf_str)
	}

	yyyy := m[1]
	yyyy_q := m[2]

	mm := m[3]
	mm_q := m[4]

	dd := m[5]
	dd_q := m[6]

	precision := edtf.NONE
	q := ""

	if yyyy_q != "" {

		q = yyyy_q

		precision.AddFlag(edtf.ANNUAL)
	}

	if mm_q != "" {

		yyyy_q = mm_q
		q = mm_q

		precision.AddFlag(edtf.ANNUAL)
		precision.AddFlag(edtf.MONTHLY)
	}

	if dd_q != "" {

		mm_q = dd_q
		yyyy_q = dd_q
		q = dd_q

		precision.AddFlag(edtf.ANNUAL)
		precision.AddFlag(edtf.MONTHLY)
		precision.AddFlag(edtf.DAILY)
	}

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
		// pass
	}

	return d, nil
}

/*

Qualification of Individual Component

A qualification character to the immediate left of a component applies to that component only.

    Example 4                   ‘?2004-06-~11’
    year uncertain; month known; day approximate
    Example 5                   ‘2004-%06-11’
    month uncertain and approximate; year and day known

*/

func IsIndividualQualification(edtf_str string) bool {
	return re.IndividualQualification.MatchString(edtf_str)
}

func ParseIndividualQualification(edtf_str string) (*edtf.EDTFDate, error) {

	/*

		INDIVIDUAL ?2004-06-~11 7 ?2004-06-~11,?,2004,,06,~,11
		INDIVIDUAL 2004-%06-11 7 2004-%06-11,,2004,%,06,,11

	*/

	m := re.IndividualQualification.FindStringSubmatch(edtf_str)

	if len(m) != 7 {
		return nil, edtf.Invalid(INDIVIDUAL_QUALIFICATION, edtf_str)
	}

	yyyy := m[2]
	yyyy_q := m[1]

	mm := m[4]
	mm_q := m[3]

	dd := m[6]
	dd_q := m[5]

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

	uncertain := edtf.NONE
	approximate := edtf.NONE

	switch yyyy_q {
	case edtf.UNCERTAIN:
		uncertain.AddFlag(edtf.ANNUAL)
	case edtf.APPROXIMATE:
		approximate.AddFlag(edtf.ANNUAL)
	case edtf.UNCERTAIN_AND_APPROXIMATE:
		uncertain.AddFlag(edtf.ANNUAL)
		approximate.AddFlag(edtf.ANNUAL)
	default:
		// pass
	}

	switch mm_q {
	case edtf.UNCERTAIN:
		uncertain.AddFlag(edtf.MONTHLY)
	case edtf.APPROXIMATE:
		approximate.AddFlag(edtf.MONTHLY)
	case edtf.UNCERTAIN_AND_APPROXIMATE:
		uncertain.AddFlag(edtf.MONTHLY)
		approximate.AddFlag(edtf.MONTHLY)
	default:
		// pass
	}

	switch dd_q {
	case edtf.UNCERTAIN:
		uncertain.AddFlag(edtf.DAILY)
	case edtf.APPROXIMATE:
		approximate.AddFlag(edtf.DAILY)
	case edtf.UNCERTAIN_AND_APPROXIMATE:
		uncertain.AddFlag(edtf.DAILY)
		approximate.AddFlag(edtf.DAILY)
	default:
		// pass
	}

	if uncertain != edtf.NONE {
		d.Start.Lower.Uncertain = uncertain
		d.Start.Upper.Uncertain = uncertain
		d.End.Lower.Uncertain = uncertain
		d.End.Upper.Uncertain = uncertain
	}

	if approximate != edtf.NONE {
		d.Start.Lower.Approximate = approximate
		d.Start.Upper.Approximate = approximate
		d.End.Lower.Approximate = approximate
		d.End.Upper.Approximate = approximate
	}

	return d, nil
}
