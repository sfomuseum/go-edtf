package level1

import (
	"testing"
)

var test_strings map[string][]string

func init() {

	test_strings = map[string][]string{
		"prefixed_calendar_year": []string{
			"Y170000002",
			"Y-17000002",
		},
		"season": []string{
			"2001-01",
			"2001-24",
			"Spring, 2002",
			"winter, 2002",
		},
		"qualified_date": []string{
			"1984?",
			"2004-06~",
			"2004-06-11%",
		},
		"unspecified_digits": []string{
			"201X",
			"20XX",
			"2004-XX",
			"1985-04-XX",
			"1985-XX-XX",
		},
		"extended_interval_start": []string{
			"../1985-04-12",
			"../1985-04",
			"../1985",
			"/1985-04-12",
			"/1985-04",
			"/1985",
		},
		"extended_interval_end": []string{
			"1985-04-12/..",
			"1985-04/..",
			"1985/..",
			"1985-04-12/",
			"1985-04/",
			"1985/",
		},
		"negative_calendar_year": []string{
			"-1985",
		},
	}
}

func TestIsLevel1(t *testing.T) {

	valid := make([]string, 0)

	for _, candidates := range test_strings {

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

	for _, candidates := range test_strings {

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

func TestLetterPrefixedCalendarYear(t *testing.T) {

	valid, ok := test_strings["prefixed_calendar_year"]

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

func TestSeason(t *testing.T) {

	valid, ok := test_strings["season"]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseSeason(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}

}

func TestQualifiedDate(t *testing.T) {

	valid, ok := test_strings["qualified_date"]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseQualifiedDate(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}

}

func TestUnspecifiedDigits(t *testing.T) {

	valid, ok := test_strings["unspecified_digits"]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseUnspecifiedDigits(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}

}

func TestExtendedIntervalEnd(t *testing.T) {

	valid, ok := test_strings["extended_interval_end"]

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

	valid, ok := test_strings["extended_interval_start"]

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

	valid, ok := test_strings["negative_calendar_year"]

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
