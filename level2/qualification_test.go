package level2

import (
	"github.com/whosonfirst/go-edtf"
	"testing"
)

func TestQualification(t *testing.T) {
	TestGroupQualification(t)
	TestIndividualQualification(t)
}

func TestGroupQualification(t *testing.T) {

	valid, ok := Tests[GROUP_QUALIFICATION]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for input, tr := range valid {

		t.Logf("ParseGroupQualification '%s'", input)

		d, err := ParseGroupQualification(input)

		if err != nil {
			if edtf.IsNotImplemented(err) {
				t.Logf("Skipping '%s', %v", input, err)
				continue
			} else {
				t.Logf("Failed to parse '%s', %v", input, err)
				t.Fail()
				continue
			}
		}

		if tr != nil {

			err := tr.TestDate(d)

			if err != nil {
				t.Logf("Results failed tests '%s', %v", input, err)
				t.Fail()
				continue
			}
		}
	}
}

func TestIndividualQualification(t *testing.T) {

	valid, ok := Tests[INDIVIDUAL_QUALIFICATION]

	if !ok {
		t.Logf("Failed to load test strings")
	}

	for input, tr := range valid {

		t.Logf("ParseIndividualQualification '%s'", input)

		d, err := ParseIndividualQualification(input)

		if err != nil {
			if edtf.IsNotImplemented(err) {
				t.Logf("Skipping '%s', %v", input, err)
				continue
			} else {
				t.Logf("Failed to parse '%s', %v", input, err)
				t.Fail()
				continue
			}
		}

		if tr != nil {

			err := tr.TestDate(d)

			if err != nil {
				t.Logf("Results failed tests '%s', %v", input, err)
				t.Fail()
				continue
			}
		}

	}
}
