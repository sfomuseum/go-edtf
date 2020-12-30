package level2

import (
	"github.com/whosonfirst/go-edtf"
	"testing"
)

func TestUnspecifiedDigit(t *testing.T) {

	valid, ok := Tests[UNSPECIFIED_DIGIT]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseUnspecifiedDigit(str)

		if err != nil {
			if edtf.IsNotImplemented(err) {
				t.Logf("Failed to parse '%s', %v", str, err)
			} else {
				t.Fatalf("Failed to parse '%s', %v", str, err)
			}
		}
	}
}
