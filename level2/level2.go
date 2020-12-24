package level2

import (
	"errors"
	"github.com/whosonfirst/go-edtf"
	"regexp"
	"strings"
)

const LEVEL int = 2

const PATTERN_EXPONENTIAL_YEAR string = `^(?i)Y(\-?\d+)E(\d+)$`
const PATTERN_SIGNIFICANT_DIGITS string = `^(?:(\d{4})S(\d+)|Y(\d+)S(\d+)|Y(\d+)E(\d+)S(\d+))$`
const PATTERN_SUB_YEAR string = `^(\d{4})\-(2[1-9]|3[0-9]|4[0-1])$`
const PATTERN_SET_REPRESENTATIONS string = `^(\[|\{)(\.\.)?(?:(?:(\d{4})(?:-(\d{2})(?:-(\d{2}))?)?)(,|\.\.)?)+(\.\.)?(\}|\])$`
const PATTERN_GROUP_QUALIFICATION string = `^(?:(\d{4})(%|~|\?)?(?:-(\d{2})(%|~|\?)?(?:-(\d{2})(%|~|\?)?)?)?)$`
const PATTERN_INDIVIDUAL_QUALIFICATION string = `^(?:(%|~|\?)?(\d{4})(?:-(%|~|\?)?(\d{2})(?:-(%|~|\?)?(\d{2}))?)?)$`
const PATTERN_UNSPECIFIED_DIGIT string = `^([0-9X]{4})(?:-([0-9X]{2})(?:-([0-9X]{2}))?)?$`
const PATTERN_INTERVAL string = `^(%|~|\?)?([0-9X]{4})(?:-(%|~|\?)?([0-9X]{2})(?:-(%|~|\?)?([0-9X]{2}))?)?\/(%|~|\?)?([0-9X]{4})(?:-(%|~|\?)?([0-9X]{2})(?:-(%|~|\?)?([0-9X]{2}))?)?$`

var re_exponential_year *regexp.Regexp
var re_significant_digits *regexp.Regexp
var re_sub_year *regexp.Regexp
var re_set_representations *regexp.Regexp
var re_group_qualification *regexp.Regexp
var re_individual_qualification *regexp.Regexp
var re_unspecified_digit *regexp.Regexp
var re_interval *regexp.Regexp
var re_level2 *regexp.Regexp

func init() {

	re_exponential_year = regexp.MustCompile(PATTERN_EXPONENTIAL_YEAR)

	re_significant_digits = regexp.MustCompile(PATTERN_SIGNIFICANT_DIGITS)

	re_sub_year = regexp.MustCompile(PATTERN_SUB_YEAR)

	re_set_representations = regexp.MustCompile(PATTERN_SET_REPRESENTATIONS)

	re_group_qualification = regexp.MustCompile(PATTERN_GROUP_QUALIFICATION)

	re_individual_qualification = regexp.MustCompile(PATTERN_INDIVIDUAL_QUALIFICATION)

	re_unspecified_digit = regexp.MustCompile(PATTERN_UNSPECIFIED_DIGIT)

	re_interval = regexp.MustCompile(PATTERN_INTERVAL)

	level2_patterns := []string{
		PATTERN_EXPONENTIAL_YEAR,
		PATTERN_SIGNIFICANT_DIGITS,
		PATTERN_SUB_YEAR,
		PATTERN_SET_REPRESENTATIONS,
		PATTERN_GROUP_QUALIFICATION,
		PATTERN_INDIVIDUAL_QUALIFICATION,
		PATTERN_UNSPECIFIED_DIGIT,
		PATTERN_INTERVAL,
	}

	re_level2 = regexp.MustCompile(`(` + strings.Join(level2_patterns, "|") + `)`)

}

func IsLevel2(edtf_str string) bool {
	return re_level2.MatchString(edtf_str)
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

Exponential year

'Y' at the beginning of the string (which indicates "year", as in level 1) may be followed by an integer, followed by 'E' followed by a positive integer. This signifies "times 10 to the power of". Thus 17E8 means "17 times 10 to the eighth power".

    Example        ‘Y-17E7’
    the calendar year -17*10 to the seventh power= -170000000

*/

func IsExponentialYear(edtf_str string) bool {
	return re_exponential_year.MatchString(edtf_str)
}

func ParseExponentialYear(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_exponential_year.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 2 exponential year string")
	}

	return nil, nil
}

/*

Significant digits

A year (expressed in any of the three allowable forms: four-digit, 'Y' prefix, or exponential) may be followed by 'S', followed by a positive integer indicating the number of significant digits.

    Example 1      ‘1950S2’
    some year between 1900 and 1999, estimated to be 1950
    Example 2      ‘Y171010000S3’
    some year between 171010000 and 171010999, estimated to be 171010000
    Example 3       ‘Y3388E2S3’
    some year between 338000 and 338999, estimated to be 338800.

*/

func IsSignificantDigits(edtf_str string) bool {
	return re_significant_digits.MatchString(edtf_str)
}

func ParseSignificantDigits(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_significant_digits.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 2 significant digits string")
	}

	return nil, nil
}

/*

Level 2 extends the season feature of Level 1 to include the following sub-year groupings.

21     Spring (independent of location)
22     Summer (independent of location)
23     Autumn (independent of location)
24     Winter (independent of location)
25     Spring - Northern Hemisphere
26     Summer - Northern Hemisphere
27     Autumn - Northern Hemisphere
28     Winter - Northern Hemisphere
29     Spring - Southern Hemisphere
30     Summer - Southern Hemisphere
31     Autumn - Southern Hemisphere
32     Winter - Southern Hemisphere
33     Quarter 1 (3 months in duration)
34     Quarter 2 (3 months in duration)
35     Quarter 3 (3 months in duration)
36     Quarter 4 (3 months in duration)
37     Quadrimester 1 (4 months in duration)
38     Quadrimester 2 (4 months in duration)
39     Quadrimester 3 (4 months in duration)
40     Semestral 1 (6 months in duration)
41     Semestral 2 (6 months in duration)

    Example        ‘2001-34’
    second quarter of 2001

*/

func IsSubYearGrouping(edtf_str string) bool {
	return re_sub_year.MatchString(edtf_str)
}

func ParseSubYearGroupings(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_sub_year.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 2 sub year groupings string")
	}

	return nil, nil
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
	return re_set_representations.MatchString(edtf_str)
}

func ParseSetRepresentations(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_set_representations.MatchString(edtf_str) {
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
	return re_group_qualification.MatchString(edtf_str)
}

func ParseGroupQualification(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_group_qualification.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 2 group qualification string")
	}

	return nil, nil
}

func IsIndividualQualification(edtf_str string) bool {
	return re_individual_qualification.MatchString(edtf_str)
}

func ParseIndividualQualification(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_individual_qualification.MatchString(edtf_str) {
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
	return re_unspecified_digit.MatchString(edtf_str)
}

func ParseUnspecifiedDigit(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_unspecified_digit.MatchString(edtf_str) {
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
	return re_interval.MatchString(edtf_str)
}

func ParseInterval(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_interval.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 2 interval string")
	}

	return nil, nil
}
