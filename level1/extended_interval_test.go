package level1

import (
	"testing"
)

func TestExtendedIntervalEnd(t *testing.T) {

	valid, ok := Tests[EXTENDED_INTERVAL_END]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseExtendedIntervalEnd(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}

}

func TestExtendedIntervalStart(t *testing.T) {

	valid, ok := Tests[EXTENDED_INTERVAL_START]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseExtendedIntervalStart(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}

}
