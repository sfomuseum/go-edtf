package level0

import (
	"github.com/whosonfirst/go-edtf/tests"
)

var Tests map[string]map[string]*tests.TestResult = map[string]map[string]*tests.TestResult{
	DATE: map[string]*tests.TestResult{
		"1985-04-12": nil,
		"1985-04":    nil,
		"1985":       nil,
		"-0400":      nil,
		"-1200-06":   nil,
	},
	DATE_AND_TIME: map[string]*tests.TestResult{
		"1985-04-12T23:20:30":       nil,
		"1985-04-12T23:20:30Z":      nil,
		"1985-04-12T23:20:30-04":    nil,
		"1985-04-12T23:20:30+04:30": nil,
		"-1972-04-12T23:20:28":      nil,
	},
	TIME_INTERVAL: map[string]*tests.TestResult{
		"1964/2008":             nil,
		"2004-06/2006-08":       nil,
		"2004-02-01/2005-02-08": nil,
		"2004-02-01/2005-02":    nil,
		"2004-02-01/2005":       nil,
		"2005/2006-02":          nil,
		"-0200/0200":            nil,
		"-1200-06/0200-05-02":   nil,
	},
}
