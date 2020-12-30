package re

import (
	"regexp"
)

var Year *regexp.Regexp

func init() {
	Year = regexp.MustCompile(`^` + PATTERN_YEAR + `$`)
}
