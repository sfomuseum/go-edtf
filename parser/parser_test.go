package parser

import (
	"github.com/whosonfirst/go-edtf/level0"
	"github.com/whosonfirst/go-edtf/level1"
	"github.com/whosonfirst/go-edtf/level2"
	"testing"
)

func TestIsValid(t *testing.T) {

	for label, tests := range level0.Tests {

		for _, str := range tests {

			if !IsValid(str) {
				t.Fatalf("Invalid level 0 string '%s' (%s)", str, label)
			}
		}
	}

	for label, tests := range level1.Tests {

		for _, str := range tests {

			if !IsValid(str) {
				t.Fatalf("Invalid level 1 string '%s' (%s)", str, label)
			}
		}
	}

	for label, tests := range level2.Tests {

		for _, str := range tests {

			if !IsValid(str) {
				t.Fatalf("Invalid level 2 string '%s' (%s)", str, label)
			}
		}
	}

}

func TestParseString(t *testing.T) {

	for label, tests := range level0.Tests {

		for _, str := range tests {

			_, err := ParseString(str)

			if err != nil {
				t.Fatalf("Invalid level 0 string '%s' (%s), %v", str, label, err)
			}
		}
	}

	for label, tests := range level1.Tests {

		for _, str := range tests {

			_, err := ParseString(str)

			if err != nil {
				t.Fatalf("Invalid level 1 string '%s' (%s), %v", str, label, err)
			}
		}
	}

	for label, tests := range level2.Tests {

		for _, str := range tests {

			_, err := ParseString(str)

			if err != nil {
				t.Fatalf("Invalid level 2 string '%s' (%s), %v", str, label, err)
			}
		}
	}

}
