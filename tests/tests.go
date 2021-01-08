package tests

import (
	"fmt"
	"github.com/sfomuseum/go-edtf"
	"time"
)

type TestResult struct {
	options TestResultOptions
}

type TestResultOptions struct {
	StartLowerTimeRFC3339 string
	StartUpperTimeRFC3339 string
	EndLowerTimeRFC3339   string
	EndUpperTimeRFC3339   string
	EndLowerTimeUnix      int64
	StartUpperTimeUnix    int64
	StartLowerTimeUnix    int64
	EndUpperTimeUnix      int64
	StartLowerUncertain   edtf.Precision
	StartUpperUncertain   edtf.Precision
	EndLowerUncertain     edtf.Precision
	EndUpperUncertain     edtf.Precision
	StartLowerApproximate edtf.Precision
	StartUpperApproximate edtf.Precision
	EndLowerApproximate   edtf.Precision
	EndUpperApproximate   edtf.Precision
	StartLowerPrecision   edtf.Precision
	StartUpperPrecision   edtf.Precision
	EndLowerPrecision     edtf.Precision
	EndUpperPrecision     edtf.Precision
	StartLowerIsOpen      bool
	StartUpperIsOpen      bool
	EndLowerIsOpen        bool
	EndUpperIsOpen        bool
	StartLowerIsUnknown   bool
	StartUpperIsUnknown   bool
	EndLowerIsUnknown     bool
	EndUpperIsUnknown     bool
	StartLowerInclusivity edtf.Precision
	StartUpperInclusivity edtf.Precision
	EndLowerInclusivity   edtf.Precision
	EndUpperInclusivity   edtf.Precision
}

func NewTestResult(opts TestResultOptions) *TestResult {

	r := &TestResult{
		options: opts,
	}

	return r
}

func (r *TestResult) TestDate(d *edtf.EDTFDate) error {

	/*

		if d.Start.Lower.Time != nil {
			fmt.Printf("[%s][start.lower] %s %d\n", d.String(), d.Start.Lower.Time.Format(time.RFC3339), d.Start.Lower.Time.Unix())
		}

		if d.Start.Upper.Time != nil {
			fmt.Printf("[%s][start.upper] %s %d\n", d.String(), d.Start.Lower.Time.Format(time.RFC3339), d.Start.Lower.Time.Unix())
		}

		if d.End.Lower.Time != nil {
			fmt.Printf("[%s][end.lower] %s %d\n", d.String(), d.End.Lower.Time.Format(time.RFC3339), d.End.Lower.Time.Unix())
		}

		if d.End.Upper.Time != nil {
			fmt.Printf("[%s][end.upper] %s %d\n", d.String(), d.End.Lower.Time.Format(time.RFC3339), d.End.Lower.Time.Unix())
		}

	*/

	err := r.testRFC3339All(d)

	if err != nil {
		return err
	}

	err = r.testUnixAll(d)

	if err != nil {
		return err
	}

	err = r.testPrecisionAll(d)

	if err != nil {
		return err
	}

	err = r.testUncertainAll(d)

	if err != nil {
		return err
	}

	err = r.testApproximateAll(d)

	if err != nil {
		return err
	}

	err = r.testIsOpenAll(d)

	if err != nil {
		return err
	}

	err = r.testIsUnknownAll(d)

	if err != nil {
		return err
	}

	err = r.testInclusivityAll(d)

	if err != nil {
		return err
	}

	return nil
}

func (r *TestResult) testIsOpenAll(d *edtf.EDTFDate) error {

	err := r.testBoolean(d.Start.Lower.Open, r.options.StartLowerIsOpen)

	if err != nil {
		return fmt.Errorf("Invalid StartLowerIsOpen flag, %v", err)
	}

	err = r.testBoolean(d.Start.Upper.Open, r.options.StartUpperIsOpen)

	if err != nil {
		return fmt.Errorf("Invalid StartUpperIsOpen flag, %v", err)
	}

	err = r.testBoolean(d.End.Lower.Open, r.options.EndLowerIsOpen)

	if err != nil {
		return fmt.Errorf("Invalid EndLowerIsOpen flag, %v", err)
	}

	err = r.testBoolean(d.End.Upper.Open, r.options.EndUpperIsOpen)

	if err != nil {
		return fmt.Errorf("Invalid EndUpperIsOpen flag, %v", err)
	}

	return nil
}

