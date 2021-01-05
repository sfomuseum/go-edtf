package level1

import (
	"testing"
)

func TestUnspecifiedDigits(t *testing.T) {

	valid, ok := Tests[UNSPECIFIED_DIGITS]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for input, tr := range valid {

		d, err := ParseUnspecifiedDigits(input)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", input, err)
		}

		if tr != nil {

			err := tr.TestDate(d)

			if err != nil {
				t.Fatalf("Results failed tests '%s', %v", input, err)
			}
		}

	}

}
