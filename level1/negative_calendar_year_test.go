package level1

import (
	"testing"
)

func TestNegativeCalendarYear(t *testing.T) {

	valid, ok := Tests[NEGATIVE_CALENDAR_YEAR]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for input, tr := range valid {

		t.Logf("ParseNegativeCalendarYear '%s'", input)

		d, err := ParseNegativeCalendarYear(input)

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
