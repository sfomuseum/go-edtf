package level1

import (
	"testing"
)

func TestSeason(t *testing.T) {

	valid, ok := Tests[SEASON]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for input, _ := range valid {

		_, err := ParseSeason(input)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", input, err)
		}
	}

}
