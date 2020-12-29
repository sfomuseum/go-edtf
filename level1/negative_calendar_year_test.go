package level1

import (
	"testing"
)

func TestNegativeCalendarYear(t *testing.T) {

	valid, ok := Tests["negative_calendar_year"]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseNegativeCalendarYear(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}

}
