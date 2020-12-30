package re

import (
	"regexp"
	"strings"
)

var Date *regexp.Regexp
var DateAndTime *regexp.Regexp
var TimeInterval *regexp.Regexp

var Level0 *regexp.Regexp

func init() {

	Date = regexp.MustCompile(`^` + PATTERN_DATE + `$`)

	DateAndTime = regexp.MustCompile(`^` + PATTERN_DATE_AND_TIME + `$`)

	TimeInterval = regexp.MustCompile(`^` + PATTERN_TIME_INTERVAL + `$`)

	level0_patterns := []string{
		PATTERN_DATE,
		PATTERN_DATE_AND_TIME,
		PATTERN_TIME_INTERVAL,
	}

	Level0 = regexp.MustCompile(`^(` + strings.Join(level0_patterns, "|") + `)$`)
}
