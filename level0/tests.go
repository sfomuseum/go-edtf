package level0

import (
	"github.com/whosonfirst/go-edtf/tests"
)

var Tests map[string]map[string]*tests.TestResult = map[string]map[string]*tests.TestResult{
	DATE: map[string]*tests.TestResult{
		"1985-04-12": tests.NewTestResult(tests.TestResultOptions{}),
		"1985-04":    tests.NewTestResult(tests.TestResultOptions{}),
		"1985":       tests.NewTestResult(tests.TestResultOptions{}),
		"-0400":      tests.NewTestResult(tests.TestResultOptions{}),
		"-1200-06":   tests.NewTestResult(tests.TestResultOptions{}),
	},
	DATE_AND_TIME: map[string]*tests.TestResult{
		"1985-04-12T23:20:30":       tests.NewTestResult(tests.TestResultOptions{}),
		"1985-04-12T23:20:30Z":      tests.NewTestResult(tests.TestResultOptions{}),
		"1985-04-12T23:20:30-04":    tests.NewTestResult(tests.TestResultOptions{}),
		"1985-04-12T23:20:30+04:30": tests.NewTestResult(tests.TestResultOptions{}),
		"-1972-04-12T23:20:28":      tests.NewTestResult(tests.TestResultOptions{}),
	},
	TIME_INTERVAL: map[string]*tests.TestResult{
		"1964/2008":             tests.NewTestResult(tests.TestResultOptions{}),
		"2004-06/2006-08":       tests.NewTestResult(tests.TestResultOptions{}),
		"2004-02-01/2005-02-08": tests.NewTestResult(tests.TestResultOptions{}),
		"2004-02-01/2005-02":    tests.NewTestResult(tests.TestResultOptions{}),
		"2004-02-01/2005":       tests.NewTestResult(tests.TestResultOptions{}),
		"2005/2006-02":          tests.NewTestResult(tests.TestResultOptions{}),
		"-0200/0200":            tests.NewTestResult(tests.TestResultOptions{}),
		"-1200-06/0200-05-02":   tests.NewTestResult(tests.TestResultOptions{}),
	},
}
