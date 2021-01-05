package level2

import (
	"github.com/whosonfirst/go-edtf/tests"
)

var Tests map[string]map[string]*tests.TestResult = map[string]map[string]*tests.TestResult{
	EXPONENTIAL_YEAR: map[string]*tests.TestResult{
		"Y-17E7": nil, // TO DO - https://github.com/whosonfirst/go-edtf/issues/5
		"Y10E7":  nil, // TO DO
		"Y20E2":  nil,
	},
	SIGNIFICANT_DIGITS: map[string]*tests.TestResult{
		"1950S2":       nil,
		"Y171010000S3": nil,
		"Y-20E2S3":     nil,
		"Y3388E2S3":    nil,
	},
	SUB_YEAR_GROUPINGS: map[string]*tests.TestResult{
		"2001-34": nil,
		// "second quarter of 2001": nil,	// TO DO
	},
	SET_REPRESENTATIONS: map[string]*tests.TestResult{
		"[1667,1668,1670..1672]":      nil,
		"[..1760-12-03]":              nil,
		"[1760-12..]":                 nil,
		"[1760-01,1760-02,1760-12..]": nil,
		"[1667,1760-12]":              nil,
		"[..1984]":                    nil,
		"{1667,1668,1670..1672}":      nil,
		"{1960,1961-12}":              nil,
		"{..1984}":                    nil,
	},
	GROUP_QUALIFICATION: map[string]*tests.TestResult{
		"2004-06-11%": nil,
		"2004-06~-11": nil,
		"2004?-06-11": nil,
	},
	INDIVIDUAL_QUALIFICATION: map[string]*tests.TestResult{
		"?2004-06-~11": nil,
		"2004-%06-11":  nil,
	},
	UNSPECIFIED_DIGIT: map[string]*tests.TestResult{
		"156X-12-25": nil,
		"15XX-12-25": nil,
		// "XXXX-12-XX": nil,	// TO DO
		"1XXX-XX": nil,
		"1XXX-12": nil,
		"1984-1X": nil,
	},
	INTERVAL: map[string]*tests.TestResult{
		"2004-06-~01/2004-06-~20": nil,
		"2004-06-XX/2004-07-03":   nil,
	},
}
