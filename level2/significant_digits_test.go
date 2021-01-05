package level2

import (
	"github.com/whosonfirst/go-edtf"
	"testing"
)

func TestSignificantDigits(t *testing.T) {

	valid, ok := Tests[SIGNIFICANT_DIGITS]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for input, tr := range valid {

		d, err := ParseSignificantDigits(input)

		if err != nil {

			if edtf.IsNotImplemented(err) || edtf.IsUnsupported(err) {
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
