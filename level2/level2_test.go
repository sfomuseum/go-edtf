package level2

import (
	"github.com/whosonfirst/go-edtf"
	"testing"
)

func TestIsLevel2(t *testing.T) {

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

		ok := IsLevel2(str)

		if !ok {
			t.Fatalf("Expected '%s' to parse as Level 2 string", str)
		}

		ParseString(str)
	}

	for _, str := range invalid {

		ok := IsLevel2(str)

		if ok {
			t.Fatalf("Expected '%s' to not parse as Level 2 string", str)
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

			if edtf.IsNotImplemented(err) || edtf.IsUnsupported(err) {
				t.Logf("Skipping '%s', %v", str, err)
			} else {
				t.Fatalf("Failed to parse '%s', %v", str, err)
			}
		}
	}
}
