package level1

import (
	"github.com/sfomuseum/go-edtf"
	"testing"
)

func TestLetterPrefixedCalendarYear(t *testing.T) {

	valid, ok := Tests[LETTER_PREFIXED_CALENDAR_YEAR]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for input, tr := range valid {

		t.Logf("ParseLetterPrefixedCalendarYear '%s'", input)

		d, err := ParseLetterPrefixedCalendarYear(input)

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
