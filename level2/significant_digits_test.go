package level2

import (
	"github.com/sfomuseum/go-edtf"
	"testing"
)

func TestSignificantDigits(t *testing.T) {

	valid, ok := Tests[SIGNIFICANT_DIGITS]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for input, tr := range valid {

		t.Logf("ParseSignificantDigits '%s'", input)

		d, err := ParseSignificantDigits(input)

		if err != nil {

			if edtf.IsNotImplemented(err) || edtf.IsUnsupported(err) {
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
