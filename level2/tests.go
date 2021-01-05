package level2

import (
	"github.com/whosonfirst/go-edtf/tests"
)

var Tests map[string]map[string]*tests.TestResult = map[string]map[string]*tests.TestResult{
	EXPONENTIAL_YEAR: map[string]*tests.TestResult{
		"Y-17E7": tests.NewTestResult(tests.TestResultOptions{}), // TO DO - https://github.com/whosonfirst/go-edtf/issues/5
		"Y10E7":  tests.NewTestResult(tests.TestResultOptions{}), // TO DO
		"Y20E2":  tests.NewTestResult(tests.TestResultOptions{}),
	},
	SIGNIFICANT_DIGITS: map[string]*tests.TestResult{
		"1950S2":       tests.NewTestResult(tests.TestResultOptions{}),
		"Y171010000S3": tests.NewTestResult(tests.TestResultOptions{}),
		"Y-20E2S3":     tests.NewTestResult(tests.TestResultOptions{}),
		"Y3388E2S3":    tests.NewTestResult(tests.TestResultOptions{}),
	},
	SUB_YEAR_GROUPINGS: map[string]*tests.TestResult{
		"2001-34": tests.NewTestResult(tests.TestResultOptions{}),
		// "second quarter of 2001": tests.NewTestResult(tests.TestResultOptions{}),	// TO DO
	},
	SET_REPRESENTATIONS: map[string]*tests.TestResult{
		"[1667,1668,1670..1672]":      tests.NewTestResult(tests.TestResultOptions{}),
		"[..1760-12-03]":              tests.NewTestResult(tests.TestResultOptions{}),
		"[1760-12..]":                 tests.NewTestResult(tests.TestResultOptions{}),
		"[1760-01,1760-02,1760-12..]": tests.NewTestResult(tests.TestResultOptions{}),
		"[1667,1760-12]":              tests.NewTestResult(tests.TestResultOptions{}),
		"[..1984]":                    tests.NewTestResult(tests.TestResultOptions{}),
		"{1667,1668,1670..1672}":      tests.NewTestResult(tests.TestResultOptions{}),
		"{1960,1961-12}":              tests.NewTestResult(tests.TestResultOptions{}),
		"{..1984}":                    tests.NewTestResult(tests.TestResultOptions{}),
	},
	GROUP_QUALIFICATION: map[string]*tests.TestResult{
		"2004-06-11%": tests.NewTestResult(tests.TestResultOptions{}),
		"2004-06~-11": tests.NewTestResult(tests.TestResultOptions{}),
		"2004?-06-11": tests.NewTestResult(tests.TestResultOptions{}),
	},
	INDIVIDUAL_QUALIFICATION: map[string]*tests.TestResult{
		"?2004-06-~11": tests.NewTestResult(tests.TestResultOptions{}),
		"2004-%06-11":  tests.NewTestResult(tests.TestResultOptions{}),
	},
	UNSPECIFIED_DIGIT: map[string]*tests.TestResult{
		"156X-12-25": tests.NewTestResult(tests.TestResultOptions{}),
		"15XX-12-25": tests.NewTestResult(tests.TestResultOptions{}),
		// "XXXX-12-XX": tests.NewTestResult(tests.TestResultOptions{}),	// TO DO
		"1XXX-XX": tests.NewTestResult(tests.TestResultOptions{}),
		"1XXX-12": tests.NewTestResult(tests.TestResultOptions{}),
		"1984-1X": tests.NewTestResult(tests.TestResultOptions{}),
	},
	INTERVAL: map[string]*tests.TestResult{
		"2004-06-~01/2004-06-~20": tests.NewTestResult(tests.TestResultOptions{}),
		"2004-06-XX/2004-07-03":   tests.NewTestResult(tests.TestResultOptions{}),
	},
}
