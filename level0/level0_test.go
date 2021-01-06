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

		t.Logf("IsLevel0 '%s'", input)

		ok := IsLevel0(input)

		if !ok {
			t.Logf("Expected '%s' to parse as Level 0 string", input)
			t.Fail()
			continue
		}

		ParseString(input)
	}

	for _, input := range invalid {

		t.Logf("IsLevel0 (false) '%s'", input)

		ok := IsLevel0(input)

		if ok {
			t.Logf("Expected '%s' to not parse as Level 0 string", input)
			t.Fail()
			continue
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

		_, err := ParseString(input)

		if err != nil {
			t.Logf("Failed to parse '%s', %v", input, err)
			t.Fail()
			continue
		}
	}
}
