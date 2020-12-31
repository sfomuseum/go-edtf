package re

import (
	"regexp"
	"strings"
)

var ExponentialYear *regexp.Regexp

var SignificantDigits *regexp.Regexp
var SubYearGrouping *regexp.Regexp
var SetRepresentations *regexp.Regexp
var GroupQualification *regexp.Regexp
var IndividualQualification *regexp.Regexp
var UnspecifiedDigit *regexp.Regexp
var Interval *regexp.Regexp
var Level2 *regexp.Regexp

func init() {

	ExponentialYear = regexp.MustCompile(`^` + PATTERN_EXPONENTIAL_YEAR + `$`)

	SignificantDigits = regexp.MustCompile(`^` + PATTERN_SIGNIFICANT_DIGITS + `$`)

	SubYearGrouping = regexp.MustCompile(`^` + PATTERN_SUB_YEAR_GROUPING + `$`)

	SetRepresentations = regexp.MustCompile(`^` + PATTERN_SET_REPRESENTATIONS + `$`)

	GroupQualification = regexp.MustCompile(`^` + PATTERN_GROUP_QUALIFICATION + `$`)

	IndividualQualification = regexp.MustCompile(`^` + PATTERN_INDIVIDUAL_QUALIFICATION + `$`)

	UnspecifiedDigit = regexp.MustCompile(`^` + PATTERN_UNSPECIFIED_DIGIT + `$`)

	Interval = regexp.MustCompile(`^` + PATTERN_INTERVAL + `$`)

	level2_patterns := []string{
		PATTERN_EXPONENTIAL_YEAR,
		PATTERN_SIGNIFICANT_DIGITS,
		PATTERN_SUB_YEAR_GROUPING,
		PATTERN_SET_REPRESENTATIONS,
		PATTERN_GROUP_QUALIFICATION,
		PATTERN_INDIVIDUAL_QUALIFICATION,
		PATTERN_UNSPECIFIED_DIGIT,
		PATTERN_INTERVAL,
	}

	Level2 = regexp.MustCompile(`^` + `(` + strings.Join(level2_patterns, "|") + `)`)
}
