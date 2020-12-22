package level1

import (
	// "errors"
	// "regexp"
	"testing"
)

func TestLetterPrefixedCalendarYear(t *testing.T) {

	valid := []string{
		"Y170000002",
		"Y-17000002",
	}

	for _, str := range valid {

		_, err := ParseLetterPrefixedCalendarYear(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}
}

func TestSeason(t *testing.T) {

	valid := []string{
		"2001-01",
		"2001-24",
		"Spring, 2002",
		"winter, 2002",
	}

	for _, str := range valid {

		_, err := ParseSeason(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}

}

func TestQualifiedDate(t *testing.T) {

	valid := []string{
		"1984?",
		"2004-06~",
		"2004-06-11%",
	}

	for _, str := range valid {

		_, err := ParseQualifiedDate(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}

}

/*
func TestUnspecifiedDigits(t *testing.T) {

	valid := []string{
		"Y170000002",
		"Y-170000002",
	}

	for _, str := range valid {

		_, err := ParseLetterPrefixedCalendarYear(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}

}

func TestExtendedInterval(t *testing.T) {

	valid := []string{
		"Y170000002",
		"Y-170000002",
	}

	for _, str := range valid {

		_, err := ParseLetterPrefixedCalendarYear(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}

}

func TestNegativeCalendarYear(t *testing.T) {

	valid := []string{
		"Y170000002",
		"Y-170000002",
	}

	for _, str := range valid {

		_, err := ParseLetterPrefixedCalendarYear(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}

}
*/
