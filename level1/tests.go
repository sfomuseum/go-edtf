package level1

import (
	"github.com/whosonfirst/go-edtf/tests"
)

var Tests map[string]map[string]*tests.TestResult = map[string]map[string]*tests.TestResult{
	LETTER_PREFIXED_CALENDAR_YEAR: map[string]*tests.TestResult{
		"Y170000002": tests.NewTestResult(tests.TestResultOptions{}), // TO DO
		"Y-17000002": tests.NewTestResult(tests.TestResultOptions{}), // TO DO
		"Y1700":      tests.NewTestResult(tests.TestResultOptions{}),
		"Y-1200":     tests.NewTestResult(tests.TestResultOptions{}),
	},
	SEASON: map[string]*tests.TestResult{
		"2001-01": tests.NewTestResult(tests.TestResultOptions{}),
		"2001-24": tests.NewTestResult(tests.TestResultOptions{}),
		// "-0011-24": tests.NewTestResult(tests.TestResultOptions{}),	// TO DO
		// "-0301-05": tests.NewTestResult(tests.TestResultOptions{}),	// TO DO
		"Spring, 2002": tests.NewTestResult(tests.TestResultOptions{}),
		"winter, 2002": tests.NewTestResult(tests.TestResultOptions{}),
		// "Summer, -1980": tests.NewTestResult(tests.TestResultOptions{}),	// TO DO
	},
	QUALIFIED_DATE: map[string]*tests.TestResult{
		"1984?":       tests.NewTestResult(tests.TestResultOptions{}),
		"2004-06~":    tests.NewTestResult(tests.TestResultOptions{}),
		"2004-06-11%": tests.NewTestResult(tests.TestResultOptions{}),
	},
	UNSPECIFIED_DIGITS: map[string]*tests.TestResult{
		"201X":       tests.NewTestResult(tests.TestResultOptions{}),
		"20XX":       tests.NewTestResult(tests.TestResultOptions{}),
		"2004-XX":    tests.NewTestResult(tests.TestResultOptions{}),
		"1985-04-XX": tests.NewTestResult(tests.TestResultOptions{}),
		"1985-XX-XX": tests.NewTestResult(tests.TestResultOptions{}),
	},
	EXTENDED_INTERVAL_START: map[string]*tests.TestResult{
		"../1985-04-12": tests.NewTestResult(tests.TestResultOptions{}),
		"../1985-04":    tests.NewTestResult(tests.TestResultOptions{}),
		"../1985":       tests.NewTestResult(tests.TestResultOptions{}),
		"/1985-04-12":   tests.NewTestResult(tests.TestResultOptions{}),
		"/1985-04":      tests.NewTestResult(tests.TestResultOptions{}),
		"/1985":         tests.NewTestResult(tests.TestResultOptions{}),
	},
	EXTENDED_INTERVAL_END: map[string]*tests.TestResult{
		"1985-04-12/..": tests.NewTestResult(tests.TestResultOptions{}),
		"1985-04/..":    tests.NewTestResult(tests.TestResultOptions{}),
		"1985/..":       tests.NewTestResult(tests.TestResultOptions{}),
		"1985-04-12/":   tests.NewTestResult(tests.TestResultOptions{}),
		"1985-04/":      tests.NewTestResult(tests.TestResultOptions{}),
		"1985/":         tests.NewTestResult(tests.TestResultOptions{}),
	},
	NEGATIVE_CALENDAR_YEAR: map[string]*tests.TestResult{
		"-1985": tests.NewTestResult(tests.TestResultOptions{}),
	},
}
