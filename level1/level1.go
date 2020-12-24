package level1

import (
	"errors"
	_ "fmt"
	"github.com/whosonfirst/go-edtf"
	"regexp"
	"strings"
)

const LEVEL int = 1

const PATTERN_CALENDAR_YEAR string = `^Y(\-)?(\d+)$`
const PATTERN_SEASON string = `^(\d{4})-(0[1-9]|2[1-4])|(?i)(spring|summer|fall|winter)\s*,\s*(\d{4})$`
const PATTERN_QUALIFIED string = `^(?:(\d{4})(?:-(\d{2})(?:-(\d{2}))?)?)(\?|~|%)$`
const PATTERN_UNSPECIFIED string = `^(?:(\d{3})(X)|(\d{2})(XX)|(\d{4})-(XX)|(\d{4})\-(\d{2})\-(XX)|(\d{4})\-(XX)\-(XX))$`
const PATTERN_INTERVAL_START = `^(\.\.)?\/(?:(\d{4})(?:-(\d{2})(?:-(\d{2}))?)?)$`
const PATTERN_INTERVAL_END = `^(?:(\d{4})(?:-(\d{2})(?:-(\d{2}))?)?)\/(\.\.)?$`
const PATTERN_NEGATIVE_YEAR = `^\-(\d{4})$`

var re_calendaryear *regexp.Regexp
var re_season *regexp.Regexp
var re_qualified *regexp.Regexp
var re_unspecified *regexp.Regexp
var re_interval_end *regexp.Regexp
var re_interval_start *regexp.Regexp
var re_negative_year *regexp.Regexp
var re_level1 *regexp.Regexp

func init() {

	re_calendaryear = regexp.MustCompile(PATTERN_CALENDAR_YEAR)

	re_season = regexp.MustCompile(PATTERN_SEASON)

	re_qualified = regexp.MustCompile(PATTERN_QUALIFIED)

	re_unspecified = regexp.MustCompile(PATTERN_UNSPECIFIED)

	re_interval_start = regexp.MustCompile(PATTERN_INTERVAL_START)

	re_interval_end = regexp.MustCompile(PATTERN_INTERVAL_END)

	re_negative_year = regexp.MustCompile(PATTERN_NEGATIVE_YEAR)

	level1_patterns := []string{
		PATTERN_CALENDAR_YEAR,
		PATTERN_SEASON,
		PATTERN_QUALIFIED,
		PATTERN_UNSPECIFIED,
		PATTERN_INTERVAL_START,
		PATTERN_INTERVAL_END,
		PATTERN_NEGATIVE_YEAR,
	}

	re_level1 = regexp.MustCompile(`(` + strings.Join(level1_patterns, "|") + `)`)
}

func IsLevel1(edtf_str string) bool {
	return re_level1.MatchString(edtf_str)
}

func ParseString(edtf_str string) (*edtf.EDTFDate, error) {

	if IsLetterPrefixedCalendarYear(edtf_str) {
		return ParseLetterPrefixedCalendarYear(edtf_str)
	}

	if IsSeason(edtf_str) {
		return ParseSeason(edtf_str)
	}

	if IsQualifiedDate(edtf_str) {
		return ParseQualifiedDate(edtf_str)
	}

	if IsUnspecifiedDigits(edtf_str) {
		return ParseUnspecifiedDigits(edtf_str)
	}

	if IsNegativeCalendarYear(edtf_str) {
		return ParseNegativeCalendarYear(edtf_str)
	}

	if IsExtendedInterval(edtf_str) {
		return ParseExtendedInterval(edtf_str)
	}

	return nil, errors.New("Invalid or unsupported Level 1 EDTF string")
}

/*

'Y' may be used at the beginning of the date string to signify that the date is a year, when (and only when) the year exceeds four digits, i.e. for years later than 9999 or earlier than -9999.

    Example 1             'Y170000002' is the year 170000002
    Example 2             'Y-170000002' is the year -170000002

*/

func IsLetterPrefixedCalendarYear(edtf_str string) bool {
	return re_calendaryear.MatchString(edtf_str)
}

func ParseLetterPrefixedCalendarYear(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_calendaryear.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 1 calendar year string")
	}

	return nil, nil
}

/*

Seasons

The values 21, 22, 23, 24 may be used used to signify ' Spring', 'Summer', 'Autumn', 'Winter', respectively, in place of a month value (01 through 12) for a year-and-month format string.

    Example                   2001-21     Spring, 2001

*/

