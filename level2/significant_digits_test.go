package level2

import (
	"testing"
)

func TestSignificantDigits(t *testing.T) {

	valid, ok := Tests[SIGNIFICANT_DIGITS]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseSignificantDigits(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}
}
