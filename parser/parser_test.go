package parser

import (
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/level0"
	"github.com/sfomuseum/go-edtf/level1"
	"github.com/sfomuseum/go-edtf/level2"
	"testing"
)

func TestLevels(t *testing.T) {

	for label, tests := range level0.Tests {

		for input, _ := range tests {

			if !level0.IsLevel0(input) {
				t.Fatalf("Invalid level 0 string '%s' (%s)", input, label)
			}
		}
	}

	for label, tests := range level1.Tests {

		for input, _ := range tests {

			if !level1.IsLevel1(input) {
				t.Fatalf("Invalid level 1 string '%s' (%s)", input, label)
			}
		}
	}

	for label, tests := range level2.Tests {

		for input, _ := range tests {

			if !level2.IsLevel2(input) {
				t.Fatalf("Invalid level 2 string '%s' (%s)", input, label)
			}
		}
	}

}

func TestIsValid(t *testing.T) {

	for label, tests := range Tests {

		for input, _ := range tests {

			if !IsValid(input) {
				t.Fatalf("Invalid string '%s' (%s)", input, label)
			}
		}
	}

	for label, tests := range level0.Tests {

		for input, _ := range tests {

			if !IsValid(input) {
				t.Fatalf("Invalid level 0 string '%s' (%s)", input, label)
			}
		}
	}

	for label, tests := range level1.Tests {

		for input, _ := range tests {

			if !IsValid(input) {
				t.Fatalf("Invalid level 1 string '%s' (%s)", input, label)
			}
		}
	}

	for label, tests := range level2.Tests {

		for input, _ := range tests {

			if !IsValid(input) {
				t.Fatalf("Invalid level 2 string '%s' (%s)", input, label)
			}
		}
	}

}

func TestParseString(t *testing.T) {

	for label, tests := range Tests {

		for input, _ := range tests {

			_, err := ParseString(input)

			if err != nil {

				if edtf.IsNotImplemented(err) || edtf.IsUnsupported(err) {
					t.Logf("Skipping level 0 string '%s' (%s), %v", input, label, err)
				} else {
					t.Fatalf("Failed to parse level 0 string '%s' (%s), %v", input, label, err)
				}
			}
		}
	}

	for label, tests := range level0.Tests {

		for input, _ := range tests {

			_, err := ParseString(input)

			if err != nil {

				if edtf.IsNotImplemented(err) || edtf.IsUnsupported(err) {
					t.Logf("Skipping level 0 string '%s' (%s), %v", input, label, err)
				} else {
					t.Fatalf("Failed to parse level 0 string '%s' (%s), %v", input, label, err)
				}
			}
		}
	}

	for label, tests := range level1.Tests {

		for input, _ := range tests {

			_, err := ParseString(input)

			if err != nil {

				if edtf.IsNotImplemented(err) || edtf.IsUnsupported(err) {
					t.Logf("Skipping level 1 string '%s' (%s), %v", input, label, err)
				} else {
					t.Fatalf("Failed to parse level 1 string '%s' (%s), %v", input, label, err)
				}
			}
		}
	}

	for label, tests := range level2.Tests {

		for input, _ := range tests {

			_, err := ParseString(input)

			if err != nil {

				if edtf.IsNotImplemented(err) || edtf.IsUnsupported(err) {
					t.Logf("Skipping level 2 string '%s' (%s), %v", input, label, err)
				} else {
					t.Fatalf("Failed to parse level 2 string '%s' (%s), %v", input, label, err)
				}
			}
		}
	}

}

func TestIsOpen(t *testing.T) {

	is_open := []string{
		edtf.OPEN,
		edtf.OPEN_2012,
	}

	not_open := []string{
		"",
		"2021-09-28",
	}

	for _, s := range is_open {

		if !IsOpen(s) {
			t.Fatalf("String '%s' is considered open but reported as not open", s)
		}
	}

	for _, s := range not_open {

		if IsOpen(s) {
			t.Fatalf("String '%s' is not open but reported as being open", s)
		}
	}

}

func TestIsUnspecified(t *testing.T) {

	is_unspecified := []string{
		edtf.UNSPECIFIED,
		edtf.UNSPECIFIED_2012,
	}

	not_unspecified := []string{
		"2021-09-28",
	}

	for _, s := range is_unspecified {

		if !IsUnspecified(s) {
			t.Fatalf("String '%s' is considered unspecified but reported as not unspecified", s)
		}
	}

	for _, s := range not_unspecified {

		if IsUnspecified(s) {
			t.Fatalf("String '%s' is not unspecified but reported as being unspecified", s)
		}
	}

}

func TestIsUnknown(t *testing.T) {

	is_unknown := []string{
		edtf.UNKNOWN,
		edtf.UNKNOWN_2012,
	}

	not_unknown := []string{
		"2021-09-28",
	}

	for _, s := range is_unknown {

		if !IsUnknown(s) {
			t.Fatalf("String '%s' is considered unknown but reported as not unknown", s)
		}
	}

	for _, s := range not_unknown {

		if IsUnknown(s) {
			t.Fatalf("String '%s' is not unknown but reported as being unknown", s)
		}
	}

}
