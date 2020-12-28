package level1

import (
	"testing"
)

func TestIsLevel1(t *testing.T) {

	valid := make([]string, 0)

	for _, candidates := range Tests {

		for _, str := range candidates {
			valid = append(valid, str)
		}
	}

	invalid := []string{
		"Dec. 1970",
		"Last week",
		// "Y170000002",
		"c 1960",
	}

	for _, str := range valid {

		ok := IsLevel1(str)

		if !ok {
			t.Fatalf("Expected '%s' to parse as Level 1 string", str)
		}

		ParseString(str)
	}

	for _, str := range invalid {

		ok := IsLevel1(str)

		if ok {
			t.Fatalf("Expected '%s' to not parse as Level 1 string", str)
		}
	}

}

func TestParseString(t *testing.T) {

	valid := make([]string, 0)

	for _, candidates := range Tests {

		for _, str := range candidates {
			valid = append(valid, str)
		}
	}

	for _, str := range valid {

		_, err := ParseString(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}

		/*
			if d.String() != str {
				t.Fatalf("Failed to stringify '%s', %v", str, err)
			}
		*/
	}
}

func TestExtendedIntervalEnd(t *testing.T) {

	valid, ok := Tests["extended_interval_end"]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseExtendedIntervalEnd(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}

}

func TestExtendedIntervalStart(t *testing.T) {

	valid, ok := Tests["extended_interval_start"]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseExtendedIntervalStart(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}

}

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