func (r *TestResult) testIsUnknownAll(d *edtf.EDTFDate) error {

	err := r.testBoolean(d.Start.Lower.Unknown, r.options.StartLowerIsUnknown)

	if err != nil {
		return fmt.Errorf("Invalid StartLowerIsUnknown flag, %v", err)
	}

	err = r.testBoolean(d.Start.Upper.Unknown, r.options.StartUpperIsUnknown)

	if err != nil {
		return fmt.Errorf("Invalid StartUpperIsUnknown flag, %v", err)
	}

	err = r.testBoolean(d.End.Lower.Unknown, r.options.EndLowerIsUnknown)

	if err != nil {
		return fmt.Errorf("Invalid EndLowerIsUnknown flag, %v", err)
	}

	err = r.testBoolean(d.End.Upper.Unknown, r.options.EndUpperIsUnknown)

	if err != nil {
		return fmt.Errorf("Invalid EndUpperIsUnknown flag, %v", err)
	}

	return nil
}

func (r *TestResult) testBoolean(candidate bool, expected bool) error {

	if candidate != expected {
		return fmt.Errorf("Boolean test failed, expected '%t' but got '%t'", expected, candidate)
	}

	return nil
}

func (r *TestResult) testInclusivityAll(d *edtf.EDTFDate) error {

	err := r.testPrecision(d.Start.Lower.Inclusivity, r.options.StartLowerInclusivity)

	if err != nil {
		return fmt.Errorf("Invalid StartLowerInclusivity flag, %v", err)
	}

	err = r.testPrecision(d.Start.Upper.Inclusivity, r.options.StartUpperInclusivity)

	if err != nil {
		return fmt.Errorf("Invalid StartUpperInclusivity flag, %v", err)
	}

	err = r.testPrecision(d.End.Lower.Inclusivity, r.options.EndLowerInclusivity)

	if err != nil {
		return fmt.Errorf("Invalid EndLowerInclusivity flag, %v", err)
	}

	err = r.testPrecision(d.End.Upper.Inclusivity, r.options.EndUpperInclusivity)

	if err != nil {
		return fmt.Errorf("Invalid EndUpperInclusivity flag, %v", err)
	}

	return nil
}

func (r *TestResult) testPrecisionAll(d *edtf.EDTFDate) error {

	err := r.testPrecision(d.Start.Lower.Precision, r.options.StartLowerPrecision)

	if err != nil {
		return fmt.Errorf("Invalid StartLowerPrecision flag, %v", err)
	}

	err = r.testPrecision(d.Start.Upper.Precision, r.options.StartUpperPrecision)

	if err != nil {
		return fmt.Errorf("Invalid StartUpperPrecision flag, %v", err)
	}

	err = r.testPrecision(d.End.Lower.Precision, r.options.EndLowerPrecision)

	if err != nil {
		return fmt.Errorf("Invalid EndLowerPrecision flag, %v", err)
	}

	err = r.testPrecision(d.End.Upper.Precision, r.options.EndUpperPrecision)

	if err != nil {
		return fmt.Errorf("Invalid EndUpperPrecision flag, %v", err)
	}

	return nil
}

func (r *TestResult) testUncertainAll(d *edtf.EDTFDate) error {

	err := r.testPrecision(d.Start.Lower.Uncertain, r.options.StartLowerUncertain)

	if err != nil {
		return fmt.Errorf("Invalid StartLowerUncertain flag, %v", err)
	}

	err = r.testPrecision(d.Start.Upper.Uncertain, r.options.StartUpperUncertain)

	if err != nil {
		return fmt.Errorf("Invalid StartUpperUncertain flag, %v", err)
	}

	err = r.testPrecision(d.End.Lower.Uncertain, r.options.EndLowerUncertain)

	if err != nil {
		return fmt.Errorf("Invalid EndLowerUncertain flag, %v", err)
	}

	err = r.testPrecision(d.End.Upper.Uncertain, r.options.EndUpperUncertain)

	if err != nil {
		return fmt.Errorf("Invalid EndUpperUncertain flag, %v", err)
	}

	return nil
}

