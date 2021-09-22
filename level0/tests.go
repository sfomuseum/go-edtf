package level0

import (
	"github.com/sfomuseum/go-edtf/tests"
)

var Tests map[string]map[string]*tests.TestResult = map[string]map[string]*tests.TestResult{
	DATE: map[string]*tests.TestResult{
		"1985-04-12": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1985-04-12T00:00:00Z",
			StartUpperTimeRFC3339: "1985-04-12T00:00:00Z",
			EndLowerTimeRFC3339:   "1985-04-12T23:59:59Z",
			EndUpperTimeRFC3339:   "1985-04-12T23:59:59Z",
		}),
		"1985-04": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1985-04-01T00:00:00Z",
			StartUpperTimeRFC3339: "1985-04-01T23:59:59Z",
			EndLowerTimeRFC3339:   "1985-04-30T00:00:00Z",
			EndUpperTimeRFC3339:   "1985-04-30T23:59:59Z",
		}),
		"1985": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1985-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "1985-01-01T23:59:59Z",
			EndLowerTimeRFC3339:   "1985-12-31T00:00:00Z",
			EndUpperTimeRFC3339:   "1985-12-31T23:59:59Z",
		}),
		"-0400": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "-0400-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "-0400-01-01T23:59:59Z",
			EndLowerTimeRFC3339:   "-0400-12-31T00:00:00Z",
			EndUpperTimeRFC3339:   "-0400-12-31T23:59:59Z",
		}),
		"-1200-06": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "-1200-06-01T00:00:00Z",
			StartUpperTimeRFC3339: "-1200-06-01T23:59:59Z",
			EndLowerTimeRFC3339:   "-1200-06-30T00:00:00Z",
			EndUpperTimeRFC3339:   "-1200-06-30T23:59:59Z",
		}),
	},
	DATE_AND_TIME: map[string]*tests.TestResult{
		"1985-04-12T23:20:30": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1985-04-12T23:20:30Z",
			StartUpperTimeRFC3339: "1985-04-12T23:20:30Z",
			EndLowerTimeRFC3339:   "1985-04-12T23:20:30Z",
			EndUpperTimeRFC3339:   "1985-04-12T23:20:30Z",
		}),
		"2021-09-20T21:14:00Z": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2021-09-20T21:14:00Z",
			StartUpperTimeRFC3339: "2021-09-20T21:14:00Z",
			EndLowerTimeRFC3339:   "2021-09-20T21:14:00Z",
			EndUpperTimeRFC3339:   "2021-09-20T21:14:00Z",
		}),
		"1985-04-12T23:20:30Z": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1985-04-12T23:20:30Z",
			StartUpperTimeRFC3339: "1985-04-12T23:20:30Z",
			EndLowerTimeRFC3339:   "1985-04-12T23:20:30Z",
			EndUpperTimeRFC3339:   "1985-04-12T23:20:30Z",
		}),
		"1985-04-12T23:20:30-04": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1985-04-13T03:20:30Z",
			StartUpperTimeRFC3339: "1985-04-13T03:20:30Z",
			EndLowerTimeRFC3339:   "1985-04-13T03:20:30Z",
			EndUpperTimeRFC3339:   "1985-04-13T03:20:30Z",
		}),
		"1985-04-12T23:20:30+04:30": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1985-04-12T18:50:30Z",
			StartUpperTimeRFC3339: "1985-04-12T18:50:30Z",
			EndLowerTimeRFC3339:   "1985-04-12T18:50:30Z",
			EndUpperTimeRFC3339:   "1985-04-12T18:50:30Z",
		}),
		"-1972-04-12T23:20:28": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "-1972-04-12T23:20:28Z",
			StartUpperTimeRFC3339: "-1972-04-12T23:20:28Z",
			EndLowerTimeRFC3339:   "-1972-04-12T23:20:28Z",
			EndUpperTimeRFC3339:   "-1972-04-12T23:20:28Z",
		}),
	},
	TIME_INTERVAL: map[string]*tests.TestResult{
		"1964/2008": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1964-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "1964-12-31T23:59:59Z",
			EndLowerTimeRFC3339:   "2008-01-01T00:00:00Z",
			EndUpperTimeRFC3339:   "2008-12-31T23:59:59Z",
		}),
		"2004-06/2006-08": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2004-06-01T00:00:00Z",
			StartUpperTimeRFC3339: "2004-06-30T23:59:59Z",
			EndLowerTimeRFC3339:   "2006-08-01T00:00:00Z",
			EndUpperTimeRFC3339:   "2006-08-31T23:59:59Z",
		}),
		"2004-02-01/2005-02-08": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2004-02-01T00:00:00Z",
			StartUpperTimeRFC3339: "2004-02-01T23:59:59Z",
			EndLowerTimeRFC3339:   "2005-02-08T00:00:00Z",
			EndUpperTimeRFC3339:   "2005-02-08T23:59:59Z",
		}),
		"2004-02-01/2005-02": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2004-02-01T00:00:00Z",
			StartUpperTimeRFC3339: "2004-02-01T23:59:59Z",
			EndLowerTimeRFC3339:   "2005-02-01T00:00:00Z",
			EndUpperTimeRFC3339:   "2005-02-28T23:59:59Z",
		}),
		"2004-02-01/2005": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2004-02-01T00:00:00Z",
			StartUpperTimeRFC3339: "2004-02-01T23:59:59Z",
			EndLowerTimeRFC3339:   "2005-01-01T00:00:00Z",
			EndUpperTimeRFC3339:   "2005-12-31T23:59:59Z",
		}),
		"2005/2020-02": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2005-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "2005-12-31T23:59:59Z",
			EndLowerTimeRFC3339:   "2020-02-01T00:00:00Z",
			EndUpperTimeRFC3339:   "2020-02-29T23:59:59Z", // leap year
		}),
		"-0200/0200": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "-0200-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "-0200-12-31T23:59:59Z",
			EndLowerTimeRFC3339:   "0200-01-01T00:00:00Z",
			EndUpperTimeRFC3339:   "0200-12-31T23:59:59Z",
		}),
		"-1200-06/0200-05-02": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "-1200-06-01T00:00:00Z",
			StartUpperTimeRFC3339: "-1200-06-30T23:59:59Z",
			EndLowerTimeRFC3339:   "0200-05-02T00:00:00Z",
			EndUpperTimeRFC3339:   "0200-05-02T23:59:59Z",
		}),
	},
}
