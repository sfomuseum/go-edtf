package tests

import (
	"github.com/whosonfirst/go-edtf"
)

type TestResult struct {
	Start *TestDateRange
	End   *TestDateRange
}

type TestDateRange struct {
	Lower *TestDate
	Upper *TestDate
}

type TestDate struct {
	Time string
}

type TestResultOptions struct {
}

func NewTestResult(opts TestResultOptions) *TestResult {

	start := NewTestDateRange(opts)
	end := NewTestDateRange(opts)

	r := &TestResult{
		Start: start,
		End:   end,
	}

	return r
}

func (r *TestResult) TestDate(d *edtf.EDTFDate) error {

	return nil
}

func NewTestDateRange(opts TestResultOptions) *TestDateRange {

	lower := NewTestDate(opts)
	upper := NewTestDate(opts)

	r := &TestDateRange{
		Lower: lower,
		Upper: upper,
	}

	return r
}

func NewTestDate(opts TestResultOptions) *TestDate {

	d := &TestDate{}

	return d
}
