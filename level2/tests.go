package level2

import (
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/tests"
)

var Tests map[string]map[string]*tests.TestResult = map[string]map[string]*tests.TestResult{
	EXPONENTIAL_YEAR: map[string]*tests.TestResult{
		"Y-17E7": tests.NewTestResult(tests.TestResultOptions{}), // TO DO - https://github.com/sfomuseum/go-edtf/issues/5
		"Y10E7":  tests.NewTestResult(tests.TestResultOptions{}), // TO DO
		"Y20E2": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2000-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "2000-01-01T23:59:59Z",
			EndLowerTimeRFC3339:   "2000-12-31T00:00:00Z",
			EndUpperTimeRFC3339:   "2000-12-31T23:59:59Z",
		}),
	},
	SIGNIFICANT_DIGITS: map[string]*tests.TestResult{
		"1950S2": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1900-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "1900-12-31T23:59:59Z",
			EndLowerTimeRFC3339:   "1999-01-01T00:00:00Z",
			EndUpperTimeRFC3339:   "1999-12-31T23:59:59Z",
		}),
		"Y171010000S3": tests.NewTestResult(tests.TestResultOptions{}),
		"Y-20E2S3": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "-2999-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "-2999-12-31T23:59:59Z",
			EndLowerTimeRFC3339:   "-2000-01-01T00:00:00Z",
			EndUpperTimeRFC3339:   "-2000-12-31T23:59:59Z",
		}),
		"Y3388E2S3": tests.NewTestResult(tests.TestResultOptions{}),
	},
	SUB_YEAR_GROUPINGS: map[string]*tests.TestResult{
		"2001-34": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2001-04-01T00:00:00Z",
			StartUpperTimeRFC3339: "2001-04-01T23:59:59Z",
			EndLowerTimeRFC3339:   "2001-06-30T00:00:00Z",
			EndUpperTimeRFC3339:   "2001-06-30T23:59:59Z",
		}),
		"2019-28": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2019-12-01T00:00:00Z",
			StartUpperTimeRFC3339: "2019-12-01T23:59:59Z",
			EndLowerTimeRFC3339:   "2020-02-01T00:00:00Z",
			EndUpperTimeRFC3339:   "2020-02-29T23:59:59Z",
		}),
		// "second quarter of 2001": tests.NewTestResult(tests.TestResultOptions{}),	// TO DO
	},
	SET_REPRESENTATIONS: map[string]*tests.TestResult{
		"[1760-01,1760-02,1760-12..]": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1760-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "1760-12-31T23:59:59Z",
			EndLowerIsOpen:        true,
			EndUpperIsOpen:        true,
			StartLowerInclusivity: edtf.ANY,
		}),
		"[1667,1668,1670..1672]": tests.NewTestResult(tests.TestResultOptions{
			// THIS FEELS WRONG...LIKE IT'S BACKWARDS
			StartLowerTimeRFC3339: "1667-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "1667-12-31T23:59:59Z",
			EndLowerTimeRFC3339:   "1672-01-01T00:00:00Z",
			EndUpperTimeRFC3339:   "1672-12-31T23:59:59Z",
			StartLowerInclusivity: edtf.ANY,
			EndUpperInclusivity:   edtf.ANY,
		}),
		"[..1760-12-03]": tests.NewTestResult(tests.TestResultOptions{
			EndLowerTimeRFC3339: "1760-12-03T00:00:00Z",
			EndUpperTimeRFC3339: "1760-12-03T23:59:59Z",
			StartLowerIsOpen:    true,
			StartUpperIsOpen:    true,
			EndUpperInclusivity: edtf.ANY,
		}),
		"[1760-12..]": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1760-12-01T00:00:00Z",
			StartUpperTimeRFC3339: "1760-12-31T23:59:59Z",
			EndLowerIsOpen:        true,
			EndUpperIsOpen:        true,
			StartUpperInclusivity: edtf.ANY,
		}),
		"[1667,1760-12]": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1667-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "1667-12-31T23:59:59Z",
			EndLowerTimeRFC3339:   "1760-12-01T00:00:00Z",
			EndUpperTimeRFC3339:   "1760-12-31T23:59:59Z",
			StartUpperInclusivity: edtf.ANY,
			EndLowerInclusivity:   edtf.ANY,
		}),

		"[..1984]": tests.NewTestResult(tests.TestResultOptions{
			StartLowerIsOpen:    true,
			StartUpperIsOpen:    true,
			EndLowerTimeRFC3339: "1984-01-01T00:00:00Z",
			EndUpperTimeRFC3339: "1984-12-31T23:59:59Z",
			EndLowerInclusivity: edtf.ANY,
		}),
		"{1667,1668,1670..1672}": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1667-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "1667-12-31T23:59:59Z",
			EndLowerTimeRFC3339:   "1672-01-01T00:00:00Z",
			EndUpperTimeRFC3339:   "1672-12-31T23:59:59Z",
			StartUpperInclusivity: edtf.ALL,
			EndLowerInclusivity:   edtf.ALL,
		}),
		"{1960,1961-12}": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1960-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "1960-12-31T23:59:59Z",
			EndLowerTimeRFC3339:   "1961-12-01T00:00:00Z",
			EndUpperTimeRFC3339:   "1961-12-31T23:59:59Z",
			StartUpperInclusivity: edtf.ALL,
			EndLowerInclusivity:   edtf.ALL,
		}),
		"{..1984}": tests.NewTestResult(tests.TestResultOptions{
			StartLowerIsOpen:    true,
			StartUpperIsOpen:    true,
			EndLowerTimeRFC3339: "1984-01-01T00:00:00Z",
			EndUpperTimeRFC3339: "1984-12-31T23:59:59Z",
			EndLowerInclusivity: edtf.ALL,
		}),
	},
	GROUP_QUALIFICATION: map[string]*tests.TestResult{
		"2004-06-11%": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2004-06-11T00:00:00Z",
			StartUpperTimeRFC3339: "2004-06-11T00:00:00Z",
			EndLowerTimeRFC3339:   "2004-06-11T23:59:59Z",
			EndUpperTimeRFC3339:   "2004-06-11T23:59:59Z",
			StartUpperUncertain:   edtf.DAY,
			StartUpperApproximate: edtf.DAY,
			EndLowerApproximate:   edtf.YEAR,
			EndLowerUncertain:     edtf.MONTH,
		}),
		"2004-06~-11": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2004-06-11T00:00:00Z",
			StartUpperTimeRFC3339: "2004-06-11T00:00:00Z",
			EndLowerTimeRFC3339:   "2004-06-11T23:59:59Z",
			EndUpperTimeRFC3339:   "2004-06-11T23:59:59Z",
			EndUpperApproximate:   edtf.MONTH,
			EndLowerApproximate:   edtf.YEAR,
		}),
		"2004?-06-11": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2004-06-11T00:00:00Z",
			StartUpperTimeRFC3339: "2004-06-11T00:00:00Z",
			EndLowerTimeRFC3339:   "2004-06-11T23:59:59Z",
			EndUpperTimeRFC3339:   "2004-06-11T23:59:59Z",
			EndLowerUncertain:     edtf.YEAR,
		}),
	},
	INDIVIDUAL_QUALIFICATION: map[string]*tests.TestResult{
		"?2004-06-~11": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2004-06-11T00:00:00Z",
			StartUpperTimeRFC3339: "2004-06-11T00:00:00Z",
			EndLowerTimeRFC3339:   "2004-06-11T23:59:59Z",
			EndUpperTimeRFC3339:   "2004-06-11T23:59:59Z",
			EndUpperApproximate:   edtf.DAY,
		}),
		"2004-%06-11": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2004-06-11T00:00:00Z",
			StartUpperTimeRFC3339: "2004-06-11T00:00:00Z",
			EndLowerTimeRFC3339:   "2004-06-11T23:59:59Z",
			EndUpperTimeRFC3339:   "2004-06-11T23:59:59Z",
			EndUpperApproximate:   edtf.MONTH,
			EndUpperUncertain:     edtf.MONTH,
		}),
	},
	UNSPECIFIED_DIGIT: map[string]*tests.TestResult{
		"156X-12-25": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1560-12-25T00:00:00Z",
			StartUpperTimeRFC3339: "1560-12-25T23:59:59Z",
			EndLowerTimeRFC3339:   "1569-12-25T00:00:00Z",
			EndUpperTimeRFC3339:   "1569-12-25T23:59:59Z",
			StartUpperPrecision:   edtf.DECADE,
		}),
		"15XX-12-25": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1500-12-25T00:00:00Z",
			StartUpperTimeRFC3339: "1500-12-25T23:59:59Z",
			EndLowerTimeRFC3339:   "1599-12-25T00:00:00Z",
			EndUpperTimeRFC3339:   "1599-12-25T23:59:59Z",
			StartUpperPrecision:   edtf.CENTURY,
		}),
		// "XXXX-12-XX": tests.NewTestResult(tests.TestResultOptions{}),	// TO DO
		"1XXX-XX": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1000-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "1000-01-01T23:59:59Z",
			EndLowerTimeRFC3339:   "1999-12-31T00:00:00Z",
			EndUpperTimeRFC3339:   "1999-12-31T23:59:59Z",
			StartUpperPrecision:   edtf.MILLENIUM,
		}),
		"1XXX-12": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1000-12-01T00:00:00Z",
			StartUpperTimeRFC3339: "1000-12-01T23:59:59Z",
			EndLowerTimeRFC3339:   "1999-12-31T00:00:00Z",
			EndUpperTimeRFC3339:   "1999-12-31T23:59:59Z",
			StartUpperPrecision:   edtf.MILLENIUM,
		}),
		"1984-1X": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "1984-10-01T00:00:00Z",
			StartUpperTimeRFC3339: "1984-10-01T23:59:59Z",
			EndLowerTimeRFC3339:   "1984-12-31T00:00:00Z",
			EndUpperTimeRFC3339:   "1984-12-31T23:59:59Z",
			StartUpperPrecision:   edtf.MONTH,
		}),
	},
	INTERVAL: map[string]*tests.TestResult{
		"2004-06-~01/2004-06-~20": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2004-06-01T00:00:00Z",
			StartUpperTimeRFC3339: "2004-06-01T23:59:59Z",
			EndLowerTimeRFC3339:   "2004-06-20T00:00:00Z",
			EndUpperTimeRFC3339:   "2004-06-20T23:59:59Z",
			EndUpperApproximate:   edtf.DAY,
		}),
		"2004-06-XX/2004-07-03": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "2004-06-01T00:00:00Z",
			StartUpperTimeRFC3339: "2004-06-30T23:59:59Z",
			EndLowerTimeRFC3339:   "2004-07-03T00:00:00Z",
			EndUpperTimeRFC3339:   "2004-07-03T23:59:59Z",
		}),
		"~-0100/~2020": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "-0100-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "-0100-12-31T23:59:59Z",
			EndLowerTimeRFC3339:   "2020-01-01T00:00:00Z",
			EndUpperTimeRFC3339:   "2020-12-31T23:59:59Z",
		}),
		"~-0100/~-0010": tests.NewTestResult(tests.TestResultOptions{
			StartLowerTimeRFC3339: "-0100-01-01T00:00:00Z",
			StartUpperTimeRFC3339: "-0100-12-31T23:59:59Z",
			EndLowerTimeRFC3339:   "-0010-01-01T00:00:00Z",
			EndUpperTimeRFC3339:   "-0010-12-31T23:59:59Z",
		}),
	},
}
