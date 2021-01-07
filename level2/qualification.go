package level2

import (
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/common"
	"github.com/sfomuseum/go-edtf/re"
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

	sp, err := common.DateSpanFromEDTF(edtf_str)

	if err != nil {
		return nil, err
	}

	d := &edtf.EDTFDate{
		Start:   sp.Start,
		End:     sp.End,
		EDTF:    edtf_str,
		Level:   LEVEL,
		Feature: GROUP_QUALIFICATION,
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

	if !re.IndividualQualification.MatchString(edtf_str) {
		return nil, edtf.Invalid(INDIVIDUAL_QUALIFICATION, edtf_str)
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
		Feature: INDIVIDUAL_QUALIFICATION,
	}

	return d, nil
}
