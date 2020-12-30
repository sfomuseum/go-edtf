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
		// https://github.com/whosonfirst/go-edtf/issues/5
		// "Y-17E7",
		// "Y10E7",
		"Y2E3",
	},
	SIGNIFICANT_DIGITS: []string{
		"1950S2",
		"Y171010000S3",
		"Y-1S3",
		"Y3388E2S3",
		"Y-20E2S3",
	},
	SUB_YEAR_GROUPINGS: []string{
		"2001-34",
		// TO DO
		// "second quarter of 2001"
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
		// TO DO
		// "XXXX-12-XX",
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

/*

Set representation

    Square brackets wrap a single-choice list (select one member).
    Curly brackets wrap an inclusive list (all members included).
    Members of the set are separated by commas.
    No spaces are allowed, anywhere within the expression.
    Double-dots indicates all the values between the two values it separates, inclusive.
    Double-dot at the beginning or end of the list means "on or before" or "on or after" respectively.
    Elements immediately preceeding and/or following as well as the elements represented by a double-dot, all have the same precision. Otherwise, different elements may have different precisions

One of a set

    Example 1       [1667,1668,1670..1672]
    One of the years 1667, 1668, 1670, 1671, 1672
    Example 2         [..1760-12-03]
    December 3, 1760; or some earlier date
    Example 3          [1760-12..]
    December 1760, or some later month
    Example 4         [1760-01,1760-02,1760-12..]
    January or February of 1760 or December 1760 or some later month
    Example 5          [1667,1760-12]
    Either the year 1667 or the month December of 1760.
    Example 6         [..1984]
    The year 1984 or an earlier year

All Members

    Example 7          {1667,1668,1670..1672}
    All of the years 1667, 1668, 1670, 1671, 1672
    Example 8            {1960,1961-12}
    The year 1960 and the month December of 1961.
    Example 9         {..1984}
    The year 1984 and all earlier years

*/

func IsSetRepresentation(edtf_str string) bool {
	return re.SetRepresentations.MatchString(edtf_str)
}

func ParseSetRepresentations(edtf_str string) (*edtf.EDTFDate, error) {

	if !re.SetRepresentations.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 2 set representation string")
	}

	return nil, nil
}

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
		return nil, errors.New("Invalid Level 2 group qualification string")
	}

	return nil, nil
}

func IsIndividualQualification(edtf_str string) bool {
	return re.IndividualQualification.MatchString(edtf_str)
}

func ParseIndividualQualification(edtf_str string) (*edtf.EDTFDate, error) {

	if !re.IndividualQualification.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 2 individual qualification string")
	}

	return nil, nil
}

/*

Unspecified Digit

For level 2 the unspecified digit, 'X', may occur anywhere within a component.

    Example 1                 ‘156X-12-25’
    December 25 sometime during the 1560s
    Example 2                 ‘15XX-12-25’
    December 25 sometime during the 1500s
    Example 3                ‘XXXX-12-XX’
    Some day in December in some year
    Example 4                 '1XXX-XX’
    Some month during the 1000s
    Example 5                  ‘1XXX-12’
    Some December during the 1000s
    Example 6                  ‘1984-1X’
    October, November, or December 1984

*/

func IsUnspecifiedDigit(edtf_str string) bool {
	return re.UnspecifiedDigit.MatchString(edtf_str)
}

func ParseUnspecifiedDigit(edtf_str string) (*edtf.EDTFDate, error) {

	if !re.UnspecifiedDigit.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 2 unspecified digit string")
	}

	return nil, nil
}

/*

For Level 2 portions of a date within an interval may be designated as approximate, uncertain, or unspecified.

    Example 1                 ‘2004-06-~01/2004-06-~20’
    An interval in June 2004 beginning approximately the first and ending approximately the 20th
    Example 2                 ‘2004-06-XX/2004-07-03’
    An interval beginning on an unspecified day in June 2004 and ending July 3.


*/

func IsInterval(edtf_str string) bool {
	return re.Interval.MatchString(edtf_str)
}

func ParseInterval(edtf_str string) (*edtf.EDTFDate, error) {

	if !re.Interval.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 2 interval string")
	}

	return nil, nil
}