func IsSeason(edtf_str string) bool {
	return re_season.MatchString(edtf_str)
}

func ParseSeason(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_season.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 1 season string")
	}

	return nil, nil
}

/*

Qualification of a date (complete)

The characters '?', '~' and '%' are used to mean "uncertain", "approximate", and "uncertain" as well as "approximate", respectively. These characters may occur only at the end of the date string and apply to the entire date.

    Example 1             '1984?'             year uncertain (possibly the year 1984, but not definitely)
    Example 2              '2004-06~''       year-month approximate
    Example 3        '2004-06-11%'          entire date (year-month-day) uncertain and approximate

*/

func IsQualifiedDate(edtf_str string) bool {
	return re_qualified.MatchString(edtf_str)
}

func ParseQualifiedDate(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_qualified.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 1 qualified date string")
	}

	return nil, nil
}

/*

Unspecified digit(s) from the right

The character 'X' may be used in place of one or more rightmost digits to indicate that the value of that digit is unspecified, for the following cases:

    A year with one or two (rightmost) unspecified digits in a year-only expression (year precision)
    Example 1       ‘201X’
    Example 2       ‘20XX’
    Year specified, month unspecified in a year-month expression (month precision)
    Example 3       ‘2004-XX’
    Year and month specified, day unspecified in a year-month-day expression (day precision)
    Example 4       ‘1985-04-XX’
    Year specified, day and month unspecified in a year-month-day expression  (day precision)
    Example 5       ‘1985-XX-XX’


*/

func IsUnspecifiedDigits(edtf_str string) bool {
	return re_unspecified.MatchString(edtf_str)
}

func ParseUnspecifiedDigits(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_unspecified.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 1 unspecified digits string")
	}

	return nil, nil
}

/*

Extended Interval (L1)

    A null string may be used for the start or end date when it is unknown.
    Double-dot (“..”) may be used when either the start or end date is not specified, either because there is none or for any other reason.
    A modifier may appear at the end of the date to indicate "uncertain" and/or "approximate"

Open end time interval

    Example 1          ‘1985-04-12/..’
    interval starting at 1985 April 12th with day precision; end open
    Example 2          ‘1985-04/..’
    interval starting at 1985 April with month precision; end open
    Example 3          ‘1985/..’
    interval starting at year 1985 with year precision; end open

Open start time interval

    Example 4          ‘../1985-04-12’
    interval with open start; ending 1985 April 12th with day precision
    Example 5          ‘../1985-04’
    interval with open start; ending 1985 April with month precision
    Example 6          ‘../1985’
    interval with open start; ending at year 1985 with year precision

Time interval with unknown end

    Example 7          ‘1985-04-12/’
    interval starting 1985 April 12th with day precision; end unknown
    Example 8          ‘1985-04/’
    interval starting 1985 April with month precision; end unknown
    Example 9          ‘1985/’
    interval starting year 1985 with year precision; end unknown

Time interval with unknown start

    Example 10       ‘/1985-04-12’
    interval with unknown start; ending 1985 April 12th with day precision
    Example 11       ‘/1985-04’
    interval with unknown start; ending 1985 April with month precision
    Example 12       ‘/1985’
    interval with unknown start; ending year 1985 with year precision

*/

func IsExtendedInterval(edtf_str string) bool {

	if re_interval_end.MatchString(edtf_str) {
		return true
	}

	if re_interval_start.MatchString(edtf_str) {
		return true
	}

	return true
}

func ParseExtendedInterval(edtf_str string) (*edtf.EDTFDate, error) {

	if re_interval_start.MatchString(edtf_str) {
		return ParseExtendedIntervalStart(edtf_str)
	}

	if re_interval_end.MatchString(edtf_str) {
		return ParseExtendedIntervalEnd(edtf_str)
	}

	return nil, errors.New("Invalid extended interval string")
}

func ParseExtendedIntervalEnd(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_interval_end.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 1 extended interval (end) string")
	}

	return nil, nil
}

func ParseExtendedIntervalStart(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_interval_start.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 1 extended interval (start) string")
	}

	return nil, nil
}

/*

 Negative calendar year

    Example 1       ‘-1985’

Note: ISO 8601 Part 1 does not support negative year.

*/

func IsNegativeCalendarYear(edtf_str string) bool {
	return re_negative_year.MatchString(edtf_str)
}

func ParseNegativeCalendarYear(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_negative_year.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 1 negative year string")
	}

	return nil, nil
}
