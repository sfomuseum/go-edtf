package parser

import (
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/tests"
)

var Tests map[string]map[string]*tests.TestResult = map[string]map[string]*tests.TestResult{
	"Unknown": map[string]*tests.TestResult{
		edtf.UNKNOWN: tests.NewTestResult(tests.TestResultOptions{
			StartLowerIsUnknown: true,
			StartUpperIsUnknown: true,
			EndLowerIsUnknown:   true,
			EndUpperIsUnknown:   true,
		}),
	},
	"Open": map[string]*tests.TestResult{
		edtf.OPEN: tests.NewTestResult(tests.TestResultOptions{
			StartLowerIsOpen: true,
			StartUpperIsOpen: true,
			EndLowerIsOpen:   true,
			EndUpperIsOpen:   true,
		}),
	},
}
