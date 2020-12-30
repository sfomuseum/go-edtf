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
		// Go is incapable of parsing these dates
		// "Y170000002",
		// "Y-17000002",
		"Y1700",
		"Y-1200",
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
