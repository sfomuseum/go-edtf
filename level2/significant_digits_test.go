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

	for _, str := range valid {

		_, err := ParseSignificantDigits(str)

		if err != nil {

			if edtf.IsNotImplemented(err) || edtf.IsUnsupported(err) {
				t.Logf("Skipping '%s', %v", str, err)
			} else {
				t.Fatalf("Failed to parse '%s', %v", str, err)
			}
		}
	}
}
