package level1

import (
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/tests"
)

var Tests map[string]map[string]*tests.TestResult = map[string]map[string]*tests.TestResult{
	LETTER_PREFIXED_CALENDAR_YEAR: map[string]*tests.TestResult{
		"Y170000002": tests.NewTestResult(tests.TestResultOptions{}), // TO DO
		"Y-17000002": tests.NewTestResult(tests.TestResultOptions{}), // TO DO
		"Y1700": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1700-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "1700-01-01T23:59:59Z",
			EndLowerTimeRFC3339:   "1700-12-31T00:00:00Z",
			EndUpperTimeRFC3339:   "1700-12-31T23:59:59Z",
		}),
		"Y-1200": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "-1200-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "-1200-01-01T23:59:59Z",
			EndLowerTimeRFC3339:   "-1200-12-31T00:00:00Z",
			EndUpperTimeRFC3339:   "-1200-12-31T23:59:59Z",
		}),
	},
	SEASON: map[string]*tests.TestResult{
		"2001-01": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2001-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "2001-01-01T23:59:59Z",
			EndLowerTimeRFC3339:   "2001-01-31T00:00:00Z",
			EndUpperTimeRFC3339:   "2001-01-31T23:59:59Z",
		}),
		"2019-24": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2019-12-01T00:00:00Z",
			StartUpperTimeRFC3339: "2019-12-01T23:59:59Z",
			EndLowerTimeRFC3339:   "2020-02-29T00:00:00Z",
			EndUpperTimeRFC3339:   "2020-02-29T23:59:59Z", // leap year
		}),
		"Spring, 2002": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2002-03-01T00:00:00Z",
			StartUpperTimeRFC3339: "2002-03-01T23:59:59Z",
			EndLowerTimeRFC3339:   "2002-05-31T00:00:00Z",
			EndUpperTimeRFC3339:   "2002-05-31T23:59:59Z",
		}),
		"winter, 2002": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2002-12-01T00:00:00Z",
			StartUpperTimeRFC3339: "2002-12-01T23:59:59Z",
			EndLowerTimeRFC3339:   "2003-02-28T00:00:00Z",
			EndUpperTimeRFC3339:   "2003-02-28T23:59:59Z",
		}),
		/*
			"Summer, -1980": tests.NewTestResult(tests.TestResultOptions{
				StartLowerTimeRFC3339: "-1980-06-01T00:00:00Z",
				StartUpperTimeRFC3339: "-1980-06-01T23:59:59Z",
				EndLowerTimeRFC3339:   "-1980-08-31T00:00:00Z",
				EndUpperTimeRFC3339:   "-1980-08-31T23:59:59Z",
			}),
		*/
	},
	QUALIFIED_DATE: map[string]*tests.TestResult{
		"1984?": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1984-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "1984-01-01T23:59:59Z",
			EndLowerTimeRFC3339:   "1984-12-31T00:00:00Z",
			EndUpperTimeRFC3339:   "1984-12-31T23:59:59Z",
			StartUpperUncertain:   edtf.ANNUAL,
		}),
		"2004-06~": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2004-06-01T00:00:00Z",
			StartUpperTimeRFC3339: "2004-06-01T23:59:59Z",
			EndLowerTimeRFC3339:   "2004-06-30T00:00:00Z",
			EndUpperTimeRFC3339:   "2004-06-30T23:59:59Z",
			EndLowerApproximate:   edtf.MONTHLY,
		}),
		"2004-06-11%": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2004-06-11T00:00:00Z",
			StartUpperTimeRFC3339: "2004-06-11T00:00:00Z",
			EndLowerTimeRFC3339:   "2004-06-11T23:59:59Z",
			EndUpperTimeRFC3339:   "2004-06-11T23:59:59Z",
			EndLowerUncertain:     edtf.DAILY,
			EndLowerApproximate:   edtf.DAILY,
		}),
	},
	UNSPECIFIED_DIGITS: map[string]*tests.TestResult{
		"201X": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
		"20XX": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
		"2004-XX": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
		"1985-04-XX": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
		"1985-XX-XX": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
	},
	EXTENDED_INTERVAL_START: map[string]*tests.TestResult{
		"../1985-04-12": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
		"../1985-04": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
		"../1985": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
		"/1985-04-12": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
		"/1985-04": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
		"/1985": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
	},
	EXTENDED_INTERVAL_END: map[string]*tests.TestResult{
		"1985-04-12/..": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
		"1985-04/..": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
		"1985/..": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
		"1985-04-12/": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
		"1985-04/": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
		"1985/": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
	},
	NEGATIVE_CALENDAR_YEAR: map[string]*tests.TestResult{
		"-1985": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "",
			StartUpperTimeRFC3339: "",
			EndLowerTimeRFC3339:   "",
			EndUpperTimeRFC3339:   "",
		}),
	},
}
