package level0

import (
	"testing"
)

func TestIsLevel0(t *testing.T) {

	valid := make([]string, 0)

	for _, candidates := range Tests {

		for input, _ := range candidates {
			valid = append(valid, input)
		}
	}

	invalid := []string{
		"Dec. 1970",
		"Last week",
		"Y170000002",
		"c 1960",
	}

	for _, input := range valid {

		ok := IsLevel0(input)

		if !ok {
			t.Fatalf("Expected '%s' to parse as Level 0 string", input)
		}

		ParseString(input)
	}

	for _, input := range invalid {

		ok := IsLevel0(input)

		if ok {
			t.Fatalf("Expected '%s' to not parse as Level 0 string", input)
		}
	}

}

func TestParseString(t *testing.T) {

	valid := make([]string, 0)

	for _, candidates := range Tests {

		for input, _ := range candidates {
			valid = append(valid, input)
		}
	}

	for _, input := range valid {

		d, err := ParseString(input)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", input, err)
		}

		if d.String() != input {
			t.Fatalf("Failed to stringify '%s', %v", input, err)
		}

	}
}
