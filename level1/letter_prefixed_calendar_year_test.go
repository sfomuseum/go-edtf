package level1

import (
	"github.com/whosonfirst/go-edtf"
	"testing"
)

func TestLetterPrefixedCalendarYear(t *testing.T) {

	valid, ok := Tests[LETTER_PREFIXED_CALENDAR_YEAR]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for input, tr := range valid {

		d, err := ParseLetterPrefixedCalendarYear(input)

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
