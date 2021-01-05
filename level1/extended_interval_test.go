package level1

import (
	"testing"
)

func TestExtendedIntervalEnd(t *testing.T) {

	valid, ok := Tests[EXTENDED_INTERVAL_END]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for input, _ := range valid {

		_, err := ParseExtendedIntervalEnd(input)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", input, err)
		}
	}

}

func TestExtendedIntervalStart(t *testing.T) {

	valid, ok := Tests[EXTENDED_INTERVAL_START]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for input, _ := range valid {

		_, err := ParseExtendedIntervalStart(input)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", input, err)
		}
	}

}
