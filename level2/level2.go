package level2

import (
	"errors"
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/re"
)

const LEVEL int = 2

const EXPONENTIAL_YEAR string = "exponential year"
const SIGNIFICANT_DIGITS string = "significant digits"
const SUB_YEAR_GROUPINGS string = "sub year groupings"
const SET_REPRESENTATIONS string = "set representations"
const GROUP_QUALIFICATION string = "group qualification"
const INDIVIDUAL_QUALIFICATION string = "individual qualification"
const UNSPECIFIED_DIGIT string = "unspecified digit"
const INTERVAL string = "interval"

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

	return "", errors.New("Invalid or unsupported Level 2 string")
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

	return nil, errors.New("Invalid or unsupported Level 2 string")
}
