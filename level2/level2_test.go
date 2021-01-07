package level2

import (
	"github.com/sfomuseum/go-edtf"
	"testing"
)

func TestIsLevel2(t *testing.T) {

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

		ok := IsLevel2(input)

		if !ok {
			t.Fatalf("Expected '%s' to parse as Level 2 string", input)
		}

		ParseString(input)
	}

	for _, input := range invalid {

		ok := IsLevel2(input)

		if ok {
			t.Fatalf("Expected '%s' to not parse as Level 2 string", input)
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

			if edtf.IsNotImplemented(err) || edtf.IsUnsupported(err) {
				t.Logf("Skipping '%s', %v", input, err)
			} else {
				t.Fatalf("Failed to parse '%s', %v", input, err)
			}
		}
	}
}
