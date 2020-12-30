package re

import (
	"regexp"
	"strings"
)

var LetterPrefixedCalendarYear *regexp.Regexp
var Season *regexp.Regexp
var QualifiedDate *regexp.Regexp
var UnspecifiedDigits *regexp.Regexp
var IntervalEnd *regexp.Regexp
var IntervalStart *regexp.Regexp
var NegativeYear *regexp.Regexp
var Level1 *regexp.Regexp

func init() {

	LetterPrefixedCalendarYear = regexp.MustCompile(`^` + PATTERN_LETTER_PREFIXED_CALENDAR_YEAR + `$`)

	Season = regexp.MustCompile(`^` + PATTERN_SEASON + `$`)

	QualifiedDate = regexp.MustCompile(`^` + PATTERN_QUALIFIED_DATE + `$`)

	UnspecifiedDigits = regexp.MustCompile(`^` + PATTERN_UNSPECIFIED_DIGITS + `$`)

	IntervalStart = regexp.MustCompile(`^` + PATTERN_INTERVAL_START + `$`)

	IntervalEnd = regexp.MustCompile(`^` + PATTERN_INTERVAL_END + `$`)

	NegativeYear = regexp.MustCompile(`^` + PATTERN_NEGATIVE_YEAR + `$`)

	level1_patterns := []string{
		PATTERN_LETTER_PREFIXED_CALENDAR_YEAR,
		PATTERN_SEASON,
		PATTERN_QUALIFIED_DATE,
		PATTERN_UNSPECIFIED_DIGITS,
		PATTERN_INTERVAL_START,
		PATTERN_INTERVAL_END,
		PATTERN_NEGATIVE_YEAR,
	}

	Level1 = regexp.MustCompile(`^(` + strings.Join(level1_patterns, "|") + `)$`)
}
