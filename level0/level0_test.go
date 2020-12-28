package level0

import (
	"testing"
	"time"
)

func TestIsLevel0(t *testing.T) {

	valid := make([]string, 0)

	for _, candidates := range Tests {

		for _, str := range candidates {
			valid = append(valid, str)
		}
	}

	invalid := []string{
		"Dec. 1970",
		"Last week",
		"Y170000002",
		"c 1960",
	}

	for _, str := range valid {

		ok := IsLevel0(str)

		if !ok {
			t.Fatalf("Expected '%s' to parse as Level 0 string", str)
		}

		ParseString(str)
	}

	for _, str := range invalid {

		ok := IsLevel0(str)

		if ok {
			t.Fatalf("Expected '%s' to not parse as Level 0 string", str)
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

		d, err := ParseString(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}

		if d.String() != str {
			t.Fatalf("Failed to stringify '%s', %v", str, err)
		}

	}
}
