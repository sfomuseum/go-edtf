package level2

import (
	"github.com/whosonfirst/go-edtf"
	"testing"
)

func TestUnspecifiedDigit(t *testing.T) {

	valid, ok := Tests[UNSPECIFIED_DIGIT]

	if !ok {
		t.Logf("Failed to load test strings")
	}

	for input, tr := range valid {

		t.Logf("ParseUnspecifiedDigit '%s'", input)

		d, err := ParseUnspecifiedDigit(input)

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
