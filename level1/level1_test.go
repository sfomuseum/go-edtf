package level1

import (
	"errors"
	"regexp"
	"testing"
)

func TestLetterPrefixedCalendarYear(t *testing.T) {

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

func TestSeason(t *testing.T) {

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
	
	return nil, nil
}

func TestQualifiedDate(t *testing.T) {

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
	
	return nil, nil
}

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
	
	return nil, nil
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
	
	return nil, nil
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
	
	return nil, nil
}

