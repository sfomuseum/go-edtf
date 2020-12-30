package level2

import (
	_ "fmt"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/re"
	_ "strings"
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

Qualification of Individual Component

A qualification character to the immediate left of a component applies to that component only.

    Example 4                   ‘?2004-06-~11’
    year uncertain; month known; day approximate
    Example 5                   ‘2004-%06-11’
    month uncertain and approximate; year and day known

*/

func IsGroupQualification(edtf_str string) bool {
	return re.GroupQualification.MatchString(edtf_str)
}

func ParseGroupQualification(edtf_str string) (*edtf.EDTFDate, error) {

	if !re.GroupQualification.MatchString(edtf_str) {
		return nil, edtf.Invalid(GROUP_QUALIFICATION, edtf_str)
	}

	return nil, edtf.NotImplemented(GROUP_QUALIFICATION, edtf_str)
}

func IsIndividualQualification(edtf_str string) bool {
	return re.IndividualQualification.MatchString(edtf_str)
}

func ParseIndividualQualification(edtf_str string) (*edtf.EDTFDate, error) {

	if !re.IndividualQualification.MatchString(edtf_str) {
		return nil, edtf.Invalid(INDIVIDUAL_QUALIFICATION, edtf_str)
	}

	return nil, edtf.NotImplemented(INDIVIDUAL_QUALIFICATION, edtf_str)
}
