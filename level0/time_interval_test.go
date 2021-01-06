package level0

import (
	"testing"
)

func TestParseTimeInterval(t *testing.T) {

	valid, ok := Tests[TIME_INTERVAL]

	if !ok {
		t.Fatal("Failed to load test strings")
	}

	for input, tr := range valid {

		t.Logf("ParseTimeInterval '%s'", input)

		d, err := ParseTimeInterval(input)

		if err != nil {
			t.Logf("Failed to parse '%s', %v", input, err)
			t.Fail()
			continue
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
