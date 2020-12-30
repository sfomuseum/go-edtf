package level2

import (
	"errors"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/re"
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

var Tests map[string][]string = map[string][]string{
	EXPONENTIAL_YEAR: []string{
		"Y-17E7", // TO DO - https://github.com/whosonfirst/go-edtf/issues/5
		"Y10E7",  // TO DO
		"Y20E2",
	},
	SIGNIFICANT_DIGITS: []string{
		"1950S2",
		"Y171010000S3",
		"Y-20E2S3",
		"Y3388E2S3",
		"Y-20E2S3",
	},
	SUB_YEAR_GROUPINGS: []string{
		"2001-34",
		// "second quarter of 2001",	// TO DO
	},
	SET_REPRESENTATIONS: []string{
		"[1667,1668,1670..1672]",
		"[..1760-12-03]",
		"[1760-12..]",
		"[1760-01,1760-02,1760-12..]",
		"[1667,1760-12]",
		"[..1984]",
		"{1667,1668,1670..1672}",
		"{1960,1961-12}",
		"{..1984}",
	},
	GROUP_QUALIFICATION: []string{
		"2004-06-11%",
		"2004-06~-11",
		"2004?-06-11",
	},
	INDIVIDUAL_QUALIFICATION: []string{
		"?2004-06-~11",
		"2004-%06-11",
	},
	UNSPECIFIED_DIGIT: []string{
		"156X-12-25",
		"15XX-12-25",
		// "XXXX-12-XX",	// TO DO
		"1XXX-XX",
		"1XXX-12",
		"1984-1X",
	},
	INTERVAL: []string{
		"2004-06-~01/2004-06-~20",
		"2004-06-XX/2004-07-03",
	},
}

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
