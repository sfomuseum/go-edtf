package level2

import (
	"github.com/whosonfirst/go-edtf"
	"testing"
)

func TestGroupQualification(t *testing.T) {

	valid, ok := Tests[GROUP_QUALIFICATION]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for input, tr := range valid {

		d, err := ParseGroupQualification(input)

		if err != nil {
			if edtf.IsNotImplemented(err) {
				t.Logf("Skipping '%s', %v", input, err)
				continue
			} else {
				t.Fatalf("Failed to parse '%s', %v", input, err)
			}
		}

		if tr != nil {

			err := tr.TestDate(d)

			if err != nil {
				t.Fatalf("Results failed tests '%s', %v", input, err)
			}
		}
	}
}

func TestIndividualQualification(t *testing.T) {

	valid, ok := Tests[INDIVIDUAL_QUALIFICATION]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for input, tr := range valid {

		d, err := ParseIndividualQualification(input)

		if err != nil {
			if edtf.IsNotImplemented(err) {
				t.Logf("Skipping '%s', %v", input, err)
				continue
			} else {
				t.Fatalf("Failed to parse '%s', %v", input, err)
			}
		}

		if tr != nil {

			err := tr.TestDate(d)

			if err != nil {
				t.Fatalf("Results failed tests '%s', %v", input, err)
			}
		}

	}
}
