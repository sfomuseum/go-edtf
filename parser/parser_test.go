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
