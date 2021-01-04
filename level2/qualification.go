package level2

import (
	"fmt"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/common"
	"github.com/whosonfirst/go-edtf/re"
	"strings"
)

/*

Qualification

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

	if dd_q != "" {
		mm_q = dd_q
		yyyy_q = dd_q
	}

	if mm_q != "" {
		yyyy_q = mm_q
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

	switch yyyy_q {
	case edtf.UNCERTAIN:
		d.Start.Lower.Uncertain = edtf.ANNUAL
		d.Start.Upper.Uncertain = edtf.ANNUAL
		d.End.Lower.Uncertain = edtf.ANNUAL
		d.End.Upper.Uncertain = edtf.ANNUAL
	case edtf.APPROXIMATE:
		d.Start.Lower.Approximate = edtf.ANNUAL
		d.Start.Upper.Approximate = edtf.ANNUAL
		d.End.Lower.Approximate = edtf.ANNUAL
		d.End.Upper.Approximate = edtf.ANNUAL
	case edtf.UNCERTAIN_AND_APPROXIMATE:
		d.Start.Lower.Uncertain = edtf.ANNUAL
		d.Start.Upper.Uncertain = edtf.ANNUAL
		d.End.Lower.Uncertain = edtf.ANNUAL
		d.End.Upper.Uncertain = edtf.ANNUAL
		d.Start.Lower.Approximate = edtf.ANNUAL
		d.Start.Upper.Approximate = edtf.ANNUAL
		d.End.Lower.Approximate = edtf.ANNUAL
		d.End.Upper.Approximate = edtf.ANNUAL
	default:
		// pass
	}

	switch mm_q {
	case edtf.UNCERTAIN:
		d.Start.Lower.Uncertain = edtf.MONTHLY
		d.Start.Upper.Uncertain = edtf.MONTHLY
		d.End.Lower.Uncertain = edtf.MONTHLY
		d.End.Upper.Uncertain = edtf.MONTHLY
	case edtf.APPROXIMATE:
		d.Start.Lower.Approximate = edtf.MONTHLY
		d.Start.Upper.Approximate = edtf.MONTHLY
		d.End.Lower.Approximate = edtf.MONTHLY
		d.End.Upper.Approximate = edtf.MONTHLY
	case edtf.UNCERTAIN_AND_APPROXIMATE:
		d.Start.Lower.Uncertain = edtf.MONTHLY
		d.Start.Upper.Uncertain = edtf.MONTHLY
		d.End.Lower.Uncertain = edtf.MONTHLY
		d.End.Upper.Uncertain = edtf.MONTHLY
		d.Start.Lower.Approximate = edtf.MONTHLY
		d.Start.Upper.Approximate = edtf.MONTHLY
		d.End.Lower.Approximate = edtf.MONTHLY
		d.End.Upper.Approximate = edtf.MONTHLY
	default:
		// pass
	}

	switch dd_q {
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

	fmt.Println("INDIVIDUAL", edtf_str, len(m), strings.Join(m, ","))

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

	// FIX ME : account for multiple values...

	switch yyyy_q {
	case edtf.UNCERTAIN:
		d.Start.Lower.Uncertain = edtf.ANNUAL
		d.Start.Upper.Uncertain = edtf.ANNUAL
		d.End.Lower.Uncertain = edtf.ANNUAL
		d.End.Upper.Uncertain = edtf.ANNUAL
	case edtf.APPROXIMATE:
		d.Start.Lower.Approximate = edtf.ANNUAL
		d.Start.Upper.Approximate = edtf.ANNUAL
		d.End.Lower.Approximate = edtf.ANNUAL
		d.End.Upper.Approximate = edtf.ANNUAL
	case edtf.UNCERTAIN_AND_APPROXIMATE:
		d.Start.Lower.Uncertain = edtf.ANNUAL
		d.Start.Upper.Uncertain = edtf.ANNUAL
		d.End.Lower.Uncertain = edtf.ANNUAL
		d.End.Upper.Uncertain = edtf.ANNUAL
		d.Start.Lower.Approximate = edtf.ANNUAL
		d.Start.Upper.Approximate = edtf.ANNUAL
		d.End.Lower.Approximate = edtf.ANNUAL
		d.End.Upper.Approximate = edtf.ANNUAL
	default:
		// pass
	}

	switch mm_q {
	case edtf.UNCERTAIN:
		d.Start.Lower.Uncertain = edtf.MONTHLY
		d.Start.Upper.Uncertain = edtf.MONTHLY
		d.End.Lower.Uncertain = edtf.MONTHLY
		d.End.Upper.Uncertain = edtf.MONTHLY
	case edtf.APPROXIMATE:
		d.Start.Lower.Approximate = edtf.MONTHLY
		d.Start.Upper.Approximate = edtf.MONTHLY
		d.End.Lower.Approximate = edtf.MONTHLY
		d.End.Upper.Approximate = edtf.MONTHLY
	case edtf.UNCERTAIN_AND_APPROXIMATE:
		d.Start.Lower.Uncertain = edtf.MONTHLY
		d.Start.Upper.Uncertain = edtf.MONTHLY
		d.End.Lower.Uncertain = edtf.MONTHLY
		d.End.Upper.Uncertain = edtf.MONTHLY
		d.Start.Lower.Approximate = edtf.MONTHLY
		d.Start.Upper.Approximate = edtf.MONTHLY
		d.End.Lower.Approximate = edtf.MONTHLY
		d.End.Upper.Approximate = edtf.MONTHLY
	default:
		// pass
	}

	switch dd_q {
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
		// pass
	}

	return d, nil
}
