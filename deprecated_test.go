package edtf

import (
	"testing"
)

func TestIsDeprecated(t *testing.T) {

	is_deprecated := []string{
		OPEN_2012,
		UNSPECIFIED_2012,
	}

	not_deprecated := []string{
		OPEN,
		UNSPECIFIED,
	}

	for _, s := range is_deprecated {

		if !IsDeprecated(s) {
			t.Fatalf("%s is expected to be deprecated", s)
		}
	}

	for _, s := range not_deprecated {

		if IsDeprecated(s) {
			t.Fatalf("%s is expected not deprecated", s)
		}
	}

}

func TestReplaceDeprecated(t *testing.T) {

	for old, expected := range deprecated {

		new, err := ReplaceDeprecated(old)

		if err != nil {
			t.Fatalf("Failed to replace deprecated string (%s), %v", old, err)
		}

		if new != expected {
			t.Fatalf("Unexpected replacement for %s. Expected '%s' but got '%s'", old, expected, new)
		}
	}

}
