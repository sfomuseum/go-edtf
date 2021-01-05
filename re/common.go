package re

import (
	"regexp"
)

var Year *regexp.Regexp

var YMD *regexp.Regexp

var QualifiedIndividual *regexp.Regexp
var QualifiedGroup *regexp.Regexp

func init() {
	Year = regexp.MustCompile(`^` + PATTERN_YEAR + `$`)

	YMD = regexp.MustCompile(`^` + PATTERN_YMD_X + `$`)

	QualifiedIndividual = regexp.MustCompile(`^(` + PATTERN_QUALIFIER + `)?` + PATTERN_DATE_X + `$`)

	QualifiedGroup = regexp.MustCompile(`^` + PATTERN_DATE_X + `(` + PATTERN_QUALIFIER + `)?$`)
}
