package level1

import (
	"testing"
)

func TestUnspecifiedDigits(t *testing.T) {

	valid, ok := Tests[UNSPECIFIED_DIGITS]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseUnspecifiedDigits(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}

}