func (r *TestResult) testApproximateAll(d *edtf.EDTFDate) error {

	err := r.testPrecision(d.Start.Lower.Approximate, r.options.StartLowerApproximate)

	if err != nil {
		return fmt.Errorf("Invalid StartLowerApproximate flag, %v", err)
	}

	err = r.testPrecision(d.Start.Upper.Approximate, r.options.StartUpperApproximate)

	if err != nil {
		return fmt.Errorf("Invalid StartUpperApproximate flag, %v", err)
	}

	err = r.testPrecision(d.End.Lower.Approximate, r.options.EndLowerApproximate)

	if err != nil {
		return fmt.Errorf("Invalid EndLowerApproximate flag, %v", err)
	}

	err = r.testPrecision(d.End.Upper.Approximate, r.options.EndUpperApproximate)

	if err != nil {
		return fmt.Errorf("Invalid EndUpperApproximate flag, %v", err)
	}

	return nil
}

func (r *TestResult) testPrecision(flags edtf.Precision, expected edtf.Precision) error {

	if expected == edtf.NONE {
		return nil
	}

	if !flags.HasFlag(expected) {
		return fmt.Errorf("Missing flag %v", expected)
	}

	return nil
}

func (r *TestResult) testRFC3339All(d *edtf.EDTFDate) error {

	if r.options.StartLowerTimeRFC3339 != "" {

		err := r.testRFC3339(r.options.StartLowerTimeRFC3339, d.Start.Lower.Timestamp)

		if err != nil {
			return fmt.Errorf("Failed StartLowerTimeRFC3339 test, %v", err)
		}
	}

	if r.options.StartUpperTimeRFC3339 != "" {

		err := r.testRFC3339(r.options.StartUpperTimeRFC3339, d.Start.Upper.Timestamp)

		if err != nil {
			return fmt.Errorf("Failed StartUpperTimeRFC3339 test, %v", err)
		}
	}

	if r.options.EndLowerTimeRFC3339 != "" {

		err := r.testRFC3339(r.options.EndLowerTimeRFC3339, d.End.Lower.Timestamp)

		if err != nil {
			return fmt.Errorf("Failed EndLowerTimeRFC3339 test, %v", err)
		}
	}

	if r.options.EndUpperTimeRFC3339 != "" {

		err := r.testRFC3339(r.options.EndUpperTimeRFC3339, d.End.Upper.Timestamp)

		if err != nil {
			return fmt.Errorf("Failed EndUpperTimeRFC3339 test, %v", err)
		}
	}

	return nil
}

func (r *TestResult) testRFC3339(expected string, ts *edtf.Timestamp) error {

	if ts == nil {
		return fmt.Errorf("Missing edtf.Timestamp instance")
	}

	t := ts.Time()

	t_str := t.Format(time.RFC3339)

	if t_str != expected {
		return fmt.Errorf("Invalid RFC3339 time, expected '%s' but got '%s'", expected, t_str)
	}

	return nil
}

func (r *TestResult) testUnixAll(d *edtf.EDTFDate) error {

	if r.options.StartLowerTimeUnix != 0 {

		err := r.testUnix(r.options.StartLowerTimeUnix, d.Start.Lower.Timestamp)

		if err != nil {
			return fmt.Errorf("Failed StartLowerTimeUnix test, %v", err)
		}
	}

	if r.options.StartUpperTimeUnix != 0 {

		err := r.testUnix(r.options.StartUpperTimeUnix, d.Start.Upper.Timestamp)

		if err != nil {
			return fmt.Errorf("Failed StartUpperTimeUnix test, %v", err)
		}
	}

	if r.options.EndLowerTimeUnix != 0 {

		err := r.testUnix(r.options.EndLowerTimeUnix, d.End.Lower.Timestamp)

		if err != nil {
			return fmt.Errorf("Failed EndLowerTimeUnix test, %v", err)
		}
	}

	if r.options.EndUpperTimeUnix != 0 {

		err := r.testUnix(r.options.EndUpperTimeUnix, d.End.Upper.Timestamp)

		if err != nil {
			return fmt.Errorf("Failed EndUpperTimeUnix test, %v", err)
		}
	}

	return nil
}

func (r *TestResult) testUnix(expected int64, ts *edtf.Timestamp) error {

	if ts == nil {
		return fmt.Errorf("Missing edtf.Timestamp instance")
	}

	ts_unix := ts.Unix()

	if ts_unix != expected {
		return fmt.Errorf("Invalid Unix time, expected '%d' but got '%d'", expected, ts_unix)
	}

	return nil
}
