package level1

import (
	"testing"
)

func TestLetterPrefixedCalendarYear(t *testing.T) {

	valid, ok := Tests[LETTER_PREFIXED_CALENDAR_YEAR]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseLetterPrefixedCalendarYear(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}
}
