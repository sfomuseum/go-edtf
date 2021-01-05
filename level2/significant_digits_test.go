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

	for input, _ := range valid {

		_, err := ParseSignificantDigits(input)

		if err != nil {

			if edtf.IsNotImplemented(err) || edtf.IsUnsupported(err) {
				t.Logf("Skipping '%s', %v", input, err)
			} else {
				t.Fatalf("Failed to parse '%s', %v", input, err)
			}
		}
	}
}
