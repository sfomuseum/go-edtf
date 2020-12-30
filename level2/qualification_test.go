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

	for _, str := range valid {

		_, err := ParseGroupQualification(str)

		if err != nil {
			if edtf.IsNotImplemented(err) {
				t.Logf("Failed to parse '%s', %v", str, err)
			} else {
				t.Fatalf("Failed to parse '%s', %v", str, err)
			}
		}
	}
}

func TestIndividualQualification(t *testing.T) {

	valid, ok := Tests[INDIVIDUAL_QUALIFICATION]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseIndividualQualification(str)

		if err != nil {
			if edtf.IsNotImplemented(err) {
				t.Logf("Failed to parse '%s', %v", str, err)
			} else {
				t.Fatalf("Failed to parse '%s', %v", str, err)
			}
		}
	}
}
