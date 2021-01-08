package level2

import (
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/re"
)

const LEVEL int = 2

const EXPONENTIAL_YEAR string = "Exponential year"
const SIGNIFICANT_DIGITS string = "Significant digits"
const SUB_YEAR_GROUPINGS string = "Sub-year groupings"
const SET_REPRESENTATIONS string = "Set representation"
const GROUP_QUALIFICATION string = "Qualification (Group)"
const INDIVIDUAL_QUALIFICATION string = "Qualification (Individual)"
const UNSPECIFIED_DIGIT string = "Unspecified Digit"
const INTERVAL string = "Interval"

func IsLevel2(edtf_str string) bool {
	return re.Level2.MatchString(edtf_str)
}

func Matches(edtf_str string) (string, error) {

	if IsExponentialYear(edtf_str) {
		return EXPONENTIAL_YEAR, nil
	}

	if IsSignificantDigits(edtf_str) {
		return SIGNIFICANT_DIGITS, nil
	}

	if IsSubYearGrouping(edtf_str) {
		return SUB_YEAR_GROUPINGS, nil
	}

	if IsSetRepresentation(edtf_str) {
		return SET_REPRESENTATIONS, nil
	}

	if IsGroupQualification(edtf_str) {
		return GROUP_QUALIFICATION, nil
	}

	if IsIndividualQualification(edtf_str) {
		return INDIVIDUAL_QUALIFICATION, nil
	}

	if IsUnspecifiedDigit(edtf_str) {
		return UNSPECIFIED_DIGIT, nil
	}

	if IsInterval(edtf_str) {
		return INTERVAL, nil
	}

	return "", edtf.Invalid("Invalid or unsupported Level 2 string", edtf_str)
}

func ParseString(edtf_str string) (*edtf.EDTFDate, error) {

	if IsExponentialYear(edtf_str) {
		return ParseExponentialYear(edtf_str)
	}

	if IsSignificantDigits(edtf_str) {
		return ParseSignificantDigits(edtf_str)
	}

	if IsSubYearGrouping(edtf_str) {
		return ParseSubYearGroupings(edtf_str)
	}

	if IsSetRepresentation(edtf_str) {
		return ParseSetRepresentations(edtf_str)
	}

	if IsGroupQualification(edtf_str) {
		return ParseGroupQualification(edtf_str)
	}

	if IsIndividualQualification(edtf_str) {
		return ParseIndividualQualification(edtf_str)
	}

	if IsUnspecifiedDigit(edtf_str) {
		return ParseUnspecifiedDigit(edtf_str)
	}

	if IsInterval(edtf_str) {
		return ParseInterval(edtf_str)
	}

	return nil, edtf.Invalid("Invalid or unsupported Level 2 string", edtf_str)
}
