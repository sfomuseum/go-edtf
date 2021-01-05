package level1

import (
	"github.com/whosonfirst/go-edtf/tests"
)

var Tests map[string]map[string]*tests.TestResult = map[string]map[string]*tests.TestResult{
	LETTER_PREFIXED_CALENDAR_YEAR: map[string]*tests.TestResult{
		"Y170000002": nil, // TO DO
		"Y-17000002": nil, // TO DO
		"Y1700":      nil,
		"Y-1200":     nil,
	},
	SEASON: map[string]*tests.TestResult{
		"2001-01": nil,
		"2001-24": nil,
		// "-0011-24": nil,	// TO DO
		// "-0301-05": nil,	// TO DO
		"Spring, 2002": nil,
		"winter, 2002": nil,
		// "Summer, -1980": nil,	// TO DO
	},
	QUALIFIED_DATE: map[string]*tests.TestResult{
		"1984?":       nil,
		"2004-06~":    nil,
		"2004-06-11%": nil,
	},
	UNSPECIFIED_DIGITS: map[string]*tests.TestResult{
		"201X":       nil,
		"20XX":       nil,
		"2004-XX":    nil,
		"1985-04-XX": nil,
		"1985-XX-XX": nil,
	},
	EXTENDED_INTERVAL_START: map[string]*tests.TestResult{
		"../1985-04-12": nil,
		"../1985-04":    nil,
		"../1985":       nil,
		"/1985-04-12":   nil,
		"/1985-04":      nil,
		"/1985":         nil,
	},
	EXTENDED_INTERVAL_END: map[string]*tests.TestResult{
		"1985-04-12/..": nil,
		"1985-04/..":    nil,
		"1985/..":       nil,
		"1985-04-12/":   nil,
		"1985-04/":      nil,
		"1985/":         nil,
	},
	NEGATIVE_CALENDAR_YEAR: map[string]*tests.TestResult{
		"-1985": nil,
	},
}
