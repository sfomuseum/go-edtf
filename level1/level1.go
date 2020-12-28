package level1

import (
	"errors"
	"github.com/whosonfirst/go-edtf"
	"regexp"
	"strings"
)

const LEVEL int = 1

const PATTERN_CALENDAR_YEAR string = `^Y(\-)?(\d+)$`
const PATTERN_SEASON string = `^(\d{4})-(0[1-9]|2[1-4])|(?i)(spring|summer|fall|winter)\s*,\s*(\d{4})$`
const PATTERN_QUALIFIED string = `^(?:(\d{4})(?:-(\d{2})(?:-(\d{2}))?)?)(\?|~|%)$`

// const PATTERN_UNSPECIFIED string = `^(?:(\d{3})(X)|(\d{2})(XX)|(\d{4})-(XX)|(\d{4})\-(\d{2})\-(XX)|(\d{4})\-(XX)\-(XX))$`

const PATTERN_UNSPECIFIED string = `^(?:([0-9X]{4})(?:-([0-9X]{2})(?:-([0-9X]{2}))?)?)$`

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

var Tests map[string][]string = map[string][]string{
	"prefixed_calendar_year": []string{
		// "Y170000002",
		// "Y-17000002",
	},
	"season": []string{
		"2001-01",
		"2001-24",
		"Spring, 2002",
		"winter, 2002",
	},
	"qualified_date": []string{
		"1984?",
		"2004-06~",
		"2004-06-11%",
	},
	"unspecified_digits": []string{
		"201X",
		"20XX",
		"2004-XX",
		"1985-04-XX",
		"1985-XX-XX",
	},
	"extended_interval_start": []string{
		"../1985-04-12",
		"../1985-04",
		"../1985",
		"/1985-04-12",
		"/1985-04",
		"/1985",
	},
	"extended_interval_end": []string{
		"1985-04-12/..",
		"1985-04/..",
		"1985/..",
		"1985-04-12/",
		"1985-04/",
		"1985/",
	},
	"negative_calendar_year": []string{
		"-1985",
	},
}

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
